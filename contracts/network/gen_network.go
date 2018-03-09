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
const NetworkContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setBaseDepositUpperBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"setBaseDepositLowerBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsUpperBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDepositUpperBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availability\",\"outputs\":[{\"name\":\"available\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDepositLowerBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorsLowerBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorsUpperBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsLowerBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"},{\"name\":\"_maxValidators\",\"type\":\"uint256\"},{\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// NetworkContractBin is the compiled bytecode used for deploying new contracts.
const NetworkContractBin = `6060604052604051608080620014fc83398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508334101515156200008c57600080fd5b600182101515156200009d57600080fd5b60008110151515620000ae57600080fd5b836001819055508160048190555082600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600b8190555060026001548115156200011357fe5b0460038190555060026001540260028190555060026004548115156200013557fe5b04600681905550600260045402600581905550620001688334620001726401000000000262000d3a176401000000009004565b505050506200049b565b620001928282620001966401000000000262000d48176401000000009004565b5050565b600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018054806001018281620001ec9190620003dd565b9160005260206000209060020201600060408051908101604052808581526020016000815250909190915060008201518160000155602082015181600101555050506001600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff021916908315150217905550600160078054806001018281620002a1919062000412565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555062000354620003586401000000000262000ef0176401000000009004565b5050565b6007604051808280548015620003c457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162000379575b5050915050604051809103902060098160001916905550565b8154818355818115116200040d576002028160020283600052602060002091820191016200040c919062000441565b5b505050565b8154818355818115116200043c578183600052602060002091820191016200043b919062000473565b5b505050565b6200047091905b808211156200046c5760008082016000905560018201600090555060020162000448565b5090565b90565b6200049891905b80821115620004945760008160009055506001016200047a565b5090565b90565b61105180620004ab6000396000f300606060405260043610610149576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461014e57806308ac5256146101775780631cbd5487146101a05780633ccfd60b146101c35780634b1d422f146101d857806369474625146101fb5780636cf6d675146102245780637071688a1461024d57806376792094146102765780639bb2ea5a14610299578063a6554f1a146102bc578063a7f0b3de146102e5578063b774cb1e1461033a578063c22a933c1461036b578063c9b539001461038e578063cbfb94ce146103b7578063cefddda9146103e0578063d0e30db014610431578063d66d9e191461043b578063e4f4410b14610450578063e7277b1214610479578063e7a60a9c146104a2578063e99cc6961461050c578063f2fde38b1461052f578063facd743b14610568575b600080fd5b341561015957600080fd5b6101616105b9565b6040518082815260200191505060405180910390f35b341561018257600080fd5b61018a610690565b6040518082815260200191505060405180910390f35b34156101ab57600080fd5b6101c16004808035906020019091905050610696565b005b34156101ce57600080fd5b6101d661070c565b005b34156101e357600080fd5b6101f96004808035906020019091905050610722565b005b341561020657600080fd5b61020e610798565b6040518082815260200191505060405180910390f35b341561022f57600080fd5b61023761079e565b6040518082815260200191505060405180910390f35b341561025857600080fd5b6102606107a4565b6040518082815260200191505060405180910390f35b341561028157600080fd5b61029760048080359060200190919050506107b1565b005b34156102a457600080fd5b6102ba6004808035906020019091905050610827565b005b34156102c757600080fd5b6102cf6108e4565b6040518082815260200191505060405180910390f35b34156102f057600080fd5b6102f86108ea565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561034557600080fd5b61034d610910565b60405180826000191660001916815260200191505060405180910390f35b341561037657600080fd5b61038c6004808035906020019091905050610916565b005b341561039957600080fd5b6103a161099c565b6040518082815260200191505060405180910390f35b34156103c257600080fd5b6103ca6109ad565b6040518082815260200191505060405180910390f35b34156103eb57600080fd5b610417600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506109b3565b604051808215151515815260200191505060405180910390f35b610439610a0d565b005b341561044657600080fd5b61044e610a3b565b005b341561045b57600080fd5b610463610a51565b6040518082815260200191505060405180910390f35b341561048457600080fd5b61048c610a57565b6040518082815260200191505060405180910390f35b34156104ad57600080fd5b6104c36004808035906020019091905050610a5d565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b341561051757600080fd5b61052d6004808035906020019091905050610b15565b005b341561053a57600080fd5b610566600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b8b565b005b341561057357600080fd5b61059f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ce1565b604051808215151515815260200191505060405180910390f35b60008060006105c661099c565b11156105d657600154915061068c565b6008600060076001600780549050038154811015156105f157fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561067657fe5b9060005260206000209060020201600001540191505b5090565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106f157600080fd5b600354811015151561070257600080fd5b8060028190555050565b61071533610ce1565b151561072057600080fd5b565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561077d57600080fd5b600254811115151561078e57600080fd5b8060038190555050565b60015481565b600b5481565b6000600780549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561080c57600080fd5b600654811015151561081d57600080fd5b8060058190555050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561088557600080fd5b82600654811015801561089a57506005548111155b15156108a557600080fd5b60006108af61099c565b14156108d75783600454039250600091505b828210156108d65781806001019250506108c1565b5b8360048190555050505050565b60025481565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561097157600080fd5b80600354811015801561098657506002548111155b151561099157600080fd5b816001819055505050565b600060078054905060045403905090565b60035481565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b610a156105b9565b3410151515610a2357600080fd5b6000610a2d61099c565b5050610a393334610d3a565b565b610a4433610ce1565b1515610a4f57600080fd5b565b60065481565b60055481565b6000806000600784815481101515610a7157fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806002016001826002018054905003815481101515610afb57fe5b906000526020600020906002020160000154915050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b7057600080fd5b6005548111151515610b8157600080fd5b8060068190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610be657600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b610d448282610d48565b5050565b600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018054806001018281610d9c9190610f73565b9160005260206000209060020201600060408051908101604052808581526020016000815250909190915060008201518160000155602082015181600101555050506001600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff021916908315150217905550600160078054806001018281610e4f9190610fa5565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550610eec610ef0565b5050565b6007604051808280548015610f5a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610f10575b5050915050604051809103902060098160001916905550565b815481835581811511610fa057600202816002028360005260206000209182019101610f9f9190610fd1565b5b505050565b815481835581811511610fcc57818360005260206000209182019101610fcb9190611000565b5b505050565b610ffd91905b80821115610ff957600080820160009055600182016000905550600201610fd7565b5090565b90565b61102291905b8082111561101e576000816000905550600101611006565b5090565b905600a165627a7a723058209b345e70a331bbbfcc50afcecab8d5e3f7f3f68e16d4a47faddc6f7637e18b220029`

// DeployNetworkContract deploys a new Ethereum contract, binding an instance of NetworkContract to it.
func DeployNetworkContract(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _genesis common.Address, _maxValidators *big.Int, _unbondingPeriod *big.Int) (common.Address, *types.Transaction, *NetworkContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkContractBin), backend, _baseDeposit, _genesis, _maxValidators, _unbondingPeriod)
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

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available uint256)
func (_NetworkContract *NetworkContractCaller) Availability(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "availability")
	return *ret0, err
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available uint256)
func (_NetworkContract *NetworkContractSession) Availability() (*big.Int, error) {
	return _NetworkContract.Contract.Availability(&_NetworkContract.CallOpts)
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available uint256)
func (_NetworkContract *NetworkContractCallerSession) Availability() (*big.Int, error) {
	return _NetworkContract.Contract.Availability(&_NetworkContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) BaseDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDeposit(&_NetworkContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) BaseDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDeposit(&_NetworkContract.CallOpts)
}

// BaseDepositLowerBound is a free data retrieval call binding the contract method 0xcbfb94ce.
//
// Solidity: function baseDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) BaseDepositLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "baseDepositLowerBound")
	return *ret0, err
}

// BaseDepositLowerBound is a free data retrieval call binding the contract method 0xcbfb94ce.
//
// Solidity: function baseDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) BaseDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositLowerBound(&_NetworkContract.CallOpts)
}

// BaseDepositLowerBound is a free data retrieval call binding the contract method 0xcbfb94ce.
//
// Solidity: function baseDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) BaseDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositLowerBound(&_NetworkContract.CallOpts)
}

// BaseDepositUpperBound is a free data retrieval call binding the contract method 0xa6554f1a.
//
// Solidity: function baseDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) BaseDepositUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "baseDepositUpperBound")
	return *ret0, err
}

// BaseDepositUpperBound is a free data retrieval call binding the contract method 0xa6554f1a.
//
// Solidity: function baseDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) BaseDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositUpperBound(&_NetworkContract.CallOpts)
}

// BaseDepositUpperBound is a free data retrieval call binding the contract method 0xa6554f1a.
//
// Solidity: function baseDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) BaseDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositUpperBound(&_NetworkContract.CallOpts)
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

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_NetworkContract *NetworkContractCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_NetworkContract *NetworkContractSession) GetMinimumDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.GetMinimumDeposit(&_NetworkContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_NetworkContract *NetworkContractCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.GetMinimumDeposit(&_NetworkContract.CallOpts)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_NetworkContract *NetworkContractCaller) GetValidatorAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _NetworkContract.contract.Call(opts, out, "getValidatorAtIndex", index)
	return *ret, err
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_NetworkContract *NetworkContractSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _NetworkContract.Contract.GetValidatorAtIndex(&_NetworkContract.CallOpts, index)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_NetworkContract *NetworkContractCallerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _NetworkContract.Contract.GetValidatorAtIndex(&_NetworkContract.CallOpts, index)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "getValidatorCount")
	return *ret0, err
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractSession) GetValidatorCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetValidatorCount(&_NetworkContract.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCallerSession) GetValidatorCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetValidatorCount(&_NetworkContract.CallOpts)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsGenesisValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isGenesisValidator", code)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisValidator(&_NetworkContract.CallOpts, code)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisValidator(&_NetworkContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isValidator", code)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsValidator(&_NetworkContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsValidator(&_NetworkContract.CallOpts, code)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MaxValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "maxValidators")
	return *ret0, err
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MaxValidators() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidators(&_NetworkContract.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MaxValidators() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidators(&_NetworkContract.CallOpts)
}

// MaxValidatorsLowerBound is a free data retrieval call binding the contract method 0xe4f4410b.
//
// Solidity: function maxValidatorsLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MaxValidatorsLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "maxValidatorsLowerBound")
	return *ret0, err
}

// MaxValidatorsLowerBound is a free data retrieval call binding the contract method 0xe4f4410b.
//
// Solidity: function maxValidatorsLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MaxValidatorsLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsLowerBound(&_NetworkContract.CallOpts)
}

// MaxValidatorsLowerBound is a free data retrieval call binding the contract method 0xe4f4410b.
//
// Solidity: function maxValidatorsLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MaxValidatorsLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsLowerBound(&_NetworkContract.CallOpts)
}

// MaxValidatorsUpperBound is a free data retrieval call binding the contract method 0xe7277b12.
//
// Solidity: function maxValidatorsUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MaxValidatorsUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "maxValidatorsUpperBound")
	return *ret0, err
}

// MaxValidatorsUpperBound is a free data retrieval call binding the contract method 0xe7277b12.
//
// Solidity: function maxValidatorsUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MaxValidatorsUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsUpperBound(&_NetworkContract.CallOpts)
}

// MaxValidatorsUpperBound is a free data retrieval call binding the contract method 0xe7277b12.
//
// Solidity: function maxValidatorsUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MaxValidatorsUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsUpperBound(&_NetworkContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) UnbondingPeriod() (*big.Int, error) {
	return _NetworkContract.Contract.UnbondingPeriod(&_NetworkContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _NetworkContract.Contract.UnbondingPeriod(&_NetworkContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_NetworkContract *NetworkContractCaller) ValidatorsChecksum(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "validatorsChecksum")
	return *ret0, err
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_NetworkContract *NetworkContractSession) ValidatorsChecksum() ([32]byte, error) {
	return _NetworkContract.Contract.ValidatorsChecksum(&_NetworkContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_NetworkContract *NetworkContractCallerSession) ValidatorsChecksum() ([32]byte, error) {
	return _NetworkContract.Contract.ValidatorsChecksum(&_NetworkContract.CallOpts)
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

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_NetworkContract *NetworkContractTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_NetworkContract *NetworkContractSession) Leave() (*types.Transaction, error) {
	return _NetworkContract.Contract.Leave(&_NetworkContract.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_NetworkContract *NetworkContractTransactorSession) Leave() (*types.Transaction, error) {
	return _NetworkContract.Contract.Leave(&_NetworkContract.TransactOpts)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetBaseDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setBaseDeposit", deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetBaseDepositLowerBound is a paid mutator transaction binding the contract method 0x4b1d422f.
//
// Solidity: function setBaseDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetBaseDepositLowerBound(opts *bind.TransactOpts, min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setBaseDepositLowerBound", min)
}

// SetBaseDepositLowerBound is a paid mutator transaction binding the contract method 0x4b1d422f.
//
// Solidity: function setBaseDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractSession) SetBaseDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetBaseDepositLowerBound is a paid mutator transaction binding the contract method 0x4b1d422f.
//
// Solidity: function setBaseDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetBaseDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetBaseDepositUpperBound is a paid mutator transaction binding the contract method 0x1cbd5487.
//
// Solidity: function setBaseDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetBaseDepositUpperBound(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setBaseDepositUpperBound", max)
}

// SetBaseDepositUpperBound is a paid mutator transaction binding the contract method 0x1cbd5487.
//
// Solidity: function setBaseDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetBaseDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetBaseDepositUpperBound is a paid mutator transaction binding the contract method 0x1cbd5487.
//
// Solidity: function setBaseDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetBaseDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMaxValidators(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMaxValidators", max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidators(&_NetworkContract.TransactOpts, max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidators(&_NetworkContract.TransactOpts, max)
}

// SetMaxValidatorsLowerBound is a paid mutator transaction binding the contract method 0xe99cc696.
//
// Solidity: function setMaxValidatorsLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMaxValidatorsLowerBound(opts *bind.TransactOpts, min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMaxValidatorsLowerBound", min)
}

// SetMaxValidatorsLowerBound is a paid mutator transaction binding the contract method 0xe99cc696.
//
// Solidity: function setMaxValidatorsLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMaxValidatorsLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMaxValidatorsLowerBound is a paid mutator transaction binding the contract method 0xe99cc696.
//
// Solidity: function setMaxValidatorsLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMaxValidatorsLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMaxValidatorsUpperBound is a paid mutator transaction binding the contract method 0x76792094.
//
// Solidity: function setMaxValidatorsUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMaxValidatorsUpperBound(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMaxValidatorsUpperBound", max)
}

// SetMaxValidatorsUpperBound is a paid mutator transaction binding the contract method 0x76792094.
//
// Solidity: function setMaxValidatorsUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMaxValidatorsUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetMaxValidatorsUpperBound is a paid mutator transaction binding the contract method 0x76792094.
//
// Solidity: function setMaxValidatorsUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMaxValidatorsUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsUpperBound(&_NetworkContract.TransactOpts, max)
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
