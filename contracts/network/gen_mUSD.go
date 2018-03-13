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

// MusdContractABI is the input ABI used to generate the binding from.
const MusdContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"delegateAddr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maximumSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiverAddr\",\"type\":\"address\"}],\"name\":\"proposeReceiverAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersReceivers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"delegatedFrom\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"availableTo\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ownerAddr\",\"type\":\"address\"}],\"name\":\"addressesOf\",\"outputs\":[{\"name\":\"minerAddr\",\"type\":\"address\"},{\"name\":\"receiveAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownersMiners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ownerAddr\",\"type\":\"address\"},{\"name\":\"miningAddr\",\"type\":\"address\"},{\"name\":\"receiverAddr\",\"type\":\"address\"}],\"name\":\"initializeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ownerAddr\",\"type\":\"address\"}],\"name\":\"acceptMiningAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"delegateAddr\",\"type\":\"address\"}],\"name\":\"delegatedTo\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiverAddr\",\"type\":\"address\"}],\"name\":\"acceptReceiverAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"miningAddr\",\"type\":\"address\"}],\"name\":\"proposeMiningAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"miningAddr\",\"type\":\"address\"},{\"name\":\"receiverAddr\",\"type\":\"address\"}],\"name\":\"proposeAccountAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"availableForDelegationTo\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenHolders\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberTokenHolders\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toAddr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableForDelegation\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ownerAddr\",\"type\":\"address\"},{\"name\":\"receiverAddr\",\"type\":\"address\"}],\"name\":\"acceptAccountAddresses\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegateAddr\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"fromAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"toAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"managementAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"miningAddr\",\"type\":\"address\"}],\"name\":\"NewMiningAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"miningAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"receiverAddr\",\"type\":\"address\"}],\"name\":\"NewReceiverAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegateAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegateAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"OwnershipTransfer\",\"type\":\"event\"}]"

// MusdContractBin is the compiled bytecode used for deploying new contracts.
const MusdContractBin = `60606040526000600455600060055534156200001a57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040805190810160405280600481526020017f6d5553440000000000000000000000000000000000000000000000000000000081525060019080519060200190620000a792919062000107565b506040805190810160405280600481526020017f6d5553440000000000000000000000000000000000000000000000000000000081525060029080519060200190620000f592919062000107565b506340000000600481905550620001b6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200014a57805160ff19168380011785556200017b565b828001600101855582156200017b579182015b828111156200017a5782518255916020019190600101906200015d565b5b5090506200018a91906200018e565b5090565b620001b391905b80821115620001af57600081600090555060010162000195565b5090565b90565b61256780620001c66000396000f300606060405260043610610175576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063026e402b146101855780630480e58b146101df57806306fdde031461020857806308aee770146102965780630a6c3247146102cf57806318160ddd146103485780632c9a9632146103715780632db9395a146103be5780632e1590651461040b578063313ce5671461048457806338a078c4146104b357806340c10f191461055f5780634297d632146105b95780634f422ca9146106325780635254941a146106a957806365da1264146106fa578063665d4040146107475780636847d05f146107985780636a505a42146107d15780636bba590114610829578063923108d91461087657806395cc989c146108d957806395d89b4114610902578063a9059cbb14610990578063b8377644146109ea578063b8a90b0314610a37578063da57537c14610a60578063eac449d914610ad0578063f2fde38b14610b2a575b341561018057600080fd5b600080fd5b341561019057600080fd5b6101c5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610b63565b604051808215151515815260200191505060405180910390f35b34156101ea57600080fd5b6101f2610d8a565b6040518082815260200191505060405180910390f35b341561021357600080fd5b61021b610d90565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561025b578082015181840152602081019050610240565b50505050905090810190601f1680156102885780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102a157600080fd5b6102cd600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610e2e565b005b34156102da57600080fd5b610306600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610f0e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561035357600080fd5b61035b610f41565b6040518082815260200191505060405180910390f35b341561037c57600080fd5b6103a8600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610f47565b6040518082815260200191505060405180910390f35b34156103c957600080fd5b6103f5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610fcd565b6040518082815260200191505060405180910390f35b341561041657600080fd5b610442600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611098565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561048f57600080fd5b6104976110cb565b604051808260ff1660ff16815260200191505060405180910390f35b34156104be57600080fd5b6104ea600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506110de565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b341561056a57600080fd5b61059f600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506111b0565b604051808215151515815260200191505060405180910390f35b34156105c457600080fd5b6105f0600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061132c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561063d57600080fd5b6106a7600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061135f565b005b34156106b457600080fd5b6106e0600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611539565b604051808215151515815260200191505060405180910390f35b341561070557600080fd5b610731600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506117d7565b6040518082815260200191505060405180910390f35b341561075257600080fd5b61077e600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061185d565b604051808215151515815260200191505060405180910390f35b34156107a357600080fd5b6107cf600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611a7d565b005b34156107dc57600080fd5b610827600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611afe565b005b341561083457600080fd5b610860600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611b14565b6040518082815260200191505060405180910390f35b341561088157600080fd5b6108976004808035906020019091905050611b67565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156108e457600080fd5b6108ec611ba6565b6040518082815260200191505060405180910390f35b341561090d57600080fd5b610915611bb3565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561095557808201518184015260208101905061093a565b50505050905090810190601f1680156109825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561099b57600080fd5b6109d0600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050611c51565b604051808215151515815260200191505060405180910390f35b34156109f557600080fd5b610a21600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611e5a565b6040518082815260200191505060405180910390f35b3415610a4257600080fd5b610a4a611e72565b6040518082815260200191505060405180910390f35b3415610a6b57600080fd5b610ab6600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611e82565b604051808215151515815260200191505060405180910390f35b3415610adb57600080fd5b610b10600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050611ea6565b604051808215151515815260200191505060405180910390f35b3415610b3557600080fd5b610b61600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612119565b005b600081600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054031015610bf65760009050610d84565b81600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555081600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555081600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f96eafeca8c3c21ab2fa4a636b93ba20c9e22e3d222d92c6530fedc29a53671ee846040518082815260200191505060405180910390a3600190505b92915050565b60045481565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e265780601f10610dfb57610100808354040283529160200191610e26565b820191906000526020600020905b815481529060010190602001808311610e0957829003601f168201915b505050505081565b600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60106020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60055481565b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401039050919050565b600d6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900460ff1681565b6000806000600e60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250925050915091565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561120d57600080fd5b600454826005540111156112245760009050611326565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415611276576112758361226f565b5b81600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550816005600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a2600190505b92915050565b600e6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113ba57600080fd5b81600e60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60008173ffffffffffffffffffffffffffffffffffffffff16600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156115d857600090506117d2565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905533600e60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fac0b1feded7ffcd64aced5c14411a7d98c9f21dc3dbf7fa2da1451c084d2300a8233604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1600190505b919050565b6000600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60003373ffffffffffffffffffffffffffffffffffffffff16600f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156118fc5760009050611a78565b600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905581601060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507facf4b7670280bd969e5434ff6374c5ab689705c88a93a49244ac8466eb96812a3383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1600190505b919050565b33600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611b0782611a7d565b611b1081610e2e565b5050565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611b5f83610fcd565b039050919050565b600781815481101515611b7657fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600780549050905090565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c495780601f10611c1e57610100808354040283529160200191611c49565b820191906000526020600020905b815481529060010190602001808311611c2c57829003601f168201915b505050505081565b600081600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c9d33610fcd565b031015611cad5760009050611e54565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055506000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415611d9957611d983361231e565b5b81600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415611dea57611de98361226f565b5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b92915050565b60066020528060005260406000206000915090505481565b6000611e7d33611b14565b905090565b6000611e8d83611539565b8015611e9e5750611e9d8261185d565b5b905092915050565b600081600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611ef85760009050612113565b81600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611f855760009050612113565b81600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167faf2be5d3056627fcbd77a887e7ea236a5c437c5781c0c75b1f71cf3fa5cadfc4846040518082815260200191505060405180910390a3600190505b92915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561217457600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f22500af037c600dd7b720644ab6e358635085601d9ac508ad83eb2d6b2d729ca6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150565b6007805480600101828161228391906124be565b9160005260206000209001600083909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600780549050600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205491506000821415612371576124b9565b818060019003925050600760016007805490500381548110151561239157fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806007838154811015156123cf57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018201600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905560078054809190600190036124b791906124ea565b505b505050565b8154818355818115116124e5578183600052602060002091820191016124e49190612516565b5b505050565b815481835581811511612511578183600052602060002091820191016125109190612516565b5b505050565b61253891905b8082111561253457600081600090555060010161251c565b5090565b905600a165627a7a723058207e43253c088e0227063feb0c12c4b832ec84b845660e3c21fe3f90ffd33efcc20029`

// DeployMusdContract deploys a new Ethereum contract, binding an instance of MusdContract to it.
func DeployMusdContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MusdContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MusdContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MusdContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MusdContract{MusdContractCaller: MusdContractCaller{contract: contract}, MusdContractTransactor: MusdContractTransactor{contract: contract}}, nil
}

// MusdContract is an auto generated Go binding around an Ethereum contract.
type MusdContract struct {
	MusdContractCaller     // Read-only binding to the contract
	MusdContractTransactor // Write-only binding to the contract
}

// MusdContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MusdContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MusdContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MusdContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MusdContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MusdContractSession struct {
	Contract     *MusdContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MusdContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MusdContractCallerSession struct {
	Contract *MusdContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MusdContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MusdContractTransactorSession struct {
	Contract     *MusdContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MusdContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MusdContractRaw struct {
	Contract *MusdContract // Generic contract binding to access the raw methods on
}

// MusdContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MusdContractCallerRaw struct {
	Contract *MusdContractCaller // Generic read-only contract binding to access the raw methods on
}

// MusdContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MusdContractTransactorRaw struct {
	Contract *MusdContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMusdContract creates a new instance of MusdContract, bound to a specific deployed contract.
func NewMusdContract(address common.Address, backend bind.ContractBackend) (*MusdContract, error) {
	contract, err := bindMusdContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MusdContract{MusdContractCaller: MusdContractCaller{contract: contract}, MusdContractTransactor: MusdContractTransactor{contract: contract}}, nil
}

// NewMusdContractCaller creates a new read-only instance of MusdContract, bound to a specific deployed contract.
func NewMusdContractCaller(address common.Address, caller bind.ContractCaller) (*MusdContractCaller, error) {
	contract, err := bindMusdContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MusdContractCaller{contract: contract}, nil
}

// NewMusdContractTransactor creates a new write-only instance of MusdContract, bound to a specific deployed contract.
func NewMusdContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MusdContractTransactor, error) {
	contract, err := bindMusdContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MusdContractTransactor{contract: contract}, nil
}

// bindMusdContract binds a generic wrapper to an already deployed contract.
func bindMusdContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MusdContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MusdContract *MusdContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MusdContract.Contract.MusdContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MusdContract *MusdContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MusdContract.Contract.MusdContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MusdContract *MusdContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MusdContract.Contract.MusdContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MusdContract *MusdContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MusdContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MusdContract *MusdContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MusdContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MusdContract *MusdContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MusdContract.Contract.contract.Transact(opts, method, params...)
}

// AddressesOf is a free data retrieval call binding the contract method 0x38a078c4.
//
// Solidity: function addressesOf(ownerAddr address) constant returns(minerAddr address, receiveAddr address)
func (_MusdContract *MusdContractCaller) AddressesOf(opts *bind.CallOpts, ownerAddr common.Address) (struct {
	MinerAddr   common.Address
	ReceiveAddr common.Address
}, error) {
	ret := new(struct {
		MinerAddr   common.Address
		ReceiveAddr common.Address
	})
	out := ret
	err := _MusdContract.contract.Call(opts, out, "addressesOf", ownerAddr)
	return *ret, err
}

// AddressesOf is a free data retrieval call binding the contract method 0x38a078c4.
//
// Solidity: function addressesOf(ownerAddr address) constant returns(minerAddr address, receiveAddr address)
func (_MusdContract *MusdContractSession) AddressesOf(ownerAddr common.Address) (struct {
	MinerAddr   common.Address
	ReceiveAddr common.Address
}, error) {
	return _MusdContract.Contract.AddressesOf(&_MusdContract.CallOpts, ownerAddr)
}

// AddressesOf is a free data retrieval call binding the contract method 0x38a078c4.
//
// Solidity: function addressesOf(ownerAddr address) constant returns(minerAddr address, receiveAddr address)
func (_MusdContract *MusdContractCallerSession) AddressesOf(ownerAddr common.Address) (struct {
	MinerAddr   common.Address
	ReceiveAddr common.Address
}, error) {
	return _MusdContract.Contract.AddressesOf(&_MusdContract.CallOpts, ownerAddr)
}

// AvailableForDelegation is a free data retrieval call binding the contract method 0xb8a90b03.
//
// Solidity: function availableForDelegation() constant returns(amount uint256)
func (_MusdContract *MusdContractCaller) AvailableForDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "availableForDelegation")
	return *ret0, err
}

// AvailableForDelegation is a free data retrieval call binding the contract method 0xb8a90b03.
//
// Solidity: function availableForDelegation() constant returns(amount uint256)
func (_MusdContract *MusdContractSession) AvailableForDelegation() (*big.Int, error) {
	return _MusdContract.Contract.AvailableForDelegation(&_MusdContract.CallOpts)
}

// AvailableForDelegation is a free data retrieval call binding the contract method 0xb8a90b03.
//
// Solidity: function availableForDelegation() constant returns(amount uint256)
func (_MusdContract *MusdContractCallerSession) AvailableForDelegation() (*big.Int, error) {
	return _MusdContract.Contract.AvailableForDelegation(&_MusdContract.CallOpts)
}

// AvailableForDelegationTo is a free data retrieval call binding the contract method 0x6bba5901.
//
// Solidity: function availableForDelegationTo(addr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCaller) AvailableForDelegationTo(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "availableForDelegationTo", addr)
	return *ret0, err
}

// AvailableForDelegationTo is a free data retrieval call binding the contract method 0x6bba5901.
//
// Solidity: function availableForDelegationTo(addr address) constant returns(amount uint256)
func (_MusdContract *MusdContractSession) AvailableForDelegationTo(addr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.AvailableForDelegationTo(&_MusdContract.CallOpts, addr)
}

// AvailableForDelegationTo is a free data retrieval call binding the contract method 0x6bba5901.
//
// Solidity: function availableForDelegationTo(addr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCallerSession) AvailableForDelegationTo(addr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.AvailableForDelegationTo(&_MusdContract.CallOpts, addr)
}

// AvailableTo is a free data retrieval call binding the contract method 0x2db9395a.
//
// Solidity: function availableTo(addr address) constant returns(balance uint256)
func (_MusdContract *MusdContractCaller) AvailableTo(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "availableTo", addr)
	return *ret0, err
}

// AvailableTo is a free data retrieval call binding the contract method 0x2db9395a.
//
// Solidity: function availableTo(addr address) constant returns(balance uint256)
func (_MusdContract *MusdContractSession) AvailableTo(addr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.AvailableTo(&_MusdContract.CallOpts, addr)
}

// AvailableTo is a free data retrieval call binding the contract method 0x2db9395a.
//
// Solidity: function availableTo(addr address) constant returns(balance uint256)
func (_MusdContract *MusdContractCallerSession) AvailableTo(addr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.AvailableTo(&_MusdContract.CallOpts, addr)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MusdContract *MusdContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MusdContract *MusdContractSession) Decimals() (uint8, error) {
	return _MusdContract.Contract.Decimals(&_MusdContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MusdContract *MusdContractCallerSession) Decimals() (uint8, error) {
	return _MusdContract.Contract.Decimals(&_MusdContract.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x2c9a9632.
//
// Solidity: function delegatedFrom(delegatorAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCaller) DelegatedFrom(opts *bind.CallOpts, delegatorAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "delegatedFrom", delegatorAddr)
	return *ret0, err
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x2c9a9632.
//
// Solidity: function delegatedFrom(delegatorAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractSession) DelegatedFrom(delegatorAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedFrom(&_MusdContract.CallOpts, delegatorAddr)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x2c9a9632.
//
// Solidity: function delegatedFrom(delegatorAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCallerSession) DelegatedFrom(delegatorAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedFrom(&_MusdContract.CallOpts, delegatorAddr)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(delegateAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCaller) DelegatedTo(opts *bind.CallOpts, delegateAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "delegatedTo", delegateAddr)
	return *ret0, err
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(delegateAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractSession) DelegatedTo(delegateAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedTo(&_MusdContract.CallOpts, delegateAddr)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(delegateAddr address) constant returns(amount uint256)
func (_MusdContract *MusdContractCallerSession) DelegatedTo(delegateAddr common.Address) (*big.Int, error) {
	return _MusdContract.Contract.DelegatedTo(&_MusdContract.CallOpts, delegateAddr)
}

// MaximumSupply is a free data retrieval call binding the contract method 0x0480e58b.
//
// Solidity: function maximumSupply() constant returns(uint256)
func (_MusdContract *MusdContractCaller) MaximumSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "maximumSupply")
	return *ret0, err
}

// MaximumSupply is a free data retrieval call binding the contract method 0x0480e58b.
//
// Solidity: function maximumSupply() constant returns(uint256)
func (_MusdContract *MusdContractSession) MaximumSupply() (*big.Int, error) {
	return _MusdContract.Contract.MaximumSupply(&_MusdContract.CallOpts)
}

// MaximumSupply is a free data retrieval call binding the contract method 0x0480e58b.
//
// Solidity: function maximumSupply() constant returns(uint256)
func (_MusdContract *MusdContractCallerSession) MaximumSupply() (*big.Int, error) {
	return _MusdContract.Contract.MaximumSupply(&_MusdContract.CallOpts)
}

// MinersOwners is a free data retrieval call binding the contract method 0x2e159065.
//
// Solidity: function minersOwners( address) constant returns(address)
func (_MusdContract *MusdContractCaller) MinersOwners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "minersOwners", arg0)
	return *ret0, err
}

// MinersOwners is a free data retrieval call binding the contract method 0x2e159065.
//
// Solidity: function minersOwners( address) constant returns(address)
func (_MusdContract *MusdContractSession) MinersOwners(arg0 common.Address) (common.Address, error) {
	return _MusdContract.Contract.MinersOwners(&_MusdContract.CallOpts, arg0)
}

// MinersOwners is a free data retrieval call binding the contract method 0x2e159065.
//
// Solidity: function minersOwners( address) constant returns(address)
func (_MusdContract *MusdContractCallerSession) MinersOwners(arg0 common.Address) (common.Address, error) {
	return _MusdContract.Contract.MinersOwners(&_MusdContract.CallOpts, arg0)
}

// MinersReceivers is a free data retrieval call binding the contract method 0x0a6c3247.
//
// Solidity: function minersReceivers( address) constant returns(address)
func (_MusdContract *MusdContractCaller) MinersReceivers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "minersReceivers", arg0)
	return *ret0, err
}

// MinersReceivers is a free data retrieval call binding the contract method 0x0a6c3247.
//
// Solidity: function minersReceivers( address) constant returns(address)
func (_MusdContract *MusdContractSession) MinersReceivers(arg0 common.Address) (common.Address, error) {
	return _MusdContract.Contract.MinersReceivers(&_MusdContract.CallOpts, arg0)
}

// MinersReceivers is a free data retrieval call binding the contract method 0x0a6c3247.
//
// Solidity: function minersReceivers( address) constant returns(address)
func (_MusdContract *MusdContractCallerSession) MinersReceivers(arg0 common.Address) (common.Address, error) {
	return _MusdContract.Contract.MinersReceivers(&_MusdContract.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MusdContract *MusdContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MusdContract *MusdContractSession) Name() (string, error) {
	return _MusdContract.Contract.Name(&_MusdContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MusdContract *MusdContractCallerSession) Name() (string, error) {
	return _MusdContract.Contract.Name(&_MusdContract.CallOpts)
}

// NumberTokenHolders is a free data retrieval call binding the contract method 0x95cc989c.
//
// Solidity: function numberTokenHolders() constant returns(count uint256)
func (_MusdContract *MusdContractCaller) NumberTokenHolders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "numberTokenHolders")
	return *ret0, err
}

// NumberTokenHolders is a free data retrieval call binding the contract method 0x95cc989c.
//
// Solidity: function numberTokenHolders() constant returns(count uint256)
func (_MusdContract *MusdContractSession) NumberTokenHolders() (*big.Int, error) {
	return _MusdContract.Contract.NumberTokenHolders(&_MusdContract.CallOpts)
}

// NumberTokenHolders is a free data retrieval call binding the contract method 0x95cc989c.
//
// Solidity: function numberTokenHolders() constant returns(count uint256)
func (_MusdContract *MusdContractCallerSession) NumberTokenHolders() (*big.Int, error) {
	return _MusdContract.Contract.NumberTokenHolders(&_MusdContract.CallOpts)
}

// OwnedBy is a free data retrieval call binding the contract method 0xb8377644.
//
// Solidity: function ownedBy( address) constant returns(uint256)
func (_MusdContract *MusdContractCaller) OwnedBy(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "ownedBy", arg0)
	return *ret0, err
}

// OwnedBy is a free data retrieval call binding the contract method 0xb8377644.
//
// Solidity: function ownedBy( address) constant returns(uint256)
func (_MusdContract *MusdContractSession) OwnedBy(arg0 common.Address) (*big.Int, error) {
	return _MusdContract.Contract.OwnedBy(&_MusdContract.CallOpts, arg0)
}

// OwnedBy is a free data retrieval call binding the contract method 0xb8377644.
//
// Solidity: function ownedBy( address) constant returns(uint256)
func (_MusdContract *MusdContractCallerSession) OwnedBy(arg0 common.Address) (*big.Int, error) {
	return _MusdContract.Contract.OwnedBy(&_MusdContract.CallOpts, arg0)
}

// OwnersMiners is a free data retrieval call binding the contract method 0x4297d632.
//
// Solidity: function ownersMiners( address) constant returns(address)
func (_MusdContract *MusdContractCaller) OwnersMiners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "ownersMiners", arg0)
	return *ret0, err
}

// OwnersMiners is a free data retrieval call binding the contract method 0x4297d632.
//
// Solidity: function ownersMiners( address) constant returns(address)
func (_MusdContract *MusdContractSession) OwnersMiners(arg0 common.Address) (common.Address, error) {
	return _MusdContract.Contract.OwnersMiners(&_MusdContract.CallOpts, arg0)
}

// OwnersMiners is a free data retrieval call binding the contract method 0x4297d632.
//
// Solidity: function ownersMiners( address) constant returns(address)
func (_MusdContract *MusdContractCallerSession) OwnersMiners(arg0 common.Address) (common.Address, error) {
	return _MusdContract.Contract.OwnersMiners(&_MusdContract.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MusdContract *MusdContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MusdContract *MusdContractSession) Symbol() (string, error) {
	return _MusdContract.Contract.Symbol(&_MusdContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MusdContract *MusdContractCallerSession) Symbol() (string, error) {
	return _MusdContract.Contract.Symbol(&_MusdContract.CallOpts)
}

// TokenHolders is a free data retrieval call binding the contract method 0x923108d9.
//
// Solidity: function tokenHolders( uint256) constant returns(address)
func (_MusdContract *MusdContractCaller) TokenHolders(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "tokenHolders", arg0)
	return *ret0, err
}

// TokenHolders is a free data retrieval call binding the contract method 0x923108d9.
//
// Solidity: function tokenHolders( uint256) constant returns(address)
func (_MusdContract *MusdContractSession) TokenHolders(arg0 *big.Int) (common.Address, error) {
	return _MusdContract.Contract.TokenHolders(&_MusdContract.CallOpts, arg0)
}

// TokenHolders is a free data retrieval call binding the contract method 0x923108d9.
//
// Solidity: function tokenHolders( uint256) constant returns(address)
func (_MusdContract *MusdContractCallerSession) TokenHolders(arg0 *big.Int) (common.Address, error) {
	return _MusdContract.Contract.TokenHolders(&_MusdContract.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MusdContract *MusdContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MusdContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MusdContract *MusdContractSession) TotalSupply() (*big.Int, error) {
	return _MusdContract.Contract.TotalSupply(&_MusdContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MusdContract *MusdContractCallerSession) TotalSupply() (*big.Int, error) {
	return _MusdContract.Contract.TotalSupply(&_MusdContract.CallOpts)
}

// AcceptAccountAddresses is a paid mutator transaction binding the contract method 0xda57537c.
//
// Solidity: function acceptAccountAddresses(ownerAddr address, receiverAddr address) returns(success bool)
func (_MusdContract *MusdContractTransactor) AcceptAccountAddresses(opts *bind.TransactOpts, ownerAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "acceptAccountAddresses", ownerAddr, receiverAddr)
}

// AcceptAccountAddresses is a paid mutator transaction binding the contract method 0xda57537c.
//
// Solidity: function acceptAccountAddresses(ownerAddr address, receiverAddr address) returns(success bool)
func (_MusdContract *MusdContractSession) AcceptAccountAddresses(ownerAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.AcceptAccountAddresses(&_MusdContract.TransactOpts, ownerAddr, receiverAddr)
}

// AcceptAccountAddresses is a paid mutator transaction binding the contract method 0xda57537c.
//
// Solidity: function acceptAccountAddresses(ownerAddr address, receiverAddr address) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) AcceptAccountAddresses(ownerAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.AcceptAccountAddresses(&_MusdContract.TransactOpts, ownerAddr, receiverAddr)
}

// AcceptMiningAddress is a paid mutator transaction binding the contract method 0x5254941a.
//
// Solidity: function acceptMiningAddress(ownerAddr address) returns(success bool)
func (_MusdContract *MusdContractTransactor) AcceptMiningAddress(opts *bind.TransactOpts, ownerAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "acceptMiningAddress", ownerAddr)
}

// AcceptMiningAddress is a paid mutator transaction binding the contract method 0x5254941a.
//
// Solidity: function acceptMiningAddress(ownerAddr address) returns(success bool)
func (_MusdContract *MusdContractSession) AcceptMiningAddress(ownerAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.AcceptMiningAddress(&_MusdContract.TransactOpts, ownerAddr)
}

// AcceptMiningAddress is a paid mutator transaction binding the contract method 0x5254941a.
//
// Solidity: function acceptMiningAddress(ownerAddr address) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) AcceptMiningAddress(ownerAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.AcceptMiningAddress(&_MusdContract.TransactOpts, ownerAddr)
}

// AcceptReceiverAddress is a paid mutator transaction binding the contract method 0x665d4040.
//
// Solidity: function acceptReceiverAddress(receiverAddr address) returns(success bool)
func (_MusdContract *MusdContractTransactor) AcceptReceiverAddress(opts *bind.TransactOpts, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "acceptReceiverAddress", receiverAddr)
}

// AcceptReceiverAddress is a paid mutator transaction binding the contract method 0x665d4040.
//
// Solidity: function acceptReceiverAddress(receiverAddr address) returns(success bool)
func (_MusdContract *MusdContractSession) AcceptReceiverAddress(receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.AcceptReceiverAddress(&_MusdContract.TransactOpts, receiverAddr)
}

// AcceptReceiverAddress is a paid mutator transaction binding the contract method 0x665d4040.
//
// Solidity: function acceptReceiverAddress(receiverAddr address) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) AcceptReceiverAddress(receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.AcceptReceiverAddress(&_MusdContract.TransactOpts, receiverAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) Delegate(opts *bind.TransactOpts, delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "delegate", delegateAddr, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) Delegate(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Delegate(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) Delegate(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Delegate(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// InitializeAccount is a paid mutator transaction binding the contract method 0x4f422ca9.
//
// Solidity: function initializeAccount(ownerAddr address, miningAddr address, receiverAddr address) returns()
func (_MusdContract *MusdContractTransactor) InitializeAccount(opts *bind.TransactOpts, ownerAddr common.Address, miningAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "initializeAccount", ownerAddr, miningAddr, receiverAddr)
}

// InitializeAccount is a paid mutator transaction binding the contract method 0x4f422ca9.
//
// Solidity: function initializeAccount(ownerAddr address, miningAddr address, receiverAddr address) returns()
func (_MusdContract *MusdContractSession) InitializeAccount(ownerAddr common.Address, miningAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.InitializeAccount(&_MusdContract.TransactOpts, ownerAddr, miningAddr, receiverAddr)
}

// InitializeAccount is a paid mutator transaction binding the contract method 0x4f422ca9.
//
// Solidity: function initializeAccount(ownerAddr address, miningAddr address, receiverAddr address) returns()
func (_MusdContract *MusdContractTransactorSession) InitializeAccount(ownerAddr common.Address, miningAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.InitializeAccount(&_MusdContract.TransactOpts, ownerAddr, miningAddr, receiverAddr)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(addr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) Mint(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "mint", addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(addr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Mint(&_MusdContract.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(addr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Mint(&_MusdContract.TransactOpts, addr, amount)
}

// ProposeAccountAddresses is a paid mutator transaction binding the contract method 0x6a505a42.
//
// Solidity: function proposeAccountAddresses(miningAddr address, receiverAddr address) returns()
func (_MusdContract *MusdContractTransactor) ProposeAccountAddresses(opts *bind.TransactOpts, miningAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "proposeAccountAddresses", miningAddr, receiverAddr)
}

// ProposeAccountAddresses is a paid mutator transaction binding the contract method 0x6a505a42.
//
// Solidity: function proposeAccountAddresses(miningAddr address, receiverAddr address) returns()
func (_MusdContract *MusdContractSession) ProposeAccountAddresses(miningAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.ProposeAccountAddresses(&_MusdContract.TransactOpts, miningAddr, receiverAddr)
}

// ProposeAccountAddresses is a paid mutator transaction binding the contract method 0x6a505a42.
//
// Solidity: function proposeAccountAddresses(miningAddr address, receiverAddr address) returns()
func (_MusdContract *MusdContractTransactorSession) ProposeAccountAddresses(miningAddr common.Address, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.ProposeAccountAddresses(&_MusdContract.TransactOpts, miningAddr, receiverAddr)
}

// ProposeMiningAddress is a paid mutator transaction binding the contract method 0x6847d05f.
//
// Solidity: function proposeMiningAddress(miningAddr address) returns()
func (_MusdContract *MusdContractTransactor) ProposeMiningAddress(opts *bind.TransactOpts, miningAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "proposeMiningAddress", miningAddr)
}

// ProposeMiningAddress is a paid mutator transaction binding the contract method 0x6847d05f.
//
// Solidity: function proposeMiningAddress(miningAddr address) returns()
func (_MusdContract *MusdContractSession) ProposeMiningAddress(miningAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.ProposeMiningAddress(&_MusdContract.TransactOpts, miningAddr)
}

// ProposeMiningAddress is a paid mutator transaction binding the contract method 0x6847d05f.
//
// Solidity: function proposeMiningAddress(miningAddr address) returns()
func (_MusdContract *MusdContractTransactorSession) ProposeMiningAddress(miningAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.ProposeMiningAddress(&_MusdContract.TransactOpts, miningAddr)
}

// ProposeReceiverAddress is a paid mutator transaction binding the contract method 0x08aee770.
//
// Solidity: function proposeReceiverAddress(receiverAddr address) returns()
func (_MusdContract *MusdContractTransactor) ProposeReceiverAddress(opts *bind.TransactOpts, receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "proposeReceiverAddress", receiverAddr)
}

// ProposeReceiverAddress is a paid mutator transaction binding the contract method 0x08aee770.
//
// Solidity: function proposeReceiverAddress(receiverAddr address) returns()
func (_MusdContract *MusdContractSession) ProposeReceiverAddress(receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.ProposeReceiverAddress(&_MusdContract.TransactOpts, receiverAddr)
}

// ProposeReceiverAddress is a paid mutator transaction binding the contract method 0x08aee770.
//
// Solidity: function proposeReceiverAddress(receiverAddr address) returns()
func (_MusdContract *MusdContractTransactorSession) ProposeReceiverAddress(receiverAddr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.ProposeReceiverAddress(&_MusdContract.TransactOpts, receiverAddr)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) Revoke(opts *bind.TransactOpts, delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "revoke", delegateAddr, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) Revoke(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Revoke(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(delegateAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) Revoke(delegateAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Revoke(&_MusdContract.TransactOpts, delegateAddr, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(toAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactor) Transfer(opts *bind.TransactOpts, toAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "transfer", toAddr, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(toAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractSession) Transfer(toAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Transfer(&_MusdContract.TransactOpts, toAddr, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(toAddr address, amount uint256) returns(success bool)
func (_MusdContract *MusdContractTransactorSession) Transfer(toAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MusdContract.Contract.Transfer(&_MusdContract.TransactOpts, toAddr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_MusdContract *MusdContractTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _MusdContract.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_MusdContract *MusdContractSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.TransferOwnership(&_MusdContract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(addr address) returns()
func (_MusdContract *MusdContractTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _MusdContract.Contract.TransferOwnership(&_MusdContract.TransactOpts, addr)
}
