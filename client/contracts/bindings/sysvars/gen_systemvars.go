// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sysvars

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// SystemVarsABI is the input ABI used to generate the binding from.
const SystemVarsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"oracleDeduction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencySupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevCurrencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"uint256\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SystemVarsBin is the compiled bytecode used for deploying new contracts.
const SystemVarsBin = `608060405234801561001057600080fd5b5060405160408061041483398101806040528101908080519060200190929190805190602001909291905050508160008190555081600181905550806003819055508060028190555050506103aa8061006a6000396000f30060806040526004361061008d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062bf32ca1461009257806321873631146100d35780632af4f9c0146100fe5780632d380242146101295780636df566d714610154578063a035b1fe1461017f578063b0c6363d146101aa578063fc634f4b146101d5575b600080fd5b34801561009e57600080fd5b506100bd60048036038101908080359060200190929190505050610200565b6040518082815260200191505060405180910390f35b3480156100df57600080fd5b506100e8610219565b6040518082815260200191505060405180910390f35b34801561010a57600080fd5b50610113610249565b6040518082815260200191505060405180910390f35b34801561013557600080fd5b5061013e61024f565b6040518082815260200191505060405180910390f35b34801561016057600080fd5b506101696102d7565b6040518082815260200191505060405180910390f35b34801561018b57600080fd5b506101946102dd565b6040518082815260200191505060405180910390f35b3480156101b657600080fd5b506101bf6102e7565b6040518082815260200191505060405180910390f35b3480156101e157600080fd5b506101ea6102ed565b6040518082815260200191505060405180910390f35b600060648260040281151561021157fe5b049050919050565b6000610244670de0b6b3a76400003073ffffffffffffffffffffffffffffffffffffffff16316102f3565b905090565b60035481565b6000806001804301141561026e57680246ddf9797668000091506102d3565b61271060035481151561027d57fe5b04905060005460015411801561029c5750670de0b6b3a7640000600054115b156102bd576102b681600354016102b161030c565b6102f3565b91506102d3565b6102d0816003540364e8d4a5100061034d565b91505b5090565b60015481565b6000600154905090565b60025481565b60005481565b60008183106103025781610304565b825b905092915050565b600060018043011180156103245750610323610367565b5b61033757680471fa858b9e080000610348565b61271060025481151561034657fe5b045b905090565b60008183101561035d578161035f565b825b905092915050565b600069d3c21bcecceda100000060025410159050905600a165627a7a7230582032afe135b275dd9af341167c5c6dbdbcc9a3aa4c81db37430c971692d7c370e20029`

// DeploySystemVars deploys a new Kowala contract, binding an instance of SystemVars to it.
func DeploySystemVars(auth *bind.TransactOpts, backend bind.ContractBackend, _initialPrice *big.Int, _initialSupply *big.Int) (common.Address, *types.Transaction, *SystemVars, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemVarsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemVarsBin), backend, _initialPrice, _initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemVars{SystemVarsCaller: SystemVarsCaller{contract: contract}, SystemVarsTransactor: SystemVarsTransactor{contract: contract}, SystemVarsFilterer: SystemVarsFilterer{contract: contract}}, nil
}

// SystemVars is an auto generated Go binding around a Kowala contract.
type SystemVars struct {
	SystemVarsCaller     // Read-only binding to the contract
	SystemVarsTransactor // Write-only binding to the contract
	SystemVarsFilterer   // Log filterer for contract events
}

// SystemVarsCaller is an auto generated read-only Go binding around a Kowala contract.
type SystemVarsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemVarsTransactor is an auto generated write-only Go binding around a Kowala contract.
type SystemVarsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemVarsFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type SystemVarsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemVarsSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type SystemVarsSession struct {
	Contract     *SystemVars       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemVarsCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type SystemVarsCallerSession struct {
	Contract *SystemVarsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SystemVarsTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type SystemVarsTransactorSession struct {
	Contract     *SystemVarsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SystemVarsRaw is an auto generated low-level Go binding around a Kowala contract.
type SystemVarsRaw struct {
	Contract *SystemVars // Generic contract binding to access the raw methods on
}

// SystemVarsCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type SystemVarsCallerRaw struct {
	Contract *SystemVarsCaller // Generic read-only contract binding to access the raw methods on
}

// SystemVarsTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type SystemVarsTransactorRaw struct {
	Contract *SystemVarsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemVars creates a new instance of SystemVars, bound to a specific deployed contract.
func NewSystemVars(address common.Address, backend bind.ContractBackend) (*SystemVars, error) {
	contract, err := bindSystemVars(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemVars{SystemVarsCaller: SystemVarsCaller{contract: contract}, SystemVarsTransactor: SystemVarsTransactor{contract: contract}, SystemVarsFilterer: SystemVarsFilterer{contract: contract}}, nil
}

// NewSystemVarsCaller creates a new read-only instance of SystemVars, bound to a specific deployed contract.
func NewSystemVarsCaller(address common.Address, caller bind.ContractCaller) (*SystemVarsCaller, error) {
	contract, err := bindSystemVars(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemVarsCaller{contract: contract}, nil
}

// NewSystemVarsTransactor creates a new write-only instance of SystemVars, bound to a specific deployed contract.
func NewSystemVarsTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemVarsTransactor, error) {
	contract, err := bindSystemVars(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemVarsTransactor{contract: contract}, nil
}

// NewSystemVarsFilterer creates a new log filterer instance of SystemVars, bound to a specific deployed contract.
func NewSystemVarsFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemVarsFilterer, error) {
	contract, err := bindSystemVars(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemVarsFilterer{contract: contract}, nil
}

// bindSystemVars binds a generic wrapper to an already deployed contract.
func bindSystemVars(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemVarsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemVars *SystemVarsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemVars.Contract.SystemVarsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemVars *SystemVarsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemVars.Contract.SystemVarsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemVars *SystemVarsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemVars.Contract.SystemVarsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemVars *SystemVarsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemVars.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemVars *SystemVarsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemVars.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemVars *SystemVarsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemVars.Contract.contract.Transact(opts, method, params...)
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) CurrencyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "currencyPrice")
	return *ret0, err
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsSession) CurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.CurrencyPrice(&_SystemVars.CallOpts)
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) CurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.CurrencyPrice(&_SystemVars.CallOpts)
}

// CurrencySupply is a free data retrieval call binding the contract method 0xb0c6363d.
//
// Solidity: function currencySupply() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) CurrencySupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "currencySupply")
	return *ret0, err
}

// CurrencySupply is a free data retrieval call binding the contract method 0xb0c6363d.
//
// Solidity: function currencySupply() constant returns(uint256)
func (_SystemVars *SystemVarsSession) CurrencySupply() (*big.Int, error) {
	return _SystemVars.Contract.CurrencySupply(&_SystemVars.CallOpts)
}

// CurrencySupply is a free data retrieval call binding the contract method 0xb0c6363d.
//
// Solidity: function currencySupply() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) CurrencySupply() (*big.Int, error) {
	return _SystemVars.Contract.CurrencySupply(&_SystemVars.CallOpts)
}

// MintedAmount is a free data retrieval call binding the contract method 0x2d380242.
//
// Solidity: function mintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) MintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "mintedAmount")
	return *ret0, err
}

// MintedAmount is a free data retrieval call binding the contract method 0x2d380242.
//
// Solidity: function mintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsSession) MintedAmount() (*big.Int, error) {
	return _SystemVars.Contract.MintedAmount(&_SystemVars.CallOpts)
}

// MintedAmount is a free data retrieval call binding the contract method 0x2d380242.
//
// Solidity: function mintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) MintedAmount() (*big.Int, error) {
	return _SystemVars.Contract.MintedAmount(&_SystemVars.CallOpts)
}

// MintedReward is a free data retrieval call binding the contract method 0x2af4f9c0.
//
// Solidity: function mintedReward() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) MintedReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "mintedReward")
	return *ret0, err
}

// MintedReward is a free data retrieval call binding the contract method 0x2af4f9c0.
//
// Solidity: function mintedReward() constant returns(uint256)
func (_SystemVars *SystemVarsSession) MintedReward() (*big.Int, error) {
	return _SystemVars.Contract.MintedReward(&_SystemVars.CallOpts)
}

// MintedReward is a free data retrieval call binding the contract method 0x2af4f9c0.
//
// Solidity: function mintedReward() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) MintedReward() (*big.Int, error) {
	return _SystemVars.Contract.MintedReward(&_SystemVars.CallOpts)
}

// OracleDeduction is a free data retrieval call binding the contract method 0x00bf32ca.
//
// Solidity: function oracleDeduction(mintedAmount uint256) constant returns(uint256)
func (_SystemVars *SystemVarsCaller) OracleDeduction(opts *bind.CallOpts, mintedAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "oracleDeduction", mintedAmount)
	return *ret0, err
}

// OracleDeduction is a free data retrieval call binding the contract method 0x00bf32ca.
//
// Solidity: function oracleDeduction(mintedAmount uint256) constant returns(uint256)
func (_SystemVars *SystemVarsSession) OracleDeduction(mintedAmount *big.Int) (*big.Int, error) {
	return _SystemVars.Contract.OracleDeduction(&_SystemVars.CallOpts, mintedAmount)
}

// OracleDeduction is a free data retrieval call binding the contract method 0x00bf32ca.
//
// Solidity: function oracleDeduction(mintedAmount uint256) constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) OracleDeduction(mintedAmount *big.Int) (*big.Int, error) {
	return _SystemVars.Contract.OracleDeduction(&_SystemVars.CallOpts, mintedAmount)
}

// OracleReward is a free data retrieval call binding the contract method 0x21873631.
//
// Solidity: function oracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) OracleReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "oracleReward")
	return *ret0, err
}

// OracleReward is a free data retrieval call binding the contract method 0x21873631.
//
// Solidity: function oracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsSession) OracleReward() (*big.Int, error) {
	return _SystemVars.Contract.OracleReward(&_SystemVars.CallOpts)
}

// OracleReward is a free data retrieval call binding the contract method 0x21873631.
//
// Solidity: function oracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) OracleReward() (*big.Int, error) {
	return _SystemVars.Contract.OracleReward(&_SystemVars.CallOpts)
}

// PrevCurrencyPrice is a free data retrieval call binding the contract method 0xfc634f4b.
//
// Solidity: function prevCurrencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) PrevCurrencyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "prevCurrencyPrice")
	return *ret0, err
}

// PrevCurrencyPrice is a free data retrieval call binding the contract method 0xfc634f4b.
//
// Solidity: function prevCurrencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsSession) PrevCurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.PrevCurrencyPrice(&_SystemVars.CallOpts)
}

// PrevCurrencyPrice is a free data retrieval call binding the contract method 0xfc634f4b.
//
// Solidity: function prevCurrencyPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) PrevCurrencyPrice() (*big.Int, error) {
	return _SystemVars.Contract.PrevCurrencyPrice(&_SystemVars.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(price uint256)
func (_SystemVars *SystemVarsCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(price uint256)
func (_SystemVars *SystemVarsSession) Price() (*big.Int, error) {
	return _SystemVars.Contract.Price(&_SystemVars.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(price uint256)
func (_SystemVars *SystemVarsCallerSession) Price() (*big.Int, error) {
	return _SystemVars.Contract.Price(&_SystemVars.CallOpts)
}
