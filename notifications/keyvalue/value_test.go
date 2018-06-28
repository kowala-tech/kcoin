package keyvalue

import "testing"
import "github.com/stretchr/testify/require"

func wrappedMemoryKV() Value {
	kv := NewMemoryKeyValue()
	return WrapKeyValue(kv, "test")
}

func Test_MemoryValueStorage_StoreString(t *testing.T) {
	v := wrappedMemoryKV()
	str := "Hello world"

	err := v.PutString(str)
	require.NoError(t, err)

	received, err := v.GetString()
	require.NoError(t, err)
	require.Equal(t, received, str)
}

func Test_MemoryValueStorage_StoreInt(t *testing.T) {
	v := wrappedMemoryKV()
	var n int64 = 42

	err := v.PutInt64(n)
	require.NoError(t, err)

	received, err := v.GetInt64()
	require.NoError(t, err)
	require.Equal(t, received, n)
}
