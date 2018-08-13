package stability

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/stability/Stability.sol
//go:generate ../../../build/bin/abigen -abi build/Stability.abi -bin build/Stability.bin -pkg stability -type Stability -out ./gen_stability.go

var mapStabilityContractToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type StabilityContract interface {
}

type stability struct {
	*StabilitySession
}

// @TODO(rgeraldes) - temporary method
func (sc *stability) Domain() string {
	return ""
}

// Bind returns a binding to the current stability contract
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, ok := mapStabilityContractToAddr[chainID.Uint64()]
	if !ok {
		return nil, bindings.ErrNoAddress
	}

	contract, err := NewStability(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &stability{
		&StabilitySession{
			Contract: contract,
			CallOpts: bind.CallOpts{},
		},
	}, nil
}
