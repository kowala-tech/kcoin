package oracle

import (
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

var (
	errNoAddress = errors.New("there isn't an address for the provided chain ID")
)

var mapOracleMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Manager interface {
	Price() (*big.Int, error)
	GetOracleCount() (*big.Int, error)
}

// Binding returns a binding to the current oracle mgr
func Binding(contractBackend bind.ContractBackend, chainID *big.Int) (*OracleMgrSession, error) {
	addr, ok := mapOracleMgrToAddr[chainID.Uint64()]
	if !ok {
		return nil, errNoAddress
	}

	mgr, err := NewOracleMgr(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &OracleMgrSession{
		Contract: mgr,
		CallOpts: bind.CallOpts{},
	}, nil
}
