package oracle

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

var mapOracleMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Manager interface {
	RegisterOracle(user *ecdsa.PrivateKey) (*types.Transaction, error)
	SubmitPrice(user *ecdsa.PrivateKey, price *big.Int) (*types.Transaction, error)
	Submissions() ([]common.Address, error)
	AveragePrice() (*big.Int, error)
	GetOracleCount() (*big.Int, error)
	GetOracleAtIndex(*big.Int) (common.Address, *big.Int, bool, error)
	MaxNumOracles() (*big.Int, error)
	Address() common.Address
}

type manager struct {
	*OracleMgrSession
	addr common.Address
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
		OracleMgrSession: &OracleMgrSession{
			Contract: mgr,
			CallOpts: bind.CallOpts{},
		},
		addr: addr,
	}, nil
}

func (mgr *manager) Address() common.Address {
	return mgr.addr
}

func (mgr *manager) Submissions() ([]common.Address, error) {
	numSubmissions, err := mgr.GetNumSubmissions()
	if err != nil {
		return nil, err
	}

	submissions := make([]common.Address, numSubmissions.Uint64())
	for i := int64(0); i < numSubmissions.Int64(); i++ {
		submission, err := mgr.GetSubmissionAtIndex(big.NewInt(i))
		if err != nil {
			return nil, err
		}
		submissions[i] = submission
	}

	return submissions, nil
}

func (mgr *manager) RegisterOracle(user *ecdsa.PrivateKey) (*types.Transaction, error) {
	return mgr.OracleMgrSession.Contract.RegisterOracle(bind.NewKeyedTransactor(user))
}

func (mgr *manager) SubmitPrice(user *ecdsa.PrivateKey, price *big.Int) (*types.Transaction, error) {
	return mgr.OracleMgrSession.Contract.SubmitPrice(bind.NewKeyedTransactor(user), price)
}

func (mgr *manager) GetOracleAtIndex(index *big.Int) (common.Address, *big.Int, bool, error) {
	values, err := mgr.OracleMgrSession.GetOracleAtIndex(index)
	return values.Code, values.Deposit, values.Price, err
}
