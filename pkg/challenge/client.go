package challenge

import (
	"net/http"
	"time"
)

type ClientType string

const (
	CLIENT_NULL ClientType = "null"
	CLIENT_WEB  ClientType = "web"
)

type Client interface {
	Send(Request) Response
}

func NewClient(ctype ...ClientType) Client {
	var clientType ClientType
	if len(ctype) > 0 {
		clientType = ctype[0]
	}

	if CLIENT_WEB == clientType {
		client := &http.Client{Timeout: time.Second * 10}
		return WebClient{http: client}
	}
	return NullClient{}
}

type NullClient struct{}

func (nc NullClient) Send(req Request) Response {
	return Response{StatusCode: STATUS_PASSTHROUGH}
}

type WebClient struct {
	http *http.Client
}

func (wc WebClient) Send(req Request) Response {
	request := req.Build()
	response, err := wc.http.Do(&request)
	if err != nil {
		return Response{StatusCode: STATUS_FAILED}
	}
	return NewResponseFromHttp(response)
}
