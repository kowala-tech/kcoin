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

// OracleMgrABI is the input ABI used to generate the binding from.
const OracleMgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxNumOracles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getOracleAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"_isOracle\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerOracle\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"availableAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracleCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"syncFrequency\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"addPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialPrice\",\"type\":\"uint256\"},{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxNumOracles\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"},{\"name\":\"_syncFrequency\",\"type\":\"uint256\"},{\"name\":\"_updatePeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OracleMgrBin is the compiled bytecode used for deploying new contracts.
const OracleMgrBin = `606060405260008060146101000a81548160ff021916908315150217905550341561002957600080fd5b60405160c08061152a83398101604052808051906020019091908051906020019091908051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000861115156100c157600080fd5b6000841115156100d057600080fd5b600082101515156100e057600080fd5b600081101515156100f057600080fd5b856005819055508460028190555083600381905550620151808302600481905550816001819055505050505050506113fd8061012d6000396000f30060606040526004361061011c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062fe7b1114610121578063035cf1421461014a57806309fe9d39146101735780630a3cb663146101dd578063252f7be914610206578063339d2590146102575780633ed0a373146102615780633f4ba83a1461029f5780633f4e4251146102b45780635c975abb146102dd578063694746251461030a5780638456cb59146103335780638da5cb5b146103485780639363a1411461039d57806397584b3e146103c6578063a035b1fe146103f3578063aded41ec1461041c578063cdee7e0714610431578063e9f0ee561461045a578063f2fde38b1461047d578063f93a2eb2146104b6575b600080fd5b341561012c57600080fd5b6101346104cb565b6040518082815260200191505060405180910390f35b341561015557600080fd5b61015d6104d1565b6040518082815260200191505060405180910390f35b341561017e57600080fd5b61019460048080359060200190919050506105a5565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34156101e857600080fd5b6101f061065d565b6040518082815260200191505060405180910390f35b341561021157600080fd5b61023d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610663565b604051808215151515815260200191505060405180910390f35b61025f6106bc565b005b341561026c57600080fd5b6102826004808035906020019091905050610726565b604051808381526020018281526020019250505060405180910390f35b34156102aa57600080fd5b6102b261079e565b005b34156102bf57600080fd5b6102c761085c565b6040518082815260200191505060405180910390f35b34156102e857600080fd5b6102f0610869565b604051808215151515815260200191505060405180910390f35b341561031557600080fd5b61031d61087c565b6040518082815260200191505060405180910390f35b341561033e57600080fd5b610346610882565b005b341561035357600080fd5b61035b610942565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103a857600080fd5b6103b0610967565b6040518082815260200191505060405180910390f35b34156103d157600080fd5b6103d96109b4565b604051808215151515815260200191505060405180910390f35b34156103fe57600080fd5b6104066109c7565b6040518082815260200191505060405180910390f35b341561042757600080fd5b61042f6109cd565b005b341561043c57600080fd5b610444610b28565b6040518082815260200191505060405180910390f35b341561046557600080fd5b61047b6004808035906020019091905050610b2e565b005b341561048857600080fd5b6104b4600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b79565b005b34156104c157600080fd5b6104c9610cce565b005b60035481565b6000806104dc6109b4565b156104eb5760025491506105a1565b60066000600760016007805490500381548110151561050657fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561058b57fe5b9060005260206000209060020201600001540191505b5090565b60008060006007848154811015156105b957fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600201600182600201805490500381548110151561064357fe5b906000526020600020906002020160000154915050915091565b60045481565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b600060149054906101000a900460ff161515156106d857600080fd5b6106e133610663565b1515156106ed57600080fd5b6106f56104d1565b341015151561070357600080fd5b61070b6109b4565b151561071a57610719610d09565b5b6107243334610d56565b565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018481548110151561077a57fe5b90600052602060002090600202019050806000015481600101549250925050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107f957600080fd5b600060149054906101000a900460ff16151561081457600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6000600780549050905090565b600060149054906101000a900460ff1681565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108dd57600080fd5b600060149054906101000a900460ff161515156108f957600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506003540311905090565b60055481565b60008060008060149054906101000a900460ff161515156109ed57600080fd5b6000925060009150600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b808054905082108015610a6d575060008183815481101515610a5857fe5b90600052602060002090600202016001015414155b15610acf578082815481101515610a8057fe5b906000526020600020906002020160010154421015610a9e57610acf565b8082815481101515610aac57fe5b906000526020600020906002020160000154830192508180600101925050610a3a565b610ad93383611068565b6000831115610b23573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f193505050501515610b2257600080fd5b5b505050565b60015481565b600060149054906101000a900460ff16151515610b4a57600080fd5b610b5333610663565b1515610b5e57600080fd5b80600081111515610b6e57600080fd5b816005819055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bd457600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610c1057600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060149054906101000a900460ff16151515610cea57600080fd5b610cf333610663565b1515610cfe57600080fd5b610d0733611155565b565b610d546007600160078054905003815481101515610d2357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611155565b565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078054806001018281610db391906112c1565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff021916908315150217905550836002018054806001018281610e3d91906112ed565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b60008311156110605760066000600760018603815481101515610ea657fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816002016001836002018054905003815481101515610f2957fe5b90600052602060002090600202019050806000015485111515610f4b57611060565b600760018403815481101515610f5d57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600784815481101515610f9857fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600760018503815481101515610ff457fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050610e87565b505050505050565b60008060008084141561107a5761114e565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b826002018054905081101561113c5782600201818154811015156110e357fe5b9060005260206000209060020201836002018381548110151561110257fe5b90600052602060002090600202016000820154816000015560018201548160010155905050818060010192505080806001019150506110c3565b81836002018161114c919061131f565b505b5050505050565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160078054905003811015611254576007600182018154811015156111c357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007828154811015156111fe57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806001019150506111a1565b60078054809190600190036112699190611351565b5060008260010160006101000a81548160ff02191690831515021790555060045442018260020160018460020180549050038154811015156112a757fe5b906000526020600020906002020160010181905550505050565b8154818355818115116112e8578183600052602060002091820191016112e7919061137d565b5b505050565b81548183558181151161131a5760020281600202836000526020600020918201910161131991906113a2565b5b505050565b81548183558181151161134c5760020281600202836000526020600020918201910161134b91906113a2565b5b505050565b81548183558181151161137857818360005260206000209182019101611377919061137d565b5b505050565b61139f91905b8082111561139b576000816000905550600101611383565b5090565b90565b6113ce91905b808211156113ca576000808201600090556001820160009055506002016113a8565b5090565b905600a165627a7a72305820d95b776a4e9eb15f553ef66351775699796cccbd3876fdf900911890c86d17180029`

// DeployOracleMgr deploys a new Ethereum contract, binding an instance of OracleMgr to it.
func DeployOracleMgr(auth *bind.TransactOpts, backend bind.ContractBackend, _initialPrice *big.Int, _baseDeposit *big.Int, _maxNumOracles *big.Int, _freezePeriod *big.Int, _syncFrequency *big.Int, _updatePeriod *big.Int) (common.Address, *types.Transaction, *OracleMgr, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleMgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleMgrBin), backend, _initialPrice, _baseDeposit, _maxNumOracles, _freezePeriod, _syncFrequency, _updatePeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleMgr{OracleMgrCaller: OracleMgrCaller{contract: contract}, OracleMgrTransactor: OracleMgrTransactor{contract: contract}}, nil
}

// OracleMgr is an auto generated Go binding around an Ethereum contract.
type OracleMgr struct {
	OracleMgrCaller     // Read-only binding to the contract
	OracleMgrTransactor // Write-only binding to the contract
}

// OracleMgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleMgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleMgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleMgrSession struct {
	Contract     *OracleMgr        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleMgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleMgrCallerSession struct {
	Contract *OracleMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OracleMgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleMgrTransactorSession struct {
	Contract     *OracleMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OracleMgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleMgrRaw struct {
	Contract *OracleMgr // Generic contract binding to access the raw methods on
}

// OracleMgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleMgrCallerRaw struct {
	Contract *OracleMgrCaller // Generic read-only contract binding to access the raw methods on
}

// OracleMgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleMgrTransactorRaw struct {
	Contract *OracleMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleMgr creates a new instance of OracleMgr, bound to a specific deployed contract.
func NewOracleMgr(address common.Address, backend bind.ContractBackend) (*OracleMgr, error) {
	contract, err := bindOracleMgr(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleMgr{OracleMgrCaller: OracleMgrCaller{contract: contract}, OracleMgrTransactor: OracleMgrTransactor{contract: contract}}, nil
}

// NewOracleMgrCaller creates a new read-only instance of OracleMgr, bound to a specific deployed contract.
func NewOracleMgrCaller(address common.Address, caller bind.ContractCaller) (*OracleMgrCaller, error) {
	contract, err := bindOracleMgr(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OracleMgrCaller{contract: contract}, nil
}

// NewOracleMgrTransactor creates a new write-only instance of OracleMgr, bound to a specific deployed contract.
func NewOracleMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleMgrTransactor, error) {
	contract, err := bindOracleMgr(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OracleMgrTransactor{contract: contract}, nil
}

// bindOracleMgr binds a generic wrapper to an already deployed contract.
func bindOracleMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleMgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleMgr *OracleMgrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleMgr.Contract.OracleMgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleMgr *OracleMgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.Contract.OracleMgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleMgr *OracleMgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleMgr.Contract.OracleMgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleMgr *OracleMgrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleMgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleMgr *OracleMgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleMgr *OracleMgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleMgr.Contract.contract.Transact(opts, method, params...)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleMgr *OracleMgrCaller) _hasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleMgr *OracleMgrSession) _hasAvailability() (bool, error) {
	return _OracleMgr.Contract._hasAvailability(&_OracleMgr.CallOpts)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_OracleMgr *OracleMgrCallerSession) _hasAvailability() (bool, error) {
	return _OracleMgr.Contract._hasAvailability(&_OracleMgr.CallOpts)
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleMgr *OracleMgrCaller) _isOracle(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "_isOracle", identity)
	return *ret0, err
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleMgr *OracleMgrSession) _isOracle(identity common.Address) (bool, error) {
	return _OracleMgr.Contract._isOracle(&_OracleMgr.CallOpts, identity)
}

// _isOracle is a free data retrieval call binding the contract method 0x252f7be9.
//
// Solidity: function _isOracle(identity address) constant returns(isIndeed bool)
func (_OracleMgr *OracleMgrCallerSession) _isOracle(identity common.Address) (bool, error) {
	return _OracleMgr.Contract._isOracle(&_OracleMgr.CallOpts, identity)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleMgr *OracleMgrCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleMgr *OracleMgrSession) BaseDeposit() (*big.Int, error) {
	return _OracleMgr.Contract.BaseDeposit(&_OracleMgr.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_OracleMgr *OracleMgrCallerSession) BaseDeposit() (*big.Int, error) {
	return _OracleMgr.Contract.BaseDeposit(&_OracleMgr.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleMgr *OracleMgrCaller) FreezePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "freezePeriod")
	return *ret0, err
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleMgr *OracleMgrSession) FreezePeriod() (*big.Int, error) {
	return _OracleMgr.Contract.FreezePeriod(&_OracleMgr.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_OracleMgr *OracleMgrCallerSession) FreezePeriod() (*big.Int, error) {
	return _OracleMgr.Contract.FreezePeriod(&_OracleMgr.CallOpts)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_OracleMgr *OracleMgrCaller) GetDepositAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	ret := new(struct {
		Amount      *big.Int
		AvailableAt *big.Int
	})
	out := ret
	err := _OracleMgr.contract.Call(opts, out, "getDepositAtIndex", index)
	return *ret, err
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_OracleMgr *OracleMgrSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	return _OracleMgr.Contract.GetDepositAtIndex(&_OracleMgr.CallOpts, index)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_OracleMgr *OracleMgrCallerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	return _OracleMgr.Contract.GetDepositAtIndex(&_OracleMgr.CallOpts, index)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrCaller) GetDepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "getDepositCount")
	return *ret0, err
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrSession) GetDepositCount() (*big.Int, error) {
	return _OracleMgr.Contract.GetDepositCount(&_OracleMgr.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrCallerSession) GetDepositCount() (*big.Int, error) {
	return _OracleMgr.Contract.GetDepositCount(&_OracleMgr.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleMgr *OracleMgrCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleMgr *OracleMgrSession) GetMinimumDeposit() (*big.Int, error) {
	return _OracleMgr.Contract.GetMinimumDeposit(&_OracleMgr.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_OracleMgr *OracleMgrCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _OracleMgr.Contract.GetMinimumDeposit(&_OracleMgr.CallOpts)
}

// GetOracleAtIndex is a free data retrieval call binding the contract method 0x09fe9d39.
//
// Solidity: function getOracleAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_OracleMgr *OracleMgrCaller) GetOracleAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _OracleMgr.contract.Call(opts, out, "getOracleAtIndex", index)
	return *ret, err
}

// GetOracleAtIndex is a free data retrieval call binding the contract method 0x09fe9d39.
//
// Solidity: function getOracleAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_OracleMgr *OracleMgrSession) GetOracleAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _OracleMgr.Contract.GetOracleAtIndex(&_OracleMgr.CallOpts, index)
}

// GetOracleAtIndex is a free data retrieval call binding the contract method 0x09fe9d39.
//
// Solidity: function getOracleAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_OracleMgr *OracleMgrCallerSession) GetOracleAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _OracleMgr.Contract.GetOracleAtIndex(&_OracleMgr.CallOpts, index)
}

// GetOracleCount is a free data retrieval call binding the contract method 0x3f4e4251.
//
// Solidity: function getOracleCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrCaller) GetOracleCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "getOracleCount")
	return *ret0, err
}

// GetOracleCount is a free data retrieval call binding the contract method 0x3f4e4251.
//
// Solidity: function getOracleCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrSession) GetOracleCount() (*big.Int, error) {
	return _OracleMgr.Contract.GetOracleCount(&_OracleMgr.CallOpts)
}

// GetOracleCount is a free data retrieval call binding the contract method 0x3f4e4251.
//
// Solidity: function getOracleCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrCallerSession) GetOracleCount() (*big.Int, error) {
	return _OracleMgr.Contract.GetOracleCount(&_OracleMgr.CallOpts)
}

// MaxNumOracles is a free data retrieval call binding the contract method 0x00fe7b11.
//
// Solidity: function maxNumOracles() constant returns(uint256)
func (_OracleMgr *OracleMgrCaller) MaxNumOracles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "maxNumOracles")
	return *ret0, err
}

// MaxNumOracles is a free data retrieval call binding the contract method 0x00fe7b11.
//
// Solidity: function maxNumOracles() constant returns(uint256)
func (_OracleMgr *OracleMgrSession) MaxNumOracles() (*big.Int, error) {
	return _OracleMgr.Contract.MaxNumOracles(&_OracleMgr.CallOpts)
}

// MaxNumOracles is a free data retrieval call binding the contract method 0x00fe7b11.
//
// Solidity: function maxNumOracles() constant returns(uint256)
func (_OracleMgr *OracleMgrCallerSession) MaxNumOracles() (*big.Int, error) {
	return _OracleMgr.Contract.MaxNumOracles(&_OracleMgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_OracleMgr *OracleMgrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_OracleMgr *OracleMgrSession) Owner() (common.Address, error) {
	return _OracleMgr.Contract.Owner(&_OracleMgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_OracleMgr *OracleMgrCallerSession) Owner() (common.Address, error) {
	return _OracleMgr.Contract.Owner(&_OracleMgr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_OracleMgr *OracleMgrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_OracleMgr *OracleMgrSession) Paused() (bool, error) {
	return _OracleMgr.Contract.Paused(&_OracleMgr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_OracleMgr *OracleMgrCallerSession) Paused() (bool, error) {
	return _OracleMgr.Contract.Paused(&_OracleMgr.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleMgr *OracleMgrCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleMgr *OracleMgrSession) Price() (*big.Int, error) {
	return _OracleMgr.Contract.Price(&_OracleMgr.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleMgr *OracleMgrCallerSession) Price() (*big.Int, error) {
	return _OracleMgr.Contract.Price(&_OracleMgr.CallOpts)
}

// SyncFrequency is a free data retrieval call binding the contract method 0xcdee7e07.
//
// Solidity: function syncFrequency() constant returns(uint256)
func (_OracleMgr *OracleMgrCaller) SyncFrequency(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "syncFrequency")
	return *ret0, err
}

// SyncFrequency is a free data retrieval call binding the contract method 0xcdee7e07.
//
// Solidity: function syncFrequency() constant returns(uint256)
func (_OracleMgr *OracleMgrSession) SyncFrequency() (*big.Int, error) {
	return _OracleMgr.Contract.SyncFrequency(&_OracleMgr.CallOpts)
}

// SyncFrequency is a free data retrieval call binding the contract method 0xcdee7e07.
//
// Solidity: function syncFrequency() constant returns(uint256)
func (_OracleMgr *OracleMgrCallerSession) SyncFrequency() (*big.Int, error) {
	return _OracleMgr.Contract.SyncFrequency(&_OracleMgr.CallOpts)
}

// AddPrice is a paid mutator transaction binding the contract method 0xe9f0ee56.
//
// Solidity: function addPrice(_price uint256) returns()
func (_OracleMgr *OracleMgrTransactor) AddPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "addPrice", _price)
}

// AddPrice is a paid mutator transaction binding the contract method 0xe9f0ee56.
//
// Solidity: function addPrice(_price uint256) returns()
func (_OracleMgr *OracleMgrSession) AddPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleMgr.Contract.AddPrice(&_OracleMgr.TransactOpts, _price)
}

// AddPrice is a paid mutator transaction binding the contract method 0xe9f0ee56.
//
// Solidity: function addPrice(_price uint256) returns()
func (_OracleMgr *OracleMgrTransactorSession) AddPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleMgr.Contract.AddPrice(&_OracleMgr.TransactOpts, _price)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleMgr *OracleMgrTransactor) DeregisterOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "deregisterOracle")
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleMgr *OracleMgrSession) DeregisterOracle() (*types.Transaction, error) {
	return _OracleMgr.Contract.DeregisterOracle(&_OracleMgr.TransactOpts)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0xf93a2eb2.
//
// Solidity: function deregisterOracle() returns()
func (_OracleMgr *OracleMgrTransactorSession) DeregisterOracle() (*types.Transaction, error) {
	return _OracleMgr.Contract.DeregisterOracle(&_OracleMgr.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleMgr *OracleMgrTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleMgr *OracleMgrSession) Pause() (*types.Transaction, error) {
	return _OracleMgr.Contract.Pause(&_OracleMgr.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleMgr *OracleMgrTransactorSession) Pause() (*types.Transaction, error) {
	return _OracleMgr.Contract.Pause(&_OracleMgr.TransactOpts)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleMgr *OracleMgrTransactor) RegisterOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "registerOracle")
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleMgr *OracleMgrSession) RegisterOracle() (*types.Transaction, error) {
	return _OracleMgr.Contract.RegisterOracle(&_OracleMgr.TransactOpts)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0x339d2590.
//
// Solidity: function registerOracle() returns()
func (_OracleMgr *OracleMgrTransactorSession) RegisterOracle() (*types.Transaction, error) {
	return _OracleMgr.Contract.RegisterOracle(&_OracleMgr.TransactOpts)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_OracleMgr *OracleMgrTransactor) ReleaseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "releaseDeposits")
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_OracleMgr *OracleMgrSession) ReleaseDeposits() (*types.Transaction, error) {
	return _OracleMgr.Contract.ReleaseDeposits(&_OracleMgr.TransactOpts)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_OracleMgr *OracleMgrTransactorSession) ReleaseDeposits() (*types.Transaction, error) {
	return _OracleMgr.Contract.ReleaseDeposits(&_OracleMgr.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_OracleMgr *OracleMgrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_OracleMgr *OracleMgrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.TransferOwnership(&_OracleMgr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_OracleMgr *OracleMgrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.TransferOwnership(&_OracleMgr.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleMgr *OracleMgrTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleMgr *OracleMgrSession) Unpause() (*types.Transaction, error) {
	return _OracleMgr.Contract.Unpause(&_OracleMgr.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleMgr *OracleMgrTransactorSession) Unpause() (*types.Transaction, error) {
	return _OracleMgr.Contract.Unpause(&_OracleMgr.TransactOpts)
}
