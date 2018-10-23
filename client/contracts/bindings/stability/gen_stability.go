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
const StabilityABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSubscriptionAtIndex\",\"outputs\":[{\"name\":\"code\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSubscriptionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"name\":\"_priceProviderAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unsubscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"name\":\"_priceProviderAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// StabilitySrcMap is used in order to generate source maps to use when we want to debug bytecode.
const StabilitySrcMap = "{\"contracts\":{\"../../truffle/contracts/stability/PriceProvider.sol:PriceProvider\":{\"bin-runtime\":\"\",\"srcmap-runtime\":\"\"},\"../../truffle/contracts/stability/Stability.sol:Stability\":{\"bin-runtime\":\"6080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063158ef93e146100ca5780633f4ba83a146100f9578063402300461461011057806341b3d185146101845780635c975abb146101af57806366419970146101de578063715018a6146102095780638456cb59146102205780638da5cb5b146102375780638f449a051461028e578063da35a26f14610298578063f2fde38b146102e5578063fcae448414610328575b600080fd5b3480156100d657600080fd5b506100df61033f565b604051808215151515815260200191505060405180910390f35b34801561010557600080fd5b5061010e610352565b005b34801561011c57600080fd5b5061013b60048036038101908080359060200190929190505050610410565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561019057600080fd5b5061019961049f565b6040518082815260200191505060405180910390f35b3480156101bb57600080fd5b506101c46104a5565b604051808215151515815260200191505060405180910390f35b3480156101ea57600080fd5b506101f36104b8565b6040518082815260200191505060405180910390f35b34801561021557600080fd5b5061021e6104c5565b005b34801561022c57600080fd5b506102356105c7565b005b34801561024357600080fd5b5061024c610687565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102966106ac565b005b3480156102a457600080fd5b506102e360048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061073b565b005b3480156102f157600080fd5b50610326600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061084d565b005b34801561033457600080fd5b5061033d6108b4565b005b600060159054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103ad57600080fd5b600060149054906101000a900460ff1615156103c857600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600080600060048481548110151561042457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020154915050915091565b60015481565b600060149054906101000a900460ff1681565b6000600480549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052057600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561062257600080fd5b600060149054906101000a900460ff1615151561063e57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff161515156106c957600080fd5b6106d233610b9b565b1561072f57600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050348160020160008282540192505081905550610738565b610737610bf4565b5b50565b600060159054906101000a900460ff161515156107e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b8160018190555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600060156101000a81548160ff0219169083151502179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108a857600080fd5b6108b181610ce1565b50565b60008060006108c233610b9b565b15156108cd57600080fd5b670de0b6b3a7640000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a035b1fe6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561095c57600080fd5b505af1158015610970573d6000803e3d6000fd5b505050506040513d602081101561098657600080fd5b8101908080519060200190929190505050101515156109a457600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250826000015491503373ffffffffffffffffffffffffffffffffffffffff166108fc84600201549081150290604051600060405180830381858888f19350505050158015610a36573d6000803e3d6000fd5b50600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090556001820160006101000a81549060ff0219169055600282016000905550506004600160048054905003815481101515610ab457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600483815481101515610af157fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506004805480919060019003610b959190610ddb565b50505050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b60006001543410151515610c0757600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600160043390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003816000018190555060018160010160006101000a81548160ff02191690831515021790555034816002018190555050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610d1d57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b815481835581811115610e0257818360005260206000209182019101610e019190610e07565b5b505050565b610e2991905b80821115610e25576000816000905550600101610e0d565b5090565b905600a165627a7a7230582003e9f8a667889c71192d8deb1e6aa5493e226a18e8f0762c904236c8053c7c100029\",\"srcmap-runtime\":\"241:3144:1:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;562:23:4;;8:9:-1;5:2;;;30:1;27;20:12;5:2;562:23:4;;;;;;;;;;;;;;;;;;;;;;;;;;;838:92:2;;8:9:-1;5:2;;;30:1;27;20:12;5:2;838:92:2;;;;;;1784:225:1;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1784:225:1;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;331:22;;8:9:-1;5:2;;;30:1;27;20:12;5:2;331:22:1;;;;;;;;;;;;;;;;;;;;;;;247:26:2;;8:9:-1;5:2;;;30:1;27;20:12;5:2;247:26:2;;;;;;;;;;;;;;;;;;;;;;;;;;;1669:109:1;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1669:109:1;;;;;;;;;;;;;;;;;;;;;;;1001:111:3;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1001:111:3;;;;;;666:90:2;;8:9:-1;5:2;;;30:1;27;20:12;5:2;666:90:2;;;;;;238:20:3;;8:9:-1;5:2;;;30:1;27;20:12;5:2;238:20:3;;;;;;;;;;;;;;;;;;;;;;;;;;;2506:276:1;;;;;;1476:187;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1476:187:1;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1274:103:3;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1274:103:3;;;;;;;;;;;;;;;;;;;;;;;;;;;;2840:543:1;;8:9:-1;5:2;;;30:1;27;20:12;5:2;2840:543:1;;;;;;562:23:4;;;;;;;;;;;;;:::o;838:92:2:-;719:5:3;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;568:6:2;;;;;;;;;;;560:15;;;;;;;;900:5;891:6;;:14;;;;;;;;;;;;;;;;;;916:9;;;;;;;;;;838:92::o;1784:225:1:-;1849:12;1863;1924:17;1894:13;1908:5;1894:20;;;;;;;;;;;;;;;;;;;;;;;;;;;1887:27;;1944:20;:26;1965:4;1944:26;;;;;;;;;;;;;;;1924:46;;1990:4;:12;;;1980:22;;1784:225;;;;:::o;331:22::-;;;;:::o;247:26:2:-;;;;;;;;;;;;;:::o;1669:109:1:-;1722:10;1751:13;:20;;;;1744:27;;1669:109;:::o;1001:111:3:-;719:5;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;1077:5;;;;;;;;;;;1058:25;;;;;;;;;;;;1105:1;1089:5;;:18;;;;;;;;;;;;;;;;;;1001:111::o;666:90:2:-;719:5:3;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;416:6:2;;;;;;;;;;;415:7;407:16;;;;;;;;729:4;720:6;;:13;;;;;;;;;;;;;;;;;;744:7;;;;;;;;;;666:90::o;238:20:3:-;;;;;;;;;;;;;:::o;2506:276:1:-;2614:17;416:6:2;;;;;;;;;;;415:7;407:16;;;;;;;;2570:28:1;2587:10;2570:16;:28::i;:::-;2566:170;;;2634:20;:32;2655:10;2634:32;;;;;;;;;;;;;;;2614:52;;2696:9;2680:4;:12;;;:25;;;;;;;;;;;2719:7;;2566:170;2754:21;:19;:21::i;:::-;429:1:2;2506:276:1;:::o;1476:187::-;714:11:4;;;;;;;;;;;713:12;705:71;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1586:11:1;1573:10;:24;;;;1637:18;1607:13;;:49;;;;;;;;;;;;;;;;;;803:4:4;789:11;;:18;;;;;;;;;;;;;;;;;;1476:187:1;;:::o;1274:103:3:-;719:5;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;1343:29;1362:9;1343:18;:29::i;:::-;1274:103;:::o;2840:543:1:-;2920:17;2982:16;3177:17;651:28;668:10;651:16;:28::i;:::-;643:37;;;;;;;;313:7;850:13;;;;;;;;;;;:19;;;:21;;;;;;;;;;;;;;;;;;;;;;;8:9:-1;5:2;;;30:1;27;20:12;5:2;850:21:1;;;;8:9:-1;5:2;;;45:16;42:1;39;24:38;77:16;74:1;67:27;5:2;850:21:1;;;;;;;13:2:-1;8:3;5:11;2:2;;;29:1;26;19:12;2:2;850:21:1;;;;;;;;;;;;;;;;:28;;842:37;;;;;;;;2940:20;:32;2961:10;2940:32;;;;;;;;;;;;;;;2920:52;;3001:4;:10;;;2982:29;;3021:10;:19;;:33;3041:4;:12;;;3021:33;;;;;;;;;;;;;;;;;;;;;;;;8:9:-1;5:2;;;45:16;42:1;39;24:38;77:16;74:1;67:27;5:2;3021:33:1;3071:20;:32;3092:10;3071:32;;;;;;;;;;;;;;;;3064:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;3197:13;3232:1;3211:13;:20;;;;:22;3197:37;;;;;;;;;;;;;;;;;;;;;;;;;;;3177:57;;3274:9;3245:13;3259:11;3245:26;;;;;;;;;;;;;;;;;;:38;;;;;;;;;;;;;;;;;;3333:11;3293:20;:31;3314:9;3293:31;;;;;;;;;;;;;;;:37;;:51;;;;3354:13;:22;;;;;;;;;;;;:::i;:::-;;2840:543;;;:::o;2015:151::-;2081:13;2113:20;:30;2134:8;2113:30;;;;;;;;;;;;;;;:46;;;;;;;;;;;;2106:53;;2015:151;;;:::o;2172:255::-;2240:17;763:10;;750:9;:23;;742:32;;;;;;;;2260:20;:32;2281:10;2260:32;;;;;;;;;;;;;;;2240:52;;2348:1;2315:13;2334:10;2315:30;;39:1:-1;33:3;27:10;23:18;57:10;52:3;45:23;79:10;72:17;;0:93;2315:30:1;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:34;2302:4;:10;;:47;;;;2382:4;2359;:20;;;:27;;;;;;;;;;;;;;;;;;2411:9;2396:4;:12;;:24;;;;2172:255;:::o;1512:171:3:-;1603:1;1582:23;;:9;:23;;;;1574:32;;;;;;;;1645:9;1617:38;;1638:5;;;;;;;;;;;1617:38;;;;;;;;;;;;1669:9;1661:5;;:17;;;;;;;;;;;;;;;;;;1512:171;:::o;241:3144:1:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;:::o\"},\"../../truffle/node_modules/openzeppelin-solidity/contracts/lifecycle/Pausable.sol:Pausable\":{\"bin-runtime\":\"608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633f4ba83a1461007d5780635c975abb14610094578063715018a6146100c35780638456cb59146100da5780638da5cb5b146100f1578063f2fde38b14610148575b600080fd5b34801561008957600080fd5b5061009261018b565b005b3480156100a057600080fd5b506100a9610249565b604051808215151515815260200191505060405180910390f35b3480156100cf57600080fd5b506100d861025c565b005b3480156100e657600080fd5b506100ef61035e565b005b3480156100fd57600080fd5b5061010661041e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561015457600080fd5b50610189600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610443565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156101e657600080fd5b600060149054906101000a900460ff16151561020157600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600060149054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102b757600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103b957600080fd5b600060149054906101000a900460ff161515156103d557600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561049e57600080fd5b6104a7816104aa565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156104e657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582040f5a2c18af98ccde0dd9369624e6d6a3e910a7a7cceea5f5577830dacd860180029\",\"srcmap-runtime\":\"177:755:2:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;838:92;;8:9:-1;5:2;;;30:1;27;20:12;5:2;838:92:2;;;;;;247:26;;8:9:-1;5:2;;;30:1;27;20:12;5:2;247:26:2;;;;;;;;;;;;;;;;;;;;;;;;;;;1001:111:3;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1001:111:3;;;;;;666:90:2;;8:9:-1;5:2;;;30:1;27;20:12;5:2;666:90:2;;;;;;238:20:3;;8:9:-1;5:2;;;30:1;27;20:12;5:2;238:20:3;;;;;;;;;;;;;;;;;;;;;;;;;;;1274:103;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1274:103:3;;;;;;;;;;;;;;;;;;;;;;;;;;;;838:92:2;719:5:3;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;568:6:2;;;;;;;;;;;560:15;;;;;;;;900:5;891:6;;:14;;;;;;;;;;;;;;;;;;916:9;;;;;;;;;;838:92::o;247:26::-;;;;;;;;;;;;;:::o;1001:111:3:-;719:5;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;1077:5;;;;;;;;;;;1058:25;;;;;;;;;;;;1105:1;1089:5;;:18;;;;;;;;;;;;;;;;;;1001:111::o;666:90:2:-;719:5:3;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;416:6:2;;;;;;;;;;;415:7;407:16;;;;;;;;729:4;720:6;;:13;;;;;;;;;;;;;;;;;;744:7;;;;;;;;;;666:90::o;238:20:3:-;;;;;;;;;;;;;:::o;1274:103::-;719:5;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;1343:29;1362:9;1343:18;:29::i;:::-;1274:103;:::o;1512:171::-;1603:1;1582:23;;:9;:23;;;;1574:32;;;;;;;;1645:9;1617:38;;1638:5;;;;;;;;;;;1617:38;;;;;;;;;;;;1669:9;1661:5;;:17;;;;;;;;;;;;;;;;;;1512:171;:::o\"},\"../../truffle/node_modules/openzeppelin-solidity/contracts/ownership/Ownable.sol:Ownable\":{\"bin-runtime\":\"608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063715018a61461005c5780638da5cb5b14610073578063f2fde38b146100ca575b600080fd5b34801561006857600080fd5b5061007161010d565b005b34801561007f57600080fd5b5061008861020f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100d657600080fd5b5061010b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610234565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561016857600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561028f57600080fd5b6102988161029b565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156102d757600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582013733afa07283c4cc5ea5981b752bbacf0cdbc441b4f06b064de9b0379f8f82a0029\",\"srcmap-runtime\":\"217:1468:3:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1001:111;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1001:111:3;;;;;;238:20;;8:9:-1;5:2;;;30:1;27;20:12;5:2;238:20:3;;;;;;;;;;;;;;;;;;;;;;;;;;;1274:103;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1274:103:3;;;;;;;;;;;;;;;;;;;;;;;;;;;;1001:111;719:5;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;1077:5;;;;;;;;;;;1058:25;;;;;;;;;;;;1105:1;1089:5;;:18;;;;;;;;;;;;;;;;;;1001:111::o;238:20::-;;;;;;;;;;;;;:::o;1274:103::-;719:5;;;;;;;;;;;705:19;;:10;:19;;;697:28;;;;;;;;1343:29;1362:9;1343:18;:29::i;:::-;1274:103;:::o;1512:171::-;1603:1;1582:23;;:9;:23;;;;1574:32;;;;;;;;1645:9;1617:38;;1638:5;;;;;;;;;;;1617:38;;;;;;;;;;;;1669:9;1661:5;;:17;;;;;;;;;;;;;;;;;;1512:171;:::o\"},\"../../truffle/node_modules/zos-lib/contracts/migrations/Initializable.sol:Initializable\":{\"bin-runtime\":\"608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063158ef93e146044575b600080fd5b348015604f57600080fd5b5060566070565b604051808215151515815260200191505060405180910390f35b6000809054906101000a900460ff16815600a165627a7a72305820240a09a31dd6de272868e252ab59cc425779f50fdbc3faf839da50e9545268f80029\",\"srcmap-runtime\":\"464:350:4:-;;;;;;;;;;;;;;;;;;;;;;;;562:23;;8:9:-1;5:2;;;30:1;27;20:12;5:2;562:23:4;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::o\"}},\"sourceList\":[\"../../truffle/contracts/stability/PriceProvider.sol\",\"../../truffle/contracts/stability/Stability.sol\",\"../../truffle/node_modules/openzeppelin-solidity/contracts/lifecycle/Pausable.sol\",\"../../truffle/node_modules/openzeppelin-solidity/contracts/ownership/Ownable.sol\",\"../../truffle/node_modules/zos-lib/contracts/migrations/Initializable.sol\"],\"version\":\"0.4.24+commit.e67f0147.Darwin.appleclang\"}"

// StabilityBin is the compiled bytecode used for deploying new contracts.
const StabilityBin = `608060405260008060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b50604051604080610f488339810180604052810190808051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160018190555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610e58806100f06000396000f3006080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063158ef93e146100ca5780633f4ba83a146100f9578063402300461461011057806341b3d185146101845780635c975abb146101af57806366419970146101de578063715018a6146102095780638456cb59146102205780638da5cb5b146102375780638f449a051461028e578063da35a26f14610298578063f2fde38b146102e5578063fcae448414610328575b600080fd5b3480156100d657600080fd5b506100df61033f565b604051808215151515815260200191505060405180910390f35b34801561010557600080fd5b5061010e610352565b005b34801561011c57600080fd5b5061013b60048036038101908080359060200190929190505050610410565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561019057600080fd5b5061019961049f565b6040518082815260200191505060405180910390f35b3480156101bb57600080fd5b506101c46104a5565b604051808215151515815260200191505060405180910390f35b3480156101ea57600080fd5b506101f36104b8565b6040518082815260200191505060405180910390f35b34801561021557600080fd5b5061021e6104c5565b005b34801561022c57600080fd5b506102356105c7565b005b34801561024357600080fd5b5061024c610687565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102966106ac565b005b3480156102a457600080fd5b506102e360048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061073b565b005b3480156102f157600080fd5b50610326600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061084d565b005b34801561033457600080fd5b5061033d6108b4565b005b600060159054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103ad57600080fd5b600060149054906101000a900460ff1615156103c857600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b600080600060048481548110151561042457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020154915050915091565b60015481565b600060149054906101000a900460ff1681565b6000600480549050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052057600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561062257600080fd5b600060149054906101000a900460ff1615151561063e57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff161515156106c957600080fd5b6106d233610b9b565b1561072f57600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050348160020160008282540192505081905550610738565b610737610bf4565b5b50565b600060159054906101000a900460ff161515156107e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b8160018190555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600060156101000a81548160ff0219169083151502179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108a857600080fd5b6108b181610ce1565b50565b60008060006108c233610b9b565b15156108cd57600080fd5b670de0b6b3a7640000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a035b1fe6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561095c57600080fd5b505af1158015610970573d6000803e3d6000fd5b505050506040513d602081101561098657600080fd5b8101908080519060200190929190505050101515156109a457600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250826000015491503373ffffffffffffffffffffffffffffffffffffffff166108fc84600201549081150290604051600060405180830381858888f19350505050158015610a36573d6000803e3d6000fd5b50600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090556001820160006101000a81549060ff0219169055600282016000905550506004600160048054905003815481101515610ab457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600483815481101515610af157fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506004805480919060019003610b959190610ddb565b50505050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff169050919050565b60006001543410151515610c0757600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600160043390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003816000018190555060018160010160006101000a81548160ff02191690831515021790555034816002018190555050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610d1d57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b815481835581811115610e0257818360005260206000209182019101610e019190610e07565b5b505050565b610e2991905b80821115610e25576000816000905550600101610e0d565b5090565b905600a165627a7a7230582003e9f8a667889c71192d8deb1e6aa5493e226a18e8f0762c904236c8053c7c100029`

// DeployStability deploys a new Kowala contract, binding an instance of Stability to it.
func DeployStability(auth *bind.TransactOpts, backend bind.ContractBackend, _minDeposit *big.Int, _priceProviderAddr common.Address) (common.Address, *types.Transaction, *Stability, error) {
	parsed, err := abi.JSON(strings.NewReader(StabilityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StabilityBin), backend, _minDeposit, _priceProviderAddr)
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
