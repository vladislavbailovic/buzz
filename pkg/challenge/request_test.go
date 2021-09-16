package challenge

import (
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
