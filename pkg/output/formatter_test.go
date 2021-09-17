package output

import (
	"buzz/pkg/challenge"
	"buzz/pkg/report"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	req := challenge.NewRequest("test.com")
	resp := challenge.Response{StatusCode: 403}

	str := Format(report.NewItem(req, resp))
	if !strings.Contains(str, "test.com") {
		t.Fatalf("expected request URL in formatted item")
	}
	if !strings.Contains(str, "403") {
		t.Fatalf("expected response status in formatted item")
	}
}
