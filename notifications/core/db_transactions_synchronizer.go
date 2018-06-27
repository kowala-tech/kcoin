package core

import (
	"math/big"

	"github.com/kowala-tech/kcoin/notifications/blockchain"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/sirupsen/logrus"
)

type DbTransactionsSynchronizer struct {
	Blockchain   blockchain.Blockchain             `inj:""`
	Persistence  persistence.TransactionRepository `inj:""`
	ValueStorage keyvalue.Value                    `inj:""`

	logger *logrus.Entry
}

func NewDbTransactionsSynchronizer(logger *logrus.Entry) *DbTransactionsSynchronizer {
	return &DbTransactionsSynchronizer{
		logger: logger.WithField("app", "core/transactions_db_synchronize"),
	}
}

func (tp *DbTransactionsSynchronizer) Start() error {
	tp.logger.Debug("Starting...")
	tp.logger.Debug("Fetching last processed block...")
	blockNum, err := tp.ValueStorage.GetInt64()
	if err != nil {
		tp.logger.WithError(err).Error("Error getting last block number from the value storage")
		return err
	}
	if blockNum > 0 {
		tp.logger.WithField("blockNum", blockNum).Info("Seeking blockchain to last processed block.")
		tp.Blockchain.Seek(big.NewInt(blockNum))
	}
	err = tp.Blockchain.OnBlock(tp)
	if err != nil {
		tp.logger.WithError(err).Error("Error registering OnBlock handler")
		return err
	}
	return tp.Blockchain.Start()
}

func (tp *DbTransactionsSynchronizer) Stop() {
	tp.logger.Debug("Stopping...")
	tp.Blockchain.Stop()
}

func (tp *DbTransactionsSynchronizer) HandleBlock(block *blockchain.Block) {
	tp.logger.
		WithField("blockNum", block.Number).
		WithField("transactionsNum", len(block.Transactions)).
		Info("Block received")
	tp.ValueStorage.PutInt64(block.Number.Int64())
	for _, tx := range block.Transactions {
		tp.Persistence.Save(tx)
		tp.logger.
			WithField("tx", tx.Hash).
			Info("Transaction saved")
	}
}
