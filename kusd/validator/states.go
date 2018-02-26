package validator

import (
	"math/big"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

// @TODO (rgeraldes) - confirm
// work is the proposer current environment and holds all of the current state information
type work struct {
	state    *state.StateDB
	header   *types.Header
	tcount   int
	txs      []*types.Transaction
	receipts []*types.Receipt
}

// stateFn represents a state function
type stateFn func() stateFn

// @NOTE (rgeraldes) - initial state
func (val *validator) notLoggedInState() stateFn {
	isGenesis, err := val.network.IsGenesisVoter(&bind.CallOpts{}, val.walletAccount.Account().Address)
	if err != nil {
		log.Crit("Failed to verify the voter information", "err", err)
		return nil
	}

	// @NOTE (rgeraldes) - sync was already done at this point and by default the investors will be
	// part of the initial set of validators - no need to make a deposit if the block number is 0
	// since these validators will be marked as voters from the start
	if !isGenesis || (isGenesis && val.chain.CurrentBlock().NumberU64() > 0) {
		// Subscribe events from blockchain
		chainHeadCh := make(chan core.ChainHeadEvent)
		chainHeadSub := val.chain.SubscribeChainHeadEvent(chainHeadCh)
		defer chainHeadSub.Unsubscribe()

		log.Info("Making Deposit")
		if err := val.makeDeposit(); err != nil {
			return nil
		}

		log.Info("Waiting confirmation to participate in the consensus")
	L:
		for {
			select {
			case _, ok := <-chainHeadCh:
				if !ok {
					// @TODO (rgeraldes) - log
					return nil
				}

				confirmed, err := val.network.IsVoter(&bind.CallOpts{}, val.walletAccount.Account().Address)
				if err != nil {
					log.Crit("Failed to verify the voter registration", "err", err)
				}

				if confirmed {
					break L
				}
			}
		}
	} else {
		// sanity check
		isVoter, err := val.network.IsVoter(&bind.CallOpts{}, val.walletAccount.Account().Address)
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
		// @TODO (rgeraldes) - log
		return nil
	}

	<-time.NewTimer(val.start.Sub(time.Now())).C

	// @NOTE (rgeraldes) - wait for txs - sync genesis validators, round zero for the first block only.
	if val.blockNumber.Cmp(big.NewInt(1)) == 0 {
		numTxs, _ := val.backend.TxPool().Stats() //
		if val.round == 0 && numTxs == 0 {        //!cs.needProofBlock(height)
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
	log.Info("Starting a new voting round", "start time", val.start, "block number", val.blockNumber, "round", val.round)

	// updates the validators weight > proposer
	val.validators.UpdateWeight()

	if val.round != 0 {
		val.round++
		val.proposal = nil
		val.block = nil
		val.blockFragments = nil
	}

	//	val.votes.SetRound(val.round + 1) // also track next round (round+1) to allow round-skipping
	return val.newProposalState
}

func (val *validator) newProposalState() stateFn {
	timeout := time.Duration(params.ProposeDuration+val.round*params.ProposeDeltaDuration) * time.Millisecond

	if val.isProposer() {
		log.Info("Proposing a new block")
		val.propose()
	} else {
		log.Info("Waiting for the proposal", "proposer", val.validators.Proposer())
		select {
		case block := <-val.blockCh:
			val.block = block
			log.Info("Received the block", "hash", val.block.Hash())
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
	// @TODO (rgeraldes) - move to a post processor state
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
		// @TODO (rgeraldes) - confirm
		return val.commitState
	}
}

func (val *validator) commitState() stateFn {
	log.Info("Commit state")

	// @TODO (rgeraldes) - replace work with unconfirmed, unjustified?

	block := val.block
	work := val.work
	chainDb := val.backend.ChainDb()

	work.state.CommitTo(chainDb, true)

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

	_, err := val.chain.WriteBlockAndState(block, val.work.receipts, val.work.state)
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

	// @TODO(rgeraldes)
	// leaves only when it has all the pre commits
	voter, err := val.network.IsVoter(&bind.CallOpts{}, val.walletAccount.Account().Address)
	if err != nil {
		log.Crit("Failed to verify if the validator is a voter", "err", err)
	}
	if !voter {
		return val.loggedOutState
	}

	return val.newElectionState
}

// @NOTE (rgeraldes) - end state
func (val *validator) loggedOutState() stateFn {
	log.Info("Logged out")

	atomic.StoreInt32(&val.validating, 0)

	return nil
}
