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

// MusdContractABI is the input ABI used to generate the binding from.
const MusdContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"delegateAddr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maximumSupply\",\"outputs\":[{\"name\":\"maxTokenSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"totalTokenSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"delegatedFrom\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegateAddr\",\"type\":\"address\"}],\"name\":\"delegatedTo\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toAddr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegateAddr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegateAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegateAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"fromAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"toAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// MusdContractBin is the compiled bytecode used for deploying new contracts.
const MusdContractBin = `60606040526000600555600060065534156200001a57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040805190810160405280600481526020017f6d5553440000000000000000000000000000000000000000000000000000000081525060019080519060200190620000a7929190620002d6565b506040805190810160405280600481526020017f6d5553440000000000000000000000000000000000000000000000000000000081525060029080519060200190620000f5929190620002d6565b5063400000006006819055506200013673d6e579085c82329c89fca7a9f012be59028ed53f6064620001a96401000000000262000e9b176401000000009004565b506200016c73497dc8a0096cf116e696ba9072516c92383770ed6064620001a96401000000000262000e9b176401000000009004565b50620001a273259be75d96876f2ada3d202722523e9cd4dd917d6064620001a96401000000000262000e9b176401000000009004565b5062000385565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156200020757600080fd5b60065482600554011115620002205760009050620002d0565b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550816005600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a2600190505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200031957805160ff19168380011785556200034a565b828001600101855582156200034a579182015b82811115620003495782518255916020019190600101906200032c565b5b5090506200035991906200035d565b5090565b6200038291905b808211156200037e57600081600090555060010162000364565b5090565b90565b61114780620003956000396000f3006060604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063026e402b146100d55780630480e58b1461012f57806306fdde031461015857806318160ddd146101e65780632c9a96321461020f578063313ce5671461025c57806365da12641461028b57806370a08231146102d857806395d89b4114610325578063a9059cbb146103b3578063eac449d91461040d578063f0dda65c14610467578063f2fde38b146104c1575b34156100d057600080fd5b600080fd5b34156100e057600080fd5b610115600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506104fa565b604051808215151515815260200191505060405180910390f35b341561013a57600080fd5b610142610721565b6040518082815260200191505060405180910390f35b341561016357600080fd5b61016b61072b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101ab578082015181840152602081019050610190565b50505050905090810190601f1680156101d85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101f157600080fd5b6101f96107c9565b6040518082815260200191505060405180910390f35b341561021a57600080fd5b610246600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506107d3565b6040518082815260200191505060405180910390f35b341561026757600080fd5b61026f610859565b604051808260ff1660ff16815260200191505060405180910390f35b341561029657600080fd5b6102c2600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061086c565b6040518082815260200191505060405180910390f35b34156102e357600080fd5b61030f600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506108f2565b6040518082815260200191505060405180910390f35b341561033057600080fd5b6103386109bd565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561037857808201518184015260208101905061035d565b50505050905090810190601f1680156103a55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103be57600080fd5b6103f3600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610a5b565b604051808215151515815260200191505060405180910390f35b341561041857600080fd5b61044d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610c28565b604051808215151515815260200191505060405180910390f35b341561047257600080fd5b6104a7600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610e9b565b604051808215151515815260200191505060405180910390f35b34156104cc57600080fd5b6104f8600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610fc5565b005b600081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403101561058d576000905061071b565b81600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555081600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555081600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f96eafeca8c3c21ab2fa4a636b93ba20c9e22e3d222d92c6530fedc29a53671ee846040518082815260200191505060405180910390a3600190505b92915050565b6000600654905090565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107c15780601f10610796576101008083540402835291602001916107c1565b820191906000526020600020905b8154815290600101906020018083116107a457829003601f168201915b505050505081565b6000600554905090565b6000600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600360009054906101000a900460ff1681565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401039050919050565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a535780601f10610a2857610100808354040283529160200191610a53565b820191906000526020600020905b815481529060010190602001808311610a3657829003601f168201915b505050505081565b6000813073ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515610af857600080fd5b5af11515610b0557600080fd5b505050604051805190501015610b1e5760009050610c22565b81600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b92915050565b600081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610c7a5760009050610e95565b81600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610d075760009050610e95565b81600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167faf2be5d3056627fcbd77a887e7ea236a5c437c5781c0c75b1f71cf3fa5cadfc4846040518082815260200191505060405180910390a3600190505b92915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ef857600080fd5b60065482600554011115610f0f5760009050610fbf565b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550816005600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a2600190505b92915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561102057600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505600a165627a7a723058206195ce20aa045a14df67c5d214a818e212febec1b9fab531c4fb45656c646d280029`

// DeployMusdContract deploys a new Ethereum contract, binding an instance of MusdContract to it.
func DeployMusdContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MusdContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MusdContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MusdContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MusdContract{MusdContractCaller: MusdContractCaller{contract: contract}, MusdContractTransactor: MusdContractTransactor{contract: contract}}, nil
}

// MusdContract is an auto generated Go binding around an Ethereum contract.
type MusdContract struct {
	MusdContractCaller     // Read-only binding to the contract
	MusdContractTransactor // Write-only binding to the contract
}

// MusdContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MusdContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MusdContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MusdContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MusdContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MusdContractSession struct {
	Contract     *MusdContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MusdContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MusdContractCallerSession struct {
	Contract *MusdContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MusdContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MusdContractTransactorSession struct {
	Contract     *MusdContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MusdContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MusdContractRaw struct {
	Contract *MusdContract // Generic contract binding to access the raw methods on
}

// MusdContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MusdContractCallerRaw struct {
	Contract *MusdContractCaller // Generic read-only contract binding to access the raw methods on
}

// MusdContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MusdContractTransactorRaw struct {
	Contract *MusdContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMusdContract creates a new instance of MusdContract, bound to a specific deployed contract.
func NewMusdContract(address common.Address, backend bind.ContractBackend) (*MusdContract, error) {
	contract, err := bindMusdContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MusdContract{MusdContractCaller: MusdContractCaller{contract: contract}, MusdContractTransactor: MusdContractTransactor{contract: contract}}, nil
}

// NewMusdContractCaller creates a new read-only instance of MusdContract, bound to a specific deployed contract.
func NewMusdContractCaller(address common.Address, caller bind.ContractCaller) (*MusdContractCaller, error) {
	contract, err := bindMusdContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MusdContractCaller{contract: contract}, nil
}

// NewMusdContractTransactor creates a new write-only instance of MusdContract, bound to a specific deployed contract.
func NewMusdContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MusdContractTransactor, error) {
	contract, err := bindMusdContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MusdContractTransactor{contract: contract}, nil
}

// bindMusdContract binds a generic wrapper to an already deployed contract.
func bindMusdContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MusdContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MusdContract *MusdContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MusdContract.Contract.MusdContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MusdContract *MusdContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MusdContract.Contract.MusdContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MusdContract *MusdContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MusdContract.Contract.MusdContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MusdContract *MusdContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MusdContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MusdContract *MusdContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MusdContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MusdContract *MusdContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MusdContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(balance uint256)
func (_MusdContract *MusdContractCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "balanceOf", addr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(balance uint256)
func (_MusdContract *MusdContractSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.BalanceOf(&_MusdContract.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(balance uint256)
func (_MusdContract *MusdContractCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.BalanceOf(&_MusdContract.CallOpts, addr)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MusdContract *MusdContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MusdContract *MusdContractSession) Decimals() (uint8, error) {
	return _MusdContract.Contract.Decimals(&_MusdContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MusdContract *MusdContractCallerSession) Decimals() (uint8, error) {
	return _MusdContract.Contract.Decimals(&_MusdContract.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x2c9a9632.
//
// Solidity: function delegatedFrom(delegatorAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCaller) DelegatedFrom(opts *bind.CallOpts, delegatorAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "delegatedFrom", delegatorAddr)
	return *ret0, err
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x2c9a9632.
//
// Solidity: function delegatedFrom(delegatorAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractSession) DelegatedFrom(delegatorAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedFrom(&_MusdContract.CallOpts, delegatorAddr)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x2c9a9632.
//
// Solidity: function delegatedFrom(delegatorAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCallerSession) DelegatedFrom(delegatorAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedFrom(&_MusdContract.CallOpts, delegatorAddr)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(delegateAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCaller) DelegatedTo(opts *bind.CallOpts, delegateAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "delegatedTo", delegateAddr)
	return *ret0, err
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(delegateAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractSession) DelegatedTo(delegateAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedTo(&_MusdContract.CallOpts, delegateAddr)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(delegateAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCallerSession) DelegatedTo(delegateAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedTo(&_MusdContract.CallOpts, delegateAddr)
}

// MaximumSupply is a free data retrieval call binding the contract method 0x0480e58b.
//
// Solidity: function maximumSupply() constant returns(maxTokenSupply uint256)
func (_MusdContract *MusdContractCaller) MaximumSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "maximumSupply")
	return *ret0, err
}

// MaximumSupply is a free data retrieval call binding the contract method 0x0480e58b.
//
// Solidity: function maximumSupply() constant returns(maxTokenSupply uint256)
func (_MusdContract *MusdContractSession) MaximumSupply() (*big.Int, error) {
	return _MusdContract.Contract.MaximumSupply(&_MusdContract.CallOpts)
}

// MaximumSupply is a free data retrieval call binding the contract method 0x0480e58b.
//
// Solidity: function maximumSupply() constant returns(maxTokenSupply uint256)
func (_MusdContract *MusdContractCallerSession) MaximumSupply() (*big.Int, error) {
	return _MusdContract.Contract.MaximumSupply(&_MusdContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MusdContract *MusdContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MusdContract *MusdContractSession) Name() (string, error) {
	return _MusdContract.Contract.Name(&_MusdContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MusdContract *MusdContractCallerSession) Name() (string, error) {
	return _MusdContract.Contract.Name(&_MusdContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MusdContract *MusdContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MusdContract *MusdContractSession) Symbol() (string, error) {
	return _MusdContract.Contract.Symbol(&_MusdContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MusdContract *MusdContractCallerSession) Symbol() (string, error) {
	return _MusdContract.Contract.Symbol(&_MusdContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalTokenSupply uint256)
func (_MusdContract *MusdContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalTokenSupply uint256)
func (_MusdContract *MusdContractSession) TotalSupply() (*big.Int, error) {
	return _MusdContract.Contract.TotalSupply(&_MusdContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalTokenSupply uint256)
func (_MusdContract *MusdContractCallerSession) TotalSupply() (*big.Int, error) {
	return _MusdContract.Contract.TotalSupply(&_MusdContract.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) Delegate(opts *bind.TransactOpts, delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "delegate", delegateAddr, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) Delegate(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Delegate(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) Delegate(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Delegate(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(addr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) MintTokens(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "mintTokens", addr, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(addr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) MintTokens(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.MintTokens(&_MusdContract.TransactOpts, addr, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(addr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) MintTokens(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.MintTokens(&_MusdContract.TransactOpts, addr, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) Revoke(opts *bind.TransactOpts, delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "revoke", delegateAddr, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) Revoke(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Revoke(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) Revoke(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Revoke(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(toAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) Transfer(opts *bind.TransactOpts, toAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "transfer", toAddr, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(toAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) Transfer(toAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Transfer(&_MusdContract.TransactOpts, toAddr, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(toAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) Transfer(toAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Transfer(&_MusdContract.TransactOpts, toAddr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_MusdContract *MusdContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_MusdContract *MusdContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.TransferOwnership(&_MusdContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_MusdContract *MusdContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.TransferOwnership(&_MusdContract.TransactOpts, addr)
}
