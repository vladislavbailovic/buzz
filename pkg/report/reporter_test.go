package report

import (
	"buzz/pkg/challenge"
	"buzz/pkg/units"
	"testing"
)

func TestInit(t *testing.T) {
	r := NewReporter()
	cast := units.Broadcaster{}
	cast.Subscribe("test", r)

	if 0 != r.report.RawSize() {
		t.Fatalf("report should have no items initially")
	}

	cast.Publish("test", challenge.NewRequest("test.com"), challenge.Response{StatusCode: 200})
	if 0 == r.report.RawSize() {
		t.Fatalf("report should have an item after event")
	}
}

func TestRunnerBroadcasts(t *testing.T) {
	report := NewReporter()
	req := challenge.NewRequest("http://test.com")
	batch := []challenge.Request{req, req, req}

	rnr := challenge.NewRunner(batch, challenge.NewClient())
	rnr.Subscribe(challenge.EVT_REQUEST, report)

	rnr.Run()
	if report.report.RawSize() != len(batch) {
		t.Fatalf("report size does not match request queue")
	}
}
