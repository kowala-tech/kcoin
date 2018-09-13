package core

import (
	"errors"
	"fmt"
	"math"
)

var (
	// ErrComputeCapacityReached is returned by the computational resources pool if the computational effort required
	// by a transaction is higher than what's left in the block.
	ErrComputeCapacityReached = errors.New("compute capacity reached")
)

// ComputationalResourcePool tracks the computational resource available in compute units during
// the execution of the transactions in a block.
type ComputationalResourcePool uint64

// AddResource makes computational resource available for execution.
func (pool *ComputationalResourcePool) AddResource(units uint64) *ComputationalResourcePool {
	if uint64(*pool) > math.MaxUint64-units {
		panic("computational resource pool pushed above uint64")
	}
	*(*uint64)(pool) += units
	return pool
}

// SubResource deducts the given computational resource units from the pool if enough units are
// available and returns an error otherwise.
func (pool *ComputationalResourcePool) SubResource(units uint64) error {
	if uint64(*pool) < units {
		return ErrComputeCapacityReached
	}
	*(*uint64)(pool) -= units
	return nil
}

// Value returns the amount of computational resource (compute units) remaining in the pool.
func (pool *ComputationalResourcePool) Value() uint64 {
	return uint64(*pool)
}

func (pool *ComputationalResourcePool) String() string {
	return fmt.Sprintf("%d", *pool)
}
