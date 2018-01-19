package validator

import (
	"math/big"
	"time"

	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
)

// Election encapsulates the consensus state for a specific block election
type Election struct {
	blockNumber *big.Int
	round       uint64

	validators     *types.ValidatorSet
	proposal       *types.Proposal
	block          *types.Block
	blockFragments *types.BlockFragments
	votingSystem   *VotingSystem // election votes since round 1

	lockedRound uint64
	lockedBlock *types.Block

	start time.Time // used to sync the validator nodes

	commitRound int

	lastCommit     *core.VotingTable // Last precommits at current block number-1
	lastValidators *types.ValidatorSet

	// inputs
	proposalCh                    chan *types.Proposal
	firstMajority, secondMajority *event.TypeMuxSubscription

	// @TODO (rgeraldes) - not sure if it will be necessary
	// proposer
	*work
}

// VotingTables represents the voting tables available for each election round
type VotingTables = [2]*core.VotingTable

// VotingSystem records the election votes since round 1
type VotingSystem struct {
	round         uint64
	votesPerRound map[uint64]VotingTables
}

// NewVotingSystem returns a new voting system
func NewVotingSystem() *VotingSystem {
	return &VotingSystem{
		round:         0,
		votesPerRound: make(map[uint64]VotingTables),
	}
}

// Add registers a vote
func (vs *VotingSystem) Add(vote *types.Vote) (bool, error) {
	// @TODO (rgeraldes) - validation
	votingTable := vs.getVoteSet(vote.Round(), vote.Type())
	votingTable.Add(vote)

	return false, nil
}

func (vs *VotingSystem) getVoteSet(round uint64, voteType types.VoteType) *core.VotingTable {
	votingTables, ok := vs.votesPerRound[round]
	if !ok {
		// @TODO (rgeraldes) - critical
		return nil
	}

	return votingTables[int(voteType)]
}
