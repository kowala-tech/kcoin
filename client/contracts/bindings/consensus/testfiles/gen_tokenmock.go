// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testfiles

import (
	"math/big"
	"strings"

	kowala "github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
)

// TokenMockABI is the input ABI used to generate the binding from.
const TokenMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"custom_fallback\",\"type\":\"string\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TokenMockBin is the compiled bytecode used for deploying new contracts.
const TokenMockBin = `608060405234801561001057600080fd5b50610669806100206000396000f30060806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461009357806318160ddd14610123578063313ce5671461014e57806370a082311461017f57806395d89b41146101d6578063a9059cbb14610266578063be45fd62146102cb578063f6368f8a14610376575b600080fd5b34801561009f57600080fd5b506100a8610467565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100e85780820151818401526020810190506100cd565b50505050905090810190601f1680156101155780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012f57600080fd5b506101386104a4565b6040518082815260200191505060405180910390f35b34801561015a57600080fd5b506101636104ac565b604051808260ff1660ff16815260200191505060405180910390f35b34801561018b57600080fd5b506101c0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104b5565b6040518082815260200191505060405180910390f35b3480156101e257600080fd5b506101eb6104fe565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561022b578082015181840152602081019050610210565b50505050905090810190601f1680156102585780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027257600080fd5b506102b1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061053b565b604051808215151515815260200191505060405180910390f35b3480156102d757600080fd5b5061035c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610590565b604051808215151515815260200191505060405180910390f35b34801561038257600080fd5b5061044d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506105e6565b604051808215151515815260200191505060405180910390f35b60606040805190810160405280600481526020017f6d6f636b00000000000000000000000000000000000000000000000000000000815250905090565b600080905090565b60006012905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606040805190810160405280600481526020017f6d6f636b00000000000000000000000000000000000000000000000000000000815250905090565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555092915050565b600082600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055509392505050565b600083600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055509493505050505600a165627a7a72305820d0574c35fb6133cfceef79b5ed24fc7f557b712bc9613f69dfb8d5b95d6049fd0029`

// DeployTokenMock deploys a new Kowala contract, binding an instance of TokenMock to it.
func DeployTokenMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenMock, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenMock{TokenMockCaller: TokenMockCaller{contract: contract}, TokenMockTransactor: TokenMockTransactor{contract: contract}, TokenMockFilterer: TokenMockFilterer{contract: contract}}, nil
}

// TokenMock is an auto generated Go binding around a Kowala contract.
type TokenMock struct {
	TokenMockCaller     // Read-only binding to the contract
	TokenMockTransactor // Write-only binding to the contract
	TokenMockFilterer   // Log filterer for contract events
}

// TokenMockCaller is an auto generated read-only Go binding around a Kowala contract.
type TokenMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMockTransactor is an auto generated write-only Go binding around a Kowala contract.
type TokenMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMockFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type TokenMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMockSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type TokenMockSession struct {
	Contract     *TokenMock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenMockCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type TokenMockCallerSession struct {
	Contract *TokenMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenMockTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type TokenMockTransactorSession struct {
	Contract     *TokenMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenMockRaw is an auto generated low-level Go binding around a Kowala contract.
type TokenMockRaw struct {
	Contract *TokenMock // Generic contract binding to access the raw methods on
}

// TokenMockCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type TokenMockCallerRaw struct {
	Contract *TokenMockCaller // Generic read-only contract binding to access the raw methods on
}

// TokenMockTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type TokenMockTransactorRaw struct {
	Contract *TokenMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenMock creates a new instance of TokenMock, bound to a specific deployed contract.
func NewTokenMock(address common.Address, backend bind.ContractBackend) (*TokenMock, error) {
	contract, err := bindTokenMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenMock{TokenMockCaller: TokenMockCaller{contract: contract}, TokenMockTransactor: TokenMockTransactor{contract: contract}, TokenMockFilterer: TokenMockFilterer{contract: contract}}, nil
}

// NewTokenMockCaller creates a new read-only instance of TokenMock, bound to a specific deployed contract.
func NewTokenMockCaller(address common.Address, caller bind.ContractCaller) (*TokenMockCaller, error) {
	contract, err := bindTokenMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenMockCaller{contract: contract}, nil
}

// NewTokenMockTransactor creates a new write-only instance of TokenMock, bound to a specific deployed contract.
func NewTokenMockTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenMockTransactor, error) {
	contract, err := bindTokenMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenMockTransactor{contract: contract}, nil
}

// NewTokenMockFilterer creates a new log filterer instance of TokenMock, bound to a specific deployed contract.
func NewTokenMockFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenMockFilterer, error) {
	contract, err := bindTokenMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenMockFilterer{contract: contract}, nil
}

// bindTokenMock binds a generic wrapper to an already deployed contract.
func bindTokenMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenMock *TokenMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenMock.Contract.TokenMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenMock *TokenMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMock.Contract.TokenMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenMock *TokenMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenMock.Contract.TokenMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenMock *TokenMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenMock *TokenMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenMock *TokenMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenMock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TokenMock *TokenMockCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenMock.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TokenMock *TokenMockSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _TokenMock.Contract.BalanceOf(&_TokenMock.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TokenMock *TokenMockCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _TokenMock.Contract.BalanceOf(&_TokenMock.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_TokenMock *TokenMockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TokenMock.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_TokenMock *TokenMockSession) Decimals() (uint8, error) {
	return _TokenMock.Contract.Decimals(&_TokenMock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_TokenMock *TokenMockCallerSession) Decimals() (uint8, error) {
	return _TokenMock.Contract.Decimals(&_TokenMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_TokenMock *TokenMockCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenMock.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_TokenMock *TokenMockSession) Name() (string, error) {
	return _TokenMock.Contract.Name(&_TokenMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_TokenMock *TokenMockCallerSession) Name() (string, error) {
	return _TokenMock.Contract.Name(&_TokenMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_TokenMock *TokenMockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenMock.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_TokenMock *TokenMockSession) Symbol() (string, error) {
	return _TokenMock.Contract.Symbol(&_TokenMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_TokenMock *TokenMockCallerSession) Symbol() (string, error) {
	return _TokenMock.Contract.Symbol(&_TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_supply uint256)
func (_TokenMock *TokenMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_supply uint256)
func (_TokenMock *TokenMockSession) TotalSupply() (*big.Int, error) {
	return _TokenMock.Contract.TotalSupply(&_TokenMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_supply uint256)
func (_TokenMock *TokenMockCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenMock.Contract.TotalSupply(&_TokenMock.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(to address, value uint256, data bytes, custom_fallback string) returns(ok bool)
func (_TokenMock *TokenMockTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, custom_fallback string) (*types.Transaction, error) {
	return _TokenMock.contract.Transact(opts, "transfer", to, value, data, custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(to address, value uint256, data bytes, custom_fallback string) returns(ok bool)
func (_TokenMock *TokenMockSession) Transfer(to common.Address, value *big.Int, data []byte, custom_fallback string) (*types.Transaction, error) {
	return _TokenMock.Contract.Transfer(&_TokenMock.TransactOpts, to, value, data, custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(to address, value uint256, data bytes, custom_fallback string) returns(ok bool)
func (_TokenMock *TokenMockTransactorSession) Transfer(to common.Address, value *big.Int, data []byte, custom_fallback string) (*types.Transaction, error) {
	return _TokenMock.Contract.Transfer(&_TokenMock.TransactOpts, to, value, data, custom_fallback)
}

// TokenMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenMock contract.
type TokenMockTransferIterator struct {
	Event *TokenMockTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMockTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenMockTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMockTransfer represents a Transfer event raised by the TokenMock contract.
type TokenMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  common.Hash
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256, data indexed bytes)
func (_TokenMock *TokenMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, data [][]byte) (*TokenMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _TokenMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &TokenMockTransferIterator{contract: _TokenMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256, data indexed bytes)
func (_TokenMock *TokenMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenMockTransfer, from []common.Address, to []common.Address, data [][]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _TokenMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMockTransfer)
				if err := _TokenMock.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
