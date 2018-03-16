package validator

import (
	"errors"
	"fmt"
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core/types"
	"math/big"
)

type ValidatorsChecksum [32]byte

// ValidationNetwork is a gateway to validators contracts on the network
type ValidationNetwork interface {
	Join(walletAccount accounts.WalletAccount, amount uint64) error
	Withdraw(walletAccount accounts.WalletAccount) error
	ValidatorsChecksum() (ValidatorsChecksum, error)
	Validators() (*types.ValidatorSet, error)
	IsGenesisVoter(address common.Address) (bool, error)
	IsVoter(address common.Address) (bool, error)
}

func NewValidationNetwork(networkContract *network.NetworkContract, chainID *big.Int) *validationNetwork {
	return &validationNetwork{
		NetworkContract: networkContract,
		chainID:         chainID,
	}
}

type validationNetwork struct {
	*network.NetworkContract
	chainID *big.Int
}

func (network *validationNetwork) Join(walletAccount accounts.WalletAccount, amount uint64) error {
	availability, err := network.Availability(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if !availability {
		return fmt.Errorf("there are no positions available at the moment")
	}

	err = network.deposit(walletAccount, amount)
	if err != nil {
		return err
	}

	return nil
}

func (network *validationNetwork) deposit(walletAccount accounts.WalletAccount, amount uint64) error {
	min, err := network.MinDeposit(&bind.CallOpts{})
	if err != nil {
		return err
	}

	var deposit big.Int
	if min.Cmp(deposit.SetUint64(amount)) > 0 {
		return fmt.Errorf("current deposit - %d - is not enough. The minimum required is %d", amount, min)
	}

	_, err = network.Deposit(network.transactDepositOpts(walletAccount, amount))
	if err != nil {
		return fmt.Errorf("failed to transact the deposit: %s", err)
	}

	return nil
}

func (network *validationNetwork) Withdraw(walletAccount accounts.WalletAccount) error {
	_, err := network.NetworkContract.Withdraw(network.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (network *validationNetwork) ValidatorsChecksum() (ValidatorsChecksum, error) {
	return network.NetworkContract.VotersChecksum(&bind.CallOpts{})
}

func (network *validationNetwork) Validators() (*types.ValidatorSet, error) {
	count, err := network.GetVoterCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	validators := make([]*types.Validator, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := network.GetVoterAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		weight := big.NewInt(0)
		validators[i] = types.NewValidator(validator.Addr, validator.Deposit.Uint64(), weight)
	}

	return types.NewValidatorSet(validators), nil
}

func (network *validationNetwork) IsGenesisVoter(address common.Address) (bool, error) {
	return network.NetworkContract.IsGenesisVoter(&bind.CallOpts{}, address)
}

func (network *validationNetwork) IsVoter(address common.Address) (bool, error) {
	return network.NetworkContract.IsVoter(&bind.CallOpts{}, address)
}

func (network *validationNetwork) transactOpts(walletAccount accounts.WalletAccount) *bind.TransactOpts {
	signerAddress := walletAccount.Account().Address
	opts := &bind.TransactOpts{
		From: signerAddress,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != signerAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			return walletAccount.SignTx(walletAccount.Account(), tx, network.chainID)
		},
	}

	return opts
}

func (network *validationNetwork) transactDepositOpts(walletAccount accounts.WalletAccount, amount uint64) *bind.TransactOpts {
	ops := network.transactOpts(walletAccount)
	var deposit big.Int
	ops.Value = deposit.SetUint64(amount)
	return ops
}
