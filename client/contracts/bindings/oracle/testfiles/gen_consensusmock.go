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

// ConsensusMockABI is the input ABI used to generate the binding from.
const ConsensusMockABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isSuperNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_superNode\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ConsensusMockBin is the compiled bytecode used for deploying new contracts.
const ConsensusMockBin = `608060405234801561001057600080fd5b5060405160208061013b83398101806040528101908080519060200190929190505050806000806101000a81548160ff0219169083151502179055505060e08061005b6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680637d0e81bf146044575b600080fd5b348015604f57600080fd5b506082600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050609c565b604051808215151515815260200191505060405180910390f35b60008060009054906101000a900460ff1690509190505600a165627a7a72305820d59a8a261ae038cc850cfe737fdfeda66c7e16cfaac8929833698a0bd0a536750029`

// DeployConsensusMock deploys a new Kowala contract, binding an instance of ConsensusMock to it.
func DeployConsensusMock(auth *bind.TransactOpts, backend bind.ContractBackend, _superNode bool) (common.Address, *types.Transaction, *ConsensusMock, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsensusMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConsensusMockBin), backend, _superNode)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConsensusMock{ConsensusMockCaller: ConsensusMockCaller{contract: contract}, ConsensusMockTransactor: ConsensusMockTransactor{contract: contract}, ConsensusMockFilterer: ConsensusMockFilterer{contract: contract}}, nil
}

// ConsensusMock is an auto generated Go binding around a Kowala contract.
type ConsensusMock struct {
	ConsensusMockCaller     // Read-only binding to the contract
	ConsensusMockTransactor // Write-only binding to the contract
	ConsensusMockFilterer   // Log filterer for contract events
}

// ConsensusMockCaller is an auto generated read-only Go binding around a Kowala contract.
type ConsensusMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusMockTransactor is an auto generated write-only Go binding around a Kowala contract.
type ConsensusMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusMockFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type ConsensusMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusMockSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type ConsensusMockSession struct {
	Contract     *ConsensusMock    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsensusMockCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type ConsensusMockCallerSession struct {
	Contract *ConsensusMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConsensusMockTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type ConsensusMockTransactorSession struct {
	Contract     *ConsensusMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConsensusMockRaw is an auto generated low-level Go binding around a Kowala contract.
type ConsensusMockRaw struct {
	Contract *ConsensusMock // Generic contract binding to access the raw methods on
}

// ConsensusMockCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type ConsensusMockCallerRaw struct {
	Contract *ConsensusMockCaller // Generic read-only contract binding to access the raw methods on
}

// ConsensusMockTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type ConsensusMockTransactorRaw struct {
	Contract *ConsensusMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsensusMock creates a new instance of ConsensusMock, bound to a specific deployed contract.
func NewConsensusMock(address common.Address, backend bind.ContractBackend) (*ConsensusMock, error) {
	contract, err := bindConsensusMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsensusMock{ConsensusMockCaller: ConsensusMockCaller{contract: contract}, ConsensusMockTransactor: ConsensusMockTransactor{contract: contract}, ConsensusMockFilterer: ConsensusMockFilterer{contract: contract}}, nil
}

// NewConsensusMockCaller creates a new read-only instance of ConsensusMock, bound to a specific deployed contract.
func NewConsensusMockCaller(address common.Address, caller bind.ContractCaller) (*ConsensusMockCaller, error) {
	contract, err := bindConsensusMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusMockCaller{contract: contract}, nil
}

// NewConsensusMockTransactor creates a new write-only instance of ConsensusMock, bound to a specific deployed contract.
func NewConsensusMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsensusMockTransactor, error) {
	contract, err := bindConsensusMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusMockTransactor{contract: contract}, nil
}

// NewConsensusMockFilterer creates a new log filterer instance of ConsensusMock, bound to a specific deployed contract.
func NewConsensusMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsensusMockFilterer, error) {
	contract, err := bindConsensusMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsensusMockFilterer{contract: contract}, nil
}

// bindConsensusMock binds a generic wrapper to an already deployed contract.
func bindConsensusMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsensusMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusMock *ConsensusMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConsensusMock.Contract.ConsensusMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusMock *ConsensusMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusMock.Contract.ConsensusMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusMock *ConsensusMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusMock.Contract.ConsensusMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusMock *ConsensusMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConsensusMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusMock *ConsensusMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusMock *ConsensusMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusMock.Contract.contract.Transact(opts, method, params...)
}

// IsSuperNode is a free data retrieval call binding the contract method 0x7d0e81bf.
//
// Solidity: function isSuperNode(identity address) constant returns(bool)
func (_ConsensusMock *ConsensusMockCaller) IsSuperNode(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ConsensusMock.contract.Call(opts, out, "isSuperNode", identity)
	return *ret0, err
}

// IsSuperNode is a free data retrieval call binding the contract method 0x7d0e81bf.
//
// Solidity: function isSuperNode(identity address) constant returns(bool)
func (_ConsensusMock *ConsensusMockSession) IsSuperNode(identity common.Address) (bool, error) {
	return _ConsensusMock.Contract.IsSuperNode(&_ConsensusMock.CallOpts, identity)
}

// IsSuperNode is a free data retrieval call binding the contract method 0x7d0e81bf.
//
// Solidity: function isSuperNode(identity address) constant returns(bool)
func (_ConsensusMock *ConsensusMockCallerSession) IsSuperNode(identity common.Address) (bool, error) {
	return _ConsensusMock.Contract.IsSuperNode(&_ConsensusMock.CallOpts, identity)
}
