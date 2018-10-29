package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	txConfirmationTimeout = 20 * time.Second
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
	Post(interface{}) error
}

type Service interface {
	Validating() bool
	Running() bool
	WaitProposal() bool
	AddProposal(proposal *types.Proposal) error
	AddVote(vote *types.Vote) error
	AddBlockFragment(blockNumber *big.Int, blockHash common.Hash, round uint64, fragment *types.BlockFragment) error
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
	canStart          int32 // can start indicates whether we can start the validation operation
	isStarted         int32 // is Start method has been already called
	waitProposalBlock int32

	// events
	eventMux    *event.TypeMux
	eventPoster eventPoster

	wg sync.WaitGroup

	handleMutex sync.Mutex
}

// New returns a new consensus validator
func New(backend Backend, consensus *consensus.Consensus, config *params.ChainConfig, eventMux *event.TypeMux, engine engine.Engine, vmConfig vm.Config) *validator {
	validator := &validator{
		config:      config,
		backend:     backend,
		chain:       backend.BlockChain(),
		engine:      engine,
		consensus:   consensus,
		eventMux:    eventMux,
		eventPoster: newPoster(eventMux),
		signer:      types.NewAndromedaSigner(config.ChainID),
		vmConfig:    vmConfig,
	}

	return validator
}

func (val *validator) Start(walletAccount accounts.WalletAccount, deposit *big.Int) {
	log.Warn("initial sync is going to run", "account", walletAccount.Account().Address.String(), "deposit", spew.Sdump(deposit))

	val.walletAccount = walletAccount
	val.deposit = deposit

	if isGenesisValidator, err := val.isGenesisValidator(); err == nil && !isGenesisValidator {
		if !val.isBlockZero() {
			if err := val.tryJoin(); err != nil {
				log.Warn("stopping validation", "err", err)
				return
			}
		}
	}
	log.Warn("initial sync done")

	atomic.StoreInt32(&val.canStart, 1)

	log.Warn("trying to start validator")
	if val.Validating() {
		log.Warn("failed to start the validator - the state machine is already running")
		return
	}

	if !atomic.CompareAndSwapInt32(&val.isStarted, 0, 1) {
		log.Warn("failed to start the validator - the state machine is already running 1")
		return
	}

	if atomic.LoadInt32(&val.canStart) == 0 {
		atomic.StoreInt32(&val.isStarted, 0)
		log.Debug("network syncing, will start validator afterwards")
		return
	}

	go val.run()
}

func (val *validator) run() {
	log.Info("Starting validation operation. run")
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

	atomic.StoreInt32(&val.isStarted, 0)
	log.Info("Consensus validator stopped")
	return nil
}

func (val *validator) SetExtra(extra []byte) error { return nil }

func (val *validator) Validating() bool {
	return atomic.LoadInt32(&val.validating) > 0
}

func (val *validator) WaitProposal() bool {
	return atomic.LoadInt32(&val.waitProposalBlock) == 1
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

	if deposit != nil {
		log.Info("setting a deposit on setDeposit call", "deposit", deposit)
	} else {
		log.Info("setting a deposit on setDeposit call", "deposit", "nil")
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

func (val *validator) init() error {
	parent := val.chain.CurrentBlock()
	val.parentBlockCreatedAt = time.Unix(parent.Time().Int64(), 0)
	val.previousRoundCreatedAt = time.Unix(time.Now().Unix(), 0)

	val.blockNumber = parent.Number().Add(parent.Number(), big.NewInt(1))
	val.round = 0

	val.lockedRound = 0
	val.lockedBlock = nil
	val.commitRound = -1

	val.checkUpdateValidators(true)

	val.blockCh = make(chan *types.Block, 1)
	val.majority = val.eventMux.Subscribe(core.NewMajorityEvent{})

	if err := val.makeCurrent(parent); err != nil {
		log.Error("Failed to create mining context", "err", err)
		return nil
	}

	return nil
}

func (val *validator) checkUpdateValidators(forceUpdate bool) {
	var (
		checksum types.VotersChecksum
		err      error
	)

	if !forceUpdate {
		checksum, err = val.consensus.ValidatorsChecksum()
		if err != nil {
			log.Crit("Failed to access the voters checksum", "err", err)
		}

		if val.votersChecksum == checksum {
			return
		}
	}

	if err = val.updateValidators(checksum); err != nil {
		log.Crit("Failed to update the validator set", "err", err)
	}

	val.votingSystem = NewVotingSystem(val.voters, val.eventPoster)
}

func (val *validator) AddProposal(proposal *types.Proposal) error {
	if !val.Validating() {
		return ErrCantAddProposalNotValidating
	}

	if val.config.ChainID.Cmp(proposal.ChainID()) != 0 {
		return fmt.Errorf("expected proposed block for chainID %v, got %v",
			val.config.ChainID.Int64(), proposal.ChainID().Int64())
	}

	if val.blockNumber.Cmp(proposal.BlockNumber()) != 0 {
		return fmt.Errorf("expected proposed block number %v, got %v",
			val.blockNumber.Int64(), proposal.BlockNumber().Int64())
	}

	if val.round != proposal.Round() {
		return fmt.Errorf("expected proposed block round %v, got %v", val.round, proposal.Round())
	}

	if val.isProposal {
		return errors.New("is proposal itself")
	}

	proposalAddress, err := types.ProposalSender(val.signer, proposal)
	if err != nil {
		return err
	}

	if val.proposer == nil {
		return errors.New("a proposer not defined yet")
	}

	if val.proposer.Address() != proposalAddress {
		return fmt.Errorf("expected proposer %v, got %v", val.proposer.Address(), proposalAddress)
	}

	val.handleMutex.Lock()

	val.proposal = proposal

	val.blockFragmentsLock.Lock()

	blockFragments, ok := val.blockFragments[proposal.Hash()]
	if !ok {
		blockFragments = types.NewDataSetFromMeta(proposal.BlockMetadata())
	}
	val.blockFragments[proposal.Hash()] = blockFragments

	val.blockFragmentsLock.Unlock()

	hasAllFragments := blockFragments.HasAll()
	val.handleMutex.Unlock()

	if hasAllFragments {
		log.Debug("addProposal has all fragments and assembling a block")
		return val.assembleBlock(proposal.Round(), proposal.BlockNumber(), proposal.Hash())
	}

	return nil
}

func (val *validator) AddVote(vote *types.Vote) error {
	log.Debug("adding vote", "vote", vote.String())

	val.handleMutex.Lock()
	defer func() {
		val.handleMutex.Unlock()
	}()

	addressVote, err := types.NewAddressVote(val.signer, vote)
	if err != nil {
		log.Debug("can't add a vote", "err", err, "vote", vote.String())
		return err
	}

	if err = val.addVote(addressVote); err != nil {
		log.Debug("cant add a vote",
			"validator", addressVote.Address(),
			"number", vote.BlockNumber().Int64(),
			"round", vote.Round(),
			"chainID", vote.ChainID().Int64(),
			"hash", vote.Hash().String(),
			"err", err.Error())
	}
	log.Debug("vote has been added", "vote", vote.String())
	return err
}

func (val *validator) addVote(addressVote types.AddressVote) error {
	if !val.Validating() {
		return ErrCantVoteNotValidating
	}

	vote := addressVote.Vote()

	if val.config.ChainID.Cmp(vote.ChainID()) != 0 {
		return fmt.Errorf("expected vote for chainID %v, got %v",
			val.config.ChainID.Int64(), vote.ChainID().Int64())
	}

	if val.blockNumber.Cmp(vote.BlockNumber()) != 0 {
		return fmt.Errorf("expected vote block number %v, got %v",
			val.blockNumber.Int64(), vote.BlockNumber().Int64())
	}

	if val.round != vote.Round() {
		return fmt.Errorf("expected vote round number %v, got %v",
			val.round, vote.Round())
	}

	if err := val.votingSystem.Add(addressVote); err != nil {
		log.Debug("cannot add the vote", "err", err)
		return err
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

func (val *validator) join() error {
	if val.deposit == nil || val.deposit.Cmp(big.NewInt(0)) <= 0 {
		err := errors.New("failed to join the network as a validator. Deposit is less or equal 0")
		return err
	}

	log.Debug("joining the network",
		"chainID", val.config.ChainID.Int64(),
		"account", val.walletAccount,
		"deposit", val.deposit,
		"block", val.block,
		"deposit", spew.Sdump(val.consensus.Deposits(val.walletAccount.Account().Address)))

	txHash, err := val.consensus.Join(val.walletAccount, val.deposit)
	if err != nil {
		log.Error("Error joining validators network", "err", err)
		return err
	}
	log.Info("Waiting confirmation to participate in the consensus")

	receipt, err := WaitForTransaction(txHash, val.backend, txConfirmationTimeout)
	if err != nil {
		log.Crit("Failed to verify the voter registration", "err", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		log.Crit("Failed to register the validator - receipt status failed")
	}
	return nil
}

func WaitForTransaction(txHash common.Hash, backend Backend, duration time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	return tx.WaitMinedWithTimeout(ctx, backend, txHash)
}

func (val *validator) leave() {
	log.Warn("leaving the network",
		"chainID", val.config.ChainID.Int64(),
		"blockNumber", val.block,
		"account", val.walletAccount)

	txHash, err := val.consensus.Leave(val.walletAccount)
	if err != nil {
		log.Error("failed to leave the election", "err", err)
	}

	receipt, err := WaitForTransaction(txHash, val.backend, txConfirmationTimeout)
	if err != nil {
		log.Error("Failed to verify the voter unregistration", "err", err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		log.Error("Failed to unregister validator - receipt status failed")
	}

	// fixme wait for the block propagation
	time.Sleep(5 * time.Second)

	log.Warn("leaved validators list", "block", val.blockNumber.Int64())
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
	tstamp := time.Now().Unix()

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

	first := types.NewVote(blockNumber, parent.Hash(), val.round, types.PreCommit)

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
		// fixme what if the proposed block will be declined? shall we somehow revert this state on the proposer's stateDB? it looks like that we should see a lot reorgs in logs.
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
	val.isProposal = true
	val.block = block

	// give some time to the network to start waiting for the proposed block
	time.Sleep(time.Duration(params.ProposeDeltaDuration) * time.Millisecond)
	if err := val.eventMux.Post(core.NewProposalEvent{Proposal: signedProposal}); err != nil {
		log.Warn("can't post an event", "err", err, "event", "NewProposalEvent", "proposal", signedProposal)
	}
	log.Debug("sending a signed proposal",
		"account", val.walletAccount.Account().Address.String(),
		"chainID", val.config.ChainID.Int64(),
		"block", signedProposal.String(),
		"round", val.round)

	for i := uint(0); i < fragments.Size(); i++ {
		blockFragmentEvent := core.NewBlockFragmentEvent{
			BlockNumber: val.blockNumber,
			BlockHash:   signedProposal.Hash(),
			Round:       val.round,
			Data:        fragments.Get(int(i)),
		}

		err := val.eventMux.Post(blockFragmentEvent)
		if err != nil {
			log.Warn("can't post an event", "err", err, "event", "NewBlockFragmentEvent", "blockFragment", blockFragmentEvent)
		}
	}

	log.Debug("all block fragments has been sent",
		"account", val.walletAccount.Account().Address.String(),
		"chainID", val.config.ChainID.Int64(),
		"block", signedProposal.String(),
		"round", val.round)
}

func (val *validator) Post(e interface{}) error {
	return val.eventMux.Post(e)
}

func (val *validator) preVote() {
	var vote common.Hash
	switch {
	case val.lockedBlock != nil:
		log.Debug("Locked Block is not nil, voting for the locked block")
		vote = val.lockedBlock.Hash()
	case val.block == nil:
		log.Warn("Proposal's block is nil, voting nil", "block", val.block == nil, "locked", val.lockedBlock == nil)
		vote = common.Hash{}
	default:
		log.Debug("Voting for the proposal's block")
		vote = val.block.Hash()
	}

	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreVote))

	// fixme: fix only for a small network - resend vote. It's necessary because we don't have a reliable sync mechanism with heartbeat
	time.Sleep(time.Duration(params.PreVoteDeltaDuration) * time.Millisecond)
	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreVote))
}

func (val *validator) preCommit() {
	var vote common.Hash

	// current leader by simple majority
	votingTable, err := val.votingSystem.getVotingTable(val.round, types.PreVote)
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
		if val.lockedBlock != nil {
			log.Warn("lockedBlock.preCommit default case",
				"equal", val.lockedBlock.Hash() == currentLeader,
				"currentLeader", currentLeader,
				"lockedHash", val.lockedBlock.Hash())
		}

		if val.block != nil {
			log.Warn("block.preCommit default case.",
				"equal", val.block.Hash() == currentLeader,
				"currentLeader", currentLeader,
				"block", val.block.Hash())
		}

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

	if val.walletAccount.Account().Address != addressVote.Address() {
		log.Crit("vote signature is incorrect", "vote", addressVote, "account", val.walletAccount.Account())
	}

	err = val.votingSystem.Add(addressVote)
	if err != nil {
		log.Debug("Failed to add own vote to voting table",
			"err", err, "blockHash", addressVote.Vote().BlockHash(), "hash", addressVote.Vote().Hash())
	}
}

func (val *validator) AddBlockFragment(blockNumber *big.Int, blockHash common.Hash, round uint64, fragment *types.BlockFragment) error {
	if !val.Validating() {
		return ErrCantAddBlockFragmentNotValidating
	}

	if val.blockNumber.Cmp(blockNumber) != 0 {
		return fmt.Errorf("expected block fragments for %d, got %d", val.blockNumber.Int64(), blockNumber.Int64())
	}

	if val.round != round {
		return fmt.Errorf("expected block fragments for round %d, got %d", val.round, round)
	}

	if val.isProposal {
		return errors.New("is proposer itself")
	}

	err, canAssemble := val.addFragment(fragment, blockHash)
	if err != nil {
		return err
	}

	if canAssemble {
		log.Debug("addBlockFragment has all fragments and assembling a block")
		return val.assembleBlock(round, blockNumber, blockHash)
	}

	return nil
}

// returns true if all fragments can be assembled now
func (val *validator) addFragment(fragment *types.BlockFragment, blockHash common.Hash) (error, bool) {
	val.blockFragmentsLock.Lock()
	defer func() {
		val.blockFragmentsLock.Unlock()
	}()

	var (
		blockFragmentsList []*types.BlockFragment
		ok                 bool
	)

	blockFragments, ok := val.blockFragments[blockHash]
	if !ok {
		// proposal block metadata has not received yet
		blockFragmentsList, ok = val.blockFragmentsStorage[blockHash]
		if !ok {
			blockFragmentsList = []*types.BlockFragment{}
		}

		val.blockFragmentsStorage[blockHash] = append(blockFragmentsList, fragment)
		return nil, false
	}

	// if val.blockFragments is set, proposal block metadata was received
	for _, blockFragment := range blockFragmentsList {
		if err := blockFragments.Add(blockFragment); err != nil {
			log.Debug("failed to add a new block fragment", "err", err.Error())
		}
	}

	if !val.HasProposer() {
		return nil, false
	}

	// cleanup
	val.blockFragmentsStorage[blockHash] = nil

	if err := blockFragments.Add(fragment); err != nil {
		err = errors.New("failed to add a new block fragment: " + err.Error())
		return err, false
	}

	return nil, blockFragments.HasAll()
}

func (val *validator) assembleBlock(round uint64, blockNumber *big.Int, blockHash common.Hash) error {
	val.handleMutex.Lock()
	if val.block != nil {
		return nil
	}

	val.blockFragmentsLock.Lock()
	defer func() {
		val.handleMutex.Unlock()

		val.blockFragmentsLock.Unlock()
	}()

	blockFragments := val.blockFragments[blockHash]

	block, err := blockFragments.Assemble()
	if err != nil {
		err = errors.New("Failed to assemble the block: " + err.Error())
		log.Warn("error while adding a new block fragment", "err", err,
			"round", round, "block", blockNumber)
		return err
	}

	// Start the parallel header verifier
	if err = val.validateProposalBlock(block, round, blockNumber); err != nil {
		return err
	}

	parent := val.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		return fmt.Errorf("wrong block received. hash %v, number %v, parentHash %v",
			block.Hash().String(), block.Number().Int64(), block.ParentHash().String())
	}

	// Process block using the parent state as reference point.
	receipts, _, usedGas, err := val.chain.Processor().Process(block, val.state, val.vmConfig)
	if err != nil {
		log.Error("Failed to process the block", "err", err,
			"round", round, "block", blockNumber, "block", block)

		log.Crit("Failed to process the block", "err", err)
	}

	val.receipts = receipts

	// Validate the state using the default validator
	err = val.chain.Validator().ValidateState(block, parent, val.state, receipts, usedGas)
	if err != nil {
		log.Crit("Failed to validate the state", "err", err, "round", round, "block", blockNumber, "parent", parent.Number().Int64(), "blockFragments", blockFragments.Count(), "receipts", receipts, "block", block)
	}

	val.block = block

	go func() {
		val.blockCh <- block
	}()

	return nil
}

func (val *validator) validateProposalBlock(block *types.Block, round uint64, blockNumber *big.Int) error {
	const nBlocks = 1
	headers := make([]*types.Header, nBlocks)
	seals := make([]bool, nBlocks)
	headers[nBlocks-1] = block.Header()
	seals[nBlocks-1] = true

	abort, results := val.engine.VerifyHeaders(val.chain, headers, seals)
	defer close(abort)

	err := <-results
	if err != nil {
		log.Debug("error while verifying a block headers",
			"err", err, "round", round, "block", blockNumber, "block",
			block, "headers", headers, "seals", seals)
		return err
	}

	err = val.chain.Validator().ValidateBody(block)
	if err != nil {
		err = errors.New("Failed to validate thr block body: " + err.Error())
		log.Debug("error while validating a block body",
			"err", err, "round", round, "block", blockNumber, "block", block)

		return err
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

func (val *validator) updateValidators(checksum [32]byte) error {
	validators, err := val.consensus.Validators()
	if err != nil {
		return err
	}

	blockNumber := int64(-1)
	if val.blockNumber != nil {
		blockNumber = val.blockNumber.Int64()
	}
	if val.voters != nil {
		log.Debug("voting. updating a list of validators", "was", val.voters.Len(), "now", validators.Len(), "block", blockNumber, "round", val.round)
	} else {
		log.Debug("voting. updating a list of validators", "was", "nil", "now", validators.Len(), "block", blockNumber, "round", val.round)
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

	receipt, err := WaitForTransaction(txHash, val.backend, txConfirmationTimeout)
	if err != nil {
		return err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("failed to redeem deposits - receipt status failed")
	}

	return nil
}

func (val *validator) HasProposer() bool {
	return val.proposer != nil
}

func roundStartTimer(blockStarts time.Time, round uint64) *time.Timer {
	startsAt := roundStartsAt(blockStarts, round)
	return time.NewTimer(startsAt.Sub(time.Now()))
}

func roundStartsAt(blockStarts time.Time, round uint64) time.Time {
	return blockStarts.Add(time.Duration(params.BlockTime*(round+1)) * time.Millisecond)
}
