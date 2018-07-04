// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// ValidatorMgrABI is the input ABI used to generate the binding from.
const ValidatorMgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningTokenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"availableAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxNumValidators\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"},{\"name\":\"_miningTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ValidatorMgrBin is the compiled bytecode used for deploying new contracts.
const ValidatorMgrBin = `608060405260008060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b50604051608080611b5e83398101806040528101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183101515156100bb57600080fd5b836001819055508260028190555062015180820260038190555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050611a35806101296000396000f300608060405260043610610149576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461014e5780630a3cb663146101795780632086ca25146101a457806327378a8c146101cf5780633e83a283146102265780633ed0a373146102b95780633f4ba83a146103015780635c975abb14610318578063671b4d4914610347578063694746251461035e5780636a911ccf146103895780637071688a146103a0578063715018a6146103cb5780638456cb59146103e25780638da5cb5b146103f95780639363a1411461045057806397584b3e1461047b5780639bb2ea5a146104aa578063aded41ec146104d7578063b774cb1e146104ee578063c22a933c14610521578063cefddda91461054e578063e7a60a9c146105a9578063f2fde38b1461061d578063facd743b14610660575b600080fd5b34801561015a57600080fd5b506101636106bb565b6040518082815260200191505060405180910390f35b34801561018557600080fd5b5061018e61078e565b6040518082815260200191505060405180910390f35b3480156101b057600080fd5b506101b9610794565b6040518082815260200191505060405180910390f35b3480156101db57600080fd5b506101e461079a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561023257600080fd5b506102b7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506107c0565b005b3480156102c557600080fd5b506102e46004803603810190808035906020019092919050505061084e565b604051808381526020018281526020019250505060405180910390f35b34801561030d57600080fd5b506103166108c6565b005b34801561032457600080fd5b5061032d610984565b604051808215151515815260200191505060405180910390f35b34801561035357600080fd5b5061035c610997565b005b34801561036a57600080fd5b50610373610a55565b6040518082815260200191505060405180910390f35b34801561039557600080fd5b5061039e610a5b565b005b3480156103ac57600080fd5b506103b5610a96565b6040518082815260200191505060405180910390f35b3480156103d757600080fd5b506103e0610aa3565b005b3480156103ee57600080fd5b506103f7610ba5565b005b34801561040557600080fd5b5061040e610c65565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561045c57600080fd5b50610465610c8a565b6040518082815260200191505060405180910390f35b34801561048757600080fd5b50610490610cd7565b604051808215151515815260200191505060405180910390f35b3480156104b657600080fd5b506104d560048036038101908080359060200190929190505050610cea565b005b3480156104e357600080fd5b506104ec610d8e565b005b3480156104fa57600080fd5b50610503610fb0565b60405180826000191660001916815260200191505060405180910390f35b34801561052d57600080fd5b5061054c60048036038101908080359060200190929190505050610fb6565b005b34801561055a57600080fd5b5061058f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061101b565b604051808215151515815260200191505060405180910390f35b3480156105b557600080fd5b506105d460048036038101908080359060200190929190505050611074565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561062957600080fd5b5061065e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061112b565b005b34801561066c57600080fd5b506106a1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611192565b604051808215151515815260200191505060405180910390f35b6000806106c6610cd7565b156106d557600154915061078a565b6006600060076001600780549050038154811015156106f057fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561077457fe5b9060005260206000209060020201600001540191505b5090565b60035481565b60025481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60408051908101604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815250600860008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050610849610997565b505050565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156108a257fe5b90600052602060002090600202019050806000015481600101549250925050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561092157600080fd5b600060149054906101000a900460ff16151561093c57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b600060149054906101000a900460ff161515156109b357600080fd5b6109e1600860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611192565b1515156109ed57600080fd5b6109f56106bb565b60086001015410151515610a0857600080fd5b610a10610cd7565b1515610a1f57610a1e6111eb565b5b610a53600860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600860010154611237565b565b60015481565b600060149054906101000a900460ff16151515610a7757600080fd5b610a8033611192565b1515610a8b57600080fd5b610a943361157b565b565b6000600780549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610afe57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c0057600080fd5b600060149054906101000a900460ff16151515610c1c57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506002540311905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d4857600080fd5b600780549050831015610d825782600780549050039150600090505b81811015610d8157610d746111eb565b8080600101915050610d64565b5b82600281905550505050565b600080600080600060149054906101000a900460ff16151515610db057600080fd5b6000935060009250600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020191505b818054905083108015610e30575060008284815481101515610e1b57fe5b90600052602060002090600202016001015414155b15610e92578183815481101515610e4357fe5b906000526020600020906002020160010154421015610e6157610e92565b8183815481101515610e6f57fe5b906000526020600020906002020160000154840193508280600101935050610dfd565b610e9c33846116ed565b6000841115610faa57600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610f6d57600080fd5b505af1158015610f81573d6000803e3d6000fd5b505050506040513d6020811015610f9757600080fd5b8101908080519060200190929190505050505b50505050565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561101157600080fd5b8060018190555050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff169050919050565b600080600060078481548110151561108857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600201600182600201805490500381548110151561111157fe5b906000526020600020906002020160000154915050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561118657600080fd5b61118f816117da565b50565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b611235600760016007805490500381548110151561120557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661157b565b565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078790806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff02191690831515021790555060004314156113325760018460010160016101000a81548160ff0219169083151502179055505b8360020160408051908101604052808781526020016000815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050836000015492505b600083111561156b57600660006007600186038154811015156113b557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915081600201600183600201805490500381548110151561143757fe5b906000526020600020906002020190508060000154851115156114595761156b565b60076001840381548110151561146b57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007848154811015156114a557fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560076001850381548110151561150057fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050611396565b6115736118d4565b505050505050565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160078054905003811015611678576007600182018154811015156115e957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078281548110151561162357fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806001019150506115c7565b600780548091906001900361168d9190611957565b5060008260010160006101000a81548160ff02191690831515021790555060035442018260020160018460020180549050038154811015156116cb57fe5b9060005260206000209060020201600101819055506116e86118d4565b505050565b6000806000808414156116ff576117d3565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b82600201805490508110156117c157826002018181548110151561176857fe5b9060005260206000209060020201836002018381548110151561178757fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050611748565b8183600201816117d19190611983565b505b5050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561181657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600760405180828054801561193e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116118f4575b5050915050604051809103902060048160001916905550565b81548183558181111561197e5781836000526020600020918201910161197d91906119b5565b5b505050565b8154818355818111156119b0576002028160020283600052602060002091820191016119af91906119da565b5b505050565b6119d791905b808211156119d35760008160009055506001016119bb565b5090565b90565b611a0691905b80821115611a02576000808201600090556001820160009055506002016119e0565b5090565b905600a165627a7a72305820709920563c5cdfc296cb30f8b13371ae91f68a20cb72d9312529f56ac0bef9fa0029`

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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorMgr *ValidatorMgrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorMgr *ValidatorMgrSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.RenounceOwnership(&_ValidatorMgr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorMgr.Contract.RenounceOwnership(&_ValidatorMgr.TransactOpts)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ValidatorMgr *ValidatorMgrTransactor) SetBaseDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "setBaseDeposit", deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ValidatorMgr *ValidatorMgrSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.SetBaseDeposit(&_ValidatorMgr.TransactOpts, deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.SetBaseDeposit(&_ValidatorMgr.TransactOpts, deposit)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ValidatorMgr *ValidatorMgrTransactor) SetMaxValidators(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "setMaxValidators", max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ValidatorMgr *ValidatorMgrSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.SetMaxValidators(&_ValidatorMgr.TransactOpts, max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.SetMaxValidators(&_ValidatorMgr.TransactOpts, max)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ValidatorMgr *ValidatorMgrTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorMgr.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ValidatorMgr *ValidatorMgrSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.TransferOwnership(&_ValidatorMgr.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ValidatorMgr *ValidatorMgrTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorMgr.Contract.TransferOwnership(&_ValidatorMgr.TransactOpts, _newOwner)
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
