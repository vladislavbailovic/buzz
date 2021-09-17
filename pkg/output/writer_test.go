package output

import (
	"buzz/pkg/challenge"
	"buzz/pkg/report"
	"testing"
)

func TestRunnerBroadcasts(t *testing.T) {
	reporter := report.NewReporter()
	writer := Writer{}
	reporter.Subscribe(report.EVT_ITEM_ADDED, writer)

	req := challenge.NewRequest("http://test.com")
	batch := []challenge.Request{req, req, req}

	rnr := challenge.NewRunner(batch, challenge.NewClient())
	rnr.Subscribe(challenge.EVT_REQUEST, reporter)

	rnr.Run()
}
