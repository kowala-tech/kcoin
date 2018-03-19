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

// ElectionContractABI is the input ABI used to generate the binding from.
const ElectionContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemFunds\",\"outputs\":[{\"name\":\"refund\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"releasedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxValidators\",\"type\":\"uint256\"},{\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// ElectionContractBin is the compiled bytecode used for deploying new contracts.
const ElectionContractBin = `60606040526040516080806200182883398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508334101515156200008c57600080fd5b600183101515156200009d57600080fd5b83600181905550826002819055508160038190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200011381346200011d6401000000000262000c8e176401000000009004565b5050505062000586565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002093506001600780548060010182816200017c9190620004c8565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff021916908315150217905550836002018054806001018281620002089190620004f7565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b60008311156200041d57600660006007600186038154811015156200027357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915081600201600160078054905003815481101515620002f557fe5b9060005260206000209060020201905080600001548511151562000319576200041d565b6007600184038154811015156200032c57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007848154811015156200036857fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600760018503815481101515620003c557fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082806001900393505062000252565b6200043b620004436401000000000262001105176401000000009004565b505050505050565b6007604051808280548015620004af57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162000464575b5050915050604051809103902060058160001916905550565b815481835581811511620004f257818360005260206000209182019101620004f191906200052c565b5b505050565b815481835581811511620005275760020281600202836000526020600020918201910162000526919062000554565b5b505050565b6200055191905b808211156200054d57600081600090555060010162000533565b5090565b90565b6200058391905b808211156200057f576000808201600090556001820160009055506002016200055b565b5090565b90565b61129280620005966000396000f300606060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461011757806308ac525614610140578063216355ec146101695780633ed0a3731461019257806369474625146101d05780636cf6d675146101f95780637071688a14610222578063893d20e81461024b5780639363a141146102a057806397584b3e146102c95780639bb2ea5a146102f6578063a7f0b3de14610319578063b688a3631461036e578063b774cb1e14610378578063c22a933c146103a9578063cefddda9146103cc578063d66d9e191461041d578063e7a60a9c14610432578063f2fde38b1461049c578063facd743b146104ed575b600080fd5b341561012257600080fd5b61012a61053e565b6040518082815260200191505060405180910390f35b341561014b57600080fd5b610153610612565b6040518082815260200191505060405180910390f35b341561017457600080fd5b61017c610618565b6040518082815260200191505060405180910390f35b341561019d57600080fd5b6101b36004808035906020019091905050610752565b604051808381526020018281526020019250505060405180910390f35b34156101db57600080fd5b6101e36107ca565b6040518082815260200191505060405180910390f35b341561020457600080fd5b61020c6107d0565b6040518082815260200191505060405180910390f35b341561022d57600080fd5b6102356107d6565b6040518082815260200191505060405180910390f35b341561025657600080fd5b61025e6107e3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102ab57600080fd5b6102b361080c565b6040518082815260200191505060405180910390f35b34156102d457600080fd5b6102dc610859565b604051808215151515815260200191505060405180910390f35b341561030157600080fd5b610317600480803590602001909190505061086c565b005b341561032457600080fd5b61032c610910565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610376610936565b005b341561038357600080fd5b61038b61096f565b60405180826000191660001916815260200191505060405180910390f35b34156103b457600080fd5b6103ca6004808035906020019091905050610975565b005b34156103d757600080fd5b610403600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506109da565b604051808215151515815260200191505060405180910390f35b341561042857600080fd5b610430610a34565b005b341561043d57600080fd5b6104536004808035906020019091905050610a53565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34156104a757600080fd5b6104d3600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b0b565b604051808215151515815260200191505060405180910390f35b34156104f857600080fd5b610524600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610be8565b604051808215151515815260200191505060405180910390f35b600080610549610859565b1561055857600154915061060e565b60066000600760016007805490500381548110151561057357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160018360020180549050038154811015156105f857fe5b9060005260206000209060020201600001540191505b5090565b60025481565b6000806000809150600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090505b80600201805490508210801561069b57506000816002018381548110151561068657fe5b90600052602060002090600202016001015414155b156107035780600201828154811015156106b157fe5b9060005260206000209060020201600101544210156106cf57610703565b80600201828154811015156106e057fe5b906000526020600020906002020160000154830192508180600101925050610662565b600083111561074d573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561074c57600080fd5b5b505090565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156107a657fe5b90600052602060002090600202019050806000015481600101549250925050915091565b60015481565b60035481565b6000600780549050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506002540311905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108ca57600080fd5b6007805490508310156109045782600780549050039150600090505b81811015610903576108f6610c41565b80806001019150506108e6565b5b82600281905550505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61093e61053e565b341015151561094c57600080fd5b610954610859565b151561096357610962610c41565b5b61096d3334610c8e565b565b60055481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109d057600080fd5b8060018190555050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b610a3d33610be8565b1515610a4857600080fd5b610a5133610f91565b565b6000806000600784815481101515610a6757fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806002016001826002018054905003815481101515610af157fe5b906000526020600020906002020160000154915050915091565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b6857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515610bdf57816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60019050919050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b610c8c6007600160078054905003815481101515610c5b57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610f91565b565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078054806001018281610ceb9190611188565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff021916908315150217905550836002018054806001018281610d7591906111b4565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115610f815760066000600760018603815481101515610dde57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915081600201600160078054905003815481101515610e5f57fe5b90600052602060002090600202019050806000015485111515610e8157610f81565b600760018403815481101515610e9357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600784815481101515610ece57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600760018503815481101515610f2a57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828060019003935050610dbf565b610f89611105565b505050505050565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b60016007805490500381101561109057600760018201815481101515610fff57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078281548110151561103a57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050610fdd565b60078054809190600190036110a591906111e6565b5060008260010160006101000a81548160ff02191690831515021790555060035442018260020160018460020180549050038154811015156110e357fe5b906000526020600020906002020160010181905550611100611105565b505050565b600760405180828054801561116f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611125575b5050915050604051809103902060058160001916905550565b8154818355818115116111af578183600052602060002091820191016111ae9190611212565b5b505050565b8154818355818115116111e1576002028160020283600052602060002091820191016111e09190611237565b5b505050565b81548183558181151161120d5781836000526020600020918201910161120c9190611212565b5b505050565b61123491905b80821115611230576000816000905550600101611218565b5090565b90565b61126391905b8082111561125f5760008082016000905560018201600090555060020161123d565b5090565b905600a165627a7a72305820c8de4548da02d24b5dc600f17d84bb221f1ebc2b6421185da881754eead230a20029`

// DeployElectionContract deploys a new Ethereum contract, binding an instance of ElectionContract to it.
func DeployElectionContract(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxValidators *big.Int, _unbondingPeriod *big.Int, _genesis common.Address) (common.Address, *types.Transaction, *ElectionContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElectionContractBin), backend, _baseDeposit, _maxValidators, _unbondingPeriod, _genesis)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElectionContract{ElectionContractCaller: ElectionContractCaller{contract: contract}, ElectionContractTransactor: ElectionContractTransactor{contract: contract}}, nil
}

// ElectionContract is an auto generated Go binding around an Ethereum contract.
type ElectionContract struct {
	ElectionContractCaller     // Read-only binding to the contract
	ElectionContractTransactor // Write-only binding to the contract
}

// ElectionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionContractSession struct {
	Contract     *ElectionContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionContractCallerSession struct {
	Contract *ElectionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ElectionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionContractTransactorSession struct {
	Contract     *ElectionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ElectionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionContractRaw struct {
	Contract *ElectionContract // Generic contract binding to access the raw methods on
}

// ElectionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionContractCallerRaw struct {
	Contract *ElectionContractCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionContractTransactorRaw struct {
	Contract *ElectionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElectionContract creates a new instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContract(address common.Address, backend bind.ContractBackend) (*ElectionContract, error) {
	contract, err := bindElectionContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElectionContract{ElectionContractCaller: ElectionContractCaller{contract: contract}, ElectionContractTransactor: ElectionContractTransactor{contract: contract}}, nil
}

// NewElectionContractCaller creates a new read-only instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContractCaller(address common.Address, caller bind.ContractCaller) (*ElectionContractCaller, error) {
	contract, err := bindElectionContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionContractCaller{contract: contract}, nil
}

// NewElectionContractTransactor creates a new write-only instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionContractTransactor, error) {
	contract, err := bindElectionContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ElectionContractTransactor{contract: contract}, nil
}

// bindElectionContract binds a generic wrapper to an already deployed contract.
func bindElectionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionContract *ElectionContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElectionContract.Contract.ElectionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionContract *ElectionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.Contract.ElectionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionContract *ElectionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionContract.Contract.ElectionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionContract *ElectionContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElectionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionContract *ElectionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionContract *ElectionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionContract.Contract.contract.Transact(opts, method, params...)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ElectionContract *ElectionContractCaller) _hasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ElectionContract *ElectionContractSession) _hasAvailability() (bool, error) {
	return _ElectionContract.Contract._hasAvailability(&_ElectionContract.CallOpts)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ElectionContract *ElectionContractCallerSession) _hasAvailability() (bool, error) {
	return _ElectionContract.Contract._hasAvailability(&_ElectionContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ElectionContract *ElectionContractCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ElectionContract *ElectionContractSession) BaseDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.BaseDeposit(&_ElectionContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) BaseDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.BaseDeposit(&_ElectionContract.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_ElectionContract *ElectionContractCaller) Genesis(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "genesis")
	return *ret0, err
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_ElectionContract *ElectionContractSession) Genesis() (common.Address, error) {
	return _ElectionContract.Contract.Genesis(&_ElectionContract.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_ElectionContract *ElectionContractCallerSession) Genesis() (common.Address, error) {
	return _ElectionContract.Contract.Genesis(&_ElectionContract.CallOpts)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_ElectionContract *ElectionContractCaller) GetDepositAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	ret := new(struct {
		Amount     *big.Int
		ReleasedAt *big.Int
	})
	out := ret
	err := _ElectionContract.contract.Call(opts, out, "getDepositAtIndex", index)
	return *ret, err
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_ElectionContract *ElectionContractSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	return _ElectionContract.Contract.GetDepositAtIndex(&_ElectionContract.CallOpts, index)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_ElectionContract *ElectionContractCallerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	return _ElectionContract.Contract.GetDepositAtIndex(&_ElectionContract.CallOpts, index)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCaller) GetDepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getDepositCount")
	return *ret0, err
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractSession) GetDepositCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetDepositCount(&_ElectionContract.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCallerSession) GetDepositCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetDepositCount(&_ElectionContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ElectionContract *ElectionContractCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ElectionContract *ElectionContractSession) GetMinimumDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.GetMinimumDeposit(&_ElectionContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ElectionContract *ElectionContractCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.GetMinimumDeposit(&_ElectionContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ElectionContract *ElectionContractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ElectionContract *ElectionContractSession) GetOwner() (common.Address, error) {
	return _ElectionContract.Contract.GetOwner(&_ElectionContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ElectionContract *ElectionContractCallerSession) GetOwner() (common.Address, error) {
	return _ElectionContract.Contract.GetOwner(&_ElectionContract.CallOpts)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ElectionContract *ElectionContractCaller) GetValidatorAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _ElectionContract.contract.Call(opts, out, "getValidatorAtIndex", index)
	return *ret, err
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ElectionContract *ElectionContractSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ElectionContract.Contract.GetValidatorAtIndex(&_ElectionContract.CallOpts, index)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ElectionContract *ElectionContractCallerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ElectionContract.Contract.GetValidatorAtIndex(&_ElectionContract.CallOpts, index)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getValidatorCount")
	return *ret0, err
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractSession) GetValidatorCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetValidatorCount(&_ElectionContract.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCallerSession) GetValidatorCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetValidatorCount(&_ElectionContract.CallOpts)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCaller) IsGenesisValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "isGenesisValidator", code)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsGenesisValidator(&_ElectionContract.CallOpts, code)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCallerSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsGenesisValidator(&_ElectionContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCaller) IsValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "isValidator", code)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractSession) IsValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsValidator(&_ElectionContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCallerSession) IsValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsValidator(&_ElectionContract.CallOpts, code)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ElectionContract *ElectionContractCaller) MaxValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "maxValidators")
	return *ret0, err
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ElectionContract *ElectionContractSession) MaxValidators() (*big.Int, error) {
	return _ElectionContract.Contract.MaxValidators(&_ElectionContract.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) MaxValidators() (*big.Int, error) {
	return _ElectionContract.Contract.MaxValidators(&_ElectionContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_ElectionContract *ElectionContractCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_ElectionContract *ElectionContractSession) UnbondingPeriod() (*big.Int, error) {
	return _ElectionContract.Contract.UnbondingPeriod(&_ElectionContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _ElectionContract.Contract.UnbondingPeriod(&_ElectionContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ElectionContract *ElectionContractCaller) ValidatorsChecksum(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "validatorsChecksum")
	return *ret0, err
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ElectionContract *ElectionContractSession) ValidatorsChecksum() ([32]byte, error) {
	return _ElectionContract.Contract.ValidatorsChecksum(&_ElectionContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ElectionContract *ElectionContractCallerSession) ValidatorsChecksum() ([32]byte, error) {
	return _ElectionContract.Contract.ValidatorsChecksum(&_ElectionContract.CallOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_ElectionContract *ElectionContractTransactor) Join(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "join")
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_ElectionContract *ElectionContractSession) Join() (*types.Transaction, error) {
	return _ElectionContract.Contract.Join(&_ElectionContract.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_ElectionContract *ElectionContractTransactorSession) Join() (*types.Transaction, error) {
	return _ElectionContract.Contract.Join(&_ElectionContract.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_ElectionContract *ElectionContractTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_ElectionContract *ElectionContractSession) Leave() (*types.Transaction, error) {
	return _ElectionContract.Contract.Leave(&_ElectionContract.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_ElectionContract *ElectionContractTransactorSession) Leave() (*types.Transaction, error) {
	return _ElectionContract.Contract.Leave(&_ElectionContract.TransactOpts)
}

// RedeemFunds is a paid mutator transaction binding the contract method 0x216355ec.
//
// Solidity: function redeemFunds() returns(refund uint256)
func (_ElectionContract *ElectionContractTransactor) RedeemFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "redeemFunds")
}

// RedeemFunds is a paid mutator transaction binding the contract method 0x216355ec.
//
// Solidity: function redeemFunds() returns(refund uint256)
func (_ElectionContract *ElectionContractSession) RedeemFunds() (*types.Transaction, error) {
	return _ElectionContract.Contract.RedeemFunds(&_ElectionContract.TransactOpts)
}

// RedeemFunds is a paid mutator transaction binding the contract method 0x216355ec.
//
// Solidity: function redeemFunds() returns(refund uint256)
func (_ElectionContract *ElectionContractTransactorSession) RedeemFunds() (*types.Transaction, error) {
	return _ElectionContract.Contract.RedeemFunds(&_ElectionContract.TransactOpts)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ElectionContract *ElectionContractTransactor) SetBaseDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "setBaseDeposit", deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ElectionContract *ElectionContractSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetBaseDeposit(&_ElectionContract.TransactOpts, deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ElectionContract *ElectionContractTransactorSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetBaseDeposit(&_ElectionContract.TransactOpts, deposit)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ElectionContract *ElectionContractTransactor) SetMaxValidators(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "setMaxValidators", max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ElectionContract *ElectionContractSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetMaxValidators(&_ElectionContract.TransactOpts, max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ElectionContract *ElectionContractTransactorSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetMaxValidators(&_ElectionContract.TransactOpts, max)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_ElectionContract *ElectionContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_ElectionContract *ElectionContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.TransferOwnership(&_ElectionContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_ElectionContract *ElectionContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.TransferOwnership(&_ElectionContract.TransactOpts, _newOwner)
}
