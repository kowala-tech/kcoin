// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
)

// ValidatorMgrABI is the input ABI used to generate the binding from.
const ValidatorMgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"superNodeAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"availableAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isSuperNode\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"knsResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxNumValidators\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"},{\"name\":\"_superNodeAmount\",\"type\":\"uint256\"},{\"name\":\"_resolverAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ValidatorMgrBin is the compiled bytecode used for deploying new contracts.
const ValidatorMgrBin = `608060405260008060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b5060405160a080611e178339810180604052810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000841115156100c457600080fd5b84600181905550836002819055506201518083026003819055508160068190555080600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507366da4ac1767b04b0d99bc94ccad6eef8da63ae9663098799626040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260128152602001807f6d696e696e67746f6b656e2e6b6f77616c61000000000000000000000000000081525060200191505060206040518083038186803b1580156101c257600080fd5b505af41580156101d6573d6000803e3d6000fd5b505050506040513d60208110156101ec57600080fd5b8101908080519060200190929190505050600581600019169055505050505050611bfc8061021b6000396000f300608060405260043610610154576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf142146101595780630a3cb663146101845780632086ca25146101af57806326833148146101da5780633e83a283146102055780633ed0a373146102985780633f4ba83a146102e05780635c975abb146102f757806369474625146103265780636a911ccf146103515780637071688a14610368578063715018a6146103935780637d0e81bf146103aa5780638456cb59146104055780638da5cb5b1461041c5780639363a1411461047357806397584b3e1461049e5780639bb2ea5a146104cd578063a2207c6a146104fa578063aded41ec14610551578063b774cb1e14610568578063c22a933c1461059b578063cefddda9146105c8578063e7a60a9c14610623578063f2fde38b14610697578063facd743b146106da575b600080fd5b34801561016557600080fd5b5061016e610735565b6040518082815260200191505060405180910390f35b34801561019057600080fd5b50610199610808565b6040518082815260200191505060405180910390f35b3480156101bb57600080fd5b506101c461080e565b6040518082815260200191505060405180910390f35b3480156101e657600080fd5b506101ef610814565b6040518082815260200191505060405180910390f35b34801561021157600080fd5b50610296600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061081a565b005b3480156102a457600080fd5b506102c3600480360381019080803590602001909291905050506108a8565b604051808381526020018281526020019250505060405180910390f35b3480156102ec57600080fd5b506102f5610920565b005b34801561030357600080fd5b5061030c6109de565b604051808215151515815260200191505060405180910390f35b34801561033257600080fd5b5061033b6109f1565b6040518082815260200191505060405180910390f35b34801561035d57600080fd5b506103666109f7565b005b34801561037457600080fd5b5061037d610a32565b6040518082815260200191505060405180910390f35b34801561039f57600080fd5b506103a8610a3f565b005b3480156103b657600080fd5b506103eb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b41565b604051808215151515815260200191505060405180910390f35b34801561041157600080fd5b5061041a610bd5565b005b34801561042857600080fd5b50610431610c95565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561047f57600080fd5b50610488610cba565b6040518082815260200191505060405180910390f35b3480156104aa57600080fd5b506104b3610d07565b604051808215151515815260200191505060405180910390f35b3480156104d957600080fd5b506104f860048036038101908080359060200190929190505050610d1a565b005b34801561050657600080fd5b5061050f610dbe565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561055d57600080fd5b50610566610de4565b005b34801561057457600080fd5b5061057d6110b9565b60405180826000191660001916815260200191505060405180910390f35b3480156105a757600080fd5b506105c6600480360381019080803590602001909291905050506110bf565b005b3480156105d457600080fd5b50610609600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611124565b604051808215151515815260200191505060405180910390f35b34801561062f57600080fd5b5061064e6004803603810190808035906020019092919050505061117d565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156106a357600080fd5b506106d8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611234565b005b3480156106e657600080fd5b5061071b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061129b565b604051808215151515815260200191505060405180910390f35b600080610740610d07565b1561074f576001549150610804565b60086000600960016009805490500381548110151561076a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160018360020180549050038154811015156107ee57fe5b9060005260206000209060020201600001540191505b5090565b60035481565b60025481565b60065481565b60408051908101604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815250600a60008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101559050506108a36112f4565b505050565b6000806000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156108fc57fe5b90600052602060002090600202019050806000015481600101549250925050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561097b57600080fd5b600060149054906101000a900460ff16151561099657600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b60015481565b600060149054906101000a900460ff16151515610a1357600080fd5b610a1c3361129b565b1515610a2757600080fd5b610a30336113b2565b565b6000600980549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a9a57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080610b4d8361129b565b1515610b5c5760009150610bcf565b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002019050600654816001838054905003815481101515610bb857fe5b906000526020600020906002020160000154101591505b50919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c3057600080fd5b600060149054906101000a900460ff16151515610c4c57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806009805490506002540311905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d7857600080fd5b600980549050831015610db25782600980549050039150600090505b81811015610db157610da4611524565b8080600101915050610d94565b5b82600281905550505050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600080600060149054906101000a900460ff16151515610e0657600080fd5b6000935060009250600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020191505b818054905083108015610e86575060008284815481101515610e7157fe5b90600052602060002090600202016001015414155b15610ee8578183815481101515610e9957fe5b906000526020600020906002020160010154421015610eb757610ee8565b8183815481101515610ec557fe5b906000526020600020906002020160000154840193508280600101935050610e53565b610ef23384611570565b60008411156110b357600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633b3b57de6005546040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015610f9657600080fd5b505af1158015610faa573d6000803e3d6000fd5b505050506040513d6020811015610fc057600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561107657600080fd5b505af115801561108a573d6000803e3d6000fd5b505050506040513d60208110156110a057600080fd5b8101908080519060200190929190505050505b50505050565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561111a57600080fd5b8060018190555050565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff169050919050565b600080600060098481548110151561119157fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600201600182600201805490500381548110151561121a57fe5b906000526020600020906002020160000154915050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561128f57600080fd5b6112988161165d565b50565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b600060149054906101000a900460ff1615151561131057600080fd5b61133e600a60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661129b565b15151561134a57600080fd5b611352610735565b600a600101541015151561136557600080fd5b61136d610d07565b151561137c5761137b611524565b5b6113b0600a60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600a60010154611757565b565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b6001600980549050038110156114af5760096001820181548110151561142057fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660098281548110151561145a57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806001019150506113fe565b60098054809190600190036114c49190611b1e565b5060008260010160006101000a81548160ff021916908315150217905550600354420182600201600184600201805490500381548110151561150257fe5b90600052602060002090600202016001018190555061151f611a9b565b505050565b61156e600960016009805490500381548110151561153e57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166113b2565b565b60008060008084141561158257611656565b600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b82600201805490508110156116445782600201818154811015156115eb57fe5b9060005260206000209060020201836002018381548110151561160a57fe5b90600052602060002090600202016000820154816000015560018201548160010155905050818060010192505080806001019150506115cb565b8183600201816116549190611b4a565b505b5050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561169957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080600080600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160098790806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff02191690831515021790555060004314156118525760018460010160016101000a81548160ff0219169083151502179055505b8360020160408051908101604052808781526020016000815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050836000015492505b6000831115611a8b57600860006009600186038154811015156118d557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915081600201600183600201805490500381548110151561195757fe5b9060005260206000209060020201905080600001548511151561197957611a8b565b60096001840381548110151561198b57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166009848154811015156119c557fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600960018503815481101515611a2057fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508282600001819055506001830384600001819055508280600190039350506118b6565b611a93611a9b565b505050505050565b6009604051808280548015611b0557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611abb575b5050915050604051809103902060048160001916905550565b815481835581811115611b4557818360005260206000209182019101611b449190611b7c565b5b505050565b815481835581811115611b7757600202816002028360005260206000209182019101611b769190611ba1565b5b505050565b611b9e91905b80821115611b9a576000816000905550600101611b82565b5090565b90565b611bcd91905b80821115611bc957600080820160009055600182016000905550600201611ba7565b5090565b905600a165627a7a72305820733cc1a4b66deb2779b2d6f359db827881d1fc1fc6b95ff0efe88faac5fd13de0029`

// DeployValidatorMgr deploys a new Kowala contract, binding an instance of ValidatorMgr to it.
func DeployValidatorMgr(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxNumValidators *big.Int, _freezePeriod *big.Int, _superNodeAmount *big.Int, _resolverAddr common.Address) (common.Address, *types.Transaction, *ValidatorMgr, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorMgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorMgrBin), backend, _baseDeposit, _maxNumValidators, _freezePeriod, _superNodeAmount, _resolverAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMgr{ValidatorMgrCaller: ValidatorMgrCaller{contract: contract}, ValidatorMgrTransactor: ValidatorMgrTransactor{contract: contract}, ValidatorMgrFilterer: ValidatorMgrFilterer{contract: contract}}, nil
}

// ValidatorMgr is an auto generated Go binding around a Kowala contract.
type ValidatorMgr struct {
	ValidatorMgrCaller     // Read-only binding to the contract
	ValidatorMgrTransactor // Write-only binding to the contract
	ValidatorMgrFilterer   // Log filterer for contract events
}

// ValidatorMgrCaller is an auto generated read-only Go binding around a Kowala contract.
type ValidatorMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrTransactor is an auto generated write-only Go binding around a Kowala contract.
type ValidatorMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type ValidatorMgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMgrSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type ValidatorMgrSession struct {
	Contract     *ValidatorMgr     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorMgrCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type ValidatorMgrCallerSession struct {
	Contract *ValidatorMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorMgrTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type ValidatorMgrTransactorSession struct {
	Contract     *ValidatorMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorMgrRaw is an auto generated low-level Go binding around a Kowala contract.
type ValidatorMgrRaw struct {
	Contract *ValidatorMgr // Generic contract binding to access the raw methods on
}

// ValidatorMgrCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type ValidatorMgrCallerRaw struct {
	Contract *ValidatorMgrCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMgrTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
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

// IsSuperNode is a free data retrieval call binding the contract method 0x7d0e81bf.
//
// Solidity: function isSuperNode(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrCaller) IsSuperNode(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "isSuperNode", code)
	return *ret0, err
}

// IsSuperNode is a free data retrieval call binding the contract method 0x7d0e81bf.
//
// Solidity: function isSuperNode(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrSession) IsSuperNode(code common.Address) (bool, error) {
	return _ValidatorMgr.Contract.IsSuperNode(&_ValidatorMgr.CallOpts, code)
}

// IsSuperNode is a free data retrieval call binding the contract method 0x7d0e81bf.
//
// Solidity: function isSuperNode(code address) constant returns(isIndeed bool)
func (_ValidatorMgr *ValidatorMgrCallerSession) IsSuperNode(code common.Address) (bool, error) {
	return _ValidatorMgr.Contract.IsSuperNode(&_ValidatorMgr.CallOpts, code)
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

// KnsResolver is a free data retrieval call binding the contract method 0xa2207c6a.
//
// Solidity: function knsResolver() constant returns(address)
func (_ValidatorMgr *ValidatorMgrCaller) KnsResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "knsResolver")
	return *ret0, err
}

// KnsResolver is a free data retrieval call binding the contract method 0xa2207c6a.
//
// Solidity: function knsResolver() constant returns(address)
func (_ValidatorMgr *ValidatorMgrSession) KnsResolver() (common.Address, error) {
	return _ValidatorMgr.Contract.KnsResolver(&_ValidatorMgr.CallOpts)
}

// KnsResolver is a free data retrieval call binding the contract method 0xa2207c6a.
//
// Solidity: function knsResolver() constant returns(address)
func (_ValidatorMgr *ValidatorMgrCallerSession) KnsResolver() (common.Address, error) {
	return _ValidatorMgr.Contract.KnsResolver(&_ValidatorMgr.CallOpts)
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

// SuperNodeAmount is a free data retrieval call binding the contract method 0x26833148.
//
// Solidity: function superNodeAmount() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCaller) SuperNodeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMgr.contract.Call(opts, out, "superNodeAmount")
	return *ret0, err
}

// SuperNodeAmount is a free data retrieval call binding the contract method 0x26833148.
//
// Solidity: function superNodeAmount() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrSession) SuperNodeAmount() (*big.Int, error) {
	return _ValidatorMgr.Contract.SuperNodeAmount(&_ValidatorMgr.CallOpts)
}

// SuperNodeAmount is a free data retrieval call binding the contract method 0x26833148.
//
// Solidity: function superNodeAmount() constant returns(uint256)
func (_ValidatorMgr *ValidatorMgrCallerSession) SuperNodeAmount() (*big.Int, error) {
	return _ValidatorMgr.Contract.SuperNodeAmount(&_ValidatorMgr.CallOpts)
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

// ValidatorMgrOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ValidatorMgr contract.
type ValidatorMgrOwnershipRenouncedIterator struct {
	Event *ValidatorMgrOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ValidatorMgrOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMgrOwnershipRenounced)
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
		it.Event = new(ValidatorMgrOwnershipRenounced)
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
func (it *ValidatorMgrOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMgrOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMgrOwnershipRenounced represents a OwnershipRenounced event raised by the ValidatorMgr contract.
type ValidatorMgrOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ValidatorMgr *ValidatorMgrFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ValidatorMgrOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ValidatorMgr.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrOwnershipRenouncedIterator{contract: _ValidatorMgr.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ValidatorMgr *ValidatorMgrFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ValidatorMgrOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ValidatorMgr.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMgrOwnershipRenounced)
				if err := _ValidatorMgr.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ValidatorMgrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorMgr contract.
type ValidatorMgrOwnershipTransferredIterator struct {
	Event *ValidatorMgrOwnershipTransferred // Event containing the contract specifics and raw log

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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
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
// Solidity: e Pause()
func (_ValidatorMgr *ValidatorMgrFilterer) FilterPause(opts *bind.FilterOpts) (*ValidatorMgrPauseIterator, error) {

	logs, sub, err := _ValidatorMgr.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrPauseIterator{contract: _ValidatorMgr.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
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

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
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
// Solidity: e Unpause()
func (_ValidatorMgr *ValidatorMgrFilterer) FilterUnpause(opts *bind.FilterOpts) (*ValidatorMgrUnpauseIterator, error) {

	logs, sub, err := _ValidatorMgr.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ValidatorMgrUnpauseIterator{contract: _ValidatorMgr.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
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
