package validator

import (
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/event"
)

// VotingState encapsulates the consensus state for a specific block election
type VotingState struct {
	blockNumber *big.Int
	round       uint64

	validators         types.Voters
	validatorsChecksum [32]byte

	proposal       *types.Proposal
	block          *types.Block
	blockFragments *types.BlockFragments
	votingSystem   *VotingSystem // election votes since round 1

	lockedRound uint64
	lockedBlock *types.Block

	start time.Time // used to sync the validator nodes

	commitRound int

	// inputs
	blockCh  chan *types.Block
	majority *event.TypeMuxSubscription

	// state changes related to the election
	*work
}

// VotingTables represents the voting tables available for each election round
type VotingTables = [2]core.VotingTable

func NewVotingTables(eventMux *event.TypeMux, voters types.Voters) VotingTables {
	majorityFunc := func() {
		go eventMux.Post(core.NewMajorityEvent{})
	}
	tables := VotingTables{}
	tables[0] = core.NewVotingTable(types.PreVote, voters, majorityFunc)
	tables[1] = core.NewVotingTable(types.PreCommit, voters, majorityFunc)
	return tables
}

// VotingSystem records the election votes since round 1
type VotingSystem struct {
	voters         types.Voters
	electionNumber *big.Int // election number
	round          uint64
	votesPerRound  map[uint64]VotingTables
	signer         types.Signer

	eventMux *event.TypeMux
}

// NewVotingSystem returns a new voting system
// @TODO (rgeraldes) - in the future replace eventMux with a subscription method
func NewVotingSystem(eventMux *event.TypeMux, signer types.Signer, electionNumber *big.Int, voters types.Voters) *VotingSystem {
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
