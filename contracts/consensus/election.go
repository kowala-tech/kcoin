package consensus

//go:generate solc --abi --bin --overwrite -o build contract/Election.sol
//go:generate abigen -abi build/Election.abi -bin build/Election.bin -pkg contract -type ElectionContract -out contract/election.go

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/contracts/consensus/contract"
	"github.com/tendermint/tendermint/types"
)

type ValidatorsChecksum [32]byte

type Election interface {
	Join(walletAccount accounts.WalletAccount, amount uint64) error
	Leave(walletAccount accounts.WalletAccount) error
	RedeemFunds(walletAccount accounts.WalletAccount) error
	ValidatorsChecksum() (ValidatorsChecksum, error)
	Validators() (*types.ValidatorSet, error)
	Deposits() error
	IsGenesisVoter(address common.Address) (bool, error)
	IsVoter(address common.Address) (bool, error)
}

func NewElection(networkContract *network.NetworkContract, chainID *big.Int) *validationNetwork {
	return &validationNetwork{
		NetworkContract: networkContract,
		chainID:         chainID,
	}
}

type election struct {
	*contract.ElectionContract
	chainID *big.Int
}

func (election *election) Join(walletAccount accounts.WalletAccount, amount uint64) error {
	min, err := network.MinDeposit(&bind.CallOpts{})
	if err != nil {
		return err
	}

	var deposit big.Int
	if min.Cmp(deposit.SetUint64(amount)) > 0 {
		return fmt.Errorf("current deposit - %d - is not enough. The minimum required is %d", amount, min)
	}

	_, err = election.Join(election.transactDepositOpts(walletAccount, amount))
	if err != nil {
		return fmt.Errorf("failed to transact the deposit: %s", err)
	}

	return nil
}

func (election *election) Leave(walletAccount accounts.WalletAccount) error {
	_, err := election.Leave(election.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (election *election) RedeemFunds(walletAccount accounts.WalletAccount) error {
	_, err := election.RedeemFunds(election.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (election *election) ValidatorsChecksum() (ValidatorsChecksum, error) {
	return election.VotersChecksum(&bind.CallOpts{})
}

func (election *election) Validators() (*types.ValidatorSet, error) {
	count, err := election.GetValidatorCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	validators := make([]*types.Validator, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := election.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		weight := big.NewInt(0)
		validators[i] = types.NewValidator(validator.Addr, validator.Deposit.Uint64(), weight)
	}

	return types.NewValidatorSet(validators), nil
}

func (election *election) Deposits() (*types.ValidatorSet, error) {}

func (election *election) IsGenesisValidator(address common.Address) (bool, error) {
	return election.IsGenesisValidator(&bind.CallOpts{}, address)
}

func (election *election) IsValidator(address common.Address) (bool, error) {
	return network.NetworkContract.IsValidator(&bind.CallOpts{}, address)
}

func (election *election) transactOpts(walletAccount accounts.WalletAccount) *bind.TransactOpts {
	signerAddress := walletAccount.Account().Address
	opts := &bind.TransactOpts{
		From: signerAddress,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != signerAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			return walletAccount.SignTx(walletAccount.Account(), tx, election.chainID)
		},
	}

	return opts
}

func (election *election) transactDepositOpts(walletAccount accounts.WalletAccount, amount uint64) *bind.TransactOpts {
	ops := election.transactOpts(walletAccount)
	var deposit big.Int
	ops.Value = deposit.SetUint64(amount)
	return ops
}
