package oracle

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

var mapOracleMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

// Manager represents an oracle manager
type Manager interface {
	CurrentPrice() (*big.Int, error)
	PreviousPrice() (*big.Int, error)
	Submissions() ([]common.Address, error)
	GetOracleCount() (*big.Int, error)
}

type manager struct {
	*OracleMgr
	addr common.Address
}

// Binding returns a binding to the current oracle mgr
func Binding(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, ok := mapOracleMgrToAddr[chainID.Uint64()]
	if !ok {
		return nil, bindings.ErrNoAddress
	}

	mgr, err := NewOracleMgr(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &manager{
		OracleMgr: mgr,
		addr:      addr,
	}, nil
}

func (mgr *manager) CurrentPrice() (*big.Int, error) {
	return mgr.Price(&bind.CallOpts{Pending: true})
}

func (mgr *manager) PreviousPrice() (*big.Int, error) {
	return mgr.Price(&bind.CallOpts{Pending: false})
}

func (mgr *manager) Address() common.Address {
	return mgr.addr
}

func (mgr *manager) Submisisons() ([]common.Address, error) {
	return nil, nil
}
