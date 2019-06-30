package pgrpc

import (
	"net"
	"time"

	"github.com/pkg/errors"
)

type listener struct {
	addr net.Addr

	connCh chan *activeConn
	stopCh chan struct{}
}

func Listen(network, address, id string, onAccept func(*net.Conn, error)) (net.Listener, error) {
	if idLen := len(id); len(id) > MAX_ID_LEN {
		return nil, errors.Errorf("id(%s) is too long", id)
	} else if idLen == 0 {
		return nil, errors.Errorf("id is empty")
	}

	var (
		ln = listener{
			connCh: make(chan *activeConn, MIN_IDLE-1),
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
			}

			aConn, err := newActiveConn(conn, id)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			ln.connCh <- aConn
			select {
			case <-aConn.init:
			case <-ln.stopCh:
				aConn.Close()
				return
			}
		}
	}()
	return &ln, nil
}

func (ln *listener) Accept() (net.Conn, error) {
	for {
		select {
		case conn := <-ln.connCh:
			return conn, nil
		case <-ln.stopCh:
			return nil, errors.New("listener has been stoped")
		}
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
	init chan struct{}

	net.Conn
}

func newActiveConn(conn net.Conn, id string) (*activeConn, error) {
	if c, ok := conn.(*net.TCPConn); ok {
		c.SetLinger(1)
		c.SetKeepAlive(true)
		c.SetKeepAlivePeriod(5 * time.Second)
	}

	aConn := activeConn{Conn: conn, init: make(chan struct{})}

	// write id
	buf := make([]byte, MAX_ID_LEN)
	copy(buf, []byte(id))

	aConn.SetDeadline(time.Now().Add(5 * time.Second))
	var err error
	for n, nn := 0, 0; nn < MAX_ID_LEN; nn += n {
		if n, err = aConn.Write(buf[nn:]); err != nil {
			close(aConn.init)
			aConn.Close()
			return nil, errors.Errorf("write id fail: %s", err)
		}
	}

	aConn.SetDeadline(time.Time{})
	return &aConn, nil
}

func (a *activeConn) Read(b []byte) (n int, err error) {
	select {
	case <-a.init:
		return a.Conn.Read(b)
	default:
		n, err = a.Conn.Read(b)
		close(a.init)
		return
	}
}
