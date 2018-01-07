package types

import "github.com/kowala-tech/kUSD/common"

// validator represents a consensus validator
type Validator struct {
	code  common.Hash // @TODO (rgeraldes) -  coinbase for now
	power uint64
}

func NewValidator(code common.Hash, power uint64) *Validator {
	return &Validator{
		code:  code,
		power: power,
	}
}

func (val *Validator) Hash() common.Hash { return rlpHash(val) }
func (val *Validator) Code() common.Hash { return val.code }
func (val *Validator) Power() uint64     { return val.power }
