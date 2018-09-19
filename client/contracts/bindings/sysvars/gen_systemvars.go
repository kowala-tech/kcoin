// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sysvars

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// SystemVarsABI is the input ABI used to generate the binding from.
const SystemVarsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"oracleDeduction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stabilityLevel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencySupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"uint256\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevCurrencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"uint256\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SystemVarsBin is the compiled bytecode used for deploying new contracts.
const SystemVarsBin = `608060405234801561001057600080fd5b506040516040806105c283398101806040528101908080519060200190929190805190602001909291905050508160018190555081600281905550806004819055508060038190555050506105588061006a6000396000f3006080604052600436106100ae576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062bf32ca146100b357806311bd50e6146100f4578063158ef93e1461011f578063218736311461014e5780632af4f9c0146101795780632d380242146101a45780636df566d7146101cf578063a035b1fe146101fa578063b0c6363d14610225578063e4a3011614610250578063fc634f4b14610287575b600080fd5b3480156100bf57600080fd5b506100de600480360381019080803590602001909291905050506102b2565b6040518082815260200191505060405180910390f35b34801561010057600080fd5b506101096102cb565b6040518082815260200191505060405180910390f35b34801561012b57600080fd5b506101346102d1565b604051808215151515815260200191505060405180910390f35b34801561015a57600080fd5b506101636102e3565b6040518082815260200191505060405180910390f35b34801561018557600080fd5b5061018e610313565b6040518082815260200191505060405180910390f35b3480156101b057600080fd5b506101b9610319565b6040518082815260200191505060405180910390f35b3480156101db57600080fd5b506101e46103a1565b6040518082815260200191505060405180910390f35b34801561020657600080fd5b5061020f6103a7565b6040518082815260200191505060405180910390f35b34801561023157600080fd5b5061023a6103b1565b6040518082815260200191505060405180910390f35b34801561025c57600080fd5b5061028560048036038101908080359060200190929190803590602001909291905050506103b7565b005b34801561029357600080fd5b5061029c61049b565b6040518082815260200191505060405180910390f35b60006064826004028115156102c357fe5b049050919050565b60055481565b6000809054906101000a900460ff1681565b600061030e670de0b6b3a76400003073ffffffffffffffffffffffffffffffffffffffff16316104a1565b905090565b60045481565b6000806001804301141561033857680246ddf97976680000915061039d565b61271060045481151561034757fe5b0490506001546002541180156103665750670de0b6b3a7640000600154115b1561038757610380816004540161037b6104ba565b6104a1565b915061039d565b61039a816004540364e8d4a510006104fb565b91505b5090565b60025481565b6000600254905090565b60035481565b6000809054906101000a900460ff16151515610461576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b8160018190555081600281905550806004819055508060038190555060016000806101000a81548160ff0219169083151502179055505050565b60015481565b60008183106104b057816104b2565b825b905092915050565b600060018043011180156104d257506104d1610515565b5b6104e557680471fa858b9e0800006104f6565b6127106003548115156104f457fe5b045b905090565b60008183101561050b578161050d565b825b905092915050565b600069d3c21bcecceda100000060035410159050905600a165627a7a723058200fb36c618f960ebd7d69d9dcb10a3d437d194087c5ea2c7d98e179a6ddff14040029`

// DeploySystemVars deploys a new Kowala contract, binding an instance of SystemVars to it.
func DeploySystemVars(auth *bind.TransactOpts, backend bind.ContractBackend, _initialPrice *big.Int, _initialSupply *big.Int) (common.Address, *types.Transaction, *SystemVars, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemVarsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemVarsBin), backend, _initialPrice, _initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemVars{SystemVarsCaller: SystemVarsCaller{contract: contract}, SystemVarsTransactor: SystemVarsTransactor{contract: contract}, SystemVarsFilterer: SystemVarsFilterer{contract: contract}}, nil
}

// SystemVars is an auto generated Go binding around a Kowala contract.
type SystemVars struct {
	SystemVarsCaller     // Read-only binding to the contract
	SystemVarsTransactor // Write-only binding to the contract
	SystemVarsFilterer   // Log filterer for contract events
}

// SystemVarsCaller is an auto generated read-only Go binding around a Kowala contract.
type SystemVarsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemVarsTransactor is an auto generated write-only Go binding around a Kowala contract.
type SystemVarsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemVarsFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type SystemVarsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemVarsSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type SystemVarsSession struct {
	Contract     *SystemVars       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemVarsCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type SystemVarsCallerSession struct {
	Contract *SystemVarsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SystemVarsTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type SystemVarsTransactorSession struct {
	Contract     *SystemVarsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SystemVarsRaw is an auto generated low-level Go binding around a Kowala contract.
type SystemVarsRaw struct {
	Contract *SystemVars // Generic contract binding to access the raw methods on
}

// SystemVarsCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type SystemVarsCallerRaw struct {
	Contract *SystemVarsCaller // Generic read-only contract binding to access the raw methods on
}

// SystemVarsTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type SystemVarsTransactorRaw struct {
	Contract *SystemVarsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemVars creates a new instance of SystemVars, bound to a specific deployed contract.
func NewSystemVars(address common.Address, backend bind.ContractBackend) (*SystemVars, error) {
	contract, err := bindSystemVars(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemVars{SystemVarsCaller: SystemVarsCaller{contract: contract}, SystemVarsTransactor: SystemVarsTransactor{contract: contract}, SystemVarsFilterer: SystemVarsFilterer{contract: contract}}, nil
}

// NewSystemVarsCaller creates a new read-only instance of SystemVars, bound to a specific deployed contract.
func NewSystemVarsCaller(address common.Address, caller bind.ContractCaller) (*SystemVarsCaller, error) {
	contract, err := bindSystemVars(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemVarsCaller{contract: contract}, nil
}

// NewSystemVarsTransactor creates a new write-only instance of SystemVars, bound to a specific deployed contract.
func NewSystemVarsTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemVarsTransactor, error) {
	contract, err := bindSystemVars(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemVarsTransactor{contract: contract}, nil
}

// NewSystemVarsFilterer creates a new log filterer instance of SystemVars, bound to a specific deployed contract.
func NewSystemVarsFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemVarsFilterer, error) {
	contract, err := bindSystemVars(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemVarsFilterer{contract: contract}, nil
}

// bindSystemVars binds a generic wrapper to an already deployed contract.
func bindSystemVars(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemVarsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemVars *SystemVarsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemVars.Contract.SystemVarsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemVars *SystemVarsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemVars.Contract.SystemVarsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemVars *SystemVarsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemVars.Contract.SystemVarsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemVars *SystemVarsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemVars.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemVars *SystemVarsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemVars.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemVars *SystemVarsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemVars.Contract.contract.Transact(opts, method, params...)
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) CurrencyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "currencyPrice")
	return *ret0, err
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsSession) CurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.CurrencyPrice(&_SystemVars.CallOpts)
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) CurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.CurrencyPrice(&_SystemVars.CallOpts)
}

// CurrencySupply is a free data retrieval call binding the contract method 0xb0c6363d.
//
// Solidity: function currencySupply() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) CurrencySupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "currencySupply")
	return *ret0, err
}

// CurrencySupply is a free data retrieval call binding the contract method 0xb0c6363d.
//
// Solidity: function currencySupply() constant returns(uint256)
func (_SystemVars *SystemVarsSession) CurrencySupply() (*big.Int, error) {
	return _SystemVars.Contract.CurrencySupply(&_SystemVars.CallOpts)
}

// CurrencySupply is a free data retrieval call binding the contract method 0xb0c6363d.
//
// Solidity: function currencySupply() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) CurrencySupply() (*big.Int, error) {
	return _SystemVars.Contract.CurrencySupply(&_SystemVars.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_SystemVars *SystemVarsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_SystemVars *SystemVarsSession) Initialized() (bool, error) {
	return _SystemVars.Contract.Initialized(&_SystemVars.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_SystemVars *SystemVarsCallerSession) Initialized() (bool, error) {
	return _SystemVars.Contract.Initialized(&_SystemVars.CallOpts)
}

// MintedAmount is a free data retrieval call binding the contract method 0x2d380242.
//
// Solidity: function mintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) MintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "mintedAmount")
	return *ret0, err
}

// MintedAmount is a free data retrieval call binding the contract method 0x2d380242.
//
// Solidity: function mintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsSession) MintedAmount() (*big.Int, error) {
	return _SystemVars.Contract.MintedAmount(&_SystemVars.CallOpts)
}

// MintedAmount is a free data retrieval call binding the contract method 0x2d380242.
//
// Solidity: function mintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) MintedAmount() (*big.Int, error) {
	return _SystemVars.Contract.MintedAmount(&_SystemVars.CallOpts)
}

// MintedReward is a free data retrieval call binding the contract method 0x2af4f9c0.
//
// Solidity: function mintedReward() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) MintedReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "mintedReward")
	return *ret0, err
}

// MintedReward is a free data retrieval call binding the contract method 0x2af4f9c0.
//
// Solidity: function mintedReward() constant returns(uint256)
func (_SystemVars *SystemVarsSession) MintedReward() (*big.Int, error) {
	return _SystemVars.Contract.MintedReward(&_SystemVars.CallOpts)
}

// MintedReward is a free data retrieval call binding the contract method 0x2af4f9c0.
//
// Solidity: function mintedReward() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) MintedReward() (*big.Int, error) {
	return _SystemVars.Contract.MintedReward(&_SystemVars.CallOpts)
}

// OracleDeduction is a free data retrieval call binding the contract method 0x00bf32ca.
//
// Solidity: function oracleDeduction(mintedAmount uint256) constant returns(uint256)
func (_SystemVars *SystemVarsCaller) OracleDeduction(opts *bind.CallOpts, mintedAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "oracleDeduction", mintedAmount)
	return *ret0, err
}

// OracleDeduction is a free data retrieval call binding the contract method 0x00bf32ca.
//
// Solidity: function oracleDeduction(mintedAmount uint256) constant returns(uint256)
func (_SystemVars *SystemVarsSession) OracleDeduction(mintedAmount *big.Int) (*big.Int, error) {
	return _SystemVars.Contract.OracleDeduction(&_SystemVars.CallOpts, mintedAmount)
}

// OracleDeduction is a free data retrieval call binding the contract method 0x00bf32ca.
//
// Solidity: function oracleDeduction(mintedAmount uint256) constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) OracleDeduction(mintedAmount *big.Int) (*big.Int, error) {
	return _SystemVars.Contract.OracleDeduction(&_SystemVars.CallOpts, mintedAmount)
}

// OracleReward is a free data retrieval call binding the contract method 0x21873631.
//
// Solidity: function oracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) OracleReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "oracleReward")
	return *ret0, err
}

// OracleReward is a free data retrieval call binding the contract method 0x21873631.
//
// Solidity: function oracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsSession) OracleReward() (*big.Int, error) {
	return _SystemVars.Contract.OracleReward(&_SystemVars.CallOpts)
}

// OracleReward is a free data retrieval call binding the contract method 0x21873631.
//
// Solidity: function oracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) OracleReward() (*big.Int, error) {
	return _SystemVars.Contract.OracleReward(&_SystemVars.CallOpts)
}

// PrevCurrencyPrice is a free data retrieval call binding the contract method 0xfc634f4b.
//
// Solidity: function prevCurrencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) PrevCurrencyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "prevCurrencyPrice")
	return *ret0, err
}

// PrevCurrencyPrice is a free data retrieval call binding the contract method 0xfc634f4b.
//
// Solidity: function prevCurrencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsSession) PrevCurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.PrevCurrencyPrice(&_SystemVars.CallOpts)
}

// PrevCurrencyPrice is a free data retrieval call binding the contract method 0xfc634f4b.
//
// Solidity: function prevCurrencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) PrevCurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.PrevCurrencyPrice(&_SystemVars.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(price uint256)
func (_SystemVars *SystemVarsCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(price uint256)
func (_SystemVars *SystemVarsSession) Price() (*big.Int, error) {
	return _SystemVars.Contract.Price(&_SystemVars.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(price uint256)
func (_SystemVars *SystemVarsCallerSession) Price() (*big.Int, error) {
	return _SystemVars.Contract.Price(&_SystemVars.CallOpts)
}

// StabilityLevel is a free data retrieval call binding the contract method 0x11bd50e6.
//
// Solidity: function stabilityLevel() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) StabilityLevel(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "stabilityLevel")
	return *ret0, err
}

// StabilityLevel is a free data retrieval call binding the contract method 0x11bd50e6.
//
// Solidity: function stabilityLevel() constant returns(uint256)
func (_SystemVars *SystemVarsSession) StabilityLevel() (*big.Int, error) {
	return _SystemVars.Contract.StabilityLevel(&_SystemVars.CallOpts)
}

// StabilityLevel is a free data retrieval call binding the contract method 0x11bd50e6.
//
// Solidity: function stabilityLevel() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) StabilityLevel() (*big.Int, error) {
	return _SystemVars.Contract.StabilityLevel(&_SystemVars.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(_initialPrice uint256, _initialSupply uint256) returns()
func (_SystemVars *SystemVarsTransactor) Initialize(opts *bind.TransactOpts, _initialPrice *big.Int, _initialSupply *big.Int) (*types.Transaction, error) {
	return _SystemVars.contract.Transact(opts, "initialize", _initialPrice, _initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(_initialPrice uint256, _initialSupply uint256) returns()
func (_SystemVars *SystemVarsSession) Initialize(_initialPrice *big.Int, _initialSupply *big.Int) (*types.Transaction, error) {
	return _SystemVars.Contract.Initialize(&_SystemVars.TransactOpts, _initialPrice, _initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(_initialPrice uint256, _initialSupply uint256) returns()
func (_SystemVars *SystemVarsTransactorSession) Initialize(_initialPrice *big.Int, _initialSupply *big.Int) (*types.Transaction, error) {
	return _SystemVars.Contract.Initialize(&_SystemVars.TransactOpts, _initialPrice, _initialSupply)
}
