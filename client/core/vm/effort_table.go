// Copyright 2017 The go-ethereum Authors
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
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/math"
	effrt "github.com/kowala-tech/kcoin/client/params/effort"
)

// memoryEffort calculates the quadratic computational effort for memory expansion. It does so
// only for the memory region that is expanded, not the total memory.
func memoryEffort(mem *Memory, newMemSize uint64) (uint64, error) {

	if newMemSize == 0 {
		return 0, nil
	}
	// The maximum that will fit in a uint64 is max_word_count - 1
	// anything above that will result in an overflow.
	// Additionally, a newMemSize which results in a
	// newMemSizeWords larger than 0x7ffffffff will cause the square operation
	// to overflow.
	// The constant 0xffffffffe0 is the highest number that can be used without
	// overflowing the effort calculation
	if newMemSize > 0xffffffffe0 {
		return 0, errEffortUintOverflow
	}

	newMemSizeWords := toWordSize(newMemSize)
	newMemSize = newMemSizeWords * 32

	if newMemSize > uint64(mem.Len()) {
		square := newMemSizeWords * newMemSizeWords
		linCoef := newMemSizeWords * effrt.Memory
		quadCoef := square / effrt.QuadCoeffDiv
		newTotalFee := linCoef + quadCoef

		fee := newTotalFee - mem.lastResourceUsage
		mem.lastResourceUsage = newTotalFee

		return fee, nil
	}
	return 0, nil
}

func constEffortFunc(effort uint64) effortFunc {
	return func(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
		return effort, nil
	}
}

func effortCallDataCopy(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if effort, overflow = math.SafeAdd(effort, EffortFastestStep); overflow {
		return 0, errEffortUintOverflow
	}

	words, overflow := bigUint64(stack.Back(2))
	if overflow {
		return 0, errEffortUintOverflow
	}

	if words, overflow = math.SafeMul(toWordSize(words), effrt.Copy); overflow {
		return 0, errEffortUintOverflow
	}

	if effort, overflow = math.SafeAdd(effort, words); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortReturnDataCopy(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if effort, overflow = math.SafeAdd(effort, EffortFastestStep); overflow {
		return 0, errEffortUintOverflow
	}

	words, overflow := bigUint64(stack.Back(2))
	if overflow {
		return 0, errEffortUintOverflow
	}

	if words, overflow = math.SafeMul(toWordSize(words), effrt.Copy); overflow {
		return 0, errEffortUintOverflow
	}

	if effort, overflow = math.SafeAdd(effort, words); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortSStore(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var (
		y, x = stack.Back(1), stack.Back(0)
		val  = vm.StateDB.GetState(contract.Address(), common.BigToHash(x))
	)
	// This checks for 3 scenario's and calculates effort accordingly
	// 1. From a zero-value address to a non-zero value         (NEW VALUE)
	// 2. From a non-zero value address to a zero-value address (DELETE)
	// 3. From a non-zero to a non-zero                         (CHANGE)
	if val == (common.Hash{}) && y.Sign() != 0 {
		// 0 => non 0
		return effrt.SstoreSet, nil
	} else if val != (common.Hash{}) && y.Sign() == 0 {
		// non 0 => 0
		vm.StateDB.AddRefund(effrt.SstoreRefund)
		return effrt.SstoreClear, nil
	} else {
		// non 0 => non 0 (or 0 => 0)
		return effrt.SstoreReset, nil
	}
}

func makeEffortLog(n uint64) effortFunc {
	return func(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
		requestedSize, overflow := bigUint64(stack.Back(1))
		if overflow {
			return 0, errEffortUintOverflow
		}

		effort, err := memoryEffort(mem, memorySize)
		if err != nil {
			return 0, err
		}

		if effort, overflow = math.SafeAdd(effort, effrt.Log); overflow {
			return 0, errEffortUintOverflow
		}
		if effort, overflow = math.SafeAdd(effort, n*effrt.LogTopic); overflow {
			return 0, errEffortUintOverflow
		}

		var memorySizeEffort uint64
		if memorySizeEffort, overflow = math.SafeMul(requestedSize, effrt.LogData); overflow {
			return 0, errEffortUintOverflow
		}
		if effort, overflow = math.SafeAdd(effort, memorySizeEffort); overflow {
			return 0, errEffortUintOverflow
		}
		return effort, nil
	}
}

func effortSha3(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}

	if effort, overflow = math.SafeAdd(effort, effrt.Sha3); overflow {
		return 0, errEffortUintOverflow
	}

	wordEffort, overflow := bigUint64(stack.Back(1))
	if overflow {
		return 0, errEffortUintOverflow
	}
	if wordEffort, overflow = math.SafeMul(toWordSize(wordEffort), effrt.Sha3Word); overflow {
		return 0, errEffortUintOverflow
	}
	if effort, overflow = math.SafeAdd(effort, wordEffort); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortCodeCopy(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if effort, overflow = math.SafeAdd(effort, EffortFastestStep); overflow {
		return 0, errEffortUintOverflow
	}

	wordEffort, overflow := bigUint64(stack.Back(2))
	if overflow {
		return 0, errEffortUintOverflow
	}
	if wordEffort, overflow = math.SafeMul(toWordSize(wordEffort), effrt.Copy); overflow {
		return 0, errEffortUintOverflow
	}
	if effort, overflow = math.SafeAdd(effort, wordEffort); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortExtCodeCopy(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}

	var overflow bool
	if effort, overflow = math.SafeAdd(effort, gt.ExtcodeCopy); overflow {
		return 0, errEffortUintOverflow
	}

	wordEffort, overflow := bigUint64(stack.Back(3))
	if overflow {
		return 0, errEffortUintOverflow
	}

	if wordEffort, overflow = math.SafeMul(toWordSize(wordEffort), effrt.Copy); overflow {
		return 0, errEffortUintOverflow
	}

	if effort, overflow = math.SafeAdd(effort, wordEffort); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortMLoad(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, errEffortUintOverflow
	}
	if effort, overflow = math.SafeAdd(effort, EffortFastestStep); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortMStore8(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, errEffortUintOverflow
	}
	if effort, overflow = math.SafeAdd(effort, EffortFastestStep); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortMStore(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, errEffortUintOverflow
	}
	if effort, overflow = math.SafeAdd(effort, EffortFastestStep); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortCreate(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var overflow bool
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}
	if effort, overflow = math.SafeAdd(effort, effrt.Create); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortBalance(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return gt.Balance, nil
}

func effortExtCodeSize(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return gt.ExtcodeSize, nil
}

func effortSLoad(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return gt.SLoad, nil
}

func effortExp(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	expByteLen := uint64((stack.data[stack.len()-2].BitLen() + 7) / 8)

	var (
		effort   = expByteLen * gt.ExpByte // no overflow check required. Max is 256 * ExpByte effort
		overflow bool
	)
	if effort, overflow = math.SafeAdd(effort, EffortSlowStep); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortCall(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var (
		effort         = gt.Calls
		transfersValue = stack.Back(2).Sign() != 0
		address        = common.BigToAddress(stack.Back(1))
	)

	if transfersValue && vm.StateDB.Empty(address) {
		effort += effrt.CallNewAccount
	}

	if transfersValue {
		effort += effrt.CallValueTransfer
	}
	memoryEffort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if effort, overflow = math.SafeAdd(effort, memoryEffort); overflow {
		return 0, errEffortUintOverflow
	}

	vm.callResourceTemp, err = callEffort(gt, contract.ComputationalResource, effort, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if effort, overflow = math.SafeAdd(effort, vm.callResourceTemp); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortCallCode(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	effort := gt.Calls
	if stack.Back(2).Sign() != 0 {
		effort += effrt.CallValueTransfer
	}
	memoryEffort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if effort, overflow = math.SafeAdd(effort, memoryEffort); overflow {
		return 0, errEffortUintOverflow
	}

	vm.callResourceTemp, err = callEffort(gt, contract.ComputationalResource, effort, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if effort, overflow = math.SafeAdd(effort, vm.callResourceTemp); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortReturn(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return memoryEffort(mem, memorySize)
}

func effortRevert(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return memoryEffort(mem, memorySize)
}

func effortSuicide(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	var effort uint64

	effort = gt.Suicide
	var (
		address = common.BigToAddress(stack.Back(0))
	)

	// if empty and transfers value
	if vm.StateDB.Empty(address) && vm.StateDB.GetBalance(contract.Address()).Sign() != 0 {
		effort += gt.CreateBySuicide
	}

	if !vm.StateDB.HasSuicided(contract.Address()) {
		vm.StateDB.AddRefund(effrt.SuicideRefund)
	}
	return effort, nil
}

func effortDelegateCall(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if effort, overflow = math.SafeAdd(effort, gt.Calls); overflow {
		return 0, errEffortUintOverflow
	}

	vm.callResourceTemp, err = callEffort(gt, contract.ComputationalResource, effort, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if effort, overflow = math.SafeAdd(effort, vm.callResourceTemp); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortStaticCall(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	effort, err := memoryEffort(mem, memorySize)
	if err != nil {
		return 0, err
	}
	var overflow bool
	if effort, overflow = math.SafeAdd(effort, gt.Calls); overflow {
		return 0, errEffortUintOverflow
	}

	vm.callResourceTemp, err = callEffort(gt, contract.ComputationalResource, effort, stack.Back(0))
	if err != nil {
		return 0, err
	}
	if effort, overflow = math.SafeAdd(effort, vm.callResourceTemp); overflow {
		return 0, errEffortUintOverflow
	}
	return effort, nil
}

func effortPush(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return EffortFastestStep, nil
}

func effortSwap(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return EffortFastestStep, nil
}

func effortDup(gt effrt.Table, vm *VM, contract *Contract, stack *Stack, mem *Memory, memorySize uint64) (uint64, error) {
	return EffortFastestStep, nil
}
