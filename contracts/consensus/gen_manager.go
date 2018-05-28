// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus

import (
	"math/big"
	"strings"

	kcoin "github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/event"
)

// ValidatorMgrABI is the input ABI used to generate the binding from.
const ValidatorMgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningTokenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"availableAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxNumValidators\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"},{\"name\":\"_miningTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ValidatorMgrBin is the compiled bytecode used for deploying new contracts.
const ValidatorMgrBin = `606060405260008060146101000a81548160ff021916908315150217905550341561002957600080fd5b6040516080806119ca83398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183101515156100b057600080fd5b836001819055508260028190555062015180820260038190555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506118ac8061011e6000396000f300606060405260043610610133576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf142146101385780630a3cb663146101615780632086ca251461018a57806327378a8c146101b35780633e83a283146102085780633ed0a3731461028d5780633f4ba83a146102cb5780635c975abb146102e0578063694746251461030d5780636a911ccf146103365780637071688a1461034b5780638456cb59146103745780638da5cb5b146103895780639363a141146103de57806397584b3e146104075780639bb2ea5a14610434578063aded41ec14610457578063b774cb1e1461046c578063c22a933c1461049d578063cefddda9146104c0578063e7a60a9c14610511578063f2fde38b1461057b578063facd743b146105b4575b600080fd5b341561014357600080fd5b61014b610605565b6040518082815260200191505060405180910390f35b341561016c57600080fd5b6101746106d9565b6040518082815260200191505060405180910390f35b341561019557600080fd5b61019d6106df565b6040518082815260200191505060405180910390f35b34156101be57600080fd5b6101c66106e5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561021357600080fd5b61028b600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190505061070b565b005b341561029857600080fd5b6102ae6004808035906020019091905050610799565b604051808381526020018281526020019250505060405180910390f35b34156102d657600080fd5b6102de610811565b005b34156102eb57600080fd5b6102f36108cf565b604051808215151515815260200191505060405180910390f35b341561031857600080fd5b6103206108e2565b6040518082815260200191505060405180910390f35b341561034157600080fd5b6103496108e8565b005b341561035657600080fd5b61035e610923565b6040518082815260200191505060405180910390f35b341561037f57600080fd5b610387610930565b005b341561039457600080fd5b61039c6109f0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103e957600080fd5b6103f1610a15565b6040518082815260200191505060405180910390f35b341561041257600080fd5b61041a610a62565b604051808215151515815260200191505060405180910390f35b341561043f57600080fd5b6104556004808035906020019091905050610a75565b005b341561046257600080fd5b61046a610b19565b005b341561047757600080fd5b61047f610d16565b60405180826000191660001916815260200191505060405180910390f35b34156104a857600080fd5b6104be6004808035906020019091905050610d1c565b005b34156104cb57600080fd5b6104f7600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610d81565b604051808215151515815260200191505060405180910390f35b341561051c57600080fd5b6105326004808035906020019091905050610dda565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b341561058657600080fd5b6105b2600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610e92565b005b34156105bf57600080fd5b6105eb600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610fe7565b604051808215151515815260200191505060405180910390f35b600080610610610a62565b1561061f5760015491506106d5565b60066000600760016007805490500381548110151561063a57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160018360020180549050038154811015156106bf57fe5b9060005260206000209060020201600001540191505b5090565b60035481565b60025481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60408051908101604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815250600860008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050610794611040565b505050565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156107ed57fe5b90600052602060002090600202019050806000015481600101549250925050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561086c57600080fd5b600060149054906101000a900460ff16151561088757600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b60015481565b600060149054906101000a900460ff1615151561090457600080fd5b61090d33610fe7565b151561091857600080fd5b610921336110fe565b565b6000600780549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561098b57600080fd5b600060149054906101000a900460ff161515156109a757600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506002540311905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ad357600080fd5b600780549050831015610b0d5782600780549050039150600090505b81811015610b0c57610aff611272565b8080600101915050610aef565b5b82600281905550505050565b600080600080600060149054906101000a900460ff16151515610b3b57600080fd5b6000935060009250600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020191505b818054905083108015610bbb575060008284815481101515610ba657fe5b90600052602060002090600202016001015414155b15610c1d578183815481101515610bce57fe5b906000526020600020906002020160010154421015610bec57610c1d565b8183815481101515610bfa57fe5b906000526020600020906002020160000154840193508280600101935050610b88565b610c2733846112bf565b6000841115610d1057600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1515610cf757600080fd5b5af11515610d0457600080fd5b50505060405180519050505b50505050565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d7757600080fd5b8060018190555050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff169050919050565b6000806000600784815481101515610dee57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806002016001826002018054905003815481101515610e7857fe5b906000526020600020906002020160000154915050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610eed57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610f2957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b600060149054906101000a900460ff1615151561105c57600080fd5b61108a600860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610fe7565b15151561109657600080fd5b61109e610605565b600860010154101515156110b157600080fd5b6110b9610a62565b15156110c8576110c7611272565b5b6110fc600860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166008600101546113ac565b565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b6001600780549050038110156111fd5760076001820181548110151561116c57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007828154811015156111a757fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808060010191505061114a565b60078054809190600190036112129190611770565b5060008260010160006101000a81548160ff021916908315150217905550600354420182600201600184600201805490500381548110151561125057fe5b90600052602060002090600202016001018190555061126d6116ed565b505050565b6112bd600760016007805490500381548110151561128c57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166110fe565b565b6000806000808414156112d1576113a5565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b826002018054905081101561139357826002018181548110151561133a57fe5b9060005260206000209060020201836002018381548110151561135957fe5b906000526020600020906002020160008201548160000155600182015481600101559050508180600101925050808060010191505061131a565b8183600201816113a3919061179c565b505b5050505050565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020935060016007805480600101828161140991906117ce565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff02191690831515021790555060004314156114a45760018460010160016101000a81548160ff0219169083151502179055505b8360020180548060010182816114ba91906117fa565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b60008311156116dd576006600060076001860381548110151561152357fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002091508160020160018360020180549050038154811015156115a657fe5b906000526020600020906002020190508060000154851115156115c8576116dd565b6007600184038154811015156115da57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078481548110151561161557fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560076001850381548110151561167157fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050611504565b6116e56116ed565b505050505050565b600760405180828054801561175757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161170d575b5050915050604051809103902060048160001916905550565b81548183558181151161179757818360005260206000209182019101611796919061182c565b5b505050565b8154818355818115116117c9576002028160020283600052602060002091820191016117c89190611851565b5b505050565b8154818355818115116117f5578183600052602060002091820191016117f4919061182c565b5b505050565b815481835581811511611827576002028160020283600052602060002091820191016118269190611851565b5b505050565b61184e91905b8082111561184a576000816000905550600101611832565b5090565b90565b61187d91905b8082111561187957600080820160009055600182016000905550600201611857565b5090565b905600a165627a7a72305820d31cdc287f51905ea7e4d93bf3b330ac11990062a7cf7cf563311a5f4858fcfd0029`

// DeployValidatorMgr deploys a new Kowala contract, binding an instance of ValidatorMgr to it.
func DeployValidatorMgr(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxNumValidators *big.Int, _freezePeriod *big.Int, _miningTokenAddr common.Address) (common.Address, *types.Transaction, *ValidatorMgr, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorMgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorMgrBin), backend, _baseDeposit, _maxNumValidators, _freezePeriod, _miningTokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMgr{ValidatorMgrCaller: ValidatorMgrCaller{contract: contract}, ValidatorMgrTransactor: ValidatorMgrTransactor{contract: contract}, ValidatorMgrFilterer: ValidatorMgrFilterer{contract: contract}}, nil
}

// ValidatorMgr is an auto generated Go binding around an Kowala contract.
type ValidatorMgr struct {
	ValidatorMgrCaller     // Read-only binding to the contract
	ValidatorMgrTransactor // Write-only binding to the contract
	ValidatorMgrFilterer   // Log filterer for contract events
}

// ValidatorMgrCaller is an auto generated read-only Go binding around an Kowala contract.
type ValidatorMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrTransactor is an auto generated write-only Go binding around an Kowala contract.
type ValidatorMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrFilterer is an auto generated log filtering Go binding around an Kowala contract events.
type ValidatorMgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrSession is an auto generated Go binding around an Kowala contract,
// with pre-set call and transact options.
type ValidatorMgrSession struct {
	Contract     *ValidatorMgr     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorMgrCallerSession is an auto generated read-only Go binding around an Kowala contract,
// with pre-set call options.
type ValidatorMgrCallerSession struct {
	Contract *ValidatorMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorMgrTransactorSession is an auto generated write-only Go binding around an Kowala contract,
// with pre-set transact options.
type ValidatorMgrTransactorSession struct {
	Contract     *ValidatorMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorMgrRaw is an auto generated low-level Go binding around an Kowala contract.
type ValidatorMgrRaw struct {
	Contract *ValidatorMgr // Generic contract binding to access the raw methods on
}

// ValidatorMgrCallerRaw is an auto generated low-level read-only Go binding around an Kowala contract.
type ValidatorMgrCallerRaw struct {
	Contract *ValidatorMgrCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMgrTransactorRaw is an auto generated low-level write-only Go binding around an Kowala contract.
type ValidatorMgrTransactorRaw struct {
	Contract *ValidatorMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorMgr creates a new instance of ValidatorMgr, bound to a specific deployed contract.
func NewValidatorMgr(address common.Address, backend bind.ContractBackend) (*ValidatorMgr, error) {
	contract, err := bindValidatorMgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgr{ValidatorMgrCaller: ValidatorMgrCaller{contract: contract}, ValidatorMgrTransactor: ValidatorMgrTransactor{contract: contract}, ValidatorMgrFilterer: ValidatorMgrFilterer{contract: contract}}, nil
}

// NewValidatorMgrCaller creates a new read-only instance of ValidatorMgr, bound to a specific deployed contract.
func NewValidatorMgrCaller(address common.Address, caller bind.ContractCaller) (*ValidatorMgrCaller, error) {
	contract, err := bindValidatorMgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrCaller{contract: contract}, nil
}

// NewValidatorMgrTransactor creates a new write-only instance of ValidatorMgr, bound to a specific deployed contract.
func NewValidatorMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorMgrTransactor, error) {
	contract, err := bindValidatorMgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrTransactor{contract: contract}, nil
}

// NewValidatorMgrFilterer creates a new log filterer instance of ValidatorMgr, bound to a specific deployed contract.
func NewValidatorMgrFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorMgrFilterer, error) {
	contract, err := bindValidatorMgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrFilterer{contract: contract}, nil
}

// bindValidatorMgr binds a generic wrapper to an already deployed contract.
func bindValidatorMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorMgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// HasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorMgr *ValidatorMgrCaller) HasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// HasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorMgr *ValidatorMgrSession) HasAvailability() (bool, error) {
	return _ValidatorMgr.Contract.HasAvailability(&_ValidatorMgr.CallOpts)
}

// HasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorMgr *ValidatorMgrCallerSession) HasAvailability() (bool, error) {
	return _ValidatorMgr.Contract.HasAvailability(&_ValidatorMgr.CallOpts)
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

// ValidatorMgrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorMgr contract.
type ValidatorMgrOwnershipTransferredIterator struct {
	Event *ValidatorMgrOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log     // Log channel receiving the found contract events
	sub  kcoin.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorMgrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMgrOwnershipTransferred)
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
		it.Event = new(ValidatorMgrOwnershipTransferred)
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
func (it *ValidatorMgrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMgrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMgrOwnershipTransferred represents a OwnershipTransferred event raised by the ValidatorMgr contract.
type ValidatorMgrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ValidatorMgr *ValidatorMgrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorMgrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorMgr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrOwnershipTransferredIterator{contract: _ValidatorMgr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ValidatorMgr *ValidatorMgrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorMgrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorMgr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMgrOwnershipTransferred)
				if err := _ValidatorMgr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ValidatorMgrPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ValidatorMgr contract.
type ValidatorMgrPauseIterator struct {
	Event *ValidatorMgrPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log     // Log channel receiving the found contract events
	sub  kcoin.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorMgrPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMgrPause)
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
		it.Event = new(ValidatorMgrPause)
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
func (it *ValidatorMgrPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMgrPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMgrPause represents a Pause event raised by the ValidatorMgr contract.
type ValidatorMgrPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ValidatorMgr *ValidatorMgrFilterer) FilterPause(opts *bind.FilterOpts) (*ValidatorMgrPauseIterator, error) {

	logs, sub, err := _ValidatorMgr.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrPauseIterator{contract: _ValidatorMgr.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ValidatorMgr *ValidatorMgrFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ValidatorMgrPause) (event.Subscription, error) {

	logs, sub, err := _ValidatorMgr.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMgrPause)
				if err := _ValidatorMgr.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ValidatorMgrUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ValidatorMgr contract.
type ValidatorMgrUnpauseIterator struct {
	Event *ValidatorMgrUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log     // Log channel receiving the found contract events
	sub  kcoin.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorMgrUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMgrUnpause)
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
		it.Event = new(ValidatorMgrUnpause)
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
func (it *ValidatorMgrUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMgrUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMgrUnpause represents a Unpause event raised by the ValidatorMgr contract.
type ValidatorMgrUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ValidatorMgr *ValidatorMgrFilterer) FilterUnpause(opts *bind.FilterOpts) (*ValidatorMgrUnpauseIterator, error) {

	logs, sub, err := _ValidatorMgr.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrUnpauseIterator{contract: _ValidatorMgr.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ValidatorMgr *ValidatorMgrFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ValidatorMgrUnpause) (event.Subscription, error) {

	logs, sub, err := _ValidatorMgr.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMgrUnpause)
				if err := _ValidatorMgr.contract.UnpackLog(event, "Unpause", log); err != nil {
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
