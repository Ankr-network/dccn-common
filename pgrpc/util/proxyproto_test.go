package util

import (
	"io/ioutil"
	"net"
	"testing"
)

func Test_ParseProxyProto1(t *testing.T) {
	c1, c2 := net.Pipe()
	go func() {
		c1.Write([]byte("PROXY TCP4 192.168.1.2 192.168.1.3 10086 10010\r\n123"))
		c1.Close()
	}()

	c2, addr, err := ParseProxyProto(c2)
	if err != nil {
		t.Error(err)
	}
	if addr != "192.168.1.2" {
		t.Errorf("parse addr fail: %s", addr)
	}

	data, err := ioutil.ReadAll(c2)
	if err != nil {
		t.Error(err)
	}
	if string(data) != "123" {
		t.Errorf("parse data fail: %s", data)
	}
}

func Test_ParseProxyProto2(t *testing.T) {
	c1, c2 := net.Pipe()
	go func() {
		c1.Write(signatureV2)
		c1.Write([]byte{0x02, 0x11, 0, 12, 192, 168, 1, 2, 192, 168, 1, 3, 10086 >> 8, 10086 & 0xFF, 10010 >> 8, 10010 & 0xFF})
		c1.Write([]byte("123"))
		c1.Close()
	}()

	c2, addr, err := ParseProxyProto(c2)
	if err != nil {
		t.Error(err)
	}
	if addr != "192.168.1.2" {
		t.Errorf("parse addr fail: %s", addr)
	}

	data, err := ioutil.ReadAll(c2)
	if err != nil {
		t.Error(err)
	}
	if string(data) != "123" {
		t.Errorf("parse data fail: %s", data)
	}
}

func Test_ParseProxyProto3(t *testing.T) {
	str := `1234567890987623456789o54324567890-098765434567890-09875434567890765
	434567890ewetyuioiuytrertyuijhvcxsertyhbvcxder6ujbvcder5678ijhgfre4567yhgdse
	4567ujhvde478ujhfder4578ijhfdre45678ijhfde45678ujbvcdrtyujvcdertyuijbvfr567u
	jhfr567uijvfr567ujbvfrt567ujbvfrt67ujbvcfrtujvfrt67ujhgfrt678ijhgft678ikjuhi
	434567890ewetyuioiuytrertyuijhvcxsertyhbvcxder6ujbvcder5678ijhgfre4567yhgdse
	4567ujhvde478ujhfder4578ijhfdre45678ijhfde45678ujbvcdrtyujvcdertyuijbvfr567u
	jhfr567uijvfr567ujbvfrt567ujbvfrt67ujbvcfrtujvfrt67ujhgfrt678ijhgft678ikjuhi
	wuehdyewqgdiuwqeghdigwedigweqidheowgciwbecyiweucbiewhdyewidhe	dhuweqdbeyud
	hiwebhcviwecyiwehcigweidhuewgdhiowehu`
	c1, c2 := net.Pipe()
	go func() {
		c1.Write([]byte("PROXY TCP4 192.168.1.2 192.168.1.3 10086 10010\r\n"))
		c1.Write([]byte(str))
		c1.Close()
	}()

	c2, addr, err := ParseProxyProto(c2)
	if err != nil {
		t.Error(err)
	}
	if addr != "192.168.1.2" {
		t.Errorf("parse addr fail: %s", addr)
	}

	data, err := ioutil.ReadAll(c2)
	if err != nil {
		t.Error(err)
	}
	if string(data) != str {
		t.Errorf("parse data fail: %s", data)
	}
}
