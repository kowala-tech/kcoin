package pubsub

import nsq "github.com/nsqio/go-nsq"
import "github.com/sirupsen/logrus"

type nsqSubscriber struct {
	handlers map[MessageHandler]struct{}

	address  string
	channel  string
	topic    string
	consumer *nsq.Consumer
	logger   *logrus.Entry
}

func NewNSQSubscriber(topic, channel, address string, logger *logrus.Entry) Subscriber {
	return &nsqSubscriber{
		handlers: map[MessageHandler]struct{}{},
		topic:    topic,
		address:  address,
		channel:  channel,
		logger:   logger.WithField("app", "pubsub/nsq_subscriber"),
	}
}

func (s *nsqSubscriber) Start() error {
	s.logger.Debug("Starting...")
	consumer, err := nsq.NewConsumer(s.topic, s.channel, nsq.NewConfig())
	if err != nil {
		s.logger.WithError(err).Error("Error starting the nsq consumer")
		return err
	}
	consumer.SetLogger(NewNSQLogger(s.logger))
	s.consumer = consumer
	consumer.AddHandler(s)

	err = consumer.ConnectToNSQD(s.address)
	if err != nil {
		s.logger.WithError(err).Error("Error connecting to nsq")
		return err
	}

	return nil
}

func (s *nsqSubscriber) Stop() {
	s.logger.Debug("Stopping...")

	s.consumer.Stop()
	<-s.consumer.StopChan
}

func (s *nsqSubscriber) AddHandler(handler MessageHandler) {
	s.handlers[handler] = struct{}{}
}

func (s *nsqSubscriber) HandleMessage(message *nsq.Message) error {
	for handler := range s.handlers {
		if err := handler.HandleMessage(s.topic, message.Body); err != nil {
			return err
		}
	}
	return nil
}
