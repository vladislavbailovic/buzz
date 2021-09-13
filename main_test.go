package main

import (
	"strconv"
	"testing"
)

func TestRandomIntHasRange(t *testing.T) {
	gx := RandomInt{&Range{1, 0}, Bounds{}, &History{""}}
	if !gx.HasNext() {
		t.Fatalf("should have next initially")
	}
	_ = gx.GetNext()
	if gx.HasNext() {
		t.Fatalf("should have reached limit")
	}
}

func TestRandomIntGeneratesNumbers(t *testing.T) {
	limit := 161
	gx := RandomInt{&Range{limit, 0}, Bounds{}, &History{""}}
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
	gx := RandomInt{&Range{limit, 0}, Bounds{}, &History{""}}
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
	gx := RandomInt{&Range{limit, 0}, Bounds{lower, upper}, &History{""}}
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
	gx := RandomInt{&Range{limit, 0}, Bounds{lower, upper}, &History{""}}
	old := gx.GetNext()
	for i := 0; i < limit; i++ {
		val := gx.GetNext()
		if old == val {
			t.Fatalf("new value should be different than old: %s (attempt %d)", val, i)
		}
		old = val
	}
}
