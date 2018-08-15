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

// DomainResolverMockABI is the input ABI used to generate the binding from.
const DomainResolverMockABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DomainResolverMockBin is the compiled bytecode used for deploying new contracts.
const DomainResolverMockBin = `6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820a7012f1005ee45bbeaed29d2b6fe4aa9bf765a4585dd3483584497d087e913eb0029`

// DeployDomainResolverMock deploys a new Kowala contract, binding an instance of DomainResolverMock to it.
func DeployDomainResolverMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DomainResolverMock, error) {
	parsed, err := abi.JSON(strings.NewReader(DomainResolverMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DomainResolverMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DomainResolverMock{DomainResolverMockCaller: DomainResolverMockCaller{contract: contract}, DomainResolverMockTransactor: DomainResolverMockTransactor{contract: contract}, DomainResolverMockFilterer: DomainResolverMockFilterer{contract: contract}}, nil
}

// DomainResolverMock is an auto generated Go binding around a Kowala contract.
type DomainResolverMock struct {
	DomainResolverMockCaller     // Read-only binding to the contract
	DomainResolverMockTransactor // Write-only binding to the contract
	DomainResolverMockFilterer   // Log filterer for contract events
}

// DomainResolverMockCaller is an auto generated read-only Go binding around a Kowala contract.
type DomainResolverMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainResolverMockTransactor is an auto generated write-only Go binding around a Kowala contract.
type DomainResolverMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainResolverMockFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type DomainResolverMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainResolverMockSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type DomainResolverMockSession struct {
	Contract     *DomainResolverMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DomainResolverMockCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type DomainResolverMockCallerSession struct {
	Contract *DomainResolverMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DomainResolverMockTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type DomainResolverMockTransactorSession struct {
	Contract     *DomainResolverMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DomainResolverMockRaw is an auto generated low-level Go binding around a Kowala contract.
type DomainResolverMockRaw struct {
	Contract *DomainResolverMock // Generic contract binding to access the raw methods on
}

// DomainResolverMockCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type DomainResolverMockCallerRaw struct {
	Contract *DomainResolverMockCaller // Generic read-only contract binding to access the raw methods on
}

// DomainResolverMockTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type DomainResolverMockTransactorRaw struct {
	Contract *DomainResolverMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomainResolverMock creates a new instance of DomainResolverMock, bound to a specific deployed contract.
func NewDomainResolverMock(address common.Address, backend bind.ContractBackend) (*DomainResolverMock, error) {
	contract, err := bindDomainResolverMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DomainResolverMock{DomainResolverMockCaller: DomainResolverMockCaller{contract: contract}, DomainResolverMockTransactor: DomainResolverMockTransactor{contract: contract}, DomainResolverMockFilterer: DomainResolverMockFilterer{contract: contract}}, nil
}

// NewDomainResolverMockCaller creates a new read-only instance of DomainResolverMock, bound to a specific deployed contract.
func NewDomainResolverMockCaller(address common.Address, caller bind.ContractCaller) (*DomainResolverMockCaller, error) {
	contract, err := bindDomainResolverMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomainResolverMockCaller{contract: contract}, nil
}

// NewDomainResolverMockTransactor creates a new write-only instance of DomainResolverMock, bound to a specific deployed contract.
func NewDomainResolverMockTransactor(address common.Address, transactor bind.ContractTransactor) (*DomainResolverMockTransactor, error) {
	contract, err := bindDomainResolverMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomainResolverMockTransactor{contract: contract}, nil
}

// NewDomainResolverMockFilterer creates a new log filterer instance of DomainResolverMock, bound to a specific deployed contract.
func NewDomainResolverMockFilterer(address common.Address, filterer bind.ContractFilterer) (*DomainResolverMockFilterer, error) {
	contract, err := bindDomainResolverMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomainResolverMockFilterer{contract: contract}, nil
}

// bindDomainResolverMock binds a generic wrapper to an already deployed contract.
func bindDomainResolverMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DomainResolverMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainResolverMock *DomainResolverMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DomainResolverMock.Contract.DomainResolverMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainResolverMock *DomainResolverMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainResolverMock.Contract.DomainResolverMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainResolverMock *DomainResolverMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainResolverMock.Contract.DomainResolverMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainResolverMock *DomainResolverMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DomainResolverMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainResolverMock *DomainResolverMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainResolverMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainResolverMock *DomainResolverMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainResolverMock.Contract.contract.Transact(opts, method, params...)
}
