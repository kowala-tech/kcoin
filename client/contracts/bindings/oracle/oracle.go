package oracle

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite --libraries NameHash:0x66DA4aC1767B04B0d99bC94CCaD6EEF8dA63Ae96 -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

var mapOracleMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Manager interface {
	GetOracleCount() (*big.Int, error)
}

type manager struct {
	*OracleMgrSession
}

// @TODO(rgeraldes) - temporary method
func (mgr *manager) Domain() string {
	return ""
}

// Bind returns a binding to the current oracle mgr
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, ok := mapOracleMgrToAddr[chainID.Uint64()]
	if !ok {
		return nil, bindings.ErrNoAddress
	}

	mgr, err := NewOracleMgr(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &manager{
		&OracleMgrSession{
			Contract: mgr,
			CallOpts: bind.CallOpts{},
		},
	}, nil
}
