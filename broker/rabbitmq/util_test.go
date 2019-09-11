package rabbitmq

import (
	"reflect"
	"testing"
)

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
