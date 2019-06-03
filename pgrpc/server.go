package pgrpc

import (
	"net"
	"time"

	"github.com/pkg/errors"
)

type listener struct {
	addr net.Addr

	connCh chan net.Conn
	stopCh chan struct{}
}

func Listen(network, address string, onAccept func(*net.Conn, error)) (net.Listener, error) {
	var (
		ln = listener{
			connCh: make(chan net.Conn),
			stopCh: make(chan struct{}),
		}
		err error
	)

	if ln.addr, err = net.ResolveTCPAddr(network, address); err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := net.Dial(network, address)
			if onAccept != nil {
				onAccept(&conn, err)
			}
			if err != nil {
				time.Sleep(5 * time.Second)
				continue
			} else {
				c := conn.(*net.TCPConn)
				c.SetKeepAlive(true)
				c.SetKeepAlivePeriod(5 * time.Second)
				c.SetLinger(1)
			}

			conn, sig := newActiveConn(conn)
			select {
			case <-sig:
				ln.connCh <- conn
			case <-ln.stopCh:
				return
			}
		}
	}()
	return &ln, nil
}

func (ln *listener) Accept() (net.Conn, error) {
	select {
	case conn := <-ln.connCh:
		return conn, nil
	case <-ln.stopCh:
		return nil, errors.New("listener has been stoped")
	}
}
func (ln *listener) Close() error {
	close(ln.stopCh)
	return nil
}
func (ln *listener) Addr() net.Addr {
	return ln.addr
}

// activeConn is a net.Conn that can detect conn first activaty
type activeConn struct {
	b    byte
	n    int
	err  error
	init bool

	net.Conn
}

func newActiveConn(conn net.Conn) (net.Conn, <-chan struct{}) {
	aConn := activeConn{Conn: conn}
	sig := make(chan struct{})

	go func() {
		buf := make([]byte, 1)
		aConn.n, aConn.err = conn.Read(buf)
		aConn.b = buf[0]
		close(sig)
	}()

	return &aConn, sig
}

func (a *activeConn) Read(b []byte) (n int, err error) {
	if !a.init {
		if len(b) == 0 {
			return 0, a.err
		}

		a.init = true
		b[0] = a.b
		return a.n, a.err
	}

	return a.Conn.Read(b)
}
