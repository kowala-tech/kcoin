package types

import (
	"math/big"
	"time"

	"errors"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/rlp"
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

type ValidatorList interface {
	UpdateWeights()
	At(i int) *Validator
	Get(addr common.Address) *Validator
	Len() int
	Proposer() *Validator
	Contains(addr common.Address) bool
	Hash() common.Hash
}

var ErrInvalidParams = errors.New("A validator set needs at least one validator")

func NewValidatorList(validators []*Validator) (*validatorList, error) {
	if len(validators) == 0 {
		return nil, ErrInvalidParams
	}

	set := &validatorList{
		validators: validators,
		proposer:   validators[0],
	}

	return set, nil
}

type validatorList struct {
	validators []*Validator
	proposer   *Validator
}

// Update updates the weight and the proposer based on the set of validators
func (set *validatorList) UpdateWeights() {
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

func (set *validatorList) At(i int) *Validator {
	if i < 0 || i >= len(set.validators) {
		return nil
	}
	return set.validators[i]
}

func (set *validatorList) Get(addr common.Address) *Validator {
	for _, validator := range set.validators {
		if validator.Address() == addr {
			return validator
		}
	}
	return nil
}

func (set *validatorList) Len() int {
	return len(set.validators)
}

func (set *validatorList) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(set.validators[i])
	return enc
}

func (set *validatorList) Hash() common.Hash {
	return DeriveSha(set)
}

func (set *validatorList) Proposer() *Validator {
	return set.proposer
}

func (set *validatorList) Contains(addr common.Address) bool {
	validator := set.Get(addr)
	return validator != nil
}

func NewDeposit(amount uint64, unixTimestamp int64) *Deposit {
	return &Deposit{
		amount:      amount,
		availableAt: time.Unix(unixTimestamp, 0),
	}
}

// Deposit represents the validator deposits at stake
type Deposit struct {
	amount      uint64
	availableAt time.Time
}

func (dep *Deposit) Amount() uint64 {
	return dep.amount
}

func (dep *Deposit) AvailableAt() time.Time {
	return dep.availableAt
}
