// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package network

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// NetworkStatsContractABI is the input ABI used to generate the binding from.
const NetworkStatsContractABI = `[{"constant":true,"inputs":[],"name":"totalMinedWei","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"oldAddr","type":"address"},{"indexed":false,"name":"newAddr","type":"address"}],"name":"OwnershipTransfer","type":"event"}]`

// NetworkStatsContractBin is the compiled bytecode used for deploying new contracts.
const NetworkStatsContractBin = `60606040526000600155336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061023b806100586000396000f30060606040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806321d3fab314610051578063f2fde38b1461007a575b600080fd5b341561005c57600080fd5b6100646100b3565b6040518082815260200191505060405180910390f35b341561008557600080fd5b6100b1600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506100b9565b005b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561011457600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505600a165627a7a72305820cb527a869e9553177f98549f6afb220768dd98195f45f8ca33d54ad5df2e053f0029`

// DeployNetworkStatsContract deploys a new Ethereum contract, binding an instance of NetworkStatsContract to it.
func DeployNetworkStatsContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NetworkStatsContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkStatsContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkStatsContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NetworkStatsContract{NetworkStatsContractCaller: NetworkStatsContractCaller{contract: contract}, NetworkStatsContractTransactor: NetworkStatsContractTransactor{contract: contract}}, nil
}

// NetworkStatsContract is an auto generated Go binding around an Ethereum contract.
type NetworkStatsContract struct {
	NetworkStatsContractCaller     // Read-only binding to the contract
	NetworkStatsContractTransactor // Write-only binding to the contract
}

// NetworkStatsContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkStatsContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkStatsContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkStatsContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkStatsContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkStatsContractSession struct {
	Contract     *NetworkStatsContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NetworkStatsContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkStatsContractCallerSession struct {
	Contract *NetworkStatsContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// NetworkStatsContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkStatsContractTransactorSession struct {
	Contract     *NetworkStatsContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// NetworkStatsContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkStatsContractRaw struct {
	Contract *NetworkStatsContract // Generic contract binding to access the raw methods on
}

// NetworkStatsContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkStatsContractCallerRaw struct {
	Contract *NetworkStatsContractCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkStatsContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkStatsContractTransactorRaw struct {
	Contract *NetworkStatsContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkStatsContract creates a new instance of NetworkStatsContract, bound to a specific deployed contract.
func NewNetworkStatsContract(address common.Address, backend bind.ContractBackend) (*NetworkStatsContract, error) {
	contract, err := bindNetworkStatsContract(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &NetworkStatsContract{NetworkStatsContractCaller: NetworkStatsContractCaller{contract: contract}, NetworkStatsContractTransactor: NetworkStatsContractTransactor{contract: contract}}, nil
}

// NewNetworkStatsContractCaller creates a new read-only instance of NetworkStatsContract, bound to a specific deployed contract.
func NewNetworkStatsContractCaller(address common.Address, caller bind.ContractCaller) (*NetworkStatsContractCaller, error) {
	contract, err := bindNetworkStatsContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkStatsContractCaller{contract: contract}, nil
}

// NewNetworkStatsContractTransactor creates a new write-only instance of NetworkStatsContract, bound to a specific deployed contract.
func NewNetworkStatsContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkStatsContractTransactor, error) {
	contract, err := bindNetworkStatsContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NetworkStatsContractTransactor{contract: contract}, nil
}

// bindNetworkStatsContract binds a generic wrapper to an already deployed contract.
func bindNetworkStatsContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkStatsContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkStatsContract *NetworkStatsContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkStatsContract.Contract.NetworkStatsContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkStatsContract *NetworkStatsContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStatsContract.Contract.NetworkStatsContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkStatsContract *NetworkStatsContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkStatsContract.Contract.NetworkStatsContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkStatsContract *NetworkStatsContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkStatsContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkStatsContract *NetworkStatsContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStatsContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkStatsContract *NetworkStatsContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkStatsContract.Contract.contract.Transact(opts, method, params...)
}

// TotalMinedWei is a free data retrieval call binding the contract method 0x21d3fab3.
//
// Solidity: function totalMinedWei() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCaller) TotalMinedWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkStatsContract.contract.Call(opts, out, "totalMinedWei")
	return *ret0, err
}

// TotalMinedWei is a free data retrieval call binding the contract method 0x21d3fab3.
//
// Solidity: function totalMinedWei() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractSession) TotalMinedWei() (*big.Int, error) {
	return _NetworkStatsContract.Contract.TotalMinedWei(&_NetworkStatsContract.CallOpts)
}

// TotalMinedWei is a free data retrieval call binding the contract method 0x21d3fab3.
//
// Solidity: function totalMinedWei() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCallerSession) TotalMinedWei() (*big.Int, error) {
	return _NetworkStatsContract.Contract.TotalMinedWei(&_NetworkStatsContract.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkStatsContract *NetworkStatsContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _NetworkStatsContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkStatsContract *NetworkStatsContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkStatsContract.Contract.TransferOwnership(&_NetworkStatsContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkStatsContract *NetworkStatsContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkStatsContract.Contract.TransferOwnership(&_NetworkStatsContract.TransactOpts, addr)
}
