package validator

import (
	"errors"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
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

func NewVotingTables(eventMux *event.TypeMux, voters types.Voters) (VotingTables, error) {
	majorityFunc := func(winnerBlock common.Hash) {
		go eventMux.Post(core.NewMajorityEvent{Winner: winnerBlock})
	}

	var err error
	tables := VotingTables{}

	// prevote
	tables[types.PreVote], err = core.NewVotingTable(types.PreVote, voters, majorityFunc)
	if err != nil {
		return tables, err
	}

	// precommit
	tables[types.PreCommit], err = core.NewVotingTable(types.PreCommit, voters, majorityFunc)
	if err != nil {
		return tables, err
	}

	return tables, nil
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
func NewVotingSystem(eventMux *event.TypeMux, electionNumber *big.Int, voters types.Voters) (*VotingSystem, error) {
	system := &VotingSystem{
		voters:         voters,
		electionNumber: electionNumber,
		round:          0,
		votesPerRound:  make(map[uint64]VotingTables),
		eventMux:       eventMux,
	}

	err := system.NewRound()
	if err != nil {
		return nil, err
	}

	return system, nil
}

func (vs *VotingSystem) NewRound() error {
	var err error
	vs.votesPerRound[vs.round], err = NewVotingTables(vs.eventMux, vs.voters)
	if err != nil {
		return err
	}
	return nil
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

func (vs *VotingSystem) Leader(round uint64, voteType types.VoteType) (common.Hash, error) {
	votingTable, err := vs.getVoteSet(round, voteType)
	if err != nil {
		return common.Hash{}, err
	}

	return votingTable.Leader(), nil
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
