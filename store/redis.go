package store

import (
	"github.com/go-redis/redis"
	"time"
)

// redisStore is use redis to store for captcha ids and their values.
type redisStore struct {
	// Redis Client Connection
	client *redis.Client

	// The Id keys prefix
	prefix string

	// Expiration time of captchas.
	expiration time.Duration
}

// NewRedisStore newRedisStoreByDetail returns a new redis store for captcha with the
// given collection threshold and expiration time (duration). The returned
// store can registered with SetCustomStore to replace the default one.
func NewRedisStore(client *redis.Client, prefix string, expiration time.Duration) Store {
	s := new(redisStore)
	s.client = client
	s.prefix = prefix
	s.expiration = expiration
	return s
}

func (s *redisStore) Set(id string, value string) {
	s.client.Set(s.prefix+id, value, s.expiration)
}

func (s *redisStore) Get(id string, clear bool) (value string) {

	var err error
	if clear {
		value, err = s.getAndClear(id)
	} else {
		value, err = s.client.Get(s.prefix + id).Result()
	}

	if err != nil {
		return ""
	}

	return
}

func (s *redisStore) getAndClear(id string) (string, error) {
	var (
		cmd *redis.StringCmd
		err error
	)

	_, err = s.client.Pipelined(func(pipeliner redis.Pipeliner) (err error) {
		cmd = pipeliner.Get(s.prefix + id)
		_, err = pipeliner.Del(s.prefix + id).Result()
		return
	})

	if err != nil {
		return "", err
	}

	return cmd.Result()
}
