package report

import "buzz/pkg/challenge"

type Item struct {
	source challenge.Request
	result challenge.Response
}

func NewItem(req challenge.Request, resp challenge.Response) Item {
	return Item{req, resp}
}

func (i Item) GetResult() challenge.Response {
	return i.result
}
func (i Item) GetSource() challenge.Request {
	return i.source
}
