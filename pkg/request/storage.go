package request

import (
	"net/http"
	"net/http/httputil"

	"buzz/pkg/units"
)

type RequestStorage struct {
	queue   []*http.Request
	history *units.LongHistory
}

func newRequestStorage(reqs []*http.Request) RequestStorage {
	history := units.NewLongHistory()
	return RequestStorage{queue: reqs, history: &history}
}

func (rs RequestStorage) getKey(req *http.Request) string {
	key, _ := httputil.DumpRequestOut(req, true)
	return string(key)
}

func (rs RequestStorage) isProcessed(req *http.Request) bool {
	return rs.history.IsKnown(rs.getKey(req))
}

func (rs RequestStorage) update(req *http.Request) {
	rs.history.Update(rs.getKey(req))
}
