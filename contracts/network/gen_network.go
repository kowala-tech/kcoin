// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// NetworkContractABI is the input ABI used to generate the binding from.
const NetworkContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDepositLowerBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"setMinDepositLowerBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsUpperBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availability\",\"outputs\":[{\"name\":\"available\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMinDepositUpperBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorsLowerBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorsUpperBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsLowerBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDepositUpperBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"},{\"name\":\"_maxValidators\",\"type\":\"uint256\"},{\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// NetworkContractBin is the compiled bytecode used for deploying new contracts.
const NetworkContractBin = `6060604052341561000f57600080fd5b604051608080610b8583398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600181905550600260015481151561009b57fe5b046003819055506002600154026002819055508160088190555060026008548115156100c357fe5b04600a8190555060026008540260098190555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600b8190555050505050610a54806101316000396000f300606060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806308ac52561461011757806341b3d1851461014057806355abf098146101695780636cf6d675146101925780636df01b6e146101bb5780637071688a146101de57806376792094146102075780638fcc9cfb1461022a5780639bb2ea5a1461024d578063a7f0b3de14610270578063b774cb1e146102c5578063c9b53900146102f6578063cefddda91461031f578063e188f27614610370578063e4f4410b14610393578063e7277b12146103bc578063e99cc696146103e5578063ebfa771614610408578063f2fde38b14610431578063facd743b1461046a575b600080fd5b341561012257600080fd5b61012a6104bb565b6040518082815260200191505060405180910390f35b341561014b57600080fd5b6101536104c1565b6040518082815260200191505060405180910390f35b341561017457600080fd5b61017c6104c7565b6040518082815260200191505060405180910390f35b341561019d57600080fd5b6101a56104cd565b6040518082815260200191505060405180910390f35b34156101c657600080fd5b6101dc60048080359060200190919050506104d3565b005b34156101e957600080fd5b6101f1610549565b6040518082815260200191505060405180910390f35b341561021257600080fd5b6102286004808035906020019091905050610556565b005b341561023557600080fd5b61024b60048080359060200190919050506105cc565b005b341561025857600080fd5b61026e6004808035906020019091905050610650565b005b341561027b57600080fd5b6102836106e4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102d057600080fd5b6102d861070a565b60405180826000191660001916815260200191505060405180910390f35b341561030157600080fd5b610309610710565b6040518082815260200191505060405180910390f35b341561032a57600080fd5b610356600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610721565b604051808215151515815260200191505060405180910390f35b341561037b57600080fd5b610391600480803590602001909190505061077b565b005b341561039e57600080fd5b6103a66107f1565b6040518082815260200191505060405180910390f35b34156103c757600080fd5b6103cf6107f7565b6040518082815260200191505060405180910390f35b34156103f057600080fd5b61040660048080359060200190919050506107fd565b005b341561041357600080fd5b61041b610873565b6040518082815260200191505060405180910390f35b341561043c57600080fd5b610468600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610879565b005b341561047557600080fd5b6104a1600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506109cf565b604051808215151515815260200191505060405180910390f35b60085481565b60015481565b60035481565b600b5481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052e57600080fd5b600254811115151561053f57600080fd5b8060038190555050565b6000600680549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105b157600080fd5b600a5481101515156105c257600080fd5b8060098190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561062757600080fd5b600354811015801561063b57506002548111155b151561064657600080fd5b8060018190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106ab57600080fd5b600a5481101580156106bf57506009548111155b15156106ca57600080fd5b6008548111156106e057806008819055506106e1565b5b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b600060068054905060085403905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107d657600080fd5b60035481101515156107e757600080fd5b8060028190555050565b600a5481565b60095481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561085857600080fd5b600954811115151561086957600080fd5b80600a8190555050565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108d457600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff1690509190505600a165627a7a7230582028f19b516a6189846615f2f295ea68df150abd35a06723b9ab913dcb1cd755ad0029`

// DeployNetworkContract deploys a new Ethereum contract, binding an instance of NetworkContract to it.
func DeployNetworkContract(auth *bind.TransactOpts, backend bind.ContractBackend, _minDeposit *big.Int, _genesis common.Address, _maxValidators *big.Int, _unbondingPeriod *big.Int) (common.Address, *types.Transaction, *NetworkContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkContractBin), backend, _minDeposit, _genesis, _maxValidators, _unbondingPeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NetworkContract{NetworkContractCaller: NetworkContractCaller{contract: contract}, NetworkContractTransactor: NetworkContractTransactor{contract: contract}}, nil
}

// NetworkContract is an auto generated Go binding around an Ethereum contract.
type NetworkContract struct {
	NetworkContractCaller     // Read-only binding to the contract
	NetworkContractTransactor // Write-only binding to the contract
}

// NetworkContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkContractSession struct {
	Contract     *NetworkContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkContractCallerSession struct {
	Contract *NetworkContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkContractTransactorSession struct {
	Contract     *NetworkContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkContractRaw struct {
	Contract *NetworkContract // Generic contract binding to access the raw methods on
}

// NetworkContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkContractCallerRaw struct {
	Contract *NetworkContractCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkContractTransactorRaw struct {
	Contract *NetworkContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkContract creates a new instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContract(address common.Address, backend bind.ContractBackend) (*NetworkContract, error) {
	contract, err := bindNetworkContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkContract{NetworkContractCaller: NetworkContractCaller{contract: contract}, NetworkContractTransactor: NetworkContractTransactor{contract: contract}}, nil
}

// NewNetworkContractCaller creates a new read-only instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContractCaller(address common.Address, caller bind.ContractCaller) (*NetworkContractCaller, error) {
	contract, err := bindNetworkContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkContractCaller{contract: contract}, nil
}

// NewNetworkContractTransactor creates a new write-only instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkContractTransactor, error) {
	contract, err := bindNetworkContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NetworkContractTransactor{contract: contract}, nil
}

// bindNetworkContract binds a generic wrapper to an already deployed contract.
func bindNetworkContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContract *NetworkContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkContract.Contract.NetworkContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContract *NetworkContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.Contract.NetworkContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContract *NetworkContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContract.Contract.NetworkContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContract *NetworkContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContract *NetworkContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContract *NetworkContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContract.Contract.contract.Transact(opts, method, params...)
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available uint256)
func (_NetworkContract *NetworkContractCaller) Availability(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "availability")
	return *ret0, err
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available uint256)
func (_NetworkContract *NetworkContractSession) Availability() (*big.Int, error) {
	return _NetworkContract.Contract.Availability(&_NetworkContract.CallOpts)
}

// Availability is a free data retrieval call binding the contract method 0xc9b53900.
//
// Solidity: function availability() constant returns(available uint256)
func (_NetworkContract *NetworkContractCallerSession) Availability() (*big.Int, error) {
	return _NetworkContract.Contract.Availability(&_NetworkContract.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_NetworkContract *NetworkContractCaller) Genesis(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "genesis")
	return *ret0, err
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_NetworkContract *NetworkContractSession) Genesis() (common.Address, error) {
	return _NetworkContract.Contract.Genesis(&_NetworkContract.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_NetworkContract *NetworkContractCallerSession) Genesis() (common.Address, error) {
	return _NetworkContract.Contract.Genesis(&_NetworkContract.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "getValidatorCount")
	return *ret0, err
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractSession) GetValidatorCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetValidatorCount(&_NetworkContract.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCallerSession) GetValidatorCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetValidatorCount(&_NetworkContract.CallOpts)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(account address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsGenesisValidator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isGenesisValidator", account)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(account address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsGenesisValidator(account common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisValidator(&_NetworkContract.CallOpts, account)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(account address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsGenesisValidator(account common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisValidator(&_NetworkContract.CallOpts, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsValidator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isValidator", addr)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsValidator(addr common.Address) (bool, error) {
	return _NetworkContract.Contract.IsValidator(&_NetworkContract.CallOpts, addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(addr address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsValidator(addr common.Address) (bool, error) {
	return _NetworkContract.Contract.IsValidator(&_NetworkContract.CallOpts, addr)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MaxValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "maxValidators")
	return *ret0, err
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MaxValidators() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidators(&_NetworkContract.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MaxValidators() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidators(&_NetworkContract.CallOpts)
}

// MaxValidatorsLowerBound is a free data retrieval call binding the contract method 0xe4f4410b.
//
// Solidity: function maxValidatorsLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MaxValidatorsLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "maxValidatorsLowerBound")
	return *ret0, err
}

// MaxValidatorsLowerBound is a free data retrieval call binding the contract method 0xe4f4410b.
//
// Solidity: function maxValidatorsLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MaxValidatorsLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsLowerBound(&_NetworkContract.CallOpts)
}

// MaxValidatorsLowerBound is a free data retrieval call binding the contract method 0xe4f4410b.
//
// Solidity: function maxValidatorsLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MaxValidatorsLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsLowerBound(&_NetworkContract.CallOpts)
}

// MaxValidatorsUpperBound is a free data retrieval call binding the contract method 0xe7277b12.
//
// Solidity: function maxValidatorsUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MaxValidatorsUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "maxValidatorsUpperBound")
	return *ret0, err
}

// MaxValidatorsUpperBound is a free data retrieval call binding the contract method 0xe7277b12.
//
// Solidity: function maxValidatorsUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MaxValidatorsUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsUpperBound(&_NetworkContract.CallOpts)
}

// MaxValidatorsUpperBound is a free data retrieval call binding the contract method 0xe7277b12.
//
// Solidity: function maxValidatorsUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MaxValidatorsUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MaxValidatorsUpperBound(&_NetworkContract.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "minDeposit")
	return *ret0, err
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MinDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.MinDeposit(&_NetworkContract.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MinDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.MinDeposit(&_NetworkContract.CallOpts)
}

// MinDepositLowerBound is a free data retrieval call binding the contract method 0x55abf098.
//
// Solidity: function minDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MinDepositLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "minDepositLowerBound")
	return *ret0, err
}

// MinDepositLowerBound is a free data retrieval call binding the contract method 0x55abf098.
//
// Solidity: function minDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MinDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositLowerBound(&_NetworkContract.CallOpts)
}

// MinDepositLowerBound is a free data retrieval call binding the contract method 0x55abf098.
//
// Solidity: function minDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MinDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositLowerBound(&_NetworkContract.CallOpts)
}

// MinDepositUpperBound is a free data retrieval call binding the contract method 0xebfa7716.
//
// Solidity: function minDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) MinDepositUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "minDepositUpperBound")
	return *ret0, err
}

// MinDepositUpperBound is a free data retrieval call binding the contract method 0xebfa7716.
//
// Solidity: function minDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) MinDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositUpperBound(&_NetworkContract.CallOpts)
}

// MinDepositUpperBound is a free data retrieval call binding the contract method 0xebfa7716.
//
// Solidity: function minDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) MinDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.MinDepositUpperBound(&_NetworkContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) UnbondingPeriod() (*big.Int, error) {
	return _NetworkContract.Contract.UnbondingPeriod(&_NetworkContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _NetworkContract.Contract.UnbondingPeriod(&_NetworkContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_NetworkContract *NetworkContractCaller) ValidatorsChecksum(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "validatorsChecksum")
	return *ret0, err
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_NetworkContract *NetworkContractSession) ValidatorsChecksum() ([32]byte, error) {
	return _NetworkContract.Contract.ValidatorsChecksum(&_NetworkContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_NetworkContract *NetworkContractCallerSession) ValidatorsChecksum() ([32]byte, error) {
	return _NetworkContract.Contract.ValidatorsChecksum(&_NetworkContract.CallOpts)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMaxValidators(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMaxValidators", max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidators(&_NetworkContract.TransactOpts, max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidators(&_NetworkContract.TransactOpts, max)
}

// SetMaxValidatorsLowerBound is a paid mutator transaction binding the contract method 0xe99cc696.
//
// Solidity: function setMaxValidatorsLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMaxValidatorsLowerBound(opts *bind.TransactOpts, min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMaxValidatorsLowerBound", min)
}

// SetMaxValidatorsLowerBound is a paid mutator transaction binding the contract method 0xe99cc696.
//
// Solidity: function setMaxValidatorsLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMaxValidatorsLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMaxValidatorsLowerBound is a paid mutator transaction binding the contract method 0xe99cc696.
//
// Solidity: function setMaxValidatorsLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMaxValidatorsLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMaxValidatorsUpperBound is a paid mutator transaction binding the contract method 0x76792094.
//
// Solidity: function setMaxValidatorsUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMaxValidatorsUpperBound(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMaxValidatorsUpperBound", max)
}

// SetMaxValidatorsUpperBound is a paid mutator transaction binding the contract method 0x76792094.
//
// Solidity: function setMaxValidatorsUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMaxValidatorsUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetMaxValidatorsUpperBound is a paid mutator transaction binding the contract method 0x76792094.
//
// Solidity: function setMaxValidatorsUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMaxValidatorsUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMaxValidatorsUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMinDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMinDeposit", deposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMinDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMinDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetMinDepositLowerBound is a paid mutator transaction binding the contract method 0x6df01b6e.
//
// Solidity: function setMinDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMinDepositLowerBound(opts *bind.TransactOpts, min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMinDepositLowerBound", min)
}

// SetMinDepositLowerBound is a paid mutator transaction binding the contract method 0x6df01b6e.
//
// Solidity: function setMinDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMinDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMinDepositLowerBound is a paid mutator transaction binding the contract method 0x6df01b6e.
//
// Solidity: function setMinDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMinDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetMinDepositUpperBound is a paid mutator transaction binding the contract method 0xe188f276.
//
// Solidity: function setMinDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetMinDepositUpperBound(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMinDepositUpperBound", max)
}

// SetMinDepositUpperBound is a paid mutator transaction binding the contract method 0xe188f276.
//
// Solidity: function setMinDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetMinDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetMinDepositUpperBound is a paid mutator transaction binding the contract method 0xe188f276.
//
// Solidity: function setMinDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMinDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMinDepositUpperBound(&_NetworkContract.TransactOpts, max)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContract *NetworkContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContract *NetworkContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.TransferOwnership(&_NetworkContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_NetworkContract *NetworkContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.TransferOwnership(&_NetworkContract.TransactOpts, addr)
}
