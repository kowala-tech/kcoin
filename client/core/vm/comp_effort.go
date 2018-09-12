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

	"github.com/kowala-tech/kcoin/client/params/effort"
)

// Computational Effort
const (
	EffortQuickStep   uint64 = 2
	EffortFastestStep uint64 = 3
	EffortFastStep    uint64 = 5
	EffortMidStep     uint64 = 8
	EffortSlowStep    uint64 = 10
	EffortExtStep     uint64 = 20

	EffortReturn       uint64 = 0
	EffortStop         uint64 = 0
	EffortContractByte uint64 = 200
)

// calcEffort returns the actual computational effort of the call.
func callEffort(table effort.Table, availableResource, base uint64, callCost *big.Int) (uint64, error) {
	if table.CreateBySuicide > 0 {
		availableResource = availableResource - base
		resource := availableResource - availableResource/64
		// If the bit length exceeds 64 bit we know that the newly calculated "computational effort" for EIP150
		// is smaller than the requested amount. Therefore we return the new computational effort instead
		// of returning an error.
		if callCost.BitLen() > 64 || resource < callCost.Uint64() {
			return resource, nil
		}
	}
	if callCost.BitLen() > 64 {
		return 0, errEffortUintOverflow
	}

	return callCost.Uint64(), nil
}
