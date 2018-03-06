package validator

import (
	"math/big"
	"fmt"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
	"sync"
	"time"
)

// validating represents a consensus validating state
type validating struct {
	Election           // consensus state
	maxTransitions int // max number of state transitions (tests) 0 - unlimited

	*context

	wg sync.WaitGroup
}

// NewValidating returns a new consensus validating
func newValidating(context *context) *validating {
	return &validating{context: context}
}

func (val *validating) Start() (Validator, error) {
	go func() {
		log.Info("Starting validation operation")
		val.wg.Add(1)

		defer val.wg.Done()

		log.Info("Starting the consensus state machine")
		for state, numTransitions := val.notLoggedInState, 0; state != nil; numTransitions++ {
			state = state()
			if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
				break
			}
		}
	}()

	return val, nil
}

func (val *validating) Stop() (Validator, error) {
	val.withdraw()
	val.wg.Wait() // waits until the validating is no longer registered as a voter.

	log.Info("Consensus validating stopped")
	return val, nil
}

func (val *validating) SetExtra(extra []byte) error {
	val.extra = extra
	return nil
}

func (val *validating) Validating() bool {
	return true
}

func (val *validating) SetCoinbase(address common.Address) error {
	return ErrCantSetCoinbaseOnStartedValidator
}

func (val *validating) SetDeposit(deposit uint64) error {
	return ErrCantSetDepositOnStartedValidator
}

// Pending returns the currently pending block and associated state.
func (val *validating) Pending() (*types.Block, *state.StateDB) {
	curState, err := val.chain.State()
	if err != nil {
		log.Crit("Failed to fetch the latest state", "err", err)
	}

	return val.chain.CurrentBlock(), curState
}

func (val *validating) PendingBlock() *types.Block {
	return val.chain.CurrentBlock()
}

func (val *validating) restoreLastCommit() {
	checksum, err := val.contract.VotersChecksum(&bind.CallOpts{})
	if err != nil {
		log.Crit("Failed to access the voters checksum", "err", err)
	}

	if err := val.updateValidators(checksum, true); err != nil {
		log.Crit("Failed to update the validating set", "err", err)
	}

	currentBlock := val.chain.CurrentBlock()
	if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
		return
	}
}

func (val *validating) init() error {
	parent := val.chain.CurrentBlock()

	checksum, err := val.contract.VotersChecksum(&bind.CallOpts{})
	if err != nil {
		log.Crit("Failed to access the voters checksum", "err", err)
	}

	if val.validatorsChecksum != checksum {
		if err := val.updateValidators(checksum, true); err != nil {
			log.Crit("Failed to update the validating set", "err", err)
		}
	}

	// @NOTE (rgeraldes) - start is not relevant for the first block as the first election will
	// wait until we have transactions
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

func (val *validating) isProposer() bool {
	return val.validators.Proposer() == val.walletAccount.Account().Address
}

func (val *validating) AddProposal(proposal *types.Proposal) error {
	if !val.Validating() {
		return ErrCantAddProposalNotValidating
	}

	log.Info("Received Proposal")

	val.proposal = proposal
	val.blockFragments = types.NewDataSetFromMeta(proposal.BlockMetadata())

	return nil
}

func (val *validating) AddVote(vote *types.Vote) error {
	if !val.Validating() {
		return ErrCantVoteNotValidating
	}

	if err := val.addVote(vote); err != nil {
		switch err {
		}
	}
	return nil
}

func (val *validating) addVote(vote *types.Vote) error {
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

func (val *validating) commitTransactions(mux *event.TypeMux, txs *types.TransactionsByPriceAndNonce, bc *core.BlockChain, coinbase common.Address) {
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
			// NewValidating head notification data race between the transaction pool and miner, shift
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

func (val *validating) commitTransaction(tx *types.Transaction, bc *core.BlockChain, coinbase common.Address, gp *core.GasPool) (error, []*types.Log) {
	snap := val.state.Snapshot()

	receipt, _, err := core.ApplyTransaction(val.chainConfig, bc, &coinbase, gp, val.state, val.header, tx, val.header.GasUsed, vm.Config{})
	if err != nil {
		val.state.RevertToSnapshot(snap)
		return err, nil
	}
	val.txs = append(val.txs, tx)
	val.receipts = append(val.receipts, receipt)

	return nil, receipt.Logs
}

func (val *validating) makeDeposit() error {
	min, err := val.contract.MinDeposit(&bind.CallOpts{})
	if err != nil {
		return err
	}

	var deposit big.Int
	if min.Cmp(deposit.SetUint64(val.deposit)) > 0 {
		log.Warn("Current deposit is not enough", "deposit", val.deposit, "minimum required", min)
		// @TODO (rgeraldes) - error handling?
		return fmt.Errorf("Current deposit - %d - is not enough. The minimum required is %d", val.deposit, min)
	}

	availability, err := val.contract.Availability(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if !availability {
		return fmt.Errorf("There are not positions available at the moment")
	}

	options := getTransactionOpts(val.walletAccount, deposit.SetUint64(val.deposit), val.chainConfig.ChainID)
	_, err = val.contract.Deposit(options)
	if err != nil {
		return fmt.Errorf("Failed to transact the deposit: %x", err)
	}

	return nil
}

func (val *validating) withdraw() {
	options := getTransactionOpts(val.walletAccount, nil, val.chainConfig.ChainID)
	_, err := val.contract.Withdraw(options)
	if err != nil {
		log.Error("Failed to withdraw from the election", "err", err)
	}
}

func (val *validating) createProposalBlock() *types.Block {
	if val.lockedBlock != nil {
		log.Info("Picking a locked block")
		return val.lockedBlock
	}
	return val.createBlock()
}

func (val *validating) createBlock() *types.Block {
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
	val.commitTransactions(val.eventMux, txs, val.chain, val.walletAccount.Account().Address)

	// Create the new block to seal with the consensus engine
	var block *types.Block
	if block, err = val.engine.Finalize(val.chain, header, val.state, val.txs, commit, val.receipts); err != nil {
		log.Crit("Failed to finalize block for sealing", "err", err)
	}

	return block
}

func (val *validating) propose() {
	block := val.createProposalBlock()

	//lockedRound, lockedBlock := val.votes.LockingInfo()
	lockedRound := 1
	lockedBlock := common.Hash{}

	// @TODO (rgeraldes) - review int/int64; address situation where validators size might be zero (no peers)
	// @NOTE (rgeraldes) - (for now size = block size) number of block fragments = number of validators - self
	fragments, err := block.AsFragments(int(block.Size().Int64()) /*/val.validators.Size() - 1 */)
	if err != nil {
		// @TODO(rgeraldes) - analyse consequences
		log.Crit("Failed to get the block as a set of fragments of information", "err", err)
	}

	proposal := types.NewProposal(val.blockNumber, val.round, fragments.Metadata(), lockedRound, lockedBlock)

	signedProposal, err := val.walletAccount.SignProposal(val.walletAccount.Account(), proposal, val.chainConfig.ChainID)
	if err != nil {
		log.Crit("Failed to sign the proposal", "err", err)
	}

	val.proposal = signedProposal
	val.block = block

	val.eventMux.Post(core.NewProposalEvent{Proposal: proposal})

	// post block segments events
	// @TODO(rgeraldes) - review types int/uint
	for i := uint(0); i < fragments.Size(); i++ {
		val.eventMux.Post(core.NewBlockFragmentEvent{
			BlockNumber: val.blockNumber,
			Round:       val.round,
			Data:        fragments.Get(int(i)),
		})
	}

}

func (val *validating) preVote() {
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

func (val *validating) preCommit() {
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
	case winner == val.lockedBlock.Hash():
		log.Debug("Majority of validators pre-voted the locked block")
		// update locked block round
		val.lockedRound = val.round
		// vote on the pre-vote election winner
		vote = winner
	case winner == val.block.Hash():
		log.Debug("Majority of validators pre-voted the proposed block")
		// lock block
		val.lockedRound = val.round
		val.lockedBlock = val.block
		// vote on the pre-vote election winner
		vote = winner
		// we don't have the current block (fetch)
		// @TODO (tendermint): in the future save the POL prevotes for justification.
	default:
		// fetch block, unlock, precommit
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

func (val *validating) vote(vote *types.Vote) {
	signedVote, err := val.walletAccount.SignVote(val.walletAccount.Account(), vote, val.chainConfig.ChainID)
	if err != nil {
		log.Crit("Failed to sign the vote", "err", err)
	}

	val.votingSystem.Add(signedVote, true)
}

func (val *validating) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error {
	if !val.Validating() {
		return ErrCantAddBlockFragmentNotValidating
	}
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
		parent := val.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)

		// Process block using the parent state as reference point.
		receipts, _, usedGas, err := val.chain.Processor().Process(block, val.state, val.vmConfig)
		if err != nil {
			log.Crit("Failed to process the block", "err", err)
			//bc.reportBlock(block, receipts, err)
			//return i, err
		}
		val.receipts = receipts

		// Validate the state using the default validating
		err = val.chain.Validator().ValidateState(block, parent, val.state, receipts, usedGas)
		if err != nil {
			log.Crit("Failed to validate the state", "err", err)
			//bc.reportBlock(block, receipts, err)
			//return i, err
		}

		val.block = block

		go func() { val.blockCh <- block }()
	}
	return nil
}

func (val *validating) makeCurrent(parent *types.Block) error {
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

func (val *validating) updateValidators(checksum [32]byte, genesis bool) error {
	count, err := val.contract.GetVoterCount(&bind.CallOpts{})
	if err != nil {
		return err
	}

	val.validatorsChecksum = checksum
	validators := make([]*types.Validator, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := val.contract.GetVoterAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return err
		}

		var weight *big.Int
		weight = big.NewInt(0)

		validators[i] = types.NewValidator(validator.Addr, validator.Deposit.Uint64(), weight)
	}
	val.validators = types.NewValidatorSet(validators)
	val.validatorsChecksum = checksum

	return nil
}
