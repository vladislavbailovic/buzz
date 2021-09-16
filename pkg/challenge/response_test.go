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

func TestResponseConversionFactory(t *testing.T) {
	resp := http.Response{StatusCode: 200}
	r := NewResponseFromHttp(&resp)
	if resp.StatusCode != r.StatusCode {
		t.Fatalf("factory status code mismatch")
	}
}
