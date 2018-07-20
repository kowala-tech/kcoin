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

// IncompatibleABI is the input ABI used to generate the binding from.
const IncompatibleABI = "[]"

// IncompatibleBin is the compiled bytecode used for deploying new contracts.
const IncompatibleBin = `6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a723058209ca218a8fe7b2d879a8a603f315967018cb655d3f3c836b4a9f8b4536e5734a60029`

// DeployIncompatible deploys a new Kowala contract, binding an instance of Incompatible to it.
func DeployIncompatible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Incompatible, error) {
	parsed, err := abi.JSON(strings.NewReader(IncompatibleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncompatibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Incompatible{IncompatibleCaller: IncompatibleCaller{contract: contract}, IncompatibleTransactor: IncompatibleTransactor{contract: contract}, IncompatibleFilterer: IncompatibleFilterer{contract: contract}}, nil
}

// Incompatible is an auto generated Go binding around a Kowala contract.
type Incompatible struct {
	IncompatibleCaller     // Read-only binding to the contract
	IncompatibleTransactor // Write-only binding to the contract
	IncompatibleFilterer   // Log filterer for contract events
}

// IncompatibleCaller is an auto generated read-only Go binding around a Kowala contract.
type IncompatibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncompatibleTransactor is an auto generated write-only Go binding around a Kowala contract.
type IncompatibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncompatibleFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type IncompatibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncompatibleSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type IncompatibleSession struct {
	Contract     *Incompatible     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncompatibleCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type IncompatibleCallerSession struct {
	Contract *IncompatibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IncompatibleTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type IncompatibleTransactorSession struct {
	Contract     *IncompatibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IncompatibleRaw is an auto generated low-level Go binding around a Kowala contract.
type IncompatibleRaw struct {
	Contract *Incompatible // Generic contract binding to access the raw methods on
}

// IncompatibleCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type IncompatibleCallerRaw struct {
	Contract *IncompatibleCaller // Generic read-only contract binding to access the raw methods on
}

// IncompatibleTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type IncompatibleTransactorRaw struct {
	Contract *IncompatibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncompatible creates a new instance of Incompatible, bound to a specific deployed contract.
func NewIncompatible(address common.Address, backend bind.ContractBackend) (*Incompatible, error) {
	contract, err := bindIncompatible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Incompatible{IncompatibleCaller: IncompatibleCaller{contract: contract}, IncompatibleTransactor: IncompatibleTransactor{contract: contract}, IncompatibleFilterer: IncompatibleFilterer{contract: contract}}, nil
}

// NewIncompatibleCaller creates a new read-only instance of Incompatible, bound to a specific deployed contract.
func NewIncompatibleCaller(address common.Address, caller bind.ContractCaller) (*IncompatibleCaller, error) {
	contract, err := bindIncompatible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncompatibleCaller{contract: contract}, nil
}

// NewIncompatibleTransactor creates a new write-only instance of Incompatible, bound to a specific deployed contract.
func NewIncompatibleTransactor(address common.Address, transactor bind.ContractTransactor) (*IncompatibleTransactor, error) {
	contract, err := bindIncompatible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncompatibleTransactor{contract: contract}, nil
}

// NewIncompatibleFilterer creates a new log filterer instance of Incompatible, bound to a specific deployed contract.
func NewIncompatibleFilterer(address common.Address, filterer bind.ContractFilterer) (*IncompatibleFilterer, error) {
	contract, err := bindIncompatible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncompatibleFilterer{contract: contract}, nil
}

// bindIncompatible binds a generic wrapper to an already deployed contract.
func bindIncompatible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncompatibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incompatible *IncompatibleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incompatible.Contract.IncompatibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incompatible *IncompatibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incompatible.Contract.IncompatibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incompatible *IncompatibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incompatible.Contract.IncompatibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incompatible *IncompatibleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incompatible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incompatible *IncompatibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incompatible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incompatible *IncompatibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incompatible.Contract.contract.Transact(opts, method, params...)
}
