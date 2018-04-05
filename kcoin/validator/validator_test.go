package validator

import (
	"testing"

	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/kcoin/validator/mocks"
	"github.com/stretchr/testify/require"
)

func TestAddVote_NotValidating(t *testing.T) {
	vote := &types.Vote{}
	validator := &validator{
		validating: 0,
	}
	require.Error(t, ErrCantVoteNotValidating, validator.AddVote(vote))
}

func TestAddVote_Validating(t *testing.T) {
	vote := &types.Vote{}
	votingSystem := &mocks.VotingSystem{}
	votingSystem.On("Add", &types.Vote{}).Return(nil)

	validator := &validator{
		validating: 1,
		VotingState: VotingState{
			votingSystem: votingSystem,
		},
	}

	require.NoError(t, validator.AddVote(vote))
}

func TestAddVote_Validating_DuplicateVote(t *testing.T) {}
