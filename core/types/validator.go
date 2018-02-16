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

type ValidatorSet struct {
	validators []*Validator
	proposer   *Validator

	//cache
	membership map[common.Address]*Validator
}

// @TODO (rgeraldes) - size needs to be > 0
func NewValidatorSet(validators []*Validator) *ValidatorSet {
	set := &ValidatorSet{
		validators: validators,
		membership: make(map[common.Address]*Validator, len(validators)),
	}

	for _, validator := range validators {
		set.membership[validator.address] = validator
	}

	return set
}

// Update updates the weight and the proposer based on the set of validators
func (set *ValidatorSet) UpdateWeight() {
	pq := make(common.PriorityQueue, 0)
	heap.Init(&pq)

	for _, validator := range set.validators {
		validator.weight = validator.weight.Add(validator.weight, big.NewInt(int64(validator.deposit)))
		// @TODO (rgeraldes) - review types, possible overflow
		heap.Push(&pq, &common.Item{Priority: int(validator.weight.Int64()), Value: validator})
	}

	item := heap.Pop(&pq).(*common.Item)
	set.proposer = item.Value.(*Validator)

	// decrement the validator weight since he has been selected
	set.proposer.weight.Sub(set.proposer.weight, big.NewInt(int64(set.proposer.deposit)))
}

func (set *ValidatorSet) AtIndex(i int) *Validator {
	return set.validators[i]
}

func (set *ValidatorSet) Get(addr common.Address) *Validator {
	return set.membership[addr]
}

func (set *ValidatorSet) Size() int {
	return len(set.validators)
}

func (set *ValidatorSet) Proposer() common.Address {
	return set.proposer.address
}

func (set *ValidatorSet) Contains(addr common.Address) bool {
	_, ok := set.membership[addr]
	return ok
}
