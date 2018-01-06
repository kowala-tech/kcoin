package validator

import (
	"sync"
	"sync/atomic"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusddb"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

// Backend wraps all methods required for mining.
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() kusddb.Database
}

// Validator represents a consensus validator
type Validator struct {
	electionMu sync.Mutex
	election   // consensus state

	validating int32
	coinbase   common.Address
	deposit    uint64

	kusd Backend

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync
}

func New(kusd Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine) *Validator {
	return &Validator{
		kusd: kusd,
	}
}

func (val *Validator) Start(coinbase common.Address, deposit uint64) {
	atomic.StoreInt32(&val.shouldStart, 1)
	val.coinbase = coinbase
	val.deposit = deposit

	if atomic.LoadInt32(&val.canStart) == 0 {
		log.Info("Network syncing, will start validator afterwards")
		return
	}
	atomic.StoreInt32(&val.validating, 1)

	// go val.handle()
	// go val.run()
}

func (val *Validator) Stop() {
	//val.worker.stop()
	atomic.StoreInt32(&val.validating, 0)
	atomic.StoreInt32(&val.shouldStart, 0)
}

func (val *Validator) SetExtra(extra []byte) error { return nil }

func (val *Validator) Validating() bool {
	return atomic.LoadInt32(&val.validating) > 0
}

func (val *Validator) SetCoinbase(addr common.Address) {
	val.coinbase = addr
}

func (val *Validator) SetDeposit(deposit uint64) {
	val.deposit = deposit
}

// @TODO (rgeraldes) - not sure if pending makes much sense in pos context with low network latencies. review
// Pending returns the currently pending block and associated state.
func (self *Validator) Pending() (*types.Block, *state.StateDB) { return nil, nil }

// @TODO (rgeraldes) - same as Pending()
// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (self *Validator) PendingBlock() *types.Block {
	self.electionMu.Lock()
	defer self.electionMu.Unlock()

	/*
		if atomic.LoadInt32(&self.validating) == 0 {
			return types.NewBlock(
				self.current.header,
				self.current.txs,
				nil,
				self.current.receipts,
			)
		}
		return self.current.Block
	*/
	return nil
}
