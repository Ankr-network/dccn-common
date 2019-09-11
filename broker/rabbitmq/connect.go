package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

type connection struct {
	confirm       bool
	url           string
	conn          *amqp.Connection
	channel       *amqp.Channel
	connCloseChan chan *amqp.Error
	chnCloseChan  chan *amqp.Error
	changeChChan  chan struct{}
	stopChan      chan struct{}
	confirmChan   chan amqp.Confirmation
}

func newConnection(url string, confirm bool, conn *amqp.Connection) (*connection, error) {
	connCloseCh := make(chan *amqp.Error)
	chnCloseCh := make(chan *amqp.Error)

	conn.NotifyClose(connCloseCh)

	channel, err := channel(conn)
	if err != nil {
		return nil, err
	}

	if confirm {
		if err := channel.Confirm(false); err != nil {
			return nil, err
		}
	}

	channel.NotifyClose(chnCloseCh)

	r := &connection{
		url:           url,
		conn:          conn,
		channel:       channel,
		connCloseChan: connCloseCh,
		changeChChan:  make(chan struct{}),
		stopChan:      make(chan struct{}),
	}

	if confirm {
		r.confirmChan = channel.NotifyPublish(make(chan amqp.Confirmation))
	}

	return r, nil
}

func (r *connection) connect() {
	select {
	case <-r.stopChan:
		log.Printf("close connect")
		return
	default:
		// no op
	}

	go func() {
		for {
			log.Printf("waiting connection close")
			select {
			case amqpErr := <-r.connCloseChan:
				log.Printf("connection close: %v", amqpErr)
			case <-r.stopChan:
				return
			}

			var conn *amqp.Connection
			for {
				var err error
				conn, err = dial(r.url)
				if err != nil {
					log.Printf("dial error: %v", err)
					time.Sleep(2 * time.Second)
					continue
				}
				break
			}

			log.Print("new connection")
			r.conn = conn
			r.connCloseChan = make(chan *amqp.Error, 1)
			conn.NotifyClose(r.connCloseChan)
			r.changeChChan <- struct{}{}
		}
	}()

	go func() {
		for {
			log.Printf("waiting channel close")
			select {
			case amqpErr := <-r.chnCloseChan:
				log.Printf("channel close: %v", amqpErr)
			case <-r.changeChChan:
				log.Printf("channel reset for connection close")
			case <-r.stopChan:
				return
			}
			var c *amqp.Channel
			for {
				var err error
				c, err = channel(r.conn)
				if err != nil {
					log.Printf("channel error: %v", err)
					time.Sleep(4 * time.Second)
					continue
				}
				if r.confirm {
					if err := c.Confirm(false); err != nil {
						log.Printf("channel.Confirm error: %v", err)
						time.Sleep(4 * time.Second)
						continue
					}
					r.confirmChan = c.NotifyPublish(make(chan amqp.Confirmation, 1))
				}
				break
			}
			log.Print("new channel")
			r.channel = c
			r.chnCloseChan = make(chan *amqp.Error, 1)
			c.NotifyClose(r.chnCloseChan)
		}
	}()
}

func (r *connection) close() error {
	close(r.stopChan)
	if err := r.channel.Close(); err != nil {
		return err
	}
	if err := r.conn.Close(); err != nil {
		return err
	}
	return nil
}
