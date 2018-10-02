package mining

import (
	"bytes"
	"errors"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/mining/validator"
)

// PrivateValidatorAPI provides private RPC methods to control the validator.
// These methods can be abused by external users and must be considered insecure for use by untrusted users.
type PrivateValidatorAPI struct {
	service *Service
}

// NewPrivateValidatorAPI create a new RPC service which controls the validator of this node.
func NewPrivateValidatorAPI(service *Service) *PrivateValidatorAPI {
	return &PrivateValidatorAPI{service: service}
}

// Start the validator.
func (api *PrivateValidatorAPI) Start(deposit *hexutil.Big) error {
	// Start the validator and return
	if !api.service.IsValidating() {
		// Propagate the initial price point to the transaction pool
		api.service.lock.RLock()
		price := api.service.gasPrice
		api.service.lock.RUnlock()

		bigint := deposit.ToInt()
		if bigint.Cmp(big.NewInt(0)) != 0 {
			err := api.service.SetDeposit(bigint)
			if err != nil && err != validator.ErrIsNotRunning {
				return err
			}
		}

		api.service.txPool.SetGasPrice(price)
		return api.service.StartValidating()
	}
	return nil
}

// Stop the validator
func (api *PrivateValidatorAPI) Stop() bool {
	api.service.StopValidating()
	return true
}

// SetExtra sets the extra data string that is included when this validator proposes a block.
func (api *PrivateValidatorAPI) SetExtra(extra string) (bool, error) {
	if err := api.service.Validator().SetExtra([]byte(extra)); err != nil {
		return false, err
	}
	return true, nil
}

// SetGasPrice sets the minimum accepted gas price for the validator.
func (api *PrivateValidatorAPI) SetGasPrice(gasPrice hexutil.Big) bool {
	api.service.lock.Lock()
	api.service.gasPrice = (*big.Int)(&gasPrice)
	api.service.lock.Unlock()

	api.service.txPool.SetGasPrice((*big.Int)(&gasPrice))
	return true
}

// SetCoinbase sets the coinbase of the validator
func (api *PrivateValidatorAPI) SetCoinbase(coinbase common.Address) bool {
	api.service.SetCoinbase(coinbase)
	return true
}

// GetMinimumDeposit gets the minimum deposit required to take a slot as a validator
func (api *PrivateValidatorAPI) GetMinimumDeposit() (*big.Int, error) {
	return api.service.GetMinimumDeposit()
}

// GetDepositsResult is the result of a validator_getDeposits API call.
type GetDepositsResult struct {
	Deposits []depositEntry `json:"deposits"`
}

type depositEntry struct {
	Amount      *big.Int `json:"value"`
	AvailableAt string   `json:",omitempty"`
}

// GetDeposits returns the validator deposits
func (api *PrivateValidatorAPI) GetDeposits(address *common.Address) (GetDepositsResult, error) {
	rawDeposits, err := api.service.Validator().Deposits(address)
	if err != nil {
		return GetDepositsResult{}, err
	}

	return depositsToResponse(rawDeposits), nil
}

func depositsToResponse(rawDeposits []*types.Deposit) GetDepositsResult {
	deposits := make([]depositEntry, len(rawDeposits))

	for i, deposit := range rawDeposits {
		// @NOTE (rgeraldes) - zero values are not shown for this field
		var availableAt string

		if deposit.AvailableAtTimeUnix() != 0 {
			availableAt = time.Unix(deposit.AvailableAtTimeUnix(), 0).String()
		}

		deposits[i] = depositEntry{
			Amount:      deposit.Amount(),
			AvailableAt: availableAt,
		}
	}

	return GetDepositsResult{Deposits: deposits}
}

// IsValidating returns the validator is currently validating
func (api *PrivateValidatorAPI) IsValidating() bool {
	return api.service.IsValidating()
}

// IsValidating returns the validator is currently running
func (api *PrivateValidatorAPI) IsRunning() bool {
	return api.service.IsRunning()
}

// RedeemDeposits requests a transfer of the unlocked deposits back
// to the validator account
func (api *PrivateValidatorAPI) RedeemDeposits() error {
	return api.service.Validator().RedeemDeposits()
}

// Coinbase is the address that consensus rewards will be send to
func (api *PrivateValidatorAPI) Coinbase() (common.Address, error) {
	return api.service.Coinbase()
}

// TransferArgs represents the arguments to transfer tokens.
type TransferArgs struct {
	From           common.Address  `json:"from"`
	To             *common.Address `json:"to"`
	Value          *hexutil.Big    `json:"value"`
	Data           hexutil.Bytes   `json:"data"`
	CustomFallback string          `json:"fallback"`
}

// PublicTokenAPI exposes a collection of methods related to tokens
type PublicTokenAPI struct {
	accountMgr *accounts.Manager
	consensus  *consensus.Consensus
	chainID    *big.Int
}

func NewPublicTokenAPI(accountMgr *accounts.Manager, c *consensus.Consensus, chainID *big.Int) *PublicTokenAPI {
	return &PublicTokenAPI{
		accountMgr: accountMgr,
		consensus:  c,
		chainID:    chainID,
	}
}

func (api *PublicTokenAPI) GetBalance(target common.Address) (*big.Int, error) {
	return api.consensus.Token().BalanceOf(target)
}

func (api *PublicTokenAPI) Transfer(args TransferArgs) (common.Hash, error) {
	_, walletAccount, err := api.getWallet(args.From)
	if err != nil {
		return common.Hash{}, err
	}

	if args.Value == nil {
		args.Value = new(hexutil.Big)
	}

	return api.consensus.Token().Transfer(walletAccount, *args.To, (*big.Int)(args.Value), args.Data, args.CustomFallback)
}

func (api *PublicTokenAPI) Mint(from, to common.Address, value *hexutil.Big) (common.Hash, error) {
	if value == nil {
		return common.Hash{}, errors.New("a number of tokens should be specified")
	}

	account, walletAccount, err := api.getWallet(from)
	if err != nil {
		return common.Hash{}, err
	}

	tOpts := &accounts.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return walletAccount.SignTx(*account, tx, api.chainID)
		},
	}

	return api.consensus.Mint(tOpts, to, value.ToInt())
}

func (api *PublicTokenAPI) Confirm(from common.Address, transactionID *hexutil.Big) (common.Hash, error) {
	account, walletAccount, err := api.getWallet(from)
	if err != nil {
		return common.Hash{}, err
	}

	tOpts := &accounts.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return walletAccount.SignTx(*account, tx, api.chainID)
		},
	}

	return api.consensus.Confirm(tOpts, transactionID.ToInt())
}

func (api *PublicTokenAPI) Cap() (*big.Int, error) {
	return api.consensus.Token().Cap()
}

func (api *PublicTokenAPI) TotalSupply() (*big.Int, error) {
	return api.consensus.Token().TotalSupply()
}

func (api *PublicTokenAPI) MintingFinished() (bool, error) {
	return api.consensus.Token().MintingFinished()
}

type PendingMintTransaction struct {
	Id        *big.Int       `json:"id"`
	To        common.Address `json:"to",omitempty`
	Amount    *big.Int       `json:"amount",omitempty`
	Confirmed bool           `json:"confirmed"`
}

type PendingMintTransactions []PendingMintTransaction

func (api *PublicTokenAPI) MintList() (ret PendingMintTransactions, err error) {

	if err := api.consensus.MintInit(); err != nil {
		return ret, err
	}

	multiSig := api.consensus.MultiSigWalletContract()

	if multiSig == nil {
		return ret, errors.New("can't get multi sig contract")
	}

	max, err := multiSig.GetTransactionCount(&bind.CallOpts{}, true, true)

	if err != nil {
		return ret, err
	}

	ids, err := multiSig.GetTransactionIds(&bind.CallOpts{}, big.NewInt(0), max, true, true)

	if err != nil {
		return ret, err
	}

	mintMethodId := crypto.Keccak256([]byte("mint(address,uint256)"))[:4]

	for _, id := range ids {
		output, err := multiSig.Transactions(&bind.CallOpts{}, id)

		if err != nil {
			return ret, err
		}

		if !bytes.Equal(output.Data[:4], mintMethodId) {
			continue
		}

		amount := new(big.Int)
		amount.SetBytes(output.Data[37:])

		ret = append(ret, PendingMintTransaction{
			Id:        id,
			To:        common.BytesToAddress(output.Data[4:36]),
			Amount:    amount,
			Confirmed: output.Executed,
		})
	}

	return ret, nil
}

func (api *PublicTokenAPI) getWallet(addr common.Address) (*accounts.Account, accounts.WalletAccount, error) {
	// Look up the wallet containing the requested signer
	for _, wallet := range api.accountMgr.Wallets() {
		for _, account := range wallet.Accounts() {
			if account.Address == addr {
				walletAccount, err := accounts.NewWalletAccount(wallet, account)
				if err != nil {
					return nil, nil, err
				}
				return &account, walletAccount, nil
			}
		}
	}
	return nil, nil, errors.New("account not found in any wallet")
}
