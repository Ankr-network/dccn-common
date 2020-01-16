// Package broker is an interface used for asynchronous messaging
package broker

// Broker is an interface used for asynchronous messaging.
type Broker interface {
	Publisher(topic string, reliable bool) (Publisher, error)
	Subscribe(name, topic string, reliable, requeue bool, handler interface{}) error
}

type Publisher interface {
	Publish(m interface{}) error
}
