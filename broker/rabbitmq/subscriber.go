package rabbitmq

import (
	"github.com/streadway/amqp"
)

type rabbitSubscriber struct {
	reliable bool
	name     string
	url      string
	topic    string
	conn     *Connection
	channel  *Channel
}

func (s *rabbitSubscriber) Connect() error {
	conn, err := Dial(s.url)
	if err != nil {
		return err
	}

	channel, err := conn.Channel(true)
	if err != nil {
		if err := conn.Close(); err != nil {
			logger.Printf("Connection.Close error %v", err)
		}
		return err
	}

	if err := exchangeDeclare(defaultExchange, channel.Channel); err != nil {
		if err := conn.Close(); err != nil {
			logger.Printf("Connection.Close error %v", err)
		}
		return err
	}

	if err := queueDeclare(s.name, s.reliable, channel.Channel); err != nil {
		if err := conn.Close(); err != nil {
			logger.Printf("Connection.Close error %v", err)
		}
		return err
	}

	if err := queueBind(s.name, s.topic, defaultExchange, channel.Channel); err != nil {
		if err := conn.Close(); err != nil {
			logger.Printf("Connection.Close error %v", err)
		}
		return err
	}

	s.conn = conn
	s.channel = channel

	return nil
}

func (s *rabbitSubscriber) Close() error {
	return s.conn.Close()
}

func (s *rabbitSubscriber) Consume() (<-chan amqp.Delivery, error) {
	return s.channel.Consume(s.name, "", !s.reliable, false, false, false, nil)
}
