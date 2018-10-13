package mapping

import (
	"fmt"
	"strconv"
	"strings"
)

const Separator = ";"
const PartSeparator = ":"

const ByteOffsetConverterKey = "byteOffsetConverter"
const SourceRangeLengthConverterKey = "sourceRangeLengthConverter"
const FileIndexConverterKey = "fileIndexConverter"

type SourceMapInstruction struct {
	byteOffsetStart   int
	sourceRangeLength int
	fileIndex         int
	typeJump          string
}

func ParseSourceMap(sm string) ([]SourceMapInstruction, error) {
	items := strings.Split(sm, Separator)

	var sourceItem []SourceMapInstruction

	typeJumpConverter := &StringSourceMapConverter{}
	intConverters := getIntConverters()

	for _, item := range items {
		p := strings.Split(item, PartSeparator)

		if len(item) == 0 {
			sourceItem = append(sourceItem, SourceMapInstruction{
				byteOffsetStart:   intConverters[0].lastItem,
				sourceRangeLength: intConverters[1].lastItem,
				fileIndex:         intConverters[2].lastItem,
				typeJump:          typeJumpConverter.lastItem,
			})
			continue
		}

		includesTypeJump := len(p) == 4

		var err error
		var counterLegth int
		var typeJump string

		if includesTypeJump {
			counterLegth = len(p) - 1
			typeJump, err = typeJumpConverter.Extract(p[3])
			if err != nil {
				return nil, fmt.Errorf("error extracting type jump: %s", err)
			}
		} else {
			counterLegth = len(p)
			typeJump = typeJumpConverter.lastItem
		}

		for i := 0; i < counterLegth; i++ {
			intConverters[i].convertedValue, err = intConverters[i].Extract(p[i])
			if err != nil {
				return nil, fmt.Errorf("error extracting byte offset: %s", err)
			}
		}

		sourceItem = append(sourceItem, SourceMapInstruction{
			byteOffsetStart:   intConverters[0].convertedValue,
			sourceRangeLength: intConverters[1].convertedValue,
			fileIndex:         intConverters[2].convertedValue,
			typeJump:          typeJump,
		})
	}

	return sourceItem, nil
}

func getIntConverters() []*struct {
	name string
	*IntSourceMapConverter
	convertedValue int
} {
	return []*struct {
		name string
		*IntSourceMapConverter
		convertedValue int
	}{
		{
			ByteOffsetConverterKey,
			&IntSourceMapConverter{},
			0,
		},
		{
			SourceRangeLengthConverterKey,
			&IntSourceMapConverter{},
			0,
		},
		{
			FileIndexConverterKey,
			&IntSourceMapConverter{},
			0,
		},
	}
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
