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
	"github.com/kowala-tech/kUSD/params"
)

// Backend wraps all methods required for mining.
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() kusddb.Database
}

type Validator struct {
	consensusMu sync.Mutex

	validating int32
	coinbase   common.Address
}

func New(kusd Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine) *Validator {
	return &Validator{}
}

func (self *Validator) Start(coinbase common.Address) {}
func (self *Validator) Stop()                         {}

func (self *Validator) SetExtra(extra []byte) error { return nil }
func (self *Validator) Validating() bool {
	return atomic.LoadInt32(&self.validating) > 0
}

func (self *Validator) SetCoinbase(addr common.Address) {
	self.coinbase = addr
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
	self.consensusMu.Lock()
	defer self.consensusMu.Unlock()

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
