package validator

import (
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/core/types/mocks"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVotingTables_Return2Tables(t *testing.T) {
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, common.Big0, big.NewInt(1))})
	require.NoError(t, err)
	votingTables, err := NewVotingTables(nil, voters)
	require.NoError(t, err)

	assert.NotNil(t, votingTables[0])
	assert.NotNil(t, votingTables[1])
}

func TestNewVotingSystem_CreatesNewRound(t *testing.T) {
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, common.Big0, big.NewInt(1))})
	require.NoError(t, err)
	votingSystem, err := NewVotingSystem(nil, big.NewInt(1), voters)
	require.NoError(t, err)

	assert.NotNil(t, votingSystem.votesPerRound[0])
}

func TestVotingSystem_AddVoteWrongRoundReturnsError(t *testing.T) {
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, common.Big0, big.NewInt(1))})
	require.NoError(t, err)
	vote := types.NewVote(big.NewInt(1), common.Hash{}, 1, types.PreCommit)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	votingSystem, err := NewVotingSystem(nil, big.NewInt(1), voters)
	require.NoError(t, err)

	err = votingSystem.Add(addressVote)

	assert.Error(t, err)
}

func TestVotingSystem_AddVoteWrongVoteTypeReturnsError(t *testing.T) {
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, common.Big0, big.NewInt(1))})
	require.NoError(t, err)
	vote := types.NewVote(big.NewInt(1), common.Hash{}, 0, types.PreCommit+1)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	votingSystem, err := NewVotingSystem(nil, big.NewInt(1), voters)
	require.NoError(t, err)

	err = votingSystem.Add(addressVote)

	assert.Error(t, err)
}

func TestVotingSystem_AddVoteAddsVoteToTable(t *testing.T) {
	vote := types.NewVote(big.NewInt(1), common.Hash{}, 0, types.PreVote)
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, common.Big0, big.NewInt(1))})
	require.NoError(t, err)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	addressVote.On("Address").Return(address)
	votingSystem, err := NewVotingSystem(&event.TypeMux{}, big.NewInt(1), voters)
	assert.NoError(t, err)

	err = votingSystem.Add(addressVote)
	assert.NoError(t, err)

	addressVote.AssertExpectations(t)
}
