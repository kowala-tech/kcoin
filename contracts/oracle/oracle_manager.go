package oracle

//go:generate solc --abi --bin --overwrite -o build contract/OracleManager.sol
//go:generate abigen -abi build/OracleManager.abi -bin build/OracleManager.bin -pkg contract -type OracleManagerContract -out contract/oracle_manager.go

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/oracle/contract"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/params"
)

// @TODO (rgeraldes) - include address of the testnet
var mapChainIDToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress(""),
}

// PriceFeed represents a price submitted by an oracle
type PriceFeed struct {
}

type PriceFeeds []*PriceFeed

type OracleManager interface {
	SubmitPrice(walletAccount accounts.WalletAccount, price uint64) error
	RegisterOracle(walletAccount accounts.WalletAccount, amount uint64) error
	DeregisterOracle(walletAccount accounts.WalletAccount) error
	RedeemDeposits(walletAccount accounts.WalletAccount) error
	Submissions() (PriceFeeds, error)
}

type oracleManager struct {
	*contract.OracleManagerContract
	chainID *big.Int
}

func NewOracleManager(contractBackend bind.ContractBackend, chainID *big.Int) (*oracleManager, error) {
	contract, err := contract.NewOracleManagerContract(mapChainIDToAddr[chainID.Uint64()], contractBackend)
	if err != nil {
		return nil, err
	}

	manager := &oracleManager{
		OracleManagerContract: contract,
		chainID:               chainID,
	}

	return manager, nil
}

func (manager *oracleManager) RegisterOracle(walletAccount accounts.WalletAccount, amount uint64) error {
	_, err := manager.OracleManagerContract.RegisterOracle(manager.transactDepositOpts(walletAccount, amount))
	if err != nil {
		return fmt.Errorf("failed to transact the deposit: %s", err)
	}

	return nil
}

func (manager *oracleManager) DeregisterOracle(walletAccount accounts.WalletAccount) error {
	_, err := manager.OracleManagerContract.DeregisterOracle(manager.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (manager *oracleManager) RedeemDeposits(walletAccount accounts.WalletAccount) error {
	_, err := manager.OracleManagerContract.RedeemDeposits(manager.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (manager *oracleManager) SubmitPrice(walletAccount accounts.WalletAccount, price *big.Int) error {
	_, err := manager.OracleManagerContract.SubmitPrice(manager.transactOpts(walletAccount), price)
	if err != nil {
		return err
	}

	return nil
}

func (manager *oracleManager) transactOpts(walletAccount accounts.WalletAccount) *bind.TransactOpts {
	signerAddress := walletAccount.Account().Address
	opts := &bind.TransactOpts{
		From: signerAddress,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != signerAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			return walletAccount.SignTx(walletAccount.Account(), tx, manager.chainID)
		},
	}

	return opts
}

func (manager *oracleManager) transactDepositOpts(walletAccount accounts.WalletAccount, amount uint64) *bind.TransactOpts {
	ops := manager.transactOpts(walletAccount)
	var deposit big.Int
	ops.Value = deposit.SetUint64(amount)
	return ops
}
