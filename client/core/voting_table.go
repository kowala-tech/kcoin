package core

import (
	"errors"
	"fmt"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
)

var ErrDuplicateVote = errors.New("duplicate vote")

type VotingTable interface {
	Add(vote types.AddressVote) error
}

type votingTable struct {
	voteType types.VoteType
	voters   types.Voters
	votes    types.Votes
	quorum   QuorumFunc
	majority QuorumReachedFunc
}

func NewVotingTable(voteType types.VoteType, voters types.Voters, majority QuorumReachedFunc) (*votingTable, error) {
	if voters == nil {
		return nil, errors.New("cant create a voting table with nil voters")
	}

	return &votingTable{
		voteType: voteType,
		voters:   voters,
		votes:    types.Votes{},
		quorum:   TwoThirdsPlusOneVoteQuorum,
		majority: majority,
	}, nil
}

func (table *votingTable) Add(vote types.AddressVote) error {
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
	voteHash := vote.Hash()
	for _, tableVote := range table.votes {
		if tableVote.Hash() == voteHash {
			log.Error(fmt.Sprintf("a duplicate vote error: %s", vote.String()))
			return true
		}
	}
	return false
}

func (table *votingTable) isVoter(address common.Address) bool {
	return table.voters.Contains(address)
}

func (table *votingTable) hasQuorum() bool {
	return table.quorum(len(table.votes), table.voters.Len())
}

type QuorumReachedFunc func()

type QuorumFunc func(votes, voters int) bool

func TwoThirdsPlusOneVoteQuorum(votes, voters int) bool {
	return votes >= voters*2/3+1
}
