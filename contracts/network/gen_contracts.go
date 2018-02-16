// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network

import (
	"strings"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// ContractsContractABI is the input ABI used to generate the binding from.
const ContractsContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNetworkContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_mToken\",\"type\":\"address\"},{\"name\":\"_priceOracle\",\"type\":\"address\"},{\"name\":\"_networkContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// ContractsContractBin is the compiled bytecode used for deploying new contracts.
const ContractsContractBin = `6060604052341561000f57600080fd5b60405160608061078e83398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505061063c806101526000396000f300606060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632630c12f14610088578063530e784f146100dd578063599b934814610116578063b882c4d81461014f578063c3b6f93914610188578063f2fde38b146101dd578063f842c21314610216575b600080fd5b341561009357600080fd5b61009b61026b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156100e857600080fd5b610114600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610291565b005b341561012157600080fd5b61014d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610330565b005b341561015a57600080fd5b610186600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506103cf565b005b341561019357600080fd5b61019b61046e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156101e857600080fd5b610214600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610494565b005b341561022157600080fd5b6102296105ea565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102ec57600080fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561038b57600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561042a57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156104ef57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820d84d03b069abb5e2de5b0b95ea514ad5a98466d8a98e70aee8045345f304eb310029`

// DeployContractsContract deploys a new Ethereum contract, binding an instance of ContractsContract to it.
func DeployContractsContract(auth *bind.TransactOpts, backend bind.ContractBackend, _mToken common.Address, _priceOracle common.Address, _networkContract common.Address) (common.Address, *types.Transaction, *ContractsContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractsContractBin), backend, _mToken, _priceOracle, _networkContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractsContract{ContractsContractCaller: ContractsContractCaller{contract: contract}, ContractsContractTransactor: ContractsContractTransactor{contract: contract}}, nil
}

// ContractsContract is an auto generated Go binding around an Ethereum contract.
type ContractsContract struct {
	ContractsContractCaller     // Read-only binding to the contract
	ContractsContractTransactor // Write-only binding to the contract
}

// ContractsContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsContractSession struct {
	Contract     *ContractsContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractsContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsContractCallerSession struct {
	Contract *ContractsContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractsContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsContractTransactorSession struct {
	Contract     *ContractsContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractsContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsContractRaw struct {
	Contract *ContractsContract // Generic contract binding to access the raw methods on
}

// ContractsContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsContractCallerRaw struct {
	Contract *ContractsContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsContractTransactorRaw struct {
	Contract *ContractsContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractsContract creates a new instance of ContractsContract, bound to a specific deployed contract.
func NewContractsContract(address common.Address, backend bind.ContractBackend) (*ContractsContract, error) {
	contract, err := bindContractsContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractsContract{ContractsContractCaller: ContractsContractCaller{contract: contract}, ContractsContractTransactor: ContractsContractTransactor{contract: contract}}, nil
}

// NewContractsContractCaller creates a new read-only instance of ContractsContract, bound to a specific deployed contract.
func NewContractsContractCaller(address common.Address, caller bind.ContractCaller) (*ContractsContractCaller, error) {
	contract, err := bindContractsContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsContractCaller{contract: contract}, nil
}

// NewContractsContractTransactor creates a new write-only instance of ContractsContract, bound to a specific deployed contract.
func NewContractsContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsContractTransactor, error) {
	contract, err := bindContractsContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ContractsContractTransactor{contract: contract}, nil
}

// bindContractsContract binds a generic wrapper to an already deployed contract.
func bindContractsContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractsContract *ContractsContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractsContract.Contract.ContractsContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractsContract *ContractsContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractsContract.Contract.ContractsContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractsContract *ContractsContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractsContract.Contract.ContractsContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractsContract *ContractsContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractsContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractsContract *ContractsContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractsContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractsContract *ContractsContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractsContract.Contract.contract.Transact(opts, method, params...)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() constant returns(address)
func (_ContractsContract *ContractsContractCaller) MToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractsContract.contract.Call(opts, out, "mToken")
	return *ret0, err
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() constant returns(address)
func (_ContractsContract *ContractsContractSession) MToken() (common.Address, error) {
	return _ContractsContract.Contract.MToken(&_ContractsContract.CallOpts)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() constant returns(address)
func (_ContractsContract *ContractsContractCallerSession) MToken() (common.Address, error) {
	return _ContractsContract.Contract.MToken(&_ContractsContract.CallOpts)
}

// NetworkContract is a free data retrieval call binding the contract method 0xf842c213.
//
// Solidity: function networkContract() constant returns(address)
func (_ContractsContract *ContractsContractCaller) NetworkContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractsContract.contract.Call(opts, out, "networkContract")
	return *ret0, err
}

// NetworkContract is a free data retrieval call binding the contract method 0xf842c213.
//
// Solidity: function networkContract() constant returns(address)
func (_ContractsContract *ContractsContractSession) NetworkContract() (common.Address, error) {
	return _ContractsContract.Contract.NetworkContract(&_ContractsContract.CallOpts)
}

// NetworkContract is a free data retrieval call binding the contract method 0xf842c213.
//
// Solidity: function networkContract() constant returns(address)
func (_ContractsContract *ContractsContractCallerSession) NetworkContract() (common.Address, error) {
	return _ContractsContract.Contract.NetworkContract(&_ContractsContract.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() constant returns(address)
func (_ContractsContract *ContractsContractCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractsContract.contract.Call(opts, out, "priceOracle")
	return *ret0, err
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() constant returns(address)
func (_ContractsContract *ContractsContractSession) PriceOracle() (common.Address, error) {
	return _ContractsContract.Contract.PriceOracle(&_ContractsContract.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() constant returns(address)
func (_ContractsContract *ContractsContractCallerSession) PriceOracle() (common.Address, error) {
	return _ContractsContract.Contract.PriceOracle(&_ContractsContract.CallOpts)
}

// SetMToken is a paid mutator transaction binding the contract method 0xb882c4d8.
//
// Solidity: function setMToken(addr address) returns()
func (_ContractsContract *ContractsContractTransactor) SetMToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.contract.Transact(opts, "setMToken", addr)
}

// SetMToken is a paid mutator transaction binding the contract method 0xb882c4d8.
//
// Solidity: function setMToken(addr address) returns()
func (_ContractsContract *ContractsContractSession) SetMToken(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.SetMToken(&_ContractsContract.TransactOpts, addr)
}

// SetMToken is a paid mutator transaction binding the contract method 0xb882c4d8.
//
// Solidity: function setMToken(addr address) returns()
func (_ContractsContract *ContractsContractTransactorSession) SetMToken(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.SetMToken(&_ContractsContract.TransactOpts, addr)
}

// SetNetworkContract is a paid mutator transaction binding the contract method 0x599b9348.
//
// Solidity: function setNetworkContract(addr address) returns()
func (_ContractsContract *ContractsContractTransactor) SetNetworkContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.contract.Transact(opts, "setNetworkContract", addr)
}

// SetNetworkContract is a paid mutator transaction binding the contract method 0x599b9348.
//
// Solidity: function setNetworkContract(addr address) returns()
func (_ContractsContract *ContractsContractSession) SetNetworkContract(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.SetNetworkContract(&_ContractsContract.TransactOpts, addr)
}

// SetNetworkContract is a paid mutator transaction binding the contract method 0x599b9348.
//
// Solidity: function setNetworkContract(addr address) returns()
func (_ContractsContract *ContractsContractTransactorSession) SetNetworkContract(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.SetNetworkContract(&_ContractsContract.TransactOpts, addr)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(addr address) returns()
func (_ContractsContract *ContractsContractTransactor) SetPriceOracle(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.contract.Transact(opts, "setPriceOracle", addr)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(addr address) returns()
func (_ContractsContract *ContractsContractSession) SetPriceOracle(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.SetPriceOracle(&_ContractsContract.TransactOpts, addr)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(addr address) returns()
func (_ContractsContract *ContractsContractTransactorSession) SetPriceOracle(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.SetPriceOracle(&_ContractsContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_ContractsContract *ContractsContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_ContractsContract *ContractsContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.TransferOwnership(&_ContractsContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_ContractsContract *ContractsContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _ContractsContract.Contract.TransferOwnership(&_ContractsContract.TransactOpts, addr)
}
