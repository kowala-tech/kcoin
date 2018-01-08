package validator

import (
	"math/big"
	"time"

	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

// election encapsulates the consensus state for a specific block election
type election struct {
	blockNumber *big.Int
	round       int

	validators    *types.Validators
	proposal      *types.Proposal
	proposalBlock *types.Block

	lockedRound int
	lockedBlock *types.Block

	start time.Time // used to sync the validator nodes

	commitRound int

	lastCommit     *core.VotingTable // Last precommits at Height-1
	lastValidators *types.Validators

	// proposer
	tcount   int
	txs      []*types.Transaction
	receipts []*types.Receipt
}

// stateFn represents a state function
type stateFn func() stateFn

// @NOTE (rgeraldes) - initial state
func (val *Validator) notLoggedInState() stateFn {
	log.Info("Initial state")

	/*
		// @TODO (rgeraldes) - to do as soon as the contract for the dynamic validator is up
		// one shot type of confirmation loop
		headSub := val.eventMux.Subscribe(ChainHeadEvent{})
		defer headSub.Unsubscribe()
		// @NOTE (rgeraldes) - calls pos contract
		val.deposit()

		for {
			select {
				ev := <- latestBlock.Chan():
				// if transaction was committed
					// if successful return val.newBlockState
					// if not successful (ex: not enough funds) return nil
					// log.Error()
			}
		}
	*/

	val.restoreLastCommit()

	return val.newElectionState
}

func (val *Validator) newElectionState() stateFn {
	log.Info("Starting the election for a new block", "block number", val)

	// update state machine based on current state
	val.init()

	<-time.NewTimer(val.start.Sub(time.Now())).C

	// @NOTE (rgeraldes) - wait for txs to be available in the txPool for the first block
	// before we enter the proposal state. If the last block changed the app hash,
	// we may need an empty "proof" block, and the proposal state immediately.
	// @TODO (replace with configuration) - as soon as we are aware
	// of the full use case
	wait := true //&& !cs.needProofBlock(height)
	numTxs, _ := val.kusd.TxPool().Stats()
	if val.blockNumber.Cmp(big.NewInt(0)) == 0 && numTxs > 0 && wait {
		log.Info("Waiting for transactions")
		txSub := val.eventMux.Subscribe(core.TxPreEvent{})
		defer txSub.Unsubscribe()
		<-txSub.Chan()
	}

	return val.newRoundState
}

func (val *Validator) newRoundState() stateFn {
	log.Info("Starting a new election round", "start time", val.start, "block number", val.blockNumber, "round", val.round)

	if val.round != 0 {
		val.proposal = nil
		val.proposalBlock = nil
		//		val.proposalBlockFragments = nil
	}
	//	val.votes.SetRound(val.round + 1) // also track next round (round+1) to allow round-skipping

	return val.newProposalState
}

func (val *Validator) newProposalState() stateFn {
	log.Info("Starting a new proposal")

	timeout := time.Duration(params.ProposeDuration+uint64(val.round)*params.ProposeDeltaDuration) * time.Millisecond

	if val.isProposer() {
		log.Info("Proposing a new block", "hash", val.proposal.Hash())
		val.propose()
	} else {
		select {
		//case val.proposalSub.Chan():
		//	log.Info("Received a new proposal", "block number", val.proposal.BlockNumber(), "hash", val.proposal.Hash())
		case <-time.After(timeout):
			log.Info("Timeout expired", "duration", timeout)
		}
	}

	return val.preVoteState
}

func (val *Validator) preVoteState() stateFn {
	log.Info("Starting the pre vote election")

	val.prevote()

	return val.preVoteWaitState
}

func (val *Validator) preVoteWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-vote election")

	timeout := time.Duration(params.PreVoteDuration+uint64(val.round)*params.PreVoteDeltaDuration) * time.Millisecond

	select {
	//case val.preVoteMajSub.Chan():
	//	log.Info("There's a majority")

	case <-time.After(timeout):
		log.Info("Timeout expired", "duratiom", timeout)
	}

	return val.preCommitState
}

func (val *Validator) preCommitState() stateFn {
	log.Info("Starting the pre commit election")

	val.precommit()

	return val.preCommitWaitState
}

func (val *Validator) preCommitWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-commit election")

	timeout := time.Duration(params.PreCommitDuration+uint64(val.round)+params.PreCommitDeltaDuration) * time.Millisecond

	select {
	/*
		case val.preCommitMajSub.Chan():
			log.Info("There's a majority")

			if val.proposalBlock == nil {
				return val.newRoundState
			}
			return val.commitState
	*/
	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
		return val.newRoundState
	}
}

func (val *Validator) commitState() stateFn {
	log.Info("Committing the election result")

	stat, err := val.chain.WriteBlock(val.proposalBlock)
	if err != nil {
		log.Crit("Failed writing block to chain", "err", err)
	}

	// @TODO (rgeraldes) - review type
	// state updates
	val.commitRound = int(val.round)

	// @TODO(rgeraldes)
	/*
		// leaves only when it has all the pre commits
		select {}
	*/

	// @TODO (rgeraldes) - VALIDATOR CONTRACT

	// @TODO if the validator is not part of the validator state anymore, return nil
	// if not part of the voters jump to left election
	// return val.leftElectionsState
	// return val.newElectionState

	return val.newElectionState
}

// @NOTE (rgeraldes) - end state
func (val *Validator) loggedOutState() stateFn {
	val.wg.Done()
	return nil
}
