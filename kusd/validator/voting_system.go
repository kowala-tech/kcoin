package validator

import (
	"math/big"

	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
)

// VotingSystem records the election votes since round 1
type VotingSystem struct {
	voters         types.ValidatorList
	electionNumber *big.Int // election number
	round          uint64
	votesPerRound  map[uint64]VotingTables
	signer         types.Signer

	eventMux *event.TypeMux
}

// NewVotingSystem returns a new voting system
// @TODO (rgeraldes) - in the future replace eventMux with a subscription method
func NewVotingSystem(eventMux *event.TypeMux, signer types.Signer, electionNumber *big.Int, voters types.ValidatorList) *VotingSystem {
	system := &VotingSystem{
		voters:         voters,
		electionNumber: electionNumber,
		round:          0,
		votesPerRound:  make(map[uint64]VotingTables),
		eventMux:       eventMux,
		signer:         signer,
	}

	system.NewRound()

	return system
}

func (vs *VotingSystem) NewRound() {
	vs.votesPerRound[vs.round] = NewVotingTables(vs.eventMux, vs.voters)
}

// Add registers a vote
func (vs *VotingSystem) Add(vote *types.Vote) error {
	votingTable := vs.getVoteSet(vote.Round(), vote.Type())

	signedVote, err := types.NewSignedVote(vs.signer, vote)
	if err != nil {
		return err
	}

	err = votingTable.Add(signedVote)
	if err != nil {
		return err
	}

	go vs.eventMux.Post(core.NewVoteEvent{Vote: vote})

	return nil
}

func (vs *VotingSystem) getVoteSet(round uint64, voteType types.VoteType) core.VotingTable {
	votingTables, ok := vs.votesPerRound[round]
	if !ok {
		// @TODO (rgeraldes) - critical
		return nil
	}

	return votingTables[int(voteType)]
}
