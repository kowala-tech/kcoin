package consensus

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/token"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/consensus/mgr/ValidatorMgr.sol
//go:generate ../../../build/bin/abigen -abi build/ValidatorMgr.abi -bin build/ValidatorMgr.bin -pkg consensus -type ValidatorMgr -out ./gen_manager.go
//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ ../../truffle/contracts/consensus/token/MiningToken.sol
//go:generate ../../../build/bin/abigen -abi build/MiningToken.abi -bin build/MiningToken.bin -pkg consensus -type MiningToken -out ./gen_mtoken.go

const RegistrationHandler = "registerValidator(address,uint256,bytes)"

var (
	DefaultData = []byte("not_zero")
)

var mapValidatorMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x80eDa603028fe504B57D14d947c8087c1798D800"),
}

var mapMiningTokenToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x6f04441A6eD440Cc139a4E33402b438C27E97F4B"),
}

// ValidatorsChecksum lets a validator know if there are changes in the validator set
type ValidatorsChecksum [32]byte

// Consensus is a gateway to the validators contracts on the network
type Consensus interface {
	Join(walletAccount accounts.WalletAccount, amount *big.Int) error
	Leave(walletAccount accounts.WalletAccount) error
	RedeemDeposits(walletAccount accounts.WalletAccount) error
	ValidatorsChecksum() (ValidatorsChecksum, error)
	Validators() (types.Voters, error)
	GetValidatorCount() (*big.Int, error)
	MaxValidators() (*big.Int, error)
	Deposits(address common.Address) ([]*types.Deposit, error)
	IsGenesisValidator(address common.Address) (bool, error)
	IsValidator(address common.Address) (bool, error)
	MinimumDeposit() (*big.Int, error)
	Token() token.Token
}

type mUSD struct {
	*MiningToken
	chainID *big.Int
}

func NewMUSD(contractBackend bind.ContractBackend, chainID *big.Int) (*mUSD, error) {
	mtoken, err := NewMiningToken(mapMiningTokenToAddr[chainID.Uint64()], contractBackend)
	if err != nil {
		return nil, err
	}
	return &mUSD{MiningToken: mtoken, chainID: chainID}, nil
}

func (tkn *mUSD) Transfer(walletAccount accounts.WalletAccount, to common.Address, value *big.Int, data []byte, customFallback string) (common.Hash, error) {
	tx, err := tkn.MiningToken.Transfer(transactOpts(walletAccount, tkn.chainID), to, value, data, customFallback)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), err
}

func (tkn *mUSD) BalanceOf(target common.Address) (*big.Int, error) {
	return tkn.MiningToken.BalanceOf(&bind.CallOpts{}, target)
}

func (tkn *mUSD) Name() (string, error) {
	return tkn.MiningToken.Name(&bind.CallOpts{})
}

type consensus struct {
	manager     *ValidatorMgr
	managerAddr common.Address
	mtoken      token.Token
	chainID     *big.Int
	addr        common.Address
}

// Binding returns a binding to the current consensus engine
func Binding(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, ok := mapValidatorMgrToAddr[chainID.Uint64()]
	if !ok {
		return nil, bindings.ErrNoAddress
	}

	manager, err := NewValidatorMgr(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	mUSD, err := NewMUSD(contractBackend, chainID)
	if err != nil {
		return nil, err
	}

	return &consensus{
		manager:     manager,
		managerAddr: addr,
		mtoken:      mUSD,
		chainID:     chainID,
		addr:        addr,
	}, nil
}

func (cs *consensus) Join(walletAccount accounts.WalletAccount, deposit *big.Int) error {
	log.Warn(fmt.Sprintf("Joining the network %v with a deposit %v. Account %q",
		cs.chainID.String(), deposit.String(), walletAccount.Account().Address.String()))
	_, err := cs.mtoken.Transfer(walletAccount, cs.managerAddr, deposit, []byte("not_zero"), RegistrationHandler)
	if err != nil {
		return fmt.Errorf("failed to transact the deposit: %s", err)
	}

	return nil
}

func (cs *consensus) Leave(walletAccount accounts.WalletAccount) error {
	log.Warn(fmt.Sprintf("Leaving the network %v. Account %q",
		cs.chainID.String(), walletAccount.Account().Address.String()))
	_, err := cs.manager.DeregisterValidator(transactOpts(walletAccount, cs.chainID))
	if err != nil {
		return err
	}

	return nil
}

func (cs *consensus) RedeemDeposits(walletAccount accounts.WalletAccount) error {
	log.Warn(fmt.Sprintf("Redeem deposit from the network %v. Account %q",
		cs.chainID.String(), walletAccount.Account().Address.String()))
	_, err := cs.manager.ReleaseDeposits(transactOpts(walletAccount, cs.chainID))
	if err != nil {
		return err
	}

	return nil
}

func (cs *consensus) ValidatorsChecksum() (ValidatorsChecksum, error) {
	return cs.manager.ValidatorsChecksum(&bind.CallOpts{})
}

func (cs *consensus) Validators() (types.Voters, error) {
	count, err := cs.manager.GetValidatorCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	voters := make([]*types.Voter, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := cs.manager.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		weight := big.NewInt(0)
		voters[i] = types.NewVoter(validator.Code, validator.Deposit, weight)
	}

	return types.NewVoters(voters)
}

func (cs *consensus) Deposits(addr common.Address) ([]*types.Deposit, error) {
	count, err := cs.manager.GetDepositCount(&bind.CallOpts{From: addr})
	if err != nil {
		return nil, err
	}

	deposits := make([]*types.Deposit, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		deposit, err := cs.manager.GetDepositAtIndex(&bind.CallOpts{From: addr}, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		deposits[i] = types.NewDeposit(deposit.Amount, deposit.AvailableAt.Int64())
	}

	return deposits, nil
}

func (cs *consensus) IsGenesisValidator(address common.Address) (bool, error) {
	return cs.manager.IsGenesisValidator(&bind.CallOpts{}, address)
}

func (cs *consensus) IsValidator(address common.Address) (bool, error) {
	return cs.manager.IsValidator(&bind.CallOpts{}, address)
}

func (cs *consensus) MinimumDeposit() (*big.Int, error) {
	return cs.manager.GetMinimumDeposit(&bind.CallOpts{})
}

func (cs *consensus) GetValidatorCount() (*big.Int, error) {
	return cs.manager.GetValidatorCount(&bind.CallOpts{})
}

func (cs *consensus) MaxValidators() (*big.Int, error) {
	return cs.manager.MaxNumValidators(&bind.CallOpts{})
}

func (cs *consensus) Token() token.Token {
	return cs.mtoken
}

func (cs *consensus) Address() common.Address {
	return cs.addr
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
