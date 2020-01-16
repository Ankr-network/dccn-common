package rabbitmq

import (
	"errors"
	"reflect"

	"github.com/golang/protobuf/proto"
)

var (
	ErrMessageIsNotProtoMessage = errors.New("message must be proto.Message")
	ErrInvalidHandler           = errors.New("invalid handler, must be func and have single proto.Message implement and return single error")
	ErrPublishMessageNotAck     = errors.New("message not ack by broker")
	ErrPublishTimeout           = errors.New("publish message wait confirmation timeout")
	ErrProtoMarshal             = errors.New("marshal proto failed")
	ErrChannelConfirm           = errors.New("set confirm mode failed")
	ErrPublisherConn            = errors.New("publisher connection failed")
	ErrPublishConnClose         = errors.New("publisher connection already close")
	ErrSubscriberConnClose      = errors.New("subscriber connection already close")

	protoMessageType         = reflect.TypeOf((*proto.Message)(nil)).Elem()
	errorType                = reflect.TypeOf((*error)(nil)).Elem()
	errTypeIsNotPtr          = errors.New("type must be pointer")
	errTypeIsNotProtoMessage = errors.New("type must be proto.Message")
	errTypeIsNotError        = errors.New("type must be error")
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

func checkIsProtoMessage(t reflect.Type) error {
	if t.Kind() != reflect.Ptr {
		return errTypeIsNotPtr
	}

	if !t.Implements(protoMessageType) {
		return errTypeIsNotProtoMessage
	}
	return nil
}

func checkIsError(t reflect.Type) error {
	if !t.Implements(errorType) {
		return errTypeIsNotError
	}
	return nil
}
