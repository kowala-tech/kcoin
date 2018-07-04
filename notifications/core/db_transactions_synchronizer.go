package core

import (
	"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/sirupsen/logrus"
)

type DbTransactionsPersistance struct {
	Persistence  persistence.TransactionRepository `inj:""`

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
}

func (tp *DbTransactionsPersistance) Stop() {
	tp.logger.Debug("Stopping...")
}
