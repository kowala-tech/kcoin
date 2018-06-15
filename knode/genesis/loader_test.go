package genesis

import (
	"testing"

	"bufio"
	"bytes"
	"encoding/json"
	"github.com/kowala-tech/kcoin/core"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
)

func TestLoaderFromConfig(t *testing.T) {
	t.Run("We get a main net block", func(t *testing.T) {
		block, err := NetworkGenesisBlock("", "kusd", MainNetwork)
		require.NoError(t, err, "Unexpected error when creating genesis block")

		require.Equal(t, getNetwork(MainNetwork), block.Config.ChainID)
	})

	t.Run("We get a test net block", func(t *testing.T) {
		block, err := NetworkGenesisBlock("", "kusd", TestNetwork)
		require.NoError(t, err, "Unexpected error when creating genesis block")

		require.Equal(t, getNetwork(TestNetwork), block.Config.ChainID)
	})
}

func TestLoaderFromFile(t *testing.T) {
	deterministicBlock, _ := NetworkGenesisBlock("", "kusd", MainNetwork)
	t.Run("We get a main net block from a saved file", func(t *testing.T) {
		if *update {
			jsonConfig := jsonEncodeGenesisBlock(deterministicBlock, t)
			updateGenesisGolden(t, genesisBlockFilename(), jsonConfig)
		}

		loadedBlock, err := NetworkGenesisBlock(genesisBlockFilename(), "", "")
		require.NoError(t, err, "Unexpected error when creating genesis block")

		require.Equal(t, deterministicBlock.Config.ChainID, loadedBlock.Config.ChainID)
		require.Equal(t, deterministicBlock.Coinbase.Bytes(), loadedBlock.Coinbase.Bytes())
	})
}

func updateGenesisGolden(t *testing.T, filename string, jsonConfig bytes.Buffer) {
	t.Logf("updated golden file for %s", filename)
	if err := ioutil.WriteFile(filename, jsonConfig.Bytes(), 0644); err != nil {
		t.Fatalf("failed to update golden file: %s", err)
	}
}

func genesisBlockFilename() string {
	return filepath.Join("testfiles", "genesis.json.golden")
}

func jsonEncodeGenesisBlock(block *core.Genesis, t *testing.T) bytes.Buffer {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := json.NewEncoder(w).Encode(block)
	if err != nil {
		t.Fatalf("failed encoding genesis to json: %s", err)
	}
	w.Flush()
	return b
}
