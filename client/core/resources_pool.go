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

// CompResourcesPool tracks the computational resources available in compute units during
// the execution of the transactions in a block.
type CompResourcesPool uint64

// AddResources makes computational resources available for execution.
func (pool *CompResourcesPool) AddResources(units uint64) *CompResourcesPool {
	if uint64(*pool) > math.MaxUint64-units {
		panic("computational resources pool pushed above uint64")
	}
	*(*uint64)(pool) += units
	return pool
}

// SubResources deducts the given compute units from the pool if enough computational resources are
// available and returns an error otherwise.
func (pool *CompResourcesPool) SubResources(units uint64) error {
	if uint64(*pool) < units {
		return ErrComputeCapacityReached
	}
	*(*uint64)(pool) -= units
	return nil
}

// Resources returns the computational resources in compute units remaining in the pool.
func (pool *CompResourcesPool) Resources() uint64 {
	return uint64(*pool)
}

func (pool *CompResourcesPool) String() string {
	return fmt.Sprintf("%d", *pool)
}
