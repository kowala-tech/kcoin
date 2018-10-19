package core

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/common"
	"math/big"
	"github.com/stretchr/testify/require"
	"github.com/kowala-tech/kcoin/client/core/types/mocks"
)

func TestTwoThirdsPlusOneVoteQuorum(t *testing.T) {
	testCases := []struct {
		voters    int
		votes     int
		hasQuorum bool
	}{
		{1, 0, false},
		{1, 1, true},
		{2, 1, false},
		{2, 2, true},
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
			assert.Equal(t, tc.hasQuorum, TwoThirdsPlusOneVoteQuorum(int64(tc.votes), int64(tc.voters)))
		})
	}
}

func TestVotingTable_Add_CheckIsVoterAndVoteNotSeen_CallsQuorum(t *testing.T) {
	quorum := false
	voterAddress := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")

	voter := types.NewVoter(voterAddress, common.Big0, big.NewInt(1))
	voters, err := types.NewVoters([]*types.Voter{voter})
	require.NoError(t, err)

	votingTable := NewVotingTable(
		types.PreVote,
		voters,
		func(winner common.Hash) {
			quorum = true
		},
	)

	signedVote := &mocks.AddressVote{}
	signedVote.On("Address").Return(voterAddress)
	signedVote.On("Vote").Return(types.NewVote(big.NewInt(1), common.HexToHash("123"), 0, types.PreCommit))

	err = votingTable.Add(signedVote)

	assert.NoError(t, err)
	assert.Equal(t, voters, votingTable.voters)
	assert.Equal(t, 1, votingTable.votes.Len())
	assert.True(t, quorum)
}

func TestVotingTable_Add_DoubleVoteFromAddressReturnsError(t *testing.T) {
	voterAddress := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")

	voter := types.NewVoter(voterAddress, common.Big0, big.NewInt(1))
	voters, err := types.NewVoters([]*types.Voter{voter})
	require.NoError(t, err)

	votingTable := NewVotingTable(
		types.PreVote,
		voters,
		func(winner common.Hash) {},
	)

	signedVote := &mocks.AddressVote{}
	signedVote.On("Address").Return(voterAddress)
	signedVote.On("Vote").Return(types.NewVote(big.NewInt(1), common.HexToHash("123"), 0, types.PreCommit))

	err = votingTable.Add(signedVote)
	assert.NoError(t, err)

	err = votingTable.Add(signedVote)

	assert.EqualError(t, err, "duplicate vote")
	assert.Equal(t, voters, votingTable.voters)
	assert.Equal(t, 1, votingTable.votes.Len())
}

func TestVotingTable_Add_VoteFromNonVoterReturnsError(t *testing.T) {
	voterAddress := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")
	nonVoterAddress := common.HexToAddress("0x6aaeb6053f3e94c9b9a09f33669435e7ef1beaed")

	voter := types.NewVoter(voterAddress, common.Big0, big.NewInt(1))
	voters, err := types.NewVoters([]*types.Voter{voter})
	require.NoError(t, err)

	votingTable := NewVotingTable(
		types.PreVote,
		voters,
		func(winner common.Hash) {
			assert.Fail(t, "unexpected Quorum reached call")
		},
	)

	signedVote := &mocks.AddressVote{}
	signedVote.On("Address").Return(nonVoterAddress)
	signedVote.On("Vote").Return(types.NewVote(big.NewInt(1), common.HexToHash("123"), 0, types.PreCommit))

	err = votingTable.Add(signedVote)

	assert.EqualError(t, err, "voter address not found in voting table: 0x0000000000000000000000006aaeb6053f3e94c9b9a09f33669435e7ef1beaed")
	assert.Equal(t, voters, votingTable.voters)
	assert.Equal(t, 0, votingTable.votes.Len())
}
