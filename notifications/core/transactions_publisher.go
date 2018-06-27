package core

import (
	"math/big"

	"github.com/gogo/protobuf/proto"
	"github.com/kowala-tech/kcoin/notifications/blockchain"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
	"github.com/sirupsen/logrus"
)

type TransactionsPublisher struct {
	Blockchain   blockchain.Blockchain `inj:""`
	Publisher    pubsub.Publisher      `inj:""`
	ValueStorage keyvalue.Value        `inj:""`

	logger *logrus.Entry
}

func NewTransactionsPublisher(logger *logrus.Entry) *TransactionsPublisher {
	return &TransactionsPublisher{
		logger: logger.WithField("app", "core/transactions_publisher"),
	}
}

func (tp *TransactionsPublisher) Start() error {
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

func (tp *TransactionsPublisher) Stop() {
	tp.logger.Debug("Stopping publisher...")
	tp.Publisher.Stop()
	tp.logger.Debug("Stopping blockchain...")
	tp.Blockchain.Stop()
}

func (tp *TransactionsPublisher) HandleBlock(block *blockchain.Block) {
	tp.logger.
		WithField("blockNum", block.Number).
		WithField("transactionsNum", len(block.Transactions)).
		Info("Block received")
	tp.ValueStorage.PutInt64(block.Number.Int64())
	for _, tx := range block.Transactions {
		data, err := proto.Marshal(tx)
		if err != nil {
			// TODO: Handle this
			continue
		}
		tp.Publisher.Publish("transactions", []byte(data))
	}
}
