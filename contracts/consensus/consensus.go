package consensus

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/token"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/params"
)

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/contracts/=/usr/local/include/solidity/ contracts/ValidatorManager.sol
//go:generate abigen -abi build/ValidatorManager.abi -bin build/ValidatorManager.bin -pkg consensus -type ValidatorManager -out ./gen_manager.go

// RegistrationHandler represesnts the manager handler responsible for the registration
// of a user in the consensus. For token trasnfers (ERC233 model) we must specify the target method.
const RegistrationHandler = "RegisterValidator()"

// MapChainIDToAddr maps the contract address (const) per network
var MapChainIDToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x80eDa603028fe504B57D14d947c8087c1798D800"),
}

// ValidatorsChecksum lets a validator know if there are changes in the validator set
type ValidatorsChecksum [32]byte

// Consensus is a gateway to validators contracts on the network
type Consensus interface {
	Join(walletAccount accounts.WalletAccount, amount uint64) error
	Leave(walletAccount accounts.WalletAccount) error
	RedeemDeposits(walletAccount accounts.WalletAccount) error
	ValidatorsChecksum() (ValidatorsChecksum, error)
	Validators() (types.Voters, error)
	Deposits(address common.Address) ([]*types.Deposit, error)
	IsGenesisValidator(address common.Address) (bool, error)
	IsValidator(address common.Address) (bool, error)
	MinimumDeposit() (uint64, error)
}

type consensus struct {
	manager     *ValidatorManager
	managerAddr common.Address
	account     *token.MUSD
	chainID     *big.Int
}

// Instance returnsan instance of the current consensus engine
func Instance(contractBackend bind.ContractBackend, chainID *big.Int) (*consensus, error) {
	addr := MapChainIDToAddr[chainID.Uint64()]

	manager, err := NewValidatorManager(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	account, err := token.Instance(contractBackend, chainID)
	if err != nil {
		return nil, err
	}

	return &consensus{
		manager:     manager,
		managerAddr: addr,
		account:     account,
		chainID:     chainID,
	}, nil
}

func (consensus *consensus) Join(walletAccount accounts.WalletAccount, amount uint64) error {
	_, err := consensus.account.Transfer(consensus.transactOpts(walletAccount), consensus.managerAddr, new(big.Int).SetUint64(amount), []byte(RegistrationHandler), "tokenReceiver")
	if err != nil {
		return fmt.Errorf("failed to transact the deposit: %s", err)
	}

	return nil
}

func (consensus *consensus) Leave(walletAccount accounts.WalletAccount) error {
	_, err := consensus.manager.DeregisterValidator(consensus.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (consensus *consensus) RedeemDeposits(walletAccount accounts.WalletAccount) error {
	_, err := consensus.manager.ReleaseDeposits(consensus.transactOpts(walletAccount))
	if err != nil {
		return err
	}

	return nil
}

func (consensus *consensus) ValidatorsChecksum() (ValidatorsChecksum, error) {
	return consensus.manager.ValidatorsChecksum(&bind.CallOpts{})
}

func (consensus *consensus) Validators() (types.Voters, error) {
	count, err := consensus.manager.GetValidatorCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	voters := make([]*types.Voter, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := consensus.manager.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		weight := big.NewInt(0)
		voters[i] = types.NewVoter(validator.Code, validator.Deposit.Uint64(), weight)
	}

	return types.NewVoters(voters)
}

func (consensus *consensus) Deposits(addr common.Address) ([]*types.Deposit, error) {
	count, err := consensus.manager.GetDepositCount(&bind.CallOpts{From: addr})
	if err != nil {
		return nil, err
	}

	deposits := make([]*types.Deposit, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		deposit, err := consensus.manager.GetDepositAtIndex(&bind.CallOpts{From: addr}, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		deposits[i] = types.NewDeposit(deposit.Amount.Uint64(), deposit.AvailableAt.Int64())
	}

	return deposits, nil
}

func (consensus *consensus) IsGenesisValidator(address common.Address) (bool, error) {
	return consensus.manager.IsGenesisValidator(&bind.CallOpts{}, address)
}

func (consensus *consensus) IsValidator(address common.Address) (bool, error) {
	return consensus.manager.IsValidator(&bind.CallOpts{}, address)
}

func (consensus *consensus) transactOpts(walletAccount accounts.WalletAccount) *bind.TransactOpts {
	signerAddress := walletAccount.Account().Address
	opts := &bind.TransactOpts{
		From: signerAddress,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != signerAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			return walletAccount.SignTx(walletAccount.Account(), tx, consensus.chainID)
		},
	}

	return opts
}

func (consensus *consensus) MinimumDeposit() (uint64, error) {
	rawMinDeposit, err := consensus.manager.GetMinimumDeposit(&bind.CallOpts{})
	return rawMinDeposit.Uint64(), err
}

func (consensus *consensus) transactDepositOpts(walletAccount accounts.WalletAccount, amount uint64) *bind.TransactOpts {
	ops := consensus.transactOpts(walletAccount)
	var deposit big.Int
	ops.Value = deposit.SetUint64(amount)
	return ops
}
