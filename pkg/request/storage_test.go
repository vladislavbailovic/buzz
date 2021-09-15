package request

import (
	"net/http"
	"testing"
)

func TestRequestKeyGeneration(t *testing.T) {
	rs := RequestStorage{}
	req1, _ := http.NewRequest("GET", "http://test.com", nil)
	req2, _ := http.NewRequest("POST", "http://test.com", nil)

	key1 := rs.getKey(req1)
	if key1 == "" {
		t.Fatalf("failed serializing request 1")
	}

	key2 := rs.getKey(req2)
	if key2 == "" {
		t.Fatalf("failed serializing request 1")
	}

	if key1 == key2 {
		t.Fatalf("keys should differ between GET and POST: %s", key1)
	}
}

func TestRequestStorageHistoryUpdates(t *testing.T) {
	rs := newRequestStorage([]*http.Request{})
	req, _ := http.NewRequest("GET", "http://test.com", nil)

	if rs.isProcessed(req) {
		t.Fatalf("should not be processed initially")
	}

	rs.update(req)
	if !rs.isProcessed(req) {
		t.Fatalf("should be processed post-update")
	}
}
