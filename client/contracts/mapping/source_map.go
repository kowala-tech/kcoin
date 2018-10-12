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
	var lastByteOffset, lastSourceRange, lastFileIndex int
	var lastTypeJump string
	var err error

	for _, item := range items {
		p := strings.Split(item, PartSeparator)

		if len(item) == 0 {
			sourceItem = append(sourceItem, SourceMapInstruction{
				byteOffsetStart:   lastByteOffset,
				sourceRangeLength: lastSourceRange,
				fileIndex:         lastFileIndex,
				typeJump:          lastTypeJump,
			})
			continue
		}

		var byteOffset int
		if p[0] == "" {
			byteOffset = lastByteOffset
		} else {
			byteOffset, err = strconv.Atoi(p[0])
			if err != nil {
				return nil, fmt.Errorf("error extracting byte offset: %s", err)
			}
			lastByteOffset = byteOffset
		}

		var sourceRangeLength int
		if p[1] == "" {
			sourceRangeLength = lastSourceRange
		} else {
			sourceRangeLength, err = strconv.Atoi(p[1])
			if err != nil {
				return nil, fmt.Errorf("error extracting source range length: %s", err)
			}

			lastSourceRange = sourceRangeLength
		}

		var fileIndex int
		if p[2] == "" {
			fileIndex = lastFileIndex
		} else {
			fileIndex, err = strconv.Atoi(p[2])
			if err != nil {
				return nil, fmt.Errorf("error extracting file index: %s", err)
			}
			lastFileIndex = fileIndex
		}

		var typeJump string
		if p[3] == "" {
			typeJump = lastTypeJump
		} else {
			typeJump = p[3]
			lastTypeJump = typeJump
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
