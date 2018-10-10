package mapping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SourceMapper struct {
	instructions []Instruction
	files        []string
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

type Instruction struct {
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

	return &SourceMapper{
		files: parseFiles(sourceMap),
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
