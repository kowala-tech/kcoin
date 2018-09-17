// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stability

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

// StabilityABI is the input ABI used to generate the binding from.
const StabilityABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSubscriptionAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSubscriptionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"name\":\"_priceProviderAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unsubscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"name\":\"_initialReward\",\"type\":\"uint256\"},{\"name\":\"_priceProviderAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// StabilityBin is the compiled bytecode used for deploying new contracts.
const StabilityBin = `608060405260008060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b50604051606080610fef833981018060405281019080805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826001819055508160028190555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050610eed806101026000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063158ef93e146100d55780633f4ba83a14610104578063402300461461011b57806341b3d1851461018f5780635c975abb146101ba57806366419970146101e9578063715018a6146102145780638456cb591461022b5780638da5cb5b146102425780638f449a0514610299578063abee967c146102a3578063da35a26f146102ce578063f2fde38b1461031b578063fcae44841461035e575b600080fd5b3480156100e157600080fd5b506100ea610375565b604051808215151515815260200191505060405180910390f35b34801561011057600080fd5b50610119610388565b005b34801561012757600080fd5b5061014660048036038101908080359060200190929190505050610446565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561019b57600080fd5b506101a46104d5565b6040518082815260200191505060405180910390f35b3480156101c657600080fd5b506101cf6104db565b604051808215151515815260200191505060405180910390f35b3480156101f557600080fd5b506101fe6104ee565b6040518082815260200191505060405180910390f35b34801561022057600080fd5b506102296104fb565b005b34801561023757600080fd5b506102406105fd565b005b34801561024e57600080fd5b506102576106bd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102a16106e2565b005b3480156102af57600080fd5b506102b8610771565b6040518082815260200191505060405180910390f35b3480156102da57600080fd5b5061031960048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610777565b005b34801561032757600080fd5b5061035c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610889565b005b34801561036a57600080fd5b506103736108f0565b005b600060159054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103e357600080fd5b600060149054906101000a900460ff1615156103fe57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600080600060058481548110151561045a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020154915050915091565b60015481565b600060149054906101000a900460ff1681565b6000600580549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561055657600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561065857600080fd5b600060149054906101000a900460ff1615151561067457600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff161515156106ff57600080fd5b61070833610c25565b1561076557600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905034816002016000828254019250508190555061076e565b61076d610c7e565b5b50565b60025481565b600060159054906101000a900460ff16151515610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b8160018190555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600060156101000a81548160ff0219169083151502179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108e457600080fd5b6108ed81610d76565b50565b60008060006108fe33610c25565b151561090957600080fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020925082600001549150670de0b6b3a7640000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a035b1fe6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156109e057600080fd5b505af11580156109f4573d6000803e3d6000fd5b505050506040513d6020811015610a0a57600080fd5b8101908080519060200190929190505050101515610a6e573373ffffffffffffffffffffffffffffffffffffffff166108fc84600301549081150290604051600060405180830381858888f19350505050158015610a6c573d6000803e3d6000fd5b505b3373ffffffffffffffffffffffffffffffffffffffff166108fc84600201549081150290604051600060405180830381858888f19350505050158015610ab8573d6000803e3d6000fd5b50600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090556001820160006101000a81549060ff02191690556002820160009055600382016000905550506005600160058054905003815481101515610b3e57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600583815481101515610b7b57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506005805480919060019003610c1f9190610e70565b50505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b60006001543410151515610c9157600080fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600160053390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003816000018190555060018160010160006101000a81548160ff021916908315150217905550348160020181905550600254816003018190555050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610db257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b815481835581811115610e9757818360005260206000209182019101610e969190610e9c565b5b505050565b610ebe91905b80821115610eba576000816000905550600101610ea2565b5090565b905600a165627a7a7230582063cbab29fd1b591cea4ecd0150ffaff1b422e33494cb38487948ad4e0a5e9c210029`

// DeployStability deploys a new Kowala contract, binding an instance of Stability to it.
func DeployStability(auth *bind.TransactOpts, backend bind.ContractBackend, _minDeposit *big.Int, _initialReward *big.Int, _priceProviderAddr common.Address) (common.Address, *types.Transaction, *Stability, error) {
	parsed, err := abi.JSON(strings.NewReader(StabilityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StabilityBin), backend, _minDeposit, _initialReward, _priceProviderAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stability{StabilityCaller: StabilityCaller{contract: contract}, StabilityTransactor: StabilityTransactor{contract: contract}, StabilityFilterer: StabilityFilterer{contract: contract}}, nil
}

// Stability is an auto generated Go binding around a Kowala contract.
type Stability struct {
	StabilityCaller     // Read-only binding to the contract
	StabilityTransactor // Write-only binding to the contract
	StabilityFilterer   // Log filterer for contract events
}

// StabilityCaller is an auto generated read-only Go binding around a Kowala contract.
type StabilityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilityTransactor is an auto generated write-only Go binding around a Kowala contract.
type StabilityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilityFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type StabilityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilitySession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type StabilitySession struct {
	Contract     *Stability        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StabilityCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type StabilityCallerSession struct {
	Contract *StabilityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StabilityTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type StabilityTransactorSession struct {
	Contract     *StabilityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StabilityRaw is an auto generated low-level Go binding around a Kowala contract.
type StabilityRaw struct {
	Contract *Stability // Generic contract binding to access the raw methods on
}

// StabilityCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type StabilityCallerRaw struct {
	Contract *StabilityCaller // Generic read-only contract binding to access the raw methods on
}

// StabilityTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type StabilityTransactorRaw struct {
	Contract *StabilityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStability creates a new instance of Stability, bound to a specific deployed contract.
func NewStability(address common.Address, backend bind.ContractBackend) (*Stability, error) {
	contract, err := bindStability(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stability{StabilityCaller: StabilityCaller{contract: contract}, StabilityTransactor: StabilityTransactor{contract: contract}, StabilityFilterer: StabilityFilterer{contract: contract}}, nil
}

// NewStabilityCaller creates a new read-only instance of Stability, bound to a specific deployed contract.
func NewStabilityCaller(address common.Address, caller bind.ContractCaller) (*StabilityCaller, error) {
	contract, err := bindStability(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StabilityCaller{contract: contract}, nil
}

// NewStabilityTransactor creates a new write-only instance of Stability, bound to a specific deployed contract.
func NewStabilityTransactor(address common.Address, transactor bind.ContractTransactor) (*StabilityTransactor, error) {
	contract, err := bindStability(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StabilityTransactor{contract: contract}, nil
}

// NewStabilityFilterer creates a new log filterer instance of Stability, bound to a specific deployed contract.
func NewStabilityFilterer(address common.Address, filterer bind.ContractFilterer) (*StabilityFilterer, error) {
	contract, err := bindStability(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StabilityFilterer{contract: contract}, nil
}

// bindStability binds a generic wrapper to an already deployed contract.
func bindStability(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StabilityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stability *StabilityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stability.Contract.StabilityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stability *StabilityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.Contract.StabilityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stability *StabilityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stability.Contract.StabilityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stability *StabilityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stability.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stability *StabilityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stability *StabilityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stability.Contract.contract.Transact(opts, method, params...)
}

// GetSubscriptionAtIndex is a free data retrieval call binding the contract method 0x40230046.
//
// Solidity: function getSubscriptionAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_Stability *StabilityCaller) GetSubscriptionAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _Stability.contract.Call(opts, out, "getSubscriptionAtIndex", index)
	return *ret, err
}

// GetSubscriptionAtIndex is a free data retrieval call binding the contract method 0x40230046.
//
// Solidity: function getSubscriptionAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_Stability *StabilitySession) GetSubscriptionAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _Stability.Contract.GetSubscriptionAtIndex(&_Stability.CallOpts, index)
}

// GetSubscriptionAtIndex is a free data retrieval call binding the contract method 0x40230046.
//
// Solidity: function getSubscriptionAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_Stability *StabilityCallerSession) GetSubscriptionAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _Stability.Contract.GetSubscriptionAtIndex(&_Stability.CallOpts, index)
}

// GetSubscriptionCount is a free data retrieval call binding the contract method 0x66419970.
//
// Solidity: function getSubscriptionCount() constant returns(count uint256)
func (_Stability *StabilityCaller) GetSubscriptionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "getSubscriptionCount")
	return *ret0, err
}

// GetSubscriptionCount is a free data retrieval call binding the contract method 0x66419970.
//
// Solidity: function getSubscriptionCount() constant returns(count uint256)
func (_Stability *StabilitySession) GetSubscriptionCount() (*big.Int, error) {
	return _Stability.Contract.GetSubscriptionCount(&_Stability.CallOpts)
}

// GetSubscriptionCount is a free data retrieval call binding the contract method 0x66419970.
//
// Solidity: function getSubscriptionCount() constant returns(count uint256)
func (_Stability *StabilityCallerSession) GetSubscriptionCount() (*big.Int, error) {
	return _Stability.Contract.GetSubscriptionCount(&_Stability.CallOpts)
}

// InitialReward is a free data retrieval call binding the contract method 0xabee967c.
//
// Solidity: function initialReward() constant returns(uint256)
func (_Stability *StabilityCaller) InitialReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "initialReward")
	return *ret0, err
}

// InitialReward is a free data retrieval call binding the contract method 0xabee967c.
//
// Solidity: function initialReward() constant returns(uint256)
func (_Stability *StabilitySession) InitialReward() (*big.Int, error) {
	return _Stability.Contract.InitialReward(&_Stability.CallOpts)
}

// InitialReward is a free data retrieval call binding the contract method 0xabee967c.
//
// Solidity: function initialReward() constant returns(uint256)
func (_Stability *StabilityCallerSession) InitialReward() (*big.Int, error) {
	return _Stability.Contract.InitialReward(&_Stability.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Stability *StabilityCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Stability *StabilitySession) Initialized() (bool, error) {
	return _Stability.Contract.Initialized(&_Stability.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Stability *StabilityCallerSession) Initialized() (bool, error) {
	return _Stability.Contract.Initialized(&_Stability.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_Stability *StabilityCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "minDeposit")
	return *ret0, err
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_Stability *StabilitySession) MinDeposit() (*big.Int, error) {
	return _Stability.Contract.MinDeposit(&_Stability.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_Stability *StabilityCallerSession) MinDeposit() (*big.Int, error) {
	return _Stability.Contract.MinDeposit(&_Stability.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stability *StabilityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stability *StabilitySession) Owner() (common.Address, error) {
	return _Stability.Contract.Owner(&_Stability.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stability *StabilityCallerSession) Owner() (common.Address, error) {
	return _Stability.Contract.Owner(&_Stability.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Stability *StabilityCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stability.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Stability *StabilitySession) Paused() (bool, error) {
	return _Stability.Contract.Paused(&_Stability.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Stability *StabilityCallerSession) Paused() (bool, error) {
	return _Stability.Contract.Paused(&_Stability.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(_minDeposit uint256, _priceProviderAddr address) returns()
func (_Stability *StabilityTransactor) Initialize(opts *bind.TransactOpts, _minDeposit *big.Int, _priceProviderAddr common.Address) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "initialize", _minDeposit, _priceProviderAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(_minDeposit uint256, _priceProviderAddr address) returns()
func (_Stability *StabilitySession) Initialize(_minDeposit *big.Int, _priceProviderAddr common.Address) (*types.Transaction, error) {
	return _Stability.Contract.Initialize(&_Stability.TransactOpts, _minDeposit, _priceProviderAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(_minDeposit uint256, _priceProviderAddr address) returns()
func (_Stability *StabilityTransactorSession) Initialize(_minDeposit *big.Int, _priceProviderAddr common.Address) (*types.Transaction, error) {
	return _Stability.Contract.Initialize(&_Stability.TransactOpts, _minDeposit, _priceProviderAddr)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stability *StabilityTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stability *StabilitySession) Pause() (*types.Transaction, error) {
	return _Stability.Contract.Pause(&_Stability.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stability *StabilityTransactorSession) Pause() (*types.Transaction, error) {
	return _Stability.Contract.Pause(&_Stability.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stability *StabilityTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stability *StabilitySession) RenounceOwnership() (*types.Transaction, error) {
	return _Stability.Contract.RenounceOwnership(&_Stability.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stability *StabilityTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stability.Contract.RenounceOwnership(&_Stability.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Stability *StabilityTransactor) Subscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "subscribe")
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Stability *StabilitySession) Subscribe() (*types.Transaction, error) {
	return _Stability.Contract.Subscribe(&_Stability.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Stability *StabilityTransactorSession) Subscribe() (*types.Transaction, error) {
	return _Stability.Contract.Subscribe(&_Stability.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Stability *StabilityTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Stability *StabilitySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Stability.Contract.TransferOwnership(&_Stability.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Stability *StabilityTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Stability.Contract.TransferOwnership(&_Stability.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Stability *StabilityTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Stability *StabilitySession) Unpause() (*types.Transaction, error) {
	return _Stability.Contract.Unpause(&_Stability.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Stability *StabilityTransactorSession) Unpause() (*types.Transaction, error) {
	return _Stability.Contract.Unpause(&_Stability.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Stability *StabilityTransactor) Unsubscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stability.contract.Transact(opts, "unsubscribe")
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Stability *StabilitySession) Unsubscribe() (*types.Transaction, error) {
	return _Stability.Contract.Unsubscribe(&_Stability.TransactOpts)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0xfcae4484.
//
// Solidity: function unsubscribe() returns()
func (_Stability *StabilityTransactorSession) Unsubscribe() (*types.Transaction, error) {
	return _Stability.Contract.Unsubscribe(&_Stability.TransactOpts)
}

// StabilityOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Stability contract.
type StabilityOwnershipRenouncedIterator struct {
	Event *StabilityOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *StabilityOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityOwnershipRenounced)
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
		it.Event = new(StabilityOwnershipRenounced)
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
func (it *StabilityOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityOwnershipRenounced represents a OwnershipRenounced event raised by the Stability contract.
type StabilityOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Stability *StabilityFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*StabilityOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Stability.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StabilityOwnershipRenouncedIterator{contract: _Stability.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Stability *StabilityFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *StabilityOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Stability.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityOwnershipRenounced)
				if err := _Stability.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// StabilityOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stability contract.
type StabilityOwnershipTransferredIterator struct {
	Event *StabilityOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StabilityOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityOwnershipTransferred)
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
		it.Event = new(StabilityOwnershipTransferred)
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
func (it *StabilityOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityOwnershipTransferred represents a OwnershipTransferred event raised by the Stability contract.
type StabilityOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Stability *StabilityFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StabilityOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stability.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StabilityOwnershipTransferredIterator{contract: _Stability.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Stability *StabilityFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StabilityOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stability.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityOwnershipTransferred)
				if err := _Stability.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// StabilityPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Stability contract.
type StabilityPauseIterator struct {
	Event *StabilityPause // Event containing the contract specifics and raw log

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
func (it *StabilityPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityPause)
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
		it.Event = new(StabilityPause)
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
func (it *StabilityPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityPause represents a Pause event raised by the Stability contract.
type StabilityPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Stability *StabilityFilterer) FilterPause(opts *bind.FilterOpts) (*StabilityPauseIterator, error) {

	logs, sub, err := _Stability.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &StabilityPauseIterator{contract: _Stability.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Stability *StabilityFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *StabilityPause) (event.Subscription, error) {

	logs, sub, err := _Stability.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityPause)
				if err := _Stability.contract.UnpackLog(event, "Pause", log); err != nil {
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

// StabilityUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Stability contract.
type StabilityUnpauseIterator struct {
	Event *StabilityUnpause // Event containing the contract specifics and raw log

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
func (it *StabilityUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilityUnpause)
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
		it.Event = new(StabilityUnpause)
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
func (it *StabilityUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilityUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilityUnpause represents a Unpause event raised by the Stability contract.
type StabilityUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Stability *StabilityFilterer) FilterUnpause(opts *bind.FilterOpts) (*StabilityUnpauseIterator, error) {

	logs, sub, err := _Stability.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &StabilityUnpauseIterator{contract: _Stability.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Stability *StabilityFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *StabilityUnpause) (event.Subscription, error) {

	logs, sub, err := _Stability.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilityUnpause)
				if err := _Stability.contract.UnpackLog(event, "Unpause", log); err != nil {
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
