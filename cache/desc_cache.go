package cache

type DescCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
