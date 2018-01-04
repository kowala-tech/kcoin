package validator

import (
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
	"github.com/tendermint/tendermint/types"
)

// Backend wraps all methods required for elections.
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
}

// consensus encapsulates the consensus state during the state machine execution
type consensus struct {
	// basic info
	blockNumber *big.Int
	round       int

	// time
	start time.Time
	end   time.Time

	// current election
	proposal      *types.Proposal
	proposalBlock *types.Block
	//proposalBlockFragments
	votes

	validators types.Validators

	// locking mechanism
	lockedRound int
	lockedBlock *types.Block

	// last commit
	commitRound int
	prevCommit  *core.VoteSet

	// proposer
	header   *types.Header
	txs      []*types.Transaction
	receipts []*types.Receipt
}

// Validator represents a consensus validator
type Validator struct {
	// internal consensus state
	consensus

	code common.Address

	// basic info
	validating int32
	account    *account.Account
	deposit    int
	wallet     Wallet // sign proposals & votes based on the sender account

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync

	// state machine
	stateMu        sync.Mutex
	stateMachine   state
	maxTransitions int // max number of state transitions (tests) 0 - unlimited

	// blockchain & misc
	config *params.ChainConfig
	chain  *core.BlockChain
	kusd   Backend

	state *state.StateDB // apply state changes here

	// consensus events
	// state machine inputs
	eventMux *event.TypeMux
	//preVoteElectionSub event.Subscription
	//preCommitElectionSub	event.Subscription
	//voteCountDoneSub event.Subscription
	//proposalSub event.Subscription
}

// New returns a new validator object
func New(kusd Backend, config *params.ChainConfig, eventMux *event.TypeMux) *Validator {
	validator := &Validator{
		kusd:     backend,
		config:   config,
		mux:      eventMux,
		signer:   types.NewAndromedaSigner(config.ChainID()),
		chain:    kusd.BlockChain(),
		canStart: 1,
	}

	// ensure that validator is synced (one shot type of update loop)
	go validator.sync()

	return validator
}

// Start
func (self *Validator) Start(coinbase common.Address, deposit int) {
	atomic.StoreInt32(&self.shouldStart, 1)

	if atomic.LoadInt32(&self.canStart) == 0 {
		log.Info("Network syncing, will start validator afterwards")
		return
	}

	log.Info("Starting validation operation")

	atomic.StoreInt32(&self.validating, 1)

	go self.run()
	go self.handle()
}

/*
// WithMaxStateTransactions limits the validator state transactions
func (self *Validator) WithMaxStateTransactions(maxTransitions int) {
	self.maxTransitions = maxTransitions
}

// sync keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your validation operation for as long as the DOS continues.
func (self *Validator) sync() {
	events := self.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
out:
	for ev := range events.Chan() {
		switch ev.Data.(type) {
		case downloader.StartEvent:
			atomic.StoreInt32(&self.canStart, 0)
			if self.IsValidating() {
				self.Stop()
				atomic.StoreInt32(&self.shouldStart, 1)
				log.Info("Validation aborted due to sync")
			}
		case downloader.DoneEvent, downloader.FailedEvent:
			shouldStart := atomic.LoadInt32(&self.shouldStart) == 1

			atomic.StoreInt32(&self.canStart, 1)
			atomic.StoreInt32(&self.shouldStart, 0)
			if shouldStart {
				self.Start(self.coinbase, self.deposit)
			}
			// unsubscribe. we're only interested in this event once
			events.Unsubscribe()
			// stop immediately and ignore all further pending events
			break out
		}
	}
}

// run starts the validator's state machine
func (self *Validator) run() {
	log.Info("Starting the state machine")
	for state, numTransitions := self.InitialState, 0; state != nil; numTransitions++ {
		// @TODO(rgeraldes) - publish old state if necessary - need to review sync process
		state = state()
		if self.maxTransitions > 0 && numTransitions == self.maxNumTransitions {
			break
		}
	}
}

// handle handles consensus events which may cause state transitions
func (self *Validator) handle() {
	for _, ev := range self.consensusEventSub.Chan() {
		switch data := ev.Data.(type) {
		case core.ConsensusEvent:
			switch msg := data.Msg.(type) {
			case *types.Proposal:
				self.setProposal(msg)
			}
	}
}







/*




func (self *Validator) Stop() {
	log.Info("Stopping consensus validator")

	// @NOTE (rgeraldes) - state machine needs to be stopped first
	self.machine.Stop()

	// Wait for the state machine to come down.
	self.wg.Wait()

	// quits handle
	self.consensusEventSub.Unsubscribe()

	atomic.StoreInt32(&self.validating, 0)
	atomic.StoreInt32(&self.shouldStart, 0)

	log.Info("Kowala consensus validator stopped")
}


func (self *Validator) saveVote(vote *types.Vote) {

}

func (self *Validator) setProposal(proposal *types.Proposal) {
	// not relevant
	if proposal.Height != self.Height && proposal.Round != self.Round {
		return
	}

	// proposer sent two proposals
	if self.proposal != nil {
		// @TODO (rgeraldes) - punish the proposer
		return
	}

	// if the proposal is already known, discard it
	hash := proposal.Hash()
	if self.all[hash] != nil {
		log.Trace("Discarding already known proposal", "hash", hash)
		return false, fmt.Errorf("known transaction: %x", hash)
	}

	// if the proposal fails validation, discard it
	if err := self.validateProposal(proposal); err != nil {
		log.Trace("Discarding invalid proposal", "hash", hash, "err", err)
	}

	self.proposal = proposal

	return
}

func (self *Validator) processBlockChunk() {

}

func (self *Validator) restoreLastcommit() {
	currentBlock := self.chain.CurrentBlock()

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
}

// commitTransactions
func (self *Validator) commitTransactions(mux *event.TypeMux, txs *types.TransactionsByPriceAndNonce, bc *core.BlockChain, coinbase common.Address) {
	gp := new(core.GasPool).AddGas(self.header.GasLimit)

	var coalescedLogs []*types.Log

	for {
		// Retrieve the next transaction and abort if all done
		tx := txs.Peek()
		if tx == nil {
			break
		}
		// Error may be ignored here. The error has already been checked
		// during transaction acceptance is the transaction pool.
		//
		// We use the eip155 signer regardless of the current hf.
		from, _ := types.Sender(self.signer, tx)

		// Start executing the transaction
		self.state.Prepare(tx.Hash(), common.Hash{}, self.tcount)

		err, logs := self.commitTransaction(tx, bc, coinbase, gp)
		switch err {
		case core.ErrGasLimitReached:
			// Pop the current out-of-gas transaction without shifting in the next from the account
			log.Trace("Gas limit exceeded for current block", "sender", from)
			txs.Pop()

		case core.ErrNonceTooLow:
			// New head notification data race between the transaction pool and miner, shift
			log.Trace("Skipping transaction with low nonce", "sender", from, "nonce", tx.Nonce())
			txs.Shift()

		case core.ErrNonceTooHigh:
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Trace("Skipping account with hight nonce", "sender", from, "nonce", tx.Nonce())
			txs.Pop()

		case nil:
			// Everything ok, collect the logs and shift in the next transaction from the same account
			coalescedLogs = append(coalescedLogs, logs...)
			self.tcount++
			txs.Shift()

		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
		}
	}

	if len(coalescedLogs) > 0 || self.tcount > 0 {
		// make a copy, the state caches the logs and these logs get "upgraded" from pending to mined
		// logs by filling in the block hash when the block was mined by the local miner. This can
		// cause a race condition if a log was "upgraded" before the PendingLogsEvent is processed.
		cpy := make([]*types.Log, len(coalescedLogs))
		for i, l := range coalescedLogs {
			cpy[i] = new(types.Log)
			*cpy[i] = *l
		}
		go func(logs []*types.Log, tcount int) {
			if len(logs) > 0 {
				mux.Post(core.PendingLogsEvent{Logs: logs})
			}
			if tcount > 0 {
				mux.Post(core.PendingStateEvent{})
			}
		}(cpy, env.tcount)
	}
}

// commitTransaction
func (self *Validator) commitTransaction(tx *types.Transaction, bc *core.BlockChain, coinbase common.Address, gp *core.GasPool) (error, []*types.Log) {
	snap := self.state.Snapshot()

	receipt, _, err := core.ApplyTransaction(self.config, bc, &coinbase, gp, self.state, self.header, tx, self.header.GasUsed, vm.Config{})
	if err != nil {
		self.state.RevertToSnapshot(snap)
		return err, nil
	}
	self.txs = append(self.txs, tx)
	self.receipts = append(self.receipts, receipt)

	return nil, receipt.Logs
}

// accumulateRewards accumulates the rewards for the validators
func (self *Validator) accumulateRewards(config *params.ChainConfig, state *state.StateDB, header *types.Header) {
	// @NOTE (rgeraldes) - original value present in core/fees.go
	// @NOTE (rgeraldes) - calls Helio's contract
}

// init
func (self *Validator) init() {
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

}

func (self *Validator) deposit() {
	// @TODO(rgeraldes) - to complete
}

func (self *Validator) withdraw() {
	// @TODO(rgeraldes) - to complete
}

func (self *Validator) propose() {
	var block *types.Block
	if self.lockedBlock != nil {
		block = self.block
	} else {
		// create a new block
		block = self.newBlock()
	}

	// new proposal
	lockedRound, lockedBlock := self.votes.LockingInfo()
	proposal := types.NewProposal(m.height, m.round, block.Fragment().Metadata(), lockedRound, lockedBlock)
	if types.SignProposal(proposal, self.signer, )
	
	// sign proposal
	if err := ; err != nil {
		log.Error("proposal: error signing proposal")
		return
	}

	// update state
	m.proposal = proposal
	m.block = block

	// post proposal event
	m.events.Post(core.NewProposalEvent{Proposal: proposal})

	// post block segments events
	for i := 0; i < SegmentedBlock.NumSegments(); i ++ {
		m.events.Post(core.NewBlockSegmentEvent{m.Height, m.Round, SegmentedBlock.GetSegment(i)})
	}
}

func (self *Validator) newBlock() *types.Block {
	// new block header
	parent := self.chain.CurrentBlock()
	state, err := self.chain.StateAt(parent.Root())
	if err != nil {
		log.Error("Failed to fetch the current state", "err", err)
		return
	}
	blockNumber := parent.Number()
	tstart := time.Now()
	tstamp := tstart.Unix()
	if parent.Time().Cmp(new(big.Int).SetInt64(tstamp)) >= 0 {
		tstamp = parent.Time().Int64() + 1
	}

	header = &types.Header{
		ParentHash: parent.Hash(),
		Number:     blockNumber.Add(blockNumber, common.Big1),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Time:       big.NewInt(tstamp),
	}

	var commit *Commit
	if blockNumber.Cmp(big.NewInt(1)) == 0 {
		commit = &types.Commit{}
	} else {
		commit = self.lastCommit.Proof()
	}

	pending, err := self.eth.TxPool().Pending()
	if err != nil {
		log.Error("Failed to fetch pending transactions", "err", err)
		return
	}

	txs := types.NewTransactionsByPriceAndNonce(self.signer, pending)
	self.commitTransactions(self.eventMux, txs, self.chain, self.coinbase)

	// accumulate any block rewards and commit the final state root
	accumulateRewards(self.config, state, header)
	header.Root = state.IntermediateRoot(true)

	// create proposal block
	block := types.NewBlock(header, txs, receipts, commit)

	// Update the block hash in all logs since it is now available and not when the
	// receipt/log of individual transactions were created.
	for _, r := range work.receipts {
		for _, l := range r.Logs {
			l.BlockHash = block.Hash()
		}
	}
	for _, log := range work.state.Logs() {
		log.BlockHash = block.Hash()
	}

	return block
}

func (self *Validator) vote(typ types.VoteType, block common.hash) {
	self.eventMuxPost(core.NewVoteEvent{Vote: self.newSignedVote(typ, block)})
}

// prevote votes for a pre vote candidate
func (self *Validator) prevote() {
	var vote common.Hash

	switch {
	// prevote locked block
	case m.lockedBlock != nil:
		log.Debug("")
		vote = m.lockedBlock.Hash()

	// proposal's block is nil, prevote nil
	case m.block == nil:
		log.Debug("")
		vote = nil

	// pre vote the proposal's block
	default:
		log.Debug("")
		vote = m.block.Hash()
	}

	m.vote(types.PreVote, v.block.Hash())
}

// preCommit votes on a candidate for the pre commit election
func (self *Validator) precommit() stateFn {
	var vote common.Hash // nil by default

	// access prevotes
	winner := common.Hash{}

	switch {
		// no majority
		case !self.hasPolka():

		// majority pre-voted nil
		case winner == nil:
			log.Debug("majority of validators pre-voted nil")
			// unlock locked block
			if self.lockedBlock != nil {
				m.lockedRound = 0
				m.lockedBlock = nil
			}

		// majority pre-voted the locked block
		case winner == lockedBlock.Hash():
			log.Debug("majority of validators pre-voted the locked block")
			// update locked block round
			m.lockedRound = m.round

			// vote on the pre-vote election winner
			vote = winner

		// majority pre-voted the proposed block
		case winner == block.Hash():
			log.Debug("majority of validators pre-voted the proposed block")
			// lock block
			m.lockedRound = round
			m.lockedBlock = cs.ProposalBlock

			// vote on the pre-vote election winner
			vote = winner

		// we don't have the current block (fetch)
		// @TODO (tendermint): in the future save the POL prevotes for justification.
		// fetch block, unlock, precommit
		default:
			// unlock locked block
			cs.LockedRound = 0
			cs.LockedBlock = nil
			cs.LockedBlockParts = nil
			if !cs.ProposalBlockParts.HasHeader(blockID.PartsHeader) {
				cs.ProposalBlock = nil
				cs.ProposalBlockParts = types.NewPartSetFromHeader(blockID.PartsHeader)
			}

		}

	m.vote(PreCommit, vote)
}