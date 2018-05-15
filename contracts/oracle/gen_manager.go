// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
)

// OracleManagerABI is the input ABI used to generate the binding from.
const OracleManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxNumOracles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"_isOracle\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerOracle\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"addPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxNumOracles\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OracleManagerBin is the compiled bytecode used for deploying new contracts.
const OracleManagerBin = `606060405260008060146101000a81548160ff021916908315150217905550670de0b6b3a7640000600655341561003557600080fd5b6040516060806111f183398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000821115156100b257600080fd5b8260018190555081600281905550620151808102600381905550505050611113806100de6000396000f3006060604052600436106100e5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062fe7b11146100ea578063035cf142146101135780630a3cb6631461013c578063252f7be914610165578063339d2590146101b65780633f4ba83a146101c05780635c975abb146101d557806369474625146102025780638456cb591461022b5780638da5cb5b1461024057806397584b3e14610295578063a035b1fe146102c2578063aded41ec146102eb578063e9f0ee5614610300578063f2fde38b14610323578063f93a2eb21461035c575b600080fd5b34156100f557600080fd5b6100fd610371565b6040518082815260200191505060405180910390f35b341561011e57600080fd5b610126610377565b6040518082815260200191505060405180910390f35b341561014757600080fd5b61014f61044b565b6040518082815260200191505060405180910390f35b341561017057600080fd5b61019c600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610451565b604051808215151515815260200191505060405180910390f35b6101be6104aa565b005b34156101cb57600080fd5b6101d3610514565b005b34156101e057600080fd5b6101e86105d2565b604051808215151515815260200191505060405180910390f35b341561020d57600080fd5b6102156105e5565b6040518082815260200191505060405180910390f35b341561023657600080fd5b61023e6105eb565b005b341561024b57600080fd5b6102536106ab565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102a057600080fd5b6102a86106d0565b604051808215151515815260200191505060405180910390f35b34156102cd57600080fd5b6102d56106e3565b6040518082815260200191505060405180910390f35b34156102f657600080fd5b6102fe6106e9565b005b341561030b57600080fd5b6103216004808035906020019091905050610844565b005b341561032e57600080fd5b61035a600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061088f565b005b341561036757600080fd5b61036f6109e4565b005b60025481565b6000806103826106d0565b15610391576001549150610447565b6004600060056001600580549050038154811015156103ac57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561043157fe5b9060005260206000209060020201600001540191505b5090565b60035481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b600060149054906101000a900460ff161515156104c657600080fd5b6104cf33610451565b1515156104db57600080fd5b6104e3610377565b34101515156104f157600080fd5b6104f96106d0565b151561050857610507610a1f565b5b6105123334610a6c565b565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561056f57600080fd5b600060149054906101000a900460ff16151561058a57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561064657600080fd5b600060149054906101000a900460ff1615151561066257600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806005805490506002540311905090565b60065481565b60008060008060149054906101000a900460ff1615151561070957600080fd5b6000925060009150600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b80805490508210801561078957506000818381548110151561077457fe5b90600052602060002090600202016001015414155b156107eb57808281548110151561079c57fe5b9060005260206000209060020201600101544210156107ba576107eb565b80828154811015156107c857fe5b906000526020600020906002020160000154830192508180600101925050610756565b6107f53383610d7e565b600083111561083f573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561083e57600080fd5b5b505050565b600060149054906101000a900460ff1615151561086057600080fd5b61086933610451565b151561087457600080fd5b8060008111151561088457600080fd5b816006819055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108ea57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561092657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060149054906101000a900460ff16151515610a0057600080fd5b610a0933610451565b1515610a1457600080fd5b610a1d33610e6b565b565b610a6a6005600160058054905003815481101515610a3957fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610e6b565b565b600080600080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160058054806001018281610ac99190610fd7565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff021916908315150217905550836002018054806001018281610b539190611003565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115610d765760046000600560018603815481101515610bbc57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816002016001836002018054905003815481101515610c3f57fe5b90600052602060002090600202019050806000015485111515610c6157610d76565b600560018403815481101515610c7357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600584815481101515610cae57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600560018503815481101515610d0a57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050610b9d565b505050505050565b600080600080841415610d9057610e64565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b8260020180549050811015610e52578260020181815481101515610df957fe5b90600052602060002090600202018360020183815481101515610e1857fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050610dd9565b818360020181610e629190611035565b505b5050505050565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160058054905003811015610f6a57600560018201815481101515610ed957fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600582815481101515610f1457fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050610eb7565b6005805480919060019003610f7f9190611067565b5060008260010160006101000a81548160ff0219169083151502179055506003544201826002016001846002018054905003815481101515610fbd57fe5b906000526020600020906002020160010181905550505050565b815481835581811511610ffe57818360005260206000209182019101610ffd9190611093565b5b505050565b8154818355818115116110305760020281600202836000526020600020918201910161102f91906110b8565b5b505050565b8154818355818115116110625760020281600202836000526020600020918201910161106191906110b8565b5b505050565b81548183558181151161108e5781836000526020600020918201910161108d9190611093565b5b505050565b6110b591905b808211156110b1576000816000905550600101611099565b5090565b90565b6110e491905b808211156110e0576000808201600090556001820160009055506002016110be565b5090565b905600a165627a7a72305820c9d96a77bf272417cb0556d37596ca8317bc3514acc94016e29d7cef9f4d55ac0029`

// DeployOracleManager deploys a new Ethereum contract, binding an instance of OracleManager to it.
func DeployOracleManager(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxNumOracles *big.Int, _freezePeriod *big.Int) (common.Address, *types.Transaction, *OracleManager, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleManagerBin), backend, _baseDeposit, _maxNumOracles, _freezePeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleManager{OracleManagerCaller: OracleManagerCaller{contract: contract}, OracleManagerTransactor: OracleManagerTransactor{contract: contract}}, nil
}

// OracleManager is an auto generated Go binding around an Ethereum contract.
type OracleManager struct {
	OracleManagerCaller     // Read-only binding to the contract
	OracleManagerTransactor // Write-only binding to the contract
}

// OracleManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleManagerSession struct {
	Contract     *OracleManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleManagerCallerSession struct {
	Contract *OracleManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OracleManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleManagerTransactorSession struct {
	Contract     *OracleManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OracleManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleManagerRaw struct {
	Contract *OracleManager // Generic contract binding to access the raw methods on
}

// OracleManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleManagerCallerRaw struct {
	Contract *OracleManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OracleManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleManagerTransactorRaw struct {
	Contract *OracleManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleManager creates a new instance of OracleManager, bound to a specific deployed contract.
func NewOracleManager(address common.Address, backend bind.ContractBackend) (*OracleManager, error) {
	contract, err := bindOracleManager(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleManager{OracleManagerCaller: OracleManagerCaller{contract: contract}, OracleManagerTransactor: OracleManagerTransactor{contract: contract}}, nil
}

// NewOracleManagerCaller creates a new read-only instance of OracleManager, bound to a specific deployed contract.
func NewOracleManagerCaller(address common.Address, caller bind.ContractCaller) (*OracleManagerCaller, error) {
	contract, err := bindOracleManager(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OracleManagerCaller{contract: contract}, nil
}

// NewOracleManagerTransactor creates a new write-only instance of OracleManager, bound to a specific deployed contract.
func NewOracleManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleManagerTransactor, error) {
	contract, err := bindOracleManager(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OracleManagerTransactor{contract: contract}, nil
}

// bindOracleManager binds a generic wrapper to an already deployed contract.
func bindOracleManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleManager *OracleManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleManager.Contract.OracleManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleManager *OracleManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManager.Contract.OracleManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleManager *OracleManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleManager.Contract.OracleManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleManager *OracleManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleManager *OracleManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleManager *OracleManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleManager.Contract.contract.Transact(opts, method, params...)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleManager *OracleManagerCaller) _hasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleManager *OracleManagerSession) _hasAvailability() (bool, error) {
	return _OracleManager.Contract._hasAvailability(&_OracleManager.CallOpts)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleManager *OracleManagerCallerSession) _hasAvailability() (bool, error) {
	return _OracleManager.Contract._hasAvailability(&_OracleManager.CallOpts)
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleManager *OracleManagerCaller) _isOracle(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "_isOracle", identity)
	return *ret0, err
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleManager *OracleManagerSession) _isOracle(identity common.Address) (bool, error) {
	return _OracleManager.Contract._isOracle(&_OracleManager.CallOpts, identity)
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleManager *OracleManagerCallerSession) _isOracle(identity common.Address) (bool, error) {
	return _OracleManager.Contract._isOracle(&_OracleManager.CallOpts, identity)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleManager *OracleManagerCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleManager *OracleManagerSession) BaseDeposit() (*big.Int, error) {
	return _OracleManager.Contract.BaseDeposit(&_OracleManager.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleManager *OracleManagerCallerSession) BaseDeposit() (*big.Int, error) {
	return _OracleManager.Contract.BaseDeposit(&_OracleManager.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleManager *OracleManagerCaller) FreezePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "freezePeriod")
	return *ret0, err
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleManager *OracleManagerSession) FreezePeriod() (*big.Int, error) {
	return _OracleManager.Contract.FreezePeriod(&_OracleManager.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleManager *OracleManagerCallerSession) FreezePeriod() (*big.Int, error) {
	return _OracleManager.Contract.FreezePeriod(&_OracleManager.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleManager *OracleManagerCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleManager *OracleManagerSession) GetMinimumDeposit() (*big.Int, error) {
	return _OracleManager.Contract.GetMinimumDeposit(&_OracleManager.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleManager *OracleManagerCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _OracleManager.Contract.GetMinimumDeposit(&_OracleManager.CallOpts)
}

// MaxNumOracles is a free data retrieval call binding the contract method 0x00fe7b11.
//
// Solidity: function maxNumOracles() constant returns(uint256)
func (_OracleManager *OracleManagerCaller) MaxNumOracles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "maxNumOracles")
	return *ret0, err
}

// MaxNumOracles is a free data retrieval call binding the contract method 0x00fe7b11.
//
// Solidity: function maxNumOracles() constant returns(uint256)
func (_OracleManager *OracleManagerSession) MaxNumOracles() (*big.Int, error) {
	return _OracleManager.Contract.MaxNumOracles(&_OracleManager.CallOpts)
}

// MaxNumOracles is a free data retrieval call binding the contract method 0x00fe7b11.
//
// Solidity: function maxNumOracles() constant returns(uint256)
func (_OracleManager *OracleManagerCallerSession) MaxNumOracles() (*big.Int, error) {
	return _OracleManager.Contract.MaxNumOracles(&_OracleManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_OracleManager *OracleManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_OracleManager *OracleManagerSession) Owner() (common.Address, error) {
	return _OracleManager.Contract.Owner(&_OracleManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_OracleManager *OracleManagerCallerSession) Owner() (common.Address, error) {
	return _OracleManager.Contract.Owner(&_OracleManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_OracleManager *OracleManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_OracleManager *OracleManagerSession) Paused() (bool, error) {
	return _OracleManager.Contract.Paused(&_OracleManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_OracleManager *OracleManagerCallerSession) Paused() (bool, error) {
	return _OracleManager.Contract.Paused(&_OracleManager.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleManager *OracleManagerCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleManager *OracleManagerSession) Price() (*big.Int, error) {
	return _OracleManager.Contract.Price(&_OracleManager.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleManager *OracleManagerCallerSession) Price() (*big.Int, error) {
	return _OracleManager.Contract.Price(&_OracleManager.CallOpts)
}

// AddPrice is a paid mutator transaction binding the contract method 0xe9f0ee56.
//
// Solidity: function addPrice(_price uint256) returns()
func (_OracleManager *OracleManagerTransactor) AddPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "addPrice", _price)
}

// AddPrice is a paid mutator transaction binding the contract method 0xe9f0ee56.
//
// Solidity: function addPrice(_price uint256) returns()
func (_OracleManager *OracleManagerSession) AddPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleManager.Contract.AddPrice(&_OracleManager.TransactOpts, _price)
}

// AddPrice is a paid mutator transaction binding the contract method 0xe9f0ee56.
//
// Solidity: function addPrice(_price uint256) returns()
func (_OracleManager *OracleManagerTransactorSession) AddPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleManager.Contract.AddPrice(&_OracleManager.TransactOpts, _price)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleManager *OracleManagerTransactor) DeregisterOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "deregisterOracle")
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleManager *OracleManagerSession) DeregisterOracle() (*types.Transaction, error) {
	return _OracleManager.Contract.DeregisterOracle(&_OracleManager.TransactOpts)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleManager *OracleManagerTransactorSession) DeregisterOracle() (*types.Transaction, error) {
	return _OracleManager.Contract.DeregisterOracle(&_OracleManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleManager *OracleManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleManager *OracleManagerSession) Pause() (*types.Transaction, error) {
	return _OracleManager.Contract.Pause(&_OracleManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleManager *OracleManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _OracleManager.Contract.Pause(&_OracleManager.TransactOpts)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleManager *OracleManagerTransactor) RegisterOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "registerOracle")
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleManager *OracleManagerSession) RegisterOracle() (*types.Transaction, error) {
	return _OracleManager.Contract.RegisterOracle(&_OracleManager.TransactOpts)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleManager *OracleManagerTransactorSession) RegisterOracle() (*types.Transaction, error) {
	return _OracleManager.Contract.RegisterOracle(&_OracleManager.TransactOpts)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_OracleManager *OracleManagerTransactor) ReleaseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "releaseDeposits")
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_OracleManager *OracleManagerSession) ReleaseDeposits() (*types.Transaction, error) {
	return _OracleManager.Contract.ReleaseDeposits(&_OracleManager.TransactOpts)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_OracleManager *OracleManagerTransactorSession) ReleaseDeposits() (*types.Transaction, error) {
	return _OracleManager.Contract.ReleaseDeposits(&_OracleManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_OracleManager *OracleManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_OracleManager *OracleManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleManager.Contract.TransferOwnership(&_OracleManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_OracleManager *OracleManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleManager.Contract.TransferOwnership(&_OracleManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleManager *OracleManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleManager *OracleManagerSession) Unpause() (*types.Transaction, error) {
	return _OracleManager.Contract.Unpause(&_OracleManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleManager *OracleManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _OracleManager.Contract.Unpause(&_OracleManager.TransactOpts)
}
