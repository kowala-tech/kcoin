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
const NetworkContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"lastPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVoterCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTERS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isGenesisVoter\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastBlockReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupplyWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVoterAtIndex\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NetworkContractBin is the compiled bytecode used for deploying new contracts.
const NetworkContractBin = `6060604052670de0b6b3a764000060005560006001556000600255600a600655341561002a57600080fd5b60008060008073d6e579085c82329c89fca7a9f012be59028ed53f935073497dc8a0096cf116e696ba9072516c92383770ed9250606491506064905081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061010b848361013164010000000002610884176401000000009004565b610128838261013164010000000002610884176401000000009004565b50505050610278565b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060016005805480600101828161018e9190610227565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055505050565b81548183558181151161024e5781836000526020600020918201910161024d9190610253565b5b505050565b61027591905b80821115610271576000816000905550600101610259565b5090565b90565b610a23806102876000396000f3006060604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063053f14da146100ca57806311174a29146100f35780633ccfd60b1461011c5780633ceed6921461013157806341b3d1851461015a5780635334ecb3146101835780635798a6d5146101d457806370a8f25b146101fd578063a7771ee314610226578063b80bec5814610277578063c9b53900146102e1578063d0e30db01461030e578063d4f50f9814610318575b600080fd5b34156100d557600080fd5b6100dd61036c565b6040518082815260200191505060405180910390f35b34156100fe57600080fd5b610106610372565b6040518082815260200191505060405180910390f35b341561012757600080fd5b61012f61037f565b005b341561013c57600080fd5b610144610420565b6040518082815260200191505060405180910390f35b341561016557600080fd5b61016d610425565b6040518082815260200191505060405180910390f35b341561018e57600080fd5b6101ba600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061042b565b604051808215151515815260200191505060405180910390f35b34156101df57600080fd5b6101e7610476565b6040518082815260200191505060405180910390f35b341561020857600080fd5b61021061047c565b6040518082815260200191505060405180910390f35b341561023157600080fd5b61025d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610482565b604051808215151515815260200191505060405180910390f35b341561028257600080fd5b610298600480803590602001909190505061054d565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34156102ec57600080fd5b6102f46105d7565b604051808215151515815260200191505060405180910390f35b6103166105e7565b005b341561032357600080fd5b61034f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610696565b604051808381526020018281526020019250505060405180910390f35b60025481565b6000600580549050905090565b61038833610482565b151561039357600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549081150290604051600060405180830381858888f19350505050151561041557600080fd5b61041e3361073c565b565b600381565b60065481565b600080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054119050919050565b60015481565b60005481565b600080600580549050141561049a5760009050610548565b8173ffffffffffffffffffffffffffffffffffffffff166005600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015481548110151561050257fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505b919050565b60008060058381548110151561055f57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050915091565b6000600360058054905010905090565b60006105f233610482565b1515156105fe57600080fd5b6106073361042b565b151561063757600360058054905010151561062157600080fd5b600654341015151561063257600080fd5b610689565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905080341015151561068857600080fd5b5b6106933334610884565b50565b6000806106a283610482565b15156106ad57600080fd5b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549150600560016005805490500381548110151561079b57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806005838154811015156107d957fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010181905550600580548091906001900361087e919061097a565b50505050565b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506001600580548060010182816108e191906109a6565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055505050565b8154818355818115116109a1578183600052602060002091820191016109a091906109d2565b5b505050565b8154818355818115116109cd578183600052602060002091820191016109cc91906109d2565b5b505050565b6109f491905b808211156109f05760008160009055506001016109d8565b5090565b905600a165627a7a72305820d8db80174128b7675571a8304fe7c1535e80f10ab09680e0eb7e9edaa5cb5c450029`

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
