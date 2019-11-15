// Package rabbitmq provides a RabbitMQ broker
package rabbitmq

import (
	"github.com/Ankr-network/dccn-common/broker"
	"github.com/golang/protobuf/proto"
	"regexp"
	"time"
)

const (
	defaultRabbitURL    = "amqp://guest:guest@127.0.0.1:5672"
	defaultExchange     = "ankr.micro"
	defaultExchangeType = "topic"
	nackDelay           = 5
)

var (
	rabbitURLRegx = regexp.MustCompile(`^amqp(s)?://.*`)
)

type rabbitBroker struct {
	url string
}

func NewBroker(url ...string) broker.Broker {
	var host string
	if len(url) == 0 {
		host = defaultRabbitURL
	} else {
		host = url[0]
	}

	if !rabbitURLRegx.MatchString(host) {
		return nil
	}
	return &rabbitBroker{host}
}

/*
1. connection
2. channel over connection
3. exchange declare
4. queue declare
5. bind exchange, queue
*/
func (r *rabbitBroker) Publisher(topic string, reliable bool) (broker.Publisher, error) {
	p := newPublisher(reliable, r.url, topic)
	return p, nil
}

func (r *rabbitBroker) Subscribe(name, topic string, reliable, requeue bool, handler interface{}) error {
	h, err := newHandler(handler)
	if err != nil {
		return err
	}
	s := newSubscriber(reliable, name, r.url, topic)

	go func() {
		for {
			if !s.isConnected {
				time.Sleep(reconnectDelay)
				continue
			}
			delivery, err := s.Consume()
			if err != nil {
				logger.Printf("Subscriber, channel.Consume: %+v", err)
				continue
			}
			logger.Println("Subscriber start to consumer ...")
			for d := range delivery {
				var err error
				msg := h.newMessage()
				if err = proto.Unmarshal(d.Body, msg); err != nil {
					logger.Printf("failed to unmarshal body, err: %+v", err)
					continue
				}
				err = h.call(msg)
				if err != nil {
					logger.Printf("failed to handler msg, msg: %+v, err: %+v", msg, err)
					err = d.Nack(false, requeue)
					if err != nil {
						logger.Printf("failed to Nack, err: %+v", err)
					}
					continue
				}
				//logger.Printf("handler message successfully, msg: %+v", msg)
				if s.reliable {
					err = d.Ack(false)
					if err != nil {
						logger.Printf("failed to Ack, err: %+v", err)
					}
				}
			}

		}
	}()
	return nil
}