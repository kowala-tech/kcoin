package validator

import (
	"math/big"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/state"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/log"
	"github.com/kowala-tech/kcoin/params"
	"github.com/kowala-tech/kcoin/kcoin/wal/wal"
	"fmt"
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
	//fixme: возможно восстановление надо делать тут!!!
	fmt.Println(9999999999999999)
	isGenesis, err := val.election.IsGenesisValidator(val.walletAccount.Account().Address)
	if err != nil {
		fmt.Println(111, err, val.walletAccount.Account().Address.String())
		log.Warn("Failed to verify the voter information", "err", err)
		return nil
	}

	// @NOTE (rgeraldes) - sync was already done at this point and by default the investors will be
	// part of the initial set of validators - no need to make a deposit if the block number is 0
	// since these validators will be marked as voters from the start
	fmt.Println("===================== notLoggedInState", isGenesis, val.chain.CurrentBlock().NumberU64())
	if !isGenesis || (isGenesis && val.chain.CurrentBlock().NumberU64() > 0) {
		fmt.Println("===================== notLoggedInState ==== 1", val.chain.CurrentBlock().NumberU64())
		fmt.Println(val.election.IsValidator(val.walletAccount.Account().Address))
		fmt.Println("===")
		fmt.Println("******************************** Alloc ACCOUNT", val.walletAccount.Account().Address.String())
		chainHeadCh := make(chan core.ChainHeadEvent)
		chainHeadSub := val.chain.SubscribeChainHeadEvent(chainHeadCh)
		defer chainHeadSub.Unsubscribe()

		voters, err := val.election.Validators()
		fmt.Println("**", voters.Len(), err)
		if err := val.election.Join(val.walletAccount, val.deposit); err != nil {
			fmt.Println(2222, err)
			log.Error("Error joining validators network", "err", err)
			return nil
		}
		fmt.Println("DONE election.Join")

		log.Info("Waiting confirmation to participate in the consensus")
	L:
		for {
			fmt.Println("//////////////")

			select {
			case _, ok := <-chainHeadCh:
				if !ok {
					return nil
				}

				confirmed, err := val.election.IsValidator(val.walletAccount.Account().Address)
				if err != nil {
					fmt.Println(333)
					log.Crit("Failed to verify the voter registration", "err", err)
				}

				if confirmed {
					break L
				}
			}
		}
	} else {
		fmt.Println("===================== notLoggedInState ==== 2")
		isVoter, err := val.election.IsValidator(val.walletAccount.Account().Address)
		if err != nil {
			fmt.Println(444)
			log.Crit("Failed to verify the voter information", "err", err)
			return nil
		}
		if !isVoter {
			fmt.Println(555)
			log.Crit("Invalid genesis - genesis validator needs to be registered as a voter", "address", val.walletAccount.Account().Address)
		}

		log.Info("Deposit is not necessary for a genesis validator (first block)")
	}

	log.Info("Starting validation operation")
	atomic.StoreInt32(&val.validating, 1)

	log.Info("Voter has been accepted in the election")
	//todo: обратить внимание!!! возможно надо тут восстанавливать состояние!!!
	val.restoreLastCommit()
	fmt.Println(6666)
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
			txCh := make(chan core.TxPreEvent)
			txSub := val.backend.TxPool().SubscribeTxPreEvent(txCh)
			defer txSub.Unsubscribe()
			<-txCh
		}
	}

	val.wal.Save(wal.BlockStart(val.blockNumber))

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
		log.Info("Waiting for the proposal", proposer.Address())
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
		return val.commitState
	}
}

func (val *validator) commitState() stateFn {
	log.Info("Commit state")

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

	voter, err := val.election.IsValidator(val.walletAccount.Account().Address)
	if err != nil {
		log.Crit("Failed to verify if the validator is a voter", "err", err)
	}
	if !voter {
		return val.loggedOutState
	}

	val.wal.Save(wal.BlockCommit(block.Number()))

	return val.newElectionState
}

func (val *validator) loggedOutState() stateFn {
	log.Info("Logged out")
	fmt.Println("00000000000000000000000000000000000000000000000000000000000000000")

	atomic.StoreInt32(&val.validating, 0)

	return nil
}
