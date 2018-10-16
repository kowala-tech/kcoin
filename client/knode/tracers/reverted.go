package tracers

import (
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client/contracts/mapping"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/log"
)

type EvmRevertedTracer struct {
}

func (*EvmRevertedTracer) CaptureStart(from common.Address, to common.Address, call bool, input []byte, gas uint64, value *big.Int) error {
	return nil
}

func (*EvmRevertedTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	return nil
}

func (e *EvmRevertedTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	contractName, cErr := e.getContractByAddr(contract.Address(), env)
	if cErr != nil {
		contractName = "Undetected Contract"
	}

	callerContractName, cErr := e.getContractByAddr(contract.CallerAddress, env)
	if cErr != nil {
		callerContractName = "Undetected Contract"
	}

	if err.Error() == "evm: execution reverted" {
		log.Error(
			fmt.Sprintf(
				"error with transaction from address: %s (%s) to address: %s (%s) {opcode: %s (%s) pc: %d Error msg: %s}\n",
				contract.CallerAddress.String(),
				callerContractName,
				contract.Address().String(),
				contractName,
				op.String(),
				fmt.Sprintf("%s%s", "0x", common.Bytes2Hex([]byte{byte(op)})),
				pc,
				err,
			),
		)
	}

	_, err = mapping.NewFromCombinedRuntime("../../contracts/bindings/build/combined.json")
	if err != nil {
		return fmt.Errorf("error %s", err)
	}

	return nil
}

func (*EvmRevertedTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return nil
}

func (e *EvmRevertedTracer) getContractByAddr(addr common.Address, env *vm.EVM) (string, error) {
	contractName, err := bindings.GetContractByAddr(addr)
	if err == nil {
		return contractName, nil
	}

	contractName, err = kns.GetContractFromRegisteredDomains(addr, env)
	if err != nil {
		return "", err
	}

	return contractName, nil
}
