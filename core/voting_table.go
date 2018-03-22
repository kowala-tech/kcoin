package core

import (
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
	"errors"
)

/*
	VoteSet helps collect signatures from validators at each height + round for a
	predefined vote type.

	We need VoteSet to be able to keep track of conflicting votes when validators
	double-sign.  Yet, we can't keep track of *all* the votes seen, as that could
	be a DoS attack vector.

	There are two storage areas for votes.
	1. voteSet.votes
	2. voteSet.votesByBlock

	`.votes` is the "canonical" list of votes.  It always has at least one vote,
	if a vote from a validator had been seen at all.  Usually it keeps track of
	the first vote seen, but when a 2/3 majority is found, votes for that get
	priority and are copied over from `.votesByBlock`.

	`.votesByBlock` keeps track of a list of votes for a particular block.  There
	are two ways a &blockVotes{} gets created in `.votesByBlock`.
	1. the first vote seen by a validator was for the particular block.
	2. a peer claims to have seen 2/3 majority for the particular block.

	Since the first vote from a validator will always get added in `.votesByBlock`
	, all votes in `.votes` will have a corresponding entry in `.votesByBlock`.

	When a &blockVotes{} in `.votesByBlock` reaches a 2/3 majority quorum, its
	votes are copied into `.votes`.

	All this is memory bounded because conflicting votes only get added if a peer
	told us to track that block, each peer only gets to tell us 1 such block, and,
	there's only a limited number of peers.

	NOTE: Assumes that the sum total of voting power does not exceed MaxUInt64.
*/

type VotingTable struct {
	voteType types.VoteType
	voters   types.ValidatorList
	votes    []*types.Vote // Primary votes to share
	signer   types.Signer
	quorum   QuorumFunc
	eventMux *event.TypeMux
}

func NewVotingTable(eventMux *event.TypeMux, signer types.Signer, voteType types.VoteType, voters types.ValidatorList) *VotingTable {
	return &VotingTable{
		voteType: voteType,
		voters:   voters,
		votes:    make([]*types.Vote, voters.Size()),
		eventMux: eventMux,
		signer:   signer,
		quorum:   TwoThirdsPlusOneVoteQuorum,
	}
}

func (table *VotingTable) Add(vote *types.Vote, local bool) error {
	// check signature
	from, err := types.VoteSender(table.signer, vote) // already validated & cached
	if err != nil {
		return err
	}

	if !table.isVoterParticipating(from) {
		return errors.New("voter address not found in voting table")
	}

	if table.hasVoted(vote.Hash()) {
		return errors.New("conflict code in voting table add")
	}

	table.votes = append(table.votes, vote)

	go table.eventMux.Post(NewVoteEvent{Vote: vote})

	if table.haveQuorum() {
		go table.eventMux.Post(NewMajorityEvent{})
	}

	return nil
}

func (table *VotingTable) hasVoted(voteHash common.Hash) bool {
	for _, vote := range table.votes {
		if vote.Hash() == voteHash {
			return true
		}
	}
	return false
}

func (table *VotingTable) isVoterParticipating(address common.Address) bool {
	return table.voters.Contains(address)
}

func (table *VotingTable) haveQuorum() bool {
	return table.quorum(len(table.votes), table.voters.Size())
}

type QuorumFunc func(votes, voters int) bool

func TwoThirdsPlusOneVoteQuorum(votes, voters int) bool {
	return votes > voters*2/3+1
}
