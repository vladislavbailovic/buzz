package report

import "buzz/pkg/challenge"

type Assembly struct {
	items   []Item
	filters []Filter
}

type Item struct {
	source challenge.Response
}

type Filter interface {
	Apply(Item) bool
}

func (a *Assembly) AddResponse(resp challenge.Response) {
	item := Item{resp}
	a.Add(item)
}
func (a *Assembly) Add(item Item) {
	a.items = append(a.items, item)
}
func (a Assembly) RawSize() int {
	return len(a.items)
}
func (a Assembly) GetItems() []Item {
	if len(a.filters) == 0 {
		return a.items
	}
	result := []Item{}
	for _, item := range a.items {
		isRuledOut := false
		for _, filter := range a.filters {
			if !filter.Apply(item) {
				isRuledOut = true
				continue
			}
		}
		if !isRuledOut {
			result = append(result, item)
		}
	}
	return result
}
func (a *Assembly) AddFilter(filter Filter) {
	a.filters = append(a.filters, filter)
}
