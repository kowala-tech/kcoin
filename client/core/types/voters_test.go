package types

import (
	"fmt"
	"math/big"
	"testing"

	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var voterSet = [4]*Voter{
	makeVoter("0x1000000000000000000000000000000000000000", 100, 100),
	makeVoter("0x2000000000000000000000000000000000000000", 101, 101),
	makeVoter("0x3000000000000000000000000000000000000000", 99, 99),
	makeVoter("0x4000000000000000000000000000000000000000", 99, 99),
}

func TestVoter_Properties(t *testing.T) {
	address := common.Address{}
	deposit := new(big.Int).SetUint64(100)
	weight := &big.Int{}
	voter := NewVoter(address, deposit, weight)

	assert.Equal(t, address, voter.Address())
	assert.Equal(t, deposit, voter.Deposit())
	assert.Equal(t, weight, voter.Weight())
}

func TestVoters_EmptyReturnsError(t *testing.T) {
	voters, err := NewVoters(nil)

	require.Error(t, err)
	require.Nil(t, voters)
}

func TestVoters_GetAtNegativeIndexReturnsNil(t *testing.T) {
	voters, err := NewVoters([]*Voter{voterSet[0]})
	require.NoError(t, err)

	voter := voters.At(-1)

	assert.Nil(t, voter)
}

func TestVoters_GetAtOverLastReturnsNil(t *testing.T) {
	voters, err := NewVoters([]*Voter{voterSet[0]})
	require.NoError(t, err)

	voterAt := voters.At(0)
	assert.Equal(t, voterSet[0], voterAt)

	voterAt = voters.At(1)
	assert.Nil(t, voterAt)
}

func TestVoters_One(t *testing.T) {
	voters, err := NewVoters([]*Voter{voterSet[0]})

	require.NoError(t, err)
	assert.Equal(t, 1, voters.Len())
	assert.Equal(t, voterSet[0], voters.At(0))
	assert.Equal(t, voterSet[0], voters.Get(voterSet[0].Address()))
	assert.Equal(t, true, voters.Contains(voterSet[0].Address()))
	assert.Equal(t, voterSet[0], voters.NextProposer())
}

func TestVoters_UpdateWeightChangesProposer(t *testing.T) {
	voters, err := NewVoters([]*Voter{voterSet[0], voterSet[1], voterSet[2]})

	require.NoError(t, err)
	proposer := voters.NextProposer()
	assert.Equal(t, voterSet[1], proposer)
	assert.Equal(t, big.NewInt(101), proposer.weight)
	assert.Equal(t, big.NewInt(200), voters.At(0).weight)
	assert.Equal(t, big.NewInt(101), voters.At(1).weight)
	assert.Equal(t, big.NewInt(198), voters.At(2).weight)
	assert.Equal(t, 3, voters.Len())
}

func TestVoters_UpdateWeightChangesProposerElections(t *testing.T) {
	voters, err := NewVoters([]*Voter{voterSet[0], voterSet[1], voterSet[2]})
	require.NoError(t, err)
	require.Equal(t, 3, voters.Len())

	elections := []struct {
		proposerWeight *big.Int
		voter1weight   *big.Int
		voter2weight   *big.Int
		voter3weight   *big.Int
	}{
		{big.NewInt(200), big.NewInt(200), big.NewInt(202), big.NewInt(297)},
		{big.NewInt(297), big.NewInt(300), big.NewInt(303), big.NewInt(297)},
		{big.NewInt(303), big.NewInt(400), big.NewInt(303), big.NewInt(396)},
	}

	for round, tc := range elections {
		t.Run(fmt.Sprintf("round %d", round), func(t *testing.T) {
			proposer := voters.NextProposer()
			assert.Equal(t, tc.proposerWeight, proposer.weight)
			assert.Equal(t, tc.voter1weight, voters.At(0).weight)
			assert.Equal(t, tc.voter2weight, voters.At(1).weight)
			assert.Equal(t, tc.voter3weight, voters.At(2).weight)
		})
	}
}

func TestVoters_IsHashable(t *testing.T) {
	voters1, err := NewVoters([]*Voter{voterSet[0], voterSet[1], voterSet[2]})
	require.NoError(t, err)

	voters2, err := NewVoters([]*Voter{voterSet[1], voterSet[2], voterSet[3]})
	require.NoError(t, err)

	assert.NotEqual(t, voters1.Hash(), voters2.Hash())
}

func TestNewDeposit(t *testing.T) {
	amount := new(big.Int).SetUint64(100)
	now := time.Now().Unix()
	deposit := NewDeposit(amount, now)

	assert.Equal(t, amount, deposit.Amount())
	assert.Equal(t, now, deposit.AvailableAtTimeUnix())
}

func makeVoter(hexAddress string, deposit uint64, weight uint64) *Voter {
	address := common.HexToAddress(hexAddress)
	return NewVoter(address, new(big.Int).SetUint64(deposit), new(big.Int).SetUint64(weight))
}
