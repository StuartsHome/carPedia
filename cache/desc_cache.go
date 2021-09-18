package cache

type DescCache interface {
	Set(key string, value *Desc)
	Get(key string) *Desc
}
