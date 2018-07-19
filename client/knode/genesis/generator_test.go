package genesis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestGenerateIsDeterministic(t *testing.T) {
	options := Networks["kusd"][MainNetwork]

	generatedGenesis, err := Generate(options)
	assert.NoError(t, err)
	require.NotNil(t, generatedGenesis)

	generatedGenesisTwo, err := Generate(options)
	assert.NoError(t, err)
	require.NotNil(t, generatedGenesisTwo)

	assert.Equal(t, getHashFromGenesisBlock(generatedGenesis), getHashFromGenesisBlock(generatedGenesisTwo))
}

func TestGenerateIsDeterministicHasDifferentHash(t *testing.T) {
	options := Networks["kusd"][MainNetwork]
	generatedGenesis, err := Generate(options)
	assert.NoError(t, err)
	require.NotNil(t, generatedGenesis)

	options.ExtraData = "Something different in this config"
	generatedGenesisTwo, err := Generate(options)
	assert.NoError(t, err)
	require.NotNil(t, generatedGenesisTwo)

	assert.NotEqual(t, getHashFromGenesisBlock(generatedGenesis), getHashFromGenesisBlock(generatedGenesisTwo))
}
