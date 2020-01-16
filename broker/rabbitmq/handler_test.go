package rabbitmq

import (
	"errors"
	"reflect"
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

type testProtoMessage struct{}

func (t *testProtoMessage) Reset() {
	panic("implement me")
}

func (t *testProtoMessage) String() string {
	return "testProtoMessage"
}

func (t *testProtoMessage) ProtoMessage() {
	panic("implement me")
}

type testErr struct {
	error
}

func TestCheckIsProtoMessage(t *testing.T) {
	a := testProtoMessage{}
	err := checkIsProtoMessage(reflect.TypeOf(a))
	if err == nil {
		t.Error("expect err")
	}
	if err != errTypeIsNotPtr {
		t.Errorf("expect %v but %v", errTypeIsNotPtr, err)
	}

	err = checkIsProtoMessage(reflect.TypeOf(&struct{}{}))
	if err == nil {
		t.Error("expect err")
	}
	if err != errTypeIsNotProtoMessage {
		t.Errorf("expect %v but %v", errTypeIsNotProtoMessage, err)
	}

	if err := checkIsProtoMessage(reflect.TypeOf(&testProtoMessage{})); err != nil {
		t.Error(err)
	}
}

func TestCheckIsError(t *testing.T) {
	a := struct{}{}
	err := checkIsError(reflect.TypeOf(a))
	if err == nil {
		t.Error("expect err")
	}
	if err != errTypeIsNotError {
		t.Errorf("expect %v but %v", errTypeIsNotError, err)
	}

	if err := checkIsError(reflect.TypeOf(testErr{})); err != nil {
		t.Error(err)
	}
}
