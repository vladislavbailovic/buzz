package challenge

import (
	"net/http"
)

type Data struct {
	Method  string
	Url     string
	Headers map[string]string
}

func (b Data) Build() http.Request {
	r, _ := http.NewRequest(b.Method, b.Url, nil)
	for key, val := range b.Headers {
		r.Header.Set(key, val)
	}
	return *r
}

func NewData(url string) Data {
	return Data{
		Method:  "GET",
		Url:     url,
		Headers: map[string]string{"user-agent": "test"},
	}
}
