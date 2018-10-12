package mapping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kowala-tech/kcoin/client/common"
)

type SourceMapper struct {
	contracts             []Contract
	sourceMapInstructions []SourceMapInstruction
	files                 []string
}

type Contract struct {
	instructions          []Instruction
	sourceMapInstructions []SourceMapInstruction
}

type JSONSourceMap struct {
	Contracts  map[string]contract `json:"contracts"`
	Version    string              `json:"version"`
	SourceList []string            `json:"sourceList"`
}

type contract struct {
	BinRuntime    string `json:"bin-runtime"`
	SrcMapRuntime string `json:"srcmap-runtime"`
}

func NewFromSourceMap(filePath string) (*SourceMapper, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading source map: %s", err)
	}

	sourceMap := JSONSourceMap{}

	decoder := json.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&sourceMap)
	if err != nil {
		return nil, fmt.Errorf("error decoding source map: %s", err)
	}

	contracts, err := parseContracts(sourceMap)
	if err != nil {
		return nil, fmt.Errorf("error parsing contract data: %s", err)
	}

	return &SourceMapper{
		files:     parseFiles(sourceMap),
		contracts: contracts,
	}, nil
}

func (sm *SourceMapper) GetFileByIndex(index int) (string, error) {
	if len(sm.files) <= index {
		return "", fmt.Errorf("invalid index for file")
	}

	return sm.files[index], nil
}

func parseFiles(jsm JSONSourceMap) []string {
	var files []string

	for _, file := range jsm.SourceList {
		files = append(files, file)
	}

	return files
}

func parseContracts(jsm JSONSourceMap) ([]Contract, error) {
	var contracts []Contract

	for _, c := range jsm.Contracts {
		smi, err := ParseSourceMap(c.SrcMapRuntime)
		if err != nil {
			return nil, fmt.Errorf("error parsing source map: %s", err)
		}

		ins, err := ParseByteCode(common.Hex2Bytes(c.BinRuntime))
		if err != nil {
			return nil, fmt.Errorf("error parsing bytecode: %s", err)
		}

		contracts = append(contracts, Contract{
			sourceMapInstructions: smi,
			instructions:          ins,
		})
	}

	return contracts, nil
}
