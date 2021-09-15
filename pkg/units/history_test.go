package units

import "testing"

func TestHistoryUpdates(t *testing.T) {
	h := NewHistory()
	expected := "test"
	if h.IsKnown(expected) {
		t.Fatalf("history should be empty at first")
	}

	h.Update(expected)
	if !h.IsKnown(expected) {
		t.Fatalf("history should be updated")
	}
}

func TestLongHistoryUpdates(t *testing.T) {
	h := NewLongHistory()
	expected := "test"
	if h.IsKnown(expected) {
		t.Fatalf("history should be empty at first")
	}

	h.Update(expected)
	if !h.IsKnown(expected) {
		t.Fatalf("history should be updated")
	}
}
