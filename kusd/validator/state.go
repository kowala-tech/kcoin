package validator

import (
	"time"

	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

// state represents a consensus state
type State byte

const (
	NewBlock State = iota
	NewRound
	Proposal
	PreVote
	PreCommit
	Commit
)

// stateFn represents a state function
type stateFn func() stateFn

// deposit state makes a deposit to participate in the consenus and keeps
// track of the latest blocks in order to confirm the transaction.
func (self *Validator) InitialState() stateFn {
	log.Info("Initial state")
	// @TODO (rgeraldes) - to do as soon as the contract for the dynamic validator set is up

	/*
		// one shot type of confirmation loop
		latestBlockSub := self.eventMux.Subscribe(ChainHeadEvent{})
		defer latestBlockSub.Unsubscribe()

		// @NOTE (rgeraldes) - calls pos contract
		self.deposit()

		for {
			select {
				ev := <- latestBlock.Chan():
				// if transaction was committed
					// if successful return self.newBlockState
					// if not successful (ex: not enough funds) return nil
					// log.Error()
			}
		}
	*/

	self.restoreLastCommit()

	return self.newBlockState
}

func (self *Validator) newBlockState() stateFn {
	log.Info("Starting a new block", "block number", self.blockNumber)

	// update machine based on the blockchain state
	// @TODO (rgeraldes) - add new validator set
	self.init()

	<-time.NewTimer(self.start.Sub(time.Now())).C

	// @NOTE(rgeraldes) - wait for txs to be available in the txPool
	// before we enterPropose in round 0. If the last block changed the app hash,
	// we may need an empty "proof" block, and enterPropose immediately.
	// @TODO (replace with configuration) - as soon as we are aware
	// of the full usecase
	wait := true //&& !cs.needProofBlock(height)
	if len(self.kusd.TxPool().Pending()) > 0 && wait {
		log.Info("Waiting for transactions")
		txSub := self.eventMux.Subscribe(TxPreEvent{})
		defer txSub.Unsubscribe()
		<-txSub.Chan()
	}

	return self.newRoundState
}

func (self *Validator) newRoundState() stateFn {
	log.Info("Starting a new consensus round", "start time", self.start, "block number", self.blockNumber, "round", self.round)

	if self.round != 0 {
		self.proposal = nil
		self.proposalBlock = nil
		self.proposalBlockFragments = nil
	}
	self.votes.SetRound(self.round + 1) // also track next round (round+1) to allow round-skipping

	return self.newProposalState
}

func (self *Validator) newProposalState() stateFn {
	log.Info("Starting a new proposal")

	timeout := (params.ProposeDuration + self.round*ProposeDeltaDuration) * time.Millisecond

	if self.IsProposer() {
		log.Info("Proposing a new block", "hash", self.proposal.Hash())
		self.propose()
	} else {
		select {
		case self.proposalSub.Chan():
			log.Info("Received a new proposal", "block number", self.proposal.BlockNumber(), "hash", self.proposal.Hash())
		case time.After(timeout):
			log.Info("Timeout expired", "duration", timeout)
		}
	}

	return self.preVoteState
}

func (self *Validator) preVoteState() stateFn {
	log.Info("Starting the pre vote election")
	
	self.prevote()

	return self.preVoteWaitState
}

func (self *Validator) preVoteWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-vote election")

	timeout := (PreVoteDuration + self.round*PreVoteDeltaDuration) * time.Millisecond

	select {
	case self.preVoteMajSub.Chan():
		log.Info("There's a majority")

	case time.After(timeout):
		log.Info("Timeout expired", "duratiom", timeout)
	}

	return self.preCommitState
}

func (self *Validator) preCommitState() stateFn {
	log.Info("Starting the pre commit election")

	self.precommit()

	return self.preCommitWaitState
}

func (self *Validator) preCommitWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-commit election")

	timeout := (PreCommitDuration + self.round + PreCommitDeltaDuration) * time.Millisecond

	select {
	case self.preCommitMajSub.Chan():
		log.Info("There's a majority")

		if self.proposalBlock == nil {
			return self.newRoundState
		} else {
			return self.commitState
		}

	case time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
		return self.newRoundState
	}
}

func (self *Validator) commitState() stateFn {
		log.Info("About to commit the new head block")

		// state updates
		self.state.CommitRound = v.state.Round
		self.state.commitTime = time.Now()

		// Write block using a batch.
		batch := self.chain.chainDb.NewBatch()
		if err := core.WriteBlock(batch, block); err != nil {
			log.Crit("Failed writing block to chain", "err", err)
		}

		// @TODO(rgeraldes)
		/*
		// leaves only when it has all the pre commits
		select {}
		*/

	return self.newBlockState
}