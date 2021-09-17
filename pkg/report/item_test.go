package report

import (
	"buzz/pkg/challenge"
	"testing"
)

func TestItem(t *testing.T) {
	req := challenge.NewRequest("test.com")
	resp := challenge.Response{StatusCode: 302}
	item := NewItem(req, resp)

	if item.GetSource().Url != req.Url {
		t.Fatalf("expected source to be the same as original request")
	}

	if item.GetResult().StatusCode != resp.StatusCode {
		t.Fatalf("expected result to be the same as original response")
	}
}
