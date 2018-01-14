// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// VoterRegistryABI is the input ABI used to generate the binding from.
const VoterRegistryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"isVoter\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VoterRegistryBin is the compiled bytecode used for deploying new contracts.
const VoterRegistryBin = `0x6060604052600360005560006001556000600255341561001e57600080fd5b6103138061002d6000396000f3006060604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633ccfd60b811461007c578063636bfbab14610091578063a3ec138d146100b6578063a7771ee3146100ef578063c9b5390014610122578063d0e30db014610135575b600080fd5b341561008757600080fd5b61008f61013d565b005b341561009c57600080fd5b6100a46101ac565b60405190815260200160405180910390f35b34156100c157600080fd5b6100d5600160a060020a03600435166101b2565b604051918252151560208201526040908101905180910390f35b34156100fa57600080fd5b61010e600160a060020a03600435166101ce565b604051901515815260200160405180910390f35b341561012d57600080fd5b61010e6101ef565b61008f6101f9565b336000610149826101ce565b151561015457600080fd5b50600160a060020a03811660009081526003602052604090205461017782610237565b600160a060020a03821681156108fc0282604051600060405180830381858888f1935050505015156101a857600080fd5b5050565b60025490565b6003602052600090815260409020805460019091015460ff1682565b600160a060020a031660009081526003602052604090206001015460ff1690565b6000546001541090565b610202336101ce565b1561020c57600080fd5b6000546001541061021c57600080fd5b60025434101561022b57600080fd5b610235333461028d565b565b6001805460001901905560408051908101604090815260008083526020808401829052600160a060020a03331682526003905220815181556020820151600191909101805460ff19169115159190911790555050565b60018054810190556040805190810160409081528282526001602080840191909152600160a060020a0385166000908152600390915220815181556020820151600191909101805460ff19169115159190911790555050505600a165627a7a72305820457a380a4f1ac12f7c13a9d2d275f616f901da8425d9f64231aa50809aa8e30d0029`

// DeployVoterRegistry deploys a new Ethereum contract, binding an instance of VoterRegistry to it.
func DeployVoterRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VoterRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VoterRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VoterRegistry{VoterRegistryCaller: VoterRegistryCaller{contract: contract}, VoterRegistryTransactor: VoterRegistryTransactor{contract: contract}}, nil
}

// VoterRegistry is an auto generated Go binding around an Ethereum contract.
type VoterRegistry struct {
	VoterRegistryCaller     // Read-only binding to the contract
	VoterRegistryTransactor // Write-only binding to the contract
}

// VoterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterRegistrySession struct {
	Contract     *VoterRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterRegistryCallerSession struct {
	Contract *VoterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VoterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterRegistryTransactorSession struct {
	Contract     *VoterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VoterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterRegistryRaw struct {
	Contract *VoterRegistry // Generic contract binding to access the raw methods on
}

// VoterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterRegistryCallerRaw struct {
	Contract *VoterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// VoterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterRegistryTransactorRaw struct {
	Contract *VoterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterRegistry creates a new instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistry(address common.Address, backend bind.ContractBackend) (*VoterRegistry, error) {
	contract, err := bindVoterRegistry(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterRegistry{VoterRegistryCaller: VoterRegistryCaller{contract: contract}, VoterRegistryTransactor: VoterRegistryTransactor{contract: contract}}, nil
}

// NewVoterRegistryCaller creates a new read-only instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryCaller(address common.Address, caller bind.ContractCaller) (*VoterRegistryCaller, error) {
	contract, err := bindVoterRegistry(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryCaller{contract: contract}, nil
}

// NewVoterRegistryTransactor creates a new write-only instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterRegistryTransactor, error) {
	contract, err := bindVoterRegistry(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryTransactor{contract: contract}, nil
}

// bindVoterRegistry binds a generic wrapper to an already deployed contract.
func bindVoterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterRegistry *VoterRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterRegistry.Contract.VoterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterRegistry *VoterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.Contract.VoterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterRegistry *VoterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterRegistry.Contract.VoterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterRegistry *VoterRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterRegistry *VoterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterRegistry *VoterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterRegistry.Contract.contract.Transact(opts, method, params...)
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(bool)
func (_VoterRegistry *VoterRegistryCaller) Availability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VoterRegistry.contract.Call(opts, out, "availability")
	return *ret0, err
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(bool)
func (_VoterRegistry *VoterRegistrySession) Availability() (bool, error) {
	return _VoterRegistry.Contract.Availability(&_VoterRegistry.CallOpts)
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) Availability() (bool, error) {
	return _VoterRegistry.Contract.Availability(&_VoterRegistry.CallOpts)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VoterRegistry *VoterRegistryCaller) IsVoter(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VoterRegistry.contract.Call(opts, out, "isVoter", addr)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VoterRegistry *VoterRegistrySession) IsVoter(addr common.Address) (bool, error) {
	return _VoterRegistry.Contract.IsVoter(&_VoterRegistry.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) IsVoter(addr common.Address) (bool, error) {
	return _VoterRegistry.Contract.IsVoter(&_VoterRegistry.CallOpts, addr)
}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() constant returns(uint256)
func (_VoterRegistry *VoterRegistryCaller) MinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VoterRegistry.contract.Call(opts, out, "minimumDeposit")
	return *ret0, err
}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() constant returns(uint256)
func (_VoterRegistry *VoterRegistrySession) MinimumDeposit() (*big.Int, error) {
	return _VoterRegistry.Contract.MinimumDeposit(&_VoterRegistry.CallOpts)
}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() constant returns(uint256)
func (_VoterRegistry *VoterRegistryCallerSession) MinimumDeposit() (*big.Int, error) {
	return _VoterRegistry.Contract.MinimumDeposit(&_VoterRegistry.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(deposit uint256, isVoter bool)
func (_VoterRegistry *VoterRegistryCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit *big.Int
	IsVoter bool
}, error) {
	ret := new(struct {
		Deposit *big.Int
		IsVoter bool
	})
	out := ret
	err := _VoterRegistry.contract.Call(opts, out, "voters", arg0)
	return *ret, err
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(deposit uint256, isVoter bool)
func (_VoterRegistry *VoterRegistrySession) Voters(arg0 common.Address) (struct {
	Deposit *big.Int
	IsVoter bool
}, error) {
	return _VoterRegistry.Contract.Voters(&_VoterRegistry.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(deposit uint256, isVoter bool)
func (_VoterRegistry *VoterRegistryCallerSession) Voters(arg0 common.Address) (struct {
	Deposit *big.Int
	IsVoter bool
}, error) {
	return _VoterRegistry.Contract.Voters(&_VoterRegistry.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_VoterRegistry *VoterRegistryTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_VoterRegistry *VoterRegistrySession) Deposit() (*types.Transaction, error) {
	return _VoterRegistry.Contract.Deposit(&_VoterRegistry.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_VoterRegistry *VoterRegistryTransactorSession) Deposit() (*types.Transaction, error) {
	return _VoterRegistry.Contract.Deposit(&_VoterRegistry.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VoterRegistry *VoterRegistryTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VoterRegistry *VoterRegistrySession) Withdraw() (*types.Transaction, error) {
	return _VoterRegistry.Contract.Withdraw(&_VoterRegistry.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VoterRegistry *VoterRegistryTransactorSession) Withdraw() (*types.Transaction, error) {
	return _VoterRegistry.Contract.Withdraw(&_VoterRegistry.TransactOpts)
}
