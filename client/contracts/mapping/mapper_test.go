package mapping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeCanCreateAMapAndGetItsFiles(t *testing.T) {
	mapper, err := NewFromSourceMap("files/combined.json")
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
