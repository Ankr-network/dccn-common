package pgrpc

import (
	context "context"
	"io"
	"net"
	"sync"

	"github.com/pkg/errors"
	grpc "google.golang.org/grpc"
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
				val.(*pool).Put(nil, conn)
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

	cc := val.(*pool).Get()
	if cc == nil {
		return nil, errors.Errorf("no available connection point to %s", key)
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

			cc := pool.Get()
			// no available connection
			if cc == nil {
				return
			}

			if err := fn(key, cc); err != nil {
				cc.Close()
				return
			}
			pool.Put(cc, nil)

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
const MAX_IDLE_CONN = 2

type pool struct {
	key   string
	conns []net.Conn
	ccs   []*grpc.ClientConn
	opts  []grpc.DialOption

	mu sync.Mutex
}

func (s *pool) Get() *grpc.ClientConn {
	s.mu.Lock()
	if length := len(s.ccs); length != 0 {
		// pop the last one
		cc := s.ccs[length-1]
		s.ccs = s.ccs[:length-1]
		s.mu.Unlock()
		return cc
	}

	length := len(s.conns)
	if length == 0 {
		s.mu.Unlock()
		return nil
	}

	conn := s.conns[length-1]
	s.conns = s.conns[:length-1]
	s.mu.Unlock()

	// dial client conn
	opts := append(s.opts, grpc.WithContextDialer(
		func(context.Context, string) (net.Conn, error) { return conn, nil }))
	cc, err := grpc.DialContext(context.Background(), conn.RemoteAddr().String(), opts...)
	if err != nil {
		panic(err)
	}
	return cc
}

func (s *pool) Put(cc *grpc.ClientConn, conn net.Conn) {
	if cc != nil && conn != nil {
		panic("just support put back only one connection")
	}

	var dropped io.Closer
	s.mu.Lock()

	if count := len(s.ccs) + len(s.conns); count == MAX_IDLE_CONN {
		if len(s.ccs) != 0 { // drop client conn first
			dropped = s.ccs[0]
			s.ccs = s.ccs[1:]
		} else {
			dropped = s.conns[0]
			s.conns = s.conns[1:]
		}
	}

	if cc != nil {
		s.ccs = append(s.ccs, cc)
	}
	if conn != nil {
		s.conns = append(s.conns, conn)
	}

	s.mu.Unlock()
	if dropped != nil {
		dropped.Close()
	}
}
