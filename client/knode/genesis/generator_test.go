package genesis

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateIsDeterministic(t *testing.T) {
	options := Networks["kusd"][MainNetwork]

	generatedGenesis, err := Generate(options)
	require.NoError(t, err)

	generatedGenesisTwo, err := Generate(options)
	require.NoError(t, err)

	assert.Equal(t, getHashFromGenesisBlock(generatedGenesis), getHashFromGenesisBlock(generatedGenesisTwo))
}

func TestGenerateIsDeterministicHasDifferentHash(t *testing.T) {
	options := Networks["kusd"][MainNetwork]
	generatedGenesis, err := Generate(options)
	require.NoError(t, err)

	options.ExtraData = "Something different in this config"
	generatedGenesisTwo, err := Generate(options)
	require.NoError(t, err)

	assert.NotEqual(t, getHashFromGenesisBlock(generatedGenesis), getHashFromGenesisBlock(generatedGenesisTwo))
}
