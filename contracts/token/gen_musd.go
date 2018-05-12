// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
)

// MUSDABI is the input ABI used to generate the binding from.
const MUSDABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_custom_fallback\",\"type\":\"string\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MUSDBin is the compiled bytecode used for deploying new contracts.
const MUSDBin = `60606040526000600660146101000a81548160ff02191690831515021790555034156200002b57600080fd5b6b03782dace9d900000000000033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000811115156200008957600080fd5b80600781905550506040805190810160405280600481526020017f6d5553440000000000000000000000000000000000000000000000000000000081525060029080519060200190620000de929190620004a1565b506040805190810160405280600481526020017f6d55534400000000000000000000000000000000000000000000000000000000815250600390805190602001906200012c929190620004a1565b506012600460006101000a81548160ff021916908360ff160217905550620001887380eda603028fe504b57d14d947c8087c1798d8006a853a0d2313c000000000006200018f64010000000002620006ba176401000000009004565b5062000550565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515620001ee57600080fd5b600660149054906101000a900460ff161515156200020b57600080fd5b6007546200023383600554620002696401000000000262000e47179091906401000000009004565b111515156200024157600080fd5b620002618383620002866401000000000262000e63176401000000009004565b905092915050565b600081830190508281101515156200027d57fe5b80905092915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515620002e557600080fd5b600660149054906101000a900460ff161515156200030257600080fd5b6200032782600554620002696401000000000262000e47179091906401000000009004565b6005819055506200038f82600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620002696401000000000262000e47179091906401000000009004565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a260405180600001905060405180910390208373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16856040518082815260200191505060405180910390a46001905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004e457805160ff191683800117855562000515565b8280016001018555821562000515579182015b8281111562000514578251825591602001919060010190620004f7565b5b50905062000524919062000528565b5090565b6200054d91905b80821115620005495760008160009055506001016200052f565b5090565b90565b61152080620005606000396000f3006060604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806305d2035b146100d557806306fdde031461010257806318160ddd14610190578063313ce567146101b9578063355274ea146101e857806340c10f191461021157806370a082311461026b5780637d64bcb4146102b85780638da5cb5b146102e557806395d89b411461033a578063a9059cbb146103c8578063be45fd6214610422578063f2fde38b146104bf578063f6368f8a146104f8575b600080fd5b34156100e057600080fd5b6100e86105d8565b604051808215151515815260200191505060405180910390f35b341561010d57600080fd5b6101156105eb565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561015557808201518184015260208101905061013a565b50505050905090810190601f1680156101825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561019b57600080fd5b6101a3610693565b6040518082815260200191505060405180910390f35b34156101c457600080fd5b6101cc61069d565b604051808260ff1660ff16815260200191505060405180910390f35b34156101f357600080fd5b6101fb6106b4565b6040518082815260200191505060405180910390f35b341561021c57600080fd5b610251600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506106ba565b604051808215151515815260200191505060405180910390f35b341561027657600080fd5b6102a2600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061076b565b6040518082815260200191505060405180910390f35b34156102c357600080fd5b6102cb6107b4565b604051808215151515815260200191505060405180910390f35b34156102f057600080fd5b6102f861087c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561034557600080fd5b61034d6108a2565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561038d578082015181840152602081019050610372565b50505050905090810190601f1680156103ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103d357600080fd5b610408600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061094a565b604051808215151515815260200191505060405180910390f35b341561042d57600080fd5b6104a5600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610989565b604051808215151515815260200191505060405180910390f35b34156104ca57600080fd5b6104f6600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506109c0565b005b341561050357600080fd5b6105be600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610b18565b604051808215151515815260200191505060405180910390f35b600660149054906101000a900460ff1681565b6105f36114cc565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106895780601f1061065e57610100808354040283529160200191610689565b820191906000526020600020905b81548152906001019060200180831161066c57829003601f168201915b5050505050905090565b6000600554905090565b6000600460009054906101000a900460ff16905090565b60075481565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561071857600080fd5b600660149054906101000a900460ff1615151561073457600080fd5b60075461074c83600554610e4790919063ffffffff16565b1115151561075957600080fd5b6107638383610e63565b905092915050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561081257600080fd5b600660149054906101000a900460ff1615151561082e57600080fd5b6001600660146101000a81548160ff0219169083151502179055507fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a16001905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6108aa6114cc565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109405780601f1061091557610100808354040283529160200191610940565b820191906000526020600020905b81548152906001019060200180831161092357829003601f168201915b5050505050905090565b60006109546114e0565b61095d8461105c565b156109745761096d84848361106f565b9150610982565b61097f848483611323565b91505b5092915050565b60006109948461105c565b156109ab576109a484848461106f565b90506109b9565b6109b6848484611323565b90505b9392505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a1c57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610a5857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000610b238561105c565b15610e315783610b323361076b565b1015610b3d57600080fd5b610b8f84600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546114b390919063ffffffff16565b50610be284600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e4790919063ffffffff16565b508473ffffffffffffffffffffffffffffffffffffffff166000836040518082805190602001908083835b602083101515610c325780518252602082019150602081019050602083039250610c0d565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207c01000000000000000000000000000000000000000000000000000000009004903387876040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828051906020019080838360005b83811015610d13578082015181840152602081019050610cf8565b50505050905090810190601f168015610d405780820380516001836020036101000a031916815260200191505b50935050505060006040518083038185885af193505050501515610d6057fe5b826040518082805190602001908083835b602083101515610d965780518252602082019150602081019050602083039250610d71565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16876040518082815260200191505060405180910390a460019050610e3f565b610e3c858585611323565b90505b949350505050565b60008183019050828110151515610e5a57fe5b80905092915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ec157600080fd5b600660149054906101000a900460ff16151515610edd57600080fd5b610ef282600554610e4790919063ffffffff16565b600581905550610f4a82600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e4790919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a260405180600001905060405180910390208373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16856040518082815260200191505060405180910390a46001905092915050565b600080823b905060008111915050919050565b6000808361107c3361076b565b101561108757600080fd5b6110d984600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546114b390919063ffffffff16565b5061112c84600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e4790919063ffffffff16565b508490508073ffffffffffffffffffffffffffffffffffffffff1663c0ee0b8a3386866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111f25780820151818401526020810190506111d7565b50505050905090810190601f16801561121f5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b151561123f57600080fd5b5af1151561124c57600080fd5b505050826040518082805190602001908083835b6020831015156112855780518252602082019150602081019050602083039250611260565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16876040518082815260200191505060405180910390a460019150509392505050565b60008261132f3361076b565b101561133a57600080fd5b61138c83600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546114b390919063ffffffff16565b506113df83600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e4790919063ffffffff16565b50816040518082805190602001908083835b60208310151561141657805182526020820191506020810190506020830392506113f1565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16866040518082815260200191505060405180910390a4600190509392505050565b60008282111515156114c157fe5b818303905092915050565b602060405190810160405280600081525090565b6020604051908101604052806000815250905600a165627a7a7230582036e9bdda21182d2b0d9d0755e0874b7914f74718de255d889c3a75364eac1b7d0029`

// DeployMUSD deploys a new Ethereum contract, binding an instance of MUSD to it.
func DeployMUSD(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(MUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MUSDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MUSD{MUSDCaller: MUSDCaller{contract: contract}, MUSDTransactor: MUSDTransactor{contract: contract}}, nil
}

// MUSD is an auto generated Go binding around an Ethereum contract.
type MUSD struct {
	MUSDCaller     // Read-only binding to the contract
	MUSDTransactor // Write-only binding to the contract
}

// MUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type MUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MUSDSession struct {
	Contract     *MUSD             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MUSDCallerSession struct {
	Contract *MUSDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MUSDTransactorSession struct {
	Contract     *MUSDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type MUSDRaw struct {
	Contract *MUSD // Generic contract binding to access the raw methods on
}

// MUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MUSDCallerRaw struct {
	Contract *MUSDCaller // Generic read-only contract binding to access the raw methods on
}

// MUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MUSDTransactorRaw struct {
	Contract *MUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMUSD creates a new instance of MUSD, bound to a specific deployed contract.
func NewMUSD(address common.Address, backend bind.ContractBackend) (*MUSD, error) {
	contract, err := bindMUSD(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MUSD{MUSDCaller: MUSDCaller{contract: contract}, MUSDTransactor: MUSDTransactor{contract: contract}}, nil
}

// NewMUSDCaller creates a new read-only instance of MUSD, bound to a specific deployed contract.
func NewMUSDCaller(address common.Address, caller bind.ContractCaller) (*MUSDCaller, error) {
	contract, err := bindMUSD(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MUSDCaller{contract: contract}, nil
}

// NewMUSDTransactor creates a new write-only instance of MUSD, bound to a specific deployed contract.
func NewMUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*MUSDTransactor, error) {
	contract, err := bindMUSD(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MUSDTransactor{contract: contract}, nil
}

// bindMUSD binds a generic wrapper to an already deployed contract.
func bindMUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MUSD *MUSDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MUSD.Contract.MUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MUSD *MUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MUSD.Contract.MUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MUSD *MUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MUSD.Contract.MUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MUSD *MUSDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MUSD *MUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MUSD *MUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MUSD.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MUSD *MUSDCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MUSD *MUSDSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MUSD.Contract.BalanceOf(&_MUSD.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MUSD *MUSDCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MUSD.Contract.BalanceOf(&_MUSD.CallOpts, _owner)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MUSD *MUSDCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MUSD *MUSDSession) Cap() (*big.Int, error) {
	return _MUSD.Contract.Cap(&_MUSD.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_MUSD *MUSDCallerSession) Cap() (*big.Int, error) {
	return _MUSD.Contract.Cap(&_MUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MUSD *MUSDCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MUSD *MUSDSession) Decimals() (uint8, error) {
	return _MUSD.Contract.Decimals(&_MUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(_decimals uint8)
func (_MUSD *MUSDCallerSession) Decimals() (uint8, error) {
	return _MUSD.Contract.Decimals(&_MUSD.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MUSD *MUSDCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MUSD *MUSDSession) MintingFinished() (bool, error) {
	return _MUSD.Contract.MintingFinished(&_MUSD.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MUSD *MUSDCallerSession) MintingFinished() (bool, error) {
	return _MUSD.Contract.MintingFinished(&_MUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MUSD *MUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MUSD *MUSDSession) Name() (string, error) {
	return _MUSD.Contract.Name(&_MUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_MUSD *MUSDCallerSession) Name() (string, error) {
	return _MUSD.Contract.Name(&_MUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MUSD *MUSDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MUSD *MUSDSession) Owner() (common.Address, error) {
	return _MUSD.Contract.Owner(&_MUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MUSD *MUSDCallerSession) Owner() (common.Address, error) {
	return _MUSD.Contract.Owner(&_MUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MUSD *MUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MUSD *MUSDSession) Symbol() (string, error) {
	return _MUSD.Contract.Symbol(&_MUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_MUSD *MUSDCallerSession) Symbol() (string, error) {
	return _MUSD.Contract.Symbol(&_MUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MUSD *MUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MUSD.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MUSD *MUSDSession) TotalSupply() (*big.Int, error) {
	return _MUSD.Contract.TotalSupply(&_MUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(_totalSupply uint256)
func (_MUSD *MUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _MUSD.Contract.TotalSupply(&_MUSD.CallOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MUSD *MUSDTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MUSD.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MUSD *MUSDSession) FinishMinting() (*types.Transaction, error) {
	return _MUSD.Contract.FinishMinting(&_MUSD.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MUSD *MUSDTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _MUSD.Contract.FinishMinting(&_MUSD.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MUSD *MUSDTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MUSD.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MUSD *MUSDSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MUSD.Contract.Mint(&_MUSD.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MUSD *MUSDTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MUSD.Contract.Mint(&_MUSD.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MUSD *MUSDTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MUSD.contract.Transact(opts, "transfer", _to, _value, _data, _custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MUSD *MUSDSession) Transfer(_to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MUSD.Contract.Transfer(&_MUSD.TransactOpts, _to, _value, _data, _custom_fallback)
}

// Transfer is a paid mutator transaction binding the contract method 0xf6368f8a.
//
// Solidity: function transfer(_to address, _value uint256, _data bytes, _custom_fallback string) returns(success bool)
func (_MUSD *MUSDTransactorSession) Transfer(_to common.Address, _value *big.Int, _data []byte, _custom_fallback string) (*types.Transaction, error) {
	return _MUSD.Contract.Transfer(&_MUSD.TransactOpts, _to, _value, _data, _custom_fallback)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MUSD *MUSDTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MUSD.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MUSD *MUSDSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MUSD.Contract.TransferOwnership(&_MUSD.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MUSD *MUSDTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MUSD.Contract.TransferOwnership(&_MUSD.TransactOpts, newOwner)
}
