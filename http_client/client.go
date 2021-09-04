package http_client

import (
	"net/http"
	"net/url"
	"time"
)

//go:generate go run "github.com/vektra/mockery/cmd/mockery" -case=underscore -outpkg mock_client -output ../mock/mock_client -name=Client

type Client interface {
	Do(request *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	PostForm(uri string, data url.Values) (resp *http.Response, err error)
}

var _ Client = &http.Client{}

func New() Client {
	return makeHttpClient()
}

func makeHttpClient() *http.Client {
	transport := &http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	return &http.Client{
		Transport: transport,
	}
}
