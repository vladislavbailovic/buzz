package challenge

import (
	"fmt"
	"strings"
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

func TestDoubleFuzz(t *testing.T) {
	gx1 := source.NewRandomInt(10)
	gx2 := source.NewRandomInt(10)
	b := NewRequest("SOURCE_1/SOURCE_2.com")
	f := Fuzzer{original: b}

	actual := f.Fuzz([]source.Generator{gx1, gx2})
	if len(actual) != 10 {
		t.Fatalf("expected set amount of fuzzed requests")
	}
	for idx, req := range actual {
		if strings.Contains(req.Url, "SOURCE_") {
			t.Fatalf("position %d: expected fuzzed url: %s", idx, req.Url)
		}
	}
}

func TestMultiFuzz(t *testing.T) {
	gx := source.NewRandomInt(10)
	sources := []source.Generator{gx}
	rqpaths := []string{"SOURCE_1"}
	var request Request

	for i := 1; i < 25; i++ {
		rqpaths = append(rqpaths, fmt.Sprintf("SOURCE_%d", i))
		sources = append(sources, source.NewRandomInt(10))
		request = NewRequest(strings.Join(rqpaths, "/"))
		f := Fuzzer{original: request}

		actual := f.Fuzz(sources)
		if len(actual) != 10 {
			t.Fatalf("expected set amount of fuzzed requests")
		}
		for idx, req := range actual {
			t.Log(req.Url)
			if strings.Contains(req.Url, "SOURCE_") {
				t.Fatalf("position %d: expected fuzzed url: %s", idx, req.Url)
			}
		}
	}
}
