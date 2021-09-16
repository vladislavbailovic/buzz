package challenge

import "testing"

func TestNewClientReturnsNullClientByDefault(t *testing.T) {
	var nc Client
	nc = NewClient()

	_, ok1 := nc.(WebClient)
	if ok1 {
		t.Fatalf("should not be web client by default")
	}

	_, ok2 := nc.(NullClient)
	if !ok2 {
		t.Fatalf("should be null client by default")
	}
}

func TestNewClientReturnsNullClientWhenRequested(t *testing.T) {
	var nc Client
	nc = NewClient(CLIENT_NULL)

	_, ok1 := nc.(WebClient)
	if ok1 {
		t.Fatalf("should not be web client when null requested")
	}

	_, ok2 := nc.(NullClient)
	if !ok2 {
		t.Fatalf("should be null client when requested")
	}
}

func TestNewClientReturnsWebClientWhenRequested(t *testing.T) {
	var nc Client
	nc = NewClient(CLIENT_WEB)

	_, ok1 := nc.(NullClient)
	if ok1 {
		t.Fatalf("should not be null client when web requested")
	}

	_, ok2 := nc.(WebClient)
	if !ok2 {
		t.Fatalf("should be web client when requested")
	}
}

func TestWebClientSend(t *testing.T) {
	nc := NewClient(CLIENT_WEB)
	req := NewRequest("whatever")

	resp := nc.Send(req)
	t.Log(resp)
}
