package validator

import (
	"errors"
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

var (
	ErrCantStopNonStartedValidator       = errors.New("can't stop validator, not started")
	ErrCantSetCoinbaseOnStartedValidator = errors.New("can't set coinbase, already started validating")
)

// Backend wraps all methods required for mining.
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() kusddb.Database
}

type Validator interface {
	Start(coinbase common.Address, deposit uint64)
	Stop() error
	SetExtra(extra []byte) error
	Validating() bool
	SetCoinbase(addr common.Address) error
	SetDeposit(deposit uint64)
	Pending() (*types.Block, *state.StateDB)
	PendingBlock() *types.Block
	Deposits() ([]*types.Deposit, error)
	RedeemDeposits() error
}

// work is the proposer current environment and holds all of the current state information
type work struct {
	state    *state.StateDB
	header   *types.Header
	tcount   int
	txs      []*types.Transaction
	receipts []*types.Receipt
}

// validator represents a consensus validator
type validator struct {
	election       *election // consensus election
	maxTransitions int       // max number of state transitions (tests) 0 - unlimited

	running    int32
	validating int32
	deposit    uint64

	signer types.Signer

	// blockchain
	backend  Backend
	chain    *core.BlockChain
	config   *params.ChainConfig
	engine   consensus.Engine
	vmConfig vm.Config

	walletAccount accounts.WalletAccount

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync

	// events
	eventMux    *event.TypeMux
	majorityCh  chan core.NewMajorityEvent
	majoritySub event.Subscription
	proposalCh  chan core.ProposalEvent
	proposalSub event.Subscription

	// state changes related to the election
	proposal    *types.Proposal
	block       *types.Block
	lockedRound uint64
	lockedBlock *types.Block
	commitRound int
	*work

	wg sync.WaitGroup
}

// New returns a new consensus validator
func New(walletAccount accounts.WalletAccount, backend Backend, election *election, config *params.ChainConfig, eventMux *event.TypeMux, engine consensus.Engine, vmConfig vm.Config) *validator {
	validator := &validator{
		config:        config,
		backend:       backend,
		chain:         backend.BlockChain(),
		engine:        engine,
		election:      election,
		eventMux:      eventMux,
		signer:        types.NewAndromedaSigner(config.ChainID),
		vmConfig:      vmConfig,
		canStart:      0,
		walletAccount: walletAccount,
	}

	go validator.sync()

	return validator
}

func (val *validator) sync() {
	if err := SyncWaiter(val.eventMux); err != nil {
		log.Warn("Failed to sync with network", "err", err)
	} else {
		val.finishedSync()
	}
}

func (val *validator) finishedSync() {
	start := atomic.LoadInt32(&val.shouldStart) == 1
	atomic.StoreInt32(&val.canStart, 1)
	atomic.StoreInt32(&val.shouldStart, 0)
	if start {
		val.Start(val.walletAccount.Account().Address, val.deposit)
	}
}

func (val *validator) Start(coinbase common.Address, deposit uint64) {
	if val.Validating() {
		log.Warn("Failed to start the validator - the state machine is already running")
		return
	}

	atomic.StoreInt32(&val.shouldStart, 1)

	newWalletAccount, _ := accounts.NewWalletAccount(val.walletAccount, accounts.Account{Address: coinbase})
	val.walletAccount = newWalletAccount
	val.deposit = deposit

	if atomic.LoadInt32(&val.canStart) == 0 {
		log.Info("Network syncing, will start validator afterwards")
		return
	}

	go val.run()
}

func (val *validator) run() {
	log.Info("Starting validation operation")
	val.wg.Add(1)
	atomic.StoreInt32(&val.running, 1)

	defer func() {
		val.wg.Done()
		atomic.StoreInt32(&val.running, 0)
	}()

	log.Info("Starting the consensus state machine")
	for state, numTransitions := val.notLoggedInState, 0; state != nil; numTransitions++ {
		state = state()
		if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
			break
		}
	}
}

func (val *validator) Stop() error {
	if !val.Validating() {
		return ErrCantStopNonStartedValidator
	}
	log.Info("Stopping consensus validator")

	val.leave()
	val.wg.Wait() // waits until the validator is no longer registered as a voter.

	atomic.StoreInt32(&val.shouldStart, 0)
	log.Info("Consensus validator stopped")
	return nil
}

func (val *validator) SetExtra(extra []byte) error { return nil }

func (val *validator) Validating() bool {
	return atomic.LoadInt32(&val.validating) > 0
}

func (val *validator) SetCoinbase(address common.Address) error {
	if val.Validating() {
		return ErrCantSetCoinbaseOnStartedValidator
	}
	newWalletAccount, err := accounts.NewWalletAccount(val.walletAccount, accounts.Account{Address: address})
	if err != nil {
		return err
	}
	val.walletAccount = newWalletAccount
	return nil
}

func (val *validator) SetDeposit(deposit uint64) {
	val.deposit = deposit
}

// Pending returns the currently pending block and associated state.
func (val *validator) Pending() (*types.Block, *state.StateDB) {
	state, err := val.chain.State()
	if err != nil {
		log.Crit("Failed to fetch the latest state", "err", err)
	}

	return val.chain.CurrentBlock(), state
}

func (val *validator) PendingBlock() *types.Block {
	return val.chain.CurrentBlock()
}

func (val *validator) restoreLastCommit() {
	checksum, err := val.election.ValidatorsChecksum()
	if err != nil {
		log.Crit("Failed to access the voters checksum", "err", err)
	}

	if err := val.election.updateValidators(checksum, true); err != nil {
		log.Crit("Failed to update the validator set", "err", err)
	}

	currentBlock := val.chain.CurrentBlock()
	if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
		return
	}
}

func (val *validator) isProposer() bool {
	return val.election.Proposer() == val.walletAccount.Account().Address
}

func (val *validator) commitTransactions(mux *event.TypeMux, txs *types.TransactionsByPriceAndNonce, bc *core.BlockChain, coinbase common.Address) {
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
		from, _ := types.TxSender(val.signer, tx)

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
		}(cpy, val.tcount)
	}
}

func (val *validator) commitTransaction(tx *types.Transaction, bc *core.BlockChain, coinbase common.Address, gp *core.GasPool) (error, []*types.Log) {
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

func (val *validator) leave() {
	err := val.election.Leave(val.walletAccount)
	if err != nil {
		log.Error("failed to leave the election", "err", err)
	}
}

func (val *validator) createProposalBlock() *types.Block {
	if val.lockedBlock != nil {
		log.Info("Picking a locked block")
		return val.lockedBlock
	}
	return val.createBlock()
}

func (val *validator) createBlock() *types.Block {
	log.Info("Creating a new block")
	// new block header
	parent := val.chain.CurrentBlock()
	blockNumber := parent.Number()
	tstart := time.Now()
	tstamp := tstart.Unix()
	if parent.Time().Cmp(new(big.Int).SetInt64(tstamp)) >= 0 {
		tstamp = parent.Time().Int64() + 1
	}
	header := &types.Header{
		ParentHash: parent.Hash(),
		Coinbase:   val.walletAccount.Account().Address,
		Number:     blockNumber.Add(blockNumber, common.Big1),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Time:       big.NewInt(tstamp),
	}
	val.header = header

	var commit *types.Commit

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
	}

	if err := val.engine.Prepare(val.chain, header); err != nil {
		log.Error("Failed to prepare header for mining", "err", err)
		return nil
	}

	pending, err := val.backend.TxPool().Pending()
	if err != nil {
		log.Crit("Failed to fetch pending transactions", "err", err)
	}

	txs := types.NewTransactionsByPriceAndNonce(val.signer, pending)
	val.commitTransactions(val.eventMux, txs, val.chain, val.walletAccount.Account().Address)

	// Create the new block to seal with the consensus engine
	var block *types.Block
	if block, err = val.engine.Finalize(val.chain, header, val.state, val.txs, commit, val.receipts); err != nil {
		log.Crit("Failed to finalize block for sealing", "err", err)
	}

	return block
}

func (val *validator) propose() {
	block := val.createProposalBlock()

	lockedRound := 1
	lockedBlock := common.Hash{}
	electionNumber := val.election.Number()
	electionRound := val.election.Round()

	fragments, err := block.AsFragments(int(block.Size().Int64()))
	if err != nil {
		log.Crit("Failed to get the block as a set of fragments of information", "err", err)
	}

	proposal := types.NewProposal(electionNumber, electionRound, fragments.Metadata(), lockedRound, lockedBlock)

	signedProposal, err := val.walletAccount.SignProposal(val.walletAccount.Account(), proposal, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the proposal", "err", err)
	}

	val.proposal = signedProposal
	val.block = block

	val.eventMux.Post(core.NewProposalEvent{Proposal: proposal})

	for i := uint(0); i < fragments.Size(); i++ {
		val.eventMux.Post(core.NewBlockFragmentEvent{
			BlockNumber: electionNumber,
			Round:       electionRound,
			Data:        fragments.Get(int(i)),
		})
	}

}

func (val *validator) preVote() {
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

	val.vote(types.NewVote(val.election.Number(), vote, val.election.Round(), types.PreVote))
}

func (val *validator) preCommit() {
	var vote common.Hash
	// access prevotes
	winner := common.Hash{}
	switch {
	// no majority
	// majority pre-voted nil
	case winner == common.Hash{}:
		log.Debug("Majority of validators pre-voted nil")
		// unlock locked block
		if val.lockedBlock != nil {
			val.lockedRound = 0
			val.lockedBlock = nil
		}
	case winner == val.lockedBlock.Hash():
		log.Debug("Majority of validators pre-voted the locked block")
		// update locked block round
		val.lockedRound = val.election.Round()
		// vote on the pre-vote election winner
		vote = winner
	case winner == val.block.Hash():
		log.Debug("Majority of validators pre-voted the proposed block")
		// lock block
		val.lockedRound = val.election.Round()
		val.lockedBlock = val.block
		// vote on the pre-vote election winner
		vote = winner
		// we don't have the current block (fetch)
	default:
		// fetch block, unlock, precommit
		// unlock locked block
		val.lockedRound = 0
		val.lockedBlock = nil
		val.block = nil
	}

	val.vote(types.NewVote(val.election.Number(), vote, val.election.Round(), types.PreCommit))
}

func (val *validator) vote(vote *types.Vote) {
	signedVote, err := val.walletAccount.SignVote(val.walletAccount.Account(), vote, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the vote", "err", err)
	}

	if err := val.election.Vote(signedVote); err != nil {
		log.Warn("Failed to add own vote to voting table", "err", err)
	}
}

func (val *validator) makeCurrent(parent *types.Block) error {
	state, err := val.chain.StateAt(parent.Root())
	if err != nil {
		return err
	}
	work := &work{
		state: state,
	}

	// Keep track of transactions which return errors so they can be removed
	work.tcount = 0
	val.work = work
	return nil
}

func (val *validator) Deposits() ([]*types.Deposit, error) {
	return val.election.Deposits(val.walletAccount.Account().Address)
}

func (val *validator) RedeemDeposits() error {
	return val.election.RedeemDeposits(val.walletAccount)
}

func (val *validator) init() error {
	parent := val.backend.BlockChain().CurrentBlock()

	if err := val.election.init(); err != nil {
		return err
	}

	val.lockedRound = 0
	val.lockedBlock = nil
	val.commitRound = -1

	proposalCh := make(chan core.ProposalEvent)
	val.proposalSub = val.election.SubscribeProposalEvent(proposalCh)

	majortyCh := make(chan core.NewMajorityEvent)
	val.majoritySub = val.election.SubscribeMajorityEvent(majortyCh)

	if err := val.makeCurrent(parent); err != nil {
		log.Error("Failed to create mining context", "err", err)
		return nil
	}
	return nil
}
