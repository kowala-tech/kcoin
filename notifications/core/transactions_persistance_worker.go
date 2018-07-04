package core

import (
	"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
	"github.com/sirupsen/logrus"
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

	return nil
}

func (tp *TransactionsPersistanceWorker) Stop() {
	tp.logger.Debug("Stopping...")
}
