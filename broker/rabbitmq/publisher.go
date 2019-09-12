package rabbitmq

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

type rabbitPublisher struct {
	reliable bool
	url      string
	topic    string
	conn     *Connection
}

func (p *rabbitPublisher) Connect() error {
	conn, err := Dial(p.url)
	if err != nil {
		return err
	}

	channel, err := conn.Channel(false)
	if err != nil {
		if err := conn.Close(); err != nil {
			logger.Printf("Connection.Close error: %v", err)
		}
		return err
	}
	defer func() {
		if err := channel.Close(); err != nil {
			logger.Printf("Channel.Close error: %v", err)
		}
	}()

	if err := exchangeDeclare(defaultExchange, channel.Channel); err != nil {
		if err := conn.Close(); err != nil {
			logger.Printf("Connection.Close error: %v", err)
		}
		return err
	}

	p.conn = conn

	return nil
}

func (p *rabbitPublisher) Close() error {
	return p.conn.Close()
}

func (p *rabbitPublisher) Publish(m interface{}) error {
	// ever single channel for publish
	channel, err := p.conn.Channel(false)
	if err != nil {
		return err
	}
	defer func() {
		if err := channel.Close(); err != nil {
			logger.Printf("Channel.Close error %v", err)
		}
	}()

	msg, ok := m.(proto.Message)
	if !ok {
		return ErrMessageIsNotProtoMessage
	}
	body, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("proto.Marshal error: %w", err)
	}
	publishing := amqp.Publishing{
		ContentType: "application/protobuf",
		Body:        body,
	}

	if p.reliable {
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("Channel.Confirm error: %w", err)
		}

		confirmCh := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		publishing.DeliveryMode = amqp.Persistent

		if err := channel.Publish(defaultExchange, p.topic, false, false, publishing); err != nil {
			return fmt.Errorf("channel.Publish error: %w", err)
		}

		confirm := <-confirmCh
		if !confirm.Ack {
			return ErrPublishMessageNotAck
		}
	} else {
		if err := channel.Publish(defaultExchange, p.topic, false, false, publishing); err != nil {
			return fmt.Errorf("channel.Publish error: %w", err)
		}
	}

	return nil
}
