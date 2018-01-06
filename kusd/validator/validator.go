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

// election encapsulates the consensus election round state
type election struct {
}

type Validator struct {
	electionMu sync.Mutex
	election   // consensus state

	validating int32
	coinbase   common.Address
	deposit    int

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
