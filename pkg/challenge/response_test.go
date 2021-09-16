package challenge

import (
	"net/http"
	"testing"
)

func TestResponseBuildReturnsHttpResponse(t *testing.T) {
	var resp http.Response
	r := Response{StatusCode: 200}

	resp = r.Build()
	if resp.StatusCode != r.StatusCode {
		t.Fatalf("status code mismatch")
	}

}
