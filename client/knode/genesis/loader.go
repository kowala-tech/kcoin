package genesis

import (
	"encoding/json"
	"os"

	"github.com/kowala-tech/kcoin/core"
	"github.com/pkg/errors"
)

//NetworkGenesisBlock returns a block to use as genesis based on the kcoin
//and on the type of network (test or main). If a filepath is specified it ignores
//the other params and loads from a file.
func NetworkGenesisBlock(filePath, kcoin, networkType string) (*core.Genesis, error) {
	if len(filePath) == 0 {
		return loadFromConfig(kcoin, networkType)
	} else { // Extract from file.
		return loadFromFile(filePath)
	}
}

func loadFromFile(filePath string) (*core.Genesis, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read genesis file")
	}
	defer file.Close()

	genesis := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		return nil, errors.Wrap(err, "invalid genesis file")
	}

	return genesis, nil
}

func loadFromConfig(kcoin, networkType string) (*core.Genesis, error) {
	kcoinOptions, ok := Networks[kcoin]
	if !ok {
		return nil, errors.New("invalid kcoin")
	}

	genesisOpts, ok := kcoinOptions[networkType]
	if !ok {
		return nil, errors.New("invalid network options")
	}

	return Generate(genesisOpts)
}
