// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

import (
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// NameHashABI is the input ABI used to generate the binding from.
const NameHashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"namehash\",\"outputs\":[{\"name\":\"_namehash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NameHashBin is the compiled bytecode used for deploying new contracts.
const NameHashBin = `6105df610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f3007300000000000000000000000000000000000000003014608060405260043610610058576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063098799621461005d575b600080fd5b6100b7600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506100d5565b60405180826000191660001916815260200191505060405180910390f35b60006100df610599565b60006100e9610599565b606060006100f6876102c3565b945060006001029350610108856102f1565b15156102b65761014c6040805190810160405280600181526020017f2e000000000000000000000000000000000000000000000000000000000000008152506102c3565b92506001610163848761030190919063ffffffff16565b0160405190808252806020026020018201604052801561019757816020015b60608152602001906001900390816101825790505b509150600090505b81518110156101eb576101c36101be848761037890919063ffffffff16565b610392565b82828151811015156101d157fe5b90602001906020020181905250808060010191505061019f565b600090505b81518110156102b55783826001838551030381518110151561020e57fe5b906020019060200201516040518082805190602001908083835b60208310151561024d5780518252602082019150602081019050602083039250610228565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206040518083600019166000191681526020018260001916600019168152602001925050506040518091039020935080806001019150506101f0565b5b8395505050505050919050565b6102cb610599565b600060208301905060408051908101604052808451815260200182815250915050919050565b6000808260000151149050919050565b600080826000015161032585600001518660200151866000015187602001516103f4565b0190505b83600001518460200151018111151561037157818060010192505082600001516103698560200151830386600001510383866000015187602001516103f4565b019050610329565b5092915050565b610380610599565b61038b8383836104b0565b5092915050565b606080600083600001516040519080825280601f01601f1916602001820160405280156103ce5781602001602082028038833980820191505090505b5091506020820190506103ea818560200151866000015161054e565b8192505050919050565b6000806000806000888711151561049e576020871115156104555760018760200360080260020a031980875116888b038a018a96505b81838851161461044a5760018701965080600188031061042a578b8b0196505b5050508394506104a4565b8686209150879350600092505b8689038311151561049d57868420905080600019168260001916141561048a578394506104a4565b6001840193508280600101935050610462565b5b88880194505b50505050949350505050565b6104b8610599565b60006104d685600001518660200151866000015187602001516103f4565b90508460200151836020018181525050846020015181038360000181815250508460000151856020015101811415610518576000856000018181525050610543565b8360000151836000015101856000018181510391508181525050836000015181018560200181815250505b829150509392505050565b60005b6020821015156105765782518452602084019350602083019250602082039150610551565b6001826020036101000a0390508019835116818551168181178652505050505050565b6040805190810160405280600081526020016000815250905600a165627a7a7230582063ad7cf2d7de4f92c6ed298edfc7f7ac8d26ded437713776207c49c31963fd3a0029`

// DeployNameHash deploys a new Kowala contract, binding an instance of NameHash to it.
func DeployNameHash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NameHash, error) {
	parsed, err := abi.JSON(strings.NewReader(NameHashABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameHashBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NameHash{NameHashCaller: NameHashCaller{contract: contract}, NameHashTransactor: NameHashTransactor{contract: contract}, NameHashFilterer: NameHashFilterer{contract: contract}}, nil
}

// NameHash is an auto generated Go binding around a Kowala contract.
type NameHash struct {
	NameHashCaller     // Read-only binding to the contract
	NameHashTransactor // Write-only binding to the contract
	NameHashFilterer   // Log filterer for contract events
}

// NameHashCaller is an auto generated read-only Go binding around a Kowala contract.
type NameHashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameHashTransactor is an auto generated write-only Go binding around a Kowala contract.
type NameHashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameHashFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type NameHashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameHashSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type NameHashSession struct {
	Contract     *NameHash         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameHashCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type NameHashCallerSession struct {
	Contract *NameHashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NameHashTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type NameHashTransactorSession struct {
	Contract     *NameHashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NameHashRaw is an auto generated low-level Go binding around a Kowala contract.
type NameHashRaw struct {
	Contract *NameHash // Generic contract binding to access the raw methods on
}

// NameHashCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type NameHashCallerRaw struct {
	Contract *NameHashCaller // Generic read-only contract binding to access the raw methods on
}

// NameHashTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type NameHashTransactorRaw struct {
	Contract *NameHashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameHash creates a new instance of NameHash, bound to a specific deployed contract.
func NewNameHash(address common.Address, backend bind.ContractBackend) (*NameHash, error) {
	contract, err := bindNameHash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameHash{NameHashCaller: NameHashCaller{contract: contract}, NameHashTransactor: NameHashTransactor{contract: contract}, NameHashFilterer: NameHashFilterer{contract: contract}}, nil
}

// NewNameHashCaller creates a new read-only instance of NameHash, bound to a specific deployed contract.
func NewNameHashCaller(address common.Address, caller bind.ContractCaller) (*NameHashCaller, error) {
	contract, err := bindNameHash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameHashCaller{contract: contract}, nil
}

// NewNameHashTransactor creates a new write-only instance of NameHash, bound to a specific deployed contract.
func NewNameHashTransactor(address common.Address, transactor bind.ContractTransactor) (*NameHashTransactor, error) {
	contract, err := bindNameHash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameHashTransactor{contract: contract}, nil
}

// NewNameHashFilterer creates a new log filterer instance of NameHash, bound to a specific deployed contract.
func NewNameHashFilterer(address common.Address, filterer bind.ContractFilterer) (*NameHashFilterer, error) {
	contract, err := bindNameHash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameHashFilterer{contract: contract}, nil
}

// bindNameHash binds a generic wrapper to an already deployed contract.
func bindNameHash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameHashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameHash *NameHashRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameHash.Contract.NameHashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameHash *NameHashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameHash.Contract.NameHashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameHash *NameHashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameHash.Contract.NameHashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameHash *NameHashCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameHash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameHash *NameHashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameHash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameHash *NameHashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameHash.Contract.contract.Transact(opts, method, params...)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(_node string) constant returns(_namehash bytes32)
func (_NameHash *NameHashCaller) Namehash(opts *bind.CallOpts, _node string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NameHash.contract.Call(opts, out, "namehash", _node)
	return *ret0, err
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(_node string) constant returns(_namehash bytes32)
func (_NameHash *NameHashSession) Namehash(_node string) ([32]byte, error) {
	return _NameHash.Contract.Namehash(&_NameHash.CallOpts, _node)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(_node string) constant returns(_namehash bytes32)
func (_NameHash *NameHashCallerSession) Namehash(_node string) ([32]byte, error) {
	return _NameHash.Contract.Namehash(&_NameHash.CallOpts, _node)
}
