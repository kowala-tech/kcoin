package knode

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/common/math"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/bloombits"
	"github.com/kowala-tech/kcoin/core/state"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/core/vm"
	"github.com/kowala-tech/kcoin/event"
	"github.com/kowala-tech/kcoin/kcoindb"
	"github.com/kowala-tech/kcoin/knode/downloader"
	"github.com/kowala-tech/kcoin/knode/gasprice"
	"github.com/kowala-tech/kcoin/params"
	"github.com/kowala-tech/kcoin/rpc"
)

// KowalaApiBackend implements kcoinapi.Backend for full nodes
type KowalaApiBackend struct {
	kcoin *Kowala
	gpo   *gasprice.Oracle
}

func (b *KowalaApiBackend) ChainConfig() *params.ChainConfig {
	return b.kcoin.chainConfig
}

func (b *KowalaApiBackend) CurrentBlock() *types.Block {
	return b.kcoin.blockchain.CurrentBlock()
}

func (b *KowalaApiBackend) SetHead(number uint64) {
	b.kcoin.protocolManager.downloader.Cancel()
	b.kcoin.blockchain.SetHead(number)
}

func (b *KowalaApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the validator
	if blockNr == rpc.PendingBlockNumber {
		block := b.kcoin.validator.PendingBlock()
		return block.Header(), nil
	}

	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.kcoin.blockchain.CurrentBlock().Header(), nil
	}

	return b.kcoin.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *KowalaApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the validator
	if blockNr == rpc.PendingBlockNumber {
		block := b.kcoin.validator.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.kcoin.blockchain.CurrentBlock(), nil
	}
	return b.kcoin.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *KowalaApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the validator
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.kcoin.validator.Pending()
		return state, block.Header(), nil
	}

	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.kcoin.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *KowalaApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.kcoin.blockchain.GetBlockByHash(blockHash), nil
}

func (b *KowalaApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.kcoin.chainDb, blockHash, core.GetBlockNumber(b.kcoin.chainDb, blockHash)), nil
}

func (b *KowalaApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.kcoin.BlockChain(), nil)
	return vm.NewEVM(context, state, b.kcoin.chainConfig, vmCfg), vmError, nil
}

func (b *KowalaApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *KowalaApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeChainEvent(ch)
}

func (b *KowalaApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *KowalaApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.kcoin.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *KowalaApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.kcoin.BlockChain().SubscribeLogsEvent(ch)
}

func (b *KowalaApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.kcoin.txPool.AddLocal(signedTx)
}

func (b *KowalaApiBackend) GetPoolTransactions() (types.Transactions, error) {
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

func (b *KowalaApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.kcoin.txPool.Get(hash)
}

func (b *KowalaApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.kcoin.txPool.State().GetNonce(addr), nil
}

func (b *KowalaApiBackend) Stats() (pending int, queued int) {
	return b.kcoin.txPool.Stats()
}

func (b *KowalaApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.kcoin.TxPool().Content()
}

func (b *KowalaApiBackend) SubscribeTxPreEvent(ch chan<- core.TxPreEvent) event.Subscription {
	return b.kcoin.TxPool().SubscribeTxPreEvent(ch)
}

func (b *KowalaApiBackend) Downloader() *downloader.Downloader {
	return b.kcoin.Downloader()
}

func (b *KowalaApiBackend) ProtocolVersion() int {
	return b.kcoin.EthVersion()
}

func (b *KowalaApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *KowalaApiBackend) ChainDb() kcoindb.Database {
	return b.kcoin.ChainDb()
}

func (b *KowalaApiBackend) EventMux() *event.TypeMux {
	return b.kcoin.EventMux()
}

func (b *KowalaApiBackend) AccountManager() *accounts.Manager {
	return b.kcoin.AccountManager()
}

func (b *KowalaApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.kcoin.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *KowalaApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.kcoin.bloomRequests)
	}
}
