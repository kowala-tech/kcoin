package core

import (
	"errors"
	"math"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/log"
	effrt "github.com/kowala-tech/kcoin/client/params/effort"
)

var (
	errInsufficientBalanceForComputationalResource = errors.New("insufficient balance to pay for computational resource")
)

/*
The State Transitioning Model

A state transition is a change made when a transaction is applied to the current world state
The state transitioning model does all all the necessary work to work out a valid new state root.

1) Nonce handling
2) Pre pay computational resource
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
	crpool          *ComputationalResourcePool
	msg             Message
	resource        uint64
	initialResource uint64
	value           *big.Int
	data            []byte
	state           vm.StateDB
	vm              *vm.VM
}

// Message represents a message sent to a contract.
type Message interface {
	From() common.Address
	To() *common.Address

	ComputeLimit() uint64
	Value() *big.Int

	Nonce() uint64
	CheckNonce() bool
	Data() []byte
}

// IntrinsicEffort computes the 'intrinsic computational effort' for a message with the given data.
func IntrinsicEffort(data []byte, contractCreation bool) (uint64, error) {
	// Set the starting effort for the raw transaction
	var effort uint64
	if contractCreation {
		effort = effrt.TxContractCreation
	} else {
		effort = effrt.Tx
	}
	// Bump the required effort by the amount of transactional data
	if len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		if (math.MaxUint64-effort)/effrt.TxDataNonZero < nz {
			return 0, vm.ErrOutOfComputationalResource
		}
		effort += nz * effrt.TxDataNonZero

		z := uint64(len(data)) - nz
		if (math.MaxUint64-effort)/effrt.TxDataZero < z {
			return 0, vm.ErrOutOfComputationalResource
		}
		effort += z * effrt.TxDataZero
	}
	return effort, nil
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(vm *vm.VM, msg Message, crpool *ComputationalResourcePool) *StateTransition {
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
// the computational resource usage (which includes refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(vm *vm.VM, msg Message, crpool *ComputationalResourcePool) ([]byte, uint64, bool, error) {
	return NewStateTransition(vm, msg, crpool).TransitionDb()
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	if st.msg == nil || st.msg.To() == nil /* contract creation */ {
		return common.Address{}
	}
	return *st.msg.To()
}

func (st *StateTransition) useResource(amount uint64) error {
	if st.resource < amount {
		return vm.ErrOutOfComputationalResource
	}
	st.resource -= amount

	return nil
}

func (st *StateTransition) buyResource() error {
	mgval := new(big.Int).Mul(new(big.Int).SetUint64(st.msg.ComputeLimit()), st.vm.ComputeUnitPrice)
	if st.state.GetBalance(st.msg.From()).Cmp(mgval) < 0 {
		return errInsufficientBalanceForComputationalResource
	}
	if err := st.crpool.SubResource(st.msg.ComputeLimit()); err != nil {
		return err
	}
	st.resource += st.msg.ComputeLimit()

	st.initialResource = st.msg.ComputeLimit()
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
	return st.buyResource()
}

// TransitionDb will transition the state by applying the current message and
// returning the result including the the used computational resource. It returns an error if it
// failed. An error indicates a consensus issue.
func (st *StateTransition) TransitionDb() (ret []byte, resourceUsage uint64, failed bool, err error) {
	if err = st.preCheck(); err != nil {
		return
	}
	msg := st.msg
	sender := vm.AccountRef(msg.From())
	contractCreation := msg.To() == nil

	// Pay intrinsic computational effort
	effort, err := IntrinsicEffort(st.data, contractCreation)
	if err != nil {
		return nil, 0, false, err
	}
	if err = st.useResource(effort); err != nil {
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
		ret, _, st.resource, vmerr = env.Create(sender, st.data, st.resource, st.value)
	} else {
		// Increment the nonce for the next transaction
		st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)
		ret, st.resource, vmerr = env.Call(sender, st.to(), st.data, st.resource, st.value)
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
	st.refundResource()
	st.state.AddBalance(st.vm.Coinbase, new(big.Int).Mul(new(big.Int).SetUint64(st.resourceUsed()), st.vm.ComputeUnitPrice))

	return ret, st.resourceUsed(), vmerr != nil, err
}

func (st *StateTransition) refundResource() {
	// Apply refund counter, capped to half of the used computational resource.
	refund := st.resourceUsed() / 2
	if refund > st.state.GetRefund() {
		refund = st.state.GetRefund()
	}
	st.resource += refund

	// Return kUSD for remaining computational resource.
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.resource), st.vm.ComputeUnitPrice)
	st.state.AddBalance(st.msg.From(), remaining)

	// Also return remaining computational resource to the block resource counter so it is
	// available for the next transaction.
	st.crpool.AddResource(st.resource)
}

// resourceUsed returns the amount of computational resource used up by the state transition.
func (st *StateTransition) resourceUsed() uint64 {
	return st.initialResource - st.resource
}
