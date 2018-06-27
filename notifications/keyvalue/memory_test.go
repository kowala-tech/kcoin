package keyvalue

import "testing"
import "github.com/stretchr/testify/require"

func Test_MemoryKeyValueStorage_StoreString(t *testing.T) {
	kv := NewMemoryKeyValue()
	str := "Hello world"

	err := kv.PutString("str", str)
	require.NoError(t, err)

	received, err := kv.GetString("str")
	require.NoError(t, err)
	require.Equal(t, received, str)
}

func Test_MemoryKeyValueStorage_StoreInt(t *testing.T) {
	kv := NewMemoryKeyValue()
	var n int64 = 42

	err := kv.PutInt64("n", n)
	require.NoError(t, err)

	received, err := kv.GetInt64("n")
	require.NoError(t, err)
	require.Equal(t, received, n)
}
