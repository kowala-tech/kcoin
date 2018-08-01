// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stability

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

// StabilityABI is the input ABI used to generate the binding from.
const StabilityABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"hasSubscription\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unsubscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_systemVarsAddr\",\"type\":\"address\"},{\"name\":\"_minDeposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// StabilityBin is the compiled bytecode used for deploying new contracts.
const StabilityBin = `608060405260008060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b50604051604080610ad58339810180604052810190808051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060018190555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506109e5806100f06000396000f300608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633f4ba83a1461009e5780635c975abb146100b5578063715018a6146100e45780637514e80e146100fb5780638456cb59146101565780638da5cb5b1461016d5780638f449a05146101c4578063f2fde38b146101ce578063fcae448414610211575b600080fd5b3480156100aa57600080fd5b506100b3610228565b005b3480156100c157600080fd5b506100ca6102e6565b604051808215151515815260200191505060405180910390f35b3480156100f057600080fd5b506100f96102f9565b005b34801561010757600080fd5b5061013c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103fb565b604051808215151515815260200191505060405180910390f35b34801561016257600080fd5b5061016b610454565b005b34801561017957600080fd5b50610182610514565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101cc610539565b005b3480156101da57600080fd5b5061020f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105c8565b005b34801561021d57600080fd5b5061022661062f565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561028357600080fd5b600060149054906101000a900460ff16151561029e57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561035457600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156104af57600080fd5b600060149054906101000a900460ff161515156104cb57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff1615151561055657600080fd5b61055f336103fb565b156105bc57600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090503481600201600082825401925050819055506105c5565b6105c46107d2565b5b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561062357600080fd5b61062c816108bf565b50565b600061063a336103fb565b151561064557600080fd5b670de0b6b3a7640000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a035b1fe6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156106d457600080fd5b505af11580156106e8573d6000803e3d6000fd5b505050506040513d60208110156106fe57600080fd5b81019080805190602001909291905050501015151561071c57600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090503373ffffffffffffffffffffffffffffffffffffffff166108fc82600201549081150290604051600060405180830381858888f193505050501580156107a7573d6000803e3d6000fd5b506000816002018190555060008160010160006101000a81548160ff02191690831515021790555050565b600060015434101515156107e557600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600160043390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003816000018190555060018160010160006101000a81548160ff02191690831515021790555034816002018190555050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156108fb57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820010e6a2ebbb4c152051567a58499d51f2e7ab5ca3a624441b6f0dc1ccedc1abb0029`

// DeployStability deploys a new Kowala contract, binding an instance of Stability to it.
func DeployStability(auth *bind.TransactOpts, backend bind.ContractBackend, _systemVarsAddr common.Address, _minDeposit *big.Int) (common.Address, *types.Transaction, *Stability, error) {
	parsed, err := abi.JSON(strings.NewReader(StabilityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StabilityBin), backend, _systemVarsAddr, _minDeposit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stability{StabilityCaller: StabilityCaller{contract: contract}, StabilityTransactor: StabilityTransactor{contract: contract}, StabilityFilterer: StabilityFilterer{contract: contract}}, nil
}

// Stability is an auto generated Go binding around a Kowala contract.
type Stability struct {
	StabilityCaller     // Read-only binding to the contract
	StabilityTransactor // Write-only binding to the contract
	StabilityFilterer   // Log filterer for contract events
}

// StabilityCaller is an auto generated read-only Go binding around a Kowala contract.
type StabilityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilityTransactor is an auto generated write-only Go binding around a Kowala contract.
type StabilityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilityFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type StabilityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilitySession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type StabilitySession struct {
	Contract     *Stability        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StabilityCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type StabilityCallerSession struct {
	Contract *StabilityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StabilityTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type StabilityTransactorSession struct {
	Contract     *StabilityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StabilityRaw is an auto generated low-level Go binding around a Kowala contract.
type StabilityRaw struct {
	Contract *Stability // Generic contract binding to access the raw methods on
}

// StabilityCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type StabilityCallerRaw struct {
	Contract *StabilityCaller // Generic read-only contract binding to access the raw methods on
}

// StabilityTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type StabilityTransactorRaw struct {
	Contract *StabilityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStability creates a new instance of Stability, bound to a specific deployed contract.
func NewStability(address common.Address, backend bind.ContractBackend) (*Stability, error) {
	contract, err := bindStability(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stability{StabilityCaller: StabilityCaller{contract: contract}, StabilityTransactor: StabilityTransactor{contract: contract}, StabilityFilterer: StabilityFilterer{contract: contract}}, nil
}

// NewStabilityCaller creates a new read-only instance of Stability, bound to a specific deployed contract.
func NewStabilityCaller(address common.Address, caller bind.ContractCaller) (*StabilityCaller, error) {
	contract, err := bindStability(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StabilityCaller{contract: contract}, nil
}

// NewStabilityTransactor creates a new write-only instance of Stability, bound to a specific deployed contract.
func NewStabilityTransactor(address common.Address, transactor bind.ContractTransactor) (*StabilityTransactor, error) {
	contract, err := bindStability(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StabilityTransactor{contract: contract}, nil
}

// NewStabilityFilterer creates a new log filterer instance of Stability, bound to a specific deployed contract.
func NewStabilityFilterer(address common.Address, filterer bind.ContractFilterer) (*StabilityFilterer, error) {
	contract, err := bindStability(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StabilityFilterer{contract: contract}, nil
}

// bindStability binds a generic wrapper to an already deployed contract.
func bindStability(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StabilityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stability *StabilityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stability.Contract.StabilityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stability *StabilityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.Contract.StabilityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stability *StabilityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stability.Contract.StabilityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stability *StabilityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stability.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stability *StabilityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stability *StabilityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stability.Contract.contract.Transact(opts, method, params...)
}

// HasSubscription is a free data retrieval call binding the contract method 0x7514e80e.
//
// Solidity: function hasSubscription(identity address) constant returns(isIndeed bool)
func (_Stability *StabilityCaller) HasSubscription(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "hasSubscription", identity)
	return *ret0, err
}

// HasSubscription is a free data retrieval call binding the contract method 0x7514e80e.
//
// Solidity: function hasSubscription(identity address) constant returns(isIndeed bool)
func (_Stability *StabilitySession) HasSubscription(identity common.Address) (bool, error) {
	return _Stability.Contract.HasSubscription(&_Stability.CallOpts, identity)
}

// HasSubscription is a free data retrieval call binding the contract method 0x7514e80e.
//
// Solidity: function hasSubscription(identity address) constant returns(isIndeed bool)
func (_Stability *StabilityCallerSession) HasSubscription(identity common.Address) (bool, error) {
	return _Stability.Contract.HasSubscription(&_Stability.CallOpts, identity)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stability *StabilityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stability *StabilitySession) Owner() (common.Address, error) {
	return _Stability.Contract.Owner(&_Stability.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stability *StabilityCallerSession) Owner() (common.Address, error) {
	return _Stability.Contract.Owner(&_Stability.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Stability *StabilityCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Stability *StabilitySession) Paused() (bool, error) {
	return _Stability.Contract.Paused(&_Stability.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Stability *StabilityCallerSession) Paused() (bool, error) {
	return _Stability.Contract.Paused(&_Stability.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stability *StabilityTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stability *StabilitySession) Pause() (*types.Transaction, error) {
	return _Stability.Contract.Pause(&_Stability.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stability *StabilityTransactorSession) Pause() (*types.Transaction, error) {
	return _Stability.Contract.Pause(&_Stability.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stability *StabilityTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stability *StabilitySession) RenounceOwnership() (*types.Transaction, error) {
	return _Stability.Contract.RenounceOwnership(&_Stability.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stability *StabilityTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stability.Contract.RenounceOwnership(&_Stability.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Stability *StabilityTransactor) Subscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "subscribe")
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Stability *StabilitySession) Subscribe() (*types.Transaction, error) {
	return _Stability.Contract.Subscribe(&_Stability.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Stability *StabilityTransactorSession) Subscribe() (*types.Transaction, error) {
	return _Stability.Contract.Subscribe(&_Stability.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Stability *StabilityTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Stability *StabilitySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Stability.Contract.TransferOwnership(&_Stability.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Stability *StabilityTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Stability.Contract.TransferOwnership(&_Stability.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Stability *StabilityTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Stability *StabilitySession) Unpause() (*types.Transaction, error) {
	return _Stability.Contract.Unpause(&_Stability.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Stability *StabilityTransactorSession) Unpause() (*types.Transaction, error) {
	return _Stability.Contract.Unpause(&_Stability.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Stability *StabilityTransactor) Unsubscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "unsubscribe")
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Stability *StabilitySession) Unsubscribe() (*types.Transaction, error) {
	return _Stability.Contract.Unsubscribe(&_Stability.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Stability *StabilityTransactorSession) Unsubscribe() (*types.Transaction, error) {
	return _Stability.Contract.Unsubscribe(&_Stability.TransactOpts)
}

// StabilityOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Stability contract.
type StabilityOwnershipRenouncedIterator struct {
	Event *StabilityOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *StabilityOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityOwnershipRenounced)
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
		it.Event = new(StabilityOwnershipRenounced)
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
func (it *StabilityOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityOwnershipRenounced represents a OwnershipRenounced event raised by the Stability contract.
type StabilityOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Stability *StabilityFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*StabilityOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Stability.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StabilityOwnershipRenouncedIterator{contract: _Stability.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Stability *StabilityFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *StabilityOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Stability.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityOwnershipRenounced)
				if err := _Stability.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// StabilityOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stability contract.
type StabilityOwnershipTransferredIterator struct {
	Event *StabilityOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StabilityOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityOwnershipTransferred)
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
		it.Event = new(StabilityOwnershipTransferred)
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
func (it *StabilityOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityOwnershipTransferred represents a OwnershipTransferred event raised by the Stability contract.
type StabilityOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Stability *StabilityFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StabilityOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stability.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StabilityOwnershipTransferredIterator{contract: _Stability.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Stability *StabilityFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StabilityOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stability.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityOwnershipTransferred)
				if err := _Stability.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// StabilityPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Stability contract.
type StabilityPauseIterator struct {
	Event *StabilityPause // Event containing the contract specifics and raw log

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
func (it *StabilityPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityPause)
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
		it.Event = new(StabilityPause)
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
func (it *StabilityPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityPause represents a Pause event raised by the Stability contract.
type StabilityPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Stability *StabilityFilterer) FilterPause(opts *bind.FilterOpts) (*StabilityPauseIterator, error) {

	logs, sub, err := _Stability.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &StabilityPauseIterator{contract: _Stability.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Stability *StabilityFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *StabilityPause) (event.Subscription, error) {

	logs, sub, err := _Stability.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityPause)
				if err := _Stability.contract.UnpackLog(event, "Pause", log); err != nil {
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

// StabilityUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Stability contract.
type StabilityUnpauseIterator struct {
	Event *StabilityUnpause // Event containing the contract specifics and raw log

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
func (it *StabilityUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityUnpause)
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
		it.Event = new(StabilityUnpause)
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
func (it *StabilityUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityUnpause represents a Unpause event raised by the Stability contract.
type StabilityUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Stability *StabilityFilterer) FilterUnpause(opts *bind.FilterOpts) (*StabilityUnpauseIterator, error) {

	logs, sub, err := _Stability.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &StabilityUnpauseIterator{contract: _Stability.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Stability *StabilityFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *StabilityUnpause) (event.Subscription, error) {

	logs, sub, err := _Stability.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityUnpause)
				if err := _Stability.contract.UnpackLog(event, "Unpause", log); err != nil {
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
