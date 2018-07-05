package core

import (
	"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
	"github.com/sirupsen/logrus"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/golang/protobuf/proto"
	)

type TransactionsPersistanceWorker struct {
	Persistence persistence.TransactionRepository `inj:""`
	Subscriber  pubsub.Subscriber                 `inj:""`

	logger *logrus.Entry
}

func NewTransactionsPersistanceWorker(logger *logrus.Entry) *TransactionsPersistanceWorker {
	return &TransactionsPersistanceWorker{
		logger: logger.WithField("app", "core/transactions_db_persistence"),
	}
}

func (tp *TransactionsPersistanceWorker) Start() error {
	tp.logger.Debug("Starting...")
	tp.logger.Debug("Getting blocks from queue...")

	tp.Subscriber.AddHandler(tp)
	err := tp.Subscriber.Start()
	if err != nil {
		return nil
	}

	return nil
}

func (tp *TransactionsPersistanceWorker) Stop() {
	tp.logger.Debug("Stopping...")
	tp.Subscriber.Stop()
}


func (tp *TransactionsPersistanceWorker) HandleMessage(topic string, data []byte) error {
	var tx = new(protocolbuffer.Transaction)

	err := proto.Unmarshal(data, tx)
	if err != nil {
		return err
	}

	tp.logger.Debug("Saving transaction received: %s", tx.Hash)
	tp.Persistence.Save(tx)

	return nil
}
