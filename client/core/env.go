package core

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
)

// Message represents a message sent to an account.
type Message interface {
	From() common.Address
	To() *common.Address

	GasPrice() *big.Int
	Gas() uint64
	Value() *big.Int
	ComputeFee() *big.Int

	Nonce() uint64
	CheckNonce() bool
	Data() []byte
}

// createExecutionEnv returns a Signer based on the given chain config and block number.
func createExecuctionEnv(gasPool *GasPool) ExecutionEnv {
	var env ExecutionEnv
	switch {
	default:
		env = newDefaultEnv(gasPool)
	}
	return env
}

// ExecutionEnv represents an execution environment for messages
type ExecutionEnv interface {
	Exec(msg Message) (data []byte, gasUsage uint64, stabilityFee *big.Int, failed bool, err error)
}

// ApplyMessage applies a message in a given execution environment
func ApplyMessage(env ExecutionEnv, msg Message) ([]byte, uint64, *big.Int, bool, error) {
	return env.Exec(msg)
}

func newDefaultEnv(gasPool *GasPool) *defaultEnv {
	return &defaultEnv{
		gasPool: gasPool,
	}
}

type AccountManager interface {
	AddBalance(addr common.Address, amount *big.Int)
	SubBalance(addr common.Address, amount *big.Int)
	SetNonce(addr common.Address, nonce uint64)
	GetNonce(addr common.Address) (nonce uint64)
}

type defaultEnv struct {
	AccountManager
	gasPool          *GasPool
	stabilizationLvl uint64
	initialGas       uint64
	gasLeft          uint64
}

func (env *defaultEnv) Exec(msg Message) (data []byte, gasUsage uint64, stabilityFee *big.Int, failed bool, err error) {
	if err := validateMsg(msg, env.AccountManager); err != nil {
		return nil, 0, nil, false, err
	}

	// prepay stability fee
	if env.stabilizationLvl {
		computeFee := new(big.Int).Mul(new(big.Int).SetUint64(st.msg.Gas()), st.gasPrice)
		if stabilityFee := stability.CalcFee(computeFee, st.vm.StabilizationLevel(), st.msg.Value()); stabilityFee.Cmp(common.Big0) > 0 {
			if st.state.GetBalance(st.msg.From()).Cmp(stabilityFee) < 0 {
				return errInsufficientBalanceForStabilityFee
			}
			st.state.SubBalance(st.msg.From(), stabilityFee)
			st.stabilityFee = stabilityFee
		}
	}

	return nil, 0, nil, false, nil
}

func validateMsg(msg Message, accountMgr AccountManager) error {
	// Make sure this transaction's nonce is correct.
	if msg.CheckNonce() {
		nonce := accountMgr.GetNonce(msg.From())
		if nonce < msg.Nonce() {
			return ErrNonceTooHigh
		} else if nonce > msg.Nonce() {
			return ErrNonceTooLow
		}
	}

	return nil
}
