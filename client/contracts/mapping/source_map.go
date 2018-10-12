package mapping

import (
	"fmt"
	"strconv"
	"strings"
)

const Separator = ";"
const PartSeparator = ":"

type SourceMapInstruction struct {
	byteOffsetStart   int
	sourceRangeLength int
	fileIndex         int
	typeJump          string
}

func ParseSourceMap(sm string) ([]SourceMapInstruction, error) {
	items := strings.Split(sm, Separator)

	var sourceItem []SourceMapInstruction

	byteOffsetConverter := &IntSourceMapConverter{}
	sourceRangeLengthConverter := &IntSourceMapConverter{}
	fileIndexConverter := &IntSourceMapConverter{}
	typeJumpConverter := &StringSourceMapConverter{}

	for _, item := range items {
		p := strings.Split(item, PartSeparator)

		if len(item) == 0 {
			sourceItem = append(sourceItem, SourceMapInstruction{
				byteOffsetStart:   byteOffsetConverter.lastItem,
				sourceRangeLength: sourceRangeLengthConverter.lastItem,
				fileIndex:         fileIndexConverter.lastItem,
				typeJump:          typeJumpConverter.lastItem,
			})
			continue
		}

		byteOffset, err := byteOffsetConverter.Extract(p[0])
		if err != nil {
			return nil, fmt.Errorf("error extracting byte offset: %s", err)
		}

		sourceRangeLength, err := sourceRangeLengthConverter.Extract(p[1])
		if err != nil {
			return nil, fmt.Errorf("error extracting source range length: %s", err)
		}

		fileIndex, err := fileIndexConverter.Extract(p[2])
		if err != nil {
			return nil, fmt.Errorf("error extracting file index: %s", err)
		}

		typeJump, err := typeJumpConverter.Extract(p[3])
		if err != nil {
			return nil, fmt.Errorf("error extracting type jump: %s", err)
		}

		sourceItem = append(sourceItem, SourceMapInstruction{
			byteOffsetStart:   byteOffset,
			sourceRangeLength: sourceRangeLength,
			fileIndex:         fileIndex,
			typeJump:          typeJump,
		})
	}

	return sourceItem, nil
}

type IntSourceMapConverter struct {
	lastItem int
}

func (e *IntSourceMapConverter) Extract(item string) (int, error) {
	if item == "" {
		return e.lastItem, nil
	}

	extractedItem, err := strconv.Atoi(item)
	if err != nil {
		return 0, fmt.Errorf("error extracting item %s: %s", item, err)
	}

	e.lastItem = extractedItem

	return extractedItem, nil
}

type StringSourceMapConverter struct {
	lastItem string
}

func (s *StringSourceMapConverter) Extract(item string) (string, error) {
	if item == "" {
		return s.lastItem, nil
	}

	s.lastItem = item

	return item, nil
}
