package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	c := DefaultConfig()
	assert.Equal(t, uint16(8080), c.listenPort)
}
