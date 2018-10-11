package knode

import (
	"context"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/math"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/bloombits"
	"github.com/kowala-tech/kcoin/client/core/rawdb"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/kcoindb"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/rpc"
	"github.com/kowala-tech/kcoin/client/services/knode/downloader"
)

// KowalaAPIBackend implements kcoinapi.Backend for full nodes
type KowalaAPIBackend struct {
	kcoin *Kowala
}

// ChainConfig returns the active chain configuration.
func (b *KowalaAPIBackend) ChainConfig() *params.ChainConfig {
	return b.kcoin.chainConfig
}

func (b *KowalaAPIBackend) CurrentBlock() *types.Block {
	return b.kcoin.blockchain.CurrentBlock()
}

func (b *KowalaAPIBackend) SetHead(number uint64) {
	b.kcoin.protocolManager.downloader.Cancel()
	b.kcoin.blockchain.SetHead(number)
}

func (b *KowalaAPIBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.kcoin.blockchain.CurrentBlock().Header(), nil
	}

	return b.kcoin.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *KowalaAPIBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.kcoin.blockchain.CurrentBlock(), nil
	}
	return b.kcoin.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *KowalaAPIBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.kcoin.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *KowalaAPIBackend) GetBlock(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.kcoin.blockchain.GetBlockByHash(hash), nil
}

func (b *KowalaAPIBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	if number := rawdb.ReadHeaderNumber(b.kcoin.chainDb, hash); number != nil {
		return rawdb.ReadReceipts(b.kcoin.chainDb, hash, *number), nil
	}
	return nil, nil
}

func (b *KowalaAPIBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	number := rawdb.ReadHeaderNumber(b.kcoin.chainDb, hash)
	if number == nil {
		return nil, nil
	}
	receipts := rawdb.ReadReceipts(b.kcoin.chainDb, hash, *number)
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (b *KowalaAPIBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.kcoin.BlockChain(), nil)
	return vm.NewEVM(context, state, b.kcoin.chainConfig, vmCfg), vmError, nil
}

func (b *KowalaAPIBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *KowalaAPIBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeChainEvent(ch)
}

func (b *KowalaAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *KowalaAPIBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *KowalaAPIBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.kcoin.BlockChain().SubscribeLogsEvent(ch)
}

func (b *KowalaAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.kcoin.txPool.AddLocal(signedTx)
}

func (b *KowalaAPIBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.kcoin.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *KowalaAPIBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.kcoin.txPool.Get(hash)
}

func (b *KowalaAPIBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.kcoin.txPool.State().GetNonce(addr), nil
}

func (b *KowalaAPIBackend) Stats() (pending int, queued int) {
	return b.kcoin.txPool.Stats()
}

func (b *KowalaAPIBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.kcoin.TxPool().Content()
}

func (b *KowalaAPIBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.kcoin.TxPool().SubscribeNewTxsEvent(ch)
}

func (b *KowalaAPIBackend) Downloader() *downloader.Downloader {
	return b.kcoin.Downloader()
}

func (b *KowalaAPIBackend) ProtocolVersion() int {
	return b.kcoin.EthVersion()
}

func (b *KowalaAPIBackend) ChainDb() kcoindb.Database {
	return b.kcoin.ChainDb()
}

func (b *KowalaAPIBackend) EventMux() *event.TypeMux {
	return b.kcoin.EventMux()
}

func (b *KowalaAPIBackend) AccountManager() *accounts.Manager {
	return b.kcoin.AccountManager()
}

func (b *KowalaAPIBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.kcoin.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *KowalaAPIBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.kcoin.bloomRequests)
	}
}
