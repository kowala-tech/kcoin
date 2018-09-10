package core

import (
	"errors"
	"math"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

var (
	errInsufficientBalanceForCompResouces = errors.New("insufficient balance to pay for the required computational resources")
)

/*
The State Transitioning Model

A state transition is a change made when a transaction is applied to the current world state
The state transitioning model does all all the necessary work to work out a valid new state root.

1) Nonce handling
2) Pre pay computational effort required for the transaction
3) Create a new state object if the recipient is \0*32
4) Value transfer
== If contract creation ==
  4a) Attempt to run transaction data
  4b) If valid, use result as code for the new state object
== end ==
5) Run Script section
6) Derive new state root
*/

type StateTransition struct {
	crpool               *CompResourcesPool
	msg                  Message
	compResources        uint64
	initialCompResources uint64
	value                *big.Int
	data                 []byte
	state                vm.StateDB
	vm                   *vm.VM
}

// Message represents a message sent to a contract.
type Message interface {
	From() common.Address
	To() *common.Address

	ComputationalEffort() uint64
	Value() *big.Int

	Nonce() uint64
	CheckNonce() bool
	Data() []byte
}

// IntrinsicCompEffort computes the intrinsic computational effort (in compute units)
// required for a message with the given data.
func IntrinsicCompEffort(data []byte, contractCreation bool) (uint64, error) {
	// Set the starting compute units required for the raw transaction
	var effort uint64
	if contractCreation {
		effort = params.TxContractCreationCompEffort
	} else {
		effort = params.TxCompEffort
	}
	// Bump the required compute units by the amount of transactional data
	if len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		if (math.MaxUint64-effort)/params.TxDataNonZeroCompEffort < nz {
			return 0, vm.ErrOutOfComputationalResources
		}
		effort += nz * params.TxDataNonZeroCompEffort

		z := uint64(len(data)) - nz
		if (math.MaxUint64-effort)/params.TxDataZeroCompEffort < z {
			return 0, vm.ErrOutOfComputationalResources
		}
		effort += z * params.TxDataZeroCompEffort
	}
	return effort, nil
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(vm *vm.VM, msg Message, crpool *CompResourcesPool) *StateTransition {
	return &StateTransition{
		crpool: crpool,
		vm:     vm,
		msg:    msg,
		value:  msg.Value(),
		data:   msg.Data(),
		state:  vm.StateDB,
	}
}

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any VM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(vm *vm.VM, msg Message, crpool *CompResourcesPool) ([]byte, uint64, bool, error) {
	return NewStateTransition(vm, msg, crpool).TransitionDb()
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	if st.msg == nil || st.msg.To() == nil /* contract creation */ {
		return common.Address{}
	}
	return *st.msg.To()
}

func (st *StateTransition) useCompResources(units uint64) error {
	if st.compResources < units {
		return vm.ErrOutOfComputationalResources
	}
	st.compResources -= units

	return nil
}

func (st *StateTransition) buyCompResources() error {
	mgval := new(big.Int).Mul(new(big.Int).SetUint64(st.msg.ComputationalEffort()), st.vm.ComputeUnitPrice)
	if st.state.GetBalance(st.msg.From()).Cmp(mgval) < 0 {
		return errInsufficientBalanceForCompResouces
	}
	st.crpool.AddResources(st.msg.ComputationalEffort())
	st.compResources += st.msg.ComputationalEffort()

	st.initialCompResources = st.msg.ComputationalEffort()
	st.state.SubBalance(st.msg.From(), mgval)
	return nil
}

func (st *StateTransition) preCheck() error {
	// Make sure this transaction's nonce is correct.
	if st.msg.CheckNonce() {
		nonce := st.state.GetNonce(st.msg.From())
		if nonce < st.msg.Nonce() {
			return ErrNonceTooHigh
		} else if nonce > st.msg.Nonce() {
			return ErrNonceTooLow
		}
	}
	return st.buyCompResources()
}

// TransitionDb will transition the state by applying the current message and
// returning the result including the the used compute units. It returns an error if it
// failed. An error indicates a consensus issue.
func (st *StateTransition) TransitionDb() (ret []byte, usedComputeUnits uint64, failed bool, err error) {
	if err = st.preCheck(); err != nil {
		return
	}
	msg := st.msg
	sender := vm.AccountRef(msg.From())
	contractCreation := msg.To() == nil

	// Pay intrinsic compute units
	effort, err := IntrinsicCompEffort(st.data, contractCreation)
	if err != nil {
		return nil, 0, false, err
	}
	if err = st.useCompResources(effort); err != nil {
		return nil, 0, false, err
	}

	var (
		env = st.vm
		// vm errors do not effect consensus and are therefor
		// not assigned to err, except for insufficient balance
		// error.
		vmerr error
	)
	if contractCreation {
		ret, _, st.compResources, vmerr = env.Create(sender, st.data, st.compResources, st.value)
	} else {
		// Increment the nonce for the next transaction
		st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)
		ret, st.compResources, vmerr = env.Call(sender, st.to(), st.data, st.compResources, st.value)
	}
	if vmerr != nil {
		log.Debug("VM returned with error", "err", vmerr)
		// The only possible consensus-error would be if there wasn't
		// sufficient balance to make the transfer happen. The first
		// balance transfer may never fail.
		if vmerr == vm.ErrInsufficientBalance {
			return nil, 0, false, vmerr
		}
	}
	st.refundCompResources()
	st.state.AddBalance(st.vm.Coinbase, new(big.Int).Mul(new(big.Int).SetUint64(st.compResourcesUsed()), st.vm.ComputeUnitPrice))

	return ret, st.compResourcesUsed(), vmerr != nil, err
}

func (st *StateTransition) refundCompResources() {
	// Apply refund counter, capped to half of the used gas.
	refund := st.compResourcesUsed() / 2
	if refund > st.state.GetRefund() {
		refund = st.state.GetRefund()
	}
	st.compResources += refund

	// Return kUSD for remaining computational resources.
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.compResources), st.vm.ComputeUnitPrice)
	st.state.AddBalance(st.msg.From(), remaining)

	// Also return remaining gas to the block gas counter so it is
	// available for the next transaction.
	st.crpool.AddResources(st.compResources)
}

// gasUsed returns the computation resources (in compute units) used up by the state transition.
func (st *StateTransition) compResourcesUsed() uint64 {
	return st.initialCompResources - st.compResources
}
