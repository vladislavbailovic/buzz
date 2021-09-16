package challenge

import (
	"sync"
)

type Runner struct {
	store        *RequestStorage
	batchSize    int
	currentBatch int
	client       Client
}

func NewRunner(reqs []Request, client Client) Runner {
	store := newRequestStorage(reqs)
	return Runner{store: &store, batchSize: 5, client: client}
}

func (rnr *Runner) Run() []Response {
	var batch []Request
	var results []Response

	for _, req := range rnr.store.queue {
		batch = append(batch, req)
		if len(batch) == rnr.batchSize {
			for _, resp := range rnr.runBatch(batch) {
				results = append(results, resp)
			}
			batch = []Request{}
		}
	}

	if len(batch) > 0 {
		for _, resp := range rnr.runBatch(batch) {
			results = append(results, resp)
		}
	}

	return results
}

func (rnr *Runner) runBatch(reqs []Request) []Response {
	var result []Response
	var wg sync.WaitGroup
	var mux sync.Mutex

	rnr.currentBatch += 1
	for _, req := range reqs {
		wg.Add(1)
		go func(req Request) {
			defer wg.Done()
			defer mux.Unlock()

			mux.Lock()
			response := rnr.runRequest(req)
			result = append(result, response)
		}(req)
	}

	wg.Wait()
	return result
}

func (rnr *Runner) runRequest(req Request) Response {
	if rnr.store.isProcessed(req) {
		return Response{StatusCode: STATUS_ALREADY_REQUESTED}
	}

	rnr.store.update(req)
	return rnr.client.Send(req)
}
