// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// PriceOracleContractABI is the input ABI used to generate the binding from.
const PriceOracleContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fiatAmount\",\"type\":\"uint256\"}],\"name\":\"priceForFiat\",\"outputs\":[{\"name\":\"_cryptoAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cryptoAmount\",\"type\":\"uint256\"}],\"name\":\"priceForCrypto\",\"outputs\":[{\"name\":\"_fiatAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cryptoAmount\",\"type\":\"uint256\"},{\"name\":\"_fiatAmount\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoName\",\"type\":\"string\"},{\"name\":\"_cryptoSymbol\",\"type\":\"string\"},{\"name\":\"_cryptoDecimals\",\"type\":\"uint8\"},{\"name\":\"_cryptoAmount\",\"type\":\"uint256\"},{\"name\":\"_fiatName\",\"type\":\"string\"},{\"name\":\"_fiatSymbol\",\"type\":\"string\"},{\"name\":\"_fiatDecimals\",\"type\":\"uint8\"},{\"name\":\"_fiatAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cryptoPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fiatPrice\",\"type\":\"uint256\"}],\"name\":\"NewPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// PriceOracleContractBin is the compiled bytecode used for deploying new contracts.
const PriceOracleContractBin = `6060604052341561000f57600080fd5b604051610b45380380610b4583398101604052808051820191906020018051820191906020018051906020019091908051906020019091908051820191906020018051820191906020018051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555087600190805190602001906100c392919061015a565b5086600290805190602001906100da92919061015a565b5085600360006101000a81548160ff021916908360ff16021790555084600781905550836004908051906020019061011392919061015a565b50826005908051906020019061012a92919061015a565b5081600660006101000a81548160ff021916908360ff1602179055508060088190555050505050505050506101ff565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061019b57805160ff19168380011785556101c9565b828001600101855582156101c9579182015b828111156101c85782518255916020019190600101906101ad565b5b5090506101d691906101da565b5090565b6101fc91905b808211156101f85760008160009055506001016101e0565b5090565b90565b6109378061020e6000396000f3006060604052600436106100a3576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680624c8e14146100a85780630c272196146100d7578063126d413c1461010e57806327631d021461019c5780632af9501f1461022a5780636857e1c5146102b8578063c9ebff92146102ef578063f2fde38b1461037d578063f5890e46146103b6578063f7d97577146103e5575b600080fd5b34156100b357600080fd5b6100bb610429565b604051808260ff1660ff16815260200191505060405180910390f35b34156100e257600080fd5b6100f8600480803590602001909190505061043c565b6040518082815260200191505060405180910390f35b341561011957600080fd5b610121610457565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610161578082015181840152602081019050610146565b50505050905090810190601f16801561018e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101a757600080fd5b6101af6104f5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101ef5780820151818401526020810190506101d4565b50505050905090810190601f16801561021c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561023557600080fd5b61023d610593565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561027d578082015181840152602081019050610262565b50505050905090810190601f1680156102aa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102c357600080fd5b6102d96004808035906020019091905050610631565b6040518082815260200191505060405180910390f35b34156102fa57600080fd5b61030261064c565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610342578082015181840152602081019050610327565b50505050905090810190601f16801561036f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561038857600080fd5b6103b4600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506106ea565b005b34156103c157600080fd5b6103c9610840565b604051808260ff1660ff16815260200191505060405180910390f35b34156103f057600080fd5b61040f6004808035906020019091908035906020019091905050610853565b604051808215151515815260200191505060405180910390f35b600360009054906101000a900460ff1681565b6000600854600754830281151561044f57fe5b049050919050565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104ed5780601f106104c2576101008083540402835291602001916104ed565b820191906000526020600020905b8154815290600101906020018083116104d057829003601f168201915b505050505081565b60058054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561058b5780601f106105605761010080835404028352916020019161058b565b820191906000526020600020905b81548152906001019060200180831161056e57829003601f168201915b505050505081565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106295780601f106105fe57610100808354040283529160200191610629565b820191906000526020600020905b81548152906001019060200180831161060c57829003601f168201915b505050505081565b6000600754600854830281151561064457fe5b049050919050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106e25780601f106106b7576101008083540402835291602001916106e2565b820191906000526020600020905b8154815290600101906020018083116106c557829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561074557600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b600660009054906101000a900460ff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108b057600080fd5b82600781905550816008819055507fb9362b96e28efbb7a7e63bb4a97faf9924ec0394635feff8588a6ae2a5f784fe600754600854604051808381526020018281526020019250505060405180910390a160019050929150505600a165627a7a723058201907ada9794bf4a046040bc09af5c9b4a4a8e21bb7a56a667000971c96a11d450029`

// DeployPriceOracleContract deploys a new Ethereum contract, binding an instance of PriceOracleContract to it.
func DeployPriceOracleContract(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoName string, _cryptoSymbol string, _cryptoDecimals uint8, _cryptoAmount *big.Int, _fiatName string, _fiatSymbol string, _fiatDecimals uint8, _fiatAmount *big.Int) (common.Address, *types.Transaction, *PriceOracleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOracleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceOracleContractBin), backend, _cryptoName, _cryptoSymbol, _cryptoDecimals, _cryptoAmount, _fiatName, _fiatSymbol, _fiatDecimals, _fiatAmount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceOracleContract{PriceOracleContractCaller: PriceOracleContractCaller{contract: contract}, PriceOracleContractTransactor: PriceOracleContractTransactor{contract: contract}}, nil
}

// PriceOracleContract is an auto generated Go binding around an Ethereum contract.
type PriceOracleContract struct {
	PriceOracleContractCaller     // Read-only binding to the contract
	PriceOracleContractTransactor // Write-only binding to the contract
}

// PriceOracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOracleContractSession struct {
	Contract     *PriceOracleContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PriceOracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOracleContractCallerSession struct {
	Contract *PriceOracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PriceOracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOracleContractTransactorSession struct {
	Contract     *PriceOracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PriceOracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOracleContractRaw struct {
	Contract *PriceOracleContract // Generic contract binding to access the raw methods on
}

// PriceOracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOracleContractCallerRaw struct {
	Contract *PriceOracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// PriceOracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOracleContractTransactorRaw struct {
	Contract *PriceOracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOracleContract creates a new instance of PriceOracleContract, bound to a specific deployed contract.
func NewPriceOracleContract(address common.Address, backend bind.ContractBackend) (*PriceOracleContract, error) {
	contract, err := bindPriceOracleContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOracleContract{PriceOracleContractCaller: PriceOracleContractCaller{contract: contract}, PriceOracleContractTransactor: PriceOracleContractTransactor{contract: contract}}, nil
}

// NewPriceOracleContractCaller creates a new read-only instance of PriceOracleContract, bound to a specific deployed contract.
func NewPriceOracleContractCaller(address common.Address, caller bind.ContractCaller) (*PriceOracleContractCaller, error) {
	contract, err := bindPriceOracleContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleContractCaller{contract: contract}, nil
}

// NewPriceOracleContractTransactor creates a new write-only instance of PriceOracleContract, bound to a specific deployed contract.
func NewPriceOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceOracleContractTransactor, error) {
	contract, err := bindPriceOracleContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PriceOracleContractTransactor{contract: contract}, nil
}

// bindPriceOracleContract binds a generic wrapper to an already deployed contract.
func bindPriceOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOracleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracleContract *PriceOracleContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceOracleContract.Contract.PriceOracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracleContract *PriceOracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.PriceOracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracleContract *PriceOracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.PriceOracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracleContract *PriceOracleContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceOracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracleContract *PriceOracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracleContract *PriceOracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.contract.Transact(opts, method, params...)
}

// CryptoDecimals is a free data retrieval call binding the contract method 0x004c8e14.
//
// Solidity: function cryptoDecimals() constant returns(uint8)
func (_PriceOracleContract *PriceOracleContractCaller) CryptoDecimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "cryptoDecimals")
	return *ret0, err
}

// CryptoDecimals is a free data retrieval call binding the contract method 0x004c8e14.
//
// Solidity: function cryptoDecimals() constant returns(uint8)
func (_PriceOracleContract *PriceOracleContractSession) CryptoDecimals() (uint8, error) {
	return _PriceOracleContract.Contract.CryptoDecimals(&_PriceOracleContract.CallOpts)
}

// CryptoDecimals is a free data retrieval call binding the contract method 0x004c8e14.
//
// Solidity: function cryptoDecimals() constant returns(uint8)
func (_PriceOracleContract *PriceOracleContractCallerSession) CryptoDecimals() (uint8, error) {
	return _PriceOracleContract.Contract.CryptoDecimals(&_PriceOracleContract.CallOpts)
}

// CryptoName is a free data retrieval call binding the contract method 0xc9ebff92.
//
// Solidity: function cryptoName() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCaller) CryptoName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "cryptoName")
	return *ret0, err
}

// CryptoName is a free data retrieval call binding the contract method 0xc9ebff92.
//
// Solidity: function cryptoName() constant returns(string)
func (_PriceOracleContract *PriceOracleContractSession) CryptoName() (string, error) {
	return _PriceOracleContract.Contract.CryptoName(&_PriceOracleContract.CallOpts)
}

// CryptoName is a free data retrieval call binding the contract method 0xc9ebff92.
//
// Solidity: function cryptoName() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCallerSession) CryptoName() (string, error) {
	return _PriceOracleContract.Contract.CryptoName(&_PriceOracleContract.CallOpts)
}

// CryptoSymbol is a free data retrieval call binding the contract method 0x2af9501f.
//
// Solidity: function cryptoSymbol() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCaller) CryptoSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "cryptoSymbol")
	return *ret0, err
}

// CryptoSymbol is a free data retrieval call binding the contract method 0x2af9501f.
//
// Solidity: function cryptoSymbol() constant returns(string)
func (_PriceOracleContract *PriceOracleContractSession) CryptoSymbol() (string, error) {
	return _PriceOracleContract.Contract.CryptoSymbol(&_PriceOracleContract.CallOpts)
}

// CryptoSymbol is a free data retrieval call binding the contract method 0x2af9501f.
//
// Solidity: function cryptoSymbol() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCallerSession) CryptoSymbol() (string, error) {
	return _PriceOracleContract.Contract.CryptoSymbol(&_PriceOracleContract.CallOpts)
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint8)
func (_PriceOracleContract *PriceOracleContractCaller) FiatDecimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "fiatDecimals")
	return *ret0, err
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint8)
func (_PriceOracleContract *PriceOracleContractSession) FiatDecimals() (uint8, error) {
	return _PriceOracleContract.Contract.FiatDecimals(&_PriceOracleContract.CallOpts)
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint8)
func (_PriceOracleContract *PriceOracleContractCallerSession) FiatDecimals() (uint8, error) {
	return _PriceOracleContract.Contract.FiatDecimals(&_PriceOracleContract.CallOpts)
}

// FiatName is a free data retrieval call binding the contract method 0x126d413c.
//
// Solidity: function fiatName() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCaller) FiatName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "fiatName")
	return *ret0, err
}

// FiatName is a free data retrieval call binding the contract method 0x126d413c.
//
// Solidity: function fiatName() constant returns(string)
func (_PriceOracleContract *PriceOracleContractSession) FiatName() (string, error) {
	return _PriceOracleContract.Contract.FiatName(&_PriceOracleContract.CallOpts)
}

// FiatName is a free data retrieval call binding the contract method 0x126d413c.
//
// Solidity: function fiatName() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCallerSession) FiatName() (string, error) {
	return _PriceOracleContract.Contract.FiatName(&_PriceOracleContract.CallOpts)
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCaller) FiatSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "fiatSymbol")
	return *ret0, err
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_PriceOracleContract *PriceOracleContractSession) FiatSymbol() (string, error) {
	return _PriceOracleContract.Contract.FiatSymbol(&_PriceOracleContract.CallOpts)
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_PriceOracleContract *PriceOracleContractCallerSession) FiatSymbol() (string, error) {
	return _PriceOracleContract.Contract.FiatSymbol(&_PriceOracleContract.CallOpts)
}

// PriceForCrypto is a free data retrieval call binding the contract method 0x6857e1c5.
//
// Solidity: function priceForCrypto(_cryptoAmount uint256) constant returns(_fiatAmount uint256)
func (_PriceOracleContract *PriceOracleContractCaller) PriceForCrypto(opts *bind.CallOpts, _cryptoAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "priceForCrypto", _cryptoAmount)
	return *ret0, err
}

// PriceForCrypto is a free data retrieval call binding the contract method 0x6857e1c5.
//
// Solidity: function priceForCrypto(_cryptoAmount uint256) constant returns(_fiatAmount uint256)
func (_PriceOracleContract *PriceOracleContractSession) PriceForCrypto(_cryptoAmount *big.Int) (*big.Int, error) {
	return _PriceOracleContract.Contract.PriceForCrypto(&_PriceOracleContract.CallOpts, _cryptoAmount)
}

// PriceForCrypto is a free data retrieval call binding the contract method 0x6857e1c5.
//
// Solidity: function priceForCrypto(_cryptoAmount uint256) constant returns(_fiatAmount uint256)
func (_PriceOracleContract *PriceOracleContractCallerSession) PriceForCrypto(_cryptoAmount *big.Int) (*big.Int, error) {
	return _PriceOracleContract.Contract.PriceForCrypto(&_PriceOracleContract.CallOpts, _cryptoAmount)
}

// PriceForFiat is a free data retrieval call binding the contract method 0x0c272196.
//
// Solidity: function priceForFiat(_fiatAmount uint256) constant returns(_cryptoAmount uint256)
func (_PriceOracleContract *PriceOracleContractCaller) PriceForFiat(opts *bind.CallOpts, _fiatAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracleContract.contract.Call(opts, out, "priceForFiat", _fiatAmount)
	return *ret0, err
}

// PriceForFiat is a free data retrieval call binding the contract method 0x0c272196.
//
// Solidity: function priceForFiat(_fiatAmount uint256) constant returns(_cryptoAmount uint256)
func (_PriceOracleContract *PriceOracleContractSession) PriceForFiat(_fiatAmount *big.Int) (*big.Int, error) {
	return _PriceOracleContract.Contract.PriceForFiat(&_PriceOracleContract.CallOpts, _fiatAmount)
}

// PriceForFiat is a free data retrieval call binding the contract method 0x0c272196.
//
// Solidity: function priceForFiat(_fiatAmount uint256) constant returns(_cryptoAmount uint256)
func (_PriceOracleContract *PriceOracleContractCallerSession) PriceForFiat(_fiatAmount *big.Int) (*big.Int, error) {
	return _PriceOracleContract.Contract.PriceForFiat(&_PriceOracleContract.CallOpts, _fiatAmount)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(_cryptoAmount uint256, _fiatAmount uint256) returns(success bool)
func (_PriceOracleContract *PriceOracleContractTransactor) SetPrice(opts *bind.TransactOpts, _cryptoAmount *big.Int, _fiatAmount *big.Int) (*types.Transaction, error) {
	return _PriceOracleContract.contract.Transact(opts, "setPrice", _cryptoAmount, _fiatAmount)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(_cryptoAmount uint256, _fiatAmount uint256) returns(success bool)
func (_PriceOracleContract *PriceOracleContractSession) SetPrice(_cryptoAmount *big.Int, _fiatAmount *big.Int) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.SetPrice(&_PriceOracleContract.TransactOpts, _cryptoAmount, _fiatAmount)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(_cryptoAmount uint256, _fiatAmount uint256) returns(success bool)
func (_PriceOracleContract *PriceOracleContractTransactorSession) SetPrice(_cryptoAmount *big.Int, _fiatAmount *big.Int) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.SetPrice(&_PriceOracleContract.TransactOpts, _cryptoAmount, _fiatAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_PriceOracleContract *PriceOracleContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PriceOracleContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_PriceOracleContract *PriceOracleContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.TransferOwnership(&_PriceOracleContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_PriceOracleContract *PriceOracleContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _PriceOracleContract.Contract.TransferOwnership(&_PriceOracleContract.TransactOpts, addr)
}
