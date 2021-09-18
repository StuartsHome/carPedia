package cache

type Desc struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PostCache struct {
}
