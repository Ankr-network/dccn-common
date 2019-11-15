package rabbitmq

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/streadway/amqp"
)

//https://gist.github.com/OhBonsai/28868448ba84c38749a55ea63f22ca77

const (
	reconnectDelay = 10 * time.Second
	resendDelay    = 5 * time.Second
	resendTime     = 3
)

type publisher struct {
	reliable      bool
	url           string
	topic         string
	isConnected   bool
	done          chan struct{}
	conn          *amqp.Connection
	channel       *amqp.Channel
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
}

func newPublisher(reliable bool, url, topic string) *publisher {
	p := &publisher{
		reliable:    reliable,
		url:         url,
		topic:       topic,
		isConnected: false,
		done:        make(chan struct{}),
	}
	go p.handleConnect()
	return p
}

func (p *publisher) handleConnect() {
	for {
		p.isConnected = false
		logger.Println("Publisher is attempting connect ...")
		for !p.connect() {
			logger.Println("Publisher failed to connect ...")
			time.Sleep(reconnectDelay)
		}
		logger.Println("Publisher connected!")
		select {
		case <-p.done:
			return
		case <-p.notifyClose:
		}
	}
}

func (p *publisher) connect() bool {
	conn, err := amqp.Dial(p.url)
	if err != nil {
		return false
	}
	channel, err := conn.Channel()
	if err != nil {
		return false
	}

	if p.reliable {
		err = channel.Confirm(false)
		if err != nil {
			return false
		}
	}

	p.conn = conn
	p.channel = channel
	p.notifyClose = make(chan *amqp.Error)
	p.notifyConfirm = make(chan amqp.Confirmation)
	p.channel.NotifyClose(p.notifyClose)
	p.channel.NotifyPublish(p.notifyConfirm)
	p.isConnected = true
	return true
}

func (p *publisher) Publish(m interface{}) error {
	if !p.isConnected {
		return ErrPublisherConn
	}

	msg, ok := m.(proto.Message)
	if !ok {
		return ErrMessageIsNotProtoMessage
	}
	body, err := proto.Marshal(msg)
	if err != nil {
		return ErrProtoMarshal
	}

	publishing := amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "application/protobuf",
		ContentEncoding: "",
		Body:            body,
		DeliveryMode:    amqp.Transient,
		Priority:        0,
	}

	if p.reliable {
		currentTime := 0
		for {
			publishing.DeliveryMode = amqp.Persistent
			if err := p.channel.Publish(
				defaultExchange,
				p.topic,
				false,
				false,
				publishing,
			); err != nil {
				currentTime += 1
				logger.Printf("Publish message(reliable) failed, retry %d time\n", currentTime)
				if currentTime < resendTime {
					time.Sleep(resendDelay)
					continue
				}
				return err
			}
			ticker := time.NewTicker(resendDelay)
			select {
			case confirm := <-p.notifyConfirm:
				if confirm.Ack {
					//logger.Printf("Publish message(reliable): msg=> %+v, DeliveryTag=>%d\n", m, confirm.DeliveryTag)
					return nil
				}
			case <-ticker.C:
				//logger.Printf("Publish message(reliable) timeout: msg=> %+v\n", m)
			}
			ticker.Stop()
			logger.Printf("Publish message(reliable): send message failed\n")
			return ErrPublishTimeout
		}
	}
	err = p.channel.Publish(
		defaultExchange,
		p.topic,
		false,
		false,
		publishing,
	)
	//logger.Printf("Publish mesage(unreliable): %+v\n", m)
	return err
}

func (p *publisher) Close() error {
	if !p.isConnected {
		return ErrPublishConnClose
	}
	err := p.channel.Close()
	if err != nil {
		return err
	}
	err = p.conn.Close()
	if err != nil {
		return err
	}
	close(p.done)
	p.isConnected = false
	return nil
}
