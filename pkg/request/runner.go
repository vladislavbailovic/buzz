package request

import (
	"net/http"
)

const (
	STATUS_ALREADY_REQUESTED = iota
	STATUS_PASSTHROUGH
	STATUS_FAILED
)

type Runner struct {
	store        *RequestStorage
	batchSize    int
	currentBatch int
}

func NewRunner(reqs []*http.Request) Runner {
	store := newRequestStorage(reqs)
	return Runner{store: &store, batchSize: 5}
}

func (rnr *Runner) Run() []http.Response {
	var batch []*http.Request
	var results []http.Response

	for _, req := range rnr.store.queue {
		batch = append(batch, req)
		if len(batch) == rnr.batchSize {
			for _, resp := range rnr.runBatch(batch) {
				results = append(results, resp)
			}
			batch = []*http.Request{}
		}
	}

	if len(batch) > 0 {
		for _, resp := range rnr.runBatch(batch) {
			results = append(results, resp)
		}
	}

	return results
}

func (rnr *Runner) runBatch(reqs []*http.Request) []http.Response {
	var result []http.Response
	listener := make(chan http.Response)

	rnr.currentBatch += 1
	for _, req := range reqs {
		go rnr.runRequest(req, listener)
	}

	for r := range listener {
		result = append(result, r)
		if len(result) == len(reqs) {
			close(listener)
		}
	}

	return result
}

func (rnr *Runner) runRequest(req *http.Request, listener chan http.Response) {
	if rnr.store.isProcessed(req) {
		listener <- http.Response{StatusCode: STATUS_ALREADY_REQUESTED}
		return
	}

	rnr.store.update(req)
	// @TODO: do actual request

	listener <- http.Response{StatusCode: STATUS_PASSTHROUGH}
}
