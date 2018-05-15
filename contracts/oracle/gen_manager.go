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
const OracleManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxNumOracles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"_isOracle\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerOracle\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"submitPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxNumOracles\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OracleManagerBin is the compiled bytecode used for deploying new contracts.
const OracleManagerBin = `606060405260008060146101000a81548160ff021916908315150217905550670de0b6b3a7640000600655341561003557600080fd5b6040516060806111c183398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000821115156100b257600080fd5b670de0b6b3a76400008302600181905550816002819055506201518081026003819055505050506110d9806100e86000396000f3006060604052600436106100da576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062fe7b11146100df578063035cf142146101085780630a3cb66314610131578063252f7be91461015a578063339d2590146101ab5780633f4ba83a146101b55780635c975abb146101ca57806369474625146101f75780638456cb59146102205780638da5cb5b1461023557806397584b3e1461028a578063986fcbe9146102b7578063aded41ec146102da578063f2fde38b146102ef578063f93a2eb214610328575b600080fd5b34156100ea57600080fd5b6100f261033d565b6040518082815260200191505060405180910390f35b341561011357600080fd5b61011b610343565b6040518082815260200191505060405180910390f35b341561013c57600080fd5b610144610417565b6040518082815260200191505060405180910390f35b341561016557600080fd5b610191600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061041d565b604051808215151515815260200191505060405180910390f35b6101b3610476565b005b34156101c057600080fd5b6101c86104e0565b005b34156101d557600080fd5b6101dd61059e565b604051808215151515815260200191505060405180910390f35b341561020257600080fd5b61020a6105b1565b6040518082815260200191505060405180910390f35b341561022b57600080fd5b6102336105b7565b005b341561024057600080fd5b610248610677565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561029557600080fd5b61029d61069c565b604051808215151515815260200191505060405180910390f35b34156102c257600080fd5b6102d860048080359060200190919050506106af565b005b34156102e557600080fd5b6102ed6106fa565b005b34156102fa57600080fd5b610326600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610855565b005b341561033357600080fd5b61033b6109aa565b005b60025481565b60008061034e61069c565b1561035d576001549150610413565b60046000600560016005805490500381548110151561037857fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160018360020180549050038154811015156103fd57fe5b9060005260206000209060020201600001540191505b5090565b60035481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b600060149054906101000a900460ff1615151561049257600080fd5b61049b3361041d565b1515156104a757600080fd5b6104af610343565b34101515156104bd57600080fd5b6104c561069c565b15156104d4576104d36109e5565b5b6104de3334610a32565b565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561053b57600080fd5b600060149054906101000a900460ff16151561055657600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561061257600080fd5b600060149054906101000a900460ff1615151561062e57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806005805490506002540311905090565b600060149054906101000a900460ff161515156106cb57600080fd5b6106d43361041d565b15156106df57600080fd5b60006006541115156106f057600080fd5b8060068190555050565b60008060008060149054906101000a900460ff1615151561071a57600080fd5b6000925060009150600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b80805490508210801561079a57506000818381548110151561078557fe5b90600052602060002090600202016001015414155b156107fc5780828154811015156107ad57fe5b9060005260206000209060020201600101544210156107cb576107fc565b80828154811015156107d957fe5b906000526020600020906002020160000154830192508180600101925050610767565b6108063383610d44565b6000831115610850573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561084f57600080fd5b5b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108b057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156108ec57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060149054906101000a900460ff161515156109c657600080fd5b6109cf3361041d565b15156109da57600080fd5b6109e333610e31565b565b610a3060056001600580549050038154811015156109ff57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610e31565b565b600080600080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160058054806001018281610a8f9190610f9d565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff021916908315150217905550836002018054806001018281610b199190610fc9565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115610d3c5760046000600560018603815481101515610b8257fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816002016001836002018054905003815481101515610c0557fe5b90600052602060002090600202019050806000015485111515610c2757610d3c565b600560018403815481101515610c3957fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600584815481101515610c7457fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600560018503815481101515610cd057fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050610b63565b505050505050565b600080600080841415610d5657610e2a565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b8260020180549050811015610e18578260020181815481101515610dbf57fe5b90600052602060002090600202018360020183815481101515610dde57fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050610d9f565b818360020181610e289190610ffb565b505b5050505050565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160058054905003811015610f3057600560018201815481101515610e9f57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600582815481101515610eda57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050610e7d565b6005805480919060019003610f45919061102d565b5060008260010160006101000a81548160ff0219169083151502179055506003544201826002016001846002018054905003815481101515610f8357fe5b906000526020600020906002020160010181905550505050565b815481835581811511610fc457818360005260206000209182019101610fc39190611059565b5b505050565b815481835581811511610ff657600202816002028360005260206000209182019101610ff5919061107e565b5b505050565b81548183558181151161102857600202816002028360005260206000209182019101611027919061107e565b5b505050565b815481835581811511611054578183600052602060002091820191016110539190611059565b5b505050565b61107b91905b8082111561107757600081600090555060010161105f565b5090565b90565b6110aa91905b808211156110a657600080820160009055600182016000905550600201611084565b5090565b905600a165627a7a7230582083a88a403113fcb0738dcb6108759dbebb1e194d10c9333ebb3afe0252dc6a0e0029`

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
