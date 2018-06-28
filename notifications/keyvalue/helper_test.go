package keyvalue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func sharedKeyValueTests(t *testing.T, kv KeyValue) {
	t.Run("StoreString", func(t *testing.T) {
		str := "Hello world"

		err := kv.PutString("str", str)
		require.NoError(t, err)

		received, err := kv.GetString("str")
		require.NoError(t, err)
		require.Equal(t, str, received)
	})
	t.Run("StoreInt", func(t *testing.T) {
		var n int64 = 42

		err := kv.PutInt64("n", n)
		require.NoError(t, err)

		received, err := kv.GetInt64("n")
		require.NoError(t, err)
		require.Equal(t, n, received)
	})

	t.Run("GetMissing", func(t *testing.T) {
		receivedInt, err := kv.GetInt64("missing")
		require.NoError(t, err)
		require.Equal(t, int64(0), receivedInt)
		receivedStr, err := kv.GetString("missing")
		require.NoError(t, err)
		require.Equal(t, "", receivedStr)
	})
}
