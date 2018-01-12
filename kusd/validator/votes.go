package validator

import (
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
)

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
		round:         1,
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
