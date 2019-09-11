package rabbitmq

import (
	"github.com/golang/protobuf/proto"
	"reflect"
)

type handler struct {
	methodValue reflect.Value
	msgType     reflect.Type
}

func newHandler(h interface{}) (*handler, error) {
	ht := reflect.TypeOf(h)
	if ht.Kind() != reflect.Func {
		return nil, ErrInvalidHandler
	}

	if ht.NumIn() != 1 {
		return nil, ErrInvalidHandler
	}

	if ht.NumOut() != 1 {
		return nil, ErrInvalidHandler
	}

	mt := ht.In(0)
	if err := checkIsProtoMessage(mt); err != nil {
		return nil, ErrInvalidHandler
	}

	et := ht.Out(0)
	if err := checkIsError(et); err != nil {
		return nil, ErrInvalidHandler
	}

	return &handler{
		methodValue: reflect.ValueOf(h),
		msgType:     mt,
	}, nil
}

func (h *handler) newMessage() proto.Message {
	return reflect.New(h.msgType.Elem()).Interface().(proto.Message)
}

func (h *handler) call(msg proto.Message) error {
	in := []reflect.Value{reflect.ValueOf(msg)}
	out := h.methodValue.Call(in)
	if out[0].IsNil() {
		return nil
	}
	return out[0].Interface().(error)
}
