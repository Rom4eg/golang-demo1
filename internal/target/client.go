package target

import (
	"io"
	"net/http"
	"net/url"
)

//go:generate mockgen -typed -destination=mock/client.go -package=mock . Client

type Client interface {
	CloseIdleConnections()
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}
