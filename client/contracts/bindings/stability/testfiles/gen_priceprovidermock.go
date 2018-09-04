// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testfiles

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// PriceProviderMockABI is the input ABI used to generate the binding from.
const PriceProviderMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PriceProviderMockBin is the compiled bytecode used for deploying new contracts.
const PriceProviderMockBin = `608060405234801561001057600080fd5b506040516020806100ea83398101806040528101908080519060200190929190505050806000819055505060a1806100496000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063a035b1fe146044575b600080fd5b348015604f57600080fd5b506056606c565b6040518082815260200191505060405180910390f35b600080549050905600a165627a7a72305820a94b85aa11c597b714b0f160bafbe045bf6fea7ae310f5df196ee77b18da44c30029`

// DeployPriceProviderMock deploys a new Kowala contract, binding an instance of PriceProviderMock to it.
func DeployPriceProviderMock(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int) (common.Address, *types.Transaction, *PriceProviderMock, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceProviderMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceProviderMockBin), backend, _price)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceProviderMock{PriceProviderMockCaller: PriceProviderMockCaller{contract: contract}, PriceProviderMockTransactor: PriceProviderMockTransactor{contract: contract}, PriceProviderMockFilterer: PriceProviderMockFilterer{contract: contract}}, nil
}

// PriceProviderMock is an auto generated Go binding around a Kowala contract.
type PriceProviderMock struct {
	PriceProviderMockCaller     // Read-only binding to the contract
	PriceProviderMockTransactor // Write-only binding to the contract
	PriceProviderMockFilterer   // Log filterer for contract events
}

// PriceProviderMockCaller is an auto generated read-only Go binding around a Kowala contract.
type PriceProviderMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceProviderMockTransactor is an auto generated write-only Go binding around a Kowala contract.
type PriceProviderMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceProviderMockFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type PriceProviderMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceProviderMockSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type PriceProviderMockSession struct {
	Contract     *PriceProviderMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PriceProviderMockCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type PriceProviderMockCallerSession struct {
	Contract *PriceProviderMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PriceProviderMockTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type PriceProviderMockTransactorSession struct {
	Contract     *PriceProviderMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PriceProviderMockRaw is an auto generated low-level Go binding around a Kowala contract.
type PriceProviderMockRaw struct {
	Contract *PriceProviderMock // Generic contract binding to access the raw methods on
}

// PriceProviderMockCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type PriceProviderMockCallerRaw struct {
	Contract *PriceProviderMockCaller // Generic read-only contract binding to access the raw methods on
}

// PriceProviderMockTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type PriceProviderMockTransactorRaw struct {
	Contract *PriceProviderMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceProviderMock creates a new instance of PriceProviderMock, bound to a specific deployed contract.
func NewPriceProviderMock(address common.Address, backend bind.ContractBackend) (*PriceProviderMock, error) {
	contract, err := bindPriceProviderMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceProviderMock{PriceProviderMockCaller: PriceProviderMockCaller{contract: contract}, PriceProviderMockTransactor: PriceProviderMockTransactor{contract: contract}, PriceProviderMockFilterer: PriceProviderMockFilterer{contract: contract}}, nil
}

// NewPriceProviderMockCaller creates a new read-only instance of PriceProviderMock, bound to a specific deployed contract.
func NewPriceProviderMockCaller(address common.Address, caller bind.ContractCaller) (*PriceProviderMockCaller, error) {
	contract, err := bindPriceProviderMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceProviderMockCaller{contract: contract}, nil
}

// NewPriceProviderMockTransactor creates a new write-only instance of PriceProviderMock, bound to a specific deployed contract.
func NewPriceProviderMockTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceProviderMockTransactor, error) {
	contract, err := bindPriceProviderMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceProviderMockTransactor{contract: contract}, nil
}

// NewPriceProviderMockFilterer creates a new log filterer instance of PriceProviderMock, bound to a specific deployed contract.
func NewPriceProviderMockFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceProviderMockFilterer, error) {
	contract, err := bindPriceProviderMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceProviderMockFilterer{contract: contract}, nil
}

// bindPriceProviderMock binds a generic wrapper to an already deployed contract.
func bindPriceProviderMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceProviderMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceProviderMock *PriceProviderMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceProviderMock.Contract.PriceProviderMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceProviderMock *PriceProviderMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceProviderMock.Contract.PriceProviderMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceProviderMock *PriceProviderMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceProviderMock.Contract.PriceProviderMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceProviderMock *PriceProviderMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceProviderMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceProviderMock *PriceProviderMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceProviderMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceProviderMock *PriceProviderMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceProviderMock.Contract.contract.Transact(opts, method, params...)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_PriceProviderMock *PriceProviderMockCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceProviderMock.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_PriceProviderMock *PriceProviderMockSession) Price() (*big.Int, error) {
	return _PriceProviderMock.Contract.Price(&_PriceProviderMock.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_PriceProviderMock *PriceProviderMockCallerSession) Price() (*big.Int, error) {
	return _PriceProviderMock.Contract.Price(&_PriceProviderMock.CallOpts)
}
