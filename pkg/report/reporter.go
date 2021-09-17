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
	resp := what[0].(challenge.Response)
	r.report.AddResponse(resp)
	r.Publish(EVT_ITEM_ADDED, r.report)
}

func NewReporter() Reporter {
	report := Assembly{}
	return Reporter{report: &report}
}
