package pgrpc

import (
	context "context"
	"net"
	"sync"

	"github.com/pkg/errors"
	grpc "google.golang.org/grpc"
)

const MAX_IDLE = 2

type Client struct {
	sync.Map
}

var DefaultClient *Client

func InitClient(network, addr string, onAccept func(*net.Conn, error), opts ...grpc.DialOption) error {
	var err error
	DefaultClient, err = NewClient(network, addr, onAccept, opts...)
	return err
}

func NewClient(network, addr string, onAccept func(*net.Conn, error), opts ...grpc.DialOption) (*Client, error) {
	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	var c = &Client{}
	go func() {
		for {
			conn, err := ln.Accept()
			if onAccept != nil {
				onAccept(&conn, err)
			}
			if err != nil {
				continue
			}

			ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())

			opts = append(opts, grpc.WithContextDialer(
				func(context.Context, string) (net.Conn, error) { return conn, nil }))
			cc, err := grpc.DialContext(context.Background(), ip, opts...)
			if err != nil {
				panic(err)
			}

			if val, ok := c.LoadOrStore(ip, &stack{ccs: []*grpc.ClientConn{cc}}); ok {
				val.(*stack).Put(cc)
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

	cc := val.(*stack).Get()
	if cc == nil {
		return nil, errors.Errorf("no available connection point to %s", key)
	}
	return cc, nil
}

func Each(fn func(key string, cc *grpc.ClientConn)) {
	DefaultClient.Each(fn)
}
func (c *Client) Each(fn func(key string, cc *grpc.ClientConn)) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	c.Range(func(key, val interface{}) bool {
		wg.Add(1)
		go func(key string, stack *stack) {
			cc := stack.Get()
			if cc == nil {
				return
			}

			defer func() {
				wg.Done()
				stack.Put(cc)
			}()

			if cc.Target() != key {
				return
			}

			fn(key, cc)
		}(key.(string), val.(*stack))
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

// stack maintain idle connections
type stack struct {
	ccs []*grpc.ClientConn
	mu  sync.Mutex
}

func (s *stack) Get() (cc *grpc.ClientConn) {
	s.mu.Lock()
	if length := len(s.ccs); length != 0 {
		cc = s.ccs[length-1]
		s.ccs = s.ccs[:length-1]
	}
	s.mu.Unlock()
	return
}
func (s *stack) Put(cc *grpc.ClientConn) {
	s.mu.Lock()
	if len(s.ccs) != MAX_IDLE {
		s.ccs = append(s.ccs, cc)
		s.mu.Unlock()
		return
	}
	s.mu.Unlock()

	cc.Close()
}
