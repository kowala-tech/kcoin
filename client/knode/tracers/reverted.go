package tracers

import (
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client/contracts/bindings"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/vm"
)

type EvmRevertedTracer struct {
}

func (*EvmRevertedTracer) CaptureStart(from common.Address, to common.Address, call bool, input []byte, gas uint64, value *big.Int) error {
	return nil
}

func (*EvmRevertedTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	return nil
}

func (*EvmRevertedTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	contractName, cErr := bindings.GetContractByAddr(contract.Address())
	if cErr != nil {
		contractName = "undetected"
	}

	if err.Error() == "evm: execution reverted" {
		fmt.Printf(
			"error with transaction from address: %s to address: %s {opcode: %s (%s) pc: %d Contract Name: %s Error msg: %s}\n",
			contract.CallerAddress.String(),
			contract.Address().String(),
			op.String(),
			fmt.Sprintf("%s%s", "0x", common.Bytes2Hex([]byte{byte(op)})),
			pc,
			contractName,
			err,
		)
	}

	return nil
}

func (*EvmRevertedTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return nil
}
