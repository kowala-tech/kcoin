package core

import (
	"github.com/gogo/protobuf/proto"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/notifier"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
	"github.com/sirupsen/logrus"
)

type Emailer struct {
	Notifier   notifier.Notifier `inj:""`
	Subscriber pubsub.Subscriber `inj:""`
	KV         keyvalue.KeyValue `inj:""`

	from   string
	logger *logrus.Entry
}

func NewEmailer(logger *logrus.Entry, from string) *Emailer {
	return &Emailer{
		from:   from,
		logger: logger.WithField("app", "core/emailer"),
	}
}

func (emailer *Emailer) Register() {
	emailer.logger.Debug("Registering handler...")
	emailer.Subscriber.AddHandler(emailer)
}

func (emailer *Emailer) Stop() error {
	emailer.logger.Debug("Stopping...")
	emailer.Subscriber.Stop()
	return nil
}

func (emailer *Emailer) HandleMessage(topic string, data []byte) error {
	var tx protocolbuffer.Transaction
	err := proto.Unmarshal(data, &tx)
	if err != nil {
		emailer.logger.WithError(err).Error("Error unmarshalling message")
		return err
	}

	email, err := emailer.KV.GetString(tx.To)
	if err != nil {
		emailer.logger.WithError(err).Error("Error reading keyvalue storage")
		return err
	}
	if email == "" {
		return nil
	}

	err = emailer.Notifier.Send(map[string]string{
		notifier.EmailFromKey: emailer.from,
		notifier.EmailToKey:   email,
	})
	if err != nil {
		emailer.logger.WithError(err).Error("Error sending email")
		return err
	}

	return nil
}
