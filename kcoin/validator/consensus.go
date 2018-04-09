package validator

import (
	"math/big"
	"time"

	"errors"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/event"
)

// VotingState encapsulates the consensus state for a specific block election
type VotingState struct {
	blockNumber *big.Int
	round       uint64

	voters         types.Voters
	votersChecksum [32]byte

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
	tables[0], _ = core.NewVotingTable(types.PreVote, voters, majorityFunc)
	tables[1], _ = core.NewVotingTable(types.PreCommit, voters, majorityFunc)
	return tables
}

// VotingSystem records the election votes since round 1
type VotingSystem struct {
	voters         types.Voters
	electionNumber *big.Int // election number
	round          uint64
	votesPerRound  map[uint64]VotingTables

	eventMux *event.TypeMux
}

// NewVotingSystem returns a new voting system
func NewVotingSystem(eventMux *event.TypeMux, electionNumber *big.Int, voters types.Voters) *VotingSystem {
	system := &VotingSystem{
		voters:         voters,
		electionNumber: electionNumber,
		round:          0,
		votesPerRound:  make(map[uint64]VotingTables),
		eventMux:       eventMux,
	}

	system.NewRound()

	return system
}

func (vs *VotingSystem) NewRound() {
	vs.votesPerRound[vs.round] = NewVotingTables(vs.eventMux, vs.voters)
}

// Add registers a vote
func (vs *VotingSystem) Add(vote types.AddressVote) error {
	votingTable, err := vs.getVoteSet(vote.Vote().Round(), vote.Vote().Type())
	if err != nil {
		return err
	}

	err = votingTable.Add(vote)
	if err != nil {
		return err
	}

	go vs.eventMux.Post(core.NewVoteEvent{Vote: vote.Vote()})

	return nil
}

func (vs *VotingSystem) getVoteSet(round uint64, voteType types.VoteType) (core.VotingTable, error) {
	votingTables, ok := vs.votesPerRound[round]
	if !ok {
		return nil, errors.New("voting table for round doesnt exists")
	}

	if uint64(voteType) > uint64(len(votingTables)-1) {
		return nil, errors.New("invalid voteType on add vote ")
	}

	return votingTables[int(voteType)], nil
}
