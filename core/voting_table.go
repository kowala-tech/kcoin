package core

import (
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
	"errors"
	"fmt"
)

var ErrDuplicateVote = errors.New("duplicate vote")

type VotingTable interface {
	Add(vote types.SignedVote) error
}

type votingTable struct {
	voteType types.VoteType
	voters   types.ValidatorList
	votes    types.Votes
	quorum   QuorumFunc
	majority QuorumReachedFunc
}

func NewVotingTable(voteType types.VoteType, voters types.ValidatorList, majority QuorumReachedFunc) *votingTable {
	return &votingTable{
		voteType: voteType,
		voters:   voters,
		votes:    types.Votes{},
		quorum:   TwoThirdsPlusOneVoteQuorum,
		majority: majority,
	}
}

func (table *votingTable) Add(vote types.SignedVote) error {
	if !table.isVoter(vote.Address()) {
		return fmt.Errorf("voter address not found in voting table: %#x", vote.Address().Hash().Str())
	}

	if table.isDuplicate(vote.Vote()) {
		return ErrDuplicateVote
	}

	table.votes = append(table.votes, vote.Vote())

	if table.hasQuorum() {
		table.majority()
	}

	return nil
}

func (table *votingTable) isDuplicate(vote *types.Vote) bool {
	for _, vote := range table.votes {
		if vote.Hash() == vote.Hash() {
			return true
		}
	}
	return false
}

func (table *votingTable) isVoter(address common.Address) bool {
	return table.voters.Contains(address)
}

func (table *votingTable) hasQuorum() bool {
	return table.quorum(len(table.votes), table.voters.Size())
}

type QuorumReachedFunc func()

type QuorumFunc func(votes, voters int) bool

func TwoThirdsPlusOneVoteQuorum(votes, voters int) bool {
	return votes >= voters*2/3+1
}
