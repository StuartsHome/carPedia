package cache

import (
	"encoding/json"
	"fmt"
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
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	// this is storing json
	client.Set(key, json, cache.expires*time.Second) // this value associated to this key will expire in the number of seconds we specify

}
func (cache *redisCache) Get(key string) *Desc {
	client := cache.getClient()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

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

var cursor uint64
var n int

func (cache *redisCache) GetAll() ([]*Desc, error) {
	var data []*Desc
	client := cache.getClient()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	keys, _, err := client.Scan(cursor, "*", 10).Result()
	if err != nil {
		return nil, err
	}
	n += len(keys)
	for _, keyVal := range keys {
		tempData := cache.Get(keyVal)
		data = append(data, tempData)
	}

	return data, nil
}
