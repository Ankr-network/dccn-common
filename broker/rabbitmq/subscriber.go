package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

type rabbitSubscriber struct {
	reliable bool
	name     string
	url      string
	topic    string
	conn     *connection
}

func (s *rabbitSubscriber) Connect() error {
	conn, err := dial(s.url)
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

	if err := queueDeclare(s.name, s.reliable, channel); err != nil {
		return err
	}

	if err := queueBind(s.name, s.topic, defaultExchange, channel); err != nil {
		return err
	}

	if s.conn, err = newConnection(s.url, false, conn); err != nil {
		return err
	}

	s.conn.connect()

	return nil
}

func (s *rabbitSubscriber) Close() error {
	return s.conn.close()
}

func (s *rabbitSubscriber) Consume() (<-chan amqp.Delivery, error) {
	deliveries := make(chan amqp.Delivery)
	go func() {
		for {
			ds, err := consume(s.name, "", !s.reliable, s.conn.channel)
			if err != nil {
				log.Printf("consume error: %v", err)
				time.Sleep(8 * time.Second)
				continue
			}
			for d := range ds {
				deliveries <- d
			}
			log.Printf("delivery chan close, reconsume")
		}
	}()
	return deliveries, nil
}
