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
const SystemVarsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"oracleDeduction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencySupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"uint256\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevCurrencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"uint256\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SystemVarsBin is the compiled bytecode used for deploying new contracts.
const SystemVarsBin = `608060405234801561001057600080fd5b50604051604080610586833981018060405281019080805190602001909291908051906020019092919050505081600181905550816002819055508060048190555080600381905550505061051c8061006a6000396000f3006080604052600436106100a3576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062bf32ca146100a8578063158ef93e146100e957806321873631146101185780632af4f9c0146101435780632d3802421461016e5780636df566d714610199578063a035b1fe146101c4578063b0c6363d146101ef578063e4a301161461021a578063fc634f4b14610251575b600080fd5b3480156100b457600080fd5b506100d36004803603810190808035906020019092919050505061027c565b6040518082815260200191505060405180910390f35b3480156100f557600080fd5b506100fe610295565b604051808215151515815260200191505060405180910390f35b34801561012457600080fd5b5061012d6102a7565b6040518082815260200191505060405180910390f35b34801561014f57600080fd5b506101586102d7565b6040518082815260200191505060405180910390f35b34801561017a57600080fd5b506101836102dd565b6040518082815260200191505060405180910390f35b3480156101a557600080fd5b506101ae610365565b6040518082815260200191505060405180910390f35b3480156101d057600080fd5b506101d961036b565b6040518082815260200191505060405180910390f35b3480156101fb57600080fd5b50610204610375565b6040518082815260200191505060405180910390f35b34801561022657600080fd5b5061024f600480360381019080803590602001909291908035906020019092919050505061037b565b005b34801561025d57600080fd5b5061026661045f565b6040518082815260200191505060405180910390f35b600060648260040281151561028d57fe5b049050919050565b6000809054906101000a900460ff1681565b60006102d2670de0b6b3a76400003073ffffffffffffffffffffffffffffffffffffffff1631610465565b905090565b60045481565b600080600180430114156102fc57680246ddf979766800009150610361565b61271060045481151561030b57fe5b04905060015460025411801561032a5750670de0b6b3a7640000600154115b1561034b57610344816004540161033f61047e565b610465565b9150610361565b61035e816004540364e8d4a510006104bf565b91505b5090565b60025481565b6000600254905090565b60035481565b6000809054906101000a900460ff16151515610425576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b8160018190555081600281905550806004819055508060038190555060016000806101000a81548160ff0219169083151502179055505050565b60015481565b60008183106104745781610476565b825b905092915050565b6000600180430111801561049657506104956104d9565b5b6104a957680471fa858b9e0800006104ba565b6127106003548115156104b857fe5b045b905090565b6000818310156104cf57816104d1565b825b905092915050565b600069d3c21bcecceda100000060035410159050905600a165627a7a723058206fe4f1f991942cb71ef574c8d7fffd57bd4a05e408f7df2f2c0ecd8fe40406390029`

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
