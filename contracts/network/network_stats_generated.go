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
const NetworkStatsContractABI = `[{"constant":true,"inputs":[],"name":"lastPrice","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"lastBlockReward","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"totalSupplyWei","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

// NetworkStatsContractBin is the compiled bytecode used for deploying new contracts.
const NetworkStatsContractBin = `60606040526000805560006001556000600255341561001d57600080fd5b6101088061002c6000396000f3006060604052600436106053576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063053f14da1460585780635798a6d514607e57806370a8f25b1460a4575b600080fd5b3415606257600080fd5b606860ca565b6040518082815260200191505060405180910390f35b3415608857600080fd5b608e60d0565b6040518082815260200191505060405180910390f35b341560ae57600080fd5b60b460d6565b6040518082815260200191505060405180910390f35b60025481565b60015481565b600054815600a165627a7a7230582092305fdcffddbc74a54c6b0f8816287f2956f15f9bf3b7af49c1b6f7b917a1c10029`

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

// LastBlockReward is a free data retrieval call binding the contract method 0x5798a6d5.
//
// Solidity: function lastBlockReward() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCaller) LastBlockReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkStatsContract.contract.Call(opts, out, "lastBlockReward")
	return *ret0, err
}

// LastBlockReward is a free data retrieval call binding the contract method 0x5798a6d5.
//
// Solidity: function lastBlockReward() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractSession) LastBlockReward() (*big.Int, error) {
	return _NetworkStatsContract.Contract.LastBlockReward(&_NetworkStatsContract.CallOpts)
}

// LastBlockReward is a free data retrieval call binding the contract method 0x5798a6d5.
//
// Solidity: function lastBlockReward() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCallerSession) LastBlockReward() (*big.Int, error) {
	return _NetworkStatsContract.Contract.LastBlockReward(&_NetworkStatsContract.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkStatsContract.contract.Call(opts, out, "lastPrice")
	return *ret0, err
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractSession) LastPrice() (*big.Int, error) {
	return _NetworkStatsContract.Contract.LastPrice(&_NetworkStatsContract.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCallerSession) LastPrice() (*big.Int, error) {
	return _NetworkStatsContract.Contract.LastPrice(&_NetworkStatsContract.CallOpts)
}

// TotalSupplyWei is a free data retrieval call binding the contract method 0x70a8f25b.
//
// Solidity: function totalSupplyWei() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCaller) TotalSupplyWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkStatsContract.contract.Call(opts, out, "totalSupplyWei")
	return *ret0, err
}

// TotalSupplyWei is a free data retrieval call binding the contract method 0x70a8f25b.
//
// Solidity: function totalSupplyWei() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractSession) TotalSupplyWei() (*big.Int, error) {
	return _NetworkStatsContract.Contract.TotalSupplyWei(&_NetworkStatsContract.CallOpts)
}

// TotalSupplyWei is a free data retrieval call binding the contract method 0x70a8f25b.
//
// Solidity: function totalSupplyWei() constant returns(uint256)
func (_NetworkStatsContract *NetworkStatsContractCallerSession) TotalSupplyWei() (*big.Int, error) {
	return _NetworkStatsContract.Contract.TotalSupplyWei(&_NetworkStatsContract.CallOpts)
}
