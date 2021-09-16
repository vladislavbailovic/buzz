package challenge

import (
	"testing"
)

func TestRunRequestUpdatesStoreHistory(t *testing.T) {
	req := NewRequest("http://test.com")
	rnr := NewRunner([]Request{req}, NewClient())

	if rnr.store.isProcessed(req) {
		t.Fatalf("should not be processed")
	}

	response1 := rnr.runRequest(req)

	if response1.StatusCode == STATUS_ALREADY_REQUESTED {
		t.Fatalf("expected status code to NOT be %d, got %d", STATUS_ALREADY_REQUESTED, response1.StatusCode)
	}

	if !rnr.store.isProcessed(req) {
		t.Fatalf("should now be processed")
	}
	response2 := rnr.runRequest(req)

	if response2.StatusCode != STATUS_ALREADY_REQUESTED {
		t.Fatalf("expected status code to be %d, got %d", STATUS_ALREADY_REQUESTED, response2.StatusCode)
	}
}

func TestRunBatchWithKnownRequests(t *testing.T) {
	req := NewRequest("http://test.com")
	batch := []Request{req, req, req}

	rnr := NewRunner(batch, NewClient())
	rnr.store.update(req)

	for idx, resp := range rnr.runBatch(batch) {
		if resp.StatusCode != STATUS_ALREADY_REQUESTED {
			t.Fatalf("expected %d to already be requested, got %d", idx, resp.StatusCode)
		}
	}
}

func TestRunWithKnownRequestsBatchSizes(t *testing.T) {
	req := NewRequest("http://test.com")

	for i := 1; i < 161; i += 5 {
		max := i % 16
		if max == 0 {
			continue
		}
		batch := []Request{}
		for j := 0; j < i; j++ {
			batch = append(batch, req)
		}
		rnr := NewRunner(batch, NewClient())
		rnr.batchSize = max

		result := rnr.Run()
		if len(result) != len(batch) {
			t.Fatalf("expected whole batch to be processed (%d) - got %d", len(batch), len(result))
		}
	}

}
