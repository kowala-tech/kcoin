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
	if err.Error() == "evm: execution reverted" {
		contractName, cErr := e.getContractByAddr(contract.Address(), env)
		if cErr != nil {
			contractName = "Undetected Contract"
		}

		callerContractName, cErr := e.getContractByAddr(contract.CallerAddress, env)
		if cErr != nil {
			callerContractName = "Undetected Contract"
		}

		mapper, err := mapping.NewFromCombinedRuntime("../contracts/bindings/sysvars/build/combined.json")
		if err != nil {
			log.Error(fmt.Sprintf("error %s", err))
			return fmt.Errorf("error %s", err)
		}

		lineContent, err := mapper.GetSolidityLineByPc(pc)
		if err != nil {
			log.Error(fmt.Sprintf("error getting line from contract: %s", err))
			return fmt.Errorf("error getting line from contract: %s", err)
		}

		log.Error(
			fmt.Sprintf(
				"error with transaction from address: %s (%s) to address: %s (%s) {opcode: %s (%s) pc: %d Aproximate solidity line content: \"%s\"Error msg: %s}\n",
				contract.CallerAddress.String(),
				callerContractName,
				contract.Address().String(),
				contractName,
				op.String(),
				fmt.Sprintf("%s%s", "0x", common.Bytes2Hex([]byte{byte(op)})),
				pc,
				lineContent,
				err,
			),
		)
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
