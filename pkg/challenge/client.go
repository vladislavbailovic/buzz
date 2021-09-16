package challenge

type ClientType string

const (
	CLIENT_NULL ClientType = "null"
	CLIENT_WEB  ClientType = "web"
)

type Client interface {
	Send(Request) Response
}

type NullClient struct{}

func (nc NullClient) Send(req Request) Response {
	return Response{StatusCode: STATUS_PASSTHROUGH}
}

func NewClient(ctype ...ClientType) Client {
	var clientType ClientType
	if len(ctype) > 0 {
		clientType = ctype[0]
	}

	if CLIENT_NULL == clientType {
		return NullClient{}
	}
	return NullClient{}
}
