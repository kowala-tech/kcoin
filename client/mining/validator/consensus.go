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
	blockCh                                  chan *types.Block
	preVoteMajorityCh, preCommitMajorityCh   chan core.NewMajorityEvent
	preVoteMajoritySub, preCommitMajoritySub event.Subscription

	// state changes related to the election
	*work
}

func (state *VotingState) VotingSystem() *VotingSystem { return state.votingSystem }

// VotingTables represents the voting tables available for each election round
type VotingTables = [2]core.VotingTable

func NewVotingTables(voters types.Voters) (VotingTables, error) {
	var err error
	tables := VotingTables{}

	// prevote
	tables[types.PreVote], err = core.NewVotingTable(types.PreVote, voters)
	if err != nil {
		return tables, err
	}

	// precommit
	tables[types.PreCommit], err = core.NewVotingTable(types.PreCommit, voters)
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
	voteFeed       event.Feed
	scope          event.SubscriptionScope
}

// NewVotingSystem returns a new voting system
func NewVotingSystem(electionNumber *big.Int, voters types.Voters) (*VotingSystem, error) {
	system := &VotingSystem{
		voters:         voters,
		electionNumber: electionNumber,
		round:          0,
		votesPerRound:  make(map[uint64]VotingTables),
	}

	err := system.NewRound()
	if err != nil {
		return nil, err
	}

	return system, nil
}

func (vs *VotingSystem) NewRound() error {
	var err error
	vs.votesPerRound[vs.round], err = NewVotingTables(vs.voters)
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

	go vs.voteFeed.Send(core.NewVoteEvent{Vote: vote.Vote()})

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

// SubscribeNewVoteEvent registers a subscription of NewVoteEvent and
// starts sending event to the given channel.
func (vs *VotingSystem) SubscribeNewVoteEvent(ch chan<- core.NewVoteEvent) event.Subscription {
	return vs.scope.Track(vs.voteFeed.Subscribe(ch))
}

// SubscribePreVoteMajority registers a subscription of NewMajorityEvent
// in the pre vote table and starts sending event to the given channel.
func (vs *VotingSystem) SubscribePreVoteMajority(ch chan<- core.NewMajorityEvent) event.Subscription {
	votingTables := vs.votesPerRound[vs.round]
	return votingTables[int(types.PreVote)].SubscribeNewMajorityEvent(ch)
}

// SubscribePreCommitMajority registers a subscription of NewMajorityEvent
// in the pre commit table and starts sending event to the given channel.
func (vs *VotingSystem) SubscribePreCommitMajority(ch chan<- core.NewMajorityEvent) event.Subscription {
	votingTables := vs.votesPerRound[vs.round]
	return votingTables[int(types.PreCommit)].SubscribeNewMajorityEvent(ch)
}
