package cache

//go:generate go run "github.com/vektra/mockery/cmd/mockery" -case=underscore -outpkg mock_redis_cache -output ../mock/mock_redis_cache -name=RedisCache
type RedisCache interface {
	Set(key string, value *Desc)
	Get(key string) *Desc
	GetAll() ([]*Desc, error)
}
