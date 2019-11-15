package rabbitmq

import (
	"time"

	"github.com/streadway/amqp"
)

type subscriber struct {
	reliable    bool
	name        string
	url         string
	topic       string
	isConnected bool
	done        chan struct{}
	conn        *amqp.Connection
	channel     *amqp.Channel
	notifyClose chan *amqp.Error
}

func newSubscriber(reliable bool, name, url, topic string) *subscriber {
	s := &subscriber{
		reliable:    reliable,
		name:        name,
		url:         url,
		topic:       topic,
		isConnected: false,
		done:        make(chan struct{}),
	}
	go s.handlerConnect()
	return s
}

func (s *subscriber) handlerConnect() {
	for {
		s.isConnected = false
		logger.Println("Subscriber is attempting connect ...")
		for !s.connect() {
			logger.Println("Subscriber failed to connect ...")
			time.Sleep(reconnectDelay)
		}
		logger.Println("Subscriber connected!")
		select {
		case <-s.done:
			return
		case <-s.notifyClose:
		}
	}
}

func (s *subscriber) connect() bool {
	conn, err := amqp.Dial(s.url)
	if err != nil {
		return false
	}
	channel, err := conn.Channel()
	if err != nil {
		return false
	}

	//Exchange Declare
	err = channel.ExchangeDeclare(
		s.name,
		defaultExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return false
	}

	args := amqp.Table{}
	if !s.reliable {
		args["x-message-ttl"] = 20 * 1000
	}
	_, err = channel.QueueDeclare(
		s.name,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return false
	}
	//queue, key, exchange, false, nil
	err = channel.QueueBind(
		s.name,
		s.topic,
		defaultExchange,
		false,
		nil,
	)
	if err != nil {
		return false
	}

	s.conn = conn
	s.channel = channel
	s.notifyClose = make(chan *amqp.Error)
	s.channel.NotifyClose(s.notifyClose)
	s.isConnected = true
	return true
}

func (s *subscriber) Consume() (<-chan amqp.Delivery, error) {
	if !s.isConnected {
		return nil, ErrSubscriberConnClose
	}
	//queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table
	return s.channel.Consume(s.name, "", !s.reliable, false, false, false, nil)
}

func (s *subscriber) Close() error {
	if !s.isConnected {
		return ErrSubscriberConnClose
	}

	err := s.channel.Close()
	if err != nil {
		return err
	}
	err = s.conn.Close()
	if err != nil {
		return err
	}
	close(s.done)
	s.isConnected = false
	return nil
}
