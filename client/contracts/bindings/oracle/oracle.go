package oracle

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

var mapOracleMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Manager interface {
	RegisterOracle(walletAccount accounts.WalletAccount) (*types.Transaction, error)
	DeregisterOracle(walletAccount accounts.WalletAccount) (*types.Transaction, error)
	GetOracleCount() (*big.Int, error)
	IsOracle(identity common.Address) (bool, error)
	HasPriceFrom(oracle common.Address) (bool, error)
}

type manager struct {
	*OracleMgr
	chainID *big.Int
}

// Binding returns a binding to the current oracle mgr
func Binding(contractBackend bind.ContractBackend, chainID *big.Int) (*manager, error) {
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
		chainID:   chainID,
	}, nil
}

func (mgr *manager) RegisterOracle(walletAccount accounts.WalletAccount) (*types.Transaction, error) {
	log.Info(fmt.Sprintf("Joining the oracle pool %v. Account %q", mgr.chainID.String(), walletAccount.Account().Address.String()))
	return mgr.OracleMgr.RegisterOracle(transactOpts(walletAccount, mgr.chainID))
}

func (mgr *manager) DeregisterOracle(walletAccount accounts.WalletAccount) (*types.Transaction, error) {
	log.Info(fmt.Sprintf("Leaving the network %v. Account %q", mgr.chainID.String(), walletAccount.Account().Address.String()))
	return mgr.OracleMgr.DeregisterOracle(transactOpts(walletAccount, mgr.chainID))
}

func (mgr *manager) GetOracleCount() (*big.Int, error) {
	return mgr.OracleMgr.GetOracleCount(&bind.CallOpts{})
}

func (mgr *manager) IsOracle(identity common.Address) (bool, error) {
	return mgr.OracleMgr.IsOracle(&bind.CallOpts{}, identity)
}

func (mgr *manager) HasPriceFrom(oracle common.Address) (bool, error) {
	return mgr.OracleMgr.HasSubmittedPrice(&bind.CallOpts{}, oracle)
}

func transactOpts(walletAccount accounts.WalletAccount, chainID *big.Int) *bind.TransactOpts {
	signerAddress := walletAccount.Account().Address
	opts := &bind.TransactOpts{
		From: signerAddress,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != signerAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			return walletAccount.SignTx(walletAccount.Account(), tx, chainID)
		},
	}

	return opts
}
