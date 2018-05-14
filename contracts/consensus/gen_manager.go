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

// ValidatorManagerABI is the input ABI used to generate the binding from.
const ValidatorManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumDeposit\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDepositAtIndex\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"availableAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenReceiver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_registerValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_hasAvailability\",\"outputs\":[{\"name\":\"available\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorsChecksum\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"setBaseDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isGenesisValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getValidatorAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesisValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_baseDeposit\",\"type\":\"uint256\"},{\"name\":\"_maxValidators\",\"type\":\"uint256\"},{\"name\":\"_freezePeriod\",\"type\":\"uint256\"},{\"name\":\"_genesis\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ValidatorManagerBin is the compiled bytecode used for deploying new contracts.
const ValidatorManagerBin = `606060405260008060146101000a81548160ff0219169083151502179055506001600c60006101000a81548160ff02191690831515021790555034156200004557600080fd5b604051608080620021d883398101604052808051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018310151515620000ce57600080fd5b670de0b6b3a764000084026001819055508260028190555062015180820260038190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000155816001546200015f640100000000026200142f176401000000009004565b50505050620005df565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600160078054806001018281620001be919062000521565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff0219169083151502179055508360020180548060010182816200024a919062000550565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115620004765760066000600760018603815481101515620002b557fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002091508160020160018360020180549050038154811015156200033957fe5b906000526020600020906002020190508060000154851115156200035d5762000476565b6007600184038154811015156200037057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600784815481101515620003ac57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856007600185038154811015156200040957fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082826000018190555060018303846000018190555082806001900393505062000294565b620004946200049c64010000000002620019aa176401000000009004565b505050505050565b60076040518082805480156200050857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620004bd575b5050915050604051809103902060058160001916905550565b8154818355818115116200054b578183600052602060002091820191016200054a919062000585565b5b505050565b81548183558181151162000580576002028160020283600052602060002091820191016200057f9190620005ad565b5b505050565b620005aa91905b80821115620005a65760008160009055506001016200058c565b5090565b90565b620005dc91905b80821115620005d857600080820160009055600182016000905550600201620005b4565b5090565b90565b611be980620005ef6000396000f300606060405260043610610149576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063035cf1421461014e57806308ac5256146101775780630a3cb663146101a05780633ed0a373146101c95780633f4ba83a14610207578063592fdb8e1461021c5780635c975abb146102a1578063671b4d49146102ce57806369474625146102e35780636a911ccf1461030c5780637071688a146103215780638456cb591461034a5780638da5cb5b1461035f5780639363a141146103b457806397584b3e146103dd5780639bb2ea5a1461040a578063aded41ec1461042d578063b774cb1e14610442578063c0ee0b8a14610473578063c22a933c146104f8578063cefddda91461051b578063e7a60a9c1461056c578063f2fde38b146105d6578063f963aeea1461060f578063facd743b14610664575b600080fd5b341561015957600080fd5b6101616106b5565b6040518082815260200191505060405180910390f35b341561018257600080fd5b61018a610789565b6040518082815260200191505060405180910390f35b34156101ab57600080fd5b6101b361078f565b6040518082815260200191505060405180910390f35b34156101d457600080fd5b6101ea6004808035906020019091905050610795565b604051808381526020018281526020019250505060405180910390f35b341561021257600080fd5b61021a610820565b005b341561022757600080fd5b61029f600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506108de565b005b34156102ac57600080fd5b6102b4610c38565b604051808215151515815260200191505060405180910390f35b34156102d957600080fd5b6102e1610c4b565b005b34156102ee57600080fd5b6102f6610cfa565b6040518082815260200191505060405180910390f35b341561031757600080fd5b61031f610d00565b005b341561032c57600080fd5b610334610d3b565b6040518082815260200191505060405180910390f35b341561035557600080fd5b61035d610d48565b005b341561036a57600080fd5b610372610e08565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103bf57600080fd5b6103c7610e2d565b6040518082815260200191505060405180910390f35b34156103e857600080fd5b6103f0610e7a565b604051808215151515815260200191505060405180910390f35b341561041557600080fd5b61042b6004808035906020019091905050610e8d565b005b341561043857600080fd5b610440610f31565b005b341561044d57600080fd5b61045561108c565b60405180826000191660001916815260200191505060405180910390f35b341561047e57600080fd5b6104f6600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050611092565b005b341561050357600080fd5b6105196004808035906020019091905050611097565b005b341561052657600080fd5b610552600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506110fc565b604051808215151515815260200191505060405180910390f35b341561057757600080fd5b61058d6004808035906020019091905050611156565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34156105e157600080fd5b61060d600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061120e565b005b341561061a57600080fd5b610622611363565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561066f57600080fd5b61069b600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611389565b604051808215151515815260200191505060405180910390f35b6000806106c0610e7a565b156106cf576001549150610785565b6006600060076001600780549050038154811015156106ea57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600201600183600201805490500381548110151561076f57fe5b9060005260206000209060020201600001540191505b5090565b60025481565b60035481565b6000806000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201848154811015156107e957fe5b90600052602060002090600202019050670de0b6b3a7640000816000015481151561081057fe5b0481600101549250925050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561087b57600080fd5b600060149054906101000a900460ff16151561089657600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060188260008151811015156108f157fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900463ffffffff169060020a02601083600181518110151561097b57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900463ffffffff169060020a026008846002815181101515610a0557fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900463ffffffff169060020a02846003815181101515610a8d57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900401010190506080604051908101604052808573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838152602001827c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815250600860008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002019080519060200190610be5929190611a2d565b5060608201518160030160006101000a81548163ffffffff02191690837c010000000000000000000000000000000000000000000000000000000090040217905550905050610c32610c4b565b50505050565b600060149054906101000a900460ff1681565b600c60009054906101000a900460ff161515610c6657600080fd5b600060149054906101000a900460ff16151515610c8257600080fd5b610c8b33611389565b151515610c9757600080fd5b610c9f6106b5565b3410151515610cad57600080fd5b610cb5610e7a565b1515610cc457610cc36113e2565b5b610cf8600860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660086001015461142f565b565b60015481565b600060149054906101000a900460ff16151515610d1c57600080fd5b610d2533611389565b1515610d3057600080fd5b610d3933611749565b565b6000600780549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610da357600080fd5b600060149054906101000a900460ff16151515610dbf57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020180549050905090565b6000806007805490506002540311905090565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610eeb57600080fd5b600780549050831015610f255782600780549050039150600090505b81811015610f2457610f176113e2565b8080600101915050610f07565b5b82600281905550505050565b60008060008060149054906101000a900460ff16151515610f5157600080fd5b6000925060009150600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190505b808054905082108015610fd1575060008183815481101515610fbc57fe5b90600052602060002090600202016001015414155b15611033578082815481101515610fe457fe5b90600052602060002090600202016001015442101561100257611033565b808281548110151561101057fe5b906000526020600020906002020160000154830192508180600101925050610f9e565b61103d33836118bd565b6000831115611087573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561108657600080fd5b5b505050565b60055481565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156110f257600080fd5b8060018190555050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b600080600060078481548110151561116a57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020160018260020180549050038154811015156111f457fe5b906000526020600020906002020160000154915050915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561126957600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156112a557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b61142d60076001600780549050038154811015156113fc57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611749565b565b600080600080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020935060016007805480600101828161148c9190611aad565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003846000018190555060018460010160006101000a81548160ff0219169083151502179055508360020180548060010182816115169190611ad9565b916000526020600020906002020160006040805190810160405280898152602001600081525090919091506000820151816000015560208201518160010155505050836000015492505b6000831115611739576006600060076001860381548110151561157f57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020915081600201600183600201805490500381548110151561160257fe5b9060005260206000209060020201905080600001548511151561162457611739565b60076001840381548110151561163657fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660078481548110151561167157fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856007600185038154811015156116cd57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828260000181905550600183038460000181905550828060019003935050611560565b6117416119aa565b505050505050565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150816000015490505b600160078054905003811015611848576007600182018154811015156117b757fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007828154811015156117f257fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050611795565b600780548091906001900361185d9190611b0b565b5060008260010160006101000a81548160ff021916908315150217905550600354420182600201600184600201805490500381548110151561189b57fe5b9060005260206000209060020201600101819055506118b86119aa565b505050565b6000806000808414156118cf576119a3565b600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600091508390505b826002018054905081101561199157826002018181548110151561193857fe5b9060005260206000209060020201836002018381548110151561195757fe5b9060005260206000209060020201600082015481600001556001820154816001015590505081806001019250508080600101915050611918565b8183600201816119a19190611b37565b505b5050505050565b6007604051808280548015611a1457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116119ca575b5050915050604051809103902060058160001916905550565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a6e57805160ff1916838001178555611a9c565b82800160010185558215611a9c579182015b82811115611a9b578251825591602001919060010190611a80565b5b509050611aa99190611b69565b5090565b815481835581811511611ad457818360005260206000209182019101611ad39190611b69565b5b505050565b815481835581811511611b0657600202816002028360005260206000209182019101611b059190611b8e565b5b505050565b815481835581811511611b3257818360005260206000209182019101611b319190611b69565b5b505050565b815481835581811511611b6457600202816002028360005260206000209182019101611b639190611b8e565b5b505050565b611b8b91905b80821115611b87576000816000905550600101611b6f565b5090565b90565b611bba91905b80821115611bb657600080820160009055600182016000905550600201611b94565b5090565b905600a165627a7a72305820b100c95c986827870ca505173732fac2a6dabe23188f3460ba9c91e8e360ed4c0029`

// DeployValidatorManager deploys a new Ethereum contract, binding an instance of ValidatorManager to it.
func DeployValidatorManager(auth *bind.TransactOpts, backend bind.ContractBackend, _baseDeposit *big.Int, _maxValidators *big.Int, _freezePeriod *big.Int, _genesis common.Address) (common.Address, *types.Transaction, *ValidatorManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorManagerBin), backend, _baseDeposit, _maxValidators, _freezePeriod, _genesis)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorManager{ValidatorManagerCaller: ValidatorManagerCaller{contract: contract}, ValidatorManagerTransactor: ValidatorManagerTransactor{contract: contract}}, nil
}

// ValidatorManager is an auto generated Go binding around an Ethereum contract.
type ValidatorManager struct {
	ValidatorManagerCaller     // Read-only binding to the contract
	ValidatorManagerTransactor // Write-only binding to the contract
}

// ValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorManagerSession struct {
	Contract     *ValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorManagerCallerSession struct {
	Contract *ValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorManagerTransactorSession struct {
	Contract     *ValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorManagerRaw struct {
	Contract *ValidatorManager // Generic contract binding to access the raw methods on
}

// ValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorManagerCallerRaw struct {
	Contract *ValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactorRaw struct {
	Contract *ValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorManager creates a new instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManager(address common.Address, backend bind.ContractBackend) (*ValidatorManager, error) {
	contract, err := bindValidatorManager(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorManager{ValidatorManagerCaller: ValidatorManagerCaller{contract: contract}, ValidatorManagerTransactor: ValidatorManagerTransactor{contract: contract}}, nil
}

// NewValidatorManagerCaller creates a new read-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*ValidatorManagerCaller, error) {
	contract, err := bindValidatorManager(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerCaller{contract: contract}, nil
}

// NewValidatorManagerTransactor creates a new write-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorManagerTransactor, error) {
	contract, err := bindValidatorManager(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerTransactor{contract: contract}, nil
}

// bindValidatorManager binds a generic wrapper to an already deployed contract.
func bindValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.ValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorManager *ValidatorManagerCaller) _hasAvailability(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "_hasAvailability")
	return *ret0, err
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorManager *ValidatorManagerSession) _hasAvailability() (bool, error) {
	return _ValidatorManager.Contract._hasAvailability(&_ValidatorManager.CallOpts)
}

// _hasAvailability is a free data retrieval call binding the contract method 0x97584b3e.
//
// Solidity: function _hasAvailability() constant returns(available bool)
func (_ValidatorManager *ValidatorManagerCallerSession) _hasAvailability() (bool, error) {
	return _ValidatorManager.Contract._hasAvailability(&_ValidatorManager.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerCaller) BaseDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "baseDeposit")
	return *ret0, err
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerSession) BaseDeposit() (*big.Int, error) {
	return _ValidatorManager.Contract.BaseDeposit(&_ValidatorManager.CallOpts)
}

// BaseDeposit is a free data retrieval call binding the contract method 0x69474625.
//
// Solidity: function baseDeposit() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) BaseDeposit() (*big.Int, error) {
	return _ValidatorManager.Contract.BaseDeposit(&_ValidatorManager.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerCaller) FreezePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "freezePeriod")
	return *ret0, err
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerSession) FreezePeriod() (*big.Int, error) {
	return _ValidatorManager.Contract.FreezePeriod(&_ValidatorManager.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) FreezePeriod() (*big.Int, error) {
	return _ValidatorManager.Contract.FreezePeriod(&_ValidatorManager.CallOpts)
}

// GenesisValidator is a free data retrieval call binding the contract method 0xf963aeea.
//
// Solidity: function genesisValidator() constant returns(address)
func (_ValidatorManager *ValidatorManagerCaller) GenesisValidator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "genesisValidator")
	return *ret0, err
}

// GenesisValidator is a free data retrieval call binding the contract method 0xf963aeea.
//
// Solidity: function genesisValidator() constant returns(address)
func (_ValidatorManager *ValidatorManagerSession) GenesisValidator() (common.Address, error) {
	return _ValidatorManager.Contract.GenesisValidator(&_ValidatorManager.CallOpts)
}

// GenesisValidator is a free data retrieval call binding the contract method 0xf963aeea.
//
// Solidity: function genesisValidator() constant returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) GenesisValidator() (common.Address, error) {
	return _ValidatorManager.Contract.GenesisValidator(&_ValidatorManager.CallOpts)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ValidatorManager *ValidatorManagerCaller) GetDepositAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	ret := new(struct {
		Amount      *big.Int
		AvailableAt *big.Int
	})
	out := ret
	err := _ValidatorManager.contract.Call(opts, out, "getDepositAtIndex", index)
	return *ret, err
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ValidatorManager *ValidatorManagerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	return _ValidatorManager.Contract.GetDepositAtIndex(&_ValidatorManager.CallOpts, index)
}

// GetDepositAtIndex is a free data retrieval call binding the contract method 0x3ed0a373.
//
// Solidity: function getDepositAtIndex(index uint256) constant returns(amount uint256, availableAt uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) GetDepositAtIndex(index *big.Int) (struct {
	Amount      *big.Int
	AvailableAt *big.Int
}, error) {
	return _ValidatorManager.Contract.GetDepositAtIndex(&_ValidatorManager.CallOpts, index)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ValidatorManager *ValidatorManagerCaller) GetDepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "getDepositCount")
	return *ret0, err
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ValidatorManager *ValidatorManagerSession) GetDepositCount() (*big.Int, error) {
	return _ValidatorManager.Contract.GetDepositCount(&_ValidatorManager.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x9363a141.
//
// Solidity: function getDepositCount() constant returns(count uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) GetDepositCount() (*big.Int, error) {
	return _ValidatorManager.Contract.GetDepositCount(&_ValidatorManager.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ValidatorManager *ValidatorManagerCaller) GetMinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "getMinimumDeposit")
	return *ret0, err
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ValidatorManager *ValidatorManagerSession) GetMinimumDeposit() (*big.Int, error) {
	return _ValidatorManager.Contract.GetMinimumDeposit(&_ValidatorManager.CallOpts)
}

// GetMinimumDeposit is a free data retrieval call binding the contract method 0x035cf142.
//
// Solidity: function getMinimumDeposit() constant returns(deposit uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) GetMinimumDeposit() (*big.Int, error) {
	return _ValidatorManager.Contract.GetMinimumDeposit(&_ValidatorManager.CallOpts)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ValidatorManager *ValidatorManagerCaller) GetValidatorAtIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	ret := new(struct {
		Code    common.Address
		Deposit *big.Int
	})
	out := ret
	err := _ValidatorManager.contract.Call(opts, out, "getValidatorAtIndex", index)
	return *ret, err
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ValidatorManager *ValidatorManagerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ValidatorManager.Contract.GetValidatorAtIndex(&_ValidatorManager.CallOpts, index)
}

// GetValidatorAtIndex is a free data retrieval call binding the contract method 0xe7a60a9c.
//
// Solidity: function getValidatorAtIndex(index uint256) constant returns(code address, deposit uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) GetValidatorAtIndex(index *big.Int) (struct {
	Code    common.Address
	Deposit *big.Int
}, error) {
	return _ValidatorManager.Contract.GetValidatorAtIndex(&_ValidatorManager.CallOpts, index)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ValidatorManager *ValidatorManagerCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "getValidatorCount")
	return *ret0, err
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ValidatorManager *ValidatorManagerSession) GetValidatorCount() (*big.Int, error) {
	return _ValidatorManager.Contract.GetValidatorCount(&_ValidatorManager.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() constant returns(count uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) GetValidatorCount() (*big.Int, error) {
	return _ValidatorManager.Contract.GetValidatorCount(&_ValidatorManager.CallOpts)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ValidatorManager *ValidatorManagerCaller) IsGenesisValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "isGenesisValidator", code)
	return *ret0, err
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ValidatorManager *ValidatorManagerSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsGenesisValidator(&_ValidatorManager.CallOpts, code)
}

// IsGenesisValidator is a free data retrieval call binding the contract method 0xcefddda9.
//
// Solidity: function isGenesisValidator(code address) constant returns(isIndeed bool)
func (_ValidatorManager *ValidatorManagerCallerSession) IsGenesisValidator(code common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsGenesisValidator(&_ValidatorManager.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ValidatorManager *ValidatorManagerCaller) IsValidator(opts *bind.CallOpts, code common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "isValidator", code)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ValidatorManager *ValidatorManagerSession) IsValidator(code common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsValidator(&_ValidatorManager.CallOpts, code)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(code address) constant returns(isIndeed bool)
func (_ValidatorManager *ValidatorManagerCallerSession) IsValidator(code common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsValidator(&_ValidatorManager.CallOpts, code)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerCaller) MaxValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "maxValidators")
	return *ret0, err
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerSession) MaxValidators() (*big.Int, error) {
	return _ValidatorManager.Contract.MaxValidators(&_ValidatorManager.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() constant returns(uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) MaxValidators() (*big.Int, error) {
	return _ValidatorManager.Contract.MaxValidators(&_ValidatorManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidatorManager *ValidatorManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidatorManager *ValidatorManagerSession) Owner() (common.Address, error) {
	return _ValidatorManager.Contract.Owner(&_ValidatorManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) Owner() (common.Address, error) {
	return _ValidatorManager.Contract.Owner(&_ValidatorManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ValidatorManager *ValidatorManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ValidatorManager *ValidatorManagerSession) Paused() (bool, error) {
	return _ValidatorManager.Contract.Paused(&_ValidatorManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ValidatorManager *ValidatorManagerCallerSession) Paused() (bool, error) {
	return _ValidatorManager.Contract.Paused(&_ValidatorManager.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ValidatorManager *ValidatorManagerCaller) ValidatorsChecksum(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ValidatorManager.contract.Call(opts, out, "validatorsChecksum")
	return *ret0, err
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ValidatorManager *ValidatorManagerSession) ValidatorsChecksum() ([32]byte, error) {
	return _ValidatorManager.Contract.ValidatorsChecksum(&_ValidatorManager.CallOpts)
}

// ValidatorsChecksum is a free data retrieval call binding the contract method 0xb774cb1e.
//
// Solidity: function validatorsChecksum() constant returns(bytes32)
func (_ValidatorManager *ValidatorManagerCallerSession) ValidatorsChecksum() ([32]byte, error) {
	return _ValidatorManager.Contract.ValidatorsChecksum(&_ValidatorManager.CallOpts)
}

// _registerValidator is a paid mutator transaction binding the contract method 0x671b4d49.
//
// Solidity: function _registerValidator() returns()
func (_ValidatorManager *ValidatorManagerTransactor) _registerValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "_registerValidator")
}

// _registerValidator is a paid mutator transaction binding the contract method 0x671b4d49.
//
// Solidity: function _registerValidator() returns()
func (_ValidatorManager *ValidatorManagerSession) _registerValidator() (*types.Transaction, error) {
	return _ValidatorManager.Contract._registerValidator(&_ValidatorManager.TransactOpts)
}

// _registerValidator is a paid mutator transaction binding the contract method 0x671b4d49.
//
// Solidity: function _registerValidator() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) _registerValidator() (*types.Transaction, error) {
	return _ValidatorManager.Contract._registerValidator(&_ValidatorManager.TransactOpts)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x6a911ccf.
//
// Solidity: function deregisterValidator() returns()
func (_ValidatorManager *ValidatorManagerTransactor) DeregisterValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "deregisterValidator")
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x6a911ccf.
//
// Solidity: function deregisterValidator() returns()
func (_ValidatorManager *ValidatorManagerSession) DeregisterValidator() (*types.Transaction, error) {
	return _ValidatorManager.Contract.DeregisterValidator(&_ValidatorManager.TransactOpts)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x6a911ccf.
//
// Solidity: function deregisterValidator() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) DeregisterValidator() (*types.Transaction, error) {
	return _ValidatorManager.Contract.DeregisterValidator(&_ValidatorManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ValidatorManager *ValidatorManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ValidatorManager *ValidatorManagerSession) Pause() (*types.Transaction, error) {
	return _ValidatorManager.Contract.Pause(&_ValidatorManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _ValidatorManager.Contract.Pause(&_ValidatorManager.TransactOpts)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_ValidatorManager *ValidatorManagerTransactor) ReleaseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "releaseDeposits")
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_ValidatorManager *ValidatorManagerSession) ReleaseDeposits() (*types.Transaction, error) {
	return _ValidatorManager.Contract.ReleaseDeposits(&_ValidatorManager.TransactOpts)
}

// ReleaseDeposits is a paid mutator transaction binding the contract method 0xaded41ec.
//
// Solidity: function releaseDeposits() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) ReleaseDeposits() (*types.Transaction, error) {
	return _ValidatorManager.Contract.ReleaseDeposits(&_ValidatorManager.TransactOpts)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ValidatorManager *ValidatorManagerTransactor) SetBaseDeposit(opts *bind.TransactOpts, deposit *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "setBaseDeposit", deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ValidatorManager *ValidatorManagerSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetBaseDeposit(&_ValidatorManager.TransactOpts, deposit)
}

// SetBaseDeposit is a paid mutator transaction binding the contract method 0xc22a933c.
//
// Solidity: function setBaseDeposit(deposit uint256) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) SetBaseDeposit(deposit *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetBaseDeposit(&_ValidatorManager.TransactOpts, deposit)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ValidatorManager *ValidatorManagerTransactor) SetMaxValidators(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "setMaxValidators", max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ValidatorManager *ValidatorManagerSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetMaxValidators(&_ValidatorManager.TransactOpts, max)
}

// SetMaxValidators is a paid mutator transaction binding the contract method 0x9bb2ea5a.
//
// Solidity: function setMaxValidators(max uint256) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) SetMaxValidators(max *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.SetMaxValidators(&_ValidatorManager.TransactOpts, max)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ValidatorManager *ValidatorManagerTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ValidatorManager *ValidatorManagerSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TokenFallback(&_ValidatorManager.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(_from address, _value uint256, _data bytes) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TokenFallback(&_ValidatorManager.TransactOpts, _from, _value, _data)
}

// TokenReceiver is a paid mutator transaction binding the contract method 0x592fdb8e.
//
// Solidity: function tokenReceiver(_from address, _value uint256, _data bytes) returns()
func (_ValidatorManager *ValidatorManagerTransactor) TokenReceiver(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "tokenReceiver", _from, _value, _data)
}

// TokenReceiver is a paid mutator transaction binding the contract method 0x592fdb8e.
//
// Solidity: function tokenReceiver(_from address, _value uint256, _data bytes) returns()
func (_ValidatorManager *ValidatorManagerSession) TokenReceiver(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TokenReceiver(&_ValidatorManager.TransactOpts, _from, _value, _data)
}

// TokenReceiver is a paid mutator transaction binding the contract method 0x592fdb8e.
//
// Solidity: function tokenReceiver(_from address, _value uint256, _data bytes) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) TokenReceiver(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TokenReceiver(&_ValidatorManager.TransactOpts, _from, _value, _data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ValidatorManager *ValidatorManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ValidatorManager *ValidatorManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TransferOwnership(&_ValidatorManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.TransferOwnership(&_ValidatorManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ValidatorManager *ValidatorManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ValidatorManager *ValidatorManagerSession) Unpause() (*types.Transaction, error) {
	return _ValidatorManager.Contract.Unpause(&_ValidatorManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _ValidatorManager.Contract.Unpause(&_ValidatorManager.TransactOpts)
}
