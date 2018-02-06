package types

import (
	"container/heap"
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

// @TODO (rgeraldes) - size needs to be > 0
func NewValidatorSet(validators []*Validator) *ValidatorSet {
	set := &ValidatorSet{
		validators: validators,
	}

	return set
}

// Update updates the weight and the proposer based on the validator set
func (set *ValidatorSet) UpdateWeight() {
	pq := make(common.PriorityQueue, len(set.validators))
	heap.Init(&pq)

	for _, validator := range set.validators {
		validator.weight = validator.weight.Add(validator.weight, big.NewInt(int64(validator.deposit)))
		heap.Push(&pq, &common.Item{Value: validator, Priority: int(validator.deposit)})
	}

	proposer := heap.Pop(&pq).(*Validator)
	set.proposer = proposer

	// decrement the validator weight since he has been selected
	proposer.weight.Sub(proposer.weight, big.NewInt(int64(proposer.deposit)))
}

func (set *ValidatorSet) AtIndex(i int) *Validator {
	return set.validators[i]
}

func (set *ValidatorSet) Size() int {
	return len(set.validators)
}

func (set *ValidatorSet) Proposer() common.Address {
	return set.proposer.address
}
