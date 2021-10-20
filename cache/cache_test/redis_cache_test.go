package cache_test

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/stuartshome/carpedia/cache"
	"github.com/stuartshome/carpedia/mock/mock_redis_cache"
)

var Client *redis.Client

func TestSetNilDesc(t *testing.T) {
	// Given

	/*
		both of the below work.
		New returns a pointer to the memory address of the param
		and the allocated memory will be set to zero of the value type
		e.g. false for bool, 0 for int etc.
	*/
	redis := mock_redis_cache.RedisCache{}
	// redis := new(mock_redis_cache.RedisCache)

	key := "1"
	value := cache.Desc{}
	redis.On("Set", key, &value).Return("hello")
	redis.Set(key, &value)

	// When
	newcache := cache.NewRedisCache("test:8100", 0, 0)
	newcache.Set(key, &value)

	// Then
	redis.AssertCalled(t, "Set", key, &value)
}
