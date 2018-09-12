package effort

// computational effort for different operations
const (
	ExpByte            uint64 = 10    // Times ceil(log256(exponent)) for the EXP instruction.
	Sload              uint64 = 50    // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	CallValueTransfer  uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	CallNewAccount     uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	Tx                 uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxContractCreation uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZero         uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv       uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	SstoreSet          uint64 = 20000 // Once per SLOAD operation.
	LogData            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend        uint64 = 2300  // Free computational resources given at beginning of call.

	Sha3          uint64 = 30    // Once per SHA3 operation.
	Sha3Word      uint64 = 6     // Once per word of the SHA3 operation's data.
	SstoreReset   uint64 = 5000  // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClear   uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefund  uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.
	Jumpdest      uint64 = 1     // Refunded computational resources, once per SSTORE operation if the zeroness changes to zero.
	Call          uint64 = 40    // Once per CALL operation & message call transaction.
	CreateData    uint64 = 200   //
	Exp           uint64 = 10    // Once per EXP instruction
	Log           uint64 = 375   // Per LOG* operation.
	Copy          uint64 = 3     //
	TierStep      uint64 = 0     // Once per operation, for a selection of them.
	LogTopic      uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	Create        uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	SuicideRefund uint64 = 24000 // Refunded following a suicide operation.
	Memory        uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.
	TxDataNonZero uint64 = 68    // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.

	Ecrecover            uint64 = 3000   // Elliptic curve sender recovery computational effort
	Sha256Base           uint64 = 60     // Base price for a SHA256 operation
	Sha256PerWord        uint64 = 12     // Per-word price for a SHA256 operation
	Ripemd160Base        uint64 = 600    // Base price for a RIPEMD160 operation
	Ripemd160PerWord     uint64 = 120    // Per-word price for a RIPEMD160 operation
	IdentityBase         uint64 = 15     // Base price for a data copy operation
	IdentityPerWord      uint64 = 3      // Per-work price for a data copy operation
	ModExpQuadCoeffDiv   uint64 = 20     // Divisor for the quadratic particle of the big int modular exponentiation
	Bn256Add             uint64 = 500    // computational resources needed for an elliptic curve addition
	Bn256ScalarMul       uint64 = 40000  // computational resources needed for an elliptic curve scalar multiplication
	Bn256PairingBase     uint64 = 100000 // Base price for an elliptic curve pairing check
	Bn256PairingPerPoint uint64 = 80000  // Per-point price for an elliptic curve pairing check
)
