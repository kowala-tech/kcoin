// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
)

// ElectionContractABI is the input ABI used to generate the binding from.
const ElectionContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"availableAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"reportValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesisValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxValidators\",\"type\":\"uint256\"},{\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// ElectionContractBin is the compiled bytecode used for deploying new contracts.
const ElectionContractBin = `606060405234156200001057600080fd5b60405160808062001bba83398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183101515156200009957600080fd5b670de0b6b3a764000084026001819055508260028190555062015180820260038190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000120816001546200012a6401000000000262000fb3176401000000009004565b50505050620005aa565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078054806001018281620001899190620004ec565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff0219169083151502179055508360020180548060010182816200021591906200051b565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b60008311156200044157600660006007600186038154811015156200028057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002091508160020160018360020180549050038154811015156200030457fe5b90600052602060002090600202019050806000015485111515620003285762000441565b6007600184038154811015156200033b57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007848154811015156200037757fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600760018503815481101515620003d457fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508282600001819055506001830384600001819055508280600190039350506200025f565b6200045f620004676401000000000262001441176401000000009004565b505050505050565b6007604051808280548015620004d357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162000488575b5050915050604051809103902060058160001916905550565b815481835581811511620005165781836000526020600020918201910162000515919062000550565b5b505050565b8154818355818115116200054b576002028160020283600052602060002091820191016200054a919062000578565b5b505050565b6200057591905b808211156200057157600081600090555060010162000557565b5090565b90565b620005a791905b80821115620005a3576000808201600090556001820160009055506002016200057f565b5090565b90565b61160080620005ba6000396000f300606060405260043610610128576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461012a57806308ac5256146101535780633ed0a3731461017c5780634b2c89d5146101ba5780634ee2a190146101cf57806369474625146102085780636cf6d675146102315780637071688a1461025a578063893d20e8146102835780639363a141146102d857806397584b3e146103015780639bb2ea5a1461032e578063b688a36314610351578063b774cb1e1461035b578063c22a933c1461038c578063cefddda9146103af578063d66d9e1914610400578063e7a60a9c14610415578063f2fde38b1461047f578063f963aeea146104d0578063facd743b14610525578063fe575a8714610576575b005b341561013557600080fd5b61013d6105c7565b6040518082815260200191505060405180910390f35b341561015e57600080fd5b61016661069b565b6040518082815260200191505060405180910390f35b341561018757600080fd5b61019d60048080359060200190919050506106a1565b604051808381526020018281526020019250505060405180910390f35b34156101c557600080fd5b6101cd61072c565b005b34156101da57600080fd5b610206600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610881565b005b341561021357600080fd5b61021b6108a1565b6040518082815260200191505060405180910390f35b341561023c57600080fd5b6102446108a7565b6040518082815260200191505060405180910390f35b341561026557600080fd5b61026d6108ad565b6040518082815260200191505060405180910390f35b341561028e57600080fd5b6102966108ba565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102e357600080fd5b6102eb6108e3565b6040518082815260200191505060405180910390f35b341561030c57600080fd5b610314610930565b604051808215151515815260200191505060405180910390f35b341561033957600080fd5b61034f6004808035906020019091905050610943565b005b6103596109e7565b005b341561036657600080fd5b61036e610a4a565b60405180826000191660001916815260200191505060405180910390f35b341561039757600080fd5b6103ad6004808035906020019091905050610a50565b005b34156103ba57600080fd5b6103e6600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ab5565b604051808215151515815260200191505060405180910390f35b341561040b57600080fd5b610413610b0f565b005b341561042057600080fd5b6104366004808035906020019091905050610b2e565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b341561048a57600080fd5b6104b6600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610be6565b604051808215151515815260200191505060405180910390f35b34156104db57600080fd5b6104e3610cc3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561053057600080fd5b61055c600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ce9565b604051808215151515815260200191505060405180910390f35b341561058157600080fd5b6105ad600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610d42565b604051808215151515815260200191505060405180910390f35b6000806105d2610930565b156105e1576001549150610697565b6006600060076001600780549050038154811015156105fc57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561068157fe5b9060005260206000209060020201600001540191505b5090565b60025481565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156106f557fe5b90600052602060002090600202019050670de0b6b3a7640000816000015481151561071c57fe5b0481600101549250925050915091565b600080600061073a33610d42565b15151561074657600080fd5b6000925060009150600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b8080549050821080156107c65750600081838154811015156107b157fe5b90600052602060002090600202016001015414155b156108285780828154811015156107d957fe5b9060005260206000209060020201600101544210156107f757610828565b808281548110151561080557fe5b906000526020600020906002020160000154830192508180600101925050610793565b6108323383610d9b565b600083111561087c573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561087b57600080fd5b5b505050565b61088a33610ce9565b151561089557600080fd5b61089e81610e88565b50565b60015481565b60035481565b6000600780549050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506002540311905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109a157600080fd5b6007805490508310156109db5782600780549050039150600090505b818110156109da576109cd610f66565b80806001019150506109bd565b5b82600281905550505050565b6109f033610d42565b1515156109fc57600080fd5b610a0533610ce9565b151515610a1157600080fd5b610a196105c7565b3410151515610a2757600080fd5b610a2f610930565b1515610a3e57610a3d610f66565b5b610a483334610fb3565b565b60055481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610aab57600080fd5b8060018190555050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b610b1833610ce9565b1515610b2357600080fd5b610b2c336112cd565b565b6000806000600784815481101515610b4257fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806002016001826002018054905003815481101515610bcc57fe5b906000526020600020906002020160000154915050915091565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c4357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515610cba57816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60019050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff169050919050565b600080600080841415610dad57610e81565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b8260020180549050811015610e6f578260020181815481101515610e1657fe5b90600052602060002090600202018360020183815481101515610e3557fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050610df6565b818360020181610e7f91906114c4565b505b5050505050565b610e9181610ce9565b1515610e9c57600080fd5b610ea5816112cd565b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160016101000a81548160ff02191690831515021790555060088054806001018281610f1491906114f6565b9160005260206000209001600083909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b610fb16007600160078054905003815481101515610f8057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166112cd565b565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020935060016007805480600101828161101091906114f6565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff02191690831515021790555083600201805480600101828161109a9190611522565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b60008311156112bd576006600060076001860381548110151561110357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915081600201600183600201805490500381548110151561118657fe5b906000526020600020906002020190508060000154851115156111a8576112bd565b6007600184038154811015156111ba57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007848154811015156111f557fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560076001850381548110151561125157fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508282600001819055506001830384600001819055508280600190039350506110e4565b6112c5611441565b505050505050565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b6001600780549050038110156113cc5760076001820181548110151561133b57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078281548110151561137657fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050611319565b60078054809190600190036113e19190611554565b5060008260010160006101000a81548160ff021916908315150217905550600354420182600201600184600201805490500381548110151561141f57fe5b90600052602060002090600202016001018190555061143c611441565b505050565b60076040518082805480156114ab57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611461575b5050915050604051809103902060058160001916905550565b8154818355818115116114f1576002028160020283600052602060002091820191016114f09190611580565b5b505050565b81548183558181151161151d5781836000526020600020918201910161151c91906115af565b5b505050565b81548183558181151161154f5760020281600202836000526020600020918201910161154e9190611580565b5b505050565b81548183558181151161157b5781836000526020600020918201910161157a91906115af565b5b505050565b6115ac91905b808211156115a857600080820160009055600182016000905550600201611586565b5090565b90565b6115d191905b808211156115cd5760008160009055506001016115b5565b5090565b905600a165627a7a72305820057820f1b49914cdba897a23707d15ece27aff75e2530b56b259dfc8b9d1f4f00029`

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

// GenesisValidator is a free data retrieval call binding the contract method 0xf963aeea.
//
// Solidity: function genesisValidator() constant returns(address)
func (_ElectionContract *ElectionContractCaller) GenesisValidator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "genesisValidator")
	return *ret0, err
}

// GenesisValidator is a free data retrieval call binding the contract method 0xf963aeea.
//
// Solidity: function genesisValidator() constant returns(address)
func (_ElectionContract *ElectionContractSession) GenesisValidator() (common.Address, error) {
	return _ElectionContract.Contract.GenesisValidator(&_ElectionContract.CallOpts)
}

// GenesisValidator is a free data retrieval call binding the contract method 0xf963aeea.
//
// Solidity: function genesisValidator() constant returns(address)
func (_ElectionContract *ElectionContractCallerSession) GenesisValidator() (common.Address, error) {
	return _ElectionContract.Contract.GenesisValidator(&_ElectionContract.CallOpts)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ElectionContract *ElectionContractCaller) GetDepositAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	ret := new(struct {
		Amount      *big.Int
		AvailableAt *big.Int
	})
	out := ret
	err := _ElectionContract.contract.Call(opts, out, "getDepositAtIndex", index)
	return *ret, err
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ElectionContract *ElectionContractSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	return _ElectionContract.Contract.GetDepositAtIndex(&_ElectionContract.CallOpts, index)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ElectionContract *ElectionContractCallerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
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
// Solidity: function getValidatorAtIndex(index uint256) constant returns(identity address, deposit uint256)
func (_ElectionContract *ElectionContractCaller) GetValidatorAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Identity common.Address
	Deposit  *big.Int
}, error) {
	ret := new(struct {
		Identity common.Address
		Deposit  *big.Int
	})
	out := ret
	err := _ElectionContract.contract.Call(opts, out, "getValidatorAtIndex", index)
	return *ret, err
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(identity address, deposit uint256)
func (_ElectionContract *ElectionContractSession) GetValidatorAtIndex(index *big.Int) (struct {
	Identity common.Address
	Deposit  *big.Int
}, error) {
	return _ElectionContract.Contract.GetValidatorAtIndex(&_ElectionContract.CallOpts, index)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(identity address, deposit uint256)
func (_ElectionContract *ElectionContractCallerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Identity common.Address
	Deposit  *big.Int
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

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCaller) IsBlacklisted(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "isBlacklisted", identity)
	return *ret0, err
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractSession) IsBlacklisted(identity common.Address) (bool, error) {
	return _ElectionContract.Contract.IsBlacklisted(&_ElectionContract.CallOpts, identity)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCallerSession) IsBlacklisted(identity common.Address) (bool, error) {
	return _ElectionContract.Contract.IsBlacklisted(&_ElectionContract.CallOpts, identity)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCaller) IsGenesisValidator(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "isGenesisValidator", identity)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractSession) IsGenesisValidator(identity common.Address) (bool, error) {
	return _ElectionContract.Contract.IsGenesisValidator(&_ElectionContract.CallOpts, identity)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCallerSession) IsGenesisValidator(identity common.Address) (bool, error) {
	return _ElectionContract.Contract.IsGenesisValidator(&_ElectionContract.CallOpts, identity)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCaller) IsValidator(opts *bind.CallOpts, identity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElectionContract.contract.Call(opts, out, "isValidator", identity)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractSession) IsValidator(identity common.Address) (bool, error) {
	return _ElectionContract.Contract.IsValidator(&_ElectionContract.CallOpts, identity)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(identity address) constant returns(isIndeed bool)
func (_ElectionContract *ElectionContractCallerSession) IsValidator(identity common.Address) (bool, error) {
	return _ElectionContract.Contract.IsValidator(&_ElectionContract.CallOpts, identity)
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

// RedeemDeposits is a paid mutator transaction binding the contract method 0x4b2c89d5.
//
// Solidity: function redeemDeposits() returns()
func (_ElectionContract *ElectionContractTransactor) RedeemDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "redeemDeposits")
}

// RedeemDeposits is a paid mutator transaction binding the contract method 0x4b2c89d5.
//
// Solidity: function redeemDeposits() returns()
func (_ElectionContract *ElectionContractSession) RedeemDeposits() (*types.Transaction, error) {
	return _ElectionContract.Contract.RedeemDeposits(&_ElectionContract.TransactOpts)
}

// RedeemDeposits is a paid mutator transaction binding the contract method 0x4b2c89d5.
//
// Solidity: function redeemDeposits() returns()
func (_ElectionContract *ElectionContractTransactorSession) RedeemDeposits() (*types.Transaction, error) {
	return _ElectionContract.Contract.RedeemDeposits(&_ElectionContract.TransactOpts)
}

// ReportValidator is a paid mutator transaction binding the contract method 0x4ee2a190.
//
// Solidity: function reportValidator(identity address) returns()
func (_ElectionContract *ElectionContractTransactor) ReportValidator(opts *bind.TransactOpts, identity common.Address) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "reportValidator", identity)
}

// ReportValidator is a paid mutator transaction binding the contract method 0x4ee2a190.
//
// Solidity: function reportValidator(identity address) returns()
func (_ElectionContract *ElectionContractSession) ReportValidator(identity common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.ReportValidator(&_ElectionContract.TransactOpts, identity)
}

// ReportValidator is a paid mutator transaction binding the contract method 0x4ee2a190.
//
// Solidity: function reportValidator(identity address) returns()
func (_ElectionContract *ElectionContractTransactorSession) ReportValidator(identity common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.ReportValidator(&_ElectionContract.TransactOpts, identity)
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
