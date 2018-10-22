package core

import (
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
)

type VotingTable interface {
	Add(vote types.AddressVote) error
	Leader() common.Hash
}

type votingTable struct {
	voteType types.VoteType
	voters   types.Voters
	votes    *types.VotesSet
	quorum   QuorumFunc
	majority QuorumReachedFunc
}

func NewVotingTable(voteType types.VoteType, voters types.Voters, majority QuorumReachedFunc) *votingTable {
	return &votingTable{
		voteType: voteType,
		voters:   voters,
		votes:    types.NewVotesSet(),
		quorum:   TwoThirdsPlusOneVoteQuorum,
		majority: majority,
	}
}

func (table *votingTable) Add(voteAddressed types.AddressVote) error {
	if !table.isVoter(voteAddressed.Address()) {
		return fmt.Errorf("voter address not found in voting table: 0x%x", voteAddressed.Address().Hash())
	}

	if err := table.isDuplicate(voteAddressed); err != nil {
		return err
	}

	table.votes.Add(voteAddressed)

	if table.hasQuorum() {
		table.majority(table.Leader())
	}

	return nil
}

func (table *votingTable) Leader() common.Hash {
	return table.votes.Leader()
}

func (table *votingTable) isDuplicate(voteAddressed types.AddressVote) error {
	err := table.votes.Contains(voteAddressed)
	if err != nil {
		vote := voteAddressed.Vote()
		log.Debug("a duplicate vote",
			"table", table.voteType,
			"blockHash", vote.BlockHash().String(),
			"voteHash", vote.Hash().String(),
			"validator", voteAddressed.Address().Hash().String(),
			"err", vote.String())
		return err
	}
	return nil
}

func (table *votingTable) isVoter(address common.Address) bool {
	return table.voters.Contains(address)
}

func (table *votingTable) hasQuorum() bool {
	leaderBlockVotes := table.votes.Count(table.Leader())
	isQuorum := table.quorum(int64(leaderBlockVotes), int64(table.voters.Len()))

	log.Debug("voting. hasQuorum", "leaderVotes", leaderBlockVotes, "votes", table.votes.Len(),
		"voters", table.voters.Len(), "isQuorum", isQuorum, "leader", table.Leader().String())

	return isQuorum
}

type QuorumReachedFunc func(winner common.Hash)

type QuorumFunc func(votes, voters int64) bool

func TwoThirdsPlusOneVoteQuorum(votes, voters int64) bool {
	if votes > voters {
		log.Error("the number of votes is greater than the number of voters", "votes", votes, "voters", voters)
	}

	majority := big.NewRat(2*voters,3)
	majority.Add(majority, big.NewRat(1,1))

	if majority.Cmp(big.NewRat(voters, 1)) > 0 {
		// the majority shouldn't be greater than number of voters
		majority.SetInt64(voters)
	}

	// votes >= voters*2/3+1
	return majority.Cmp(big.NewRat(votes, 1)) <= 0
}
