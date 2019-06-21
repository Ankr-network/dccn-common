package pgrpc

import (
	context "context"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Client struct {
	sync.Map
}

var DefaultClient *Client

func InitClient(network, addr string, ipHook func(*net.Conn) string, opts ...grpc.DialOption) error {
	var err error
	DefaultClient, err = NewClient(network, addr, ipHook, opts...)
	return err
}

func NewClient(network, addr string, ipHook func(*net.Conn) string, opts ...grpc.DialOption) (*Client, error) {
	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	var c = &Client{}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}

			go func(conn net.Conn) {
				if tcpConn, ok := conn.(*net.TCPConn); ok {
					tcpConn.SetLinger(1)
					tcpConn.SetKeepAlive(true)
					tcpConn.SetKeepAlivePeriod(5 * time.Second)
				}

				var ip string
				if ipHook != nil {
					ip = ipHook(&conn)
				}
				if ip == "" {
					ip, _, _ = net.SplitHostPort(conn.RemoteAddr().String())
				}

				buf := make([]byte, MAX_ID_LEN)
				conn.SetDeadline(time.Now().Add(5 * time.Second))
				if _, err := io.ReadFull(conn, buf); err != nil {
					conn.Close()
					return
				}

				conn.SetDeadline(time.Time{})
				id := strings.TrimSpace(string(buf))

				if val, ok := c.LoadOrStore(id, &pool{id: id,
					ips: []string{ip}, conns: []net.Conn{conn}, opts: opts}); ok {
					val.(*pool).PutConn(ip, conn)
				}
			}(conn)
		}
	}()
	return c, nil
}

func Dial(key string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return DefaultClient.Dial(key)
}

func (c *Client) Dial(key string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	val, ok := c.Load(key)
	if !ok {
		return nil, errors.Errorf("connection point to %s not found", key)
	}

	cc, _, err := val.(*pool).Get()
	return cc, errors.Wrap(err, "pgrpc dial")
}

func Each(fn func(key string, ips []string, cc *grpc.ClientConn, err error) error) {
	DefaultClient.Each(fn)
}
func (c *Client) Each(fn func(key string, ips []string, cc *grpc.ClientConn, err error) error) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	c.Range(func(key, val interface{}) bool {
		wg.Add(1)
		go func(key string, pool *pool) {
			defer wg.Done()

			cc, ips, err := pool.Get()
			if e := fn(key, ips, cc, err); e != nil {
				if err == nil {
					cc.Close()
				}
				return
			}
			pool.PutCC(cc)

		}(key.(string), val.(*pool))
		return true
	})
}

func PutCC(cc *grpc.ClientConn) {
	DefaultClient.PutCC(cc)
}
func (c *Client) PutCC(cc *grpc.ClientConn) {
	val, ok := c.Load(cc.Target())
	if !ok {
		return
	}

	pool := val.(*pool)
	pool.PutCC(cc)
}

// pool maintain idle connections
type pool struct {
	id    string
	ips   []string
	conns []net.Conn
	ccs   []*grpc.ClientConn
	opts  []grpc.DialOption

	mu sync.Mutex
}

func (s *pool) Get() (*grpc.ClientConn, []string, error) {
	var conn net.Conn
	var ips []string
	for i := 0; i < 10; i++ {
		s.mu.Lock()
		if len(s.ccs) != 0 {
			pick := s.ccs[0]
			s.ccs = s.ccs[1:]
			ips = s.ips
			s.mu.Unlock()
			return pick, ips, nil
		}

		// no avaiable ClientConn, try build from net.Conn
		if len(s.conns) == 0 {
			s.mu.Unlock()
			time.Sleep(200 * time.Millisecond)
			continue
		}

		conn = s.conns[0]
		s.conns = s.conns[1:]
		ips = s.ips
		s.mu.Unlock()
		break
	}

	if conn == nil {
		return nil, nil, errors.New("no available connection to " + s.id)
	}

	// dial client conn
	opts := append(s.opts, grpc.WithContextDialer(
		func(context.Context, string) (net.Conn, error) { return conn, nil }))
	cc, err := grpc.DialContext(context.Background(), s.id, opts...)
	if err != nil {
		conn.Close()
		return nil, nil, err
	}
	return cc, ips, nil
}

func (s *pool) PutCC(cc *grpc.ClientConn) {
	s.mu.Lock()
	if len(s.ccs) == (MAX_IDLE - MIN_IDLE) {
		cc.Close()
	} else {
		s.ccs = append(s.ccs, cc)
	}
	s.mu.Unlock()
}

func (s *pool) PutConn(ip string, conn net.Conn) {
	s.mu.Lock()
	ipInPool := false
	for i := range s.ips {
		if s.ips[i] == ip {
			ipInPool = true
		}
	}
	if !ipInPool {
		s.ips = append(s.ips, ip)
	}

	s.conns = append(s.conns, conn)
	s.mu.Unlock()
}
