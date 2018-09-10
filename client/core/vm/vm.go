package vm

import (
	"math/big"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
)

// emptyCodeHash is used by create to ensure deployment is disallowed to already
// deployed contract addresses (relevant after the account abstraction).
var emptyCodeHash = crypto.Keccak256Hash(nil)

type (
	// CanTransferFunc is the signature of a transfer guard function
	CanTransferFunc func(StateDB, common.Address, *big.Int) bool
	// TransferFunc is the signature of a transfer function
	TransferFunc func(StateDB, common.Address, common.Address, *big.Int)
	// GetHashFunc returns the nth block hash in the blockchain
	// and is used by the BLOCKHASH VM op code.
	GetHashFunc func(uint64) common.Hash
)

// run runs the given contract and takes care of running precompiles with a fallback to the byte code interpreter.
func run(vm *VM, contract *Contract, input []byte) ([]byte, error) {
	if contract.CodeAddr != nil {
		precompiles := PrecompiledContractsAndromeda
		if p := precompiles[*contract.CodeAddr]; p != nil {
			return RunPrecompiledContract(p, input, contract)
		}
	}
	return vm.interpreter.Run(contract, input)
}

// Context provides the VM with auxiliary information. Once provided
// it shouldn't be modified.
type Context struct {
	// CanTransfer returns whether the account contains
	// sufficient kcoin to transfer the value
	CanTransfer CanTransferFunc
	// Transfer transfers kcoin from one account to the other
	Transfer TransferFunc
	// GetHash returns the hash corresponding to n
	GetHash GetHashFunc

	// Message information
	Origin common.Address // Provides information for ORIGIN

	// Block information
	Coinbase    common.Address // Provides information for COINBASE
	BlockNumber *big.Int       // Provides information for NUMBER
	Time        *big.Int       // Provides information for TIME

	// Network information
	ComputeCapacity  uint64   // Provides information for COMPUTECAPACITY
	ComputeUnitPrice *big.Int // Provides information for COMPUTEUNITPRICE
}

// VM represents the Kowala Virtual Machine base object and provides
// the necessary tools to run a contract on the given state with
// the provided context. It should be noted that any error
// generated through any of the calls should be considered a
// revert-state-and-consume-all-computational-resources operation, no checks on
// specific errors should ever be performed. The interpreter makes
// sure that any errors generated are to be considered faulty code.
//
// The VM should never be reused and is not thread safe.
type VM struct {
	// Context provides auxiliary blockchain related information
	Context
	// StateDB gives access to the underlying state
	StateDB StateDB
	// Depth is the current call stack
	depth int

	// chainConfig contains information about the current chain
	chainConfig *params.ChainConfig
	// chain rules contains the chain rules for the current epoch
	chainRules params.Rules
	// virtual machine configuration options used to initialise the
	// vm.
	vmConfig Config
	// global (to this context) ethereum virtual machine
	// used throughout the execution of the tx.
	interpreter *Interpreter
	// abort is used to abort the VM calling operations
	// NOTE: must be set atomically
	abort int32
	// callCompResourcesTemp holds the computational resources available for the current call. This is needed because the
	// available computational resources are calculated in compResourcesCall* according to the 63/64 rule and later
	// applied in opCall*.
	callCompResourcesTemp uint64
}

// New returns a new VM. The returned VM is not thread safe and should
// only ever be used *once*.
func New(ctx Context, statedb StateDB, chainConfig *params.ChainConfig, vmConfig Config) *VM {
	vm := &VM{
		Context:     ctx,
		StateDB:     statedb,
		vmConfig:    vmConfig,
		chainConfig: chainConfig,
		chainRules:  chainConfig.Rules(ctx.BlockNumber),
	}

	vm.interpreter = NewInterpreter(vm, vmConfig)
	return vm
}

// Cancel cancels any running VM operation. This may be called concurrently and
// it's safe to be called multiple times.
func (vm *VM) Cancel() {
	atomic.StoreInt32(&vm.abort, 1)
}

// Call executes the contract associated with the addr with the given input as
// parameters. It also handles any necessary value transfer required and takes
// the necessary steps to create accounts and reverses the state in case of an
// execution error or failed value transfer.
func (vm *VM) Call(caller ContractRef, addr common.Address, input []byte, computeLimit uint64, value *big.Int) (ret []byte, leftOverCompResources uint64, err error) {
	if vm.vmConfig.NoRecursion && vm.depth > 0 {
		return nil, computeLimit, nil
	}

	// Fail if we're trying to execute above the call depth limit
	if vm.depth > int(params.CallCreateDepth) {
		return nil, computeLimit, ErrDepth
	}
	// Fail if we're trying to transfer more than the available balance
	if !vm.Context.CanTransfer(vm.StateDB, caller.Address(), value) {
		return nil, computeLimit, ErrInsufficientBalance
	}

	var (
		to       = AccountRef(addr)
		snapshot = vm.StateDB.Snapshot()
	)
	if !vm.StateDB.Exist(addr) {
		precompiles := PrecompiledContractsAndromeda
		if precompiles[addr] == nil && value.Sign() == 0 {
			// Calling a non existing account, don't do antything, but ping the tracer
			if vm.vmConfig.Debug && vm.depth == 0 {
				vm.vmConfig.Tracer.CaptureStart(caller.Address(), addr, false, input, computeLimit, value)
				vm.vmConfig.Tracer.CaptureEnd(ret, 0, 0, nil)
			}
			return nil, computeLimit, nil
		}
		vm.StateDB.CreateAccount(addr)
	}
	vm.Transfer(vm.StateDB, caller.Address(), to.Address(), value)

	// Initialise a new contract and set the code that is to be used by the VM.
	// The contract is a scoped environment for this execution context only.
	contract := NewContract(caller, to, value, computeLimit)
	contract.SetCallCode(&addr, vm.StateDB.GetCodeHash(addr), vm.StateDB.GetCode(addr))

	start := time.Now()

	// Capture the tracer start/end events in debug mode
	if vm.vmConfig.Debug && vm.depth == 0 {
		vm.vmConfig.Tracer.CaptureStart(caller.Address(), addr, false, input, computeLimit, value)

		defer func() { // Lazy evaluation of the parameters
			vm.vmConfig.Tracer.CaptureEnd(ret, computeLimit-contract.ComputationalResources, time.Since(start), err)
		}()
	}
	ret, err = run(vm, contract, input)

	// When an error was returned by the VM or when setting the creation code
	// above we revert to the snapshot and consume any computational resources remaining. Additionally
	// when we're in homestead this also counts for code storage compute limit errors.
	if err != nil {
		vm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseResources(contract.ComputationalResources)
		}
	}
	return ret, contract.ComputationalResources, err
}

// CallCode executes the contract associated with the addr with the given input
// as parameters. It also handles any necessary value transfer required and takes
// the necessary steps to create accounts and reverses the state in case of an
// execution error or failed value transfer.
//
// CallCode differs from Call in the sense that it executes the given address'
// code with the caller as context.
func (vm *VM) CallCode(caller ContractRef, addr common.Address, input []byte, computeLimit uint64, value *big.Int) (ret []byte, leftOverCompResources uint64, err error) {
	if vm.vmConfig.NoRecursion && vm.depth > 0 {
		return nil, computeLimit, nil
	}

	// Fail if we're trying to execute above the call depth limit
	if vm.depth > int(params.CallCreateDepth) {
		return nil, computeLimit, ErrDepth
	}
	// Fail if we're trying to transfer more than the available balance
	if !vm.CanTransfer(vm.StateDB, caller.Address(), value) {
		return nil, computeLimit, ErrInsufficientBalance
	}

	var (
		snapshot = vm.StateDB.Snapshot()
		to       = AccountRef(caller.Address())
	)
	// initialise a new contract and set the code that is to be used by the
	// VM. The contract is a scoped environment for this execution context
	// only.
	contract := NewContract(caller, to, value, computeLimit)
	contract.SetCallCode(&addr, vm.StateDB.GetCodeHash(addr), vm.StateDB.GetCode(addr))

	ret, err = run(vm, contract, input)
	if err != nil {
		vm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseResources(contract.ComputationalResources)
		}
	}
	return ret, contract.ComputationalResources, err
}

// DelegateCall executes the contract associated with the addr with the given input
// as parameters. It reverses the state in case of an execution error.
//
// DelegateCall differs from CallCode in the sense that it executes the given address'
// code with the caller as context and the caller is set to the caller of the caller.
func (vm *VM) DelegateCall(caller ContractRef, addr common.Address, input []byte, computeLimit uint64) (ret []byte, leftOverCompResources uint64, err error) {
	if vm.vmConfig.NoRecursion && vm.depth > 0 {
		return nil, computeLimit, nil
	}
	// Fail if we're trying to execute above the call depth limit
	if vm.depth > int(params.CallCreateDepth) {
		return nil, computeLimit, ErrDepth
	}

	var (
		snapshot = vm.StateDB.Snapshot()
		to       = AccountRef(caller.Address())
	)

	// Initialise a new contract and make initialise the delegate values
	contract := NewContract(caller, to, nil, computeLimit).AsDelegate()
	contract.SetCallCode(&addr, vm.StateDB.GetCodeHash(addr), vm.StateDB.GetCode(addr))

	ret, err = run(vm, contract, input)
	if err != nil {
		vm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseResources(contract.ComputationalResources)
		}
	}
	return ret, contract.ComputationalResources, err
}

// StaticCall executes the contract associated with the addr with the given input
// as parameters while disallowing any modifications to the state during the call.
// Opcodes that attempt to perform such modifications will result in exceptions
// instead of performing the modifications.
func (vm *VM) StaticCall(caller ContractRef, addr common.Address, input []byte, computeLimit uint64) (ret []byte, leftOverCompResources uint64, err error) {
	if vm.vmConfig.NoRecursion && vm.depth > 0 {
		return nil, computeLimit, nil
	}
	// Fail if we're trying to execute above the call depth limit
	if vm.depth > int(params.CallCreateDepth) {
		return nil, computeLimit, ErrDepth
	}
	// Make sure the readonly is only set if we aren't in readonly yet
	// this makes also sure that the readonly flag isn't removed for
	// child calls.
	if !vm.interpreter.readOnly {
		vm.interpreter.readOnly = true
		defer func() { vm.interpreter.readOnly = false }()
	}

	var (
		to       = AccountRef(addr)
		snapshot = vm.StateDB.Snapshot()
	)
	// Initialise a new contract and set the code that is to be used by the
	// VM. The contract is a scoped environment for this execution context
	// only.
	contract := NewContract(caller, to, new(big.Int), computeLimit)
	contract.SetCallCode(&addr, vm.StateDB.GetCodeHash(addr), vm.StateDB.GetCode(addr))

	// When an error was returned by the VM or when setting the creation code
	// above we revert to the snapshot and consume any computational resources remaining. Additionally
	// when we're in Homestead this also counts for code storage compute limit errors.
	ret, err = run(vm, contract, input)
	if err != nil {
		vm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseResources(contract.ComputationalResources)
		}
	}
	return ret, contract.ComputationalResources, err
}

// Create creates a new contract using code as deployment code.
func (vm *VM) Create(caller ContractRef, code []byte, computeLimit uint64, value *big.Int) (ret []byte, contractAddr common.Address, leftOverCompResources uint64, err error) {
	// Depth check execution. Fail if we're trying to execute above the
	// limit.
	if vm.depth > int(params.CallCreateDepth) {
		return nil, common.Address{}, computeLimit, ErrDepth
	}
	if !vm.CanTransfer(vm.StateDB, caller.Address(), value) {
		return nil, common.Address{}, computeLimit, ErrInsufficientBalance
	}
	// Ensure there's no existing contract already at the designated address
	nonce := vm.StateDB.GetNonce(caller.Address())
	vm.StateDB.SetNonce(caller.Address(), nonce+1)

	contractAddr = crypto.CreateAddress(caller.Address(), nonce)
	contractHash := vm.StateDB.GetCodeHash(contractAddr)
	if vm.StateDB.GetNonce(contractAddr) != 0 || (contractHash != (common.Hash{}) && contractHash != emptyCodeHash) {
		return nil, common.Address{}, 0, ErrContractAddressCollision
	}
	// Create a new account on the state
	snapshot := vm.StateDB.Snapshot()
	vm.StateDB.CreateAccount(contractAddr)
	vm.StateDB.SetNonce(contractAddr, 1)

	vm.Transfer(vm.StateDB, caller.Address(), contractAddr, value)

	// initialise a new contract and set the code that is to be used by the
	// VM. The contract is a scoped environment for this execution context
	// only.
	contract := NewContract(caller, AccountRef(contractAddr), value, computeLimit)
	contract.SetCallCode(&contractAddr, crypto.Keccak256Hash(code), code)

	if vm.vmConfig.NoRecursion && vm.depth > 0 {
		return nil, contractAddr, computeLimit, nil
	}

	if vm.vmConfig.Debug && vm.depth == 0 {
		vm.vmConfig.Tracer.CaptureStart(caller.Address(), contractAddr, true, code, computeLimit, value)
	}
	start := time.Now()

	ret, err = run(vm, contract, nil)

	// check whether the max code size has been exceeded
	maxCodeSizeExceeded := len(ret) > params.MaxCodeSize
	// if the contract creation ran successfully and no errors were returned
	// calculate the computational effort required to store the code. If the code could not
	// be stored due to not enough computational resources set an error and let it be handled
	// by the error checking condition below.
	if err == nil && !maxCodeSizeExceeded {
		createDataCompEffort := uint64(len(ret)) * params.CreateDataCompEffort
		if contract.UseResources(createDataCompEffort) {
			vm.StateDB.SetCode(contractAddr, ret)
		} else {
			err = ErrCodeStoreOutOfComputationalResources
		}
	}

	// When an error was returned by the VM or when setting the creation code
	// above we revert to the snapshot and consume any computational resources remaining. Additionally
	// when we're in homestead this also counts for code storage compute limit errors.
	if maxCodeSizeExceeded || (err != nil && err != ErrCodeStoreOutOfComputationalResources) {
		vm.StateDB.RevertToSnapshot(snapshot)
		if err != errExecutionReverted {
			contract.UseResources(contract.ComputationalResources)
		}
	}
	// Assign err if contract code size exceeds the max while the err is still empty.
	if maxCodeSizeExceeded && err == nil {
		err = errMaxCodeSizeExceeded
	}
	if vm.vmConfig.Debug && vm.depth == 0 {
		vm.vmConfig.Tracer.CaptureEnd(ret, computeLimit-contract.ComputationalResources, time.Since(start), err)
	}
	return ret, contractAddr, contract.ComputationalResources, err
}

// ChainConfig returns the environment's chain configuration
func (vm *VM) ChainConfig() *params.ChainConfig { return vm.chainConfig }

// Interpreter returns the VM interpreter
func (vm *VM) Interpreter() *Interpreter { return vm.interpreter }
