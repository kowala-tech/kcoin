// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nameservice

import (
	"strings"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
)

// NameServiceABI is the input ABI used to generate the binding from.
const NameServiceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"domain\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"lookup\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// NameServiceBin is the compiled bytecode used for deploying new contracts.
const NameServiceBin = `606060405260008060146101000a81548160ff021916908315150217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506107e58061006d6000396000f300606060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631e59c529146100885780633f4ba83a146101045780635c975abb146101195780638456cb59146101465780638da5cb5b1461015b578063f2fde38b146101b0578063f67187ac146101e9575b600080fd5b341561009357600080fd5b610102600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610286565b005b341561010f57600080fd5b61011761038d565b005b341561012457600080fd5b61012c61044b565b604051808215151515815260200191505060405180910390f35b341561015157600080fd5b61015961045e565b005b341561016657600080fd5b61016e61051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156101bb57600080fd5b6101e7600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610543565b005b34156101f457600080fd5b610244600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610698565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60408051908101604052808273ffffffffffffffffffffffffffffffffffffffff168152602001600115158152506001836040518082805190602001908083835b6020831015156102ec57805182526020820191506020810190506020830392506102c7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff0219169083151502179055509050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103e857600080fd5b600060149054906101000a900460ff16151561040357600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156104b957600080fd5b600060149054906101000a900460ff161515156104d557600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561059e57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156105da57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000816001816040518082805190602001908083835b6020831015156106d357805182526020820191506020810190506020830392506106ae565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160149054906101000a900460ff16151561072257600080fd5b6001836040518082805190602001908083835b60208310151561075a5780518252602082019150602081019050602083039250610735565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150509190505600a165627a7a723058207ff7bc9d9588dc60805b2e14d07db0d81203bb865e63bfd91fac9aa8952c56080029`

// DeployNameService deploys a new Ethereum contract, binding an instance of NameService to it.
func DeployNameService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NameService, error) {
	parsed, err := abi.JSON(strings.NewReader(NameServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NameService{NameServiceCaller: NameServiceCaller{contract: contract}, NameServiceTransactor: NameServiceTransactor{contract: contract}}, nil
}

// NameService is an auto generated Go binding around an Ethereum contract.
type NameService struct {
	NameServiceCaller     // Read-only binding to the contract
	NameServiceTransactor // Write-only binding to the contract
}

// NameServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameServiceSession struct {
	Contract     *NameService      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameServiceCallerSession struct {
	Contract *NameServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NameServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameServiceTransactorSession struct {
	Contract     *NameServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NameServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameServiceRaw struct {
	Contract *NameService // Generic contract binding to access the raw methods on
}

// NameServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameServiceCallerRaw struct {
	Contract *NameServiceCaller // Generic read-only contract binding to access the raw methods on
}

// NameServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameServiceTransactorRaw struct {
	Contract *NameServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameService creates a new instance of NameService, bound to a specific deployed contract.
func NewNameService(address common.Address, backend bind.ContractBackend) (*NameService, error) {
	contract, err := bindNameService(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameService{NameServiceCaller: NameServiceCaller{contract: contract}, NameServiceTransactor: NameServiceTransactor{contract: contract}}, nil
}

// NewNameServiceCaller creates a new read-only instance of NameService, bound to a specific deployed contract.
func NewNameServiceCaller(address common.Address, caller bind.ContractCaller) (*NameServiceCaller, error) {
	contract, err := bindNameService(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NameServiceCaller{contract: contract}, nil
}

// NewNameServiceTransactor creates a new write-only instance of NameService, bound to a specific deployed contract.
func NewNameServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*NameServiceTransactor, error) {
	contract, err := bindNameService(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NameServiceTransactor{contract: contract}, nil
}

// bindNameService binds a generic wrapper to an already deployed contract.
func bindNameService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameService *NameServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameService.Contract.NameServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameService *NameServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameService.Contract.NameServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameService *NameServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameService.Contract.NameServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameService *NameServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameService *NameServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameService *NameServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameService.Contract.contract.Transact(opts, method, params...)
}

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(domain string) constant returns(addr address)
func (_NameService *NameServiceCaller) Lookup(opts *bind.CallOpts, domain string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NameService.contract.Call(opts, out, "lookup", domain)
	return *ret0, err
}

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(domain string) constant returns(addr address)
func (_NameService *NameServiceSession) Lookup(domain string) (common.Address, error) {
	return _NameService.Contract.Lookup(&_NameService.CallOpts, domain)
}

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(domain string) constant returns(addr address)
func (_NameService *NameServiceCallerSession) Lookup(domain string) (common.Address, error) {
	return _NameService.Contract.Lookup(&_NameService.CallOpts, domain)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NameService *NameServiceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NameService.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NameService *NameServiceSession) Owner() (common.Address, error) {
	return _NameService.Contract.Owner(&_NameService.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NameService *NameServiceCallerSession) Owner() (common.Address, error) {
	return _NameService.Contract.Owner(&_NameService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_NameService *NameServiceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NameService.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_NameService *NameServiceSession) Paused() (bool, error) {
	return _NameService.Contract.Paused(&_NameService.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_NameService *NameServiceCallerSession) Paused() (bool, error) {
	return _NameService.Contract.Paused(&_NameService.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NameService *NameServiceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameService.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NameService *NameServiceSession) Pause() (*types.Transaction, error) {
	return _NameService.Contract.Pause(&_NameService.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NameService *NameServiceTransactorSession) Pause() (*types.Transaction, error) {
	return _NameService.Contract.Pause(&_NameService.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(domain string, addr address) returns()
func (_NameService *NameServiceTransactor) Register(opts *bind.TransactOpts, domain string, addr common.Address) (*types.Transaction, error) {
	return _NameService.contract.Transact(opts, "register", domain, addr)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(domain string, addr address) returns()
func (_NameService *NameServiceSession) Register(domain string, addr common.Address) (*types.Transaction, error) {
	return _NameService.Contract.Register(&_NameService.TransactOpts, domain, addr)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(domain string, addr address) returns()
func (_NameService *NameServiceTransactorSession) Register(domain string, addr common.Address) (*types.Transaction, error) {
	return _NameService.Contract.Register(&_NameService.TransactOpts, domain, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NameService *NameServiceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NameService.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NameService *NameServiceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NameService.Contract.TransferOwnership(&_NameService.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NameService *NameServiceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NameService.Contract.TransferOwnership(&_NameService.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NameService *NameServiceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameService.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NameService *NameServiceSession) Unpause() (*types.Transaction, error) {
	return _NameService.Contract.Unpause(&_NameService.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NameService *NameServiceTransactorSession) Unpause() (*types.Transaction, error) {
	return _NameService.Contract.Unpause(&_NameService.TransactOpts)
}
