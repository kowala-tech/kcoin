package params

import "math/big"

var (
	ComputeUnitPrice = new(big.Int).Mul(new(big.Int).SetUint64(400), new(big.Int).SetUint64(Shannon)) // Fixed compute unit price (400 Gwi/Shannon)
)

const (
	ComputeCapacity    uint64 = 4712388 // Compute capacity
	MinComputeCapacity uint64 = 5000    // Minimum the compute capacity may ever be.
	MaxTxSize float64 = 32 * 1024 // 32 KB 

	// computational efforts
	ExpByteCompEffort             uint64 = 10    // Times ceil(log256(exponent)) for the EXP instruction.
	SloadCompEffort               uint64 = 50    // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	CallValueTransferComputEffort uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	CallNewAccountCompEffort      uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	TxCompEffort                  uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxContractCreationCompEffort  uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroCompEffort          uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv                  uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	SstoreSetCompEffort           uint64 = 20000 // Once per SLOAD operation.
	LogDataComptEffort            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend                   uint64 = 2300  // Free computational resources given at beginning of call.

	Sha3CompEffort          uint64 = 30    // Once per SHA3 operation.
	Sha3WordCompEffort      uint64 = 6     // Once per word of the SHA3 operation's data.
	SstoreResetCompEffort   uint64 = 5000  // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClearCompEffort   uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundCompEffort  uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.
	JumpdestCompEffort      uint64 = 1     // Refunded gas, once per SSTORE operation if the zeroness changes to zero.
	CallCompEffort          uint64 = 40    // Once per CALL operation & message call transaction.
	CreateDataCompEffort    uint64 = 200   //
	CallCreateDepth         uint64 = 1024  // Maximum depth of call/create stack.
	ExpCompEffort           uint64 = 10    // Once per EXP instruction
	LogCompEffort           uint64 = 375   // Per LOG* operation.
	CopyCompEffort          uint64 = 3     //
	TierStepCompEffort      uint64 = 0     // Once per operation, for a selection of them.
	LogTopicCompEffort      uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	CreateCompEffort        uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	SuicideRefundCompEffort uint64 = 24000 // Refunded following a suicide operation.
	MemoryCompEffort        uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.
	TxDataNonZeroCompEffort uint64 = 68    // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.

	StackLimit uint64 = 1024 // Maximum size of VM stack allowed.

	// Precompiled contract gas prices

	EcrecoverCompEffort            uint64 = 3000   // Elliptic curve sender recovery gas price
	Sha256BaseCompEffort           uint64 = 60     // Base price for a SHA256 operation
	Sha256PerWordCompEffort        uint64 = 12     // Per-word price for a SHA256 operation
	Ripemd160BaseCompEffort        uint64 = 600    // Base price for a RIPEMD160 operation
	Ripemd160PerWordCompEffort     uint64 = 120    // Per-word price for a RIPEMD160 operation
	IdentityBaseCompEffort         uint64 = 15     // Base price for a data copy operation
	IdentityPerWordCompEffort      uint64 = 3      // Per-work price for a data copy operation
	ModExpQuadCoeffDiv             uint64 = 20     // Divisor for the quadratic particle of the big int modular exponentiation
	Bn256AddCompEffort             uint64 = 500    // Gas needed for an elliptic curve addition
	Bn256ScalarMulCompEffort       uint64 = 40000  // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseCompEffort     uint64 = 100000 // Base price for an elliptic curve pairing check
	Bn256PairingPerPointCompEffort uint64 = 80000  // Per-point price for an elliptic curve pairing check

	// Proof of Stake - timeouts
	ProposeDuration        uint64 = 500
	ProposeDeltaDuration   uint64 = 25
	PreVoteDuration        uint64 = 200
	PreVoteDeltaDuration   uint64 = 25
	PreCommitDuration      uint64 = 200
	PreCommitDeltaDuration uint64 = 25
	BlockTime              uint64 = 1000
)
