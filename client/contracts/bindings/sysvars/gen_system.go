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
const SystemVarsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"oracleDeduction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lowSupplyMetric\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stabilizedPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleDeductionFraction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxUnderNormalConditions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencySupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialMintedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adjustmentFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultOracleReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevCurrencyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SystemVarsBin is the compiled bytecode used for deploying new contracts.
const SystemVarsBin = `608060405234801561001057600080fd5b506040516040806105d8833981018060405281019080805190602001909291908051906020019092919050505081600081905550816001819055508060038190555080600281905550505061056e8061006a6000396000f3006080604052600436106100da576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062bf32ca146100df57806321873631146101205780632af4f9c01461014b5780632d380242146101765780636df566d7146101a157806374bf78e9146101cc5780639d26e501146101f7578063a17aad9614610222578063a47e456f1461024d578063b0c6363d14610278578063b811215e146102a3578063bf053db1146102ce578063cb6741ed146102f9578063ef6711a514610324578063fc634f4b1461034f575b600080fd5b3480156100eb57600080fd5b5061010a6004803603810190808035906020019092919050505061037a565b6040518082815260200191505060405180910390f35b34801561012c57600080fd5b50610135610393565b6040518082815260200191505060405180910390f35b34801561015757600080fd5b506101606103c3565b6040518082815260200191505060405180910390f35b34801561018257600080fd5b5061018b6103c9565b6040518082815260200191505060405180910390f35b3480156101ad57600080fd5b506101b6610451565b6040518082815260200191505060405180910390f35b3480156101d857600080fd5b506101e1610457565b6040518082815260200191505060405180910390f35b34801561020357600080fd5b5061020c610465565b6040518082815260200191505060405180910390f35b34801561022e57600080fd5b50610237610471565b6040518082815260200191505060405180910390f35b34801561025957600080fd5b50610262610476565b6040518082815260200191505060405180910390f35b34801561028457600080fd5b5061028d61047f565b6040518082815260200191505060405180910390f35b3480156102af57600080fd5b506102b8610485565b6040518082815260200191505060405180910390f35b3480156102da57600080fd5b506102e3610492565b6040518082815260200191505060405180910390f35b34801561030557600080fd5b5061030e61049f565b6040518082815260200191505060405180910390f35b34801561033057600080fd5b506103396104a5565b6040518082815260200191505060405180910390f35b34801561035b57600080fd5b506103646104b1565b6040518082815260200191505060405180910390f35b600060648260040281151561038b57fe5b049050919050565b60006103be670de0b6b3a76400003073ffffffffffffffffffffffffffffffffffffffff16316104b7565b905090565b60035481565b600080600180430114156103e857680246ddf97976680000915061044d565b6127106003548115156103f757fe5b0490506000546001541180156104165750670de0b6b3a7640000600054115b1561043757610430816003540161042b6104d0565b6104b7565b915061044d565b61044a816003540364e8d4a51000610511565b91505b5090565b60015481565b69d3c21bcecceda100000081565b670de0b6b3a764000081565b600481565b64e8d4a5100081565b60025481565b680471fa858b9e08000081565b680246ddf9797668000081565b61271081565b670de0b6b3a764000081565b60005481565b60008183106104c657816104c8565b825b905092915050565b600060018043011180156104e857506104e761052b565b5b6104fb57680471fa858b9e08000061050c565b61271060025481151561050a57fe5b045b905090565b6000818310156105215781610523565b825b905092915050565b600069d3c21bcecceda100000060025410159050905600a165627a7a72305820cfd2b1cb5dc023e2271408e11774b950bd61971cb19508d9c866eabb3611030d0029`

// DeploySystemVars deploys a new Kowala contract, binding an instance of SystemVars to it.
func DeploySystemVars(auth *bind.TransactOpts, backend bind.ContractBackend, initialPrice *big.Int, initialSupply *big.Int) (common.Address, *types.Transaction, *SystemVars, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemVarsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemVarsBin), backend, initialPrice, initialSupply)
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

// AdjustmentFactor is a free data retrieval call binding the contract method 0xcb6741ed.
//
// Solidity: function adjustmentFactor() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) AdjustmentFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "adjustmentFactor")
	return *ret0, err
}

// AdjustmentFactor is a free data retrieval call binding the contract method 0xcb6741ed.
//
// Solidity: function adjustmentFactor() constant returns(uint256)
func (_SystemVars *SystemVarsSession) AdjustmentFactor() (*big.Int, error) {
	return _SystemVars.Contract.AdjustmentFactor(&_SystemVars.CallOpts)
}

// AdjustmentFactor is a free data retrieval call binding the contract method 0xcb6741ed.
//
// Solidity: function adjustmentFactor() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) AdjustmentFactor() (*big.Int, error) {
	return _SystemVars.Contract.AdjustmentFactor(&_SystemVars.CallOpts)
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

// DefaultOracleReward is a free data retrieval call binding the contract method 0xef6711a5.
//
// Solidity: function defaultOracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) DefaultOracleReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "defaultOracleReward")
	return *ret0, err
}

// DefaultOracleReward is a free data retrieval call binding the contract method 0xef6711a5.
//
// Solidity: function defaultOracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsSession) DefaultOracleReward() (*big.Int, error) {
	return _SystemVars.Contract.DefaultOracleReward(&_SystemVars.CallOpts)
}

// DefaultOracleReward is a free data retrieval call binding the contract method 0xef6711a5.
//
// Solidity: function defaultOracleReward() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) DefaultOracleReward() (*big.Int, error) {
	return _SystemVars.Contract.DefaultOracleReward(&_SystemVars.CallOpts)
}

// InitialCap is a free data retrieval call binding the contract method 0xb811215e.
//
// Solidity: function initialCap() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) InitialCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "initialCap")
	return *ret0, err
}

// InitialCap is a free data retrieval call binding the contract method 0xb811215e.
//
// Solidity: function initialCap() constant returns(uint256)
func (_SystemVars *SystemVarsSession) InitialCap() (*big.Int, error) {
	return _SystemVars.Contract.InitialCap(&_SystemVars.CallOpts)
}

// InitialCap is a free data retrieval call binding the contract method 0xb811215e.
//
// Solidity: function initialCap() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) InitialCap() (*big.Int, error) {
	return _SystemVars.Contract.InitialCap(&_SystemVars.CallOpts)
}

// InitialMintedAmount is a free data retrieval call binding the contract method 0xbf053db1.
//
// Solidity: function initialMintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) InitialMintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "initialMintedAmount")
	return *ret0, err
}

// InitialMintedAmount is a free data retrieval call binding the contract method 0xbf053db1.
//
// Solidity: function initialMintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsSession) InitialMintedAmount() (*big.Int, error) {
	return _SystemVars.Contract.InitialMintedAmount(&_SystemVars.CallOpts)
}

// InitialMintedAmount is a free data retrieval call binding the contract method 0xbf053db1.
//
// Solidity: function initialMintedAmount() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) InitialMintedAmount() (*big.Int, error) {
	return _SystemVars.Contract.InitialMintedAmount(&_SystemVars.CallOpts)
}

// LowSupplyMetric is a free data retrieval call binding the contract method 0x74bf78e9.
//
// Solidity: function lowSupplyMetric() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) LowSupplyMetric(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "lowSupplyMetric")
	return *ret0, err
}

// LowSupplyMetric is a free data retrieval call binding the contract method 0x74bf78e9.
//
// Solidity: function lowSupplyMetric() constant returns(uint256)
func (_SystemVars *SystemVarsSession) LowSupplyMetric() (*big.Int, error) {
	return _SystemVars.Contract.LowSupplyMetric(&_SystemVars.CallOpts)
}

// LowSupplyMetric is a free data retrieval call binding the contract method 0x74bf78e9.
//
// Solidity: function lowSupplyMetric() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) LowSupplyMetric() (*big.Int, error) {
	return _SystemVars.Contract.LowSupplyMetric(&_SystemVars.CallOpts)
}

// MaxUnderNormalConditions is a free data retrieval call binding the contract method 0xa47e456f.
//
// Solidity: function maxUnderNormalConditions() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) MaxUnderNormalConditions(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "maxUnderNormalConditions")
	return *ret0, err
}

// MaxUnderNormalConditions is a free data retrieval call binding the contract method 0xa47e456f.
//
// Solidity: function maxUnderNormalConditions() constant returns(uint256)
func (_SystemVars *SystemVarsSession) MaxUnderNormalConditions() (*big.Int, error) {
	return _SystemVars.Contract.MaxUnderNormalConditions(&_SystemVars.CallOpts)
}

// MaxUnderNormalConditions is a free data retrieval call binding the contract method 0xa47e456f.
//
// Solidity: function maxUnderNormalConditions() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) MaxUnderNormalConditions() (*big.Int, error) {
	return _SystemVars.Contract.MaxUnderNormalConditions(&_SystemVars.CallOpts)
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

// OracleDeductionFraction is a free data retrieval call binding the contract method 0xa17aad96.
//
// Solidity: function oracleDeductionFraction() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) OracleDeductionFraction(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "oracleDeductionFraction")
	return *ret0, err
}

// OracleDeductionFraction is a free data retrieval call binding the contract method 0xa17aad96.
//
// Solidity: function oracleDeductionFraction() constant returns(uint256)
func (_SystemVars *SystemVarsSession) OracleDeductionFraction() (*big.Int, error) {
	return _SystemVars.Contract.OracleDeductionFraction(&_SystemVars.CallOpts)
}

// OracleDeductionFraction is a free data retrieval call binding the contract method 0xa17aad96.
//
// Solidity: function oracleDeductionFraction() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) OracleDeductionFraction() (*big.Int, error) {
	return _SystemVars.Contract.OracleDeductionFraction(&_SystemVars.CallOpts)
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

// StabilizedPrice is a free data retrieval call binding the contract method 0x9d26e501.
//
// Solidity: function stabilizedPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCaller) StabilizedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemVars.contract.Call(opts, out, "stabilizedPrice")
	return *ret0, err
}

// StabilizedPrice is a free data retrieval call binding the contract method 0x9d26e501.
//
// Solidity: function stabilizedPrice() constant returns(uint256)
func (_SystemVars *SystemVarsSession) StabilizedPrice() (*big.Int, error) {
	return _SystemVars.Contract.StabilizedPrice(&_SystemVars.CallOpts)
}

// StabilizedPrice is a free data retrieval call binding the contract method 0x9d26e501.
//
// Solidity: function stabilizedPrice() constant returns(uint256)
func (_SystemVars *SystemVarsCallerSession) StabilizedPrice() (*big.Int, error) {
	return _SystemVars.Contract.StabilizedPrice(&_SystemVars.CallOpts)
}
