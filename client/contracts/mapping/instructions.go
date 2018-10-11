package mapping

import "fmt"

func ParseByteCode(byteCode []byte) ([]Instruction, error) {
	fmt.Printf("%v\n", byteCode)
	var instructions []Instruction

	for _, byteC := range byteCode {
		instructions = append(instructions, Instruction{OpCode: []byte{byteC}})
	}

	return instructions, nil
}

func IsPush(b byte) bool {
	return b >= 0x60 && b <= 0x7f
}

type Instruction struct {
	OpCode []byte
}
