package validator

import (
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
	state     *state.StateDB
	header    *types.Header
	tcount    int
	failedTxs types.Transactions
	txs       []*types.Transaction
	receipts  []*types.Receipt
}

// stateFn represents a state function
type stateFn func() stateFn

// @NOTE (rgeraldes) - initial state
func (val *Validator) notLoggedInState() stateFn {
	isGenesis, err := val.network.IsGenesisVoter(&bind.CallOpts{Pending: false}, val.account.Address)
	if err != nil {
		log.Crit("Failed to verify the voter information", "err", err)
		return nil
	}

	// @NOTE (rgeraldes) - sync was already done at this point and by default the investors will be
	// part of the initial set of validators - no need to make a deposit if the block number is 0
	// since these validators will be marked as voters from the start
	if !isGenesis || (isGenesis && val.chain.CurrentBlock().NumberU64() > 0) {
		log.Info("Waiting confirmation to participate in the consensus")
		headSub := val.eventMux.Subscribe(core.ChainHeadEvent{})
		defer headSub.Unsubscribe()

		if err := val.makeDeposit(isGenesis); err != nil {
			return nil
		}

	L:
		for {
			select {
			case _, ok := <-headSub.Chan():
				if !ok {
					return nil
				}

				confirmed, err := val.network.IsVoter(&bind.CallOpts{}, val.account.Address)
				if err != nil {
					log.Crit("Failed to verify the voter registration", "err", err)
				}

				if confirmed {
					break L
				}
			}
		}
	} else {
		log.Info("Deposit is not necessary for a genesis validator (first block)")
	}

	val.restoreLastCommit()

	return val.newElectionState
}

func (val *Validator) newElectionState() stateFn {
	log.Info("Starting a new election")
	// update state machine based on current state
	if err := val.init(); err != nil {
		// @TODO (rgeraldes) - log
		return nil
	}

	<-time.NewTimer(val.start.Sub(time.Now())).C
	// @NOTE (rgeraldes) - wait for txs to be available in the txPool for the round 0
	// If the last block changed the app hash, we may need an empty "proof" block.
	/*
		numTxs, _ := val.backend.TxPool().Stats() //
		if val.round == 0 && numTxs == 0 {        //!cs.needProofBlock(height)
			log.Info("Waiting for transactions")
			txSub := val.eventMux.Subscribe(core.TxPreEvent{})
			defer txSub.Unsubscribe()
			<-txSub.Chan()
		}
	*/
	time.Sleep(time.Duration(30) * time.Second)

	return val.newRoundState
}

func (val *Validator) newRoundState() stateFn {
	log.Info("Starting a new voting round", "start time", val.start, "block number", val.blockNumber, "round", val.round)

	if val.round != 0 {
		val.round++
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

func (val *Validator) preVoteState() stateFn {
	log.Info("Pre vote sub-election")
	val.preVote()

	return val.preVoteWaitState
}

func (val *Validator) preVoteWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-vote sub-election")
	timeout := time.Duration(params.PreVoteDuration+uint64(val.round)*params.PreVoteDeltaDuration) * time.Millisecond

	select {
	case <-val.majority.Chan():
		log.Info("There's a majority in the pre-vote sub-election!")
	case <-time.After(timeout):
		log.Info("Timeout expired", "duration", timeout)
	}

	return val.preCommitState
}

func (val *Validator) preCommitState() stateFn {
	log.Info("Pre commit sub-election")
	val.preCommit()

	return val.preCommitWaitState
}

func (val *Validator) preCommitWaitState() stateFn {
	log.Info("Waiting for a majority in the pre-commit sub-election")
	timeout := time.Duration(params.PreCommitDuration+uint64(val.round)+params.PreCommitDeltaDuration) * time.Millisecond
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

	voter, err := val.network.IsVoter(&bind.CallOpts{}, val.account.Address)
	if err != nil {
		// @TODO (rgeraldes) - complete
		//log.Error()
	}
	if !voter {
		return val.loggedOutState
	}

	return val.newElectionState
}

// @NOTE (rgeraldes) - end state
func (val *Validator) loggedOutState() stateFn {
	return nil
}
