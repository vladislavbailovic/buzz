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

type LongHistory struct {
	prev map[string]bool
}

func (h *LongHistory) Update(val string) {
	h.prev[val] = true
}
func (h LongHistory) IsKnown(val string) bool {
	test, exists := h.prev[val]
	if !exists {
		return false
	}
	return test
}
func NewLongHistory() LongHistory {
	return LongHistory{map[string]bool{}}
}
