package genesis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateIsDeterministic(t *testing.T) {
	options := Networks["kusd"][MainNetwork]

	generatedGenesis, err := Generate(options)
	assert.NoError(t, err)

	generatedGenesisTwo, err := Generate(options)
	assert.NoError(t, err)

	assert.Equal(t, getHashFromGenesisBlock(generatedGenesis), getHashFromGenesisBlock(generatedGenesisTwo))
}

func TestGenerateIsDeterministicDifferent(t *testing.T) {
	options := Networks["kusd"][MainNetwork]
	generatedGenesis, err := Generate(options)
	assert.NoError(t, err)

	options = Networks["kusd"][TestNetwork]
	generatedGenesisTwo, err := Generate(options)
	assert.NoError(t, err)

	assert.NotEqual(t, getHashFromGenesisBlock(generatedGenesis), getHashFromGenesisBlock(generatedGenesisTwo))
}
