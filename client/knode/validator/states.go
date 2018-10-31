package validator

import (
	"bytes"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

// work is the proposer current environment and holds all of the current state information
type work struct {
	state    *state.StateDB
	header   *types.Header
	tcount   int
	txs      []*types.Transaction
	receipts []*types.Receipt
}

type stateFn func() stateFn

func (val *validator) genesisNotLoggedInState() stateFn {
	// no need to make a deposit if the block number is 0
	// since these validators will be marked as voters from the start
	if val.isBlockZero() {
		log.Info("Deposit is not necessary for a genesis validator (first block)")
		return val.startValidating
	}
	return val.notLoggedInState
}

func (val *validator) notLoggedInState() stateFn {
	if err := val.tryJoin(); err != nil {
		log.Warn("stopping validation", "err", err)
		return nil
	}

	return val.startValidating
}

func (val *validator) isValidator() bool {
	isValidator, err := val.consensus.IsValidator(val.walletAccount.Account().Address)
	if err != nil {
		log.Crit("Failed to verify if account is already a validator")
	}
	return isValidator
}

func (val *validator) tryJoin() error {
	var err error

	isValidator := val.isValidator()
	if !isValidator {
		try := 0
		err = val.join()
		for err != nil {
			log.Error("Failed to make deposit", "try", try, "err", err)

			isValidator = val.isValidator()
			if isValidator {
				break
			}

			err = val.join()

			if try >= 500 {
				return fmt.Errorf("failed to make deposit. try %d", try)
			}
			try++
		}
		log.Warn("started as a validator")
	}
	return nil
}

func (val *validator) startValidating() stateFn {
	log.Info("Starting validation operation")
	atomic.StoreInt32(&val.validating, 1)

	log.Info("Voter has been accepted in the election",
		"enode", val.walletAccount.Account().Address.String(), "deposit", val.deposit.Int64())

	return val.newElectionState
}

func (val *validator) isBlockZero() bool {
	return val.chain.CurrentBlock().NumberU64() == 0
}

func (val *validator) newElectionState() stateFn {
	log.Info("Starting a new election")
	// update state machine based on current state
	if err := val.init(); err != nil {
		log.Error("validator init error", "err", err)
		return nil
	}

	// @NOTE (rgeraldes) - wait for txs - sync genesis validators, round zero for the first block only.
	if val.blockNumber.Cmp(big.NewInt(1)) == 0 {
		numTxs, _ := val.backend.TxPool().Stats() //
		if val.round == 0 && numTxs == 0 {
			log.Info("Waiting for a TX")
			txCh := make(chan core.NewTxsEvent)
			txSub := val.backend.TxPool().SubscribeNewTxsEvent(txCh)
			defer txSub.Unsubscribe()
			<-txCh
		}
	}

	return val.newRoundState
}

func (val *validator) newRoundState() stateFn {
	log.Debug("newRoundState", "time", time.Now().Format("2006-01-02T15:04:05.000"), "block", val.blockNumber, "round", val.round)
	val.proposal = nil
	val.isProposal = false
	val.proposer = nil
	val.block = nil

	parent := val.chain.CurrentBlock()
	parentNumber := new(big.Int).Set(parent.Number())
	val.blockNumber = parentNumber.Add(parentNumber, big.NewInt(1))

	val.timers = newTimers(val.previousRoundCreatedAt, val.blockNumber, val.round)

	val.blockFragmentsLock.Lock()
	val.blockFragments = make(map[common.Hash]*types.BlockFragments)
	val.blockFragmentsStorage = make(map[common.Hash][]*types.BlockFragment)
	val.blockFragmentsLock.Unlock()

	if val.majority.Closed() {
		val.blockCh = make(chan *types.Block, 1)
		val.majority = val.eventMux.Subscribe(core.NewMajorityEvent{})
	}

	currentRoundCreatedAt := time.Now()
	if val.round > 0 {
		val.previousRoundCreatedAt = val.roundCreatedAt
	}
	val.roundCreatedAt = currentRoundCreatedAt

	log.Debug("round created at", "time", currentRoundCreatedAt.Format(millisecondFormat), "number", val.blockNumber.Int64(), "round", val.round)
	log.Debug("previous round created at", "time", val.previousRoundCreatedAt.Format(millisecondFormat), "number", val.blockNumber.Int64(), "round", val.round)

	atomic.StoreInt32(&val.waitProposalBlock, 1)
	if val.round != 0 {
		val.checkUpdateValidators(false)
		if err := val.makeCurrent(parent); err != nil {
			log.Error("Failed to create mining context", "err", err)
			return nil
		}
	}

	return val.newProposalState
}

func (val *validator) newProposalState() stateFn {
	defer func() {
		<-val.timers.getProposingTimers().endsAt.C

		atomic.StoreInt32(&val.waitProposalBlock, 0)
		log.Info("proposal stage is done",
			"account", val.walletAccount.Account().Address.String(),
			"chainID", val.config.ChainID.Int64(),
			"block", val.blockNumber.Int64(),
			"round", val.round)
	}()

	val.proposer = val.voters.NextProposer()

	log.Debug("a new proposer",
		"addr", val.proposer.Address().String(),
		"isSelf", val.walletAccount.Account().Address == val.proposer.Address(),
		"block", val.blockNumber.Int64(),
		"round", val.round)

	if val.proposer.Address() == val.walletAccount.Account().Address {
		log.Info("Proposing a new block")
		val.propose()
	} else {
		<-val.timers.getRoundStartTimer().C
		log.Info("Waiting for the proposal", "addr", val.proposer.Address())
		val.waitForProposal()
	}

	return val.preVoteState
}

func (val *validator) waitForProposal() {
	for {
		select {
		case block := <-val.blockCh:
			if val.blockNumber.Cmp(block.Number()) == 0 {
				val.block = block
				log.Debug("Received the block", "blockNumber", val.block.Number().Int64(),
					"valBlockNumber", val.blockNumber.Int64(), "hash", val.block.Hash())
				return
			}

			log.Debug("unexpected proposed block number",
				"want", val.blockNumber.Int64(),
				"got", block.Number().Int64())
		case <-val.timers.getProposingTimers().deadline.C:
			return
		}
	}
}

func (val *validator) getProposeTimeout() uint64 {
	return params.ProposeDuration + val.round*params.ProposeDeltaDuration
}

func (val *validator) preVoteState() stateFn {
	log.Info("Pre vote sub-election")

	val.preVote()

	return val.preVoteWaitState
}

func (val *validator) preVoteWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-vote sub-election")

	defer func() { <-val.timers.getPreVotingTimers().endsAt.C }()

	select {
	case <-val.majority.Chan():
		log.Info("There's a majority in the pre-vote sub-election!")
		// fixme shall we do something here with current stateDB?
	case <-val.timers.getPreVotingTimers().deadline.C:
	}

	return val.preCommitState
}

func (val *validator) preCommitState() stateFn {
	log.Info("Pre commit sub-election")

	val.preCommit()

	return val.preCommitWaitState
}

func (val *validator) preCommitWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-commit sub-election", "blockNumber", val.blockNumber, "round", val.round)

	defer func() { <-val.timers.getPreCommitTimers().endsAt.C }()

	defer val.majority.Unsubscribe()

	select {
	case blockEvent, ok := <-val.majority.Chan():
		log.Info("There's a majority in the pre-commit sub-election!",
			"blockNumber", val.blockNumber,
			"round", val.round,
			"chan", ok,
			"blockEvent", blockEvent)
		if !ok {
			log.Debug("precommit state. stop watching val.majority. starting a new round")
			val.round++
			return val.newRoundState
		}

		if val.block.IsEmpty() {
			if val.block != nil {
				log.Debug("No one block wins!", "isNil", val.block == nil, "isEmpty", bytes.Equal(val.block.Hash().Bytes(), common.Hash{}.Bytes()))
			} else {
				log.Debug("No one block wins!", "isNil", val.block == nil)
			}

			val.round++
			return val.newRoundState
		}
		return val.commitState
	case <-val.timers.getPreCommitTimers().deadline.C:
		val.round++
		return val.newRoundState
	}
}

func (val *validator) commitState() stateFn {
	log.Info("Commit state")
	electedBlock := val.block

	blockHash := electedBlock.Hash()

	// update block hash since it is now available and not when
	// the receipt/log of individual transactions were created
	for _, r := range val.work.receipts {
		for _, l := range r.Logs {
			l.BlockHash = blockHash
		}
	}
	for _, log := range val.work.state.Logs() {
		log.BlockHash = blockHash
	}

	_, err := val.chain.WriteBlockWithState(electedBlock, val.work.receipts, val.work.state)
	if err != nil {
		log.Error("Failed writing block to chain", "err", err)
		return nil
	}

	// Broadcast the block and announce chain insertion event
	val.eventPoster.EventPost(core.NewMinedBlockEvent{Block: electedBlock})

	var (
		events []interface{}
		logs   = val.work.state.Logs()
	)

	events = append(events, core.ChainEvent{Block: electedBlock, Hash: electedBlock.Hash(), Logs: logs})
	events = append(events, core.ChainHeadEvent{Block: electedBlock})
	val.chain.PostChainEvents(events, logs)

	log.Info("Commit state done", "hash", electedBlock.Hash(), "block", electedBlock.Number().Int64(), "round", val.round)

	// election state updates
	val.commitRound = int(val.round)

	voter := val.isValidator()
	if !voter {
		log.Info(fmt.Sprintf("Logging out. Account %q is not a validator", val.walletAccount.Account().Address.String()))
		return val.loggedOutState
	}

	return val.newElectionState
}

func (val *validator) loggedOutState() stateFn {
	log.Info("Logged out")

	atomic.StoreInt32(&val.validating, 0)
	return nil
}
