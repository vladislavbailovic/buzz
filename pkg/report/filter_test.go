package report

import (
	"buzz/pkg/challenge"
	"testing"
)

func TestStatusFilter(t *testing.T) {
	report := Assembly{}

	accept := challenge.Response{StatusCode: challenge.STATUS_PASSTHROUGH}
	report.AddResponse(accept)
	report.AddResponse(accept)
	report.AddResponse(accept)

	reject := challenge.Response{StatusCode: 200}
	report.AddResponse(reject)

	filter := StatusFilter{200}
	report.AddFilter(filter)

	items := report.GetItems()
	if report.RawSize() == len(items) {
		t.Fatalf("expected some items to be filtered out")
	}
	if len(items) != 3 {
		t.Fatalf("expected some items to stay in")
	}
}

func TestStatusOnlyFilter(t *testing.T) {
	report := Assembly{}

	accept := challenge.Response{StatusCode: challenge.STATUS_PASSTHROUGH}
	report.AddResponse(accept)
	report.AddResponse(accept)
	report.AddResponse(accept)

	reject := challenge.Response{StatusCode: 200}
	report.AddResponse(reject)

	filter := StatusOnlyFilter{200}
	report.AddFilter(filter)

	items := report.GetItems()
	if report.RawSize() == len(items) {
		t.Fatalf("expected some items to be filtered out")
	}
	if len(items) != 1 {
		t.Fatalf("expected some items to stay in")
	}
}
