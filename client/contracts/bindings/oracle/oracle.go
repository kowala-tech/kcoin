package oracle

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common/kns"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite --libraries NameHash:0x66DA4aC1767B04B0d99bC94CCaD6EEF8dA63Ae96 -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

type Manager struct {
	*OracleMgrSession
}

// @TODO(rgeraldes) - temporary method
func (mgr *Manager) Domain() string {
	return ""
}

// Bind returns a binding to the current oracle mgr
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, err := kns.GetAddressFromDomain(
		params.KNSDomains[params.OracleMgrDomain].FullDomain(),
		contractBackend,
	)
	if err != nil {
		return nil, bindings.ErrNoAddress
	}

	mgr, err := NewOracleMgr(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Manager{
		&OracleMgrSession{
			Contract: mgr,
			CallOpts: bind.CallOpts{},
		},
	}, nil
}
