package set

import "github.com/go-redis/redis"

type redisSet struct {
	client *redis.Client
	setKey string
}

func NewRedisSet(client *redis.Client, setKey string) Set {
	return &redisSet{
		client: client,
		setKey: setKey,
	}
}

func (set *redisSet) Add(value string) error {
	res := set.client.SAdd(set.setKey, value)
	return res.Err()
}

func (set *redisSet) Remove(value string) error {
	res := set.client.SRem(set.setKey, value)
	return res.Err()
}

func (set *redisSet) Contains(value string) (bool, error) {
	res := set.client.SIsMember(set.setKey, value)
	return res.Result()
}
