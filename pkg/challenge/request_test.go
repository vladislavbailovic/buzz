package challenge

import (
	"net/http"
	"testing"
)

func TestNewRequestSetsDefaults(t *testing.T) {
	b := NewRequest("test.com")
	if b.Method != "GET" {
		t.Fatalf("should be get by default")
	}
	_, exists := b.Headers["user-agent"]
	if !exists {
		t.Fatalf("should have UA set")
	}
}

func TestRequestBuildReturnsHttpRequest(t *testing.T) {
	var req http.Request
	b := NewRequest("http://test.com")
	req = b.Build()
	if req.Method != b.Method {
		t.Fatalf("method mismatch")
	}
}
