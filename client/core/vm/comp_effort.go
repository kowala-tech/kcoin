// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/params"
)

// computational effort
const (
	CompEffortQuickStep   uint64 = 2
	CompEffortFastestStep uint64 = 3
	CompEffortFastStep    uint64 = 5
	CompEffortMidStep     uint64 = 8
	CompEffortSlowStep    uint64 = 10
	CompEffortExtStep     uint64 = 20

	CompEffortReturn       uint64 = 0
	CompEffortStop         uint64 = 0
	CompEffortContractByte uint64 = 200
)

// calcComputationalEffort returns the actual computational effort required by the call.
func calcComputationalEffort(gasTable params.GasTable, availableComputationalResources, base uint64, callCost *big.Int) (uint64, error) {
	if gasTable.CreateBySuicide > 0 {
		availableComputationalResources = availableComputationalResources - base
		effort := availableComputationalResources - availableComputationalResources/64
		// If the bit length exceeds 64 bit we know that the newly calculated "gas" for EIP150
		// is smaller than the requested amount. Therefor we return the new gas instead
		// of returning an error.
		if callCost.BitLen() > 64 || effort < callCost.Uint64() {
			return effort, nil
		}
	}
	if callCost.BitLen() > 64 {
		return 0, errComputationalEffortUintOverflow
	}

	return callCost.Uint64(), nil
}
