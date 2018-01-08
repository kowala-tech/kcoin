// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package network

import (
	"strings"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// NetworkContractsMapContractABI is the input ABI used to generate the binding from.
const NetworkContractsMapContractABI = `[{"constant":true,"inputs":[],"name":"priceOracle","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"setNetworkStats","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"setPriceOracle","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"setMToken","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"networkStats","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"mToken","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"_mToken","type":"address"},{"name":"_priceOracle","type":"address"},{"name":"_networkStats","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"oldAddr","type":"address"},{"indexed":false,"name":"newAddr","type":"address"}],"name":"OwnershipTransfer","type":"event"}]`

// NetworkContractsMapContractBin is the compiled bytecode used for deploying new contracts.
const NetworkContractsMapContractBin = `6060604052341561000f57600080fd5b60405160608061078e83398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505061063c806101526000396000f300606060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632630c12f146100885780632e06efd2146100dd578063530e784f14610116578063b882c4d81461014f578063c021e0f614610188578063c3b6f939146101dd578063f2fde38b14610232575b600080fd5b341561009357600080fd5b61009b61026b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156100e857600080fd5b610114600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610291565b005b341561012157600080fd5b61014d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610330565b005b341561015a57600080fd5b610186600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506103cf565b005b341561019357600080fd5b61019b61046e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156101e857600080fd5b6101f0610494565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561023d57600080fd5b610269600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104ba565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102ec57600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561038b57600080fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561042a57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561051557600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505600a165627a7a7230582075d07f88864b0e4f17b91811f19f3380cde9fd87e8d17c17fc0d0c98249b49400029`

// DeployNetworkContractsMapContract deploys a new Ethereum contract, binding an instance of NetworkContractsMapContract to it.
func DeployNetworkContractsMapContract(auth *bind.TransactOpts, backend bind.ContractBackend, _mToken common.Address, _priceOracle common.Address, _networkStats common.Address) (common.Address, *types.Transaction, *NetworkContractsMapContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractsMapContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkContractsMapContractBin), backend, _mToken, _priceOracle, _networkStats)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NetworkContractsMapContract{NetworkContractsMapContractCaller: NetworkContractsMapContractCaller{contract: contract}, NetworkContractsMapContractTransactor: NetworkContractsMapContractTransactor{contract: contract}}, nil
}

// NetworkContractsMapContract is an auto generated Go binding around an Ethereum contract.
type NetworkContractsMapContract struct {
	NetworkContractsMapContractCaller     // Read-only binding to the contract
	NetworkContractsMapContractTransactor // Write-only binding to the contract
}

// NetworkContractsMapContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkContractsMapContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractsMapContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkContractsMapContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractsMapContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkContractsMapContractSession struct {
	Contract     *NetworkContractsMapContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NetworkContractsMapContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkContractsMapContractCallerSession struct {
	Contract *NetworkContractsMapContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// NetworkContractsMapContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkContractsMapContractTransactorSession struct {
	Contract     *NetworkContractsMapContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// NetworkContractsMapContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkContractsMapContractRaw struct {
	Contract *NetworkContractsMapContract // Generic contract binding to access the raw methods on
}

// NetworkContractsMapContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkContractsMapContractCallerRaw struct {
	Contract *NetworkContractsMapContractCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkContractsMapContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkContractsMapContractTransactorRaw struct {
	Contract *NetworkContractsMapContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkContractsMapContract creates a new instance of NetworkContractsMapContract, bound to a specific deployed contract.
func NewNetworkContractsMapContract(address common.Address, backend bind.ContractBackend) (*NetworkContractsMapContract, error) {
	contract, err := bindNetworkContractsMapContract(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &NetworkContractsMapContract{NetworkContractsMapContractCaller: NetworkContractsMapContractCaller{contract: contract}, NetworkContractsMapContractTransactor: NetworkContractsMapContractTransactor{contract: contract}}, nil
}

// NewNetworkContractsMapContractCaller creates a new read-only instance of NetworkContractsMapContract, bound to a specific deployed contract.
func NewNetworkContractsMapContractCaller(address common.Address, caller bind.ContractCaller) (*NetworkContractsMapContractCaller, error) {
	contract, err := bindNetworkContractsMapContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkContractsMapContractCaller{contract: contract}, nil
}

// NewNetworkContractsMapContractTransactor creates a new write-only instance of NetworkContractsMapContract, bound to a specific deployed contract.
func NewNetworkContractsMapContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkContractsMapContractTransactor, error) {
	contract, err := bindNetworkContractsMapContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NetworkContractsMapContractTransactor{contract: contract}, nil
}

// bindNetworkContractsMapContract binds a generic wrapper to an already deployed contract.
func bindNetworkContractsMapContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractsMapContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContractsMapContract *NetworkContractsMapContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkContractsMapContract.Contract.NetworkContractsMapContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContractsMapContract *NetworkContractsMapContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.NetworkContractsMapContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContractsMapContract *NetworkContractsMapContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.NetworkContractsMapContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContractsMapContract *NetworkContractsMapContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkContractsMapContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.contract.Transact(opts, method, params...)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractCaller) MToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NetworkContractsMapContract.contract.Call(opts, out, "mToken")
	return *ret0, err
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractSession) MToken() (common.Address, error) {
	return _NetworkContractsMapContract.Contract.MToken(&_NetworkContractsMapContract.CallOpts)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractCallerSession) MToken() (common.Address, error) {
	return _NetworkContractsMapContract.Contract.MToken(&_NetworkContractsMapContract.CallOpts)
}

// NetworkStats is a free data retrieval call binding the contract method 0xc021e0f6.
//
// Solidity: function networkStats() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractCaller) NetworkStats(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NetworkContractsMapContract.contract.Call(opts, out, "networkStats")
	return *ret0, err
}

// NetworkStats is a free data retrieval call binding the contract method 0xc021e0f6.
//
// Solidity: function networkStats() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractSession) NetworkStats() (common.Address, error) {
	return _NetworkContractsMapContract.Contract.NetworkStats(&_NetworkContractsMapContract.CallOpts)
}

// NetworkStats is a free data retrieval call binding the contract method 0xc021e0f6.
//
// Solidity: function networkStats() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractCallerSession) NetworkStats() (common.Address, error) {
	return _NetworkContractsMapContract.Contract.NetworkStats(&_NetworkContractsMapContract.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NetworkContractsMapContract.contract.Call(opts, out, "priceOracle")
	return *ret0, err
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractSession) PriceOracle() (common.Address, error) {
	return _NetworkContractsMapContract.Contract.PriceOracle(&_NetworkContractsMapContract.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() constant returns(address)
func (_NetworkContractsMapContract *NetworkContractsMapContractCallerSession) PriceOracle() (common.Address, error) {
	return _NetworkContractsMapContract.Contract.PriceOracle(&_NetworkContractsMapContract.CallOpts)
}

// SetMToken is a paid mutator transaction binding the contract method 0xb882c4d8.
//
// Solidity: function setMToken(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactor) SetMToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.contract.Transact(opts, "setMToken", addr)
}

// SetMToken is a paid mutator transaction binding the contract method 0xb882c4d8.
//
// Solidity: function setMToken(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractSession) SetMToken(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.SetMToken(&_NetworkContractsMapContract.TransactOpts, addr)
}

// SetMToken is a paid mutator transaction binding the contract method 0xb882c4d8.
//
// Solidity: function setMToken(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactorSession) SetMToken(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.SetMToken(&_NetworkContractsMapContract.TransactOpts, addr)
}

// SetNetworkStats is a paid mutator transaction binding the contract method 0x2e06efd2.
//
// Solidity: function setNetworkStats(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactor) SetNetworkStats(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.contract.Transact(opts, "setNetworkStats", addr)
}

// SetNetworkStats is a paid mutator transaction binding the contract method 0x2e06efd2.
//
// Solidity: function setNetworkStats(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractSession) SetNetworkStats(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.SetNetworkStats(&_NetworkContractsMapContract.TransactOpts, addr)
}

// SetNetworkStats is a paid mutator transaction binding the contract method 0x2e06efd2.
//
// Solidity: function setNetworkStats(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactorSession) SetNetworkStats(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.SetNetworkStats(&_NetworkContractsMapContract.TransactOpts, addr)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactor) SetPriceOracle(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.contract.Transact(opts, "setPriceOracle", addr)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractSession) SetPriceOracle(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.SetPriceOracle(&_NetworkContractsMapContract.TransactOpts, addr)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactorSession) SetPriceOracle(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.SetPriceOracle(&_NetworkContractsMapContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.TransferOwnership(&_NetworkContractsMapContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContractsMapContract *NetworkContractsMapContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkContractsMapContract.Contract.TransferOwnership(&_NetworkContractsMapContract.TransactOpts, addr)
}
