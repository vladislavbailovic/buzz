package report

import (
	"buzz/pkg/challenge"
	"buzz/pkg/units"
)

const (
	EVT_ITEM_ADDED units.EventType = "item:added"
)

type Reporter struct {
	units.Broadcaster
	report *Assembly
}

func (r Reporter) Listen(what ...interface{}) {
	req := what[0].(challenge.Request)
	resp := what[1].(challenge.Response)
	r.report.Add(NewItem(req, resp))
	r.Publish(EVT_ITEM_ADDED, r.report)
}

func NewReporter() Reporter {
	report := Assembly{}
	return Reporter{report: &report}
}
