package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) RedisCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *Desc) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	// this is storing json
	client.Set(key, json, cache.expires*time.Second) // this value associated to this key will expire in the number of seconds we specify

}
func (cache *redisCache) Get(key string) *Desc {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	post := Desc{}
	if err := json.Unmarshal([]byte(val), &post); err != nil {
		panic(err)
	}
	return &post
}
