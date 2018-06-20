package genesis

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

// TestNetworkConfigs ensure that who changes the network gets a failing
// test as this configs are not supposed to change
// golden files exists to ensure that the config code matches the committed config
// if we need to change network config then we should generate new golden files with
// go test . --update
func TestNetworkConfigs(t *testing.T) {
	for kcoin, network := range Networks {
		t.Run(fmt.Sprintf("Config for kcoin %s hasn't changed", kcoin), func(t *testing.T) {
			for name, config := range network {
				t.Run(fmt.Sprintf("Config for network %s hasn't changed", name), func(t *testing.T) {
					jsonConfig := jsonEncodeGenesisConfig(config, t)
					filename := filepath.Join("testfiles", kcoin+"-"+name+".json.golden")
					if *update {
						updateGolden(t, filename, jsonConfig)
					}

					g, err := ioutil.ReadFile(filename)
					if err != nil {
						t.Fatalf("failed reading .golden: %s", err)
					}

					assert.Equal(t, jsonConfig.Bytes(), g)
				})
			}
		})
	}
}

func updateGolden(t *testing.T, filename string, jsonConfig bytes.Buffer) {
	t.Logf("updated golden file for %s", filename)

	if err := ioutil.WriteFile(filename, jsonConfig.Bytes(), 0644); err != nil {
		t.Fatalf("failed to update golden file: %s", err)
	}
}

func jsonEncodeGenesisConfig(config Options, t *testing.T) bytes.Buffer {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := json.NewEncoder(w).Encode(config)
	if err != nil {
		t.Fatalf("failed encoding config to json: %s", err)
	}
	w.Flush()
	return b
}

func getHashFromGenesisBlock(genesis *core.Genesis) common.Hash {
	b, _ := genesis.ToBlock()
	return b.Hash()
}
