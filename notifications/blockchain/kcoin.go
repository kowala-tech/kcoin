package blockchain

import (
	"math/big"
	"time"

	"context"

	kcoinLib "github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/kcoinclient"
	"github.com/kowala-tech/kcoin/client/rpc"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/sirupsen/logrus"
)

type kcoin struct {
	rpcAddr         string
	pollingInterval time.Duration
	logger          *logrus.Entry

	rpcClient *rpc.Client
	client    *kcoinclient.Client
	ctx       context.Context
	ctxCancel context.CancelFunc
	closedCh  chan interface{}
	handlers  map[BlockHandler]struct{}

	latestBlock *big.Int
}

func NewKcoin(rpcAddr string, pollingIntervalSeconds int, logger *logrus.Entry) Blockchain {
	return &kcoin{
		rpcAddr:         rpcAddr,
		pollingInterval: time.Duration(pollingIntervalSeconds) * time.Second,
		closedCh:        make(chan interface{}),
		handlers:        map[BlockHandler]struct{}{},
		logger:          logger.WithField("app", "blockchain/kcoin"),
	}
}

func (k *kcoin) Start() error {
	k.logger.Debug("Starting...")
	rpcClient, err := rpc.Dial(k.rpcAddr)
	if err != nil {
		k.logger.WithError(err).Error("Error dialing to the RPC address")
		return err
	}
	k.rpcClient = rpcClient
	k.client = kcoinclient.NewClient(rpcClient)
	k.ctx, k.ctxCancel = context.WithCancel(context.Background())

	if k.latestBlock == nil {
		k.logger.Debug("Setting starting block to first")
		block, err := k.getBlock(big.NewInt(0))
		if err != nil {
			return err
		}
		k.latestBlock = block.Number()
		k.logger.WithField("blockNum", block.Number).Info("Starting block set to latest")
	}

	k.pollingLoop()

	return nil
}

func (k *kcoin) Stop() {
	k.ctxCancel() // Stop fetching more new blocks.
	select {
	case <-k.closedCh: // Waits for the loop to finish
	case <-time.After(k.pollingInterval * 3):
	}

	k.rpcClient.Close()
}

func (k *kcoin) Seek(blockNumber *big.Int) error {
	k.latestBlock = blockNumber
	return nil
}

func (k *kcoin) OnBlock(handler BlockHandler) error {
	k.handlers[handler] = struct{}{}
	return nil
}

func (k *kcoin) getBlock(blockNumber *big.Int) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	return k.client.BlockByNumber(ctx, blockNumber)
}

func (k *kcoin) pollingLoop() {
	k.logger.Debug("Running main loop...")

	for {
		select {
		case <-k.ctx.Done():
			// Close() has been called. End the infinite loop
			k.closedCh <- true
			k.logger.Debug("Ending main loop...")
			return
		default:
		}
		rawBlock, err := k.getBlock(k.latestBlock)
		if err != nil {
			if err == kcoinLib.NotFound {
				k.logger.Debug("No new block found")
				time.Sleep(k.pollingInterval)
			} else {
				k.logger.WithError(err).Error("Error fetching new block")
				time.Sleep(k.pollingInterval)
			}
		} else {
			k.logger.WithField("blockNum", rawBlock.Number().Int64()).Info("New block found")

			block := k.wrapBlock(rawBlock)

			for handler := range k.handlers {
				handler.HandleBlock(block)
			}

			k.latestBlock.Add(k.latestBlock, common.Big1)
		}
	}
}

func (k *kcoin) wrapBlock(block *types.Block) *Block {
	inTransactions := block.Transactions()
	transactions := make([]*protocolbuffer.Transaction, len(inTransactions))
	for i, tx := range inTransactions {
		to := "0x0"

		isContractCreationTx := tx.To() == nil
		if !isContractCreationTx {
			to = tx.To().String()
		}

		from, err := tx.From()
		if err != nil {
			return nil
		}

		transactions[i] = &protocolbuffer.Transaction{
			To:          to,
			From:        from.String(),
			Amount:      tx.Value().Int64(),
			Hash:        tx.Hash().String(),
			Timestamp:   block.Time().Int64(),
			GasUsed:     int64(block.GasUsed()),
			GasPrice:    tx.GasPrice().Int64(),
			BlockHeight: block.Number().Int64(),
		}
	}
	return &Block{
		Number:       block.Number(),
		Transactions: transactions,
	}
}
