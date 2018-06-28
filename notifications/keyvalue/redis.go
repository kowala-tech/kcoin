package keyvalue

import (
	"github.com/go-redis/redis"
)

type redisKV struct {
	client *redis.Client
}

// NewRedisKeyValue Returns a redis based KeyValue that uses plain GET and SET operations
func NewRedisKeyValue(client *redis.Client) KeyValue {
	return &redisKV{
		client: client,
	}
}

func (kv *redisKV) GetString(key string) (string, error) {
	res := kv.client.Get(key)
	if res.Err() == redis.Nil {
		return "", nil
	}
	return res.Result()
}

func (kv *redisKV) PutString(key string, value string) error {
	res := kv.client.Set(key, value, 0)
	return res.Err()
}

func (kv *redisKV) GetInt64(key string) (int64, error) {
	return parseInt64Key(kv, key)
}

func (kv *redisKV) PutInt64(key string, value int64) error {
	res := kv.client.Set(key, value, 0)
	return res.Err()
}

func (kv *redisKV) Delete(key string) error {
	res := kv.client.Del(key)
	return res.Err()
}
