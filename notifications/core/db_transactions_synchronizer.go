package core

import (
	"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
	"github.com/sirupsen/logrus"
)

type DbTransactionsPersistance struct {
	Persistence persistence.TransactionRepository `inj:""`
	Subscriber  pubsub.Subscriber                 `inj:""`

	logger *logrus.Entry
}

func NewDbTransactionsPersistence(logger *logrus.Entry) *DbTransactionsPersistance {
	return &DbTransactionsPersistance{
		logger: logger.WithField("app", "core/transactions_db_persistence"),
	}
}

func (tp *DbTransactionsPersistance) Start() error {
	tp.logger.Debug("Starting...")
	tp.logger.Debug("Getting blocks from queue...")

	return nil
}

func (tp *DbTransactionsPersistance) Stop() {
	tp.logger.Debug("Stopping...")
}
