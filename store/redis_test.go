package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestNewRedisStore(t *testing.T) {

	s := NewRedisStoreByDetail("127.0.0.1:6379", "", 0, "test:", 5*time.Minute)
	assert.IsType(t, &redisStore{}, s)
}

func TestRedisStore_Get(t *testing.T) {

	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	id := fmt.Sprintf("%d", rand.Int63())
	value := fmt.Sprintf("%d", rand.Int63())
	s := NewRedisStoreByDetail("127.0.0.1:6379", "", 0, "test:", 5*time.Minute)
	c.Set("test:"+id, value, 5*time.Minute)
	assert.Equal(t, value, s.Get(id, false))
	assert.Equal(t, value, c.Get("test:" + id).Val())
}

func TestRedisStore_GetAndClear(t *testing.T) {

	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	id := fmt.Sprintf("%d", rand.Int63())
	value := fmt.Sprintf("%d", rand.Int63())
	s := NewRedisStoreByDetail("127.0.0.1:6379", "", 0, "test:", 5*time.Minute)
	c.Set("test:"+id, value, 5*time.Minute)
	assert.Equal(t, value, s.Get(id, true))
	assert.Equal(t, int64(0), c.Exists("test:" + id).Val())
}

func TestNewRedisStoreByClient(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	s := NewRedisStoreByClient(client, "test:", 5*time.Minute)
	assert.IsType(t, &redisStore{}, s)
}
