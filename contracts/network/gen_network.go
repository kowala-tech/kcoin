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

// NetworkContractABI is the input ABI used to generate the binding from.
const NetworkContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDepositLowerBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"setMinDepositLowerBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMinDepositUpperBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDepositUpperBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// NetworkContractBin is the compiled bytecode used for deploying new contracts.
const NetworkContractBin = `6060604052341561000f57600080fd5b60405160408061061d83398101604052808051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600181905550600260015481151561008957fe5b0460038190555060026001540260028190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061052f806100ee6000396000f30060606040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806341b3d1851461009357806355abf098146100bc5780636df01b6e146100e55780638fcc9cfb14610108578063a7f0b3de1461012b578063e188f27614610180578063ebfa7716146101a3578063f2fde38b146101cc575b600080fd5b341561009e57600080fd5b6100a6610205565b6040518082815260200191505060405180910390f35b34156100c757600080fd5b6100cf61020b565b6040518082815260200191505060405180910390f35b34156100f057600080fd5b6101066004808035906020019091905050610211565b005b341561011357600080fd5b6101296004808035906020019091905050610287565b005b341561013657600080fd5b61013e61030b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561018b57600080fd5b6101a16004808035906020019091905050610331565b005b34156101ae57600080fd5b6101b66103a7565b6040518082815260200191505060405180910390f35b34156101d757600080fd5b610203600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506103ad565b005b60015481565b60035481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561026c57600080fd5b600254811115151561027d57600080fd5b8060038190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102e257600080fd5b60035481101580156102f657506002548111155b151561030157600080fd5b8060018190555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561038c57600080fd5b600354811015151561039d57600080fd5b8060028190555050565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561040857600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505600a165627a7a72305820aa7a613a3528b94fcd251a4afe0faff32a5e9c4c1e0c33c71a65e8fd641666060029`

// DeployNetworkContract deploys a new Ethereum contract, binding an instance of NetworkContract to it.
func DeployNetworkContract(auth *bind.TransactOpts, backend bind.ContractBackend, _minDeposit *big.Int, _genesis common.Address) (common.Address, *types.Transaction, *NetworkContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkContractBin), backend, _minDeposit, _genesis)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NetworkContract{NetworkContractCaller: NetworkContractCaller{contract: contract}, NetworkContractTransactor: NetworkContractTransactor{contract: contract}}, nil
}

// NetworkContract is an auto generated Go binding around an Ethereum contract.
type NetworkContract struct {
	NetworkContractCaller     // Read-only binding to the contract
	NetworkContractTransactor // Write-only binding to the contract
}

// NetworkContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkContractSession struct {
	Contract     *NetworkContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkContractCallerSession struct {
	Contract *NetworkContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkContractTransactorSession struct {
	Contract     *NetworkContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkContractRaw struct {
	Contract *NetworkContract // Generic contract binding to access the raw methods on
}

// NetworkContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkContractCallerRaw struct {
	Contract *NetworkContractCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkContractTransactorRaw struct {
	Contract *NetworkContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkContract creates a new instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContract(address common.Address, backend bind.ContractBackend) (*NetworkContract, error) {
	contract, err := bindNetworkContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkContract{NetworkContractCaller: NetworkContractCaller{contract: contract}, NetworkContractTransactor: NetworkContractTransactor{contract: contract}}, nil
}

// NewNetworkContractCaller creates a new read-only instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContractCaller(address common.Address, caller bind.ContractCaller) (*NetworkContractCaller, error) {
	contract, err := bindNetworkContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkContractCaller{contract: contract}, nil
}

// NewNetworkContractTransactor creates a new write-only instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkContractTransactor, error) {
	contract, err := bindNetworkContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NetworkContractTransactor{contract: contract}, nil
}

// bindNetworkContract binds a generic wrapper to an already deployed contract.
func bindNetworkContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContract *NetworkContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkContract.Contract.NetworkContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContract *NetworkContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.Contract.NetworkContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContract *NetworkContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContract.Contract.NetworkContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContract *NetworkContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContract *NetworkContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContract *NetworkContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContract.Contract.contract.Transact(opts, method, params...)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_NetworkContract *NetworkContractCaller) Genesis(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "genesis")
	return *ret0, err
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_NetworkContract *NetworkContractSession) Genesis() (common.Address, error) {
	return _NetworkContract.Contract.Genesis(&_NetworkContract.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_NetworkContract *NetworkContractCallerSession) Genesis() (common.Address, error) {
	return _NetworkContract.Contract.Genesis(&_NetworkContract.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "minDeposit")
	return *ret0, err
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MinDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.MinDeposit(&_NetworkContract.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MinDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.MinDeposit(&_NetworkContract.CallOpts)
}

// MinDepositLowerBound is a free data retrieval call binding the contract method 0x55abf098.
//
// Solidity: function minDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MinDepositLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "minDepositLowerBound")
	return *ret0, err
}

// MinDepositLowerBound is a free data retrieval call binding the contract method 0x55abf098.
//
// Solidity: function minDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MinDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositLowerBound(&_NetworkContract.CallOpts)
}

// MinDepositLowerBound is a free data retrieval call binding the contract method 0x55abf098.
//
// Solidity: function minDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MinDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositLowerBound(&_NetworkContract.CallOpts)
}

// MinDepositUpperBound is a free data retrieval call binding the contract method 0xebfa7716.
//
// Solidity: function minDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MinDepositUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "minDepositUpperBound")
	return *ret0, err
}

// MinDepositUpperBound is a free data retrieval call binding the contract method 0xebfa7716.
//
// Solidity: function minDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MinDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositUpperBound(&_NetworkContract.CallOpts)
}

// MinDepositUpperBound is a free data retrieval call binding the contract method 0xebfa7716.
//
// Solidity: function minDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MinDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositUpperBound(&_NetworkContract.CallOpts)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMinDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMinDeposit", deposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMinDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMinDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetMinDepositLowerBound is a paid mutator transaction binding the contract method 0x6df01b6e.
//
// Solidity: function setMinDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMinDepositLowerBound(opts *bind.TransactOpts, min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMinDepositLowerBound", min)
}

// SetMinDepositLowerBound is a paid mutator transaction binding the contract method 0x6df01b6e.
//
// Solidity: function setMinDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMinDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMinDepositLowerBound is a paid mutator transaction binding the contract method 0x6df01b6e.
//
// Solidity: function setMinDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMinDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMinDepositUpperBound is a paid mutator transaction binding the contract method 0xe188f276.
//
// Solidity: function setMinDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMinDepositUpperBound(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMinDepositUpperBound", max)
}

// SetMinDepositUpperBound is a paid mutator transaction binding the contract method 0xe188f276.
//
// Solidity: function setMinDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMinDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetMinDepositUpperBound is a paid mutator transaction binding the contract method 0xe188f276.
//
// Solidity: function setMinDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMinDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositUpperBound(&_NetworkContract.TransactOpts, max)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContract *NetworkContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContract *NetworkContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.TransferOwnership(&_NetworkContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContract *NetworkContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.TransferOwnership(&_NetworkContract.TransactOpts, addr)
}
