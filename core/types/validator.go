package types

import (
	"github.com/kowala-tech/kUSD/common"
)

// validator represents a consensus validator
type Validator struct {
	code  common.Address // @TODO (rgeraldes) -  coinbase for now
	power uint64
}

func NewValidator(code common.Address, power uint64) *Validator {
	return &Validator{
		code:  code,
		power: power,
	}
}

func (val *Validator) Hash() common.Hash    { return rlpHash(val) }
func (val *Validator) Code() common.Address { return val.code }
func (val *Validator) Power() uint64        { return val.power }

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
func (set *ValidatorSet) Size() int {
	return len(set.validators)
}

func (set *ValidatorSet) Proposer() common.Address {
	// @TODO (rgeraldes) complete - return the first validator for now
	return set.validators[0].Code()
}
