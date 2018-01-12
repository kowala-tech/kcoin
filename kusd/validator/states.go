package validator

import (
	"math/big"
	"time"

	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

// @TODO (rgeraldes) - unsubscribe majority subscriptions

// Election encapsulates the consensus state for a specific block election
type Election struct {
	blockNumber *big.Int
	round       uint64

	validators     *types.Validators
	proposal       *types.Proposal
	block          *types.Block
	blockFragments *types.BlockFragments
	votes          VotingSystem // election votes since round 1

	lockedRound uint64
	lockedBlock *types.Block

	start time.Time // used to sync the validator nodes

	commitRound int

	lastCommit     *core.VotingTable // Last precommits at current block number-1
	lastValidators *types.Validators

	// proposer
	tcount    int
	failedTxs types.Transactions
	txs       []*types.Transaction
	receipts  []*types.Receipt

	// inputs
	proposalCh                    chan *types.Proposal
	firstMajority, secondMajority *event.TypeMuxSubscription
}

// stateFn represents a state function
type stateFn func() stateFn

// @NOTE (rgeraldes) - initial state
func (val *Validator) notLoggedInState() stateFn {
	log.Info("Waiting confirmation to participate in the consensus")

	val.firstMajority = val.eventMux.Subscribe()
	val.secondMajority = val.eventMux.Subscribe()

	/*
		// @TODO (rgeraldes) - to do as soon as the contract for the dynamic validator is up
		// one shot type of confirmation loop
		headSub := val.eventMux.Subscribe(ChainHeadEvent{})
		defer headSub.Unsubscribe()
		// @NOTE (rgeraldes) - calls pos contract
		// @TODO (rgeraldes) - Genesis validators minimum deposit

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
	log.Info("Starting a new election")
	// update state machine based on current state
	val.init()

	<-time.NewTimer(val.start.Sub(time.Now())).C

	// @NOTE (rgeraldes) - wait for txs to be available in the txPool for the round 0
	// If the last block changed the app hash, we may need an empty "proof" block.
	numTxs, _ := val.kusd.TxPool().Stats() //
	if val.round == 0 && numTxs == 0 {     //!cs.needProofBlock(height)
		log.Info("Waiting for transactions")
		txSub := val.eventMux.Subscribe(core.TxPreEvent{})
		defer txSub.Unsubscribe()
		<-txSub.Chan()
	}

	return val.newRoundState
}

func (val *Validator) newRoundState() stateFn {
	log.Info("Starting a new election round", "start time", val.start, "block number", val.blockNumber, "round", val.round)
	val.round++

	if val.round != 1 {
		val.proposal = nil
		val.block = nil
		val.blockFragments = nil
	}

	//	val.votes.SetRound(val.round + 1) // also track next round (round+1) to allow round-skipping

	return val.newProposalState
}

func (val *Validator) newProposalState() stateFn {
	timeout := time.Duration(params.ProposeDuration+uint64(val.round)*params.ProposeDeltaDuration) * time.Millisecond

	if val.isProposer() {
		log.Info("Proposing a new block")
		val.propose()
	} else {
		log.Info("Waiting for the proposal", "proposer", nil)
		select {
		case proposal := <-val.proposalCh:
			val.proposal = proposal
			log.Info("Received a new proposal", "hash", val.proposal.Hash())
		case <-time.After(timeout):
			log.Info("Timeout expired", "duration", timeout)
		}
	}

	return val.preVoteState
}

func (val *Validator) preVoteState() stateFn {
	log.Info("Starting the pre vote election")
	val.preVote()

	return val.preVoteWaitState
}

func (val *Validator) preVoteWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-vote election")
	timeout := time.Duration(params.PreVoteDuration+uint64(val.round)*params.PreVoteDeltaDuration) * time.Millisecond

	select {

	case <-val.firstMajority.Chan():
		log.Info("There's a majority!")

	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
	}

	return val.preCommitState
}

func (val *Validator) preCommitState() stateFn {
	log.Info("Starting the pre commit election")
	val.preCommit()

	return val.preCommitWaitState
}

func (val *Validator) preCommitWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-commit election")
	timeout := time.Duration(params.PreCommitDuration+uint64(val.round)+params.PreCommitDeltaDuration) * time.Millisecond

	select {

	case <-val.secondMajority.Chan():
		log.Info("There's a majority!")
		if val.block == nil {
			return val.newRoundState
		}
		return val.commitState

	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
		return val.commitState
	}
}

func (val *Validator) commitState() stateFn {
	log.Info("Commit state")

	// @NOTE (rgeraldes) - this is full validation which should not be necessary at this stage.
	// @TODO (rgeraldes) - replace with just the necessary steps
	if _, err := val.chain.InsertChain(types.Blocks{val.block}); err != nil {
		log.Crit("Failure to commit the proposed block", "err", err)
	}
	go val.eventMux.Post(core.NewMinedBlockEvent{Block: val.block})

	// state updates
	val.commitRound = int(val.round)

	// @TODO(rgeraldes)
	// leaves only when it has all the pre commits

	// @TODO if the validator is not part of the new state log out
	// return val.leftElectionsState
	// return val.newElectionState

	return val.newElectionState
}

// @NOTE (rgeraldes) - end state
func (val *Validator) loggedOutState() stateFn {
	return nil
}
