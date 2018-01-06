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
	// state machine
	electionMu     sync.Mutex
	election           // consensus state
	maxTransitions int // max number of state transitions (tests) 0 - unlimited

	validating int32
	coinbase   common.Address
	deposit    uint64

	// blockchain
	kusd   Backend
	chain  *core.BlockChain
	config *params.ChainConfig
	engine consensus.Engine

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync

	// events
	eventMux *event.TypeMux
}

// New returns a new consensus validator
func New(kusd Backend, config *params.ChainConfig, eventMux *event.TypeMux, engine consensus.Engine) *Validator {
	return &Validator{
		kusd:     kusd,
		chain:    kusd.BlockChain(),
		config:   config,
		eventMux: eventMux,
		engine:   engine,
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

	go val.run()
	// go val.handle()
}

func (val *Validator) run() {
	log.Info("Starting the consensus state machine")
	for state, numTransitions := val.initialState, 0; state != nil; numTransitions++ {
		// @TODO(rgeraldes) - publish old/new state if necessary - need to review sync process
		state = state()
		if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
			break
		}
	}
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

func (val *Validator) restoreLastCommit() {
	/*
		currentBlock := val.chain.CurrentBlock()
		if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
			return
		}

		lastCommit := currentBlock.Commit()
		lastPreCommits := core.NewVoteSet(currentBlock.Number(), lastCommit.Round(), types.PreCommit, state.lastValidators)
		for _, preCommit := range lastCommit.PreCommits() {
			if preCommit == nil {
				continue
			}
			added, err := lastPreCommits.Add(preCommit)
			if !added || error != nil {
				// @TODO (rgeraldes) - this should not happen > complete
				log.Error("Failed to restore the latest commit")
			}
		}
		self.lastCommit = lastPrecommits
	*/
}

func (val *Validator) init() {
	/*
		// @TODO (rgeraldes) - call pos contract in order to get the latest set of validators
		// for now it's an hardcoded value the set of validators will be the same as the last round
		// self.validators =
		// self.prevValidators =
		prevPreCommits := new(*core.VoteSet)
		if self.commitRound > -1 && self.votes != nil {
			prevPreCommits = self.votes.PreCommits(self.commitRound)
		}
		parent := self.chain.CurrentBlock()

		self.blockNumber = parent.BlockNumber().Add(1)
		self.round = 0
		self.start = self.end.Add(SyncDuration * time.Millisecond)
		self.proposal = nil
		self.proposalBlock = nil
		self.proposalBlockChunks = nil
		self.lockedRound = 0
		self.lockedBlock = nil
		self.commitRound = -1
		self.prevCommit = prevPreCommits

	*/
}

func (val *Validator) isProposer() bool {
	//return val.coinbase ==
	return false
}
