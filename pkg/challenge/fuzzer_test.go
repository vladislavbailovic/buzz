package challenge

import (
	"testing"

	"buzz/pkg/source"
)

func TestFuzzer(t *testing.T) {
	gx := source.NewRandomInt(10)
	b := NewRequest("SOURCE_1.com")
	f := Fuzzer{original: b}

	actual := f.Fuzz([]source.Generator{gx})
	if len(actual) != 10 {
		t.Fatalf("expected set amount of fuzzed requests")
	}
	for _, req := range actual {
		if req.Url == "SOURCE_1.com" {
			t.Fatalf("expected fuzzed url")
		}
	}
}
