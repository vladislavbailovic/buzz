package challenge

import "net/http"

const (
	STATUS_ALREADY_REQUESTED = iota
	STATUS_PASSTHROUGH
	STATUS_FAILED
)

type Response struct {
	StatusCode int
}

func (r Response) Build() http.Response {
	return http.Response{StatusCode: r.StatusCode}
}

func NewResponseFromHttp(resp *http.Response) Response {
	return Response{StatusCode: resp.StatusCode}
}
