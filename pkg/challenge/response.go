package challenge

import "net/http"

type Response struct {
	StatusCode int
}

func (r Response) Build() http.Response {
	return http.Response{StatusCode: r.StatusCode}
}
