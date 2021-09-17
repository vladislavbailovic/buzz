package source

import "testing"

func TestStaticList(t *testing.T) {
	test := []string{
		"one", "two", "three",
	}
	src := NewStaticList(test)
	idx := 0
	for src.HasNext() {
		actual := src.GetNext()
		if test[idx] != actual {
			t.Fatalf("%d expected %s - got %s", idx, test[idx], actual)
		}
		idx++
	}

	if len(test) != src.Size() {
		t.Fatalf("error determining list size")
	}

	if src.HasNext() {
		t.Fatalf("should be at list end")
	}
	src.Reset()
	if !src.HasNext() {
		t.Fatalf("should be reset")
	}
}
