package rabbitmq

import (
	"errors"
	"testing"
)

type testSubscriber struct{}

var (
	errTest = errors.New("test error")
)

func (h *testSubscriber) handle(message *testProtoMessage) error {
	return errTest
}

func TestNewHandler(t *testing.T) {
	s := testSubscriber{}
	h, err := newHandler(s.handle)
	if err != nil {
		t.Error(err)
	}
	m := h.newMessage()
	if m.String() != "testProtoMessage" {
		t.Errorf("expect %q, but %q", "testProtoMessage", m.String())
	}

	if err := h.call(m); err != errTest {
		t.Errorf("expect %v but %v", errTest, err)
	}
}
