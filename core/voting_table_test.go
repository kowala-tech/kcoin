package core

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/common"
	"math/big"
	"github.com/stretchr/testify/require"
	"github.com/kowala-tech/kUSD/core/types/mocks"
)

func TestTwoThirdsPlusOneVoteQuorum(t *testing.T) {
	testCases := []struct {
		voters    int
		votes     int
		hasQuorum bool
	}{
		{3, 2, false},
		{3, 2, false},
		{3, 3, true},
		{10, 6, false},
		{10, 7, true},
		{100, 1, false},
		{100, 66, false},
		{100, 67, true},
		{100, 100, true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("voters %d votes %d quorum %t", tc.voters, tc.votes, tc.hasQuorum), func(t *testing.T) {
			assert.Equal(t, tc.hasQuorum, TwoThirdsPlusOneVoteQuorum(tc.votes, tc.voters))
		})
	}
}

func TestVotingTable_Add_CheckIsVoterAndVoteNotSeen_CallsQuorum(t *testing.T) {
	quorum := false
	voterAddress := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")

	validator := types.NewValidator(voterAddress, 0, big.NewInt(1))
	validatorList, err := types.NewValidatorList([]*types.Validator{validator})
	require.NoError(t, err)

	votingTable := NewVotingTable(
		types.PreVote,
		validatorList,
		func() {
			quorum = true
		},
	)

	signedVote := &mocks.SignedVote{}
	signedVote.On("Address").Return(voterAddress)
	signedVote.On("Vote").Return(types.NewVote(big.NewInt(1), common.HexToHash("123"), 0, types.PreCommit))

	err = votingTable.Add(signedVote)

	assert.NoError(t, err)
	assert.Equal(t, validatorList, votingTable.voters)
	assert.Equal(t, 1, len(votingTable.votes))
	assert.True(t, quorum)
}

func TestVotingTable_Add_DoubleVoteFromAddressReturnsError(t *testing.T) {
	voterAddress := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")

	validator := types.NewValidator(voterAddress, 0, big.NewInt(1))
	validatorList, err := types.NewValidatorList([]*types.Validator{validator})
	require.NoError(t, err)

	votingTable := NewVotingTable(
		types.PreVote,
		validatorList,
		func() {},
	)

	signedVote := &mocks.SignedVote{}
	signedVote.On("Address").Return(voterAddress)
	signedVote.On("Vote").Return(types.NewVote(big.NewInt(1), common.HexToHash("123"), 0, types.PreCommit))

	err = votingTable.Add(signedVote)
	assert.NoError(t, err)

	err = votingTable.Add(signedVote)

	assert.EqualError(t, err, "conflict vote seen before")
	assert.Equal(t, validatorList, votingTable.voters)
	assert.Equal(t, 1, len(votingTable.votes))
}

func TestVotingTable_Add_VoteFromNonValidatorReturnsError(t *testing.T) {
	voterAddress := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")
	nonVoterAddress := common.HexToAddress("0x6aaeb6053f3e94c9b9a09f33669435e7ef1beaed")

	validator := types.NewValidator(voterAddress, 0, big.NewInt(1))
	validatorList, err := types.NewValidatorList([]*types.Validator{validator})
	require.NoError(t, err)

	votingTable := NewVotingTable(
		types.PreVote,
		validatorList,
		func() {
			assert.Fail(t, "unexpected Quorum reached call")
		},
	)

	signedVote := &mocks.SignedVote{}
	signedVote.On("Address").Return(nonVoterAddress)
	signedVote.On("Vote").Return(types.NewVote(big.NewInt(1), common.HexToHash("123"), 0, types.PreCommit))

	err = votingTable.Add(signedVote)

	assert.EqualError(t, err, "voter address not found in voting table")
	assert.Equal(t, validatorList, votingTable.voters)
	assert.Equal(t, 0, len(votingTable.votes))
}
