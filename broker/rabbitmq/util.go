package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"reflect"
)

var (
	protoMessageType         = reflect.TypeOf((*proto.Message)(nil)).Elem()
	errorType                = reflect.TypeOf((*error)(nil)).Elem()
	errTypeIsNotPtr          = errors.New("type must be pointer")
	errTypeIsNotProtoMessage = errors.New("type must be proto.Message")
	errTypeIsNotError        = errors.New("type must be error")
)

func channel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		return channel, fmt.Errorf("connection.Channel error: %w", err)
	}
	return channel, nil
}

func dial(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("amqp.Dail error: %w", err)
	}
	return conn, nil
}

func exchangeDeclare(name string, channel *amqp.Channel) error {
	if err := channel.ExchangeDeclare(name, "topic", true, false, false, false, nil); err != nil {
		return fmt.Errorf("channel.ExchangeDeclare error: %w", err)
	}
	return nil
}

func queueBind(queue, key, exchange string, channel *amqp.Channel) error {
	if err := channel.QueueBind(queue, key, exchange, false, nil); err != nil {
		return fmt.Errorf("channel.QueueBind error: %w", err)
	}
	return nil
}

func queueDeclare(name string, reliable bool, channel *amqp.Channel) error {
	args := amqp.Table{}
	if !reliable {
		args["x-message-ttl"] = 20000 // 20 second
	}
	if _, err := channel.QueueDeclare(name, true, false, false, false, args); err != nil {
		return fmt.Errorf("channel.QueueDeclare error: %w", err)
	}
	return nil
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
