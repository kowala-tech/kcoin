package pubsub

//go:generate moq -out publisher_mock.go . Publisher
type Publisher interface {
	Publish(topic string, data []byte) error
	Stop()
}
