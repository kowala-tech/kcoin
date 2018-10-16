package genesis

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/pkg/errors"
)

//NetworkGenesisBlock returns a block to use as genesis based on the kcoin
//and on the type of network (test or main). If a filepath is specified it ignores
//the other params and loads from a file.
func NetworkGenesisBlock(filePath, currency, network string) (*core.Genesis, error) {

	if len(filePath) == 0 {

		// Check the cache first (see comment in networks.go)
		if curCfg, exists := LiveCurrencies[currency]; exists {

			gen := &core.Genesis{}

			if json, exists := curCfg[network]; exists {

				if err := gen.UnmarshalJSON(json); err != nil {
					return nil, fmt.Errorf("Failed to unmarshal cached genesis for %s (%s): %s", currency, network, err)
				}

				log.Info("Loading genesis config from freeze cache", "currency", currency, "network", network)

				return gen, nil

			} else {
				return nil, fmt.Errorf("Requested network %s not found in cache for currency %s", network, currency)
			}
		}

		// Otherwise load from file
		return loadFromConfig(currency, network)

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

func loadFromConfig(currency, network string) (*core.Genesis, error) {

	kcoinOptions, ok := Networks[currency]
	if !ok {
		return nil, errors.New("invalid kcoin")
	}

	genesisOpts, ok := kcoinOptions[network]
	if !ok {
		return nil, errors.New("invalid network options")
	}

	return Generate(genesisOpts)
}
