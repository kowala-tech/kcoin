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
const NetworkContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setBaseDepositUpperBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"releasedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"setBaseDepositLowerBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsUpperBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDepositUpperBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availability\",\"outputs\":[{\"name\":\"available\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDepositLowerBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorsLowerBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorsUpperBound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsLowerBound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"},{\"name\":\"_maxValidators\",\"type\":\"uint256\"},{\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// NetworkContractBin is the compiled bytecode used for deploying new contracts.
const NetworkContractBin = `60606040526040516080806200198b83398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508334101515156200008c57600080fd5b600182101515156200009d57600080fd5b60008110151515620000ae57600080fd5b836001819055508160048190555082600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600b8190555060026001548115156200011357fe5b0460038190555060026001540260028190555060026004548115156200013557fe5b0460068190555060026004540260058190555062000168833462000172640100000000026200114f176401000000009004565b5050505062000401565b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600160078054806001018281620001cd919062000343565b9160005260206000209001600086909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003816000018190555060018160010160006101000a81548160ff02191690831515021790555080600201805480600101828162000259919062000372565b916000526020600020906002020160006040805190810160405280868152602001600081525090919091506000820151816000015560208201518160010155505050620002b9620002be64010000000002620013ed176401000000009004565b505050565b60076040518082805480156200032a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620002df575b5050915050604051809103902060098160001916905550565b8154818355818115116200036d578183600052602060002091820191016200036c9190620003a7565b5b505050565b815481835581811511620003a257600202816002028360005260206000209182019101620003a19190620003cf565b5b505050565b620003cc91905b80821115620003c8576000816000905550600101620003ae565b5090565b90565b620003fe91905b80821115620003fa57600080820160009055600182016000905550600201620003d6565b5090565b90565b61157a80620004116000396000f30060606040526004361061015f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461016457806308ac52561461018d5780631cbd5487146101b65780633ccfd60b146101d95780633ed0a373146101ee5780634b1d422f1461022c578063694746251461024f5780636cf6d675146102785780637071688a146102a157806376792094146102ca5780639363a141146102ed5780639bb2ea5a14610316578063a6554f1a14610339578063a7f0b3de14610362578063b774cb1e146103b7578063c22a933c146103e8578063c9b539001461040b578063cbfb94ce14610434578063cefddda91461045d578063d0e30db0146104ae578063d66d9e19146104b8578063e4f4410b146104cd578063e7277b12146104f6578063e7a60a9c1461051f578063e99cc69614610589578063f2fde38b146105ac578063facd743b146105e5575b600080fd5b341561016f57600080fd5b610177610636565b6040518082815260200191505060405180910390f35b341561019857600080fd5b6101a061070d565b6040518082815260200191505060405180910390f35b34156101c157600080fd5b6101d76004808035906020019091905050610713565b005b34156101e457600080fd5b6101ec610789565b005b34156101f957600080fd5b61020f600480803590602001909190505061098a565b604051808381526020018281526020019250505060405180910390f35b341561023757600080fd5b61024d6004808035906020019091905050610a02565b005b341561025a57600080fd5b610262610a78565b6040518082815260200191505060405180910390f35b341561028357600080fd5b61028b610a7e565b6040518082815260200191505060405180910390f35b34156102ac57600080fd5b6102b4610a84565b6040518082815260200191505060405180910390f35b34156102d557600080fd5b6102eb6004808035906020019091905050610a91565b005b34156102f857600080fd5b610300610b07565b6040518082815260200191505060405180910390f35b341561032157600080fd5b6103376004808035906020019091905050610b54565b005b341561034457600080fd5b61034c610c19565b6040518082815260200191505060405180910390f35b341561036d57600080fd5b610375610c1f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103c257600080fd5b6103ca610c45565b60405180826000191660001916815260200191505060405180910390f35b34156103f357600080fd5b6104096004808035906020019091905050610c4b565b005b341561041657600080fd5b61041e610cd1565b6040518082815260200191505060405180910390f35b341561043f57600080fd5b610447610ce2565b6040518082815260200191505060405180910390f35b341561046857600080fd5b610494600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ce8565b604051808215151515815260200191505060405180910390f35b6104b6610d42565b005b34156104c357600080fd5b6104cb610d7d565b005b34156104d857600080fd5b6104e0610d9c565b6040518082815260200191505060405180910390f35b341561050157600080fd5b610509610da2565b6040518082815260200191505060405180910390f35b341561052a57600080fd5b6105406004808035906020019091905050610da8565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b341561059457600080fd5b6105aa6004808035906020019091905050610e60565b005b34156105b757600080fd5b6105e3600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ed6565b005b34156105f057600080fd5b61061c600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061102c565b604051808215151515815260200191505060405180910390f35b6000806000610643610cd1565b1115610653576001549150610709565b60086000600760016007805490500381548110151561066e57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160018360020180549050038154811015156106f357fe5b9060005260206000209060020201600001540191505b5090565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561076e57600080fd5b600354811015151561077f57600080fd5b8060028190555050565b60008060006107973361102c565b15156107a257600080fd5b6000925060009150600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090505b80600201805490508210801561082557506000816002018381548110151561081057fe5b90600052602060002090600202016001015414155b1561088d57806002018281548110151561083b57fe5b9060005260206000209060020201600101544210156108595761088d565b806002018281548110151561086a57fe5b9060005260206000209060020201600001548301925081806001019250506107ec565b61093b8160606040519081016040529081600082015481526020016001820160009054906101000a900460ff1615151515815260200160028201805480602002602001604051908101604052809291908181526020016000905b8282101561092d578382906000526020600020906002020160408051908101604052908160008201548152602001600182015481525050815260200190600101906108e7565b505050508152505083611085565b6000831115610985573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561098457600080fd5b5b505050565b6000806000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156109de57fe5b90600052602060002090600202019050806000015481600101549250925050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a5d57600080fd5b6002548111151515610a6e57600080fd5b8060038190555050565b60015481565b600b5481565b6000600780549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610aec57600080fd5b6006548110151515610afd57600080fd5b8060058190555050565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bb257600080fd5b826006548110158015610bc757506005548111155b1515610bd257600080fd5b6000610bdc610cd1565b1415610c0c5783600454039250600091505b82821015610c0b57610bfe611102565b8180600101925050610bee565b5b8360048190555050505050565b60025481565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ca657600080fd5b806003548110158015610cbb57506002548111155b1515610cc657600080fd5b816001819055505050565b600060078054905060045403905090565b60035481565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b610d4a610636565b3410151515610d5857600080fd5b6000610d62610cd1565b1415610d7157610d70611102565b5b610d7b333461114f565b565b610d863361102c565b1515610d9157600080fd5b610d9a33611281565b565b60065481565b60055481565b6000806000600784815481101515610dbc57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806002016001826002018054905003815481101515610e4657fe5b906000526020600020906002020160000154915050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ebb57600080fd5b6005548111151515610ecc57600080fd5b8060068190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f3157600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b6000806000831415611096576110fc565b600091508290505b8360400151518110156110fb578360400151818151811015156110bd57fe5b906020019060200201518460400151838151811015156110d957fe5b906020019060200201819052508180600101925050808060010191505061109e565b5b50505050565b61114d600760016007805490500381548110151561111c57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611281565b565b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506001600780548060010182816111a89190611470565b9160005260206000209001600086909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003816000018190555060018160010160006101000a81548160ff021916908315150217905550806002018054806001018281611232919061149c565b91600052602060002090600202016000604080519081016040528086815260200160008152509091909150600082015181600001556020820151816001015550505061127c6113ed565b505050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160078054905003811015611380576007600182018154811015156112ef57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078281548110151561132a57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806001019150506112cd565b600780548091906001900361139591906114ce565b5060008260010160006101000a81548160ff021916908315150217905550600b5442018260020160018460020180549050038154811015156113d357fe5b906000526020600020906002020160010181905550505050565b600760405180828054801561145757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161140d575b5050915050604051809103902060098160001916905550565b8154818355818115116114975781836000526020600020918201910161149691906114fa565b5b505050565b8154818355818115116114c9576002028160020283600052602060002091820191016114c8919061151f565b5b505050565b8154818355818115116114f5578183600052602060002091820191016114f491906114fa565b5b505050565b61151c91905b80821115611518576000816000905550600101611500565b5090565b90565b61154b91905b8082111561154757600080820160009055600182016000905550600201611525565b5090565b905600a165627a7a7230582053ca3fd9ba122d815f35917f4bc5199c0d5b2cd5c629c73341eddc54b8115ffc0029`

// DeployNetworkContract deploys a new Ethereum contract, binding an instance of NetworkContract to it.
func DeployNetworkContract(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _genesis common.Address, _maxValidators *big.Int, _unbondingPeriod *big.Int) (common.Address, *types.Transaction, *NetworkContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkContractBin), backend, _baseDeposit, _genesis, _maxValidators, _unbondingPeriod)
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

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) BaseDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDeposit(&_NetworkContract.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) BaseDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDeposit(&_NetworkContract.CallOpts)
}

// BaseDepositLowerBound is a free data retrieval call binding the contract method 0xcbfb94ce.
//
// Solidity: function baseDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) BaseDepositLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "baseDepositLowerBound")
	return *ret0, err
}

// BaseDepositLowerBound is a free data retrieval call binding the contract method 0xcbfb94ce.
//
// Solidity: function baseDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) BaseDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositLowerBound(&_NetworkContract.CallOpts)
}

// BaseDepositLowerBound is a free data retrieval call binding the contract method 0xcbfb94ce.
//
// Solidity: function baseDepositLowerBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) BaseDepositLowerBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositLowerBound(&_NetworkContract.CallOpts)
}

// BaseDepositUpperBound is a free data retrieval call binding the contract method 0xa6554f1a.
//
// Solidity: function baseDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCaller) BaseDepositUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "baseDepositUpperBound")
	return *ret0, err
}

// BaseDepositUpperBound is a free data retrieval call binding the contract method 0xa6554f1a.
//
// Solidity: function baseDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractSession) BaseDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositUpperBound(&_NetworkContract.CallOpts)
}

// BaseDepositUpperBound is a free data retrieval call binding the contract method 0xa6554f1a.
//
// Solidity: function baseDepositUpperBound() constant returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) BaseDepositUpperBound() (*big.Int, error) {
	return _NetworkContract.Contract.BaseDepositUpperBound(&_NetworkContract.CallOpts)
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

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_NetworkContract *NetworkContractCaller) GetDepositAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	ret := new(struct {
		Amount     *big.Int
		ReleasedAt *big.Int
	})
	out := ret
	err := _NetworkContract.contract.Call(opts, out, "getDepositAtIndex", index)
	return *ret, err
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_NetworkContract *NetworkContractSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	return _NetworkContract.Contract.GetDepositAtIndex(&_NetworkContract.CallOpts, index)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, releasedAt uint256)
func (_NetworkContract *NetworkContractCallerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount     *big.Int
	ReleasedAt *big.Int
}, error) {
	return _NetworkContract.Contract.GetDepositAtIndex(&_NetworkContract.CallOpts, index)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCaller) GetDepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "getDepositCount")
	return *ret0, err
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractSession) GetDepositCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetDepositCount(&_NetworkContract.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_NetworkContract *NetworkContractCallerSession) GetDepositCount() (*big.Int, error) {
	return _NetworkContract.Contract.GetDepositCount(&_NetworkContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_NetworkContract *NetworkContractCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_NetworkContract *NetworkContractSession) GetMinimumDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.GetMinimumDeposit(&_NetworkContract.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_NetworkContract *NetworkContractCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _NetworkContract.Contract.GetMinimumDeposit(&_NetworkContract.CallOpts)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_NetworkContract *NetworkContractCaller) GetValidatorAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _NetworkContract.contract.Call(opts, out, "getValidatorAtIndex", index)
	return *ret, err
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_NetworkContract *NetworkContractSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _NetworkContract.Contract.GetValidatorAtIndex(&_NetworkContract.CallOpts, index)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_NetworkContract *NetworkContractCallerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _NetworkContract.Contract.GetValidatorAtIndex(&_NetworkContract.CallOpts, index)
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
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsGenesisValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isGenesisValidator", code)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisValidator(&_NetworkContract.CallOpts, code)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsGenesisValidator(&_NetworkContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCaller) IsValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NetworkContract.contract.Call(opts, out, "isValidator", code)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractSession) IsValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsValidator(&_NetworkContract.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_NetworkContract *NetworkContractCallerSession) IsValidator(code common.Address) (bool, error) {
	return _NetworkContract.Contract.IsValidator(&_NetworkContract.CallOpts, code)
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

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NetworkContract *NetworkContractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NetworkContract *NetworkContractSession) Deposit() (*types.Transaction, error) {
	return _NetworkContract.Contract.Deposit(&_NetworkContract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NetworkContract *NetworkContractTransactorSession) Deposit() (*types.Transaction, error) {
	return _NetworkContract.Contract.Deposit(&_NetworkContract.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_NetworkContract *NetworkContractTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_NetworkContract *NetworkContractSession) Leave() (*types.Transaction, error) {
	return _NetworkContract.Contract.Leave(&_NetworkContract.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_NetworkContract *NetworkContractTransactorSession) Leave() (*types.Transaction, error) {
	return _NetworkContract.Contract.Leave(&_NetworkContract.TransactOpts)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetBaseDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setBaseDeposit", deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDeposit(&_NetworkContract.TransactOpts, deposit)
}

// SetBaseDepositLowerBound is a paid mutator transaction binding the contract method 0x4b1d422f.
//
// Solidity: function setBaseDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetBaseDepositLowerBound(opts *bind.TransactOpts, min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setBaseDepositLowerBound", min)
}

// SetBaseDepositLowerBound is a paid mutator transaction binding the contract method 0x4b1d422f.
//
// Solidity: function setBaseDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractSession) SetBaseDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetBaseDepositLowerBound is a paid mutator transaction binding the contract method 0x4b1d422f.
//
// Solidity: function setBaseDepositLowerBound(min uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetBaseDepositLowerBound(min *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositLowerBound(&_NetworkContract.TransactOpts, min)
}

// SetBaseDepositUpperBound is a paid mutator transaction binding the contract method 0x1cbd5487.
//
// Solidity: function setBaseDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactor) SetBaseDepositUpperBound(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setBaseDepositUpperBound", max)
}

// SetBaseDepositUpperBound is a paid mutator transaction binding the contract method 0x1cbd5487.
//
// Solidity: function setBaseDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractSession) SetBaseDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositUpperBound(&_NetworkContract.TransactOpts, max)
}

// SetBaseDepositUpperBound is a paid mutator transaction binding the contract method 0x1cbd5487.
//
// Solidity: function setBaseDepositUpperBound(max uint256) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetBaseDepositUpperBound(max *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetBaseDepositUpperBound(&_NetworkContract.TransactOpts, max)
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

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NetworkContract *NetworkContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NetworkContract *NetworkContractSession) Withdraw() (*types.Transaction, error) {
	return _NetworkContract.Contract.Withdraw(&_NetworkContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NetworkContract *NetworkContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _NetworkContract.Contract.Withdraw(&_NetworkContract.TransactOpts)
}
