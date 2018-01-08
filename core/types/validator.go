package types

import "github.com/kowala-tech/kUSD/common"

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

type Validators struct {
	validators []*Validator
	proposer   *Validator
}

func (vals Validators) Size() int {
	return len(vals.validators)
}

func (vals Validators) Proposer() common.Address {
	return vals.proposer.Code()
}
