// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// ElectionContractABI is the input ABI used to generate the binding from.
const ElectionContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"releasedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxValidators\",\"type\":\"uint256\"},{\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// ElectionContractBin is the compiled bytecode used for deploying new contracts.
const ElectionContractBin = `60606040526040516080806200198083398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508334101515156200008c57600080fd5b600183101515156200009d57600080fd5b836001819055508260028190555062015180820260038190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001188134620001226401000000000262000d81176401000000009004565b50505050620005a2565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078054806001018281620001819190620004e4565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff0219169083151502179055508360020180548060010182816200020d919062000513565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b60008311156200043957600660006007600186038154811015156200027857fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816002016001836002018054905003815481101515620002fc57fe5b90600052602060002090600202019050806000015485111515620003205762000439565b6007600184038154811015156200033357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007848154811015156200036f57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600760018503815481101515620003cc57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082826000018190555060018303846000018190555082806001900393505062000257565b620004576200045f640100000000026200120f176401000000009004565b505050505050565b6007604051808280548015620004cb57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162000480575b5050915050604051809103902060058160001916905550565b8154818355818115116200050e578183600052602060002091820191016200050d919062000548565b5b505050565b815481835581811511620005435760020281600202836000526020600020918201910162000542919062000570565b5b505050565b6200056d91905b80821115620005695760008160009055506001016200054f565b5090565b90565b6200059f91905b808211156200059b5760008082016000905560018201600090555060020162000577565b5090565b90565b6113ce80620005b26000396000f300606060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461011757806308ac525614610140578063216355ec146101695780633ed0a3731461017e57806369474625146101bc5780636cf6d675146101e55780637071688a1461020e578063893d20e8146102375780639363a1411461028c57806397584b3e146102b55780639bb2ea5a146102e2578063a7f0b3de14610305578063b688a3631461035a578063b774cb1e14610364578063c22a933c14610395578063cefddda9146103b8578063d66d9e1914610409578063e7a60a9c1461041e578063f2fde38b14610488578063facd743b146104d9575b600080fd5b341561012257600080fd5b61012a61052a565b6040518082815260200191505060405180910390f35b341561014b57600080fd5b6101536105fe565b6040518082815260200191505060405180910390f35b341561017457600080fd5b61017c610604565b005b341561018957600080fd5b61019f6004808035906020019091905050610743565b604051808381526020018281526020019250505060405180910390f35b34156101c757600080fd5b6101cf6107bb565b6040518082815260200191505060405180910390f35b34156101f057600080fd5b6101f86107c1565b6040518082815260200191505060405180910390f35b341561021957600080fd5b6102216107c7565b6040518082815260200191505060405180910390f35b341561024257600080fd5b61024a6107d4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561029757600080fd5b61029f6107fd565b6040518082815260200191505060405180910390f35b34156102c057600080fd5b6102c861084a565b604051808215151515815260200191505060405180910390f35b34156102ed57600080fd5b610303600480803590602001909190505061085d565b005b341561031057600080fd5b610318610901565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610362610927565b005b341561036f57600080fd5b610377610975565b60405180826000191660001916815260200191505060405180910390f35b34156103a057600080fd5b6103b6600480803590602001909190505061097b565b005b34156103c357600080fd5b6103ef600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506109e0565b604051808215151515815260200191505060405180910390f35b341561041457600080fd5b61041c610a3a565b005b341561042957600080fd5b61043f6004808035906020019091905050610a59565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b341561049357600080fd5b6104bf600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b11565b604051808215151515815260200191505060405180910390f35b34156104e457600080fd5b610510600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610bee565b604051808215151515815260200191505060405180910390f35b60008061053561084a565b156105445760015491506105fa565b60066000600760016007805490500381548110151561055f57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160018360020180549050038154811015156105e457fe5b9060005260206000209060020201600001540191505b5090565b60025481565b600080600080925060009150600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b80805490508210801561068857506000818381548110151561067357fe5b90600052602060002090600202016001015414155b156106ea57808281548110151561069b57fe5b9060005260206000209060020201600101544210156106b9576106ea565b80828154811015156106c757fe5b906000526020600020906002020160000154830192508180600101925050610655565b6106f43383610c47565b600083111561073e573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561073d57600080fd5b5b505050565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018481548110151561079757fe5b90600052602060002090600202019050806000015481600101549250925050915091565b60015481565b60035481565b6000600780549050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506002540311905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108bb57600080fd5b6007805490508310156108f55782600780549050039150600090505b818110156108f4576108e7610d34565b80806001019150506108d7565b5b82600281905550505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61093033610bee565b15151561093c57600080fd5b61094461052a565b341015151561095257600080fd5b61095a61084a565b151561096957610968610d34565b5b6109733334610d81565b565b60055481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109d657600080fd5b8060018190555050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b610a4333610bee565b1515610a4e57600080fd5b610a573361109b565b565b6000806000600784815481101515610a6d57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806002016001826002018054905003815481101515610af757fe5b906000526020600020906002020160000154915050915091565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b6e57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515610be557816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60019050919050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b600080600080841415610c5957610d2d565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b8260020180549050811015610d1b578260020181815481101515610cc257fe5b90600052602060002090600202018360020183815481101515610ce157fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050610ca2565b818360020181610d2b9190611292565b505b5050505050565b610d7f6007600160078054905003815481101515610d4e57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661109b565b565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078054806001018281610dde91906112c4565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff021916908315150217905550836002018054806001018281610e6891906112f0565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b600083111561108b5760066000600760018603815481101515610ed157fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816002016001836002018054905003815481101515610f5457fe5b90600052602060002090600202019050806000015485111515610f765761108b565b600760018403815481101515610f8857fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600784815481101515610fc357fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560076001850381548110151561101f57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050610eb2565b61109361120f565b505050505050565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b60016007805490500381101561119a5760076001820181548110151561110957fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078281548110151561114457fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806001019150506110e7565b60078054809190600190036111af9190611322565b5060008260010160006101000a81548160ff02191690831515021790555060035442018260020160018460020180549050038154811015156111ed57fe5b90600052602060002090600202016001018190555061120a61120f565b505050565b600760405180828054801561127957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161122f575b5050915050604051809103902060058160001916905550565b8154818355818115116112bf576002028160020283600052602060002091820191016112be919061134e565b5b505050565b8154818355818115116112eb578183600052602060002091820191016112ea919061137d565b5b505050565b81548183558181151161131d5760020281600202836000526020600020918201910161131c919061134e565b5b505050565b81548183558181151161134957818360005260206000209182019101611348919061137d565b5b505050565b61137a91905b8082111561137657600080820160009055600182016000905550600201611354565b5090565b90565b61139f91905b8082111561139b576000816000905550600101611383565b5090565b905600a165627a7a72305820517c708bb0e1efd6643434ebe4e5f9d423db745d9a63410aa2bcc47ea1f6074d0029`

// DeployElectionContract deploys a new Ethereum contract, binding an instance of ElectionContract to it.
func DeployElectionContract(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxValidators *big.Int, _unbondingPeriod *big.Int, _genesis common.Address) (common.Address, *types.Transaction, *ElectionContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElectionContractBin), backend, _baseDeposit, _maxValidators, _unbondingPeriod, _genesis)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElectionContract{ElectionContractCaller: ElectionContractCaller{contract: contract}, ElectionContractTransactor: ElectionContractTransactor{contract: contract}}, nil
}

// ElectionContract is an auto generated Go binding around an Ethereum contract.
type ElectionContract struct {
	ElectionContractCaller     // Read-only binding to the contract
	ElectionContractTransactor // Write-only binding to the contract
}

// ElectionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionContractSession struct {
	Contract     *ElectionContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionContractCallerSession struct {
	Contract *ElectionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ElectionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionContractTransactorSession struct {
	Contract     *ElectionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ElectionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionContractRaw struct {
	Contract *ElectionContract // Generic contract binding to access the raw methods on
}

// ElectionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionContractCallerRaw struct {
	Contract *ElectionContractCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionContractTransactorRaw struct {
	Contract *ElectionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElectionContract creates a new instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContract(address common.Address, backend bind.ContractBackend) (*ElectionContract, error) {
	contract, err := bindElectionContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElectionContract{ElectionContractCaller: ElectionContractCaller{contract: contract}, ElectionContractTransactor: ElectionContractTransactor{contract: contract}}, nil
}

// NewElectionContractCaller creates a new read-only instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContractCaller(address common.Address, caller bind.ContractCaller) (*ElectionContractCaller, error) {
	contract, err := bindElectionContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionContractCaller{contract: contract}, nil
}

// NewElectionContractTransactor creates a new write-only instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionContractTransactor, error) {
	contract, err := bindElectionContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ElectionContractTransactor{contract: contract}, nil
}

// bindElectionContract binds a generic wrapper to an already deployed contract.
func bindElectionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionContract *ElectionContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElectionContract.Contract.ElectionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionContract *ElectionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.Contract.ElectionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionContract *ElectionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionContract.Contract.ElectionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionContract *ElectionContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElectionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionContract *ElectionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionContract *ElectionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionContract.Contract.contract.Transact(opts, method, params...)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ElectionContract *ElectionContractCaller) _hasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ElectionContract *ElectionContractSession) _hasAvailability() (bool, error) {
	return _ElectionContract.Contract._hasAvailability(&_ElectionContract.CallOpts)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ElectionContract *ElectionContractCallerSession) _hasAvailability() (bool, error) {
	return _ElectionContract.Contract._hasAvailability(&_ElectionContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ElectionContract *ElectionContractCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ElectionContract *ElectionContractSession) BaseDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.BaseDeposit(&_ElectionContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) BaseDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.BaseDeposit(&_ElectionContract.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_ElectionContract *ElectionContractCaller) Genesis(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "genesis")
	return *ret0, err
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_ElectionContract *ElectionContractSession) Genesis() (common.Address, error) {
	return _ElectionContract.Contract.Genesis(&_ElectionContract.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0xa7f0b3de.
//
// Solidity: function genesis() constant returns(address)
func (_ElectionContract *ElectionContractCallerSession) Genesis() (common.Address, error) {
	return _ElectionContract.Contract.Genesis(&_ElectionContract.CallOpts)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_ElectionContract *ElectionContractCaller) GetDepositAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	ret := new(struct {
		Amount     *big.Int
		ReleasedAt *big.Int
	})
	out := ret
	err := _ElectionContract.contract.Call(opts, out, "getDepositAtIndex", index)
	return *ret, err
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_ElectionContract *ElectionContractSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	return _ElectionContract.Contract.GetDepositAtIndex(&_ElectionContract.CallOpts, index)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_ElectionContract *ElectionContractCallerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	return _ElectionContract.Contract.GetDepositAtIndex(&_ElectionContract.CallOpts, index)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCaller) GetDepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getDepositCount")
	return *ret0, err
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractSession) GetDepositCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetDepositCount(&_ElectionContract.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCallerSession) GetDepositCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetDepositCount(&_ElectionContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ElectionContract *ElectionContractCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ElectionContract *ElectionContractSession) GetMinimumDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.GetMinimumDeposit(&_ElectionContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ElectionContract *ElectionContractCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _ElectionContract.Contract.GetMinimumDeposit(&_ElectionContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ElectionContract *ElectionContractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ElectionContract *ElectionContractSession) GetOwner() (common.Address, error) {
	return _ElectionContract.Contract.GetOwner(&_ElectionContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_ElectionContract *ElectionContractCallerSession) GetOwner() (common.Address, error) {
	return _ElectionContract.Contract.GetOwner(&_ElectionContract.CallOpts)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ElectionContract *ElectionContractCaller) GetValidatorAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _ElectionContract.contract.Call(opts, out, "getValidatorAtIndex", index)
	return *ret, err
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ElectionContract *ElectionContractSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ElectionContract.Contract.GetValidatorAtIndex(&_ElectionContract.CallOpts, index)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ElectionContract *ElectionContractCallerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ElectionContract.Contract.GetValidatorAtIndex(&_ElectionContract.CallOpts, index)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "getValidatorCount")
	return *ret0, err
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractSession) GetValidatorCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetValidatorCount(&_ElectionContract.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ElectionContract *ElectionContractCallerSession) GetValidatorCount() (*big.Int, error) {
	return _ElectionContract.Contract.GetValidatorCount(&_ElectionContract.CallOpts)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCaller) IsGenesisValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "isGenesisValidator", code)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsGenesisValidator(&_ElectionContract.CallOpts, code)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCallerSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsGenesisValidator(&_ElectionContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCaller) IsValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "isValidator", code)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractSession) IsValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsValidator(&_ElectionContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCallerSession) IsValidator(code common.Address) (bool, error) {
	return _ElectionContract.Contract.IsValidator(&_ElectionContract.CallOpts, code)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ElectionContract *ElectionContractCaller) MaxValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "maxValidators")
	return *ret0, err
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ElectionContract *ElectionContractSession) MaxValidators() (*big.Int, error) {
	return _ElectionContract.Contract.MaxValidators(&_ElectionContract.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) MaxValidators() (*big.Int, error) {
	return _ElectionContract.Contract.MaxValidators(&_ElectionContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_ElectionContract *ElectionContractCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_ElectionContract *ElectionContractSession) UnbondingPeriod() (*big.Int, error) {
	return _ElectionContract.Contract.UnbondingPeriod(&_ElectionContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _ElectionContract.Contract.UnbondingPeriod(&_ElectionContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ElectionContract *ElectionContractCaller) ValidatorsChecksum(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "validatorsChecksum")
	return *ret0, err
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ElectionContract *ElectionContractSession) ValidatorsChecksum() ([32]byte, error) {
	return _ElectionContract.Contract.ValidatorsChecksum(&_ElectionContract.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ElectionContract *ElectionContractCallerSession) ValidatorsChecksum() ([32]byte, error) {
	return _ElectionContract.Contract.ValidatorsChecksum(&_ElectionContract.CallOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_ElectionContract *ElectionContractTransactor) Join(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "join")
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_ElectionContract *ElectionContractSession) Join() (*types.Transaction, error) {
	return _ElectionContract.Contract.Join(&_ElectionContract.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_ElectionContract *ElectionContractTransactorSession) Join() (*types.Transaction, error) {
	return _ElectionContract.Contract.Join(&_ElectionContract.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_ElectionContract *ElectionContractTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_ElectionContract *ElectionContractSession) Leave() (*types.Transaction, error) {
	return _ElectionContract.Contract.Leave(&_ElectionContract.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_ElectionContract *ElectionContractTransactorSession) Leave() (*types.Transaction, error) {
	return _ElectionContract.Contract.Leave(&_ElectionContract.TransactOpts)
}

// RedeemFunds is a paid mutator transaction binding the contract method 0x216355ec.
//
// Solidity: function redeemFunds() returns()
func (_ElectionContract *ElectionContractTransactor) RedeemFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "redeemFunds")
}

// RedeemFunds is a paid mutator transaction binding the contract method 0x216355ec.
//
// Solidity: function redeemFunds() returns()
func (_ElectionContract *ElectionContractSession) RedeemFunds() (*types.Transaction, error) {
	return _ElectionContract.Contract.RedeemFunds(&_ElectionContract.TransactOpts)
}

// RedeemFunds is a paid mutator transaction binding the contract method 0x216355ec.
//
// Solidity: function redeemFunds() returns()
func (_ElectionContract *ElectionContractTransactorSession) RedeemFunds() (*types.Transaction, error) {
	return _ElectionContract.Contract.RedeemFunds(&_ElectionContract.TransactOpts)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ElectionContract *ElectionContractTransactor) SetBaseDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "setBaseDeposit", deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ElectionContract *ElectionContractSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetBaseDeposit(&_ElectionContract.TransactOpts, deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ElectionContract *ElectionContractTransactorSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetBaseDeposit(&_ElectionContract.TransactOpts, deposit)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ElectionContract *ElectionContractTransactor) SetMaxValidators(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "setMaxValidators", max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ElectionContract *ElectionContractSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetMaxValidators(&_ElectionContract.TransactOpts, max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ElectionContract *ElectionContractTransactorSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.SetMaxValidators(&_ElectionContract.TransactOpts, max)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_ElectionContract *ElectionContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_ElectionContract *ElectionContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.TransferOwnership(&_ElectionContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_ElectionContract *ElectionContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.TransferOwnership(&_ElectionContract.TransactOpts, _newOwner)
}
