package units

type History struct {
	prev string
}

func (h *History) Update(val string) {
	h.prev = val
}
func (h History) IsKnown(val string) bool {
	return h.prev == val
}

func NewHistory() History {
	return History{""}
}
