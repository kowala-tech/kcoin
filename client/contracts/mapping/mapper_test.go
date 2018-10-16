package mapping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeCanCreateAMapAndGetItsFiles(t *testing.T) {
	mapper, err := NewFromCombinedRuntime("files/combined.json")
	assert.NoError(t, err)

	t.Run("Getting file out of bounds", func(t *testing.T) {
		_, err = mapper.GetFileByIndex(4)
		assert.EqualError(t, err, "invalid index for file")
	})

	t.Run("Getting file by index correctly", func(t *testing.T) {
		_, err := mapper.GetFileByIndex(0)
		assert.NoError(t, err)

		expectedFiles := []string{
			"../../truffle/contracts/sysvars/SystemVars.sol",
			"../../truffle/node_modules/openzeppelin-solidity/contracts/math/Math.sol",
			"../../truffle/node_modules/zos-lib/contracts/migrations/Initializable.sol",
		}

		for i, expectedFile := range expectedFiles {
			file, err := mapper.GetFileByIndex(i)
			assert.NoError(t, err)

			assert.Equal(t, expectedFile, file)
		}
	})
}

func TestWeCanCreateAMapAndGetInstructions(t *testing.T) {
	mapper, err := NewFromCombinedRuntime("files/combined.json")
	assert.NoError(t, err)

	t.Run("We Can Get Instructions By Index", func(t *testing.T) {
		_, err = mapper.GetFileByIndex(4)
		assert.EqualError(t, err, "invalid index for file")
	})
}

func TestWeCanGetTheContractByIndex(t *testing.T) {
	mapper, err := NewFromCombinedRuntime("files/combined.json")
	assert.NoError(t, err)

	contract, err := mapper.GetContractByIndex(1)
	assert.NoError(t, err)

	assert.NotNil(t, contract)
}

func TestItFailsWhenGettingContractWithOutOfBoundsIndex(t *testing.T) {
	mapper, err := NewFromCombinedRuntime("files/combined.json")
	assert.NoError(t, err)

	_, err = mapper.GetContractByIndex(50)
	assert.EqualError(t, err, "invalid index for contract")
}

func TestWeCanGetInstructionsByPcFromContract(t *testing.T) {
	contract := &Contract{
		instructions: []*Instruction{
			{OpCode: []byte{0x22, 0x23}},
			{OpCode: []byte{0x22, 0x23}},
			{OpCode: []byte{0x22, 0x23}},
		},
		sourceMapInstructions: []*SourceMapInstruction{
			{
				byteOffsetStart:   111,
				sourceRangeLength: 111,
				fileIndex:         1,
				typeJump:          "-",
			},
			{
				byteOffsetStart:   222,
				sourceRangeLength: 222,
				fileIndex:         2,
				typeJump:          "-",
			},
			{
				byteOffsetStart:   333,
				sourceRangeLength: 333,
				fileIndex:         2,
				typeJump:          "-",
			},
		},
	}

	t.Run("getting instruction by index", func(t *testing.T) {
		ins, smIns, err := contract.GetInstructionByPc(2)
		assert.NoError(t, err)

		assert.Equal(t, contract.sourceMapInstructions[2-1], smIns)
		assert.Equal(t, contract.instructions[2-1], ins)
	})

	t.Run("it fails when instruction is out of bounds", func(t *testing.T) {
		_, _, err := contract.GetInstructionByPc(10)
		assert.EqualError(t, err, "contract instruction out of bounds")
	})
}
