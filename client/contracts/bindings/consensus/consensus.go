package consensus

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/token"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite --libraries NameHash:0x3b058a1a62E59D185618f64BeBBAF3C52bf099E0 -o build zos-lib/=../../truffle/node_modules/zos-lib/ github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/consensus/mgr/ValidatorMgr.sol
//go:generate ../../../build/bin/abigen -abi build/ValidatorMgr.abi -bin build/ValidatorMgr.bin -pkg consensus -type ValidatorMgr -out ./gen_manager.go
//go:generate solc --allow-paths ., --abi --bin --overwrite -o build zos-lib/=../../truffle/node_modules/zos-lib/ github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ ../../truffle/contracts/consensus/token/MiningToken.sol
//go:generate ../../../build/bin/abigen -abi build/MiningToken.abi -bin build/MiningToken.bin -pkg consensus -type MiningToken -out ./gen_mtoken.go

const (
	RegistrationHandler = "registerValidator(address,uint256)"
	DepositHandler      = "increaseDeposit(address,uint256)"
)

var DefaultData = []byte("not_zero")

type mUSD struct {
	*MiningToken
	chainID *big.Int
}

func NewMUSD(contractBackend bind.ContractBackend, chainID *big.Int) (*mUSD, error) {
	addr, err := kns.GetAddressFromDomain(
		params.KNSDomains[params.MiningTokenDomain].FullDomain(),
		contractBackend,
	)
	if err != nil {
		return nil, err
	}

	mtoken, err := NewMiningToken(addr, contractBackend)
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

func (tkn *mUSD) Mint(opts *accounts.TransactOpts, to common.Address, value *big.Int) (common.Hash, error) {
	tx, err := tkn.MiningToken.Mint(toBind(opts), to, value)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), err
}

func toBind(opts *accounts.TransactOpts) *bind.TransactOpts {
	bindOpts := &bind.TransactOpts{
		From:     opts.From,
		Nonce:    opts.Nonce,
		Value:    opts.Value,
		GasPrice: opts.GasPrice,
		Context:  opts.Context,
		Signer:   bind.SignerFn(opts.Signer),
	}
	if opts.GasLimit != nil {
		bindOpts.GasLimit = opts.GasLimit.Uint64()
	}
	return bindOpts
}

func (tkn *mUSD) BalanceOf(target common.Address) (*big.Int, error) {
	return tkn.MiningToken.BalanceOf(&bind.CallOpts{}, target)
}

func (tkn *mUSD) Name() (string, error) {
	return tkn.MiningToken.Name(&bind.CallOpts{})
}

func (tkn *mUSD) Cap() (*big.Int, error) {
	return tkn.MiningToken.Cap(&bind.CallOpts{})
}

func (tkn *mUSD) TotalSupply() (*big.Int, error) {
	return tkn.MiningToken.TotalSupply(&bind.CallOpts{})
}

func (tkn *mUSD) MintingFinished() (bool, error) {
	return tkn.MiningToken.MintingFinished(&bind.CallOpts{})
}

// Consensus is a gateway to the validators contracts on the network
type Consensus struct {
	manager         *ValidatorMgr
	managerAddr     common.Address
	mtoken          token.Token
	chainID         *big.Int
	contractBackend bind.ContractBackend

	mtokenAddr     common.Address
	initMint       sync.Once
	multiSigWallet *ownership.MultiSigWallet
	oracle         *oracle.OracleMgr
}

// Binding returns a binding to the current Consensus engine
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, err := kns.GetAddressFromDomain(
		params.KNSDomains[params.ValidatorMgrDomain].FullDomain(),
		contractBackend,
	)
	if err != nil {
		return nil, err
	}

	mTokenAddr, err := kns.GetAddressFromDomain(
		params.KNSDomains[params.MiningTokenDomain].FullDomain(),
		contractBackend,
	)
	if err != nil {
		return nil, err
	}

	manager, err := NewValidatorMgr(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	mUSD, err := NewMUSD(contractBackend, chainID)
	if err != nil {
		return nil, err
	}

	return &Consensus{
		manager:         manager,
		managerAddr:     addr,
		mtoken:          mUSD,
		chainID:         chainID,
		contractBackend: contractBackend,
		mtokenAddr:      mTokenAddr,
	}, nil
}

func (css *Consensus) Join(walletAccount accounts.WalletAccount, deposit *big.Int) (common.Hash, error) {
	log.Warn(fmt.Sprintf("Joining the network %v with a deposit %v. Account %q",
		css.chainID.String(), deposit.String(), walletAccount.Account().Address.String()))
	hash, err := css.mtoken.Transfer(walletAccount, css.managerAddr, deposit, []byte("not_zero"), RegistrationHandler)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to transact the deposit: %s", err)
	}

	return hash, nil
}

func (css *Consensus) Leave(walletAccount accounts.WalletAccount) (common.Hash, error) {
	log.Warn(fmt.Sprintf("Leaving the network %v. Account %q",
		css.chainID.String(), walletAccount.Account().Address.String()))
	tx, err := css.manager.DeregisterValidator(transactOpts(walletAccount, css.chainID))
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (css *Consensus) RedeemDeposits(walletAccount accounts.WalletAccount) (common.Hash, error) {
	log.Warn(fmt.Sprintf("Redeem deposit from the network %v. Account %q",
		css.chainID.String(), walletAccount.Account().Address.String()))
	tx, err := css.manager.ReleaseDeposits(transactOpts(walletAccount, css.chainID))
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (css *Consensus) ValidatorsChecksum() (types.VotersChecksum, error) {
	return css.manager.ValidatorsChecksum(&bind.CallOpts{})
}

func (css *Consensus) Validators() (types.Voters, error) {
	count, err := css.manager.GetValidatorCount(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	voters := make([]*types.Voter, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := css.manager.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		weight := big.NewInt(0)
		voters[i] = types.NewVoter(validator.Code, validator.Deposit, weight)
	}

	return types.NewVoters(voters)
}

func (css *Consensus) Deposits(addr common.Address) ([]*types.Deposit, error) {
	count, err := css.manager.GetDepositCount(&bind.CallOpts{From: addr})
	if err != nil {
		return nil, err
	}

	deposits := make([]*types.Deposit, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		deposit, err := css.manager.GetDepositAtIndex(&bind.CallOpts{From: addr}, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		deposits[i] = types.NewDeposit(deposit.Amount, deposit.AvailableAt.Int64())
	}

	return deposits, nil
}

func (css *Consensus) IsGenesisValidator(address common.Address) (bool, error) {
	return css.manager.IsGenesisValidator(&bind.CallOpts{}, address)
}

func (css *Consensus) IsValidator(address common.Address) (bool, error) {
	return css.manager.IsValidator(&bind.CallOpts{}, address)
}

func (css *Consensus) MinimumDeposit() (*big.Int, error) {
	return css.manager.GetMinimumDeposit(&bind.CallOpts{})
}

func (css *Consensus) GetValidatorCount() (*big.Int, error) {
	return css.manager.GetValidatorCount(&bind.CallOpts{})
}

func (css *Consensus) MaxValidators() (*big.Int, error) {
	return css.manager.MaxNumValidators(&bind.CallOpts{})
}

func (css *Consensus) Token() token.Token {
	return css.mtoken
}

//Minter interface implementation

func (css *Consensus) MintInit() error {
	var err error
	css.initMint.Do(func() {
		if css.multiSigWallet == nil {
			addr := bindings.MultiSigWalletAddr

			var multisig *ownership.MultiSigWallet
			multisig, err = ownership.NewMultiSigWallet(addr, css.contractBackend)
			if err != nil {
				return
			}

			css.multiSigWallet = multisig
		}

		if css.oracle == nil {
			addr, errKns := kns.GetAddressFromDomain(
				params.KNSDomains[params.OracleMgrDomain].FullDomain(),
				css.contractBackend,
			)
			if err != nil {
				err = errKns
				return
			}

			var oracleMgr *oracle.OracleMgr
			oracleMgr, err = oracle.NewOracleMgr(addr, css.contractBackend)
			if err != nil {
				return
			}

			css.oracle = oracleMgr
		}
	})

	return err
}

func (css *Consensus) MultiSigWalletContract() *ownership.MultiSigWallet {
	return css.multiSigWallet
}

func (css *Consensus) Mint(opts *accounts.TransactOpts, to common.Address, value *big.Int) (common.Hash, error) {
	if err := css.MintInit(); err != nil {
		return common.Hash{}, err
	}

	tokenABI, err := abi.JSON(strings.NewReader(MiningTokenABI))
	if err != nil {
		return common.Hash{}, err
	}

	mintParams, err := tokenABI.Pack("mint", to, value)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := css.multiSigWallet.SubmitTransaction(toBind(opts), css.mtokenAddr, common.Big0, mintParams)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), err
}

func (css *Consensus) Confirm(opts *accounts.TransactOpts, transactionID *big.Int) (common.Hash, error) {
	if err := css.MintInit(); err != nil {
		return common.Hash{}, err
	}

	tx, err := css.multiSigWallet.ConfirmTransaction(toBind(opts), transactionID)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), err
}

func (css *Consensus) IncreaseDeposit(walletAccount accounts.WalletAccount, deposit *big.Int) (common.Hash, error) {
	log.Warn(fmt.Sprintf("Increasing the current deposit, with a new deposit of %v. Account %q", deposit.String(), walletAccount.Account().Address.String()))
	hash, err := css.mtoken.Transfer(walletAccount, css.managerAddr, deposit, []byte("not_zero"), DepositHandler)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to transact the new deposit: %s", err)
	}

	return hash, nil
}

// @TODO(rgeraldes) - temporary method
func (css *Consensus) Domain() string {
	return ""
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
