package genesis

import (
	"testing"

	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/kowala-tech/kcoin/client/core"
	"github.com/stretchr/testify/require"
)

func TestLoaderFromConfig(t *testing.T) {
	t.Run("We get a main net block", func(t *testing.T) {
		block, err := NetworkGenesisBlock("", "kusd", MainNetwork)
		require.NoError(t, err, "Unexpected error when creating genesis block")

		require.Equal(t, getNetwork(MainNetwork), block.Config.ChainID)
	})
}

func TestLoaderFromFile(t *testing.T) {
	deterministicBlock, err := NetworkGenesisBlock("", "kusd", MainNetwork)
	require.NoError(t, err, "Unexpected error when creating genesis block")
	t.Run("We get a block from a saved file", func(t *testing.T) {
		if *update {
			jsonConfig := jsonEncodeGenesisBlock(deterministicBlock, t)
			updateGenesisGolden(t, genesisSampleBlockFilename(), jsonConfig)
		}

		loadedBlock, err := NetworkGenesisBlock(genesisSampleBlockFilename(), "", "")
		require.NoError(t, err, "Unexpected error when creating genesis block")

		require.Equal(t, getHashFromGenesisBlock(deterministicBlock), getHashFromGenesisBlock(loadedBlock))
	})
}

func updateGenesisGolden(t *testing.T, filename string, jsonConfig bytes.Buffer) {
	t.Logf("updated golden file for %s", filename)
	if err := ioutil.WriteFile(filename, jsonConfig.Bytes(), 0644); err != nil {
		t.Fatalf("failed to update golden file: %s", err)
	}
}

func genesisSampleBlockFilename() string {
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
