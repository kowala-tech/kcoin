package mapping

//ParseByteCode parses a byte code and returns an array of instructions.
func ParseByteCode(byteCode []byte) ([]*Instruction, error) {
	var instructions []*Instruction

	numBytes := len(byteCode)

	for i := 0; i < numBytes; {
		instruction := &Instruction{}

		lengthPushBytes := 0

		if IsPush(byteCode[i]) {
			lengthPushBytes = GetLengthPushBytes(byteCode[i])
			instruction.OpCode = byteCode[i : i+lengthPushBytes+1]
		} else {
			instruction.OpCode = []byte{byteCode[i]}
		}

		instructions = append(instructions, instruction)
		i = i + lengthPushBytes + 1
	}

	return instructions, nil
}

//IsPush returns true if it is a push operation opcode
func IsPush(b byte) bool {
	return b >= 0x60 && b <= 0x7f
}

//GetLengthPushBytes returns the number of bytes the operation adds to the stack
func GetLengthPushBytes(pushOp byte) int {
	return int(pushOp) - 0x5f
}

type Instruction struct {
	OpCode []byte
}
