package log

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetContext(t *testing.T) {
	SetContext("message 1")

	assert.Equal(t, 1, len(root.ctx))
	assert.Equal(t, "message 1", root.ctx[0])
}

func TestSetContextMulti(t *testing.T) {
	SetContext("message 1", "message 2")

	assert.Equal(t, 2, len(root.ctx))
	assert.Equal(t, "message 1", root.ctx[0])
	assert.Equal(t, "message 2", root.ctx[1])
}
