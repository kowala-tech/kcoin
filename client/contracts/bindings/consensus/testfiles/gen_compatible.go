// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testfiles

import (
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// CompatibleABI is the input ABI used to generate the binding from.
const CompatibleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CompatibleBin is the compiled bytecode used for deploying new contracts.
const CompatibleBin = `6080604052348015600f57600080fd5b5060868061001e6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063f8a8fd6d146044575b600080fd5b348015604f57600080fd5b5060566058565b005b5600a165627a7a72305820c78b71ec6a158deb9adf28503d1aea5c377922c2236f7e1d38aae268fbe84ae90029`

// DeployCompatible deploys a new Kowala contract, binding an instance of Compatible to it.
func DeployCompatible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Compatible, error) {
	parsed, err := abi.JSON(strings.NewReader(CompatibleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompatibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Compatible{CompatibleCaller: CompatibleCaller{contract: contract}, CompatibleTransactor: CompatibleTransactor{contract: contract}, CompatibleFilterer: CompatibleFilterer{contract: contract}}, nil
}

// Compatible is an auto generated Go binding around a Kowala contract.
type Compatible struct {
	CompatibleCaller     // Read-only binding to the contract
	CompatibleTransactor // Write-only binding to the contract
	CompatibleFilterer   // Log filterer for contract events
}

// CompatibleCaller is an auto generated read-only Go binding around a Kowala contract.
type CompatibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibleTransactor is an auto generated write-only Go binding around a Kowala contract.
type CompatibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibleFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type CompatibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibleSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type CompatibleSession struct {
	Contract     *Compatible       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompatibleCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type CompatibleCallerSession struct {
	Contract *CompatibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CompatibleTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type CompatibleTransactorSession struct {
	Contract     *CompatibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CompatibleRaw is an auto generated low-level Go binding around a Kowala contract.
type CompatibleRaw struct {
	Contract *Compatible // Generic contract binding to access the raw methods on
}

// CompatibleCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type CompatibleCallerRaw struct {
	Contract *CompatibleCaller // Generic read-only contract binding to access the raw methods on
}

// CompatibleTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type CompatibleTransactorRaw struct {
	Contract *CompatibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompatible creates a new instance of Compatible, bound to a specific deployed contract.
func NewCompatible(address common.Address, backend bind.ContractBackend) (*Compatible, error) {
	contract, err := bindCompatible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compatible{CompatibleCaller: CompatibleCaller{contract: contract}, CompatibleTransactor: CompatibleTransactor{contract: contract}, CompatibleFilterer: CompatibleFilterer{contract: contract}}, nil
}

// NewCompatibleCaller creates a new read-only instance of Compatible, bound to a specific deployed contract.
func NewCompatibleCaller(address common.Address, caller bind.ContractCaller) (*CompatibleCaller, error) {
	contract, err := bindCompatible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibleCaller{contract: contract}, nil
}

// NewCompatibleTransactor creates a new write-only instance of Compatible, bound to a specific deployed contract.
func NewCompatibleTransactor(address common.Address, transactor bind.ContractTransactor) (*CompatibleTransactor, error) {
	contract, err := bindCompatible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibleTransactor{contract: contract}, nil
}

// NewCompatibleFilterer creates a new log filterer instance of Compatible, bound to a specific deployed contract.
func NewCompatibleFilterer(address common.Address, filterer bind.ContractFilterer) (*CompatibleFilterer, error) {
	contract, err := bindCompatible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompatibleFilterer{contract: contract}, nil
}

// bindCompatible binds a generic wrapper to an already deployed contract.
func bindCompatible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompatibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compatible *CompatibleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compatible.Contract.CompatibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compatible *CompatibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compatible.Contract.CompatibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compatible *CompatibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compatible.Contract.CompatibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compatible *CompatibleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compatible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compatible *CompatibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compatible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compatible *CompatibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compatible.Contract.contract.Transact(opts, method, params...)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Compatible *CompatibleTransactor) Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compatible.contract.Transact(opts, "test")
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Compatible *CompatibleSession) Test() (*types.Transaction, error) {
	return _Compatible.Contract.Test(&_Compatible.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Compatible *CompatibleTransactorSession) Test() (*types.Transaction, error) {
	return _Compatible.Contract.Test(&_Compatible.TransactOpts)
}
