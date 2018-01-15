package validator

import (
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/contracts/voters/contract"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusd/downloader"
	"github.com/kowala-tech/kUSD/kusddb"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

// @TODO (rgeraldes) - protect concurrent accesses

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
	Election           // consensus state
	maxTransitions int // max number of state transitions (tests) 0 - unlimited

	validating int32
	deposit    uint64

	// blockchain
	kusd   Backend
	chain  *core.BlockChain
	config *params.ChainConfig
	engine consensus.Engine

	// registry
	registry *contract.VoterRegistry

	// wallet (signer)
	account accounts.Account
	wallet  accounts.Wallet

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync

	// events
	eventMux *event.TypeMux

	wg sync.WaitGroup
}

// New returns a new consensus validator
func New(kusd Backend, config *params.ChainConfig, eventMux *event.TypeMux, engine consensus.Engine) *Validator {
	validator := &Validator{
		config:   config,
		kusd:     kusd,
		chain:    kusd.BlockChain(),
		engine:   engine,
		eventMux: eventMux,
	}

	// setup the registry
	// @TODO (rgeraldes) - complete
	/*
		registry, err := contract.NewVoterRegistry()
		if err != nil {
			// @TODO (rgeraldes) -
		}
		validator.registry = registry
	*/

	//go validator.sync()

	return validator
}

// sync keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your validation operation for as long as the DOS continues.
func (val *Validator) sync() {
	events := val.eventMux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
out:
	for ev := range events.Chan() {
		switch ev.Data.(type) {
		case downloader.StartEvent:
			atomic.StoreInt32(&val.canStart, 0)
			if val.Validating() {
				val.Stop()
				atomic.StoreInt32(&val.shouldStart, 1)
				log.Info("Validation aborted due to sync")
			}
		case downloader.DoneEvent, downloader.FailedEvent:
			shouldStart := atomic.LoadInt32(&val.shouldStart) == 1
			atomic.StoreInt32(&val.canStart, 1)
			atomic.StoreInt32(&val.shouldStart, 0)
			if shouldStart {
				val.Start(val.account.Address, val.deposit)
			}
			// unsubscribe. we're only interested in this event once
			events.Unsubscribe()
			// stop immediately and ignore all further pending events
			break out
		}
	}
}

// @TODO (rgeraldes) - add logic to prevent re-running the state machine
func (val *Validator) Start(coinbase common.Address, deposit uint64) {
	// @TODO (rgeraldes) - review logic (start is called by sync later on if the node is syncing)
	account := accounts.Account{Address: coinbase}
	wallet, err := val.kusd.AccountManager().Find(account)
	if err != nil {
		log.Crit("Failed to find a wallet", "err", err, "account", account.Address)
	}

	atomic.StoreInt32(&val.shouldStart, 1)

	val.wallet = wallet
	val.account = account
	val.deposit = deposit

	/*
		if atomic.LoadInt32(&val.canStart) == 0 {
			log.Info("Network syncing, will start validator afterwards")
			return
		}
	*/

	// val.joinElection()

	//if joined := val.joinElections; !joined {
	//log.Error("Failed to register validator")
	//}

	log.Info("Starting validation operation")
	atomic.StoreInt32(&val.validating, 1)

	// launch the state machine
	go val.run()
}

func (val *Validator) run() {
	val.wg.Add(1)
	defer val.wg.Done()

	log.Info("Starting the consensus state machine")
	for state, numTransitions := val.notLoggedInState, 0; state != nil; numTransitions++ {
		// @TODO(rgeraldes) - publish old/new state if necessary - need to review sync process
		state = state()
		if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
			break
		}
	}
	log.Info("Stopped Consensus state machine")
}

func (val *Validator) Stop() {
	log.Info("Stopping consensus validator")

	// val.leaveElection()
	val.wg.Wait()

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
	val.account.Address = addr
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

	// @TODO (rgeraldes) - VALIDATORS CONTRACT
	lastValidators := &types.Validators{}

	lastCommit := currentBlock.LastCommit()
	lastPreCommits := core.NewVotingTable(currentBlock.Number(), lastCommit.Round(), types.PreCommit, lastValidators)
	for _, preCommit := range lastCommit.Commits() {
		if preCommit == nil {
			continue
		}
		added, err := lastPreCommits.Add(preCommit)
		if !added || err != nil {
			// @TODO (rgeraldes) - this should not happen > complete
			log.Error("Failed to restore the latest commit")
		}
	}

	val.lastCommit = lastPreCommits
}

func (val *Validator) init() {
	parent := val.chain.CurrentBlock()

	val.blockNumber = parent.Number().Add(parent.Number(), big.NewInt(1))
	val.round = 0

	// @NOTE (rgeraldes) - in order to sync the nodes, the start time
	// must be the timestamp on the block + a sync interval
	// Tendermint uses a different logic that does not rely on the
	// previous information, but I think that's something necessary.
	val.start = time.Unix(parent.Time().Int64(), 0).Add(time.Duration(params.SyncDuration) * time.Millisecond)
	val.proposal = nil
	val.block = nil
	val.blockFragments = nil

	val.lockedRound = 0
	val.lockedBlock = nil
	val.commitRound = -1

	val.validators = &types.Validators{}

	// val.votes
	// @TODO (rgeraldes) - VALIDATORS CONTRACT

	// val.lastValidators

	// val.prevValidators =
	//prevPreCommits := new(*core.VoteSet)
	//if val.commitRound > -1 && val.votes != nil {
	//	prevPreCommits = val.votes.PreCommits(val.commitRound)
	//}
	// val.prevCommit = prevPreCommits

}

func (val *Validator) isProposer() bool {
	// @TODO (rgeraldes) - modify as soon as we access to the validator list
	//return val.validators.Proposer() == val.account.Address
	return true
}

func (val *Validator) AddProposal(proposal *types.Proposal) {
	// @TODO (rgeraldes) - Complete

	/*
		// not relevant
		if proposal.BlockNumber != val.blockNumber && proposal.Round != val.Round {
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


		return
	*/

	go func() { val.proposalCh <- proposal }()
}

func (val *Validator) AddVote(vote *types.Vote) {
	if err := val.addVote(vote); err != nil {
		switch err {
		}
	}
}

func (val *Validator) addVote(vote *types.Vote) error {
	// @NOTE (rgeraldes) - for now just pre-vote/pre-commit for the current block number
	added, err := val.votes.Add(vote)
	if err != nil {
		// @TODO (rgeraldes)
	}

	if added {
		switch vote.Type {
		//case PreVote:
		//case PreCommit:
		}
	}

	return nil
}

func (val *Validator) ProcessBlockFragment() {}

func (val *Validator) commitTransactions(mux *event.TypeMux, txs *types.TransactionsByPriceAndNonce, bc *core.BlockChain, coinbase common.Address) {
	/*
		gp := new(core.GasPool).AddGas(env.header.GasLimit)

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
			from, _ := types.Sender(env.signer, tx)
			// Check whether the tx is replay protected. If we're not in the EIP155 hf
			// phase, start ignoring the sender until we do.
			if tx.Protected() && !env.config.IsEIP155(env.header.Number) {
				log.Trace("Ignoring reply protected transaction", "hash", tx.Hash(), "eip155", env.config.EIP155Block)

				txs.Pop()
				continue
			}
			// Start executing the transaction
			env.state.Prepare(tx.Hash(), common.Hash{}, env.tcount)

			err, logs := env.commitTransaction(tx, bc, coinbase, gp)
			switch err {
			case core.ErrGasLimitReached:
				// Pop the current out-of-gas transaction without shifting in the next from the account
				log.Trace("Gas limit exceeded for current block", "sender", from)
				txs.Pop()

			case nil:
				// Everything ok, collect the logs and shift in the next transaction from the same account
				coalescedLogs = append(coalescedLogs, logs...)
				env.tcount++
				txs.Shift()

			default:
				// Pop the current failed transaction without shifting in the next from the account
				log.Trace("Transaction failed, will be removed", "hash", tx.Hash(), "err", err)
				env.failedTxs = append(env.failedTxs, tx)
				txs.Pop()
			}
		}

		if len(coalescedLogs) > 0 || env.tcount > 0 {
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
	*/
}

func (val *Validator) commitTransaction(tx *types.Transaction, bc *core.BlockChain, coinbase common.Address, gp *core.GasPool) (error, []*types.Log) {
	/*
		snap := env.state.Snapshot()

		receipt, _, err := core.ApplyTransaction(env.config, bc, &coinbase, gp, env.state, env.header, tx, env.header.GasUsed, vm.Config{})
		if err != nil {
			env.state.RevertToSnapshot(snap)
			return err, nil
		}
		env.txs = append(env.txs, tx)
		env.receipts = append(env.receipts, receipt)
	*/
	//	return nil, receipt.Logs
	return nil, nil
}

func (val *Validator) joinElection() bool {
	voter, err := val.registry.IsVoter(&bind.CallOpts{}, val.account.Address)
	if err != nil {
		log.Error("Failed to verify if the validator is already registered as a voter")
		return false
	}

	if voter {
		log.Warn("Validator is already registered as a voter")
		return false
	}

	val.makeDeposit()
	return true
}

func (val *Validator) leaveElection() {
	val.withdraw()
}

// @TODO (rgeraldes) - review call opts
func (val *Validator) makeDeposit() {
	// @TODO (rgeraldes) - the initial deposit value should also be
	// validated against the minimum deposit
	min, err := val.registry.MinimumDeposit(&bind.CallOpts{})
	if err != nil {
		log.Crit("Failed to verify the minimum deposit")
		return
	}

	if min.Cmp(big.NewInt(int64(val.deposit))) > 0 {
		log.Warn("Current deposit is inferior than the minimum required", "deposit", val.deposit, "minimum required", min)
		return
	}

	// check if there are spots left to vote
	available, err := val.registry.Availability(&bind.CallOpts{})
	if err != nil {
		log.Crit("Failed to verify the registry availability")
		return
	}

	if available {
		opts := &bind.TransactOpts{}
		_, err := val.registry.Deposit(opts)
		if err != nil {
			log.Crit("Failed to make a deposit to participate in the election")
		}
	} else {
		log.Info("There are not positions available at the moment.")
	}

}

func (val *Validator) withdraw() {
	opts := &bind.TransactOpts{}
	_, err := val.registry.Withdraw(opts)
	if err != nil {
		log.Error("Failed to withdraw from the election")
	}
}

func (val *Validator) createProposalBlock() *types.Block {
	if val.lockedBlock != nil {
		log.Info("Picking a locked block")
		return val.lockedBlock
	}
	return val.createBlock()
}

func (val *Validator) createBlock() *types.Block {
	log.Info("Creating a new block")
	// new block header
	parent := val.chain.CurrentBlock()
	state, err := val.chain.StateAt(parent.Root())
	if err != nil {
		log.Crit("Failed to fetch the current state", "err", err)
	}
	blockNumber := parent.Number()
	tstart := time.Now()
	tstamp := tstart.Unix()
	if parent.Time().Cmp(new(big.Int).SetInt64(tstamp)) >= 0 {
		tstamp = parent.Time().Int64() + 1
	}
	header := &types.Header{
		ParentHash: parent.Hash(),
		Coinbase:   val.account.Address,
		Number:     blockNumber.Add(blockNumber, common.Big1),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Time:       big.NewInt(tstamp),
	}

	var commit *types.Commit

	// @NOTE (rgeraldes) - temporary
	first := types.NewVote(blockNumber, parent.Hash(), 0, types.PreCommit)

	if blockNumber.Cmp(big.NewInt(1)) == 0 {
		commit = &types.Commit{
			PreCommits:     types.Votes{first},
			FirstPreCommit: first,
		}
	} else {
		commit = &types.Commit{
			PreCommits:     types.Votes{first},
			FirstPreCommit: first,
		}
		//commit = val.lastCommit.Proof()
	}

	pending, err := val.kusd.TxPool().Pending()
	if err != nil {
		log.Crit("Failed to fetch pending transactions", "err", err)
	}

	txs := types.NewTransactionsByPriceAndNonce(pending)
	val.commitTransactions(val.eventMux, txs, val.chain, val.account.Address)

	val.kusd.TxPool().RemoveBatch(val.failedTxs)

	// Create the new block to seal with the consensus engine
	var block *types.Block
	if block, err = val.engine.Finalize(val.chain, header, state, val.txs, val.receipts, commit); err != nil {
		log.Crit("Failed to finalize block for sealing", "err", err)
	}

	for _, r := range val.receipts {
		for _, l := range r.Logs {
			l.BlockHash = block.Hash()
		}
	}
	for _, log := range state.Logs() {
		log.BlockHash = block.Hash()
	}
	return block
}

func (val *Validator) propose() {
	block := val.createProposalBlock()

	//lockedRound, lockedBlock := val.votes.LockingInfo()
	lockedRound := 1
	lockedBlock := common.Hash{}

	// @TODO (rgeraldes) - review int/int64; address situation where validators size might be zero (no peers)
	// @NOTE (rgeraldes) - (for now size = block size) number of block fragments = number of validators - self
	blockFragments, err := block.AsFragments(int(block.Size().Int64()) /*/val.validators.Size() - 1 */)
	if err != nil {
		// @TODO(rgeraldes) - complete
		log.Crit("Failed to get the block as a set of fragments of information", "err", err)
	}

	proposal := types.NewProposal(val.blockNumber, val.round, blockFragments.Metadata(), lockedRound, lockedBlock)

	signedProposal, err := val.wallet.SignProposal(val.account, proposal, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the proposal", "err", err)
	}

	val.proposal = signedProposal
	val.block = block

	val.eventMux.Post(core.NewProposalEvent{Proposal: proposal})

	// post block segments events
	for i := 0; i < blockFragments.Size(); i++ {
		val.eventMux.Post(core.NewBlockFragmentEvent{val.blockNumber, val.round, blockFragments.Get(i)})
	}

}

func (val *Validator) preVote() {
	var vote common.Hash
	switch {
	case val.lockedBlock != nil:
		log.Debug("Locked Block is not nil, voting for the locked block")
		vote = val.lockedBlock.Hash()
	case val.block == nil:
		log.Debug("Proposal's block is nil, voting nil")
		vote = common.Hash{}
	default:
		log.Debug("Voting for the proposal's block")
		vote = val.block.Hash()
	}

	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreVote))
}

func (val *Validator) preCommit() {
	var vote common.Hash
	// access prevotes
	winner := common.Hash{}
	switch {
	// no majority
	//case !val.hasPolka():
	// majority pre-voted nil
	case winner == common.Hash{}:
		log.Debug("majority of validators pre-voted nil")
		// unlock locked block
		if val.lockedBlock != nil {
			val.lockedRound = 0
			val.lockedBlock = nil
		}
	// majority pre-voted the locked block
	case winner == val.lockedBlock.Hash():
		log.Debug("majority of validators pre-voted the locked block")
		// update locked block round
		val.lockedRound = val.round
		// vote on the pre-vote election winner
		vote = winner
	// majority pre-voted the proposed block
	case winner == val.block.Hash():
		log.Debug("majority of validators pre-voted the proposed block")
		// lock block
		val.lockedRound = val.round
		val.lockedBlock = val.block
		// vote on the pre-vote election winner
		vote = winner
	// we don't have the current block (fetch)
	// @TODO (tendermint): in the future save the POL prevotes for justification.
	// fetch block, unlock, precommit
	default:
		// unlock locked block
		val.lockedRound = 0
		val.lockedBlock = nil
		//val.lockedBlockParts = nil
		//if !cs.ProposalBlockParts.HasHeader(blockID.PartsHeader) {
		val.block = nil
		//val.ProposalBlockParts = types.NewPartSetFromHeader(blockID.PartsHeader)
		//}
	}

	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreCommit))
}

func (val *Validator) vote(vote *types.Vote) {
	signedVote, err := val.wallet.SignVote(val.account, vote, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the vote", "err", err)
	}
	val.eventMux.Post(core.NewVoteEvent{Vote: signedVote})
}

// @TODO (rgeraldes) - verify if round is necessary (not the case in tendermint)
func (val *Validator) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) {
	val.blockFragments.Add(fragment)

	// assemble block
	if val.blockFragments.HasAll() {
		block, err := val.blockFragments.Assemble()
		if err != nil {
			log.Crit("Failed to assemble the block", "err", err)
		}

		// @TODO (rgeraldes) - validations
		log.Info("Received the complete proposal block", "hash", block.Hash())
		val.block = block

		// @TODO (rgeraldes) - post event (to confirm that when we get to the commit state that we wait for the block)
	}
}
