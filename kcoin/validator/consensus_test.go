package validator

import (
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/core/types/mocks"
	"github.com/kowala-tech/kcoin/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestNewVotingTables_Return2Tables(t *testing.T) {
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, 0, big.NewInt(1))})
	require.NoError(t, err)
	votingTables := NewVotingTables(nil, voters)

	assert.NotNil(t, votingTables[0])
	assert.NotNil(t, votingTables[1])
}

func TestNewVotingSystem_CreatesNewRound(t *testing.T) {
	votingSystem := NewVotingSystem(nil, big.NewInt(1), nil)

	assert.NotNil(t, votingSystem.votesPerRound[0])
}

func TestVotingSystem_AddVoteWrongRoundReturnsError(t *testing.T) {
	vote := types.NewVote(big.NewInt(1), common.Hash{}, 1, types.PreCommit)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	votingSystem := NewVotingSystem(nil, big.NewInt(1), nil)

	err := votingSystem.Add(addressVote)

	assert.Error(t, err)
}

func TestVotingSystem_AddVoteWrongVoteTypeReturnsError(t *testing.T) {
	vote := types.NewVote(big.NewInt(1), common.Hash{}, 0, types.PreCommit+1)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	votingSystem := NewVotingSystem(nil, big.NewInt(1), nil)

	err := votingSystem.Add(addressVote)

	assert.Error(t, err)
}

func TestVotingSystem_AddVoteAddsVoteToTable(t *testing.T) {
	vote := types.NewVote(big.NewInt(1), common.Hash{}, 0, types.PreVote)
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, 0, big.NewInt(1))})
	require.NoError(t, err)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	addressVote.On("Address").Return(address)
	votingSystem := NewVotingSystem(&event.TypeMux{}, big.NewInt(1), voters)

	err = votingSystem.Add(addressVote)
	assert.NoError(t, err)

	addressVote.AssertExpectations(t)
}
