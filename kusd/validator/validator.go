package validator

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusd/downloader"
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
	Election           // consensus state
	maxTransitions int // max number of state transitions (tests) 0 - unlimited

	running    int32
	validating int32
	deposit    uint64

	// signer
	signer types.Signer

	// blockchain
	backend  Backend
	chain    *core.BlockChain
	config   *params.ChainConfig
	engine   consensus.Engine
	vmConfig vm.Config // @NOTE (rgeraldes) - temporary

	network *network.NetworkContract // validators contract

	account accounts.Account
	wallet  accounts.Wallet // signer

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync

	// events
	eventMux *event.TypeMux

	wg sync.WaitGroup
}

// New returns a new consensus validator
func New(backend Backend, contractBackend bind.ContractBackend, config *params.ChainConfig, eventMux *event.TypeMux, engine consensus.Engine, vmConfig vm.Config) *Validator {
	validator := &Validator{
		config:   config,
		backend:  backend,
		chain:    backend.BlockChain(),
		engine:   engine,
		eventMux: eventMux,
		signer:   types.NewAndromedaSigner(config.ChainID),
		vmConfig: vmConfig,
		canStart: 0,
	}

	// Network contract instance
	state, err := validator.chain.State()
	if err != nil {
		log.Crit("Failed to fetch the current state", "err", err)
	}
	contracts, err := network.GetContracts(state)
	if err != nil {
		log.Crit("Failed to access the network contracts", "err", err)
	}
	contract, err := network.NewNetworkContract(contracts.Network, contractBackend)
	if err != nil {
		log.Crit("Failed to load the network contract", "err", err)
	}
	validator.network = contract

	go validator.sync()

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
		case downloader.DoneEvent, downloader.FailedEvent:
			start := atomic.LoadInt32(&val.shouldStart) == 1
			atomic.StoreInt32(&val.canStart, 1)
			atomic.StoreInt32(&val.shouldStart, 0)
			if start {
				val.Start(val.account.Address, val.deposit)
			}
			// unsubscribe. we're only interested in this event once
			events.Unsubscribe()
			// stop immediately and ignore all further pending events
			break out
		}
	}
}

func (val *Validator) Start(coinbase common.Address, deposit uint64) {
	if val.Started() {
		log.Warn("Failed to start the validator - the state machine is already running")
		return
	}

	account := accounts.Account{Address: coinbase}
	wallet, err := val.backend.AccountManager().Find(account)
	if err != nil {
		log.Crit("Failed to find a wallet", "err", err, "account", account.Address)
	}

	atomic.StoreInt32(&val.shouldStart, 1)

	val.wallet = wallet
	val.account = account
	val.deposit = deposit

	voter, err := val.network.IsVoter(&bind.CallOpts{}, val.account.Address)
	if err != nil {
		log.Crit("Failed to verify if the validator is registered as a voter", "err", err)
	}

	// @NOTE (rgeraldes) - initial genesis validators are registered as voters from the start
	// and we can use that info to verify if we need to sync
	if voter {
		// terminate sync go-routine
		val.eventMux.Post(downloader.DoneEvent{})
	} else {
		if atomic.LoadInt32(&val.canStart) == 0 {
			log.Info("Network syncing, will start validator afterwards")
			return
		}
	}

	log.Info("Starting validation operation")
	atomic.StoreInt32(&val.running, 1)

	go val.run()
}

func (val *Validator) run() {
	val.wg.Add(1)
	defer val.wg.Done()

	log.Info("Starting the consensus state machine")
	for state, numTransitions := val.notLoggedInState, 0; state != nil; numTransitions++ {
		state = state()
		if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
			break
		}
	}
}

func (val *Validator) Stop() {
	log.Info("Stopping consensus validator")

	val.withdraw()
	val.wg.Wait() // waits until the validator is no longer registered as a voter.

	atomic.StoreInt32(&val.shouldStart, 0)
	atomic.StoreInt32(&val.running, 0)

	log.Info("Consensus validator stopped")
}

func (val *Validator) SetExtra(extra []byte) error { return nil }

func (val *Validator) Validating() bool {
	return atomic.LoadInt32(&val.validating) > 0
}

func (val *Validator) Started() bool {
	return atomic.LoadInt32(&val.running) > 0
}

func (val *Validator) SetCoinbase(addr common.Address) {
	val.account.Address = addr
}

func (val *Validator) SetDeposit(deposit uint64) {
	val.deposit = deposit
}

// Pending returns the currently pending block and associated state.
func (val *Validator) Pending() (*types.Block, *state.StateDB) {
	// @TODO (rgeraldes) - review
	// val.currentMu.Lock()
	// defer val.currentMu.Unlock()

	state, err := val.chain.State()
	if err != nil {
		log.Crit("Failed to fetch the latest state", "err", err)
	}

	return val.chain.CurrentBlock(), state
}

func (val *Validator) PendingBlock() *types.Block {
	// @TODO (rgeraldes) - review
	// val.currentMu.Lock()
	// defer val.currentMu.Unlock()

	return val.chain.CurrentBlock()
}

func (val *Validator) restoreLastCommit() {
	currentBlock := val.chain.CurrentBlock()
	if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
		return
	}

	// @TODO (rgeraldes) - review the following statement
	lastValidators := types.NewValidatorSet([]*types.Validator{})

	lastCommit := currentBlock.LastCommit()
	lastPreCommits := core.NewVotingTable(val.eventMux, val.signer, currentBlock.Number(), lastCommit.Round(), types.PreCommit, lastValidators)
	for _, preCommit := range lastCommit.Commits() {
		if preCommit == nil {
			continue
		}
		added, err := lastPreCommits.Add(preCommit, false)
		if !added || err != nil {
			// @TODO (rgeraldes) - this should not happen > complete
			log.Error("Failed to restore the latest commit")
		}
	}

	val.lastCommit = lastPreCommits
}

func (val *Validator) init() error {
	parent := val.chain.CurrentBlock()

	val.blockNumber = parent.Number().Add(parent.Number(), big.NewInt(1))
	val.round = 0

	// @NOTE (rgeraldes) - in order to sync the nodes, the start time
	// must be the timestamp on the block + a sync interval. Tendermint offers another
	// option. Review
	val.start = time.Unix(parent.Time().Int64(), 0).Add(time.Duration(params.SyncDuration) * time.Millisecond)
	val.proposal = nil
	val.block = nil
	val.blockFragments = nil

	val.lockedRound = 0
	val.lockedBlock = nil
	val.commitRound = -1

	// validators
	count, err := val.network.GetVoterCount(&bind.CallOpts{})
	if err != nil {
		return err
	}
	validators := make([]*types.Validator, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := val.network.GetVoterAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return err
		}
		validators[i] = types.NewValidator(validator.Addr, validator.Deposit.Uint64())
	}

	// @TODO (rgeraldes) - update list of trusted peers based on the validators
	val.validators = types.NewValidatorSet(validators)

	// voting system
	val.votingSystem = NewVotingSystem(val.eventMux, val.signer, val.blockNumber, val.validators)

	// @TODO (rgeraldes) - last validators
	// val.lastValidators

	// events
	val.blockCh = make(chan *types.Block)
	val.majority = val.eventMux.Subscribe(core.NewMajorityEvent{})

	// @TODO (rgeraldes) - review vs go-eth
	if err = val.makeCurrent(parent); err != nil {
		log.Error("Failed to create mining context", "err", err)
		return nil
	}

	return nil
}

func (val *Validator) isProposer() bool {
	return val.validators.Proposer() == val.account.Address
}

func (val *Validator) AddProposal(proposal *types.Proposal) {
	log.Info("Received Proposal")
	// @TODO (rgeraldes) - Add proposal validation

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

	val.proposal = proposal
	val.blockFragments = types.NewDataSetFromMeta(proposal.BlockMetadata())
}

func (val *Validator) AddVote(vote *types.Vote) {
	if err := val.addVote(vote); err != nil {
		switch err {
		}
	}
}

func (val *Validator) addVote(vote *types.Vote) error {
	// @NOTE (rgeraldes) - for now just pre-vote/pre-commit for the current block number
	added, err := val.votingSystem.Add(vote, false)
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

func (val *Validator) makeDeposit() error {
	min, err := val.network.MinDeposit(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if min.Cmp(big.NewInt(int64(val.deposit))) > 0 {
		return fmt.Errorf("Current deposit - %d - is not enough. The minimum required is %d", val.deposit, min)
	}

	availability, err := val.network.Availability(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if !availability {
		return fmt.Errorf("There are not positions available at the moment")
	}

	opts := &bind.TransactOpts{
		From:  val.account.Address,
		Value: min.Mul(min, big.NewInt(2)),
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// @NOTE (rgeraldes) - ignore the proposed signer as by default it will be a unprotected signer.

			if address != val.account.Address {
				return nil, errors.New("not authorized to sign this account")
			}

			return val.wallet.SignTx(val.account, tx, val.config.ChainID)
		},
		// @TODO (rgeraldes) - price prediction & limits
		GasPrice: big.NewInt(25),
		GasLimit: big.NewInt(600000),
	}

	_, err = val.network.Deposit(opts)
	if err != nil {
		log.Error("Transaction failed", "err", err)
		return fmt.Errorf("Failed to transact the deposit: %x", err)
	}

	return nil
}

func (val *Validator) withdraw() {
	opts := &bind.TransactOpts{
		From: val.account.Address,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// @NOTE (rgeraldes) - ignore the proposed signer as by default it will be a unprotected signer.

			if address != val.account.Address {
				return nil, errors.New("not authorized to sign this account")
			}

			return val.wallet.SignTx(val.account, tx, val.config.ChainID)
		},
		// @TODO (rgeraldes) - price prediction & limits
		GasPrice: big.NewInt(25),
		GasLimit: big.NewInt(600000),
	}
	_, err := val.network.Withdraw(opts)
	if err != nil {
		log.Error("Failed to withdraw from the election", "err", err)
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
	val.header = header

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

	if err := val.engine.Prepare(val.chain, header); err != nil {
		log.Error("Failed to prepare header for mining", "err", err)
		// @TODO (rgeraldes) - review returning nil
		return nil
	}

	pending, err := val.backend.TxPool().Pending()
	if err != nil {
		log.Crit("Failed to fetch pending transactions", "err", err)
	}

	txs := types.NewTransactionsByPriceAndNonce(val.signer, pending)
	val.commitTransactions(val.eventMux, txs, val.chain, val.account.Address)

	// Create the new block to seal with the consensus engine
	var block *types.Block
	if block, err = val.engine.Finalize(val.chain, header, val.state, val.txs, commit, val.receipts); err != nil {
		log.Crit("Failed to finalize block for sealing", "err", err)
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
		// @TODO(rgeraldes) - analyse consequences
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
	// @TODO(rgeraldes) - review types int/uint
	for i := uint(0); i < blockFragments.Size(); i++ {
		val.eventMux.Post(core.NewBlockFragmentEvent{
			BlockNumber: val.blockNumber,
			Round:       val.round,
			Data:        blockFragments.Get(int(i)),
		})
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
		log.Debug("Majority of validators pre-voted nil")
		// unlock locked block
		if val.lockedBlock != nil {
			val.lockedRound = 0
			val.lockedBlock = nil
		}
	// majority pre-voted the locked block
	case winner == val.lockedBlock.Hash():
		log.Debug("Majority of validators pre-voted the locked block")
		// update locked block round
		val.lockedRound = val.round
		// vote on the pre-vote election winner
		vote = winner
	// majority pre-voted the proposed block
	case winner == val.block.Hash():
		log.Debug("Majority of validators pre-voted the proposed block")
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

	val.votingSystem.Add(signedVote, true)
}

// @TODO (rgeraldes) - review the round argument
func (val *Validator) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) {
	val.blockFragments.Add(fragment)

	// @NOTE (rgeraldes) - the whole section needs to be refactored
	if val.blockFragments.HasAll() {
		block, err := val.blockFragments.Assemble()
		if err != nil {
			log.Crit("Failed to assemble the block", "err", err)
		}

		// @TODO (rgeraldes) - refactor ; based on core/blockchain.go (InsertChain)
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
		}

		// @NOTE(rgeraldes) - ignore for now (assume that the block is ok)
		/*
			if err != nil {
				if err == ErrKnownBlock {
					stats.ignored++
					continue
				}

				if err == consensus.ErrFutureBlock {
					// Allow up to MaxFuture second in the future blocks. If this limit
					// is exceeded the chain is discarded and processed at a later time
					// if given.
					max := big.NewInt(time.Now().Unix() + maxTimeFutureBlocks)
					if block.Time().Cmp(max) > 0 {
						return i, fmt.Errorf("future block: %v > %v", block.Time(), max)
					}
					bc.futureBlocks.Add(block.Hash(), block)
					stats.queued++
					continue
				}

				if err == consensus.ErrUnknownAncestor && bc.futureBlocks.Contains(block.ParentHash()) {
					bc.futureBlocks.Add(block.Hash(), block)
					stats.queued++
					continue
				}

				bc.reportBlock(block, nil, err)
				return i, err
			}
		*/
		parent := val.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)

		// Process block using the parent state as reference point.
		receipts, _, usedGas, err := val.chain.Processor().Process(block, val.state, val.vmConfig)
		if err != nil {
			log.Crit("Failed to process the block", "err", err)
			//bc.reportBlock(block, receipts, err)
			//return i, err
		}
		val.receipts = receipts

		// Validate the state using the default validator
		err = val.chain.Validator().ValidateState(block, parent, val.state, receipts, usedGas)
		if err != nil {
			log.Crit("Failed to validate the state", "err", err)
			//bc.reportBlock(block, receipts, err)
			//return i, err
		}

		val.block = block

		go func() { val.blockCh <- block }()
	}
}

func (val *Validator) makeCurrent(parent *types.Block) error {
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
