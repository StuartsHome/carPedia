package cache

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
