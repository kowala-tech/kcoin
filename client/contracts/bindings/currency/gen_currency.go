// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package currency

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// CurrencyABI is the input ABI used to generate the binding from.
const CurrencyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"prevMintedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CurrencyBin is the compiled bytecode used for deploying new contracts.
const CurrencyBin = `608060405234801561001057600080fd5b5060d58061001f6000396000f3006080604052600436106048576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062815bd514604d578063b61650b2146075575b600080fd5b348015605857600080fd5b50605f609d565b6040518082815260200191505060405180910390f35b348015608057600080fd5b50608760a3565b6040518082815260200191505060405180910390f35b60015481565b600054815600a165627a7a723058201e0164562478864d5b7624d71a94e2494de08a804c807c88c2538eee41961ec90029`

// DeployCurrency deploys a new Kowala contract, binding an instance of Currency to it.
func DeployCurrency(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Currency, error) {
	parsed, err := abi.JSON(strings.NewReader(CurrencyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CurrencyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Currency{CurrencyCaller: CurrencyCaller{contract: contract}, CurrencyTransactor: CurrencyTransactor{contract: contract}, CurrencyFilterer: CurrencyFilterer{contract: contract}}, nil
}

// Currency is an auto generated Go binding around a Kowala contract.
type Currency struct {
	CurrencyCaller     // Read-only binding to the contract
	CurrencyTransactor // Write-only binding to the contract
	CurrencyFilterer   // Log filterer for contract events
}

// CurrencyCaller is an auto generated read-only Go binding around a Kowala contract.
type CurrencyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyTransactor is an auto generated write-only Go binding around a Kowala contract.
type CurrencyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type CurrencyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencySession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type CurrencySession struct {
	Contract     *Currency         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurrencyCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type CurrencyCallerSession struct {
	Contract *CurrencyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CurrencyTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type CurrencyTransactorSession struct {
	Contract     *CurrencyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CurrencyRaw is an auto generated low-level Go binding around a Kowala contract.
type CurrencyRaw struct {
	Contract *Currency // Generic contract binding to access the raw methods on
}

// CurrencyCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type CurrencyCallerRaw struct {
	Contract *CurrencyCaller // Generic read-only contract binding to access the raw methods on
}

// CurrencyTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type CurrencyTransactorRaw struct {
	Contract *CurrencyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurrency creates a new instance of Currency, bound to a specific deployed contract.
func NewCurrency(address common.Address, backend bind.ContractBackend) (*Currency, error) {
	contract, err := bindCurrency(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Currency{CurrencyCaller: CurrencyCaller{contract: contract}, CurrencyTransactor: CurrencyTransactor{contract: contract}, CurrencyFilterer: CurrencyFilterer{contract: contract}}, nil
}

// NewCurrencyCaller creates a new read-only instance of Currency, bound to a specific deployed contract.
func NewCurrencyCaller(address common.Address, caller bind.ContractCaller) (*CurrencyCaller, error) {
	contract, err := bindCurrency(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurrencyCaller{contract: contract}, nil
}

// NewCurrencyTransactor creates a new write-only instance of Currency, bound to a specific deployed contract.
func NewCurrencyTransactor(address common.Address, transactor bind.ContractTransactor) (*CurrencyTransactor, error) {
	contract, err := bindCurrency(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurrencyTransactor{contract: contract}, nil
}

// NewCurrencyFilterer creates a new log filterer instance of Currency, bound to a specific deployed contract.
func NewCurrencyFilterer(address common.Address, filterer bind.ContractFilterer) (*CurrencyFilterer, error) {
	contract, err := bindCurrency(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurrencyFilterer{contract: contract}, nil
}

// bindCurrency binds a generic wrapper to an already deployed contract.
func bindCurrency(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurrencyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Currency *CurrencyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Currency.Contract.CurrencyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Currency *CurrencyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currency.Contract.CurrencyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Currency *CurrencyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Currency.Contract.CurrencyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Currency *CurrencyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Currency.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Currency *CurrencyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currency.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Currency *CurrencyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Currency.Contract.contract.Transact(opts, method, params...)
}

// PrevMintedAmount is a free data retrieval call binding the contract method 0x00815bd5.
//
// Solidity: function prevMintedAmount() constant returns(uint256)
func (_Currency *CurrencyCaller) PrevMintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Currency.contract.Call(opts, out, "prevMintedAmount")
	return *ret0, err
}

// PrevMintedAmount is a free data retrieval call binding the contract method 0x00815bd5.
//
// Solidity: function prevMintedAmount() constant returns(uint256)
func (_Currency *CurrencySession) PrevMintedAmount() (*big.Int, error) {
	return _Currency.Contract.PrevMintedAmount(&_Currency.CallOpts)
}

// PrevMintedAmount is a free data retrieval call binding the contract method 0x00815bd5.
//
// Solidity: function prevMintedAmount() constant returns(uint256)
func (_Currency *CurrencyCallerSession) PrevMintedAmount() (*big.Int, error) {
	return _Currency.Contract.PrevMintedAmount(&_Currency.CallOpts)
}

// PrevSupply is a free data retrieval call binding the contract method 0xb61650b2.
//
// Solidity: function prevSupply() constant returns(uint256)
func (_Currency *CurrencyCaller) PrevSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Currency.contract.Call(opts, out, "prevSupply")
	return *ret0, err
}

// PrevSupply is a free data retrieval call binding the contract method 0xb61650b2.
//
// Solidity: function prevSupply() constant returns(uint256)
func (_Currency *CurrencySession) PrevSupply() (*big.Int, error) {
	return _Currency.Contract.PrevSupply(&_Currency.CallOpts)
}

// PrevSupply is a free data retrieval call binding the contract method 0xb61650b2.
//
// Solidity: function prevSupply() constant returns(uint256)
func (_Currency *CurrencyCallerSession) PrevSupply() (*big.Int, error) {
	return _Currency.Contract.PrevSupply(&_Currency.CallOpts)
}
