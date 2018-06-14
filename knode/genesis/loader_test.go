package genesis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoader(t *testing.T) {
	t.Run("We get a main net block", func(t *testing.T) {
		block, err := NetworkGenesisBlock("", "kusd", MainNetwork)
		require.NoError(t, err, "Unexpected error when creating genesis block %s")

		require.Equal(t, getNetwork(MainNetwork), block.Config.ChainID)
	})

	t.Run("We get a test net block", func(t *testing.T) {
		block, err := NetworkGenesisBlock("", "kusd", TestNetwork)
		require.NoError(t, err, "Unexpected error when creating genesis block %s")

		require.Equal(t, getNetwork(TestNetwork), block.Config.ChainID)
	})
}
