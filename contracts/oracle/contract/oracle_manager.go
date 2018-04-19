// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
)

// OracleManagerContractABI is the input ABI used to generate the binding from.
const OracleManagerContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"_isOracle\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSubmissionAtIndex\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerOracle\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"submitPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSubmissionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxOracles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxOracles\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// OracleManagerContractBin is the compiled bytecode used for deploying new contracts.
const OracleManagerContractBin = `6060604052341561000f57600080fd5b6040516060806110e383398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001821015151561008d57600080fd5b670de0b6b3a7640000830260018190555081600281905550620151808102600381905550505050611020806100c36000396000f3006060604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf142146100d25780630a3cb663146100fb578063252f7be9146101245780633232f10814610175578063339d2590146101df5780634b2c89d5146101e957806369474625146101fe578063893d20e81461022757806397584b3e1461027c578063986fcbe9146102a95780639999d2ae146102cc578063c0d2c49d146102f5578063f2fde38b1461031e578063f93a2eb21461036f575b005b34156100dd57600080fd5b6100e5610384565b6040518082815260200191505060405180910390f35b341561010657600080fd5b61010e610458565b6040518082815260200191505060405180910390f35b341561012f57600080fd5b61015b600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061045e565b604051808215151515815260200191505060405180910390f35b341561018057600080fd5b61019660048080359060200190919050506104b7565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b6101e761050f565b005b34156101f457600080fd5b6101fc61055d565b005b341561020957600080fd5b61021161069c565b6040518082815260200191505060405180910390f35b341561023257600080fd5b61023a6106a2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561028757600080fd5b61028f6106cb565b604051808215151515815260200191505060405180910390f35b34156102b457600080fd5b6102ca60048080359060200190919050506106de565b005b34156102d757600080fd5b6102df61079d565b6040518082815260200191505060405180910390f35b341561030057600080fd5b6103086107aa565b6040518082815260200191505060405180910390f35b341561032957600080fd5b610355600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506107b0565b604051808215151515815260200191505060405180910390f35b341561037a57600080fd5b61038261088d565b005b60008061038f6106cb565b1561039e576001549150610454565b6004600060056001600580549050038154811015156103b957fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561043e57fe5b9060005260206000209060020201600001540191505b5090565b60035481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b60008060006006848154811015156104cb57fe5b906000526020600020906002020190508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692508060010154915050915091565b6105183361045e565b15151561052457600080fd5b61052c610384565b341015151561053a57600080fd5b6105426106cb565b1515610551576105506108ac565b5b61055b33346108f9565b565b600080600080925060009150600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b8080549050821080156105e15750600081838154811015156105cc57fe5b90600052602060002090600202016001015414155b156106435780828154811015156105f457fe5b90600052602060002090600202016001015442101561061257610643565b808281548110151561062057fe5b9060005260206000209060020201600001548301925081806001019250506105ae565b61064d3383610c0b565b6000831115610697573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561069657600080fd5b5b505050565b60015481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000806005805490506002540311905090565b6106e73361045e565b15156106f257600080fd5b600680548060010182816107069190610e64565b9160005260206000209060020201600060408051908101604052803373ffffffffffffffffffffffffffffffffffffffff16815260200185815250909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015550505050565b6000600680549050905090565b60025481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561080d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151561088457816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60019050919050565b6108963361045e565b15156108a157600080fd5b6108aa33610cf8565b565b6108f760056001600580549050038154811015156108c657fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610cf8565b565b600080600080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002093506001600580548060010182816109569190610e96565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff0219169083151502179055508360020180548060010182816109e09190610ec2565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115610c035760046000600560018603815481101515610a4957fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816002016001836002018054905003815481101515610acc57fe5b90600052602060002090600202019050806000015485111515610aee57610c03565b600560018403815481101515610b0057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600584815481101515610b3b57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600560018503815481101515610b9757fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050610a2a565b505050505050565b600080600080841415610c1d57610cf1565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b8260020180549050811015610cdf578260020181815481101515610c8657fe5b90600052602060002090600202018360020183815481101515610ca557fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050610c66565b818360020181610cef9190610ef4565b505b5050505050565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160058054905003811015610df757600560018201815481101515610d6657fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600582815481101515610da157fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050610d44565b6005805480919060019003610e0c9190610f26565b5060008260010160006101000a81548160ff0219169083151502179055506003544201826002016001846002018054905003815481101515610e4a57fe5b906000526020600020906002020160010181905550505050565b815481835581811511610e9157600202816002028360005260206000209182019101610e909190610f52565b5b505050565b815481835581811511610ebd57818360005260206000209182019101610ebc9190610fa0565b5b505050565b815481835581811511610eef57600202816002028360005260206000209182019101610eee9190610fc5565b5b505050565b815481835581811511610f2157600202816002028360005260206000209182019101610f209190610fc5565b5b505050565b815481835581811511610f4d57818360005260206000209182019101610f4c9190610fa0565b5b505050565b610f9d91905b80821115610f9957600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182016000905550600201610f58565b5090565b90565b610fc291905b80821115610fbe576000816000905550600101610fa6565b5090565b90565b610ff191905b80821115610fed57600080820160009055600182016000905550600201610fcb565b5090565b905600a165627a7a723058209dff0e6acc5cdbb6c4f68d0fd0c19c2d4f04cd83c1154acb5f0aa30611f6670c0029`

// DeployOracleManagerContract deploys a new Ethereum contract, binding an instance of OracleManagerContract to it.
func DeployOracleManagerContract(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxOracles *big.Int, _freezePeriod *big.Int) (common.Address, *types.Transaction, *OracleManagerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleManagerContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleManagerContractBin), backend, _baseDeposit, _maxOracles, _freezePeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleManagerContract{OracleManagerContractCaller: OracleManagerContractCaller{contract: contract}, OracleManagerContractTransactor: OracleManagerContractTransactor{contract: contract}}, nil
}

// OracleManagerContract is an auto generated Go binding around an Ethereum contract.
type OracleManagerContract struct {
	OracleManagerContractCaller     // Read-only binding to the contract
	OracleManagerContractTransactor // Write-only binding to the contract
}

// OracleManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleManagerContractSession struct {
	Contract     *OracleManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OracleManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleManagerContractCallerSession struct {
	Contract *OracleManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OracleManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleManagerContractTransactorSession struct {
	Contract     *OracleManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OracleManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleManagerContractRaw struct {
	Contract *OracleManagerContract // Generic contract binding to access the raw methods on
}

// OracleManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleManagerContractCallerRaw struct {
	Contract *OracleManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleManagerContractTransactorRaw struct {
	Contract *OracleManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleManagerContract creates a new instance of OracleManagerContract, bound to a specific deployed contract.
func NewOracleManagerContract(address common.Address, backend bind.ContractBackend) (*OracleManagerContract, error) {
	contract, err := bindOracleManagerContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleManagerContract{OracleManagerContractCaller: OracleManagerContractCaller{contract: contract}, OracleManagerContractTransactor: OracleManagerContractTransactor{contract: contract}}, nil
}

// NewOracleManagerContractCaller creates a new read-only instance of OracleManagerContract, bound to a specific deployed contract.
func NewOracleManagerContractCaller(address common.Address, caller bind.ContractCaller) (*OracleManagerContractCaller, error) {
	contract, err := bindOracleManagerContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OracleManagerContractCaller{contract: contract}, nil
}

// NewOracleManagerContractTransactor creates a new write-only instance of OracleManagerContract, bound to a specific deployed contract.
func NewOracleManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleManagerContractTransactor, error) {
	contract, err := bindOracleManagerContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OracleManagerContractTransactor{contract: contract}, nil
}

// bindOracleManagerContract binds a generic wrapper to an already deployed contract.
func bindOracleManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleManagerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleManagerContract *OracleManagerContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleManagerContract.Contract.OracleManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleManagerContract *OracleManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.OracleManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleManagerContract *OracleManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.OracleManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleManagerContract *OracleManagerContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleManagerContract *OracleManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleManagerContract *OracleManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.contract.Transact(opts, method, params...)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleManagerContract *OracleManagerContractCaller) _hasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleManagerContract *OracleManagerContractSession) _hasAvailability() (bool, error) {
	return _OracleManagerContract.Contract._hasAvailability(&_OracleManagerContract.CallOpts)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleManagerContract *OracleManagerContractCallerSession) _hasAvailability() (bool, error) {
	return _OracleManagerContract.Contract._hasAvailability(&_OracleManagerContract.CallOpts)
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleManagerContract *OracleManagerContractCaller) _isOracle(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "_isOracle", identity)
	return *ret0, err
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleManagerContract *OracleManagerContractSession) _isOracle(identity common.Address) (bool, error) {
	return _OracleManagerContract.Contract._isOracle(&_OracleManagerContract.CallOpts, identity)
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleManagerContract *OracleManagerContractCallerSession) _isOracle(identity common.Address) (bool, error) {
	return _OracleManagerContract.Contract._isOracle(&_OracleManagerContract.CallOpts, identity)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractSession) BaseDeposit() (*big.Int, error) {
	return _OracleManagerContract.Contract.BaseDeposit(&_OracleManagerContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractCallerSession) BaseDeposit() (*big.Int, error) {
	return _OracleManagerContract.Contract.BaseDeposit(&_OracleManagerContract.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractCaller) FreezePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "freezePeriod")
	return *ret0, err
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractSession) FreezePeriod() (*big.Int, error) {
	return _OracleManagerContract.Contract.FreezePeriod(&_OracleManagerContract.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractCallerSession) FreezePeriod() (*big.Int, error) {
	return _OracleManagerContract.Contract.FreezePeriod(&_OracleManagerContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleManagerContract *OracleManagerContractCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleManagerContract *OracleManagerContractSession) GetMinimumDeposit() (*big.Int, error) {
	return _OracleManagerContract.Contract.GetMinimumDeposit(&_OracleManagerContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleManagerContract *OracleManagerContractCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _OracleManagerContract.Contract.GetMinimumDeposit(&_OracleManagerContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_OracleManagerContract *OracleManagerContractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_OracleManagerContract *OracleManagerContractSession) GetOwner() (common.Address, error) {
	return _OracleManagerContract.Contract.GetOwner(&_OracleManagerContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_OracleManagerContract *OracleManagerContractCallerSession) GetOwner() (common.Address, error) {
	return _OracleManagerContract.Contract.GetOwner(&_OracleManagerContract.CallOpts)
}

// GetSubmissionAtIndex is a free data retrieval call binding the contract method 0x3232f108.
//
// Solidity: function getSubmissionAtIndex(index uint256) constant returns(sender address, price uint256)
func (_OracleManagerContract *OracleManagerContractCaller) GetSubmissionAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Sender common.Address
	Price  *big.Int
}, error) {
	ret := new(struct {
		Sender common.Address
		Price  *big.Int
	})
	out := ret
	err := _OracleManagerContract.contract.Call(opts, out, "getSubmissionAtIndex", index)
	return *ret, err
}

// GetSubmissionAtIndex is a free data retrieval call binding the contract method 0x3232f108.
//
// Solidity: function getSubmissionAtIndex(index uint256) constant returns(sender address, price uint256)
func (_OracleManagerContract *OracleManagerContractSession) GetSubmissionAtIndex(index *big.Int) (struct {
	Sender common.Address
	Price  *big.Int
}, error) {
	return _OracleManagerContract.Contract.GetSubmissionAtIndex(&_OracleManagerContract.CallOpts, index)
}

// GetSubmissionAtIndex is a free data retrieval call binding the contract method 0x3232f108.
//
// Solidity: function getSubmissionAtIndex(index uint256) constant returns(sender address, price uint256)
func (_OracleManagerContract *OracleManagerContractCallerSession) GetSubmissionAtIndex(index *big.Int) (struct {
	Sender common.Address
	Price  *big.Int
}, error) {
	return _OracleManagerContract.Contract.GetSubmissionAtIndex(&_OracleManagerContract.CallOpts, index)
}

// GetSubmissionCount is a free data retrieval call binding the contract method 0x9999d2ae.
//
// Solidity: function getSubmissionCount() constant returns(count uint256)
func (_OracleManagerContract *OracleManagerContractCaller) GetSubmissionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "getSubmissionCount")
	return *ret0, err
}

// GetSubmissionCount is a free data retrieval call binding the contract method 0x9999d2ae.
//
// Solidity: function getSubmissionCount() constant returns(count uint256)
func (_OracleManagerContract *OracleManagerContractSession) GetSubmissionCount() (*big.Int, error) {
	return _OracleManagerContract.Contract.GetSubmissionCount(&_OracleManagerContract.CallOpts)
}

// GetSubmissionCount is a free data retrieval call binding the contract method 0x9999d2ae.
//
// Solidity: function getSubmissionCount() constant returns(count uint256)
func (_OracleManagerContract *OracleManagerContractCallerSession) GetSubmissionCount() (*big.Int, error) {
	return _OracleManagerContract.Contract.GetSubmissionCount(&_OracleManagerContract.CallOpts)
}

// MaxOracles is a free data retrieval call binding the contract method 0xc0d2c49d.
//
// Solidity: function maxOracles() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractCaller) MaxOracles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManagerContract.contract.Call(opts, out, "maxOracles")
	return *ret0, err
}

// MaxOracles is a free data retrieval call binding the contract method 0xc0d2c49d.
//
// Solidity: function maxOracles() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractSession) MaxOracles() (*big.Int, error) {
	return _OracleManagerContract.Contract.MaxOracles(&_OracleManagerContract.CallOpts)
}

// MaxOracles is a free data retrieval call binding the contract method 0xc0d2c49d.
//
// Solidity: function maxOracles() constant returns(uint256)
func (_OracleManagerContract *OracleManagerContractCallerSession) MaxOracles() (*big.Int, error) {
	return _OracleManagerContract.Contract.MaxOracles(&_OracleManagerContract.CallOpts)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleManagerContract *OracleManagerContractTransactor) DeregisterOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManagerContract.contract.Transact(opts, "deregisterOracle")
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleManagerContract *OracleManagerContractSession) DeregisterOracle() (*types.Transaction, error) {
	return _OracleManagerContract.Contract.DeregisterOracle(&_OracleManagerContract.TransactOpts)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleManagerContract *OracleManagerContractTransactorSession) DeregisterOracle() (*types.Transaction, error) {
	return _OracleManagerContract.Contract.DeregisterOracle(&_OracleManagerContract.TransactOpts)
}

// RedeemDeposits is a paid mutator transaction binding the contract method 0x4b2c89d5.
//
// Solidity: function redeemDeposits() returns()
func (_OracleManagerContract *OracleManagerContractTransactor) RedeemDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManagerContract.contract.Transact(opts, "redeemDeposits")
}

// RedeemDeposits is a paid mutator transaction binding the contract method 0x4b2c89d5.
//
// Solidity: function redeemDeposits() returns()
func (_OracleManagerContract *OracleManagerContractSession) RedeemDeposits() (*types.Transaction, error) {
	return _OracleManagerContract.Contract.RedeemDeposits(&_OracleManagerContract.TransactOpts)
}

// RedeemDeposits is a paid mutator transaction binding the contract method 0x4b2c89d5.
//
// Solidity: function redeemDeposits() returns()
func (_OracleManagerContract *OracleManagerContractTransactorSession) RedeemDeposits() (*types.Transaction, error) {
	return _OracleManagerContract.Contract.RedeemDeposits(&_OracleManagerContract.TransactOpts)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleManagerContract *OracleManagerContractTransactor) RegisterOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManagerContract.contract.Transact(opts, "registerOracle")
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleManagerContract *OracleManagerContractSession) RegisterOracle() (*types.Transaction, error) {
	return _OracleManagerContract.Contract.RegisterOracle(&_OracleManagerContract.TransactOpts)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleManagerContract *OracleManagerContractTransactorSession) RegisterOracle() (*types.Transaction, error) {
	return _OracleManagerContract.Contract.RegisterOracle(&_OracleManagerContract.TransactOpts)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(price uint256) returns()
func (_OracleManagerContract *OracleManagerContractTransactor) SubmitPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _OracleManagerContract.contract.Transact(opts, "submitPrice", price)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(price uint256) returns()
func (_OracleManagerContract *OracleManagerContractSession) SubmitPrice(price *big.Int) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.SubmitPrice(&_OracleManagerContract.TransactOpts, price)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(price uint256) returns()
func (_OracleManagerContract *OracleManagerContractTransactorSession) SubmitPrice(price *big.Int) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.SubmitPrice(&_OracleManagerContract.TransactOpts, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_OracleManagerContract *OracleManagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _OracleManagerContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_OracleManagerContract *OracleManagerContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.TransferOwnership(&_OracleManagerContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_OracleManagerContract *OracleManagerContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OracleManagerContract.Contract.TransferOwnership(&_OracleManagerContract.TransactOpts, _newOwner)
}
