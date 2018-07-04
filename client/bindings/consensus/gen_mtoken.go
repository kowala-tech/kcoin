// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// MiningTokenABI is the input ABI used to generate the binding from.
const MiningTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_custom_fallback\",\"type\":\"string\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_cap\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MiningTokenBin is the compiled bytecode used for deploying new contracts.
const MiningTokenBin = `60806040526000600660146101000a81548160ff0219169083151502179055503480156200002c57600080fd5b50604051620019e9380380620019e9833981018060405281019080805182019291906020018051820192919060200180519060200190929190805190602001909291905050508133600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600081111515620000c457600080fd5b80600781905550508360029080519060200190620000e492919062000123565b508260039080519060200190620000fd92919062000123565b5080600460006101000a81548160ff021916908360ff16021790555050505050620001d2565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200016657805160ff191683800117855562000197565b8280016001018555821562000197579182015b828111156200019657825182559160200191906001019062000179565b5b509050620001a69190620001aa565b5090565b620001cf91905b80821115620001cb576000816000905550600101620001b1565b5090565b90565b61180780620001e26000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806305d2035b146100e057806306fdde031461010f57806318160ddd1461019f578063313ce567146101ca578063355274ea146101fb57806340c10f191461022657806370a082311461028b578063715018a6146102e25780637d64bcb4146102f95780638da5cb5b1461032857806395d89b411461037f578063a9059cbb1461040f578063be45fd6214610474578063f2fde38b1461051f578063f6368f8a14610562575b600080fd5b3480156100ec57600080fd5b506100f5610653565b604051808215151515815260200191505060405180910390f35b34801561011b57600080fd5b50610124610666565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610164578082015181840152602081019050610149565b50505050905090810190601f1680156101915780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ab57600080fd5b506101b4610708565b6040518082815260200191505060405180910390f35b3480156101d657600080fd5b506101df610712565b604051808260ff1660ff16815260200191505060405180910390f35b34801561020757600080fd5b50610210610729565b6040518082815260200191505060405180910390f35b34801561023257600080fd5b50610271600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061072f565b604051808215151515815260200191505060405180910390f35b34801561029757600080fd5b506102cc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107e0565b6040518082815260200191505060405180910390f35b3480156102ee57600080fd5b506102f7610829565b005b34801561030557600080fd5b5061030e61092e565b604051808215151515815260200191505060405180910390f35b34801561033457600080fd5b5061033d6109f6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561038b57600080fd5b50610394610a1c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103d45780820151818401526020810190506103b9565b50505050905090810190601f1680156104015780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561041b57600080fd5b5061045a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610abe565b604051808215151515815260200191505060405180910390f35b34801561048057600080fd5b50610505600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610af7565b604051808215151515815260200191505060405180910390f35b34801561052b57600080fd5b50610560600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b2e565b005b34801561056e57600080fd5b50610639600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610b96565b604051808215151515815260200191505060405180910390f35b600660149054906101000a900460ff1681565b606060028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106fe5780601f106106d3576101008083540402835291602001916106fe565b820191906000526020600020905b8154815290600101906020018083116106e157829003601f168201915b5050505050905090565b6000600554905090565b6000600460009054906101000a900460ff16905090565b60075481565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561078d57600080fd5b600660149054906101000a900460ff161515156107a957600080fd5b6007546107c183600554610f4990919063ffffffff16565b111515156107ce57600080fd5b6107d88383610f65565b905092915050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561088557600080fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a26000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561098c57600080fd5b600660149054906101000a900460ff161515156109a857600080fd5b6001600660146101000a81548160ff0219169083151502179055507fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a16001905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ab45780601f10610a8957610100808354040283529160200191610ab4565b820191906000526020600020905b815481529060010190602001808311610a9757829003601f168201915b5050505050905090565b60006060610acb8461115e565b15610ae257610adb848483611171565b9150610af0565b610aed8484836114b2565b91505b5092915050565b6000610b028461115e565b15610b1957610b12848484611171565b9050610b27565b610b248484846114b2565b90505b9392505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b8a57600080fd5b610b93816116c6565b50565b6000610ba18561115e565b15610f335783610bb0336107e0565b1015610bbb57600080fd5b610c0d84600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117c290919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610ca284600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f4990919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff166000836040518082805190602001908083835b602083101515610d345780518252602082019150602081019050602083039250610d0f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207c01000000000000000000000000000000000000000000000000000000009004903387876040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828051906020019080838360005b83811015610e15578082015181840152602081019050610dfa565b50505050905090810190601f168015610e425780820380516001836020036101000a031916815260200191505b50935050505060006040518083038185885af193505050501515610e6257fe5b826040518082805190602001908083835b602083101515610e985780518252602082019150602081019050602083039250610e73565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16876040518082815260200191505060405180910390a460019050610f41565b610f3e8585856114b2565b90505b949350505050565b60008183019050828110151515610f5c57fe5b80905092915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fc357600080fd5b600660149054906101000a900460ff16151515610fdf57600080fd5b610ff482600554610f4990919063ffffffff16565b60058190555061104c82600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f4990919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a260405180600001905060405180910390208373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16856040518082815260200191505060405180910390a46001905092915050565b600080823b905060008111915050919050565b6000808361117e336107e0565b101561118957600080fd5b6111db84600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117c290919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061127084600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f4990919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508490508073ffffffffffffffffffffffffffffffffffffffff1663c0ee0b8a3386866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561137857808201518184015260208101905061135d565b50505050905090810190601f1680156113a55780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b1580156113c657600080fd5b505af11580156113da573d6000803e3d6000fd5b50505050826040518082805190602001908083835b60208310151561141457805182526020820191506020810190506020830392506113ef565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16876040518082815260200191505060405180910390a460019150509392505050565b6000826114be336107e0565b10156114c957600080fd5b61151b83600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117c290919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506115b083600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f4990919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816040518082805190602001908083835b6020831015156116295780518252602082019150602081019050602083039250611604565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16866040518082815260200191505060405180910390a4600190509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561170257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008282111515156117d057fe5b8183039050929150505600a165627a7a72305820928f174cba1660d65216c3071656c9fbdd37d8d11e6d86706bb5fdef6b879a480029`

// DeployMiningToken deploys a new Ethereum contract, binding an instance of MiningToken to it.
func DeployMiningToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _cap *big.Int, _decimals uint8) (common.Address, *types.Transaction, *MiningToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MiningTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MiningTokenBin), backend, _name, _symbol, _cap, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MiningToken{MiningTokenCaller: MiningTokenCaller{contract: contract}, MiningTokenTransactor: MiningTokenTransactor{contract: contract}}, nil
}

// MiningToken is an auto generated Go binding around an Ethereum contract.
type MiningToken struct {
	MiningTokenCaller     // Read-only binding to the contract
	MiningTokenTransactor // Write-only binding to the contract
}

// MiningTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiningTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiningTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiningTokenSession struct {
	Contract     *MiningToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiningTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiningTokenCallerSession struct {
	Contract *MiningTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MiningTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiningTokenTransactorSession struct {
	Contract     *MiningTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MiningTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiningTokenRaw struct {
	Contract *MiningToken // Generic contract binding to access the raw methods on
}

// MiningTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiningTokenCallerRaw struct {
	Contract *MiningTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MiningTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiningTokenTransactorRaw struct {
	Contract *MiningTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiningToken creates a new instance of MiningToken, bound to a specific deployed contract.
func NewMiningToken(address common.Address, backend bind.ContractBackend) (*MiningToken, error) {
	contract, err := bindMiningToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiningToken{MiningTokenCaller: MiningTokenCaller{contract: contract}, MiningTokenTransactor: MiningTokenTransactor{contract: contract}}, nil
}

// NewMiningTokenCaller creates a new read-only instance of MiningToken, bound to a specific deployed contract.
func NewMiningTokenCaller(address common.Address, caller bind.ContractCaller) (*MiningTokenCaller, error) {
	contract, err := bindMiningToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MiningTokenCaller{contract: contract}, nil
}

// NewMiningTokenTransactor creates a new write-only instance of MiningToken, bound to a specific deployed contract.
func NewMiningTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MiningTokenTransactor, error) {
	contract, err := bindMiningToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MiningTokenTransactor{contract: contract}, nil
}

// bindMiningToken binds a generic wrapper to an already deployed contract.
func bindMiningToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiningTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningToken *MiningTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiningToken.Contract.MiningTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningToken *MiningTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.Contract.MiningTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningToken *MiningTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningToken.Contract.MiningTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningToken *MiningTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiningToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningToken *MiningTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningToken *MiningTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MiningToken *MiningTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MiningToken *MiningTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MiningToken.Contract.BalanceOf(&_MiningToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MiningToken *MiningTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MiningToken.Contract.BalanceOf(&_MiningToken.CallOpts, _owner)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MiningToken *MiningTokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MiningToken *MiningTokenSession) Cap() (*big.Int, error) {
	return _MiningToken.Contract.Cap(&_MiningToken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MiningToken *MiningTokenCallerSession) Cap() (*big.Int, error) {
	return _MiningToken.Contract.Cap(&_MiningToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MiningToken *MiningTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MiningToken *MiningTokenSession) Decimals() (uint8, error) {
	return _MiningToken.Contract.Decimals(&_MiningToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MiningToken *MiningTokenCallerSession) Decimals() (uint8, error) {
	return _MiningToken.Contract.Decimals(&_MiningToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MiningToken *MiningTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MiningToken *MiningTokenSession) MintingFinished() (bool, error) {
	return _MiningToken.Contract.MintingFinished(&_MiningToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MiningToken *MiningTokenCallerSession) MintingFinished() (bool, error) {
	return _MiningToken.Contract.MintingFinished(&_MiningToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MiningToken *MiningTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MiningToken *MiningTokenSession) Name() (string, error) {
	return _MiningToken.Contract.Name(&_MiningToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MiningToken *MiningTokenCallerSession) Name() (string, error) {
	return _MiningToken.Contract.Name(&_MiningToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MiningToken *MiningTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MiningToken *MiningTokenSession) Owner() (common.Address, error) {
	return _MiningToken.Contract.Owner(&_MiningToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MiningToken *MiningTokenCallerSession) Owner() (common.Address, error) {
	return _MiningToken.Contract.Owner(&_MiningToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MiningToken *MiningTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MiningToken *MiningTokenSession) Symbol() (string, error) {
	return _MiningToken.Contract.Symbol(&_MiningToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MiningToken *MiningTokenCallerSession) Symbol() (string, error) {
	return _MiningToken.Contract.Symbol(&_MiningToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MiningToken *MiningTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MiningToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MiningToken *MiningTokenSession) TotalSupply() (*big.Int, error) {
	return _MiningToken.Contract.TotalSupply(&_MiningToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MiningToken *MiningTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MiningToken.Contract.TotalSupply(&_MiningToken.CallOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MiningToken *MiningTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MiningToken *MiningTokenSession) FinishMinting() (*types.Transaction, error) {
	return _MiningToken.Contract.FinishMinting(&_MiningToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MiningToken *MiningTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _MiningToken.Contract.FinishMinting(&_MiningToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MiningToken *MiningTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MiningToken *MiningTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MiningToken.Contract.Mint(&_MiningToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MiningToken *MiningTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MiningToken.Contract.Mint(&_MiningToken.TransactOpts, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiningToken *MiningTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiningToken *MiningTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _MiningToken.Contract.RenounceOwnership(&_MiningToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MiningToken *MiningTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MiningToken.Contract.RenounceOwnership(&_MiningToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MiningToken *MiningTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "transfer", _to, _value, _data, _custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MiningToken *MiningTokenSession) Transfer(_to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MiningToken.Contract.Transfer(&_MiningToken.TransactOpts, _to, _value, _data, _custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MiningToken *MiningTokenTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MiningToken.Contract.Transfer(&_MiningToken.TransactOpts, _to, _value, _data, _custom_fallback)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MiningToken *MiningTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MiningToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MiningToken *MiningTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MiningToken.Contract.TransferOwnership(&_MiningToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MiningToken *MiningTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MiningToken.Contract.TransferOwnership(&_MiningToken.TransactOpts, _newOwner)
}
