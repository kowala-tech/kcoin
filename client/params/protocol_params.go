package params

import (
	"math/big"
)

var (
	FixedComputeUnitPrice = new(big.Int).Mul(new(big.Int).SetUint64(400), new(big.Int).SetUint64(Shannon)) // Fixed compute unit price (400 Gwi/Shannon)

	TargetComputeCapacity = GenesisComputeCapacity // The artificial target
)

const (
	ComputeCapacityBoundDivisor uint64 = 1024    // The bound divisor of the compute capacity, used in update calculations.
	MinComputeCapacity          uint64 = 5000    // Minimum the compute capacity may ever be.
	GenesisComputeCapacity      uint64 = 4712388 // Compute capacity of the Genesis block.

	ExpByteComputeUnits            uint64 = 10    // Times ceil(log256(exponent)) for the EXP instruction.
	SloadComputeUnits              uint64 = 50    // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	CallValueTransferComputeUnits  uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	CallNewAccountComputeUnits     uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	TxComputeUnits                 uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxContractCreationComputeUnits uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroComputeUnits         uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv                   uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	SstoreSetComputeUnits          uint64 = 20000 // Once per SLOAD operation.
	LogDataComputeUnits            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend                    uint64 = 2300  // Free compute units given at beginning of call.

	Sha3ComputeUnits          uint64 = 30    // Once per SHA3 operation.
	Sha3WordComputeUnits      uint64 = 6     // Once per word of the SHA3 operation's data.
	SstoreResetComputeUnits   uint64 = 5000  // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClearComputeUnits   uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundComputeUnits  uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.
	JumpdestComputeUnits      uint64 = 1     // Refunded compute units, once per SSTORE operation if the zeroness changes to zero.
	CallComputeUnits          uint64 = 40    // Once per CALL operation & message call transaction.
	CreateDataComputeUnits    uint64 = 200   //
	CallCreateDepth           uint64 = 1024  // Maximum depth of call/create stack.
	ExpComputeUnits           uint64 = 10    // Once per EXP instruction
	LogComputeUnits           uint64 = 375   // Per LOG* operation.
	CopyComputeUnits          uint64 = 3     //
	TierStepComputeUnits      uint64 = 0     // Once per operation, for a selection of them.
	LogTopicComputeUnits      uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	CreateComputeUnits        uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	SuicideRefundComputeUnits uint64 = 24000 // Refunded following a suicide operation.
	MemoryComputeUnits        uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.
	TxDataNonZeroComputeUnits uint64 = 68    // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.

	// vm

	StackLimit uint64 = 1024 // Maximum size of VM stack allowed.

	// Precompiled contract

	EcrecoverComputeUnits            uint64 = 3000   // Elliptic curve sender recovery compute units
	Sha256BaseComputeUnits           uint64 = 60     // Base price for a SHA256 operation
	Sha256PerWordComputeUnits        uint64 = 12     // Per-word price for a SHA256 operation
	Ripemd160BaseComputeUnits        uint64 = 600    // Base price for a RIPEMD160 operation
	Ripemd160PerWordComputeUnits     uint64 = 120    // Per-word price for a RIPEMD160 operation
	IdentityBaseComputeUnits         uint64 = 15     // Base price for a data copy operation
	IdentityPerWordComputeUnits      uint64 = 3      // Per-work price for a data copy operation
	ModExpQuadCoeffDiv               uint64 = 20     // Divisor for the quadratic particle of the big int modular exponentiation
	Bn256AddComputeUnits             uint64 = 500    // Gas needed for an elliptic curve addition
	Bn256ScalarMulComputeUnits       uint64 = 40000  // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseComputeUnits     uint64 = 100000 // Base price for an elliptic curve pairing check
	Bn256PairingPerPointComputeUnits uint64 = 80000  // Per-point price for an elliptic curve pairing check

	// Proof of Stake - timeouts

	ProposeDuration        uint64 = 500
	ProposeDeltaDuration   uint64 = 25
	PreVoteDuration        uint64 = 200
	PreVoteDeltaDuration   uint64 = 25
	PreCommitDuration      uint64 = 200
	PreCommitDeltaDuration uint64 = 25
	BlockTime              uint64 = 1000
)
