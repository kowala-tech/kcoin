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
const NetworkContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"lastPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVoterCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTERS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isGenesisVoter\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastBlockReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupplyWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVoterAtIndex\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votersSummary\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NetworkContractBin is the compiled bytecode used for deploying new contracts.
const NetworkContractBin = `6060604052670de0b6b3a764000060005560006001556000600255620186a0600655341561002c57600080fd5b60008060008060008073d6e579085c82329c89fca7a9f012be59028ed53f955073497dc8a0096cf116e696ba9072516c92383770ed945073d46d2023a7dde27037de5387b38b17ce1e93e3d2935060649250606491506064905082600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061016f8684610197640100000000026108ee176401000000009004565b61018c8583610197640100000000026108ee176401000000009004565b50505050505061035f565b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506001600580548060010182816101f4919061030e565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555060056040518082805480156102f357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116102a9575b50509150506040518091039020600781600019169055505050565b81548183558181151161033557818360005260206000209182019101610334919061033a565b5b505050565b61035c91905b80821115610358576000816000905550600101610340565b5090565b90565b610b0e8061036e6000396000f3006060604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063053f14da146100d557806311174a29146100fe5780633ccfd60b146101275780633ceed6921461013c57806341b3d185146101655780635334ecb31461018e5780635798a6d5146101df57806370a8f25b14610208578063a7771ee314610231578063b80bec5814610282578063c9b53900146102ec578063d0e30db014610319578063d4f50f9814610323578063f5a4180a14610377575b600080fd5b34156100e057600080fd5b6100e86103a8565b6040518082815260200191505060405180910390f35b341561010957600080fd5b6101116103ae565b6040518082815260200191505060405180910390f35b341561013257600080fd5b61013a6103bb565b005b341561014757600080fd5b61014f61045c565b6040518082815260200191505060405180910390f35b341561017057600080fd5b610178610461565b6040518082815260200191505060405180910390f35b341561019957600080fd5b6101c5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610467565b604051808215151515815260200191505060405180910390f35b34156101ea57600080fd5b6101f26104b2565b6040518082815260200191505060405180910390f35b341561021357600080fd5b61021b6104b8565b6040518082815260200191505060405180910390f35b341561023c57600080fd5b610268600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104be565b604051808215151515815260200191505060405180910390f35b341561028d57600080fd5b6102a36004808035906020019091905050610589565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34156102f757600080fd5b6102ff610613565b604051808215151515815260200191505060405180910390f35b610321610623565b005b341561032e57600080fd5b61035a600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610679565b604051808381526020018281526020019250505060405180910390f35b341561038257600080fd5b61038a61071f565b60405180826000191660001916815260200191505060405180910390f35b60025481565b6000600580549050905090565b6103c4336104be565b15156103cf57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549081150290604051600060405180830381858888f19350505050151561045157600080fd5b61045a33610725565b565b606481565b60065481565b600080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054119050919050565b60015481565b60005481565b60008060058054905014156104d65760009050610584565b8173ffffffffffffffffffffffffffffffffffffffff166005600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015481548110151561053e57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505b919050565b60008060058381548110151561059b57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050915091565b6000606460058054905010905090565b61062c336104be565b15151561063857600080fd5b600654341015151561064957600080fd5b61065233610467565b151561066d57606460058054905010151561066c57600080fd5b5b61067733346108ee565b565b600080610685836104be565b151561069057600080fd5b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b60075481565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549150600560016005805490500381548110151561078457fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806005838154811015156107c257fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555060058054809190600190036108679190610a65565b5060056040518082805480156108d257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610888575b5050915050604051809103902060078160001916905550505050565b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060016005805480600101828161094b9190610a91565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506005604051808280548015610a4a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a00575b50509150506040518091039020600781600019169055505050565b815481835581811511610a8c57818360005260206000209182019101610a8b9190610abd565b5b505050565b815481835581811511610ab857818360005260206000209182019101610ab79190610abd565b5b505050565b610adf91905b80821115610adb576000816000905550600101610ac3565b5090565b905600a165627a7a72305820f28242f7cb337dcf736bb71d71731a08b5c1abb62133c2510f74e0f9c1365d570029`

// DeployNetworkContract deploys a new Ethereum contract, binding an instance of NetworkContract to it.
func DeployNetworkContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NetworkContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkContractBin), backend)
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

// MAX_VOTERS is a free data retrieval call binding the contract method 0x3ceed692.
//
// Solidity: function MAX_VOTERS() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MAX_VOTERS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "MAX_VOTERS")
	return *ret0, err
}

// MAX_VOTERS is a free data retrieval call binding the contract method 0x3ceed692.
//
// Solidity: function MAX_VOTERS() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MAX_VOTERS() (*big.Int, error) {
	return _NetworkContract.Contract.MAX_VOTERS(&_NetworkContract.CallOpts)
}

// MAX_VOTERS is a free data retrieval call binding the contract method 0x3ceed692.
//
// Solidity: function MAX_VOTERS() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MAX_VOTERS() (*big.Int, error) {
	return _NetworkContract.Contract.MAX_VOTERS(&_NetworkContract.CallOpts)
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available bool)
func (_NetworkContract *NetworkContractCaller) Availability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "availability")
	return *ret0, err
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available bool)
func (_NetworkContract *NetworkContractSession) Availability() (bool, error) {
	return _NetworkContract.Contract.Availability(&_NetworkContract.CallOpts)
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available bool)
func (_NetworkContract *NetworkContractCallerSession) Availability() (bool, error) {
	return _NetworkContract.Contract.Availability(&_NetworkContract.CallOpts)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(addr address) constant returns(deposit uint256, index uint256)
func (_NetworkContract *NetworkContractCaller) GetVoter(opts *bind.CallOpts, addr common.Address) (struct {
	Deposit *big.Int
	Index   *big.Int
}, error) {
	ret := new(struct {
		Deposit *big.Int
		Index   *big.Int
	})
	out := ret
	err := _NetworkContract.contract.Call(opts, out, "getVoter", addr)
	return *ret, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(addr address) constant returns(deposit uint256, index uint256)
func (_NetworkContract *NetworkContractSession) GetVoter(addr common.Address) (struct {
	Deposit *big.Int
	Index   *big.Int
}, error) {
	return _NetworkContract.Contract.GetVoter(&_NetworkContract.CallOpts, addr)
}

// GetVoter is a free data retrieval call binding the contract method 0xd4f50f98.
//
// Solidity: function getVoter(addr address) constant returns(deposit uint256, index uint256)
func (_NetworkContract *NetworkContractCallerSession) GetVoter(addr common.Address) (struct {
	Deposit *big.Int
	Index   *big.Int
}, error) {
	return _NetworkContract.Contract.GetVoter(&_NetworkContract.CallOpts, addr)
}

// GetVoterAtIndex is a free data retrieval call binding the contract method 0xb80bec58.
//
// Solidity: function getVoterAtIndex(index uint256) constant returns(addr address, deposit uint256)
func (_NetworkContract *NetworkContractCaller) GetVoterAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Addr    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Addr    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _NetworkContract.contract.Call(opts, out, "getVoterAtIndex", index)
	return *ret, err
}

// GetVoterAtIndex is a free data retrieval call binding the contract method 0xb80bec58.
//
// Solidity: function getVoterAtIndex(index uint256) constant returns(addr address, deposit uint256)
func (_NetworkContract *NetworkContractSession) GetVoterAtIndex(index *big.Int) (struct {
	Addr    common.Address
	Deposit *big.Int
}, error) {
	return _NetworkContract.Contract.GetVoterAtIndex(&_NetworkContract.CallOpts, index)
}

// GetVoterAtIndex is a free data retrieval call binding the contract method 0xb80bec58.
//
// Solidity: function getVoterAtIndex(index uint256) constant returns(addr address, deposit uint256)
func (_NetworkContract *NetworkContractCallerSession) GetVoterAtIndex(index *big.Int) (struct {
	Addr    common.Address
	Deposit *big.Int
}, error) {
	return _NetworkContract.Contract.GetVoterAtIndex(&_NetworkContract.CallOpts, index)
}

// GetVoterCount is a free data retrieval call binding the contract method 0x11174a29.
//
// Solidity: function getVoterCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCaller) GetVoterCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "getVoterCount")
	return *ret0, err
}

// GetVoterCount is a free data retrieval call binding the contract method 0x11174a29.
//
// Solidity: function getVoterCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractSession) GetVoterCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetVoterCount(&_NetworkContract.CallOpts)
}

// GetVoterCount is a free data retrieval call binding the contract method 0x11174a29.
//
// Solidity: function getVoterCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCallerSession) GetVoterCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetVoterCount(&_NetworkContract.CallOpts)
}

// IsGenesisVoter is a free data retrieval call binding the contract method 0x5334ecb3.
//
// Solidity: function isGenesisVoter(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsGenesisVoter(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isGenesisVoter", addr)
	return *ret0, err
}

// IsGenesisVoter is a free data retrieval call binding the contract method 0x5334ecb3.
//
// Solidity: function isGenesisVoter(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsGenesisVoter(addr common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisVoter(&_NetworkContract.CallOpts, addr)
}

// IsGenesisVoter is a free data retrieval call binding the contract method 0x5334ecb3.
//
// Solidity: function isGenesisVoter(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsGenesisVoter(addr common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisVoter(&_NetworkContract.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsVoter(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isVoter", addr)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsVoter(addr common.Address) (bool, error) {
	return _NetworkContract.Contract.IsVoter(&_NetworkContract.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsVoter(addr common.Address) (bool, error) {
	return _NetworkContract.Contract.IsVoter(&_NetworkContract.CallOpts, addr)
}

// LastBlockReward is a free data retrieval call binding the contract method 0x5798a6d5.
//
// Solidity: function lastBlockReward() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) LastBlockReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "lastBlockReward")
	return *ret0, err
}

// LastBlockReward is a free data retrieval call binding the contract method 0x5798a6d5.
//
// Solidity: function lastBlockReward() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) LastBlockReward() (*big.Int, error) {
	return _NetworkContract.Contract.LastBlockReward(&_NetworkContract.CallOpts)
}

// LastBlockReward is a free data retrieval call binding the contract method 0x5798a6d5.
//
// Solidity: function lastBlockReward() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) LastBlockReward() (*big.Int, error) {
	return _NetworkContract.Contract.LastBlockReward(&_NetworkContract.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "lastPrice")
	return *ret0, err
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) LastPrice() (*big.Int, error) {
	return _NetworkContract.Contract.LastPrice(&_NetworkContract.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) LastPrice() (*big.Int, error) {
	return _NetworkContract.Contract.LastPrice(&_NetworkContract.CallOpts)
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

// TotalSupplyWei is a free data retrieval call binding the contract method 0x70a8f25b.
//
// Solidity: function totalSupplyWei() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) TotalSupplyWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "totalSupplyWei")
	return *ret0, err
}

// TotalSupplyWei is a free data retrieval call binding the contract method 0x70a8f25b.
//
// Solidity: function totalSupplyWei() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) TotalSupplyWei() (*big.Int, error) {
	return _NetworkContract.Contract.TotalSupplyWei(&_NetworkContract.CallOpts)
}

// TotalSupplyWei is a free data retrieval call binding the contract method 0x70a8f25b.
//
// Solidity: function totalSupplyWei() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) TotalSupplyWei() (*big.Int, error) {
	return _NetworkContract.Contract.TotalSupplyWei(&_NetworkContract.CallOpts)
}

// VotersSummary is a free data retrieval call binding the contract method 0xf5a4180a.
//
// Solidity: function votersSummary() constant returns(bytes32)
func (_NetworkContract *NetworkContractCaller) VotersSummary(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "votersSummary")
	return *ret0, err
}

// VotersSummary is a free data retrieval call binding the contract method 0xf5a4180a.
//
// Solidity: function votersSummary() constant returns(bytes32)
func (_NetworkContract *NetworkContractSession) VotersSummary() ([32]byte, error) {
	return _NetworkContract.Contract.VotersSummary(&_NetworkContract.CallOpts)
}

// VotersSummary is a free data retrieval call binding the contract method 0xf5a4180a.
//
// Solidity: function votersSummary() constant returns(bytes32)
func (_NetworkContract *NetworkContractCallerSession) VotersSummary() ([32]byte, error) {
	return _NetworkContract.Contract.VotersSummary(&_NetworkContract.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NetworkContract *NetworkContractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NetworkContract *NetworkContractSession) Deposit() (*types.Transaction, error) {
	return _NetworkContract.Contract.Deposit(&_NetworkContract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NetworkContract *NetworkContractTransactorSession) Deposit() (*types.Transaction, error) {
	return _NetworkContract.Contract.Deposit(&_NetworkContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NetworkContract *NetworkContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NetworkContract *NetworkContractSession) Withdraw() (*types.Transaction, error) {
	return _NetworkContract.Contract.Withdraw(&_NetworkContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NetworkContract *NetworkContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _NetworkContract.Contract.Withdraw(&_NetworkContract.TransactOpts)
}
