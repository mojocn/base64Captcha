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

// NewRedisStore returns a new redis store for captcha with the
// given collection threshold and expiration time (duration). The returned
// store can registered with SetCustomStore to replace the default one.
func NewRedisStore(addr string, password string, db int, prefix string, expiration time.Duration) Store {
	s := new(redisStore)
	s.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
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
		var cmd *redis.StringCmd
		_, err = s.client.Pipelined(func(pipeliner redis.Pipeliner) (err error) {
			cmd = pipeliner.Get(s.prefix + id)
			if err != nil {
				return err
			}
			_, err = pipeliner.Del(s.prefix + id).Result()
			return
		})
		value, err = cmd.Result()
	} else {
		value, err = s.client.Get(s.prefix + id).Result()
	}

	if err != nil {
		return ""
	}

	return
}
