package request

import (
	"net/http"
)

type Builder struct {
	Method  string
	Url     string
	Headers map[string]string
}

func (b Builder) Build() http.Request {
	r, _ := http.NewRequest(b.Method, b.Url, nil)
	for key, val := range b.Headers {
		r.Header.Set(key, val)
	}
	return *r
}

func NewBuilder(url string) Builder {
	return Builder{
		Method:  "GET",
		Url:     url,
		Headers: map[string]string{"user-agent": "test"},
	}
}
