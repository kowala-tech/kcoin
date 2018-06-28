package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MemorySetStorage(t *testing.T) {
	set := NewMemorySet()
	value := "Hello world"

	contains, err := set.Contains(value)
	require.NoError(t, err)
	require.False(t, contains)

	err = set.Add(value)
	require.NoError(t, err)

	contains, err = set.Contains(value)
	require.NoError(t, err)
	require.True(t, contains)

	err = set.Remove(value)
	require.NoError(t, err)

	contains, err = set.Contains(value)
	require.NoError(t, err)
	require.False(t, contains)
}
