package validator

import (
	"math/big"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

type stateFn func() stateFn

func (val *validator) notLoggedInState() stateFn {
	isGenesis, err := val.election.IsGenesisValidator(val.walletAccount.Account().Address)
	if err != nil {
		log.Warn("Failed to verify the voter information", "err", err)
		return nil
	}

	// @NOTE (rgeraldes) - sync was already done at this point and by default the investors will be
	// part of the initial set of validators - no need to make a deposit if the block number is 0
	// since these validators will be marked as voters from the start
	if !isGenesis || (isGenesis && val.chain.CurrentBlock().NumberU64() > 0) {
		if err := val.election.Join(val.walletAccount, val.deposit); err != nil {
			log.Error("Error joining validators network", "err", err)
			return nil
		}
	} else {
		isVoter, err := val.election.IsValidator(val.walletAccount.Account().Address)
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

	log.Info("Voter has been accepted in the election")
	val.restoreLastCommit()

	return val.newElectionState
}

func (val *validator) newElectionState() stateFn {
	log.Info("Starting a new election")
	// update state machine based on current state
	if err := val.init(); err != nil {
		return nil
	}

	<-time.NewTimer(val.election.start.Sub(time.Now())).C

	// @NOTE (rgeraldes) - wait for txs - sync genesis validators, round zero for the first block only.
	if val.election.blockNumber.Cmp(big.NewInt(1)) == 0 {
		numTxs, _ := val.backend.TxPool().Stats() //
		if val.election.round == 0 && numTxs == 0 {
			log.Info("Waiting for a TX")
			txCh := make(chan core.TxPreEvent)
			txSub := val.backend.TxPool().SubscribeTxPreEvent(txCh)
			defer txSub.Unsubscribe()
			<-txCh
		}
	}

	return val.newRoundState
}

func (val *validator) newRoundState() stateFn {
	log.Info("Starting a new voting round", "start time", val.election.start, "block number", val.election.blockNumber, "round", val.election.round)

	val.election.validators.UpdateWeights()

	if val.election.round != 0 {
		val.election.round++
		val.election.proposal = nil
		val.election.block = nil
		val.election.blockFragments = nil
	}

	return val.newProposalState
}

func (val *validator) newProposalState() stateFn {
	timeout := time.Duration(params.ProposeDuration+val.election.round*params.ProposeDeltaDuration) * time.Millisecond

	if val.isProposer() {
		log.Info("Proposing a new block")
		val.propose()
	} else {
		log.Info("Waiting for the proposal", "proposer", val.election.validators.Proposer())
		select {
		case event := <-val.proposalCh:
			val.block = event.Block
			log.Info("Received the block", "hash", val.election.block.Hash())
		case <-time.After(timeout):
			log.Info("Timeout expired", "duration", timeout)
		}
	}

	return val.preVoteState
}

func (val *validator) preVoteState() stateFn {
	log.Info("Pre vote sub-election")
	val.preVote()

	return val.preVoteWaitState
}

func (val *validator) preVoteWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-vote sub-election")
	timeout := time.Duration(params.PreVoteDuration+val.election.round*params.PreVoteDeltaDuration) * time.Millisecond

	select {
	case <-val.majorityCh:
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
	timeout := time.Duration(params.PreCommitDuration+val.election.round+params.PreCommitDeltaDuration) * time.Millisecond
	defer val.majoritySub.Unsubscribe()

	select {
	case <-val.majorityCh:
		log.Info("There's a majority in the pre-commit sub-election!")
		if val.election.block == nil {
			return val.newRoundState
		}
		return val.commitState
	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
		return val.commitState
	}
}

func (val *validator) commitState() stateFn {
	log.Info("Commit state")

	block := val.election.block
	work := val.work
	chainDb := val.backend.ChainDb()

	work.state.CommitTo(chainDb, true)

	// update block hash since it is now available and not when
	// the receipt/log of individual transactions were created
	for _, r := range work.receipts {
		for _, l := range r.Logs {
			l.BlockHash = block.Hash()
		}
	}
	for _, log := range work.state.Logs() {
		log.BlockHash = block.Hash()
	}

	_, err := val.chain.WriteBlockAndState(block, work.receipts, work.state)
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

	val.commitRound = int(val.election.round)

	voter, err := val.election.IsValidator(val.walletAccount.Account().Address)
	if err != nil {
		log.Crit("Failed to verify if the validator is a voter", "err", err)
	}
	if !voter {
		return val.loggedOutState
	}

	return val.newElectionState
}

func (val *validator) loggedOutState() stateFn {
	log.Info("Logged out")

	atomic.StoreInt32(&val.validating, 0)

	return nil
}
