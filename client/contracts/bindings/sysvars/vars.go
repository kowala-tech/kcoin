package sysvars

import (
	"github.com/kowala-tech/kcoin/client/log"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common/kns"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/sysvars/SystemVars.sol
//go:generate ../../../build/bin/abigen -abi build/SystemVars.abi -bin build/SystemVars.bin -pkg sysvars -type SystemVars -out ./gen_systemvars.go

type Vars struct {
	*SystemVarsSession
}

// @TODO(rgeraldes) - temporary method
func (v *Vars) Domain() string {
	return ""
}

// Bind returns a binding to the current stability contract
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, err := kns.GetAddressFromDomain(
		params.KNSDomains[params.SystemVarsDomain].FullDomain(),
		contractBackend,
	)
	if err != nil {
		log.Error("can't find SystemVar for given Network", "chainID", chainID.String())
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
