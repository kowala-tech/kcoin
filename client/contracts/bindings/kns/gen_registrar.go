// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kns

import (
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// FIFSRegistrarABI is the input ABI used to generate the binding from.
const FIFSRegistrarABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"knsAddr\",\"type\":\"address\"},{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subnode\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"knsAddr\",\"type\":\"address\"},{\"name\":\"node\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FIFSRegistrarBin is the compiled bytecode used for deploying new contracts.
const FIFSRegistrarBin = `608060405234801561001057600080fd5b50604051604080610578833981018060405281019080805190602001909291908051906020019092919050505081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806001816000191690555050506104de8061009a6000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063158ef93e1461005c578063be13f47c1461008b578063d22057a9146100dc575b600080fd5b34801561006857600080fd5b5061007161012d565b604051808215151515815260200191505060405180910390f35b34801561009757600080fd5b506100da600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919050505061013f565b005b3480156100e857600080fd5b5061012b6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610253565b005b6000809054906101000a900460ff1681565b6000809054906101000a900460ff161515156101e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b81600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806001816000191690555060016000806101000a81548160ff0219169083151502179055505050565b8160008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302571be36001548460405180836000191660001916815260200182600019166000191681526020019250505060405180910390206040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561031d57600080fd5b505af1158015610331573d6000803e3d6000fd5b505050506040513d602081101561034757600080fd5b8101908080519060200190929190505050905060008173ffffffffffffffffffffffffffffffffffffffff1614806103aa57503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15156103b557600080fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306ab592360015486866040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180846000191660001916815260200183600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050600060405180830381600087803b15801561049457600080fd5b505af11580156104a8573d6000803e3d6000fd5b50505050505050505600a165627a7a723058205053cd7ff3b82a7e6e8531ed12a7195800886445328be2a9911c057f02b1f0cf0029`

// DeployFIFSRegistrar deploys a new Kowala contract, binding an instance of FIFSRegistrar to it.
func DeployFIFSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, knsAddr common.Address, node [32]byte) (common.Address, *types.Transaction, *FIFSRegistrar, error) {
	parsed, err := abi.JSON(strings.NewReader(FIFSRegistrarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FIFSRegistrarBin), backend, knsAddr, node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FIFSRegistrar{FIFSRegistrarCaller: FIFSRegistrarCaller{contract: contract}, FIFSRegistrarTransactor: FIFSRegistrarTransactor{contract: contract}, FIFSRegistrarFilterer: FIFSRegistrarFilterer{contract: contract}}, nil
}

// FIFSRegistrar is an auto generated Go binding around a Kowala contract.
type FIFSRegistrar struct {
	FIFSRegistrarCaller     // Read-only binding to the contract
	FIFSRegistrarTransactor // Write-only binding to the contract
	FIFSRegistrarFilterer   // Log filterer for contract events
}

// FIFSRegistrarCaller is an auto generated read-only Go binding around a Kowala contract.
type FIFSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarTransactor is an auto generated write-only Go binding around a Kowala contract.
type FIFSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type FIFSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type FIFSRegistrarSession struct {
	Contract     *FIFSRegistrar    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FIFSRegistrarCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type FIFSRegistrarCallerSession struct {
	Contract *FIFSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FIFSRegistrarTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type FIFSRegistrarTransactorSession struct {
	Contract     *FIFSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FIFSRegistrarRaw is an auto generated low-level Go binding around a Kowala contract.
type FIFSRegistrarRaw struct {
	Contract *FIFSRegistrar // Generic contract binding to access the raw methods on
}

// FIFSRegistrarCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type FIFSRegistrarCallerRaw struct {
	Contract *FIFSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// FIFSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type FIFSRegistrarTransactorRaw struct {
	Contract *FIFSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFIFSRegistrar creates a new instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrar(address common.Address, backend bind.ContractBackend) (*FIFSRegistrar, error) {
	contract, err := bindFIFSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrar{FIFSRegistrarCaller: FIFSRegistrarCaller{contract: contract}, FIFSRegistrarTransactor: FIFSRegistrarTransactor{contract: contract}, FIFSRegistrarFilterer: FIFSRegistrarFilterer{contract: contract}}, nil
}

// NewFIFSRegistrarCaller creates a new read-only instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*FIFSRegistrarCaller, error) {
	contract, err := bindFIFSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarCaller{contract: contract}, nil
}

// NewFIFSRegistrarTransactor creates a new write-only instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*FIFSRegistrarTransactor, error) {
	contract, err := bindFIFSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarTransactor{contract: contract}, nil
}

// NewFIFSRegistrarFilterer creates a new log filterer instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*FIFSRegistrarFilterer, error) {
	contract, err := bindFIFSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarFilterer{contract: contract}, nil
}

// bindFIFSRegistrar binds a generic wrapper to an already deployed contract.
func bindFIFSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FIFSRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FIFSRegistrar *FIFSRegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FIFSRegistrar.Contract.FIFSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FIFSRegistrar *FIFSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.FIFSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FIFSRegistrar *FIFSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.FIFSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FIFSRegistrar *FIFSRegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FIFSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FIFSRegistrar *FIFSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FIFSRegistrar *FIFSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_FIFSRegistrar *FIFSRegistrarCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FIFSRegistrar.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_FIFSRegistrar *FIFSRegistrarSession) Initialized() (bool, error) {
	return _FIFSRegistrar.Contract.Initialized(&_FIFSRegistrar.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_FIFSRegistrar *FIFSRegistrarCallerSession) Initialized() (bool, error) {
	return _FIFSRegistrar.Contract.Initialized(&_FIFSRegistrar.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(knsAddr address, node bytes32) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactor) Initialize(opts *bind.TransactOpts, knsAddr common.Address, node [32]byte) (*types.Transaction, error) {
	return _FIFSRegistrar.contract.Transact(opts, "initialize", knsAddr, node)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(knsAddr address, node bytes32) returns()
func (_FIFSRegistrar *FIFSRegistrarSession) Initialize(knsAddr common.Address, node [32]byte) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Initialize(&_FIFSRegistrar.TransactOpts, knsAddr, node)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe13f47c.
//
// Solidity: function initialize(knsAddr address, node bytes32) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactorSession) Initialize(knsAddr common.Address, node [32]byte) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Initialize(&_FIFSRegistrar.TransactOpts, knsAddr, node)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactor) Register(opts *bind.TransactOpts, subnode [32]byte, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.contract.Transact(opts, "register", subnode, owner)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarSession) Register(subnode [32]byte, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Register(&_FIFSRegistrar.TransactOpts, subnode, owner)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactorSession) Register(subnode [32]byte, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Register(&_FIFSRegistrar.TransactOpts, subnode, owner)
}
