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

var voterSet = []*Voter{
	makeVoter("0x1000000000000000000000000000000000000000", 100, 100),
	makeVoter("0x2000000000000000000000000000000000000000", 101, 101),
	makeVoter("0x3000000000000000000000000000000000000000", 99, 99),
	makeVoter("0x4000000000000000000000000000000000000000", 99, 99),
	makeVoter("0x5000000000000000000000000000000000000000", 100, 100),
	makeVoter("0x6000000000000000000000000000000000000000", 600, 600),
	makeVoter("0x7000000000000000000000000000000000000000", 300, 300),
	makeVoter("0x8000000000000000000000000000000000000000", 330, 100),
	makeVoter("0x9000000000000000000000000000000000000000", 350, 600),
	makeVoter("0x1000000000000000000000000000000000000000", 400, 300),
}

func getVoters(indexes ...int) []*Voter {
	var vs []*Voter

	m := make(map[int]struct{})
	for _, idx := range indexes {
		m[idx] = struct{}{}
	}

	for i, v := range voterSet {
		if _, ok := m[i]; ok {
			vs = append(vs, makeVoter(v.address.String(), v.deposit.Uint64(), v.weight.Uint64()))
		}
	}

	return vs
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
	voters, err := NewVoters(getVoters(0))
	require.NoError(t, err)

	voter := voters.At(-1)

	assert.Nil(t, voter)
}

func TestVoters_GetAtOverLastReturnsNil(t *testing.T) {
	voters, err := NewVoters(getVoters(0))
	require.NoError(t, err)

	voterAt := voters.At(0)
	assert.Equal(t, voterSet[0], voterAt)

	voterAt = voters.At(1)
	assert.Nil(t, voterAt)
}

func TestVoters_One(t *testing.T) {
	voters, err := NewVoters(getVoters(0))

	require.NoError(t, err)
	assert.Equal(t, 1, voters.Len())
	assert.Equal(t, voterSet[0], voters.At(0))
	assert.Equal(t, voterSet[0], voters.Get(voterSet[0].Address()))
	assert.Equal(t, true, voters.Contains(voterSet[0].Address()))
	assert.Equal(t, voterSet[0], voters.NextProposer())
}

func TestVoters_UpdateWeightChangesProposer(t *testing.T) {
	voters, err := NewVoters(getVoters(0, 1, 2))

	require.NoError(t, err)
	proposer := voters.NextProposer()
	assert.Equal(t, voterSet[1].Address(), proposer.Address())
	assert.Equal(t, big.NewInt(-98), proposer.weight)
	assert.Equal(t, big.NewInt(200), voters.At(0).weight)
	assert.Equal(t, big.NewInt(-98), voters.At(1).weight)
	assert.Equal(t, big.NewInt(198), voters.At(2).weight)
	assert.Equal(t, 3, voters.Len())
}

func TestVoters_UpdateWeightChangesProposerWith2Voters(t *testing.T) {
	voters1, err := NewVoters(getVoters(0, 1, 2))
	require.NoError(t, err)

	voters2, err := NewVoters(getVoters(0, 1, 2))
	require.NoError(t, err)

	proposer1 := voters1.NextProposer()
	assert.Equal(t, voterSet[1].Address(), proposer1.Address())
	assert.Equal(t, big.NewInt(-98), proposer1.weight)
	assert.Equal(t, big.NewInt(200), voters1.At(0).weight)
	assert.Equal(t, big.NewInt(-98), voters1.At(1).weight)
	assert.Equal(t, big.NewInt(198), voters1.At(2).weight)
	assert.Equal(t, 3, voters1.Len())

	proposer2 := voters2.NextProposer()
	assert.Equal(t, voterSet[1].Address(), proposer2.Address())
	assert.Equal(t, big.NewInt(-98), proposer2.weight)
	assert.Equal(t, big.NewInt(200), voters2.At(0).weight)
	assert.Equal(t, big.NewInt(-98), voters2.At(1).weight)
	assert.Equal(t, big.NewInt(198), voters2.At(2).weight)
	assert.Equal(t, 3, voters2.Len())
}

func TestVoters_NewVotersReturnsSortedArray(t *testing.T) {
	voters, err := NewVoters(getVoters(1, 0, 2))
	require.NoError(t, err)

	assert.Equal(t, voterSet[0].Address().String(), voters.At(0).Address().String())
	assert.Equal(t, voterSet[1].Address().String(), voters.At(1).Address().String())
	assert.Equal(t, voterSet[2].Address().String(), voters.At(2).Address().String())
	assert.Equal(t, 3, voters.Len())
}

func TestVoters_UpdateWeightChangesProposerElections(t *testing.T) {
	voters, err := NewVoters(getVoters(0, 1, 2))
	require.NoError(t, err)
	require.Equal(t, 3, voters.Len())

	elections := []struct {
		proposerWeight  *big.Int
		proposerAddress string
		voter1weight    *big.Int
		voter2weight    *big.Int
		voter3weight    *big.Int
	}{
		{
			big.NewInt(-98),
			voterSet[1].Address().String(),
			big.NewInt(200),
			big.NewInt(-98),
			big.NewInt(198),
		},
		{
			big.NewInt(0),
			voterSet[0].Address().String(),
			big.NewInt(0),
			big.NewInt(3),
			big.NewInt(297),
		},
		{
			big.NewInt(96),
			voterSet[2].Address().String(),
			big.NewInt(100),
			big.NewInt(104),
			big.NewInt(96),
		},
	}

	for round, tc := range elections {
		t.Run(fmt.Sprintf("round %d", round), func(t *testing.T) {
			proposer := voters.NextProposer()
			assert.Equal(t, tc.proposerAddress, proposer.Address().String())
			assert.Equal(t, tc.proposerWeight.Int64(), proposer.weight.Int64())
			assert.Equal(t, tc.voter1weight.Int64(), voters.At(0).weight.Int64())
			assert.Equal(t, tc.voter2weight.Int64(), voters.At(1).weight.Int64())
			assert.Equal(t, tc.voter3weight.Int64(), voters.At(2).weight.Int64())
		})
	}
}

func TestVoters_UpdateWeightChangesProposerElectionsWith2Voters(t *testing.T) {
	voters1, err := NewVoters(getVoters(0, 1, 2))
	require.NoError(t, err)
	require.Equal(t, 3, voters1.Len())

	voters2, err := NewVoters(getVoters(0, 1, 2))
	require.NoError(t, err)
	require.Equal(t, 3, voters2.Len())

	elections := []struct {
		proposerWeight  *big.Int
		proposerAddress string
		voter1weight    *big.Int
		voter2weight    *big.Int
		voter3weight    *big.Int
	}{
		{
			big.NewInt(-98),
			voterSet[1].Address().String(),
			big.NewInt(200),
			big.NewInt(-98),
			big.NewInt(198),
		},
		{
			big.NewInt(0),
			voterSet[0].Address().String(),
			big.NewInt(0),
			big.NewInt(3),
			big.NewInt(297),
		},
		{
			big.NewInt(96),
			voterSet[2].Address().String(),
			big.NewInt(100),
			big.NewInt(104),
			big.NewInt(96),
		},
	}

	for round, tc := range elections {
		t.Run(fmt.Sprintf("round %d", round), func(t *testing.T) {
			proposer1 := voters1.NextProposer()
			proposer2 := voters2.NextProposer()

			assert.Equal(t, tc.proposerWeight.Int64(), proposer2.weight.Int64())
			assert.Equal(t, voters1.At(0).address, voters2.At(0).address)
			assert.Equal(t, voters1.At(1).address, voters2.At(1).address)
			assert.Equal(t, voters1.At(2).address, voters2.At(2).address)

			assert.Equal(t, voters1.At(0).weight.Int64(), voters2.At(0).weight.Int64())
			assert.Equal(t, voters1.At(1).weight.Int64(), voters2.At(1).weight.Int64())
			assert.Equal(t, voters1.At(2).weight.Int64(), voters2.At(2).weight.Int64())

			assert.Equal(t, tc.proposerWeight.Int64(), proposer1.weight.Int64())
			assert.Equal(t, tc.voter1weight.Int64(), voters1.At(0).weight.Int64())
			assert.Equal(t, tc.voter2weight.Int64(), voters1.At(1).weight.Int64())
			assert.Equal(t, tc.voter3weight.Int64(), voters1.At(2).weight.Int64())

			assert.Equal(t, tc.proposerWeight.Int64(), proposer2.weight.Int64())
			assert.Equal(t, tc.voter1weight.Int64(), voters2.At(0).weight.Int64())
			assert.Equal(t, tc.voter2weight.Int64(), voters2.At(1).weight.Int64())
			assert.Equal(t, tc.voter3weight.Int64(), voters2.At(2).weight.Int64())
		})
	}
}

func TestVoters_UpdateWeightChangesProposerElectionsVotersShouldBeChosenWithGivenProbability(t *testing.T) {
	voters1, err := NewVoters(getVoters(4, 5, 6))
	require.NoError(t, err)
	require.Equal(t, 3, voters1.Len())

	voters2, err := NewVoters(getVoters(4, 5, 6))
	require.NoError(t, err)
	require.Equal(t, 3, voters2.Len())

	freq := make(map[string]int)
	totalRounds := 1000000

	for i := 0; i < totalRounds; i++ {
		proposer1 := voters1.NextProposer()
		proposer2 := voters2.NextProposer()

		assert.Equal(t, proposer1.Address(), proposer2.Address())
		assert.Equal(t, proposer1.Deposit(), proposer2.Deposit())

		assert.Equal(t, voters1.At(0).weight, voters2.At(0).weight)
		assert.Equal(t, voters1.At(1).weight, voters2.At(1).weight)
		assert.Equal(t, voters1.At(2).weight, voters2.At(2).weight)

		freq[proposer1.Address().String()]++
	}

	totalDeposit := float64(voters1.At(0).Deposit().Int64() + voters1.At(1).Deposit().Int64() + voters1.At(2).Deposit().Int64())
	expectedFreq := map[string]float64{
		voters1.At(0).Address().String(): float64(voters1.At(0).deposit.Int64()) / totalDeposit,
		voters1.At(1).Address().String(): float64(voters1.At(1).deposit.Int64()) / totalDeposit,
		voters1.At(2).Address().String(): float64(voters1.At(2).deposit.Int64()) / totalDeposit,
	}

	epsilon := 0.0000001
	for address, calculatedFreq := range expectedFreq {
		count, ok := freq[address]
		assert.Truef(t, ok, "address '%v' hadn't been chosen", address)

		addressFreq := float64(count) / float64(totalRounds)
		assert.InEpsilonf(t, calculatedFreq, addressFreq, epsilon, "expected for '%v' %.4f, got %.4f",
			address, calculatedFreq, addressFreq)
	}
}

func TestVoters_IsHashable(t *testing.T) {
	voters1, err := NewVoters(getVoters(0, 1, 2))
	require.NoError(t, err)

	voters2, err := NewVoters(getVoters(1, 2, 3))
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
