package rabbitmq

import (
	"github.com/streadway/amqp"
	"sync/atomic"
	"time"
)

const (
	consumeRetryDelay        = 8
	channelReconnectDelay    = 4
	connectionReconnectDelay = 2
)

// Connection amqp.Connection wrapper
type Connection struct {
	*amqp.Connection
}

// Channel wrap amqp.Connection.Channel, get a auto reconnect channel
func (c *Connection) Channel(reconnect bool) (*Channel, error) {
	ch, err := channel(c.Connection)
	if err != nil {
		return nil, err
	}

	resultChannel := &Channel{
		Channel: ch,
	}

	if reconnect {
		go func() {
			for {
				reason, ok := <-resultChannel.Channel.NotifyClose(make(chan *amqp.Error))
				// exit this goroutine if closed by developer
				if !ok || resultChannel.IsClosed() {
					logger.Print("channel closed")
					if err := resultChannel.Close(); err != nil { // close again, ensure closed flag set when connection closed
						logger.Printf("Channel.Close error: %v", err)
					}
					break
				}
				logger.Printf("channel closed, reason: %v", reason)

				// reconnect if not closed by developer
				for {
					// wait 1s for connection reconnect
					time.Sleep(channelReconnectDelay * time.Second)

					ch, err := channel(c.Connection)
					if err == nil {
						logger.Print("channel recreate success")
						resultChannel.Channel = ch
						break
					}
					logger.Printf("channel recreate failed, err: %v", err)
				}
			}
		}()
	}

	return resultChannel, nil
}

// Dial wrap amqp.Dial, dial and get a reconnect connection
func Dial(url string) (*Connection, error) {
	conn, err := dial(url)
	if err != nil {
		return nil, err
	}

	connection := &Connection{
		Connection: conn,
	}

	go func() {
		for {
			reason, ok := <-connection.Connection.NotifyClose(make(chan *amqp.Error))
			// exit this goroutine if closed by developer
			if !ok {
				logger.Print("connection closed")
				break
			}
			logger.Printf("connection closed, reason: %v", reason)

			// reconnect if not closed by developer
			for {
				// wait 1s for reconnect
				time.Sleep(connectionReconnectDelay * time.Second)

				conn, err := dial(url)
				if err == nil {
					connection.Connection = conn
					logger.Print("reconnect success")
					break
				}

				logger.Printf("reconnect failed, err: %v", err)
			}
		}
	}()

	return connection, nil
}

// Channel amqp.Channel wrapper
type Channel struct {
	*amqp.Channel
	closed int32
}

// IsClosed indicate closed by developer
func (ch *Channel) IsClosed() bool {
	return atomic.LoadInt32(&ch.closed) == 1
}

// Close ensure closed flag set
func (ch *Channel) Close() error {
	if ch.IsClosed() {
		return amqp.ErrClosed
	}

	atomic.StoreInt32(&ch.closed, 1)

	return ch.Channel.Close()
}

// Consume wrap amqp.Channel.Consume, the returned delivery will end only when channel closed by developer
func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	deliveries := make(chan amqp.Delivery)

	go func() {
		for {
			d, err := ch.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
			if err != nil {
				logger.Printf("consume failed, err: %v", err)
				time.Sleep(consumeRetryDelay * time.Second)
				continue
			}

			for msg := range d {
				deliveries <- msg
			}

			// sleep before IsClose call. closed flag may not set before sleep.
			time.Sleep(consumeRetryDelay * time.Second)

			if ch.IsClosed() {
				break
			}
		}
	}()

	return deliveries, nil
}
