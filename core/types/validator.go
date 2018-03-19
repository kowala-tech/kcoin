package types

import (
	"math/big"

	"errors"
	"github.com/kowala-tech/kUSD/common"
)

// Validator represents a consensus validator
type Validator struct {
	address common.Address
	deposit uint64
	weight  *big.Int
}

// NewValidator returns a new validator instance
func NewValidator(address common.Address, deposit uint64, weight *big.Int) *Validator {
	return &Validator{
		address: address,
		deposit: deposit,
		weight:  weight,
	}
}

func (val *Validator) Address() common.Address { return val.address }
func (val *Validator) Deposit() uint64         { return val.deposit }
func (val *Validator) Weight() *big.Int        { return val.weight }

type ValidatorSet interface {
	UpdateWeight()
	AtIndex(i int) *Validator
	Get(addr common.Address) *Validator
	Size() int
	Proposer() *Validator
	Contains(addr common.Address) bool
}

var ErrInvalidParams = errors.New("A validator set needs at leat one validator")

func NewValidatorSet(validators []*Validator) (*validatorSet, error) {
	if len(validators) == 0 {
		return nil, ErrInvalidParams
	}

	set := &validatorSet{
		validators: validators,
		proposer:   validators[0],
	}

	return set, nil
}

type validatorSet struct {
	validators []*Validator
	proposer   *Validator
}

// Update updates the weight and the proposer based on the set of validators
func (set *validatorSet) UpdateWeight() {
	proposer := set.validators[0]

	for _, validator := range set.validators {
		validator.weight = validator.weight.Add(validator.weight, big.NewInt(int64(validator.deposit)))
		if validator.weight.Cmp(proposer.weight) > 0 {
			proposer = validator
		}
	}
	set.proposer = proposer

	// decrement the validator weight since he has been selected
	set.proposer.weight.Sub(set.proposer.weight, big.NewInt(int64(set.proposer.deposit)))
}

func (set *validatorSet) AtIndex(i int) *Validator {
	if i > len(set.validators) {
		return nil
	}
	return set.validators[i]
}

func (set *validatorSet) Get(addr common.Address) *Validator {
	for _, validator := range set.validators {
		if validator.Address() == addr {
			return validator
		}
	}
	return nil
}

func (set *validatorSet) Size() int {
	return len(set.validators)
}

func (set *validatorSet) Proposer() *Validator {
	return set.proposer
}

func (set *validatorSet) Contains(addr common.Address) bool {
	validator := set.Get(addr)
	return validator != nil
}
