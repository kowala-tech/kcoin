package sysvars

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/sysvars/SystemVars.sol
//go:generate ../../../build/bin/abigen -abi build/SystemVars.abi -bin build/SystemVars.bin -pkg sysvars -type SystemVars -out ./gen_systemvars.go

var mapSystemVarsToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Vars struct {
	*SystemVarsSession
}

// @TODO(rgeraldes) - temporary method
func (v *Vars) Domain() string {
	return ""
}

// Bind returns a binding to the current stability contract
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, ok := mapSystemVarsToAddr[chainID.Uint64()]
	if !ok {
		return nil, bindings.ErrNoAddress
	}

	contract, err := NewSystemVars(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Vars{
		&SystemVarsSession{
			Contract: contract,
			CallOpts: bind.CallOpts{},
		},
	}, nil
}
