package main

import (
	"fmt"
	"github.com/Ankr-network/dccn-common/broker/example/proto"
	"github.com/Ankr-network/dccn-common/broker/rabbitmq"
	"log"
	"time"

	"github.com/Ankr-network/dccn-common/broker"
)

var (
	topic           = "ankr.topic.hello"
	ankrBroker      broker.Broker
	helloPublisher  broker.Publisher
	helloSubscriber logHandler
)

type logHandler struct{}

func (s *logHandler) handle(h *proto.Hello) error {
	log.Printf("handle %+v", h)
	return nil
}

func init() {
	var err error
	ankrBroker = rabbitmq.NewBroker()
	if helloPublisher, err = ankrBroker.Publisher(topic); err != nil {
		log.Fatal(err)
	}
	if err := ankrBroker.Subscribe(topic, helloSubscriber.handle); err != nil {
		log.Fatal(err)
	}
}

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := proto.Hello{Name: fmt.Sprintf("No.%d", i)}
		if err := helloPublisher.Publish(&msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			log.Printf("[pub] pubbed message: %v", msg)
		}
		i++
	}
}

func main() {
	go pub()
	<-time.After(time.Second * 100)
}
