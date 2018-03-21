package types

import (
	"math/big"

	"errors"
	"github.com/kowala-tech/kUSD/common"
)

// Voter represents a consensus validator
type Voter struct {
	address common.Address
	deposit uint64
	weight  *big.Int
}

// NewVoter returns a new validator instance
func NewVoter(address common.Address, deposit uint64, weight *big.Int) *Voter {
	return &Voter{
		address: address,
		deposit: deposit,
		weight:  weight,
	}
}

func (val *Voter) Address() common.Address { return val.address }
func (val *Voter) Deposit() uint64         { return val.deposit }
func (val *Voter) Weight() *big.Int        { return val.weight }

type Voters interface {
	At(i int) *Voter
	Get(addr common.Address) *Voter
	Size() int
	NextProposer() *Voter
}

var ErrInvalidParams = errors.New("A validator set needs at least one validator")

func NewVoters(validators []*Voter) (*validatorList, error) {
	if len(validators) == 0 {
		return nil, ErrInvalidParams
	}

	set := &validatorList{
		validators: validators,
	}

	return set, nil
}

type validatorList struct {
	validators []*Voter
}

// Update updates the weight and the proposer based on the set of validators
func (set *validatorList) NextProposer() *Voter {
	proposer := set.validators[0]

	for _, validator := range set.validators {
		validator.weight = validator.weight.Add(validator.weight, big.NewInt(int64(validator.deposit)))
		if validator.weight.Cmp(proposer.weight) > 0 {
			proposer = validator
		}
	}

	// decrement the validator weight since he has been selected
	proposer.weight.Sub(proposer.weight, big.NewInt(int64(proposer.deposit)))

	return proposer
}

func (set *validatorList) At(i int) *Voter {
	if i < 0 || i >= len(set.validators) {
		return nil
	}
	return set.validators[i]
}

func (set *validatorList) Get(addr common.Address) *Voter {
	for _, validator := range set.validators {
		if validator.Address() == addr {
			return validator
		}
	}
	return nil
}

func (set *validatorList) Size() int {
	return len(set.validators)
}
