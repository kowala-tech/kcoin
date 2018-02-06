package types

import (
	"math/big"

	"github.com/kowala-tech/kUSD/common"
)

// Validator represents a consensus validator
type Validator struct {
	address common.Address
	deposit uint64
	weight  *big.Int
}

// NewValidator returns a new validator instance
func NewValidator(address common.Address, deposit uint64) *Validator {
	return &Validator{
		address: address,
		deposit: deposit,
		weight:  big.NewInt(0),
	}
}

func (val *Validator) Hash() common.Hash {
	return rlpHash([]interface{}{val.address, val.deposit})
}
func (val *Validator) Address() common.Address { return val.address }
func (val *Validator) Deposit() uint64         { return val.deposit }

type ValidatorSet struct {
	validators []*Validator
	proposer   *Validator
}

func NewValidatorSet(validators []*Validator) *ValidatorSet {
	// @TODO (rgeraldes) - size needs to be > 0
	return &ValidatorSet{
		validators: validators,
	}
}

func (set *ValidatorSet) AtIndex(i int) *Validator {
	return set.validators[i]
}

func (set *ValidatorSet) Size() int {
	return len(set.validators)
}

func (set *ValidatorSet) Proposer() common.Address {
	// @TODO (rgeraldes) complete - return the first validator for now
	return set.validators[0].Address()
}
