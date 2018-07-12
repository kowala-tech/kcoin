package validator

import (
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

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

func (val *validator) notLoggedInState() stateFn {
	isGenesis, err := val.consensus.IsGenesisValidator(val.walletAccount.Account().Address)
	if err != nil {
		log.Warn("Failed to verify the voter information", "err", err)
		return nil
	}

	// @NOTE (rgeraldes) - sync was already done at this point and by default the investors will be
	// part of the initial set of validators - no need to make a deposit if the block number is 0
	// since these validators will be marked as voters from the start
	if !isGenesis || (isGenesis && val.chain.CurrentBlock().NumberU64() > 0) {
		chainHeadCh := make(chan core.ChainHeadEvent)
		chainHeadSub := val.chain.SubscribeChainHeadEvent(chainHeadCh)
		defer chainHeadSub.Unsubscribe()

		if err := val.consensus.Join(val.walletAccount, val.deposit); err != nil {
			log.Error("Error joining validators network", "err", err)
			return nil
		}

		log.Info("Waiting confirmation to participate in the consensus")
	L:
		for {
			select {
			case _, ok := <-chainHeadCh:
				if !ok {
					return nil
				}

				confirmed, err := val.consensus.IsValidator(val.walletAccount.Account().Address)
				if err != nil {
					log.Crit("Failed to verify the voter registration", "err", err)
				}

				if confirmed {
					break L
				}
			}
		}
	} else {
		isVoter, err := val.consensus.IsValidator(val.walletAccount.Account().Address)
		if err != nil {
			log.Crit("Failed to verify the voter information", "err", err)
			return nil
		}
		if !isVoter {
			log.Crit("Invalid genesis - genesis validator needs to be registered as a voter", "address", val.walletAccount.Account().Address)
		}

		log.Info("Deposit is not necessary for a genesis validator (first block)")
	}

	log.Info("Starting validation operation")
	atomic.StoreInt32(&val.validating, 1)

	log.Info("Voter has been accepted in the election", "enode", val.walletAccount.Account().Address.String())
	val.restoreLastCommit()

	return val.newElectionState
}

func (val *validator) newElectionState() stateFn {
	log.Info("Starting a new election")
	// update state machine based on current state
	if err := val.init(); err != nil {
		return nil
	}

	<-time.NewTimer(val.start.Sub(time.Now())).C

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
	log.Info("Starting a new voting round", "start time", val.start, "block number", val.blockNumber, "round", val.round)

	val.voters.NextProposer()

	if val.round != 0 {
		val.round++
		val.proposal = nil
		val.block = nil
		val.blockFragments = nil
	}

	return val.newProposalState
}

func (val *validator) newProposalState() stateFn {
	proposer := val.voters.NextProposer()
	if proposer.Address() == val.walletAccount.Account().Address {
		log.Info("Proposing a new block")
		val.propose()
	} else {
		log.Info("Waiting for the proposal", "addr", proposer.Address())
		val.waitForProposal()
	}
	return val.preVoteState
}

func (val *validator) waitForProposal() {
	timeout := time.Duration(params.ProposeDuration+val.round*params.ProposeDeltaDuration) * time.Millisecond
	select {
	case block := <-val.blockCh:
		val.block = block
		log.Info("Received the block", "hash", val.block.Hash())
	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
	}
}

func (val *validator) preVoteState() stateFn {
	log.Info("Pre vote sub-election")
	val.preVote()

	return val.preVoteWaitState
}

func (val *validator) preVoteWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-vote sub-election")
	timeout := time.Duration(params.PreVoteDuration+val.round*params.PreVoteDeltaDuration) * time.Millisecond

	select {
	case <-val.majority.Chan():
		log.Info("There's a majority in the pre-vote sub-election!")
	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
	}

	return val.preCommitState
}

func (val *validator) preCommitState() stateFn {
	log.Info("Pre commit sub-election")
	val.preCommit()

	return val.preCommitWaitState
}

func (val *validator) preCommitWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-commit sub-election")
	timeout := time.Duration(params.PreCommitDuration+val.round+params.PreCommitDeltaDuration) * time.Millisecond
	defer val.majority.Unsubscribe()

	select {
	case <-val.majority.Chan():
		log.Info("There's a majority in the pre-commit sub-election!")
		if val.block == nil {
			return val.newRoundState
		}
		return val.commitState
	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
		return val.newRoundState
	}
}

func (val *validator) commitState() stateFn {
	log.Info("Commit state")

	block := val.block
	work := val.work

	_, err := work.state.Commit(true)
	if err != nil {
		log.Error("Failed writing block to chain", "err", err)
		return nil
	}

	// update block hash since it is now available and not when
	// the receipt/log of individual transactions were created
	for _, r := range val.work.receipts {
		for _, l := range r.Logs {
			l.BlockHash = block.Hash()
		}
	}
	for _, log := range val.work.state.Logs() {
		log.BlockHash = block.Hash()
	}

	_, err = val.chain.WriteBlockWithState(block, val.work.receipts, val.work.state)
	if err != nil {
		log.Error("Failed writing block to chain", "err", err)
		return nil
	}

	// Broadcast the block and announce chain insertion event
	go val.eventMux.Post(core.NewMinedBlockEvent{Block: block})
	var (
		events []interface{}
		logs   = work.state.Logs()
	)
	events = append(events, core.ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})
	events = append(events, core.ChainHeadEvent{Block: block})
	val.chain.PostChainEvents(events, logs)

	// election state updates
	val.commitRound = int(val.round)

	voter, err := val.consensus.IsValidator(val.walletAccount.Account().Address)
	if err != nil {
		log.Crit("Failed to verify if the validator is a voter", "err", err)
	}
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
