package source

import (
	"strconv"
	"testing"
)

func TestRandomIntHasRange(t *testing.T) {
	gx := NewRandomInt(1)
	if !gx.HasNext() {
		t.Fatalf("should have next initially")
	}
	_ = gx.GetNext()
	if gx.HasNext() {
		t.Fatalf("should have reached limit")
	}

	gx.Reset()
	if !gx.HasNext() {
		t.Fatalf("resetting the generator should have moved cursor")
	}
}

func TestRandomIntGeneratesNumbers(t *testing.T) {
	limit := 161
	gx := NewRandomInt(limit)
	for i := 0; i < limit; i++ {
		val := gx.GetNext()
		_, err := strconv.Atoi(val)
		if err != nil {
			t.Fatalf("expected number :%s", val)
		}
	}
}

func TestRandomIntGeneratesNumbersDifferentThanPrevious(t *testing.T) {
	limit := 161
	gx := NewRandomInt(limit)
	old := gx.GetNext()
	for i := 0; i < limit; i++ {
		val := gx.GetNext()
		if old == val {
			t.Fatalf("new value should be different than old: %s", val)
		}
	}
}

func TestRandomIntGeneratesNumbersWithinRange(t *testing.T) {
	limit := 161
	upper := 20
	lower := 10
	gx := NewRandomInt(limit, upper, lower)
	for i := 0; i < limit; i++ {
		val := gx.GetNext()
		actual, _ := strconv.Atoi(val)
		if actual < lower {
			t.Fatalf("should be above: %s (%d)", val, lower)
		}
		if actual > upper {
			t.Fatalf("should be below: %s (%d)", val, upper)
		}
	}
}

func TestRandomIntGeneratesNumbersDifferentThanPreviousWithinRange(t *testing.T) {
	limit := 161
	upper := 20
	lower := 10
	gx := NewRandomInt(limit, upper, lower)
	old := gx.GetNext()
	for i := 0; i < limit; i++ {
		val := gx.GetNext()
		if old == val {
			t.Fatalf("new value should be different than old: %s (attempt %d)", val, i)
		}
		old = val
	}
}

func TestRandomIntSize(t *testing.T) {
	limit := 161
	gx := NewRandomInt(limit)

	if gx.Size() != limit {
		t.Fatalf("random int size querying fail")
	}
}
