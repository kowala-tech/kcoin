// +build integration

package set

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/kowala-tech/kcoin/notifications/environment"
	"github.com/stretchr/testify/require"
)

func redisClient(t *testing.T) *redis.Client {

	envReader := environment.NewReaderOs()
	redisAddr := envReader.Read("REDIS_ADDR")

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	require.NoError(t, client.Ping().Err())
	require.NoError(t, client.FlushDb().Err())
	return client
}

func Test_RedisSetStorage(t *testing.T) {
	set := NewRedisSet(redisClient(t), "testSet")
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
