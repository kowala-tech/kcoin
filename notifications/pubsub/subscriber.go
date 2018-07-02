package pubsub

//go:generate moq -out message_handler_mock.go . MessageHandler
type MessageHandler interface {
	HandleMessage(topic string, data []byte) error
}

//go:generate moq -out subscriber_mock.go . Subscriber
type Subscriber interface {
	AddHandler(MessageHandler)
	Start() error
	Stop()
}
