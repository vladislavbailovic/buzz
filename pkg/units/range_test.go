package units

import "testing"

func TestCursorUpdates(t *testing.T) {
	expected := 161
	cursor := NewCursor(expected)

	if 0 != cursor.GetPos() {
		t.Fatalf("initial position should be zero")
	}
	if cursor.IsEof() {
		t.Fatalf("should not be at eof initially %d %d", cursor.GetPos(), cursor.GetSize())
	}

	actual := 0
	for !cursor.IsEof() {
		actual++
		cursor.Advance()
	}

	if actual != expected {
		t.Fatalf("cursor pos %d (%d) did not advance to expected %d", actual, cursor.GetPos(), expected)
	}

	if cursor.GetSize() != expected {
		t.Fatalf("cursor size %d larger than expected %d", cursor.GetSize(), expected)
	}

	if !cursor.IsEof() {
		t.Fatalf("should be at eof")
	}

	cursor.Reset()
	if cursor.IsEof() {
		t.Fatalf("should not be at eof after reset")
	}
}

func TestBoundsCreationAllZeros(t *testing.T) {
	b0 := NewBounds([]int{})
	if b0.GetMin() != 0 || b0.GetMax() != 0 {
		t.Fatalf("invalid zero boundaries")
	}
}

func TestBoundsCreationJustTop(t *testing.T) {
	b := NewBounds([]int{1})
	if b.GetMin() != 0 || b.GetMax() != 1 {
		t.Fatalf("invalid max boundaries: %d %d", b.GetMin(), b.GetMax())
	}
}

func TestBoundsCreationBoth(t *testing.T) {
	b := NewBounds([]int{2, 1})
	if b.GetMin() != 1 || b.GetMax() != 2 {
		t.Fatalf("invalid boundaries: %d %d", b.GetMin(), b.GetMax())
	}

	if 1 != b.GetDiff() {
		t.Fatalf("invalid diff")
	}
}
