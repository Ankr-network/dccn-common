package pgrpc

import (
	context "context"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type Client struct {
	sync.Map
}

var DefaultClient *Client

func InitClient(network, addr string, onAccept func(*net.Conn) string, opts ...grpc.DialOption) error {
	var err error
	DefaultClient, err = NewClient(network, addr, onAccept, opts...)
	return err
}

func NewClient(network, addr string, onAccept func(*net.Conn) string, opts ...grpc.DialOption) (*Client, error) {
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

			var key string
			if onAccept != nil {
				key = onAccept(&conn)
			}
			if key == "" {
				key, _, _ = net.SplitHostPort(conn.RemoteAddr().String())
			}

			if val, ok := c.LoadOrStore(key, &pool{key: key, conns: []net.Conn{conn}, opts: opts}); ok {
				val.(*pool).PutConn(conn)
			}
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

	cc, err := val.(*pool).Get()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func Each(fn func(key string, cc *grpc.ClientConn) error) {
	DefaultClient.Each(fn)
}
func (c *Client) Each(fn func(key string, cc *grpc.ClientConn) error) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	c.Range(func(key, val interface{}) bool {
		wg.Add(1)
		go func(key string, pool *pool) {
			defer wg.Done()

			// avoid send by alias pool in loop
			if pool.key != key {
				return
			}

			cc, err := pool.Get()
			// no available connection
			if err != nil {
				return
			}

			if err := fn(key, cc); err != nil {
				cc.Close()
				return
			}
			pool.PutCC(cc)

		}(key.(string), val.(*pool))
		return true
	})
}

func Alias(key, alias string) {
	DefaultClient.Alias(key, alias)
}
func (c *Client) Alias(key, alias string) {
	if val, ok := c.Load(key); ok {
		c.Store(alias, val)
	} else {
		c.Delete(alias)
	}
}

// pool maintain idle connections
const MAX_IDLE = 2

type pool struct {
	key   string
	conns []net.Conn
	ccs   []*grpc.ClientConn
	opts  []grpc.DialOption

	mu sync.Mutex
}

func (s *pool) Get() (*grpc.ClientConn, error) {
	var bans []*grpc.ClientConn
	defer func() {
		for i := range bans {
			bans[i].Close()
		}
	}()

	var pick *grpc.ClientConn
	var pickIdx = -1
	var retry = true

RETRY:
	s.mu.Lock()
	// find avaiable ClientConn
FOR:
	for i := range s.ccs {
		switch s.ccs[i].GetState() {
		case connectivity.Connecting:
			fallthrough
		case connectivity.TransientFailure:
			pickIdx = i

		case connectivity.Ready:
			fallthrough
		case connectivity.Idle:
			pickIdx = i
			break FOR

		default:
			bans = append(bans, s.ccs[i])
			if i+1 == len(s.ccs) {
				s.ccs = s.ccs[:i]
			} else {
				s.ccs = append(s.ccs[:i], s.ccs[i+1:]...)
			}
			goto FOR
		}
	}

	if pickIdx != -1 {
		pick = s.ccs[pickIdx]
		if pickIdx+1 == len(s.ccs) {
			s.ccs = s.ccs[:pickIdx]
		} else {
			s.ccs = append(s.ccs[:pickIdx], s.ccs[pickIdx+1:]...)
		}
	}

	if pick != nil {
		s.mu.Unlock()
		return pick, nil
	}

	// no avaiable ClientConn, try build from net.Conn
	if len(s.conns) == 0 {
		if retry {
			retry = false
			time.Sleep(200 * time.Millisecond)
			s.mu.Unlock()
			goto RETRY
		}

		s.mu.Unlock()
		return nil, errors.New("no available connection to " + s.key)
	}

	conn := s.conns[0]
	s.conns = s.conns[1:]
	s.mu.Unlock()

	// dial client conn
	opts := append(s.opts, grpc.WithContextDialer(
		func(context.Context, string) (net.Conn, error) { return conn, nil }))
	cc, err := grpc.DialContext(context.Background(), conn.RemoteAddr().String(), opts...)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return cc, nil
}

func (s *pool) PutCC(cc *grpc.ClientConn) {
	var dropped *grpc.ClientConn
	switch cc.GetState() {
	case connectivity.Ready:
		fallthrough
	case connectivity.Idle:
		s.mu.Lock()
		if len(s.ccs) == MAX_IDLE {
			dropped = s.ccs[0]
		}
		s.ccs = append(s.ccs, cc)
		s.mu.Unlock()

	case connectivity.Connecting:
		fallthrough
	case connectivity.TransientFailure:
		s.mu.Lock()
		if len(s.ccs) == MAX_IDLE {
			dropped = cc
		} else {
			s.ccs = append(s.ccs, cc)
		}
		s.mu.Unlock()

	default:
		dropped = cc
	}

	if dropped != nil {
		dropped.Close()
	}
}

func (s *pool) PutConn(conn net.Conn) {
	s.mu.Lock()
	s.conns = append(s.conns, conn)
	s.mu.Unlock()
}
