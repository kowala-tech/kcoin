package pubsub

import nsq "github.com/nsqio/go-nsq"
import "github.com/sirupsen/logrus"

type nsqPublisher struct {
	producer *nsq.Producer
	logger   *logrus.Entry
}

func NewNSQPublisher(address string, logger *logrus.Entry) (Publisher, error) {
	publisher := &nsqPublisher{
		logger: logger.WithField("app", "pubsub/nsq_publisher"),
	}

	producer, err := nsq.NewProducer(address, nsq.NewConfig())
	if err != nil {
		publisher.logger.WithError(err).Error("Error creating")
		return nil, err
	}
	producer.SetLogger(NewNSQLogger(publisher.logger))
	err = producer.Ping()
	if err != nil {
		publisher.logger.WithError(err).Error("Ping error")
		return nil, err
	}
	publisher.producer = producer
	return publisher, nil
}

func (p *nsqPublisher) Publish(topic string, data []byte) error {
	return p.producer.Publish(topic, data)
}

func (p *nsqPublisher) Stop() {
	p.logger.Debug("Stopping...")
	p.producer.Stop()
}
