package rabbitmq

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

type rabbitPublisher struct {
	url   string
	topic string
	conn  *connection
}

func (p *rabbitPublisher) Connect() error {
	conn, err := dial(p.url)
	if err != nil {
		return err
	}

	channel, err := channel(conn)
	if err != nil {
		return err
	}
	defer channel.Close()

	if err := exchangeDeclare(defaultExchange, channel); err != nil {
		return err
	}

	if p.conn, err = newConnection(p.url, conn); err != nil {
		return err
	}
	p.conn.connect()

	return nil
}

func (p *rabbitPublisher) Close() error {
	return p.conn.close()
}

func (p *rabbitPublisher) Publish(m interface{}) error {
	msg, ok := m.(proto.Message)
	if !ok {
		return ErrMessageIsNotProtoMessage
	}
	body, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("proto.Marshal error: %w", err)
	}

	if err := p.conn.channel.Publish(defaultExchange, p.topic, false, false, amqp.Publishing{
		ContentType: "application/protobuf",
		Body:        body,
	}); err != nil {
		return fmt.Errorf("channel.Publish error: %w", err)
	}
	return nil
}
