// Package rabbitmq provides a RabbitMQ broker
package rabbitmq

import (
	"github.com/Ankr-network/dccn-common/broker"
	"github.com/golang/protobuf/proto"
	"log"
	"regexp"
)

const (
	defaultRabbitURL = "amqp://guest:guest@127.0.0.1:5672"
	defaultExchange  = "ankr.micro"
)

type rabbitBroker struct {
	url string
}

func NewBroker(args ...string) broker.Broker {
	var url string
	if len(args) == 0 {
		url = defaultRabbitURL
	} else if len(args) > 1 {
		log.Fatalf("invalid explict args, expect just url")
	} else {
		url = args[0]
		if !regexp.MustCompile("^amqp(s)?://.*").MatchString(url) {
			logger.Fatalf("invalid RabbitMQ url: %s", url)
		}
	}

	return &rabbitBroker{url: url}
}

func (r *rabbitBroker) Publisher(topic string) (broker.Publisher, error) {
	p := &rabbitPublisher{
		url:   r.url,
		topic: topic,
	}
	if err := p.Connect(); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *rabbitBroker) Subscribe(name, topic string, handler interface{}) error {
	h, err := newHandler(handler)
	if err != nil {
		return err
	}

	s := rabbitSubscriber{
		name:  name,
		url:   r.url,
		topic: topic,
	}

	if err := s.Connect(); err != nil {
		return err
	}

	deliveries, err := s.Consume()
	if err != nil {
		return err
	}

	go func() {
		for d := range deliveries {
			msg := h.newMessage()
			if err := proto.Unmarshal(d.Body, msg); err != nil {
				log.Printf("proto.Unmarshal error: %v", err)
				continue
			}

			if err := h.call(msg); err != nil {
				log.Printf("handle message %v error: %v", msg, h.call(msg))
			}
		}
	}()

	return nil
}
