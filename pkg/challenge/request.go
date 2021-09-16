package challenge

import (
	"net/http"
)

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
}

func (b Request) Build() http.Request {
	r, _ := http.NewRequest(b.Method, b.Url, nil)
	for key, val := range b.Headers {
		r.Header.Set(key, val)
	}
	return *r
}

func NewRequest(url string) Request {
	return Request{
		Method:  "GET",
		Url:     url,
		Headers: map[string]string{"user-agent": "test"},
	}
}

func NewPostRequest(url string) Request {
	return Request{
		Method:  "POST",
		Url:     url,
		Headers: map[string]string{"user-agent": "test"},
	}
}
