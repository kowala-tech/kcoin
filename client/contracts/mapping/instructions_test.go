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
	byteCode := "6001"

	instructions, err := ParseByteCode(common.Hex2Bytes(byteCode))
	assert.NoError(t, err)

	assert.Len(t, instructions, 1)
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
