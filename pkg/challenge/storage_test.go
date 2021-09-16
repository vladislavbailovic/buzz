package challenge

import (
	"testing"
)

func TestRequestKeyGeneration(t *testing.T) {
	rs := RequestStorage{}
	req1 := NewRequest("http://test.com")
	req2 := NewPostRequest("http://test.com")

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
	rs := newRequestStorage([]Request{})
	req := NewRequest("http://test.com")

	if rs.isProcessed(req) {
		t.Fatalf("should not be processed initially")
	}

	rs.update(req)
	if !rs.isProcessed(req) {
		t.Fatalf("should be processed post-update")
	}
}
