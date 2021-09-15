package request

import (
	"net/http"
	"testing"
)

func TestRunRequestUpdatesStoreHistory(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://test.com", nil)
	rnr := NewRunner([]*http.Request{req})

	if rnr.store.isProcessed(req) {
		t.Fatalf("should not be processed")
	}

	status1 := make(chan http.Response)
	go rnr.runRequest(req, status1)
	response1 := <-status1

	if response1.StatusCode == STATUS_ALREADY_REQUESTED {
		t.Fatalf("expected status code to NOT be %d, got %d", STATUS_ALREADY_REQUESTED, response1.StatusCode)
	}

	if !rnr.store.isProcessed(req) {
		t.Fatalf("should now be processed")
	}
	status2 := make(chan http.Response)
	go rnr.runRequest(req, status2)
	response2 := <-status2

	if response2.StatusCode != STATUS_ALREADY_REQUESTED {
		t.Fatalf("expected status code to be %d, got %d", STATUS_ALREADY_REQUESTED, response2.StatusCode)
	}
}

func TestRunBatchWithKnownRequests(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://test.com", nil)
	batch := []*http.Request{req, req, req}

	rnr := NewRunner(batch)
	rnr.store.update(req) // so we don't actuall make requests

	for idx, resp := range rnr.runBatch(batch) {
		if resp.StatusCode != STATUS_ALREADY_REQUESTED {
			t.Fatalf("expected %d to already be requested", idx)
		}
	}
}

func TestRunWithKnownRequestsBatchSizes(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://test.com", nil)

	for i := 1; i < 161; i += 5 {
		max := i % 16
		if max == 0 {
			continue
		}
		batch := []*http.Request{}
		for j := 0; j < i; j++ {
			batch = append(batch, req)
		}
		rnr := NewRunner(batch)
		rnr.store.update(req) // so we don't actuall make requests
		rnr.batchSize = max

		result := rnr.Run()
		if len(result) != len(batch) {
			t.Fatalf("expected whole batch to be processed (%d) - got %d", len(batch), len(result))
		}
	}

}
