package p2p

type Publisher interface {
	Publish(data []byte) error
}

type Subscriber interface {
	Subscribe(topic string) error
}

type PublisherSubscriber interface {
	Publisher
	Subscriber
}

type Host interface {
	Start()
	Stop()
	PublisherSubscriber
}
