package mapping

import (
	"testing"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/stretchr/testify/assert"
)

func TestInstructionParserWith1ByteInstructions(t *testing.T) {
	byteCode := "50"

	instructions, err := ParseByteCode(common.Hex2Bytes(byteCode))
	assert.NoError(t, err)

	assert.Len(t, instructions, 1)
}

func TestInstructionParserWithMoreThan1ByteInstruction(t *testing.T) {
	insBytes := []struct {
		instruction     string
		numInstructions int
	}{
		{
			instruction:     "6001",
			numInstructions: 1,
		},
		{
			instruction:     "610102",
			numInstructions: 1,
		},
		{
			instruction:     "62010203",
			numInstructions: 1,
		},
	}

	for _, ins := range insBytes {
		instructions, err := ParseByteCode(common.Hex2Bytes(ins.instruction))
		assert.NoError(t, err)

		assert.Len(t, instructions, ins.numInstructions)
	}
}

func TestGetPushNumBytes(t *testing.T) {
	pushOps := []struct {
		pushOp   byte
		numBytes int
	}{
		{
			pushOp:   byte(0x60),
			numBytes: 1,
		},
		{
			pushOp:   byte(0x70),
			numBytes: 17,
		},
		{
			pushOp:   byte(0x7f),
			numBytes: 32,
		},
	}

	for _, pushOp := range pushOps {
		assert.Equal(t, pushOp.numBytes, GetLengthPushBytes(pushOp.pushOp))
	}
}

func TestIsPushFunc(t *testing.T) {
	ops := []struct {
		op     byte
		isPush bool
	}{
		{
			byte(0x60),
			true,
		},
		{
			byte(0x70),
			true,
		},
		{
			byte(0x7f),
			true,
		},
		{
			byte(0x80),
			false,
		},
		{
			byte(0x5f),
			false,
		},
	}

	for _, test := range ops {
		assert.Equal(t, IsPush(test.op), test.isPush, "error asserting that operation %d is push", test.op)
	}
}
