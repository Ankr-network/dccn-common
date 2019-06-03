package util

import (
	"bytes"
	"io"
	"net"

	"github.com/pkg/errors"
)

/************************************************************************/
var signatureV1 = []byte("PROXY ")
var signatureV2 = []byte{0x0D, 0x0A, 0x0D, 0x0A, 0x00, 0x0D, 0x0A, 0x51, 0x55, 0x49, 0x54, 0x0A}
var crlf = []byte("\r\n")
var space = []byte(" ")

// "PROXY UNKNOWN ffff:f...f:ffff ffff:f...f:ffff 65535 65535\r\n"
const maximumLenV1 = 107

// signature(12)+ver_cmd(1)+fam(1)+len(2)+socketX2(216)
const maximumLenV2 = 232

// V1 "PROXY UNKNOW" or V2 signature
const signatureLen = 12

func ParseProxyProto(c net.Conn) (net.Conn, string, error) {
	buf := make([]byte, maximumLenV2)
	if n, err := io.ReadFull(c, buf[:signatureLen]); err != nil {
		if err == io.ErrUnexpectedEOF {
			return &bufConn{Conn: c, buf: buf[:n], err: io.EOF}, "", nil
		}
		return &bufConn{Conn: c, buf: buf[:n], err: err}, "", nil
	}

	switch {
	case bytes.HasPrefix(buf, signatureV1):
		n, idx, err := ReadSignal(c, buf[signatureLen:maximumLenV1], crlf)
		n += signatureLen
		idx += signatureLen
		if err != nil {
			return &bufConn{Conn: c, buf: buf[:n], err: err}, "", nil
		}

		secs := bytes.Split(buf[:idx], space)
		if len(secs) < 2 {
			return nil, "", errors.New("invalid proxy protocol v1")
		}

		switch string(secs[1]) {
		case "TCP4":
			// "PROXY TCP4 255.255.255.255 255.255.255.255 65535 65535\r\n"
			fallthrough
		case "TCP6":
			// "PROXY TCP6 ffff:f...f:ffff ffff:f...f:ffff 65535 65535\r\n"
			if len(secs) != 6 {
				return nil, "", errors.New("invalid proxy protocol v1")
			}

			return &bufConn{Conn: c, buf: buf[idx+2 : n]}, string(secs[2]), nil

		case "UNKNOWN":
			// "PROXY UNKNOWN\r\n"
			// "PROXY UNKNOWN ffff:f...f:ffff ffff:f...f:ffff 65535 65535\r\n"
			return &bufConn{Conn: c, buf: buf[idx+2 : n]}, "", nil

		default:
			return nil, "", errors.New("invalid proxy protocol v1")
		}

	case bytes.HasPrefix(buf, signatureV2):
		// struct proxy_hdr_v2 {
		//     uint8_t sig[12];  /* hex 0D 0A 0D 0A 00 0D 0A 51 55 49 54 0A */
		//     uint8_t ver_cmd;  /* protocol version and command */
		//     uint8_t fam;      /* protocol family and address */
		//     uint16_t len;     /* number of following bytes part of the header */
		// };
		hdrLen := signatureLen + 4
		n, err := io.ReadFull(c, buf[signatureLen:hdrLen])
		n += signatureLen
		if err != nil {
			return &bufConn{Conn: c, buf: buf[:n], err: err}, "", nil
		}

		switch buf[13] & 0xF0 {
		case 0x00: // AF_UNSPEC
			//The receiver should ignore address information.
			return &bufConn{Conn: c, buf: buf[hdrLen:n]}, "", nil

		case 0x10: // AF_INET
			// Address length is 2*4 + 2*2 = 12 bytes.
			// struct {        /* for TCP/UDP over IPv4, len = 12 */
			//     uint32_t src_addr;
			//     uint32_t dst_addr;
			//     uint16_t src_port;
			//     uint16_t dst_port;
			// } ipv4_addr;
			nn, err := io.ReadFull(c, buf[hdrLen:hdrLen+12])
			n += nn
			if err != nil {
				return &bufConn{Conn: c, buf: buf[:n], err: err}, "", nil
			}

			ip := net.IPv4(buf[16], buf[17], buf[18], buf[19]).String()
			return c, ip, nil

		case 0x20: // AF_INET6
			// Address length is 2*16 + 2*2 = 36 bytes.
			// struct {        /* for TCP/UDP over IPv6, len = 36 */
			//      uint8_t  src_addr[16];
			//      uint8_t  dst_addr[16];
			//      uint16_t src_port;
			//      uint16_t dst_port;
			// } ipv6_addr;
			nn, err := io.ReadFull(c, buf[hdrLen:hdrLen+36])
			n += nn
			if err != nil {
				return &bufConn{Conn: c, buf: buf[:n], err: err}, "", nil
			}

			ip := net.IP(buf[16:32]).String()
			return c, ip, nil

		case 0x30: // AF_UNIX
			// Address length is 2*108 = 216 bytes.
			// struct {        /* for AF_UNIX sockets, len = 216 */
			//      uint8_t src_addr[108];
			//      uint8_t dst_addr[108];
			// } unix_addr;
			nn, err := io.ReadFull(c, buf[hdrLen:hdrLen+216])
			n += nn
			if err != nil {
				return &bufConn{Conn: c, buf: buf[:n], err: err}, "", nil
			}

			return c, string(buf[16 : 16+108]), nil
		default:
			return nil, "", errors.New("invalid proxy protocol v2")
		}

	default:
		return &bufConn{Conn: c, buf: buf[:signatureLen]}, "", nil
	}
}

/**********************************************************************/
type bufConn struct {
	net.Conn
	buf    []byte
	err    error
	offset int
	init   bool // read
}

func (c *bufConn) Read(b []byte) (n int, err error) {
	if c.init {
		return c.Conn.Read(b)
	}

	if len(c.buf) == c.offset+1 {
		return 0, c.err
	}

	n = copy(b, c.buf[c.offset:])
	c.offset += n
	if c.offset == len(c.buf) && c.err == nil {
		c.init = true
	}
	return n, nil
}

func ReadSignal(c net.Conn, b, signal []byte) (n, idx int, err error) {
	lenB := len(b)

	for {
		nn, err := c.Read(b)
		n += nn
		if err != nil {
			return n, -1, err
		}

		if idx := bytes.Index(b[:n], signal); idx >= 0 {
			return n, idx, nil
		}

		if n == lenB {
			return n, -1, nil
		}
	}
}
