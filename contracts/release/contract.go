// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package release

import (
	"math/big"
	"strings"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
)

// ReleaseOracleABI is the input ABI used to generate the binding from.
const ReleaseOracleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"proposedVersion\",\"outputs\":[{\"name\":\"major\",\"type\":\"uint32\"},{\"name\":\"minor\",\"type\":\"uint32\"},{\"name\":\"patch\",\"type\":\"uint32\"},{\"name\":\"commit\",\"type\":\"bytes20\"},{\"name\":\"pass\",\"type\":\"address[]\"},{\"name\":\"fail\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"demote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"authVotes\",\"outputs\":[{\"name\":\"promote\",\"type\":\"address[]\"},{\"name\":\"demote\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"name\":\"major\",\"type\":\"uint32\"},{\"name\":\"minor\",\"type\":\"uint32\"},{\"name\":\"patch\",\"type\":\"uint32\"},{\"name\":\"commit\",\"type\":\"bytes20\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nuke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authProposals\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"promote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"major\",\"type\":\"uint32\"},{\"name\":\"minor\",\"type\":\"uint32\"},{\"name\":\"patch\",\"type\":\"uint32\"},{\"name\":\"commit\",\"type\":\"bytes20\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReleaseOracleBin is the compiled bytecode used for deploying new contracts.
const ReleaseOracleBin = `0x608060405234801561001057600080fd5b50604051611274380380611274833981016040528051018051600090151561009057336000818152602081905260408120805460ff191660019081179091558054808201825591527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018054600160a060020a0319169091179055610138565b5060005b815181101561013857600160008084848151811015156100b057fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff191691151591909117905581516001908390839081106100f357fe5b6020908102919091018101518254600180820185556000948552929093209092018054600160a060020a031916600160a060020a039093169290921790915501610094565b505061112b806101496000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166326db7648811461009d57806346f0975a146101a95780635c3d005d1461020e57806364ed31fe146102315780639d888e86146102eb578063bc8fbbf814610346578063bf8ecf9c1461035b578063d0e0813a14610370578063d67cbec914610391575b600080fd5b3480156100a957600080fd5b506100b26103cd565b604051808763ffffffff1663ffffffff1681526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff168152602001846bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610150578181015183820152602001610138565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561018f578181015183820152602001610177565b505050509050019850505050505050505060405180910390f35b3480156101b557600080fd5b506101be6104e6565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101fa5781810151838201526020016101e2565b505050509050019250505060405180910390f35b34801561021a57600080fd5b5061022f600160a060020a0360043516610549565b005b34801561023d57600080fd5b50610252600160a060020a0360043516610557565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561029657818101518382015260200161027e565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156102d55781810151838201526020016102bd565b5050505090500194505050505060405180910390f35b3480156102f757600080fd5b50610300610635565b6040805163ffffffff9687168152948616602086015292909416838301526bffffffffffffffffffffffff19166060830152608082019290925290519081900360a00190f35b34801561035257600080fd5b5061022f6106dd565b34801561036757600080fd5b506101be6106ed565b34801561037c57600080fd5b5061022f600160a060020a036004351661074d565b34801561039d57600080fd5b5061022f63ffffffff600435811690602435811690604435166bffffffffffffffffffffffff1960643516610758565b6004546006805460408051602080840282018101909252828152600094859485948594606094859463ffffffff808216956401000000008304821695680100000000000000008404909216946c0100000000000000000000000093849004909302939092600792849183018282801561046f57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610451575b50505050509150808054806020026020016040519081016040528092919081815260200182805480156104cb57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116104ad575b50505050509050955095509550955095509550909192939495565b6060600180548060200260200160405190810160405280929190818152602001828054801561053e57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610520575b505050505090505b90565b61055481600061076c565b50565b600160a060020a03811660009081526002602090815260409182902080548351818402810184019094528084526060938493600184019284918301828280156105c957602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116105ab575b505050505091508080548060200260200160405190810160405280929190818152602001828054801561062557602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610607575b5050505050905091509150915091565b6000806000806000806008805490506000141561066157600095508594508493508392508291506106d5565b60088054600019810190811061067357fe5b600091825260209091206004909102018054600182015463ffffffff80831699506401000000008304811698506801000000000000000083041696506c0100000000000000000000000091829004909102945067ffffffffffffffff16925090505b509091929394565b6106eb600080808080610bdb565b565b6060600380548060200260200160405190810160405280929190818152602001828054801561053e57602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610520575050505050905090565b61055481600161076c565b610766848484846001610bdb565b50505050565b33600090815260208190526040812054819060ff1615610766575050600160a060020a0382166000908152600260205260408120905b81548110156107e357815433908390839081106107bb57fe5b600091825260209091200154600160a060020a031614156107db57610766565b6001016107a2565b5060005b600182015481101561082e576001820180543391908390811061080657fe5b600091825260209091200154600160a060020a0316141561082657610766565b6001016107e7565b815415801561083f57506001820154155b1561089057600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018054600160a060020a031916600160a060020a0386161790555b82156108d3578154600181810184556000848152602090209091018054600160a060020a03191633179055546002908354919004106108ce57610766565b61090c565b600182810180548083018255600082815260209020018054600160a060020a031916331790559054905460029091041061090c57610766565b8280156109325750600160a060020a03841660009081526020819052604090205460ff16155b1561099e57600160a060020a0384166000818152602081905260408120805460ff191660019081179091558054808201825591527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018054600160a060020a0319169091179055610ae3565b821580156109c45750600160a060020a03841660009081526020819052604090205460ff165b15610ae35750600160a060020a0383166000908152602081905260408120805460ff191690555b600154811015610ae35783600160a060020a0316600182815481101515610a0e57fe5b600091825260209091200154600160a060020a03161415610adb57600180546000198101908110610a3b57fe5b60009182526020909120015460018054600160a060020a039092169183908110610a6157fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790556001805490610a9d90600019830161102a565b50600060048181556005805467ffffffffffffffff1916905590600681610ac48282611053565b610ad2600183016000611053565b50505050610ae3565b6001016109eb565b600160a060020a038416600090815260026020526040812090610b068282611053565b610b14600183016000611053565b5050600090505b6003548110156107665783600160a060020a0316600382815481101515610b3e57fe5b600091825260209091200154600160a060020a03161415610bd357600380546000198101908110610b6b57fe5b60009182526020909120015460038054600160a060020a039092169183908110610b9157fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790556003805490610bcd90600019830161102a565b50610766565b600101610b1b565b33600090815260208190526040812054819060ff16156110215782158015610c035750600654155b15610c0d57611021565b6006541515610c88576004805463ffffffff191663ffffffff8981169190911767ffffffff00000000191664010000000089831602176bffffffff000000000000000019166801000000000000000091881691909102176bffffffffffffffffffffffff166c01000000000000000000000000808704021790555b828015610d12575060045463ffffffff8881169116141580610cbd575060045463ffffffff8781166401000000009092041614155b80610cdf575060045463ffffffff868116680100000000000000009092041614155b80610d1257506004546c0100000000000000000000000090819004026bffffffffffffffffffffffff1990811690851614155b15610d1c57611021565b506006905060005b8154811015610d655781543390839083908110610d3d57fe5b600091825260209091200154600160a060020a03161415610d5d57611021565b600101610d24565b5060005b6001820154811015610db05760018201805433919083908110610d8857fe5b600091825260209091200154600160a060020a03161415610da857611021565b600101610d69565b8215610df3578154600181810184556000848152602090209091018054600160a060020a0319163317905554600290835491900410610dee57611021565b610e2c565b600182810180548083018255600082815260209020018054600160a060020a0319163317905590549054600290910410610e2c57611021565b8215610fe8576005805467ffffffffffffffff42811667ffffffffffffffff199283161783556008805460018101808355600092909252600480549181027ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee38101805463ffffffff191663ffffffff9485161780825583546401000000009081900486160267ffffffff000000001990911617808255835468010000000000000000908190049095169094026bffffffff0000000000000000199094169390931780845582546c01000000000000000000000000908190048102819004026bffffffffffffffffffffffff90911617835595547ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee48701805490961694169390931790935560068054919492939290917ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee590910190610f8d9082908490611071565b5060018281018054610fa29284019190611071565b5050600060048181556005805467ffffffffffffffff191690559450925060069150829050610fd18282611053565b610fdf600183016000611053565b50505050611021565b600060048181556005805467ffffffffffffffff191690559060068161100e8282611053565b61101c600183016000611053565b505050505b50505050505050565b81548183558181111561104e5760008381526020902061104e9181019083016110c1565b505050565b508054600082559060005260206000209081019061055491906110c1565b8280548282559060005260206000209081019282156110b15760005260206000209182015b828111156110b1578254825591600101919060010190611096565b506110bd9291506110db565b5090565b61054691905b808211156110bd57600081556001016110c7565b61054691905b808211156110bd578054600160a060020a03191681556001016110e15600a165627a7a72305820bebf0b2d2cd753cb556a46ad784ae43f81e3e7de599e3d8b414397e9a76b7c110029`

// DeployReleaseOracle deploys a new Ethereum contract, binding an instance of ReleaseOracle to it.
func DeployReleaseOracle(auth *bind.TransactOpts, backend bind.ContractBackend, signers []common.Address) (common.Address, *types.Transaction, *ReleaseOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(ReleaseOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReleaseOracleBin), backend, signers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReleaseOracle{ReleaseOracleCaller: ReleaseOracleCaller{contract: contract}, ReleaseOracleTransactor: ReleaseOracleTransactor{contract: contract}, ReleaseOracleFilterer: ReleaseOracleFilterer{contract: contract}}, nil
}

// ReleaseOracle is an auto generated Go binding around an Ethereum contract.
type ReleaseOracle struct {
	ReleaseOracleCaller     // Read-only binding to the contract
	ReleaseOracleTransactor // Write-only binding to the contract
	ReleaseOracleFilterer   // Log filterer for contract events
}

// ReleaseOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReleaseOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReleaseOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReleaseOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReleaseOracleSession struct {
	Contract     *ReleaseOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReleaseOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReleaseOracleCallerSession struct {
	Contract *ReleaseOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ReleaseOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReleaseOracleTransactorSession struct {
	Contract     *ReleaseOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ReleaseOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReleaseOracleRaw struct {
	Contract *ReleaseOracle // Generic contract binding to access the raw methods on
}

// ReleaseOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReleaseOracleCallerRaw struct {
	Contract *ReleaseOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ReleaseOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReleaseOracleTransactorRaw struct {
	Contract *ReleaseOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReleaseOracle creates a new instance of ReleaseOracle, bound to a specific deployed contract.
func NewReleaseOracle(address common.Address, backend bind.ContractBackend) (*ReleaseOracle, error) {
	contract, err := bindReleaseOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReleaseOracle{ReleaseOracleCaller: ReleaseOracleCaller{contract: contract}, ReleaseOracleTransactor: ReleaseOracleTransactor{contract: contract}, ReleaseOracleFilterer: ReleaseOracleFilterer{contract: contract}}, nil
}

// NewReleaseOracleCaller creates a new read-only instance of ReleaseOracle, bound to a specific deployed contract.
func NewReleaseOracleCaller(address common.Address, caller bind.ContractCaller) (*ReleaseOracleCaller, error) {
	contract, err := bindReleaseOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseOracleCaller{contract: contract}, nil
}

// NewReleaseOracleTransactor creates a new write-only instance of ReleaseOracle, bound to a specific deployed contract.
func NewReleaseOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ReleaseOracleTransactor, error) {
	contract, err := bindReleaseOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseOracleTransactor{contract: contract}, nil
}

// NewReleaseOracleFilterer creates a new log filterer instance of ReleaseOracle, bound to a specific deployed contract.
func NewReleaseOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ReleaseOracleFilterer, error) {
	contract, err := bindReleaseOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReleaseOracleFilterer{contract: contract}, nil
}

// bindReleaseOracle binds a generic wrapper to an already deployed contract.
func bindReleaseOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReleaseOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseOracle *ReleaseOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReleaseOracle.Contract.ReleaseOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseOracle *ReleaseOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.ReleaseOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseOracle *ReleaseOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.ReleaseOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseOracle *ReleaseOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReleaseOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseOracle *ReleaseOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseOracle *ReleaseOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.contract.Transact(opts, method, params...)
}

// AuthProposals is a free data retrieval call binding the contract method 0xbf8ecf9c.
//
// Solidity: function authProposals() constant returns(address[])
func (_ReleaseOracle *ReleaseOracleCaller) AuthProposals(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ReleaseOracle.contract.Call(opts, out, "authProposals")
	return *ret0, err
}

// AuthProposals is a free data retrieval call binding the contract method 0xbf8ecf9c.
//
// Solidity: function authProposals() constant returns(address[])
func (_ReleaseOracle *ReleaseOracleSession) AuthProposals() ([]common.Address, error) {
	return _ReleaseOracle.Contract.AuthProposals(&_ReleaseOracle.CallOpts)
}

// AuthProposals is a free data retrieval call binding the contract method 0xbf8ecf9c.
//
// Solidity: function authProposals() constant returns(address[])
func (_ReleaseOracle *ReleaseOracleCallerSession) AuthProposals() ([]common.Address, error) {
	return _ReleaseOracle.Contract.AuthProposals(&_ReleaseOracle.CallOpts)
}

// AuthVotes is a free data retrieval call binding the contract method 0x64ed31fe.
//
// Solidity: function authVotes(user address) constant returns(promote address[], demote address[])
func (_ReleaseOracle *ReleaseOracleCaller) AuthVotes(opts *bind.CallOpts, user common.Address) (struct {
	Promote []common.Address
	Demote  []common.Address
}, error) {
	ret := new(struct {
		Promote []common.Address
		Demote  []common.Address
	})
	out := ret
	err := _ReleaseOracle.contract.Call(opts, out, "authVotes", user)
	return *ret, err
}

// AuthVotes is a free data retrieval call binding the contract method 0x64ed31fe.
//
// Solidity: function authVotes(user address) constant returns(promote address[], demote address[])
func (_ReleaseOracle *ReleaseOracleSession) AuthVotes(user common.Address) (struct {
	Promote []common.Address
	Demote  []common.Address
}, error) {
	return _ReleaseOracle.Contract.AuthVotes(&_ReleaseOracle.CallOpts, user)
}

// AuthVotes is a free data retrieval call binding the contract method 0x64ed31fe.
//
// Solidity: function authVotes(user address) constant returns(promote address[], demote address[])
func (_ReleaseOracle *ReleaseOracleCallerSession) AuthVotes(user common.Address) (struct {
	Promote []common.Address
	Demote  []common.Address
}, error) {
	return _ReleaseOracle.Contract.AuthVotes(&_ReleaseOracle.CallOpts, user)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() constant returns(major uint32, minor uint32, patch uint32, commit bytes20, time uint256)
func (_ReleaseOracle *ReleaseOracleCaller) CurrentVersion(opts *bind.CallOpts) (struct {
	Major  uint32
	Minor  uint32
	Patch  uint32
	Commit [20]byte
	Time   *big.Int
}, error) {
	ret := new(struct {
		Major  uint32
		Minor  uint32
		Patch  uint32
		Commit [20]byte
		Time   *big.Int
	})
	out := ret
	err := _ReleaseOracle.contract.Call(opts, out, "currentVersion")
	return *ret, err
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() constant returns(major uint32, minor uint32, patch uint32, commit bytes20, time uint256)
func (_ReleaseOracle *ReleaseOracleSession) CurrentVersion() (struct {
	Major  uint32
	Minor  uint32
	Patch  uint32
	Commit [20]byte
	Time   *big.Int
}, error) {
	return _ReleaseOracle.Contract.CurrentVersion(&_ReleaseOracle.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() constant returns(major uint32, minor uint32, patch uint32, commit bytes20, time uint256)
func (_ReleaseOracle *ReleaseOracleCallerSession) CurrentVersion() (struct {
	Major  uint32
	Minor  uint32
	Patch  uint32
	Commit [20]byte
	Time   *big.Int
}, error) {
	return _ReleaseOracle.Contract.CurrentVersion(&_ReleaseOracle.CallOpts)
}

// ProposedVersion is a free data retrieval call binding the contract method 0x26db7648.
//
// Solidity: function proposedVersion() constant returns(major uint32, minor uint32, patch uint32, commit bytes20, pass address[], fail address[])
func (_ReleaseOracle *ReleaseOracleCaller) ProposedVersion(opts *bind.CallOpts) (struct {
	Major  uint32
	Minor  uint32
	Patch  uint32
	Commit [20]byte
	Pass   []common.Address
	Fail   []common.Address
}, error) {
	ret := new(struct {
		Major  uint32
		Minor  uint32
		Patch  uint32
		Commit [20]byte
		Pass   []common.Address
		Fail   []common.Address
	})
	out := ret
	err := _ReleaseOracle.contract.Call(opts, out, "proposedVersion")
	return *ret, err
}

// ProposedVersion is a free data retrieval call binding the contract method 0x26db7648.
//
// Solidity: function proposedVersion() constant returns(major uint32, minor uint32, patch uint32, commit bytes20, pass address[], fail address[])
func (_ReleaseOracle *ReleaseOracleSession) ProposedVersion() (struct {
	Major  uint32
	Minor  uint32
	Patch  uint32
	Commit [20]byte
	Pass   []common.Address
	Fail   []common.Address
}, error) {
	return _ReleaseOracle.Contract.ProposedVersion(&_ReleaseOracle.CallOpts)
}

// ProposedVersion is a free data retrieval call binding the contract method 0x26db7648.
//
// Solidity: function proposedVersion() constant returns(major uint32, minor uint32, patch uint32, commit bytes20, pass address[], fail address[])
func (_ReleaseOracle *ReleaseOracleCallerSession) ProposedVersion() (struct {
	Major  uint32
	Minor  uint32
	Patch  uint32
	Commit [20]byte
	Pass   []common.Address
	Fail   []common.Address
}, error) {
	return _ReleaseOracle.Contract.ProposedVersion(&_ReleaseOracle.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x46f0975a.
//
// Solidity: function signers() constant returns(address[])
func (_ReleaseOracle *ReleaseOracleCaller) Signers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ReleaseOracle.contract.Call(opts, out, "signers")
	return *ret0, err
}

// Signers is a free data retrieval call binding the contract method 0x46f0975a.
//
// Solidity: function signers() constant returns(address[])
func (_ReleaseOracle *ReleaseOracleSession) Signers() ([]common.Address, error) {
	return _ReleaseOracle.Contract.Signers(&_ReleaseOracle.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x46f0975a.
//
// Solidity: function signers() constant returns(address[])
func (_ReleaseOracle *ReleaseOracleCallerSession) Signers() ([]common.Address, error) {
	return _ReleaseOracle.Contract.Signers(&_ReleaseOracle.CallOpts)
}

// Demote is a paid mutator transaction binding the contract method 0x5c3d005d.
//
// Solidity: function demote(user address) returns()
func (_ReleaseOracle *ReleaseOracleTransactor) Demote(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ReleaseOracle.contract.Transact(opts, "demote", user)
}

// Demote is a paid mutator transaction binding the contract method 0x5c3d005d.
//
// Solidity: function demote(user address) returns()
func (_ReleaseOracle *ReleaseOracleSession) Demote(user common.Address) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Demote(&_ReleaseOracle.TransactOpts, user)
}

// Demote is a paid mutator transaction binding the contract method 0x5c3d005d.
//
// Solidity: function demote(user address) returns()
func (_ReleaseOracle *ReleaseOracleTransactorSession) Demote(user common.Address) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Demote(&_ReleaseOracle.TransactOpts, user)
}

// Nuke is a paid mutator transaction binding the contract method 0xbc8fbbf8.
//
// Solidity: function nuke() returns()
func (_ReleaseOracle *ReleaseOracleTransactor) Nuke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseOracle.contract.Transact(opts, "nuke")
}

// Nuke is a paid mutator transaction binding the contract method 0xbc8fbbf8.
//
// Solidity: function nuke() returns()
func (_ReleaseOracle *ReleaseOracleSession) Nuke() (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Nuke(&_ReleaseOracle.TransactOpts)
}

// Nuke is a paid mutator transaction binding the contract method 0xbc8fbbf8.
//
// Solidity: function nuke() returns()
func (_ReleaseOracle *ReleaseOracleTransactorSession) Nuke() (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Nuke(&_ReleaseOracle.TransactOpts)
}

// Promote is a paid mutator transaction binding the contract method 0xd0e0813a.
//
// Solidity: function promote(user address) returns()
func (_ReleaseOracle *ReleaseOracleTransactor) Promote(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ReleaseOracle.contract.Transact(opts, "promote", user)
}

// Promote is a paid mutator transaction binding the contract method 0xd0e0813a.
//
// Solidity: function promote(user address) returns()
func (_ReleaseOracle *ReleaseOracleSession) Promote(user common.Address) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Promote(&_ReleaseOracle.TransactOpts, user)
}

// Promote is a paid mutator transaction binding the contract method 0xd0e0813a.
//
// Solidity: function promote(user address) returns()
func (_ReleaseOracle *ReleaseOracleTransactorSession) Promote(user common.Address) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Promote(&_ReleaseOracle.TransactOpts, user)
}

// Release is a paid mutator transaction binding the contract method 0xd67cbec9.
//
// Solidity: function release(major uint32, minor uint32, patch uint32, commit bytes20) returns()
func (_ReleaseOracle *ReleaseOracleTransactor) Release(opts *bind.TransactOpts, major uint32, minor uint32, patch uint32, commit [20]byte) (*types.Transaction, error) {
	return _ReleaseOracle.contract.Transact(opts, "release", major, minor, patch, commit)
}

// Release is a paid mutator transaction binding the contract method 0xd67cbec9.
//
// Solidity: function release(major uint32, minor uint32, patch uint32, commit bytes20) returns()
func (_ReleaseOracle *ReleaseOracleSession) Release(major uint32, minor uint32, patch uint32, commit [20]byte) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Release(&_ReleaseOracle.TransactOpts, major, minor, patch, commit)
}

// Release is a paid mutator transaction binding the contract method 0xd67cbec9.
//
// Solidity: function release(major uint32, minor uint32, patch uint32, commit bytes20) returns()
func (_ReleaseOracle *ReleaseOracleTransactorSession) Release(major uint32, minor uint32, patch uint32, commit [20]byte) (*types.Transaction, error) {
	return _ReleaseOracle.Contract.Release(&_ReleaseOracle.TransactOpts, major, minor, patch, commit)
}
