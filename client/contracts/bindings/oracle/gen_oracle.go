// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

import (
	"math/big"
	"strings"

	kowala "github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
)

// OracleMgrABI is the input ABI used to generate the binding from.
const OracleMgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxNumOracles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getOracleAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxNumOracles\",\"type\":\"uint256\"},{\"name\":\"_syncFrequency\",\"type\":\"uint256\"},{\"name\":\"_updatePeriod\",\"type\":\"uint256\"},{\"name\":\"_resolverAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracleCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"deregisterOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"submitPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"knsResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"updatePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isOracle\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPriceCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPriceAtIndex\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"oracle\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"syncFrequency\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"registerOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_maxNumOracles\",\"type\":\"uint256\"},{\"name\":\"_syncFrequency\",\"type\":\"uint256\"},{\"name\":\"_updatePeriod\",\"type\":\"uint256\"},{\"name\":\"_resolverAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OracleMgrBin is the compiled bytecode used for deploying new contracts.
const OracleMgrBin = `608060405260008060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b5060405160808061162283398101806040528101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000841115156100ba57600080fd5b60008311156100df576000821180156100d35750828211155b15156100de57600080fd5b5b83600181905550826002819055508160038190555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550733b058a1a62e59d185618f64bebbaf3c52bf099e063098799626040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260138152602001807f76616c696461746f726d67722e6b6f77616c610000000000000000000000000081525060200191505060206040518083038186803b1580156101d157600080fd5b505af41580156101e5573d6000803e3d6000fd5b505050506040513d60208110156101fb57600080fd5b810190808051906020019092919050505060068160001916905550505050506113f9806102296000396000f30060806040526004361061011c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062fe7b111461012157806309fe9d391461014c578063158ef93e146101b95780631f8d519d146101e85780633f4ba83a146102495780633f4e4251146102605780635c975abb1461028b578063715018a6146102ba5780638456cb59146102d15780638da5cb5b146102e85780638df4e3e21461033f578063986fcbe914610382578063a035b1fe146103af578063a2207c6a146103da578063a83627de14610431578063a97e5c931461045c578063c48c1a71146104b7578063c8104e01146104e2578063cdee7e0714610556578063e9babaec14610581578063f2fde38b146105c4575b600080fd5b34801561012d57600080fd5b50610136610607565b6040518082815260200191505060405180910390f35b34801561015857600080fd5b506101776004803603810190808035906020019092919050505061060d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101c557600080fd5b506101ce610693565b604051808215151515815260200191505060405180910390f35b3480156101f457600080fd5b50610247600480360381019080803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106a6565b005b34801561025557600080fd5b5061025e6108dd565b005b34801561026c57600080fd5b5061027561099b565b6040518082815260200191505060405180910390f35b34801561029757600080fd5b506102a06109a8565b604051808215151515815260200191505060405180910390f35b3480156102c657600080fd5b506102cf6109bb565b005b3480156102dd57600080fd5b506102e6610abd565b005b3480156102f457600080fd5b506102fd610b7d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034b57600080fd5b50610380600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ba2565b005b34801561038e57600080fd5b506103ad60048036038101908080359060200190929190505050610c3b565b005b3480156103bb57600080fd5b506103c4610dd3565b6040518082815260200191505060405180910390f35b3480156103e657600080fd5b506103ef610dd9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561043d57600080fd5b50610446610dff565b6040518082815260200191505060405180910390f35b34801561046857600080fd5b5061049d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e05565b604051808215151515815260200191505060405180910390f35b3480156104c357600080fd5b506104cc610e5e565b6040518082815260200191505060405180910390f35b3480156104ee57600080fd5b5061050d60048036038101908080359060200190929190505050610e6b565b604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b34801561056257600080fd5b5061056b610ec3565b6040518082815260200191505060405180910390f35b34801561058d57600080fd5b506105c2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ec9565b005b3480156105d057600080fd5b50610605600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f76565b005b60015481565b60008060088381548110151561061f57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905050919050565b600060159054906101000a900460ff1681565b600060159054906101000a900460ff16151515610751576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b60008411151561076057600080fd5b6000831115610785576000821180156107795750828211155b151561078457600080fd5b5b83600181905550826002819055508160038190555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550733b058a1a62e59d185618f64bebbaf3c52bf099e063098799626040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260138152602001807f76616c696461746f726d67722e6b6f77616c610000000000000000000000000081525060200191505060206040518083038186803b15801561087757600080fd5b505af415801561088b573d6000803e3d6000fd5b505050506040513d60208110156108a157600080fd5b8101908080519060200190929190505050600681600019169055506001600060156101000a81548160ff02191690831515021790555050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561093857600080fd5b600060149054906101000a900460ff16151561095357600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6000600880549050905090565b600060149054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a1657600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b1857600080fd5b600060149054906101000a900460ff16151515610b3457600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060149054906101000a900460ff16151515610bbe57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c1957600080fd5b80610c2381610e05565b1515610c2e57600080fd5b610c3782610fdd565b5050565b600060149054906101000a900460ff16151515610c5757600080fd5b33610c6181610e05565b1515610c6c57600080fd5b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff16151515610cc857600080fd5b6001600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160016101000a81548160ff021916908315150217905550600960408051908101604052808481526020013373ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b60045481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b6000600980549050905090565b6000806000600984815481101515610e7f57fe5b90600052602060002090600202019050806000015492508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915050915091565b60025481565b600060149054906101000a900460ff16151515610ee557600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f4057600080fd5b80610f4a81610e05565b151515610f5657600080fd5b610f5e61119b565b1515610f6957600080fd5b610f72826111ae565b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fd157600080fd5b610fda81611282565b50565b6000806000600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020925082600001549150600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090556001820160006101000a81549060ff02191690556001820160016101000a81549060ff0219169055505060086001600880549050038154811015156110b357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806008838154811015156110f057fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506008805480919060019003611194919061137c565b5050505050565b6000806008805490506001540311905090565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600160088390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003816000018190555060018160010160006101000a81548160ff0219169083151502179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156112be57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8154818355818111156113a3578183600052602060002091820191016113a291906113a8565b5b505050565b6113ca91905b808211156113c65760008160009055506001016113ae565b5090565b905600a165627a7a72305820408eb5444c69e406a31abf8f5256e82f610f9db7375050d039fb57d2a3eecce80029`

// DeployOracleMgr deploys a new Kowala contract, binding an instance of OracleMgr to it.
func DeployOracleMgr(auth *bind.TransactOpts, backend bind.ContractBackend, _maxNumOracles *big.Int, _syncFrequency *big.Int, _updatePeriod *big.Int, _resolverAddr common.Address) (common.Address, *types.Transaction, *OracleMgr, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleMgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleMgrBin), backend, _maxNumOracles, _syncFrequency, _updatePeriod, _resolverAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleMgr{OracleMgrCaller: OracleMgrCaller{contract: contract}, OracleMgrTransactor: OracleMgrTransactor{contract: contract}, OracleMgrFilterer: OracleMgrFilterer{contract: contract}}, nil
}

// OracleMgr is an auto generated Go binding around a Kowala contract.
type OracleMgr struct {
	OracleMgrCaller     // Read-only binding to the contract
	OracleMgrTransactor // Write-only binding to the contract
	OracleMgrFilterer   // Log filterer for contract events
}

// OracleMgrCaller is an auto generated read-only Go binding around a Kowala contract.
type OracleMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleMgrTransactor is an auto generated write-only Go binding around a Kowala contract.
type OracleMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleMgrFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type OracleMgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleMgrSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type OracleMgrSession struct {
	Contract     *OracleMgr        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleMgrCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type OracleMgrCallerSession struct {
	Contract *OracleMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OracleMgrTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type OracleMgrTransactorSession struct {
	Contract     *OracleMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OracleMgrRaw is an auto generated low-level Go binding around a Kowala contract.
type OracleMgrRaw struct {
	Contract *OracleMgr // Generic contract binding to access the raw methods on
}

// OracleMgrCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type OracleMgrCallerRaw struct {
	Contract *OracleMgrCaller // Generic read-only contract binding to access the raw methods on
}

// OracleMgrTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type OracleMgrTransactorRaw struct {
	Contract *OracleMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleMgr creates a new instance of OracleMgr, bound to a specific deployed contract.
func NewOracleMgr(address common.Address, backend bind.ContractBackend) (*OracleMgr, error) {
	contract, err := bindOracleMgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleMgr{OracleMgrCaller: OracleMgrCaller{contract: contract}, OracleMgrTransactor: OracleMgrTransactor{contract: contract}, OracleMgrFilterer: OracleMgrFilterer{contract: contract}}, nil
}

// NewOracleMgrCaller creates a new read-only instance of OracleMgr, bound to a specific deployed contract.
func NewOracleMgrCaller(address common.Address, caller bind.ContractCaller) (*OracleMgrCaller, error) {
	contract, err := bindOracleMgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleMgrCaller{contract: contract}, nil
}

// NewOracleMgrTransactor creates a new write-only instance of OracleMgr, bound to a specific deployed contract.
func NewOracleMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleMgrTransactor, error) {
	contract, err := bindOracleMgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleMgrTransactor{contract: contract}, nil
}

// NewOracleMgrFilterer creates a new log filterer instance of OracleMgr, bound to a specific deployed contract.
func NewOracleMgrFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleMgrFilterer, error) {
	contract, err := bindOracleMgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleMgrFilterer{contract: contract}, nil
}

// bindOracleMgr binds a generic wrapper to an already deployed contract.
func bindOracleMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleMgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetOracleAtIndex is a free data retrieval call binding the contract method 0x09fe9d39.
//
// Solidity: function getOracleAtIndex(index uint256) constant returns(code address)
func (_OracleMgr *OracleMgrCaller) GetOracleAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "getOracleAtIndex", index)
	return *ret0, err
}

// GetOracleAtIndex is a free data retrieval call binding the contract method 0x09fe9d39.
//
// Solidity: function getOracleAtIndex(index uint256) constant returns(code address)
func (_OracleMgr *OracleMgrSession) GetOracleAtIndex(index *big.Int) (common.Address, error) {
	return _OracleMgr.Contract.GetOracleAtIndex(&_OracleMgr.CallOpts, index)
}

// GetOracleAtIndex is a free data retrieval call binding the contract method 0x09fe9d39.
//
// Solidity: function getOracleAtIndex(index uint256) constant returns(code address)
func (_OracleMgr *OracleMgrCallerSession) GetOracleAtIndex(index *big.Int) (common.Address, error) {
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

// GetPriceAtIndex is a free data retrieval call binding the contract method 0xc8104e01.
//
// Solidity: function getPriceAtIndex(index uint256) constant returns(price uint256, oracle address)
func (_OracleMgr *OracleMgrCaller) GetPriceAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Price  *big.Int
	Oracle common.Address
}, error) {
	ret := new(struct {
		Price  *big.Int
		Oracle common.Address
	})
	out := ret
	err := _OracleMgr.contract.Call(opts, out, "getPriceAtIndex", index)
	return *ret, err
}

// GetPriceAtIndex is a free data retrieval call binding the contract method 0xc8104e01.
//
// Solidity: function getPriceAtIndex(index uint256) constant returns(price uint256, oracle address)
func (_OracleMgr *OracleMgrSession) GetPriceAtIndex(index *big.Int) (struct {
	Price  *big.Int
	Oracle common.Address
}, error) {
	return _OracleMgr.Contract.GetPriceAtIndex(&_OracleMgr.CallOpts, index)
}

// GetPriceAtIndex is a free data retrieval call binding the contract method 0xc8104e01.
//
// Solidity: function getPriceAtIndex(index uint256) constant returns(price uint256, oracle address)
func (_OracleMgr *OracleMgrCallerSession) GetPriceAtIndex(index *big.Int) (struct {
	Price  *big.Int
	Oracle common.Address
}, error) {
	return _OracleMgr.Contract.GetPriceAtIndex(&_OracleMgr.CallOpts, index)
}

// GetPriceCount is a free data retrieval call binding the contract method 0xc48c1a71.
//
// Solidity: function getPriceCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrCaller) GetPriceCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "getPriceCount")
	return *ret0, err
}

// GetPriceCount is a free data retrieval call binding the contract method 0xc48c1a71.
//
// Solidity: function getPriceCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrSession) GetPriceCount() (*big.Int, error) {
	return _OracleMgr.Contract.GetPriceCount(&_OracleMgr.CallOpts)
}

// GetPriceCount is a free data retrieval call binding the contract method 0xc48c1a71.
//
// Solidity: function getPriceCount() constant returns(count uint256)
func (_OracleMgr *OracleMgrCallerSession) GetPriceCount() (*big.Int, error) {
	return _OracleMgr.Contract.GetPriceCount(&_OracleMgr.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_OracleMgr *OracleMgrCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_OracleMgr *OracleMgrSession) Initialized() (bool, error) {
	return _OracleMgr.Contract.Initialized(&_OracleMgr.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_OracleMgr *OracleMgrCallerSession) Initialized() (bool, error) {
	return _OracleMgr.Contract.Initialized(&_OracleMgr.CallOpts)
}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(identity address) constant returns(isIndeed bool)
func (_OracleMgr *OracleMgrCaller) IsOracle(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "isOracle", identity)
	return *ret0, err
}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(identity address) constant returns(isIndeed bool)
func (_OracleMgr *OracleMgrSession) IsOracle(identity common.Address) (bool, error) {
	return _OracleMgr.Contract.IsOracle(&_OracleMgr.CallOpts, identity)
}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(identity address) constant returns(isIndeed bool)
func (_OracleMgr *OracleMgrCallerSession) IsOracle(identity common.Address) (bool, error) {
	return _OracleMgr.Contract.IsOracle(&_OracleMgr.CallOpts, identity)
}

// KnsResolver is a free data retrieval call binding the contract method 0xa2207c6a.
//
// Solidity: function knsResolver() constant returns(address)
func (_OracleMgr *OracleMgrCaller) KnsResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "knsResolver")
	return *ret0, err
}

// KnsResolver is a free data retrieval call binding the contract method 0xa2207c6a.
//
// Solidity: function knsResolver() constant returns(address)
func (_OracleMgr *OracleMgrSession) KnsResolver() (common.Address, error) {
	return _OracleMgr.Contract.KnsResolver(&_OracleMgr.CallOpts)
}

// KnsResolver is a free data retrieval call binding the contract method 0xa2207c6a.
//
// Solidity: function knsResolver() constant returns(address)
func (_OracleMgr *OracleMgrCallerSession) KnsResolver() (common.Address, error) {
	return _OracleMgr.Contract.KnsResolver(&_OracleMgr.CallOpts)
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

// UpdatePeriod is a free data retrieval call binding the contract method 0xa83627de.
//
// Solidity: function updatePeriod() constant returns(uint256)
func (_OracleMgr *OracleMgrCaller) UpdatePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleMgr.contract.Call(opts, out, "updatePeriod")
	return *ret0, err
}

// UpdatePeriod is a free data retrieval call binding the contract method 0xa83627de.
//
// Solidity: function updatePeriod() constant returns(uint256)
func (_OracleMgr *OracleMgrSession) UpdatePeriod() (*big.Int, error) {
	return _OracleMgr.Contract.UpdatePeriod(&_OracleMgr.CallOpts)
}

// UpdatePeriod is a free data retrieval call binding the contract method 0xa83627de.
//
// Solidity: function updatePeriod() constant returns(uint256)
func (_OracleMgr *OracleMgrCallerSession) UpdatePeriod() (*big.Int, error) {
	return _OracleMgr.Contract.UpdatePeriod(&_OracleMgr.CallOpts)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0x8df4e3e2.
//
// Solidity: function deregisterOracle(identity address) returns()
func (_OracleMgr *OracleMgrTransactor) DeregisterOracle(opts *bind.TransactOpts, identity common.Address) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "deregisterOracle", identity)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0x8df4e3e2.
//
// Solidity: function deregisterOracle(identity address) returns()
func (_OracleMgr *OracleMgrSession) DeregisterOracle(identity common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.DeregisterOracle(&_OracleMgr.TransactOpts, identity)
}

// DeregisterOracle is a paid mutator transaction binding the contract method 0x8df4e3e2.
//
// Solidity: function deregisterOracle(identity address) returns()
func (_OracleMgr *OracleMgrTransactorSession) DeregisterOracle(identity common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.DeregisterOracle(&_OracleMgr.TransactOpts, identity)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f8d519d.
//
// Solidity: function initialize(_maxNumOracles uint256, _syncFrequency uint256, _updatePeriod uint256, _resolverAddr address) returns()
func (_OracleMgr *OracleMgrTransactor) Initialize(opts *bind.TransactOpts, _maxNumOracles *big.Int, _syncFrequency *big.Int, _updatePeriod *big.Int, _resolverAddr common.Address) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "initialize", _maxNumOracles, _syncFrequency, _updatePeriod, _resolverAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f8d519d.
//
// Solidity: function initialize(_maxNumOracles uint256, _syncFrequency uint256, _updatePeriod uint256, _resolverAddr address) returns()
func (_OracleMgr *OracleMgrSession) Initialize(_maxNumOracles *big.Int, _syncFrequency *big.Int, _updatePeriod *big.Int, _resolverAddr common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.Initialize(&_OracleMgr.TransactOpts, _maxNumOracles, _syncFrequency, _updatePeriod, _resolverAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f8d519d.
//
// Solidity: function initialize(_maxNumOracles uint256, _syncFrequency uint256, _updatePeriod uint256, _resolverAddr address) returns()
func (_OracleMgr *OracleMgrTransactorSession) Initialize(_maxNumOracles *big.Int, _syncFrequency *big.Int, _updatePeriod *big.Int, _resolverAddr common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.Initialize(&_OracleMgr.TransactOpts, _maxNumOracles, _syncFrequency, _updatePeriod, _resolverAddr)
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

// RegisterOracle is a paid mutator transaction binding the contract method 0xe9babaec.
//
// Solidity: function registerOracle(identity address) returns()
func (_OracleMgr *OracleMgrTransactor) RegisterOracle(opts *bind.TransactOpts, identity common.Address) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "registerOracle", identity)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0xe9babaec.
//
// Solidity: function registerOracle(identity address) returns()
func (_OracleMgr *OracleMgrSession) RegisterOracle(identity common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.RegisterOracle(&_OracleMgr.TransactOpts, identity)
}

// RegisterOracle is a paid mutator transaction binding the contract method 0xe9babaec.
//
// Solidity: function registerOracle(identity address) returns()
func (_OracleMgr *OracleMgrTransactorSession) RegisterOracle(identity common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.RegisterOracle(&_OracleMgr.TransactOpts, identity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleMgr *OracleMgrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleMgr *OracleMgrSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleMgr.Contract.RenounceOwnership(&_OracleMgr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleMgr *OracleMgrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleMgr.Contract.RenounceOwnership(&_OracleMgr.TransactOpts)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(_price uint256) returns()
func (_OracleMgr *OracleMgrTransactor) SubmitPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "submitPrice", _price)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(_price uint256) returns()
func (_OracleMgr *OracleMgrSession) SubmitPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleMgr.Contract.SubmitPrice(&_OracleMgr.TransactOpts, _price)
}

// SubmitPrice is a paid mutator transaction binding the contract method 0x986fcbe9.
//
// Solidity: function submitPrice(_price uint256) returns()
func (_OracleMgr *OracleMgrTransactorSession) SubmitPrice(_price *big.Int) (*types.Transaction, error) {
	return _OracleMgr.Contract.SubmitPrice(&_OracleMgr.TransactOpts, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_OracleMgr *OracleMgrTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _OracleMgr.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_OracleMgr *OracleMgrSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.TransferOwnership(&_OracleMgr.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_OracleMgr *OracleMgrTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OracleMgr.Contract.TransferOwnership(&_OracleMgr.TransactOpts, _newOwner)
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

// OracleMgrOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the OracleMgr contract.
type OracleMgrOwnershipRenouncedIterator struct {
	Event *OracleMgrOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleMgrOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleMgrOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleMgrOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleMgrOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleMgrOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleMgrOwnershipRenounced represents a OwnershipRenounced event raised by the OracleMgr contract.
type OracleMgrOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_OracleMgr *OracleMgrFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OracleMgrOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _OracleMgr.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleMgrOwnershipRenouncedIterator{contract: _OracleMgr.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_OracleMgr *OracleMgrFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OracleMgrOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _OracleMgr.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleMgrOwnershipRenounced)
				if err := _OracleMgr.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OracleMgrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OracleMgr contract.
type OracleMgrOwnershipTransferredIterator struct {
	Event *OracleMgrOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleMgrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleMgrOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleMgrOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleMgrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleMgrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleMgrOwnershipTransferred represents a OwnershipTransferred event raised by the OracleMgr contract.
type OracleMgrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_OracleMgr *OracleMgrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleMgrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleMgr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleMgrOwnershipTransferredIterator{contract: _OracleMgr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_OracleMgr *OracleMgrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleMgrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleMgr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleMgrOwnershipTransferred)
				if err := _OracleMgr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OracleMgrPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the OracleMgr contract.
type OracleMgrPauseIterator struct {
	Event *OracleMgrPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleMgrPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleMgrPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleMgrPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleMgrPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleMgrPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleMgrPause represents a Pause event raised by the OracleMgr contract.
type OracleMgrPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_OracleMgr *OracleMgrFilterer) FilterPause(opts *bind.FilterOpts) (*OracleMgrPauseIterator, error) {

	logs, sub, err := _OracleMgr.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &OracleMgrPauseIterator{contract: _OracleMgr.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_OracleMgr *OracleMgrFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *OracleMgrPause) (event.Subscription, error) {

	logs, sub, err := _OracleMgr.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleMgrPause)
				if err := _OracleMgr.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OracleMgrUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the OracleMgr contract.
type OracleMgrUnpauseIterator struct {
	Event *OracleMgrUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleMgrUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleMgrUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleMgrUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleMgrUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleMgrUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleMgrUnpause represents a Unpause event raised by the OracleMgr contract.
type OracleMgrUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_OracleMgr *OracleMgrFilterer) FilterUnpause(opts *bind.FilterOpts) (*OracleMgrUnpauseIterator, error) {

	logs, sub, err := _OracleMgr.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &OracleMgrUnpauseIterator{contract: _OracleMgr.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_OracleMgr *OracleMgrFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *OracleMgrUnpause) (event.Subscription, error) {

	logs, sub, err := _OracleMgr.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleMgrUnpause)
				if err := _OracleMgr.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
