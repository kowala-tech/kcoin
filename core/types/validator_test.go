package types

import (
	"fmt"
	"github.com/kowala-tech/kcoin/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestValidator_Properties(t *testing.T) {
	address := common.Address{}
	deposit := uint64(100)
	weight := &big.Int{}
	validator := NewValidator(address, deposit, weight)

	assert.Equal(t, address, validator.Address())
	assert.Equal(t, deposit, validator.Deposit())
	assert.Equal(t, weight, validator.Weight())
}

func TestValidatorSet_EmptyReturnsError(t *testing.T) {
	validatorList, err := NewValidatorList(nil)

	require.Error(t, err)
	require.Nil(t, validatorList)
}

func TestValidatorSet_GetAtNegativeIndexReturnsNil(t *testing.T) {
	validator := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	validatorList, err := NewValidatorList([]*Validator{validator})
	require.NoError(t, err)

	validator = validatorList.At(-1)

	assert.Nil(t, validator)
}

func TestValidatorSet_GetAtOverLastReturnsNil(t *testing.T) {
	validator := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	validatorList, err := NewValidatorList([]*Validator{validator})
	require.NoError(t, err)

	validatorAt := validatorList.At(0)
	assert.Equal(t, validator, validatorAt)

	validatorAt = validatorList.At(1)
	assert.Nil(t, validatorAt)
}

func TestValidatorSet_One(t *testing.T) {
	address := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")
	deposit := uint64(100)
	weight := &big.Int{}
	validator := NewValidator(address, deposit, weight)

	validatorList, err := NewValidatorList([]*Validator{validator})

	require.NoError(t, err)
	assert.Equal(t, 1, validatorList.Size())
	assert.Equal(t, validator, validatorList.At(0))
	assert.Equal(t, validator, validatorList.Get(address))
	assert.Equal(t, true, validatorList.Contains(address))
	assert.Equal(t, validator, validatorList.Proposer())
}

func TestValidatorSet_UpdateWeightChangesProposer(t *testing.T) {
	validator := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	validator2 := makeValidator("0x6aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 101, 101)
	validator3 := makeValidator("0x7aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 99, 99)

	validatorList, err := NewValidatorList([]*Validator{validator, validator2, validator3})

	require.NoError(t, err)
	validatorList.UpdateWeights()
	assert.Equal(t, validator2, validatorList.Proposer())
	assert.Equal(t, big.NewInt(101), validatorList.Proposer().weight)
	assert.Equal(t, big.NewInt(200), validatorList.At(0).weight)
	assert.Equal(t, big.NewInt(101), validatorList.At(1).weight)
	assert.Equal(t, big.NewInt(198), validatorList.At(2).weight)
	assert.Equal(t, 3, validatorList.Size())
}

func TestValidatorSet_UpdateWeightChangesProposerElections(t *testing.T) {
	validator := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	validator2 := makeValidator("0x6aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 101, 101)
	validator3 := makeValidator("0x7aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 99, 99)

	validatorList, err := NewValidatorList([]*Validator{validator, validator2, validator3})
	require.NoError(t, err)
	require.Equal(t, 3, validatorList.Size())

	elections := []struct {
		proposerWeight   *big.Int
		validator1weight *big.Int
		validator2weight *big.Int
		validator3weight *big.Int
	}{
		{big.NewInt(101), big.NewInt(200), big.NewInt(101), big.NewInt(198)},
		{big.NewInt(200), big.NewInt(200), big.NewInt(202), big.NewInt(297)},
		{big.NewInt(297), big.NewInt(300), big.NewInt(303), big.NewInt(297)},
		{big.NewInt(303), big.NewInt(400), big.NewInt(303), big.NewInt(396)},
	}

	for round, tc := range elections {
		t.Run(fmt.Sprintf("round %d", round), func(t *testing.T) {
			validatorList.UpdateWeights()
			assert.Equal(t, tc.proposerWeight, validatorList.Proposer().weight)
			assert.Equal(t, tc.validator1weight, validatorList.At(0).weight)
			assert.Equal(t, tc.validator2weight, validatorList.At(1).weight)
			assert.Equal(t, tc.validator3weight, validatorList.At(2).weight)
		})
	}
}

func makeValidator(hexAddress string, deposit int, weight int64) *Validator {
	address := common.HexToAddress(hexAddress)
	return NewValidator(address, uint64(deposit), big.NewInt(weight))
}
