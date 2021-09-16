package report

type Filter interface {
	Apply(Item) bool
}

type StatusFilter struct {
	status int
}
type StatusOnlyFilter StatusFilter

func (sf StatusFilter) Apply(item Item) bool {
	return sf.status != item.source.StatusCode
}
func (sof StatusOnlyFilter) Apply(item Item) bool {
	return sof.status == item.source.StatusCode
}
