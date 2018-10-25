package tracers

import (
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	kns2 "github.com/kowala-tech/kcoin/client/contracts/bindings/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/proxy"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/sysvars"
	"github.com/kowala-tech/kcoin/client/contracts/mapping"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

var srcMapContractsData = map[string][]byte{
	"Proxy Factory Contract":                                  []byte(proxy.UpgradeabilityProxyFactorySrcMap),
	"Proxy KNS Registry Contract":                             []byte(kns2.KNSRegistrySrcMap),
	"Proxy Registrar Contract":                                []byte(kns2.FIFSRegistrarSrcMap),
	"Proxy Resolver Contract":                                 []byte(kns2.PublicResolverSrcMap),
	"Multisig Wallet Contract":                                []byte(ownership.MultiSigWalletSrcMap),
	params.KNSDomains[params.MultiSigDomain].FullDomain():     []byte(ownership.MultiSigWalletSrcMap),
	params.KNSDomains[params.OracleMgrDomain].FullDomain():    []byte(oracle.OracleMgrSrcMap),
	params.KNSDomains[params.ValidatorMgrDomain].FullDomain(): []byte(consensus.ValidatorMgrSrcMap),
	params.KNSDomains[params.MiningTokenDomain].FullDomain():  []byte(consensus.MiningTokenSrcMap),
	params.KNSDomains[params.SystemVarsDomain].FullDomain():   []byte(sysvars.SystemVarsSrcMap),
}

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

		lineContent := "cannot get the code line"
		srcMap, ok := srcMapContractsData[contractName]
		if ok {
			mapper, cErr := mapping.NewFromCombinedRuntime(
				srcMap,
				&mapping.Config{UseBinding: true},
			)
			if cErr != nil {
				return fmt.Errorf("error %s", cErr)
			}

			lineContent, cErr = mapper.GetSolidityLineByPc(pc)
			if cErr != nil {
				lineContent = "cannot get the code line"
			}
		}

		log.Error(
			fmt.Sprintf(
				"error with transaction from address: %s (%s) to address: %s (%s) {opcode: %s (%s) pc: %d code: %s Error msg: %s}\n",
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
