package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/tx"
	engine "github.com/kowala-tech/kcoin/client/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/kcoindb"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

var (
	ErrCantStopNonStartedValidator       = errors.New("can't stop validator, not started")
	ErrCantVoteNotValidating             = errors.New("can't vote, not validating")
	ErrCantSetCoinbaseOnStartedValidator = errors.New("can't set coinbase, already started validating")
	ErrCantAddProposalNotValidating      = errors.New("can't add proposal, not validating")
	ErrCantAddBlockFragmentNotValidating = errors.New("can't add block fragment, not validating")
	ErrIsNotRunning                      = errors.New("validator is not running")
	ErrIsRunning                         = errors.New("validator is running, cannot change its parameters")
)

var (
	txConfirmationTimeout = 10 * time.Second
)

// Backend wraps all methods required for mining.
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() kcoindb.Database
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

type Validator interface {
	Service
	Start(walletAccount accounts.WalletAccount, deposit *big.Int)
	Stop() error
	SetExtra(extra []byte) error
	SetCoinbase(walletAccount accounts.WalletAccount) error
	SetDeposit(deposit *big.Int) error
	Pending() (*types.Block, *state.StateDB)
	PendingBlock() *types.Block
	Deposits(address *common.Address) ([]*types.Deposit, error)
	RedeemDeposits() error
}

type Service interface {
	Validating() bool
	Running() bool
	AddProposal(proposal *types.Proposal) error
	AddVote(vote *types.Vote) error
	AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error
}

// validator represents a consensus validator
type validator struct {
	VotingState        // consensus internal state
	maxTransitions int // max number of state transitions (tests) 0 - unlimited

	running    int32
	validating int32
	deposit    *big.Int

	signer types.Signer

	// blockchain
	backend  Backend
	chain    *core.BlockChain
	config   *params.ChainConfig
	engine   engine.Engine
	vmConfig vm.Config

	walletAccount accounts.WalletAccount

	consensus *consensus.Consensus // consensus binding

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync

	// events
	eventMux *event.TypeMux

	wg sync.WaitGroup

	handleMutex sync.Mutex
}

// New returns a new consensus validator
func New(backend Backend, consensus *consensus.Consensus, config *params.ChainConfig, eventMux *event.TypeMux, engine engine.Engine, vmConfig vm.Config) *validator {
	validator := &validator{
		config:    config,
		backend:   backend,
		chain:     backend.BlockChain(),
		engine:    engine,
		consensus: consensus,
		eventMux:  eventMux,
		signer:    types.NewAndromedaSigner(config.ChainID),
		vmConfig:  vmConfig,
		canStart:  0,
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
		val.Start(val.walletAccount, val.deposit)
	}
}

func (val *validator) Start(walletAccount accounts.WalletAccount, deposit *big.Int) {
	if val.Validating() {
		log.Warn("failed to start the validator - the state machine is already running")
		return
	}

	atomic.StoreInt32(&val.shouldStart, 1)

	val.walletAccount = walletAccount
	val.deposit = deposit

	if atomic.LoadInt32(&val.canStart) == 0 {
		log.Info("network syncing, will start validator afterwards")
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

	initialStateFunc := val.notLoggedInState
	if isGenesisValidator, err := val.isGenesisValidator(); err != nil && isGenesisValidator {
		initialStateFunc = val.genesisNotLoggedInState
	}

	log.Info("Starting the consensus state machine")
	for state, numTransitions := initialStateFunc, 0; state != nil; numTransitions++ {
		state = state()
		if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
			break
		}
	}
}

func (val *validator) isGenesisValidator() (bool, error) {
	return val.consensus.IsGenesisValidator(val.walletAccount.Account().Address)
}

func (val *validator) Stop() error {

	if !val.Running() {
		return nil
	}

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

func (val *validator) Running() bool {
	return atomic.LoadInt32(&val.running) > 0
}

func (val *validator) SetCoinbase(walletAccount accounts.WalletAccount) error {
	if val.Validating() {
		return ErrCantSetCoinbaseOnStartedValidator
	}
	val.walletAccount = walletAccount
	return nil
}

func (val *validator) SetDeposit(deposit *big.Int) error {
	if val.Validating() {
		return ErrIsRunning
	}

	val.deposit = deposit

	return nil
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
	checksum, err := val.consensus.ValidatorsChecksum()
	if err != nil {
		log.Crit("Failed to access the voters checksum", "err", err)
	}

	if err := val.updateValidators(checksum, true); err != nil {
		log.Crit("Failed to update the validator set", "err", err)
	}

	currentBlock := val.chain.CurrentBlock()
	if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
		return
	}
}

func (val *validator) init() error {
	parent := val.chain.CurrentBlock()

	checksum, err := val.consensus.ValidatorsChecksum()
	if err != nil {
		log.Crit("Failed to access the voters checksum", "err", err)
	}

	if val.votersChecksum != checksum {
		if err := val.updateValidators(checksum, true); err != nil {
			log.Crit("Failed to update the validator set", "err", err)
		}
	}

	start := time.Unix(parent.Time().Int64(), 0)
	val.start = start.Add(time.Duration(params.BlockTime) * time.Millisecond)
	val.blockNumber = parent.Number().Add(parent.Number(), big.NewInt(1))
	val.round = 0

	val.proposal = nil
	val.block = nil
	val.blockFragments = nil

	val.lockedRound = 0
	val.lockedBlock = nil
	val.commitRound = -1

	val.votingSystem, err = NewVotingSystem(val.eventMux, val.blockNumber, val.voters)
	if err != nil {
		log.Error("Failed to create voting system", "err", err)
		return nil
	}

	val.blockCh = make(chan *types.Block)
	val.majority = val.eventMux.Subscribe(core.NewMajorityEvent{})

	if err = val.makeCurrent(parent); err != nil {
		log.Error("Failed to create mining context", "err", err)
		return nil
	}

	return nil
}

func (val *validator) AddProposal(proposal *types.Proposal) error {
	if !val.Validating() {
		return ErrCantAddProposalNotValidating
	}

	log.Info("Received Proposal")

	val.handleMutex.Lock()
	val.proposal = proposal
	val.blockFragments = types.NewDataSetFromMeta(proposal.BlockMetadata())
	val.handleMutex.Unlock()

	return nil
}

func (val *validator) AddVote(vote *types.Vote) error {
	if !val.Validating() {
		return ErrCantVoteNotValidating
	}

	val.handleMutex.Lock()
	defer val.handleMutex.Unlock()

	addressVote, err := types.NewAddressVote(val.signer, vote)
	if err != nil {
		return err
	}

	if err := val.votingSystem.Add(addressVote); err != nil {
		log.Error("cannot add the vote", "err", err)
	}

	return nil
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

	receipt, _, err := core.ApplyTransaction(val.config, bc, &coinbase, gp, val.state, val.header, tx, &val.header.GasUsed, vm.Config{})
	if err != nil {
		val.state.RevertToSnapshot(snap)
		return err, nil
	}
	val.txs = append(val.txs, tx)
	val.receipts = append(val.receipts, receipt)

	return nil, receipt.Logs
}

func (val *validator) leave() {
	txHash, err := val.consensus.Leave(val.walletAccount)
	if err != nil {
		log.Error("failed to leave the election", "err", err)
	}
	receipt, err := tx.WaitMinedWithTimeout(val.backend, txHash, txConfirmationTimeout)
	if err != nil {
		log.Error("Failed to verify the voter deregistration", "err", err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		log.Error("Failed to deregister validator - receipt status failed")
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
		ParentHash:     parent.Hash(),
		Coinbase:       val.walletAccount.Account().Address,
		Number:         blockNumber.Add(blockNumber, common.Big1),
		GasLimit:       core.CalcGasLimit(parent),
		Time:           big.NewInt(tstamp),
		ValidatorsHash: val.voters.Hash(),
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

	fragments, err := block.AsFragments(int(block.Size()))
	if err != nil {
		log.Crit("Failed to get the block as a set of fragments of information", "err", err)
	}

	proposal := types.NewProposal(val.blockNumber, val.round, fragments.Metadata(), lockedRound, lockedBlock)

	signedProposal, err := val.walletAccount.SignProposal(val.walletAccount.Account(), proposal, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the proposal", "err", err)
	}

	val.proposal = signedProposal
	val.block = block

	val.eventMux.Post(core.NewProposalEvent{Proposal: proposal})

	for i := uint(0); i < fragments.Size(); i++ {
		val.eventMux.Post(core.NewBlockFragmentEvent{
			BlockNumber: val.blockNumber,
			Round:       val.round,
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
		log.Warn("Proposal's block is nil, voting nil")
		vote = common.Hash{}
	default:
		log.Debug("Voting for the proposal's block")
		vote = val.block.Hash()
	}

	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreVote))
}

func (val *validator) preCommit() {
	var vote common.Hash

	// current leader by simple majority
	votingTable, err := val.votingSystem.getVoteSet(val.round, types.PreVote)
	if err != nil {
		log.Crit("Error while preCommit stage", "err", err)
	}

	currentLeader := votingTable.Leader()

	switch {
	// no majority
	// majority pre-voted nil
	case currentLeader == common.Hash{}:
		log.Warn("Majority of validators pre-voted nil")
		// unlock locked block
		if val.lockedBlock != nil {
			val.lockedRound = 0
			val.lockedBlock = nil
		}
	case val.lockedBlock != nil && currentLeader == val.lockedBlock.Hash():
		log.Debug("Majority of validators pre-voted the locked block", "block", val.lockedBlock.Hash())
		// update locked block round
		val.lockedRound = val.round
		// vote on the pre-vote election winner
		vote = currentLeader
	case val.block != nil && currentLeader == val.block.Hash():
		log.Debug("Majority of validators pre-voted the proposed block", "block", val.block.Hash())
		// lock block
		val.lockedRound = val.round
		val.lockedBlock = val.block
		// vote on the pre-vote election winner
		vote = currentLeader
		// we don't have the current block (fetch)
	default:
		// fetch block, unlock, precommit
		// unlock locked block
		log.Warn("preCommit default case")
		val.lockedRound = 0
		val.lockedBlock = nil
		val.block = nil
	}

	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreCommit))
}

func (val *validator) vote(vote *types.Vote) {
	signedVote, err := val.walletAccount.SignVote(val.walletAccount.Account(), vote, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the vote", "err", err)
	}

	addressVote, err := types.NewAddressVote(val.signer, signedVote)
	if err != nil {
		log.Crit("Failed to make address Vote", "err", err)
	}

	err = val.votingSystem.Add(addressVote)
	if err != nil {
		log.Error("Failed to add own vote to voting table",
			"err", err, "blockHash", addressVote.Vote().BlockHash(), "hash", addressVote.Vote().Hash())
	}
}

func (val *validator) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error {
	if !val.Validating() {
		return ErrCantAddBlockFragmentNotValidating
	}

	if err := val.blockFragments.Add(fragment); err != nil {
		err = errors.New("Failed to add a new block fragment: " + err.Error())
		return err
	}

	if val.blockFragments.HasAll() {
		block, err := val.blockFragments.Assemble()
		if err != nil {
			err = errors.New("Failed to assemble the block: " + err.Error())
			log.Error("error while adding a new block fragment", "err", err, "round", round, "block", blockNumber, "fragment", fragment)
			return err
		}

		// Start the parallel header verifier
		nBlocks := 1
		headers := make([]*types.Header, nBlocks)
		seals := make([]bool, nBlocks)
		headers[nBlocks-1] = block.Header()
		seals[nBlocks-1] = true

		abort, results := val.engine.VerifyHeaders(val.chain, headers, seals)
		defer close(abort)

		err = <-results
		if err == nil {
			err = val.chain.Validator().ValidateBody(block)
			if err != nil {
				err = errors.New("Failed to validate thr block body: " + err.Error())
				log.Error("error while validating a block body",
					"err", err, "round", round, "block", blockNumber, "fragment", fragment, "block", block)

				return err
			}
		}

		parent := val.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)

		// Process block using the parent state as reference point.
		receipts, _, usedGas, err := val.chain.Processor().Process(block, val.state, val.vmConfig)
		if err != nil {
			log.Error("Failed to process the block", "err", err,
				"round", round, "block", blockNumber, "fragment", fragment, "block", block)

			log.Crit("Failed to process the block", "err", err)
		}

		// guarded section
		val.handleMutex.Lock()
		val.receipts = receipts

		// Validate the state using the default validator
		err = val.chain.Validator().ValidateState(block, parent, val.state, receipts, usedGas)
		if err != nil {
			val.handleMutex.Unlock()

			log.Error("Failed to validate the state", "err", err,
				"round", round, "block", blockNumber, "fragment", fragment, "block", block)

			log.Crit("Failed to validate the state", "err", err)
		}

		val.block = block
		val.handleMutex.Unlock()

		go func() { val.blockCh <- block }()
	}
	return nil
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

func (val *validator) updateValidators(checksum [32]byte, genesis bool) error {
	validators, err := val.consensus.Validators()
	if err != nil {
		return err
	}

	if val.voters != nil {
		log.Debug("voting. updating a list of validators", "was", val.voters.Len(), "now", validators.Len())
	} else {
		log.Debug("voting. updating a list of validators", "was", "nil", "now", validators.Len())
	}

	val.voters = validators
	val.votersChecksum = checksum

	return nil
}

func (val *validator) Deposits(address *common.Address) ([]*types.Deposit, error) {
	if address != nil {
		return val.consensus.Deposits(*address)
	}

	if val.walletAccount == nil {
		return nil, errors.New("either address or validator.Start() required")
	}

	return val.consensus.Deposits(val.walletAccount.Account().Address)
}

func (val *validator) RedeemDeposits() error {
	if !val.Validating() {
		return ErrIsNotRunning
	}

	txHash, err := val.consensus.RedeemDeposits(val.walletAccount)
	if err != nil {
		return err
	}
	receipt, err := tx.WaitMinedWithTimeout(val.backend, txHash, txConfirmationTimeout)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("Failed to redeem deposits - receipt status failed")
	}

	return nil
}
