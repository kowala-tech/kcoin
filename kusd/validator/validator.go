package validator

import (
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/vm"
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

	wg sync.WaitGroup
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

	// launch the state machine
	val.wg.Add(1)
	go val.run()

	// launch the consensus events handler
	go val.handle()
}

func (val *Validator) run() {
	log.Info("Starting the consensus state machine")
	for state, numTransitions := val.notYetVoterState, 0; state != nil; numTransitions++ {
		// @TODO(rgeraldes) - publish old/new state if necessary - need to review sync process
		state = state()
		if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
			break
		}
	}
}

func (val *Validator) handle() {
	for _, ev := range val.consensusEventSub.Chan() {
		switch data := ev.Data.(type) {
		case core.ConsensusEvent:
			switch msg := data.Msg.(type) {
			case *types.Proposal:
				val.setProposal(msg)
			}
		}
	}
}

func (val *Validator) Stop() {
	log.Info("Stopping consensus validator")

	val.leaveElections()
	val.wg.Wait()

	// quits handle
	val.consensusEventSub.Unsubscribe()

	//val.worker.stop()
	atomic.StoreInt32(&val.validating, 0)
	atomic.StoreInt32(&val.shouldStart, 0)

	log.Info("Consensus validator stopped")
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
func (val *Validator) Pending() (*types.Block, *state.StateDB) { return nil, nil }

// @TODO (rgeraldes) - same as Pending()
// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (val *Validator) PendingBlock() *types.Block {
	val.electionMu.Lock()
	defer val.electionMu.Unlock()

	/*
		if atomic.LoadInt32(&val.validating) == 0 {
			return types.NewBlock(
				val.current.header,
				val.current.txs,
				nil,
				val.current.receipts,
			)
		}
		return val.current.Block
	*/
	return nil
}

func (val *Validator) restoreLastCommit() {
	currentBlock := val.chain.CurrentBlock()
	if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
		return
	}

	// @TODO (rgeraldes) - call contract to get the validators
	lastValidators := core.NewValidatorSet([]*core.Validator{})

	lastCommit := currentBlock.LastCommit()
	lastPreCommits := core.NewVoteSet(currentBlock.Number(), lastCommit.Round(), types.PreCommit, lastValidators)
	for _, preCommit := range lastCommit.PreCommits() {
		if preCommit == nil {
			continue
		}

		/*
			added, err := lastPreCommits.Add(preCommit)
			if !added || error != nil {
				// @TODO (rgeraldes) - this should not happen > complete
				log.Error("Failed to restore the latest commit")
			}
		*/
	}

	val.lastCommit = lastPreCommits
}

func (val *Validator) init() {
	// @TODO (rgeraldes) - call pos contract in order to get the latest set of validators
	// for now it's an hardcoded value the set of validators will be the same as the last round
	// val.validators =
	// val.prevValidators =
	prevPreCommits := new(*core.VoteSet)
	if val.commitRound > -1 && val.votes != nil {
		prevPreCommits = val.votes.PreCommits(val.commitRound)
	}
	parent := val.chain.CurrentBlock()

	val.blockNumber = parent.Number().Add(parent.Number(), big.NewInt(1))
	val.round = 0
	val.start = val.end.Add(time.Duration(params.SyncDuration) * time.Millisecond)
	val.proposal = nil
	val.proposalBlock = nil
	//val.proposalBlockChunks = nil
	val.lockedRound = 0
	val.lockedBlock = nil
	val.commitRound = -1
	val.prevCommit = prevPreCommits
}

func (val *Validator) isProposer() bool {
	//return val.coinbase ==
	return false
}

func (val *Validator) setProposal(proposal *types.Proposal) {
	// not relevant
	if proposal.Height != val.Height && proposal.Round != val.Round {
		return
	}

	// proposer sent two proposals
	if val.proposal != nil {
		// @TODO (rgeraldes) - punish the proposer ?
		return
	}

	// if the proposal is already known, discard it
	hash := proposal.Hash()
	if val.all[hash] != nil {
		log.Trace("Discarding already known proposal", "hash", hash)
		return false, fmt.Errorf("known transaction: %x", hash)
	}

	// if the proposal fails validation, discard it
	if err := val.validateProposal(proposal); err != nil {
		log.Trace("Discarding invalid proposal", "hash", hash, "err", err)
	}
	val.proposal = proposal

	return
}

func (val *Validator) processBlockChunk() {}

func (val *Validator) commitTransactions(mux *event.TypeMux, txs *types.TransactionsByPriceAndNonce, bc *core.BlockChain, coinbase common.Address) {
	gp := new(core.GasPool).AddGas(val.header.GasLimit)
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
		from, _ := types.Sender(val.signer, tx)
		// Start executing the transaction
		val.state.Prepare(tx.Hash(), common.Hash{}, val.tcount)
		err, logs := val.commitTransaction(tx, bc, coinbase, gp)
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
			val.tcount++
			txs.Shift()
		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
		}
	}
	if len(coalescedLogs) > 0 || val.tcount > 0 {
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

func (val *Validator) commitTransaction(tx *types.Transaction, bc *core.BlockChain, coinbase common.Address, gp *core.GasPool) (error, []*types.Log) {
	snap := val.state.Snapshot()
	receipt, _, err := core.ApplyTransaction(val.config, bc, &coinbase, gp, val.state, val.header, tx, val.header.GasUsed, vm.Config{})
	if err != nil {
		val.state.RevertToSnapshot(snap)
		return err, nil
	}
	val.txs = append(val.txs, tx)
	val.receipts = append(val.receipts, receipt)
	return nil, receipt.Logs
}

func (val *Validator) joinElections() {
	val.makeDeposit()
}

func (val *Validator) leaveElections() {
	val.withdrawDeposit()
}

func (val *Validator) makeDeposit() {
	// @TODO (rgeraldes) - to complete as soon as the dynamic validator set contract is ready
}

func (val *Validator) withdrawDeposit() {
	// @TODO (rgeraldes) - to complete as soon as the dynamic validator set contract is ready
}

func (val *Validator) propose() {
	var block *types.Block
	if val.lockedBlock != nil {
		block = val.block
	} else {
		block = val.newBlock()
	}
	
	// new proposal
	lockedRound, lockedBlock := val.votes.LockingInfo()
	proposal := types.NewProposal(val.height, val.round, block.Fragment().Metadata(), lockedRound, lockedBlock)
	//if types.SignProposal(proposal, val.signer, )
	
	// sign proposal
	if err := ; err != nil {
		log.Error("proposal: error signing proposal")
		return
	}

	// update state
	val.proposal = proposal
	val.block = block
	// post proposal event
	val.events.Post(core.NewProposalEvent{Proposal: proposal})
	// post block segments events
	for i := 0; i < SegmentedBlock.NumSegments(); i ++ {
		val.events.Post(core.NewBlockSegmentEvent{val.Height, val.Round, SegmentedBlock.GetSegment(i)})
	}
}

func (val *Validator) newBlock() *types.Block {
	// new block header
	parent := val.chain.CurrentBlock()
	state, err := val.chain.StateAt(parent.Root())
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
		commit = val.lastCommit.Proof()
	}
	pending, err := val.kusd.TxPool().Pending()
	if err != nil {
		log.Error("Failed to fetch pending transactions", "err", err)
		return
	}
	txs := types.NewTransactionsByPriceAndNonce(val.signer, pending)
	val.commitTransactions(val.eventMux, txs, val.chain, val.coinbase)
	// accumulate any block rewards and commit the final state root
	accumulateRewards(val.config, state, header)
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

func (val *Validator) vote(typ types.VoteType, block common.Hash) {
	val.eventMuxPost(core.NewVoteEvent{Vote: val.newSignedVote(typ, block)})
}

func (val *Validator) prevote() {
	var vote common.Hash
	switch {
	// prevote locked block
	case val.lockedBlock != nil:
		log.Debug("")
		vote = val.lockedBlock.Hash()
	// proposal's block is nil, prevote nil
	case val.block == nil:
		log.Debug("")
		vote = nil
	// pre vote the proposal's block
	default:
		log.Debug("")
		vote = val.block.Hash()
	}
	val.vote(types.PreVote, v.block.Hash())
}

// preCommit votes on a candidate for the pre commit election
func (val *Validator) precommit() stateFn {
	var vote common.Hash // nil by default
	// access prevotes
	winner := common.Hash{}
	switch {
		// no majority
		case !val.hasPolka():
		// majority pre-voted nil
		case winner == nil:
			log.Debug("majority of validators pre-voted nil")
			// unlock locked block
			if val.lockedBlock != nil {
				val.lockedRound = 0
				val.lockedBlock = nil
			}
		// majority pre-voted the locked block
		case winner == lockedBlock.Hash():
			log.Debug("majority of validators pre-voted the locked block")
			// update locked block round
			val.lockedRound = val.round
			// vote on the pre-vote election winner
			vote = winner
		// majority pre-voted the proposed block
		case winner == block.Hash():
			log.Debug("majority of validators pre-voted the proposed block")
			// lock block
			val.lockedRound = round
			val.lockedBlock = cs.ProposalBlock
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
	val.vote(PreCommit, vote)
}