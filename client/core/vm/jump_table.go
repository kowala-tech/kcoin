package vm

import (
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin/client/params/effort"
)

type (
	executionFunc       func(pc *uint64, env *VM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error)
	effortFunc          func(effort.Table, *VM, *Contract, *Stack, *Memory, uint64) (uint64, error) // last parameter is the requested memory size as a uint64
	stackValidationFunc func(*Stack) error
	memorySizeFunc      func(*Stack) *big.Int
)

var errEffortUintOverflow = errors.New("effort uint64 overflow")

type operation struct {
	// execute is the operation function
	execute executionFunc
	// computationalEffort is the effort function and returns the required computational effort for execution
	computationalEffort effortFunc
	// validateStack validates the stack (size) for the operation
	validateStack stackValidationFunc
	// memorySize returns the memory size required for the operation
	memorySize memorySizeFunc

	halts   bool // indicates whether the operation should halt further execution
	jumps   bool // indicates whether the program counter should not increment
	writes  bool // determines whether this a state modifying operation
	valid   bool // indication whether the retrieved operation is valid and known
	reverts bool // determines whether the operation reverts state (implicitly halts)
	returns bool // determines whether the operations sets the return data content
}

var (
	andromedaInstructionSet = NewAndromedaInstructionSet()
)

// NewAndromedaInstructionSet returns the andromeda instructions
// that can be executed during the andromeda phase.
func NewAndromedaInstructionSet() [256]operation {
	return [256]operation{
		STOP: {
			execute:             opStop,
			computationalEffort: constEffortFunc(0),
			validateStack:       makeStackFunc(0, 0),
			halts:               true,
			valid:               true,
		},
		ADD: {
			execute:             opAdd,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		MUL: {
			execute:             opMul,
			computationalEffort: constEffortFunc(EffortFastStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SUB: {
			execute:             opSub,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		DIV: {
			execute:             opDiv,
			computationalEffort: constEffortFunc(EffortFastStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SDIV: {
			execute:             opSdiv,
			computationalEffort: constEffortFunc(EffortFastStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		MOD: {
			execute:             opMod,
			computationalEffort: constEffortFunc(EffortFastStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SMOD: {
			execute:             opSmod,
			computationalEffort: constEffortFunc(EffortFastStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		ADDMOD: {
			execute:             opAddmod,
			computationalEffort: constEffortFunc(EffortMidStep),
			validateStack:       makeStackFunc(3, 1),
			valid:               true,
		},
		MULMOD: {
			execute:             opMulmod,
			computationalEffort: constEffortFunc(EffortMidStep),
			validateStack:       makeStackFunc(3, 1),
			valid:               true,
		},
		EXP: {
			execute:             opExp,
			computationalEffort: effortExp,
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SIGNEXTEND: {
			execute:             opSignExtend,
			computationalEffort: constEffortFunc(EffortFastStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		LT: {
			execute:             opLt,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		GT: {
			execute:             opGt,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SLT: {
			execute:             opSlt,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SGT: {
			execute:             opSgt,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		EQ: {
			execute:             opEq,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		ISZERO: {
			execute:             opIszero,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(1, 1),
			valid:               true,
		},
		AND: {
			execute:             opAnd,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		XOR: {
			execute:             opXor,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		OR: {
			execute:             opOr,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		NOT: {
			execute:             opNot,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(1, 1),
			valid:               true,
		},
		BYTE: {
			execute:             opByte,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SHA3: {
			execute:             opSha3,
			computationalEffort: effortSha3,
			validateStack:       makeStackFunc(2, 1),
			memorySize:          memorySha3,
			valid:               true,
		},
		ADDRESS: {
			execute:             opAddress,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		BALANCE: {
			execute:             opBalance,
			computationalEffort: effortBalance,
			validateStack:       makeStackFunc(1, 1),
			valid:               true,
		},
		ORIGIN: {
			execute:             opOrigin,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		CALLER: {
			execute:             opCaller,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		CALLVALUE: {
			execute:             opCallValue,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		CALLDATALOAD: {
			execute:             opCallDataLoad,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(1, 1),
			valid:               true,
		},
		CALLDATASIZE: {
			execute:             opCallDataSize,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		CALLDATACOPY: {
			execute:             opCallDataCopy,
			computationalEffort: effortCallDataCopy,
			validateStack:       makeStackFunc(3, 0),
			memorySize:          memoryCallDataCopy,
			valid:               true,
		},
		CODESIZE: {
			execute:             opCodeSize,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		CODECOPY: {
			execute:             opCodeCopy,
			computationalEffort: effortCodeCopy,
			validateStack:       makeStackFunc(3, 0),
			memorySize:          memoryCodeCopy,
			valid:               true,
		},
		COMPUTEUNITPRICE: {
			execute:             opComputeUnitPrice,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		EXTCODESIZE: {
			execute:             opExtCodeSize,
			computationalEffort: effortExtCodeSize,
			validateStack:       makeStackFunc(1, 1),
			valid:               true,
		},
		EXTCODECOPY: {
			execute:             opExtCodeCopy,
			computationalEffort: effortExtCodeCopy,
			validateStack:       makeStackFunc(4, 0),
			memorySize:          memoryExtCodeCopy,
			valid:               true,
		},
		BLOCKHASH: {
			execute:             opBlockhash,
			computationalEffort: constEffortFunc(EffortExtStep),
			validateStack:       makeStackFunc(1, 1),
			valid:               true,
		},
		COINBASE: {
			execute:             opCoinbase,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		TIMESTAMP: {
			execute:             opTimestamp,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		NUMBER: {
			execute:             opNumber,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		COMPUTECAPACITY: {
			execute:             opComputeCapacity,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		POP: {
			execute:             opPop,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(1, 0),
			valid:               true,
		},
		MLOAD: {
			execute:             opMload,
			computationalEffort: effortMLoad,
			validateStack:       makeStackFunc(1, 1),
			memorySize:          memoryMLoad,
			valid:               true,
		},
		MSTORE: {
			execute:             opMstore,
			computationalEffort: effortMStore,
			validateStack:       makeStackFunc(2, 0),
			memorySize:          memoryMStore,
			valid:               true,
		},
		MSTORE8: {
			execute:             opMstore8,
			computationalEffort: effortMStore8,
			memorySize:          memoryMStore8,
			validateStack:       makeStackFunc(2, 0),

			valid: true,
		},
		SLOAD: {
			execute:             opSload,
			computationalEffort: effortSLoad,
			validateStack:       makeStackFunc(1, 1),
			valid:               true,
		},
		SSTORE: {
			execute:             opSstore,
			computationalEffort: effortSStore,
			validateStack:       makeStackFunc(2, 0),
			valid:               true,
			writes:              true,
		},
		JUMP: {
			execute:             opJump,
			computationalEffort: constEffortFunc(EffortMidStep),
			validateStack:       makeStackFunc(1, 0),
			jumps:               true,
			valid:               true,
		},
		JUMPI: {
			execute:             opJumpi,
			computationalEffort: constEffortFunc(EffortSlowStep),
			validateStack:       makeStackFunc(2, 0),
			jumps:               true,
			valid:               true,
		},
		PC: {
			execute:             opPc,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		MSIZE: {
			execute:             opMsize,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		RESOURCELEFT: {
			execute:             opResourceLeft,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		JUMPDEST: {
			execute:             opJumpdest,
			computationalEffort: constEffortFunc(effort.Jumpdest),
			validateStack:       makeStackFunc(0, 0),
			valid:               true,
		},
		PUSH1: {
			execute:             makePush(1, 1),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH2: {
			execute:             makePush(2, 2),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH3: {
			execute:             makePush(3, 3),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH4: {
			execute:             makePush(4, 4),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH5: {
			execute:             makePush(5, 5),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH6: {
			execute:             makePush(6, 6),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH7: {
			execute:             makePush(7, 7),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH8: {
			execute:             makePush(8, 8),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH9: {
			execute:             makePush(9, 9),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH10: {
			execute:             makePush(10, 10),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH11: {
			execute:             makePush(11, 11),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH12: {
			execute:             makePush(12, 12),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH13: {
			execute:             makePush(13, 13),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH14: {
			execute:             makePush(14, 14),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH15: {
			execute:             makePush(15, 15),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH16: {
			execute:             makePush(16, 16),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH17: {
			execute:             makePush(17, 17),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH18: {
			execute:             makePush(18, 18),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH19: {
			execute:             makePush(19, 19),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH20: {
			execute:             makePush(20, 20),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH21: {
			execute:             makePush(21, 21),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH22: {
			execute:             makePush(22, 22),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH23: {
			execute:             makePush(23, 23),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH24: {
			execute:             makePush(24, 24),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH25: {
			execute:             makePush(25, 25),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH26: {
			execute:             makePush(26, 26),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH27: {
			execute:             makePush(27, 27),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH28: {
			execute:             makePush(28, 28),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH29: {
			execute:             makePush(29, 29),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH30: {
			execute:             makePush(30, 30),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH31: {
			execute:             makePush(31, 31),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		PUSH32: {
			execute:             makePush(32, 32),
			computationalEffort: effortPush,
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		DUP1: {
			execute:             makeDup(1),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(1),
			valid:               true,
		},
		DUP2: {
			execute:             makeDup(2),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(2),
			valid:               true,
		},
		DUP3: {
			execute:             makeDup(3),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(3),
			valid:               true,
		},
		DUP4: {
			execute:             makeDup(4),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(4),
			valid:               true,
		},
		DUP5: {
			execute:             makeDup(5),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(5),
			valid:               true,
		},
		DUP6: {
			execute:             makeDup(6),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(6),
			valid:               true,
		},
		DUP7: {
			execute:             makeDup(7),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(7),
			valid:               true,
		},
		DUP8: {
			execute:             makeDup(8),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(8),
			valid:               true,
		},
		DUP9: {
			execute:             makeDup(9),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(9),
			valid:               true,
		},
		DUP10: {
			execute:             makeDup(10),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(10),
			valid:               true,
		},
		DUP11: {
			execute:             makeDup(11),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(11),
			valid:               true,
		},
		DUP12: {
			execute:             makeDup(12),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(12),
			valid:               true,
		},
		DUP13: {
			execute:             makeDup(13),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(13),
			valid:               true,
		},
		DUP14: {
			execute:             makeDup(14),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(14),
			valid:               true,
		},
		DUP15: {
			execute:             makeDup(15),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(15),
			valid:               true,
		},
		DUP16: {
			execute:             makeDup(16),
			computationalEffort: effortDup,
			validateStack:       makeDupStackFunc(16),
			valid:               true,
		},
		SWAP1: {
			execute:             makeSwap(1),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(2),
			valid:               true,
		},
		SWAP2: {
			execute:             makeSwap(2),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(3),
			valid:               true,
		},
		SWAP3: {
			execute:             makeSwap(3),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(4),
			valid:               true,
		},
		SWAP4: {
			execute:             makeSwap(4),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(5),
			valid:               true,
		},
		SWAP5: {
			execute:             makeSwap(5),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(6),
			valid:               true,
		},
		SWAP6: {
			execute:             makeSwap(6),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(7),
			valid:               true,
		},
		SWAP7: {
			execute:             makeSwap(7),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(8),
			valid:               true,
		},
		SWAP8: {
			execute:             makeSwap(8),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(9),
			valid:               true,
		},
		SWAP9: {
			execute:             makeSwap(9),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(10),
			valid:               true,
		},
		SWAP10: {
			execute:             makeSwap(10),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(11),
			valid:               true,
		},
		SWAP11: {
			execute:             makeSwap(11),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(12),
			valid:               true,
		},
		SWAP12: {
			execute:             makeSwap(12),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(13),
			valid:               true,
		},
		SWAP13: {
			execute:             makeSwap(13),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(14),
			valid:               true,
		},
		SWAP14: {
			execute:             makeSwap(14),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(15),
			valid:               true,
		},
		SWAP15: {
			execute:             makeSwap(15),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(16),
			valid:               true,
		},
		SWAP16: {
			execute:             makeSwap(16),
			computationalEffort: effortSwap,
			validateStack:       makeSwapStackFunc(17),
			valid:               true,
		},
		LOG0: {
			execute:             makeLog(0),
			computationalEffort: makeEffortLog(0),
			validateStack:       makeStackFunc(2, 0),
			memorySize:          memoryLog,
			valid:               true,
			writes:              true,
		},
		LOG1: {
			execute:             makeLog(1),
			computationalEffort: makeEffortLog(1),
			validateStack:       makeStackFunc(3, 0),
			memorySize:          memoryLog,
			valid:               true,
			writes:              true,
		},
		LOG2: {
			execute:             makeLog(2),
			computationalEffort: makeEffortLog(2),
			validateStack:       makeStackFunc(4, 0),
			memorySize:          memoryLog,
			valid:               true,
			writes:              true,
		},
		LOG3: {
			execute:             makeLog(3),
			computationalEffort: makeEffortLog(3),
			validateStack:       makeStackFunc(5, 0),
			memorySize:          memoryLog,
			valid:               true,
			writes:              true,
		},
		LOG4: {
			execute:             makeLog(4),
			computationalEffort: makeEffortLog(4),
			validateStack:       makeStackFunc(6, 0),
			memorySize:          memoryLog,
			valid:               true,
			writes:              true,
		},
		CREATE: {
			execute:             opCreate,
			computationalEffort: effortCreate,
			validateStack:       makeStackFunc(3, 1),
			memorySize:          memoryCreate,
			valid:               true,
			writes:              true,
			returns:             true,
		},
		CALL: {
			execute:             opCall,
			computationalEffort: effortCall,
			validateStack:       makeStackFunc(7, 1),
			memorySize:          memoryCall,
			valid:               true,
			returns:             true,
		},
		CALLCODE: {
			execute:             opCallCode,
			computationalEffort: effortCallCode,
			validateStack:       makeStackFunc(7, 1),
			memorySize:          memoryCall,
			valid:               true,
			returns:             true,
		},
		RETURN: {
			execute:             opReturn,
			computationalEffort: effortReturn,
			validateStack:       makeStackFunc(2, 0),
			memorySize:          memoryReturn,
			halts:               true,
			valid:               true,
		},
		SELFDESTRUCT: {
			execute:             opSuicide,
			computationalEffort: effortSuicide,
			validateStack:       makeStackFunc(1, 0),
			halts:               true,
			valid:               true,
			writes:              true,
		},
		DELEGATECALL: {
			execute:             opDelegateCall,
			computationalEffort: effortDelegateCall,
			validateStack:       makeStackFunc(6, 1),
			memorySize:          memoryDelegateCall,
			valid:               true,
			returns:             true,
		},
		STATICCALL: {
			execute:             opStaticCall,
			computationalEffort: effortStaticCall,
			validateStack:       makeStackFunc(6, 1),
			memorySize:          memoryStaticCall,
			valid:               true,
			returns:             true,
		},
		RETURNDATASIZE: {
			execute:             opReturnDataSize,
			computationalEffort: constEffortFunc(EffortQuickStep),
			validateStack:       makeStackFunc(0, 1),
			valid:               true,
		},
		RETURNDATACOPY: {
			execute:             opReturnDataCopy,
			computationalEffort: effortReturnDataCopy,
			validateStack:       makeStackFunc(3, 0),
			memorySize:          memoryReturnDataCopy,
			valid:               true,
		},
		REVERT: {
			execute:             opRevert,
			computationalEffort: effortRevert,
			validateStack:       makeStackFunc(2, 0),
			memorySize:          memoryRevert,
			valid:               true,
			reverts:             true,
			returns:             true,
		},
		SHL: {
			execute:             opSHL,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SHR: {
			execute:             opSHR,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
		SAR: {
			execute:             opSAR,
			computationalEffort: constEffortFunc(EffortFastestStep),
			validateStack:       makeStackFunc(2, 1),
			valid:               true,
		},
	}
}
