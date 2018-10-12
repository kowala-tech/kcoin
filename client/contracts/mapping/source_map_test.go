package mapping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSourceMap(t *testing.T) {
	items := []struct {
		line  string
		items []SourceMapInstruction
	}{
		{
			"83:453:1:-;;;",
			[]SourceMapInstruction{
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
			},
		},
		{
			"1:2:1;:9;2:1:2;;",
			[]SourceMapInstruction{
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
				{
					byteOffsetStart:   83,
					sourceRangeLength: 453,
					fileIndex:         1,
					typeJump:          "-",
				},
			},
		},
	}

	for _, item := range items {
		mapItems, err := ParseSourceMap(item.line)
		assert.NoError(t, err)

		assert.Equal(t, item.items, mapItems)
	}
}
