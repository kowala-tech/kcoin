// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
)

// ValidatorMgrABI is the input ABI used to generate the binding from.
const ValidatorMgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningTokenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"availableAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxNumValidators\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"},{\"name\":\"_miningTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ValidatorMgrBin is the compiled bytecode used for deploying new contracts.
const ValidatorMgrBin = `606060405260008060146101000a81548160ff021916908315150217905550341561002957600080fd5b60405160808061188583398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183101515156100b057600080fd5b836001819055508260028190555062015180820260038190555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506117678061011e6000396000f300606060405260043610610128576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461012d5780630a3cb663146101565780632086ca251461017f57806327378a8c146101a85780633e83a283146101fd5780633ed0a373146102825780633f4ba83a146102c05780635c975abb146102d5578063671b4d491461030257806369474625146103175780636a911ccf146103405780637071688a146103555780638456cb591461037e5780638da5cb5b146103935780639363a141146103e857806397584b3e14610411578063aded41ec1461043e578063b774cb1e14610453578063cefddda914610484578063e7a60a9c146104d5578063f2fde38b1461053f578063facd743b14610578575b600080fd5b341561013857600080fd5b6101406105c9565b6040518082815260200191505060405180910390f35b341561016157600080fd5b61016961069d565b6040518082815260200191505060405180910390f35b341561018a57600080fd5b6101926106a3565b6040518082815260200191505060405180910390f35b34156101b357600080fd5b6101bb6106a9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561020857600080fd5b610280600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506106cf565b005b341561028d57600080fd5b6102a3600480803590602001909190505061075d565b604051808381526020018281526020019250505060405180910390f35b34156102cb57600080fd5b6102d36107d5565b005b34156102e057600080fd5b6102e8610893565b604051808215151515815260200191505060405180910390f35b341561030d57600080fd5b6103156108a6565b005b341561032257600080fd5b61032a610964565b6040518082815260200191505060405180910390f35b341561034b57600080fd5b61035361096a565b005b341561036057600080fd5b6103686109a5565b6040518082815260200191505060405180910390f35b341561038957600080fd5b6103916109b2565b005b341561039e57600080fd5b6103a6610a72565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103f357600080fd5b6103fb610a97565b6040518082815260200191505060405180910390f35b341561041c57600080fd5b610424610ae4565b604051808215151515815260200191505060405180910390f35b341561044957600080fd5b610451610af7565b005b341561045e57600080fd5b610466610cf4565b60405180826000191660001916815260200191505060405180910390f35b341561048f57600080fd5b6104bb600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610cfa565b604051808215151515815260200191505060405180910390f35b34156104e057600080fd5b6104f66004808035906020019091905050610d53565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b341561054a57600080fd5b610576600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610e0b565b005b341561058357600080fd5b6105af600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610f60565b604051808215151515815260200191505060405180910390f35b6000806105d4610ae4565b156105e3576001549150610699565b6006600060076001600780549050038154811015156105fe57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561068357fe5b9060005260206000209060020201600001540191505b5090565b60035481565b60025481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60408051908101604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815250600860008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101559050506107586108a6565b505050565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156107b157fe5b90600052602060002090600202019050806000015481600101549250925050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561083057600080fd5b600060149054906101000a900460ff16151561084b57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b600060149054906101000a900460ff161515156108c257600080fd5b6108f0600860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610f60565b1515156108fc57600080fd5b6109046105c9565b6008600101541015151561091757600080fd5b61091f610ae4565b151561092e5761092d610fb9565b5b610962600860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600860010154611006565b565b60015481565b600060149054906101000a900460ff1615151561098657600080fd5b61098f33610f60565b151561099a57600080fd5b6109a333611347565b565b6000600780549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a0d57600080fd5b600060149054906101000a900460ff16151515610a2957600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506002540311905090565b600080600080600060149054906101000a900460ff16151515610b1957600080fd5b6000935060009250600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020191505b818054905083108015610b99575060008284815481101515610b8457fe5b90600052602060002090600202016001015414155b15610bfb578183815481101515610bac57fe5b906000526020600020906002020160010154421015610bca57610bfb565b8183815481101515610bd857fe5b906000526020600020906002020160000154840193508280600101935050610b66565b610c0533846114bb565b6000841115610cee57600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1515610cd557600080fd5b5af11515610ce257600080fd5b50505060405180519050505b50505050565b60045481565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff169050919050565b6000806000600784815481101515610d6757fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806002016001826002018054905003815481101515610df157fe5b906000526020600020906002020160000154915050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e6657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610ea257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b6110046007600160078054905003815481101515610fd357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611347565b565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078054806001018281611063919061162b565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff02191690831515021790555060004314156110fe5760018460010160016101000a81548160ff0219169083151502179055505b8360020180548060010182816111149190611657565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115611337576006600060076001860381548110151561117d57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915081600201600183600201805490500381548110151561120057fe5b9060005260206000209060020201905080600001548511151561122257611337565b60076001840381548110151561123457fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078481548110151561126f57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856007600185038154811015156112cb57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082826000018190555060018303846000018190555082806001900393505061115e565b61133f6115a8565b505050505050565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160078054905003811015611446576007600182018154811015156113b557fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007828154811015156113f057fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050611393565b600780548091906001900361145b9190611689565b5060008260010160006101000a81548160ff021916908315150217905550600354420182600201600184600201805490500381548110151561149957fe5b9060005260206000209060020201600101819055506114b66115a8565b505050565b6000806000808414156114cd576115a1565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b826002018054905081101561158f57826002018181548110151561153657fe5b9060005260206000209060020201836002018381548110151561155557fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050611516565b81836002018161159f91906116b5565b505b5050505050565b600760405180828054801561161257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116115c8575b5050915050604051809103902060048160001916905550565b8154818355818115116116525781836000526020600020918201910161165191906116e7565b5b505050565b81548183558181151161168457600202816002028360005260206000209182019101611683919061170c565b5b505050565b8154818355818115116116b0578183600052602060002091820191016116af91906116e7565b5b505050565b8154818355818115116116e2576002028160020283600052602060002091820191016116e1919061170c565b5b505050565b61170991905b808211156117055760008160009055506001016116ed565b5090565b90565b61173891905b8082111561173457600080820160009055600182016000905550600201611712565b5090565b905600a165627a7a72305820467077fdbdb35430908438b68c2a737e6887fd063da3be038c0c9fde85b0d3da0029`

// DeployValidatorMgr deploys a new Ethereum contract, binding an instance of ValidatorMgr to it.
func DeployValidatorMgr(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxNumValidators *big.Int, _freezePeriod *big.Int, _miningTokenAddr common.Address) (common.Address, *types.Transaction, *ValidatorMgr, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorMgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorMgrBin), backend, _baseDeposit, _maxNumValidators, _freezePeriod, _miningTokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMgr{ValidatorMgrCaller: ValidatorMgrCaller{contract: contract}, ValidatorMgrTransactor: ValidatorMgrTransactor{contract: contract}}, nil
}

// ValidatorMgr is an auto generated Go binding around an Ethereum contract.
type ValidatorMgr struct {
	ValidatorMgrCaller     // Read-only binding to the contract
	ValidatorMgrTransactor // Write-only binding to the contract
}

// ValidatorMgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorMgrSession struct {
	Contract     *ValidatorMgr     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorMgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorMgrCallerSession struct {
	Contract *ValidatorMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorMgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorMgrTransactorSession struct {
	Contract     *ValidatorMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorMgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorMgrRaw struct {
	Contract *ValidatorMgr // Generic contract binding to access the raw methods on
}

// ValidatorMgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorMgrCallerRaw struct {
	Contract *ValidatorMgrCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorMgrTransactorRaw struct {
	Contract *ValidatorMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorMgr creates a new instance of ValidatorMgr, bound to a specific deployed contract.
func NewValidatorMgr(address common.Address, backend bind.ContractBackend) (*ValidatorMgr, error) {
	contract, err := bindValidatorMgr(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgr{ValidatorMgrCaller: ValidatorMgrCaller{contract: contract}, ValidatorMgrTransactor: ValidatorMgrTransactor{contract: contract}}, nil
}

// NewValidatorMgrCaller creates a new read-only instance of ValidatorMgr, bound to a specific deployed contract.
func NewValidatorMgrCaller(address common.Address, caller bind.ContractCaller) (*ValidatorMgrCaller, error) {
	contract, err := bindValidatorMgr(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrCaller{contract: contract}, nil
}

// NewValidatorMgrTransactor creates a new write-only instance of ValidatorMgr, bound to a specific deployed contract.
func NewValidatorMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorMgrTransactor, error) {
	contract, err := bindValidatorMgr(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrTransactor{contract: contract}, nil
}

// bindValidatorMgr binds a generic wrapper to an already deployed contract.
func bindValidatorMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorMgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMgr *ValidatorMgrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorMgr.Contract.ValidatorMgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMgr *ValidatorMgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.ValidatorMgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMgr *ValidatorMgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.ValidatorMgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMgr *ValidatorMgrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorMgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMgr *ValidatorMgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMgr *ValidatorMgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.contract.Transact(opts, method, params...)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorMgr *ValidatorMgrCaller) _hasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorMgr *ValidatorMgrSession) _hasAvailability() (bool, error) {
	return _ValidatorMgr.Contract._hasAvailability(&_ValidatorMgr.CallOpts)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorMgr *ValidatorMgrCallerSession) _hasAvailability() (bool, error) {
	return _ValidatorMgr.Contract._hasAvailability(&_ValidatorMgr.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrSession) BaseDeposit() (*big.Int, error) {
	return _ValidatorMgr.Contract.BaseDeposit(&_ValidatorMgr.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) BaseDeposit() (*big.Int, error) {
	return _ValidatorMgr.Contract.BaseDeposit(&_ValidatorMgr.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCaller) FreezePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "freezePeriod")
	return *ret0, err
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrSession) FreezePeriod() (*big.Int, error) {
	return _ValidatorMgr.Contract.FreezePeriod(&_ValidatorMgr.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) FreezePeriod() (*big.Int, error) {
	return _ValidatorMgr.Contract.FreezePeriod(&_ValidatorMgr.CallOpts)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ValidatorMgr *ValidatorMgrCaller) GetDepositAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	ret := new(struct {
		Amount      *big.Int
		AvailableAt *big.Int
	})
	out := ret
	err := _ValidatorMgr.contract.Call(opts, out, "getDepositAtIndex", index)
	return *ret, err
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ValidatorMgr *ValidatorMgrSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	return _ValidatorMgr.Contract.GetDepositAtIndex(&_ValidatorMgr.CallOpts, index)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	return _ValidatorMgr.Contract.GetDepositAtIndex(&_ValidatorMgr.CallOpts, index)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ValidatorMgr *ValidatorMgrCaller) GetDepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "getDepositCount")
	return *ret0, err
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ValidatorMgr *ValidatorMgrSession) GetDepositCount() (*big.Int, error) {
	return _ValidatorMgr.Contract.GetDepositCount(&_ValidatorMgr.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) GetDepositCount() (*big.Int, error) {
	return _ValidatorMgr.Contract.GetDepositCount(&_ValidatorMgr.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ValidatorMgr *ValidatorMgrCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ValidatorMgr *ValidatorMgrSession) GetMinimumDeposit() (*big.Int, error) {
	return _ValidatorMgr.Contract.GetMinimumDeposit(&_ValidatorMgr.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _ValidatorMgr.Contract.GetMinimumDeposit(&_ValidatorMgr.CallOpts)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ValidatorMgr *ValidatorMgrCaller) GetValidatorAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _ValidatorMgr.contract.Call(opts, out, "getValidatorAtIndex", index)
	return *ret, err
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ValidatorMgr *ValidatorMgrSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ValidatorMgr.Contract.GetValidatorAtIndex(&_ValidatorMgr.CallOpts, index)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ValidatorMgr.Contract.GetValidatorAtIndex(&_ValidatorMgr.CallOpts, index)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ValidatorMgr *ValidatorMgrCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "getValidatorCount")
	return *ret0, err
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ValidatorMgr *ValidatorMgrSession) GetValidatorCount() (*big.Int, error) {
	return _ValidatorMgr.Contract.GetValidatorCount(&_ValidatorMgr.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) GetValidatorCount() (*big.Int, error) {
	return _ValidatorMgr.Contract.GetValidatorCount(&_ValidatorMgr.CallOpts)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrCaller) IsGenesisValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "isGenesisValidator", code)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ValidatorMgr.Contract.IsGenesisValidator(&_ValidatorMgr.CallOpts, code)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrCallerSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ValidatorMgr.Contract.IsGenesisValidator(&_ValidatorMgr.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrCaller) IsValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "isValidator", code)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrSession) IsValidator(code common.Address) (bool, error) {
	return _ValidatorMgr.Contract.IsValidator(&_ValidatorMgr.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrCallerSession) IsValidator(code common.Address) (bool, error) {
	return _ValidatorMgr.Contract.IsValidator(&_ValidatorMgr.CallOpts, code)
}

// MaxNumValidators is a free data retrieval call binding the contract method 0x2086ca25.
//
// Solidity: function maxNumValidators() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCaller) MaxNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "maxNumValidators")
	return *ret0, err
}

// MaxNumValidators is a free data retrieval call binding the contract method 0x2086ca25.
//
// Solidity: function maxNumValidators() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrSession) MaxNumValidators() (*big.Int, error) {
	return _ValidatorMgr.Contract.MaxNumValidators(&_ValidatorMgr.CallOpts)
}

// MaxNumValidators is a free data retrieval call binding the contract method 0x2086ca25.
//
// Solidity: function maxNumValidators() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) MaxNumValidators() (*big.Int, error) {
	return _ValidatorMgr.Contract.MaxNumValidators(&_ValidatorMgr.CallOpts)
}

// MiningTokenAddr is a free data retrieval call binding the contract method 0x27378a8c.
//
// Solidity: function miningTokenAddr() constant returns(address)
func (_ValidatorMgr *ValidatorMgrCaller) MiningTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "miningTokenAddr")
	return *ret0, err
}

// MiningTokenAddr is a free data retrieval call binding the contract method 0x27378a8c.
//
// Solidity: function miningTokenAddr() constant returns(address)
func (_ValidatorMgr *ValidatorMgrSession) MiningTokenAddr() (common.Address, error) {
	return _ValidatorMgr.Contract.MiningTokenAddr(&_ValidatorMgr.CallOpts)
}

// MiningTokenAddr is a free data retrieval call binding the contract method 0x27378a8c.
//
// Solidity: function miningTokenAddr() constant returns(address)
func (_ValidatorMgr *ValidatorMgrCallerSession) MiningTokenAddr() (common.Address, error) {
	return _ValidatorMgr.Contract.MiningTokenAddr(&_ValidatorMgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidatorMgr *ValidatorMgrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidatorMgr *ValidatorMgrSession) Owner() (common.Address, error) {
	return _ValidatorMgr.Contract.Owner(&_ValidatorMgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidatorMgr *ValidatorMgrCallerSession) Owner() (common.Address, error) {
	return _ValidatorMgr.Contract.Owner(&_ValidatorMgr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ValidatorMgr *ValidatorMgrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ValidatorMgr *ValidatorMgrSession) Paused() (bool, error) {
	return _ValidatorMgr.Contract.Paused(&_ValidatorMgr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ValidatorMgr *ValidatorMgrCallerSession) Paused() (bool, error) {
	return _ValidatorMgr.Contract.Paused(&_ValidatorMgr.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ValidatorMgr *ValidatorMgrCaller) ValidatorsChecksum(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "validatorsChecksum")
	return *ret0, err
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ValidatorMgr *ValidatorMgrSession) ValidatorsChecksum() ([32]byte, error) {
	return _ValidatorMgr.Contract.ValidatorsChecksum(&_ValidatorMgr.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ValidatorMgr *ValidatorMgrCallerSession) ValidatorsChecksum() ([32]byte, error) {
	return _ValidatorMgr.Contract.ValidatorsChecksum(&_ValidatorMgr.CallOpts)
}

// _registerValidator is a paid mutator transaction binding the contract method 0x671b4d49.
//
// Solidity: function _registerValidator() returns()
func (_ValidatorMgr *ValidatorMgrTransactor) _registerValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "_registerValidator")
}

// _registerValidator is a paid mutator transaction binding the contract method 0x671b4d49.
//
// Solidity: function _registerValidator() returns()
func (_ValidatorMgr *ValidatorMgrSession) _registerValidator() (*types.Transaction, error) {
	return _ValidatorMgr.Contract._registerValidator(&_ValidatorMgr.TransactOpts)
}

// _registerValidator is a paid mutator transaction binding the contract method 0x671b4d49.
//
// Solidity: function _registerValidator() returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) _registerValidator() (*types.Transaction, error) {
	return _ValidatorMgr.Contract._registerValidator(&_ValidatorMgr.TransactOpts)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x6a911ccf.
//
// Solidity: function deregisterValidator() returns()
func (_ValidatorMgr *ValidatorMgrTransactor) DeregisterValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "deregisterValidator")
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x6a911ccf.
//
// Solidity: function deregisterValidator() returns()
func (_ValidatorMgr *ValidatorMgrSession) DeregisterValidator() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.DeregisterValidator(&_ValidatorMgr.TransactOpts)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x6a911ccf.
//
// Solidity: function deregisterValidator() returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) DeregisterValidator() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.DeregisterValidator(&_ValidatorMgr.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ValidatorMgr *ValidatorMgrTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ValidatorMgr *ValidatorMgrSession) Pause() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.Pause(&_ValidatorMgr.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) Pause() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.Pause(&_ValidatorMgr.TransactOpts)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3e83a283.
//
// Solidity: function registerValidator(_from address, _value uint256, _data bytes) returns()
func (_ValidatorMgr *ValidatorMgrTransactor) RegisterValidator(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "registerValidator", _from, _value, _data)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3e83a283.
//
// Solidity: function registerValidator(_from address, _value uint256, _data bytes) returns()
func (_ValidatorMgr *ValidatorMgrSession) RegisterValidator(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.RegisterValidator(&_ValidatorMgr.TransactOpts, _from, _value, _data)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3e83a283.
//
// Solidity: function registerValidator(_from address, _value uint256, _data bytes) returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) RegisterValidator(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.RegisterValidator(&_ValidatorMgr.TransactOpts, _from, _value, _data)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_ValidatorMgr *ValidatorMgrTransactor) ReleaseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "releaseDeposits")
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_ValidatorMgr *ValidatorMgrSession) ReleaseDeposits() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.ReleaseDeposits(&_ValidatorMgr.TransactOpts)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) ReleaseDeposits() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.ReleaseDeposits(&_ValidatorMgr.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ValidatorMgr *ValidatorMgrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ValidatorMgr *ValidatorMgrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.TransferOwnership(&_ValidatorMgr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.TransferOwnership(&_ValidatorMgr.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ValidatorMgr *ValidatorMgrTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ValidatorMgr *ValidatorMgrSession) Unpause() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.Unpause(&_ValidatorMgr.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) Unpause() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.Unpause(&_ValidatorMgr.TransactOpts)
}
