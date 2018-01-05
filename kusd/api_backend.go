package kusd

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/math"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/ethdb"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusd/downloader"
	"github.com/kowala-tech/kUSD/kusd/gasprice"
	"github.com/kowala-tech/kUSD/params"
	"github.com/kowala-tech/kUSD/rpc"
)

// KowalaApiBackend implements ethapi.Backend for full nodes
type KowalaApiBackend struct {
	kusd *Kowala
	gpo  *gasprice.Oracle
}

func (b *KowalaApiBackend) ChainConfig() *params.ChainConfig {
	return b.kusd.chainConfig
}

func (b *KowalaApiBackend) CurrentBlock() *types.Block {
	return b.kusd.blockchain.CurrentBlock()
}

func (b *KowalaApiBackend) SetHead(number uint64) {
	b.kusd.protocolManager.downloader.Cancel()
	b.kusd.blockchain.SetHead(number)
}

func (b *KowalaApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.kusd.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.kusd.blockchain.CurrentBlock().Header(), nil
	}
	return b.kusd.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *KowalaApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.kusd.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.kusd.blockchain.CurrentBlock(), nil
	}
	return b.kusd.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *KowalaApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.kusd.miner.Pending()
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.kusd.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *KowalaApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.kusd.blockchain.GetBlockByHash(blockHash), nil
}

func (b *KowalaApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.kusd.chainDb, blockHash, core.GetBlockNumber(b.kusd.chainDb, blockHash)), nil
}

func (b *KowalaApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.kusd.blockchain.GetTdByHash(blockHash)
}

func (b *KowalaApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.kusd.BlockChain(), nil)
	return vm.NewEVM(context, state, b.kusd.chainConfig, vmCfg), vmError, nil
}

func (b *KowalaApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.kusd.txPool.AddLocal(signedTx)
}

func (b *KowalaApiBackend) RemoveTx(txHash common.Hash) {
	b.kusd.txPool.Remove(txHash)
}

func (b *KowalaApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.kusd.txPool.Pending()
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
	return b.kusd.txPool.Get(hash)
}

func (b *KowalaApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.kusd.txPool.State().GetNonce(addr), nil
}

func (b *KowalaApiBackend) Stats() (pending int, queued int) {
	return b.kusd.txPool.Stats()
}

func (b *KowalaApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.kusd.TxPool().Content()
}

func (b *KowalaApiBackend) Downloader() *downloader.Downloader {
	return b.kusd.Downloader()
}

func (b *KowalaApiBackend) ProtocolVersion() int {
	return b.kusd.EthVersion()
}

func (b *KowalaApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *KowalaApiBackend) ChainDb() ethdb.Database {
	return b.kusd.ChainDb()
}

func (b *KowalaApiBackend) EventMux() *event.TypeMux {
	return b.kusd.EventMux()
}

func (b *KowalaApiBackend) AccountManager() *accounts.Manager {
	return b.kusd.AccountManager()
}
