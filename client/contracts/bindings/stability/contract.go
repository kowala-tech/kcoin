package stability

import (
	"github.com/kowala-tech/kcoin/client/log"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common/kns"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/stability/Stability.sol
//go:generate ../../../build/bin/abigen -abi build/Stability.abi -bin build/Stability.bin -pkg stability -type Stability -out ./gen_stability.go

type StabilityContract struct {
	*StabilitySession
}

// @TODO(rgeraldes) - temporary method
func (sc *StabilityContract) Domain() string {
	return ""
}

// Bind returns a binding to the current stability contract
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, err := kns.GetAddressFromDomain(
		params.KNSDomains[params.StabilityDomain].FullDomain(),
		contractBackend,
	)
	if err != nil {
		log.Error("can't find Stability contract for given Network", "chainID", chainID.String())
		return nil, bindings.ErrNoAddress
	}

	contract, err := NewStability(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &StabilityContract{
		&StabilitySession{
			Contract: contract,
			CallOpts: bind.CallOpts{},
		},
	}, nil
}
