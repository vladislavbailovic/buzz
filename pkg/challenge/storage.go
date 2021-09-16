package challenge

import (
	"buzz/pkg/units"
	"encoding/json"
)

type RequestStorage struct {
	queue   []Request
	history *units.LongHistory
}

func newRequestStorage(reqs []Request) RequestStorage {
	history := units.NewLongHistory()
	return RequestStorage{queue: reqs, history: &history}
}

func (rs RequestStorage) getKey(req Request) string {
	key, _ := json.Marshal(req)
	return string(key)
}

func (rs RequestStorage) isProcessed(req Request) bool {
	return rs.history.IsKnown(rs.getKey(req))
}

func (rs RequestStorage) update(req Request) {
	rs.history.Update(rs.getKey(req))
}
