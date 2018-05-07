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
const OracleManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"_isOracle\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerOracle\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"submitPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxOracles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxOracles\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OracleManagerBin is the compiled bytecode used for deploying new contracts.
const OracleManagerBin = `606060405260008060146101000a81548160ff021916908315150217905550670de0b6b3a7640000600655341561003557600080fd5b6040516060806111c383398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182101515156100b357600080fd5b670de0b6b3a76400008302600181905550816002819055506201518081026003819055505050506110da806100e96000396000f3006060604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf142146100e05780630a3cb66314610109578063252f7be914610132578063339d2590146101835780633f4ba83a1461018d5780635c975abb146101a257806369474625146101cf5780638456cb59146101f85780638da5cb5b1461020d57806397584b3e14610262578063986fcbe91461028f578063aded41ec146102b2578063c0d2c49d146102c7578063f2fde38b146102f0578063f93a2eb214610329575b600080fd5b34156100eb57600080fd5b6100f361033e565b6040518082815260200191505060405180910390f35b341561011457600080fd5b61011c610412565b6040518082815260200191505060405180910390f35b341561013d57600080fd5b610169600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610418565b604051808215151515815260200191505060405180910390f35b61018b610471565b005b341561019857600080fd5b6101a06104db565b005b34156101ad57600080fd5b6101b5610599565b604051808215151515815260200191505060405180910390f35b34156101da57600080fd5b6101e26105ac565b6040518082815260200191505060405180910390f35b341561020357600080fd5b61020b6105b2565b005b341561021857600080fd5b610220610672565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561026d57600080fd5b610275610697565b604051808215151515815260200191505060405180910390f35b341561029a57600080fd5b6102b060048080359060200190919050506106aa565b005b34156102bd57600080fd5b6102c56106f5565b005b34156102d257600080fd5b6102da610850565b6040518082815260200191505060405180910390f35b34156102fb57600080fd5b610327600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610856565b005b341561033457600080fd5b61033c6109ab565b005b600080610349610697565b1561035857600154915061040e565b60046000600560016005805490500381548110151561037357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160018360020180549050038154811015156103f857fe5b9060005260206000209060020201600001540191505b5090565b60035481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b600060149054906101000a900460ff1615151561048d57600080fd5b61049633610418565b1515156104a257600080fd5b6104aa61033e565b34101515156104b857600080fd5b6104c0610697565b15156104cf576104ce6109e6565b5b6104d93334610a33565b565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561053657600080fd5b600060149054906101000a900460ff16151561055157600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561060d57600080fd5b600060149054906101000a900460ff1615151561062957600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806005805490506002540311905090565b600060149054906101000a900460ff161515156106c657600080fd5b6106cf33610418565b15156106da57600080fd5b60006006541115156106eb57600080fd5b8060068190555050565b60008060008060149054906101000a900460ff1615151561071557600080fd5b6000925060009150600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b80805490508210801561079557506000818381548110151561078057fe5b90600052602060002090600202016001015414155b156107f75780828154811015156107a857fe5b9060005260206000209060020201600101544210156107c6576107f7565b80828154811015156107d457fe5b906000526020600020906002020160000154830192508180600101925050610762565b6108013383610d45565b600083111561084b573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561084a57600080fd5b5b505050565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108b157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156108ed57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060149054906101000a900460ff161515156109c757600080fd5b6109d033610418565b15156109db57600080fd5b6109e433610e32565b565b610a316005600160058054905003815481101515610a0057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610e32565b565b600080600080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160058054806001018281610a909190610f9e565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff021916908315150217905550836002018054806001018281610b1a9190610fca565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115610d3d5760046000600560018603815481101515610b8357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816002016001836002018054905003815481101515610c0657fe5b90600052602060002090600202019050806000015485111515610c2857610d3d565b600560018403815481101515610c3a57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600584815481101515610c7557fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600560018503815481101515610cd157fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050610b64565b505050505050565b600080600080841415610d5757610e2b565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b8260020180549050811015610e19578260020181815481101515610dc057fe5b90600052602060002090600202018360020183815481101515610ddf57fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050610da0565b818360020181610e299190610ffc565b505b5050505050565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160058054905003811015610f3157600560018201815481101515610ea057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600582815481101515610edb57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050610e7e565b6005805480919060019003610f46919061102e565b5060008260010160006101000a81548160ff0219169083151502179055506003544201826002016001846002018054905003815481101515610f8457fe5b906000526020600020906002020160010181905550505050565b815481835581811511610fc557818360005260206000209182019101610fc4919061105a565b5b505050565b815481835581811511610ff757600202816002028360005260206000209182019101610ff6919061107f565b5b505050565b81548183558181151161102957600202816002028360005260206000209182019101611028919061107f565b5b505050565b81548183558181151161105557818360005260206000209182019101611054919061105a565b5b505050565b61107c91905b80821115611078576000816000905550600101611060565b5090565b90565b6110ab91905b808211156110a757600080820160009055600182016000905550600201611085565b5090565b905600a165627a7a7230582005845cf46ea05d2c5e82fd220c6542dfef075714832dd1310ea8d608263f04400029`

// DeployOracleManager deploys a new Ethereum contract, binding an instance of OracleManager to it.
func DeployOracleManager(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxOracles *big.Int, _freezePeriod *big.Int) (common.Address, *types.Transaction, *OracleManager, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleManagerBin), backend, _baseDeposit, _maxOracles, _freezePeriod)
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

// MaxOracles is a free data retrieval call binding the contract method 0xc0d2c49d.
//
// Solidity: function maxOracles() constant returns(uint256)
func (_OracleManager *OracleManagerCaller) MaxOracles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleManager.contract.Call(opts, out, "maxOracles")
	return *ret0, err
}

// MaxOracles is a free data retrieval call binding the contract method 0xc0d2c49d.
//
// Solidity: function maxOracles() constant returns(uint256)
func (_OracleManager *OracleManagerSession) MaxOracles() (*big.Int, error) {
	return _OracleManager.Contract.MaxOracles(&_OracleManager.CallOpts)
}

// MaxOracles is a free data retrieval call binding the contract method 0xc0d2c49d.
//
// Solidity: function maxOracles() constant returns(uint256)
func (_OracleManager *OracleManagerCallerSession) MaxOracles() (*big.Int, error) {
	return _OracleManager.Contract.MaxOracles(&_OracleManager.CallOpts)
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

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(_price uint256) returns()
func (_OracleManager *OracleManagerTransactor) SubmitPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _OracleManager.contract.Transact(opts, "submitPrice", _price)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(_price uint256) returns()
func (_OracleManager *OracleManagerSession) SubmitPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleManager.Contract.SubmitPrice(&_OracleManager.TransactOpts, _price)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(_price uint256) returns()
func (_OracleManager *OracleManagerTransactorSession) SubmitPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleManager.Contract.SubmitPrice(&_OracleManager.TransactOpts, _price)
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
