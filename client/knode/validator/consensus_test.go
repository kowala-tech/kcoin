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

	eventPoster := newPoster(&event.TypeMux{})
	votingTables, err := NewVotingTables(voters, eventPoster)
	require.NoError(t, err)

	assert.NotNil(t, votingTables[0])
	assert.NotNil(t, votingTables[1])
}

func TestVotingSystem_AddVoteNewRound(t *testing.T) {
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, common.Big0, big.NewInt(1))})
	require.NoError(t, err)

	vote := types.NewVote(big.NewInt(1), common.Hash{}, 1, types.PreCommit)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	addressVote.On("Address").Return(address)

	eventPoster := newPoster(&event.TypeMux{})
	votingSystem := NewVotingSystem(voters, eventPoster)

	err = votingSystem.Add(addressVote)

	assert.NoError(t, err)
}

func TestVotingSystem_AddVoteWrongVoteTypeReturnsError(t *testing.T) {
	address := common.HexToAddress("0x1000000000000000000000000000000000000000")
	voters, err := types.NewVoters([]*types.Voter{types.NewVoter(address, common.Big0, big.NewInt(1))})
	require.NoError(t, err)
	vote := types.NewVote(big.NewInt(1), common.Hash{}, 0, types.PreCommit+1)
	addressVote := &mocks.AddressVote{}
	addressVote.On("Vote").Return(vote)
	eventPoster := newPoster(&event.TypeMux{})
	votingSystem := NewVotingSystem(voters, eventPoster)

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
	eventPoster := newPoster(&event.TypeMux{})
	votingSystem := NewVotingSystem(voters, eventPoster)

	err = votingSystem.Add(addressVote)
	assert.NoError(t, err)

	addressVote.AssertExpectations(t)
}
