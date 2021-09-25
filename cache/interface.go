package cache

type RedisCache interface {
	Set(key string, value *Desc)
	Get(key string) *Desc
}
