package challenge

import "testing"

func TestNewClientReturnsPassthroughClientByDefault(t *testing.T) {
	var nc Client
	nc = NewClient()

	_, ok1 := nc.(WebClient)
	if ok1 {
		t.Fatalf("should not be web client by default")
	}

	_, ok2 := nc.(PassthroughClient)
	if !ok2 {
		t.Fatalf("should be passthrough client by default")
	}
}

func TestNewClientReturnsPassthroughClientWhenRequested(t *testing.T) {
	var nc Client
	nc = NewClient(CLIENT_PASSTHROUGH)

	_, ok1 := nc.(WebClient)
	if ok1 {
		t.Fatalf("should not be web client when passthrough requested")
	}

	_, ok2 := nc.(PassthroughClient)
	if !ok2 {
		t.Fatalf("should be passthrough client when requested")
	}
}

func TestNewClientReturnsWebClientWhenRequested(t *testing.T) {
	var nc Client
	nc = NewClient(CLIENT_WEB)

	_, ok1 := nc.(PassthroughClient)
	if ok1 {
		t.Fatalf("should not be passthrough client when web requested")
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
