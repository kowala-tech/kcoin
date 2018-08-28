// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

import (
	"strings"

	kowala "github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
)

// UpgradeabilityProxyFactoryABI is the input ABI used to generate the binding from.
const UpgradeabilityProxyFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"createProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"implementation\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createProxyAndCall\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"}]"

// UpgradeabilityProxyFactoryBin is the compiled bytecode used for deploying new contracts.
const UpgradeabilityProxyFactoryBin = `608060405234801561001057600080fd5b50611029806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806325b5672714610051578063c6e8b4f3146100f4575b600080fd5b34801561005d57600080fd5b506100b2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61018e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061029b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000806101dc836103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561027957600080fd5b505af115801561028d573d6000803e3d6000fd5b505050508091505092915050565b6000806102a7846103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561034457600080fd5b505af1158015610358573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16348460405180828051906020019080838360005b838110156103a2578082015181840152602081019050610387565b50505050905090810190601f1680156103cf5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af19250505015156103ef57600080fd5b809150509392505050565b600080826104066104c6565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610458573d6000803e3d6000fd5b5090507efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e734981604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a180915050919050565b604051610b27806104d7833901905600608060405234801561001057600080fd5b50604051602080610b27833981018060405281019080805190602001909291905050508060405180807f6f72672e7a657070656c696e6f732e70726f78792e696d706c656d656e74617481526020017f696f6e000000000000000000000000000000000000000000000000000000000081525060230190506040518091039020600019167f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3600102600019161415156100c557fe5b6100dd81610167640100000000026401000000009004565b5060405180807f6f72672e7a657070656c696e6f732e70726f78792e61646d696e000000000000815250601a0190506040518091039020600019167f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001026000191614151561014957fe5b6101613361024c640100000000026401000000009004565b5061028e565b60006101858261027b6401000000000261084b176401000000009004565b151561021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b600080823b905060008111915050919050565b61088a8061029d6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633659cfe6146100775780634f1ef286146100ba5780635c60da1b146101085780638f2839701461015f578063f851a440146101a2575b6100756101f9565b005b34801561008357600080fd5b506100b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610213565b005b610106600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001919091929391929390505050610268565b005b34801561011457600080fd5b5061011d610308565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016b57600080fd5b506101a0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610360565b005b3480156101ae57600080fd5b506101b761051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610201610576565b61021161020c610651565b610682565b565b61021b6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025c57610257816106d9565b610265565b6102646101f9565b5b50565b6102706106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fa576102ac836106d9565b3073ffffffffffffffffffffffffffffffffffffffff163483836040518083838082843782019150509250505060006040518083038185875af19250505015156102f557600080fd5b610303565b6103026101f9565b5b505050565b60006103126106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103545761034d610651565b905061035d565b61035c6101f9565b5b90565b6103686106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561051257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610466576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f81526020017f787920746f20746865207a65726f20616464726573730000000000000000000081525060400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61048f6106a8565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050d81610748565b61051b565b61051a6101f9565b5b50565b60006105286106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561056a576105636106a8565b9050610573565b6105726101f9565b5b90565b61057e6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667281526020017f6f6d207468652070726f78792061646d696e000000000000000000000000000081525060400191505060405180910390fd5b61064f610777565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146106a3573d6000f35b3d6000fd5b6000807f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001029050805491505090565b6106e281610779565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b565b60006107848261084b565b151561081e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b600080823b9050600081119150509190505600a165627a7a72305820c3e0cc5b45ccfae3aad0a15ba061adbdf7cef6c82f830d01cfd81e1e28dc433b0029a165627a7a72305820f3ea1a680684c4182b7335e498aba130454f0b5066762adf5290c6a211d4b5980029`

// DeployUpgradeabilityProxyFactory deploys a new Kowala contract, binding an instance of UpgradeabilityProxyFactory to it.
func DeployUpgradeabilityProxyFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UpgradeabilityProxyFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeabilityProxyFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradeabilityProxyFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeabilityProxyFactory{UpgradeabilityProxyFactoryCaller: UpgradeabilityProxyFactoryCaller{contract: contract}, UpgradeabilityProxyFactoryTransactor: UpgradeabilityProxyFactoryTransactor{contract: contract}, UpgradeabilityProxyFactoryFilterer: UpgradeabilityProxyFactoryFilterer{contract: contract}}, nil
}

// UpgradeabilityProxyFactory is an auto generated Go binding around a Kowala contract.
type UpgradeabilityProxyFactory struct {
	UpgradeabilityProxyFactoryCaller     // Read-only binding to the contract
	UpgradeabilityProxyFactoryTransactor // Write-only binding to the contract
	UpgradeabilityProxyFactoryFilterer   // Log filterer for contract events
}

// UpgradeabilityProxyFactoryCaller is an auto generated read-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyFactoryTransactor is an auto generated write-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyFactoryFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type UpgradeabilityProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyFactorySession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type UpgradeabilityProxyFactorySession struct {
	Contract     *UpgradeabilityProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UpgradeabilityProxyFactoryCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type UpgradeabilityProxyFactoryCallerSession struct {
	Contract *UpgradeabilityProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// UpgradeabilityProxyFactoryTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type UpgradeabilityProxyFactoryTransactorSession struct {
	Contract     *UpgradeabilityProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// UpgradeabilityProxyFactoryRaw is an auto generated low-level Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryRaw struct {
	Contract *UpgradeabilityProxyFactory // Generic contract binding to access the raw methods on
}

// UpgradeabilityProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryCallerRaw struct {
	Contract *UpgradeabilityProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeabilityProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryTransactorRaw struct {
	Contract *UpgradeabilityProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeabilityProxyFactory creates a new instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactory(address common.Address, backend bind.ContractBackend) (*UpgradeabilityProxyFactory, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactory{UpgradeabilityProxyFactoryCaller: UpgradeabilityProxyFactoryCaller{contract: contract}, UpgradeabilityProxyFactoryTransactor: UpgradeabilityProxyFactoryTransactor{contract: contract}, UpgradeabilityProxyFactoryFilterer: UpgradeabilityProxyFactoryFilterer{contract: contract}}, nil
}

// NewUpgradeabilityProxyFactoryCaller creates a new read-only instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*UpgradeabilityProxyFactoryCaller, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryCaller{contract: contract}, nil
}

// NewUpgradeabilityProxyFactoryTransactor creates a new write-only instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeabilityProxyFactoryTransactor, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryTransactor{contract: contract}, nil
}

// NewUpgradeabilityProxyFactoryFilterer creates a new log filterer instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeabilityProxyFactoryFilterer, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryFilterer{contract: contract}, nil
}

// bindUpgradeabilityProxyFactory binds a generic wrapper to an already deployed contract.
func bindUpgradeabilityProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeabilityProxyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradeabilityProxyFactory.Contract.UpgradeabilityProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.UpgradeabilityProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.UpgradeabilityProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradeabilityProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(admin address, implementation address) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactor) CreateProxy(opts *bind.TransactOpts, admin common.Address, implementation common.Address) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.contract.Transact(opts, "createProxy", admin, implementation)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(admin address, implementation address) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactorySession) CreateProxy(admin common.Address, implementation common.Address) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxy(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(admin address, implementation address) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorSession) CreateProxy(admin common.Address, implementation common.Address) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxy(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation)
}

// CreateProxyAndCall is a paid mutator transaction binding the contract method 0xc6e8b4f3.
//
// Solidity: function createProxyAndCall(admin address, implementation address, data bytes) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactor) CreateProxyAndCall(opts *bind.TransactOpts, admin common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.contract.Transact(opts, "createProxyAndCall", admin, implementation, data)
}

// CreateProxyAndCall is a paid mutator transaction binding the contract method 0xc6e8b4f3.
//
// Solidity: function createProxyAndCall(admin address, implementation address, data bytes) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactorySession) CreateProxyAndCall(admin common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxyAndCall(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation, data)
}

// CreateProxyAndCall is a paid mutator transaction binding the contract method 0xc6e8b4f3.
//
// Solidity: function createProxyAndCall(admin address, implementation address, data bytes) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorSession) CreateProxyAndCall(admin common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxyAndCall(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation, data)
}

// UpgradeabilityProxyFactoryProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the UpgradeabilityProxyFactory contract.
type UpgradeabilityProxyFactoryProxyCreatedIterator struct {
	Event *UpgradeabilityProxyFactoryProxyCreated // Event containing the contract specifics and raw log

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
func (it *UpgradeabilityProxyFactoryProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeabilityProxyFactoryProxyCreated)
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
		it.Event = new(UpgradeabilityProxyFactoryProxyCreated)
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
func (it *UpgradeabilityProxyFactoryProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeabilityProxyFactoryProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeabilityProxyFactoryProxyCreated represents a ProxyCreated event raised by the UpgradeabilityProxyFactory contract.
type UpgradeabilityProxyFactoryProxyCreated struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: e ProxyCreated(proxy address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryFilterer) FilterProxyCreated(opts *bind.FilterOpts) (*UpgradeabilityProxyFactoryProxyCreatedIterator, error) {

	logs, sub, err := _UpgradeabilityProxyFactory.contract.FilterLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryProxyCreatedIterator{contract: _UpgradeabilityProxyFactory.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: e ProxyCreated(proxy address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *UpgradeabilityProxyFactoryProxyCreated) (event.Subscription, error) {

	logs, sub, err := _UpgradeabilityProxyFactory.contract.WatchLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeabilityProxyFactoryProxyCreated)
				if err := _UpgradeabilityProxyFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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
