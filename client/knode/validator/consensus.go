package validator

import (
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/log"
)

// VotingState encapsulates the consensus state for a specific block election
type VotingState struct {
	blockNumber *big.Int
	round       uint64

	voters         types.Voters
	votersChecksum [32]byte

	proposer              *types.Voter
	proposal              *types.Proposal
	isProposal            bool
	block                 *types.Block
	blockFragmentsLock    sync.RWMutex
	blockFragments        map[common.Hash]*types.BlockFragments
	blockFragmentsStorage map[common.Hash][]*types.BlockFragment
	votingSystem          *VotingSystem // election votes since round 1

	lockedRound uint64
	lockedBlock *types.Block

	parentBlockCreatedAt   time.Time // used to sync the validator nodes
	previousRoundCreatedAt time.Time // used to sync the validator nodes
	roundCreatedAt         time.Time // used to sync the validator nodes

	commitRound int

	// inputs
	blockCh  chan *types.Block
	majority *event.TypeMuxSubscription

	// state changes related to the election
	*work
}

const votingTypes = 2

// VotingTables represents the voting tables available for each election round
type VotingTables = [votingTypes]core.VotingTable

func NewVotingTables(voters types.Voters, poster eventPoster) (*VotingTables, error) {
	majorityFunc := contextMajorityFunc(poster)

	if voters == nil {
		return nil, errors.New("cant create a voting table with nil voters")
	}

	return &VotingTables{
		types.PreVote:   core.NewVotingTable(types.PreVote, voters, majorityFunc("prevote")),
		types.PreCommit: core.NewVotingTable(types.PreCommit, voters, majorityFunc("precommit")),
	}, nil
}

func contextMajorityFunc(poster eventPoster) func(msg string) func(common.Hash) {
	return func(msg string) func(common.Hash) {
		return func(winnerBlock common.Hash) {
			log.Debug("voting majority established", "log", msg)
			poster.EventPost(core.NewMajorityEvent{Winner: winnerBlock})
		}
	}
}

// VotingSystem records the election votes since round 1
type VotingSystem struct {
	voters        types.Voters
	round         uint64
	votesPerRound map[uint64]*VotingTables
	poster        eventPoster
}

// NewVotingSystem returns a new voting system
func NewVotingSystem(voters types.Voters, poster eventPoster) *VotingSystem {
	return &VotingSystem{
		voters:        voters,
		votesPerRound: make(map[uint64]*VotingTables),
		poster:        poster,
	}
}

// Add registers a vote
func (vs *VotingSystem) Add(vote types.AddressVote) error {
	votingTable, err := vs.getVotingTable(vote.Vote().Round(), vote.Vote().Type())
	if err != nil {
		return err
	}

	err = votingTable.Add(vote)
	if err != nil {
		return err
	}

	vs.poster.EventPost(core.NewVoteEvent{Vote: vote.Vote()})

	return nil
}

func (vs *VotingSystem) Leader(round uint64, voteType types.VoteType) (common.Hash, error) {
	votingTable, err := vs.getVotingTable(round, voteType)
	if err != nil {
		return common.Hash{}, err
	}

	return votingTable.Leader(), nil
}

func (vs *VotingSystem) getVotingTable(round uint64, voteType types.VoteType) (core.VotingTable, error) {
	if uint64(voteType) > votingTypes-1 {
		return nil, errors.New("invalid voteType on add vote")
	}

	if _, ok := vs.votesPerRound[round]; !ok {
		votingTables, err := NewVotingTables(vs.voters, vs.poster)
		if err != nil {
			return nil, err
		}
		vs.votesPerRound[vs.round] = votingTables
	}

	return vs.votesPerRound[vs.round][voteType], nil
}
