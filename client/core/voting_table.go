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
	votes    *types.VotesSet
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
		votes:    types.NewVotesSet(),
		quorum:   TwoThirdsPlusOneVoteQuorum,
		majority: majority,
	}, nil
}

func (table *votingTable) Add(voteAddressed types.AddressVote) error {
	if !table.isVoter(voteAddressed.Address()) {
		return fmt.Errorf("voter address not found in voting table: 0x%x", voteAddressed.Address().Hash())
	}

	vote := voteAddressed.Vote()
	if table.isDuplicate(vote) {
		log.Error(fmt.Sprintf("a duplicate vote in voting table %v; blockHash %v; voteHash %v. Error: %s",
			table.voteType, vote.BlockHash(), vote.Hash(), vote.String()))
		return ErrDuplicateVote
	}

	table.votes.Add(vote)

	if table.hasQuorum() {
		table.majority()
	}

	return nil
}

func (table *votingTable) isDuplicate(vote *types.Vote) bool {
	return table.votes.Contains(vote.Hash())
}

func (table *votingTable) isVoter(address common.Address) bool {
	return table.voters.Contains(address)
}

func (table *votingTable) hasQuorum() bool {
	return table.quorum(table.votes.Len(), table.voters.Len())
}

type QuorumReachedFunc func()

type QuorumFunc func(votes, voters int) bool

func TwoThirdsPlusOneVoteQuorum(votes, voters int) bool {
	return votes >= voters*2/3+1
}
