package ankrmicro

// Publisher is syntactic sugar for publishing
type Publisher struct {
	topic string
}

// Publish pushes data to a broker
func (p *Publisher) Publish(data interface{}) error {
	return Send(p.topic, data)
}

func (p *Publisher)GetTopic() string {
	return p.topic
}

// NewService creates and returns a new Service based on the packages within.
func NewService() GRPCService {
	return NewGRPCService()
}

// NewPublisher returns a new Publisher
func NewPublisher(topic string) *Publisher {
	return &Publisher{topic}
}

// RegisterSubscriber is syntactic sugar for registering a subscriber
func RegisterSubscriber(topic string, handler interface{}) error {
	go Receive(topic, handler)
	return nil
}
