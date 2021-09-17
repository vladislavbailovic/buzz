package challenge

import (
	"fmt"
	"math"
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
	expected := int(math.Pow(float64(10), float64(2)))
	if len(actual) != expected {
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

	for i := 2; i < 5; i++ {
		rqpaths = append(rqpaths, fmt.Sprintf("SOURCE_%d", i))
		sources = append(sources, source.NewRandomInt(10))
		request = NewRequest(strings.Join(rqpaths, "/"))
		f := Fuzzer{original: request}

		actual := f.Fuzz(sources)
		expected := int(math.Pow(float64(gx.Size()), float64(i)))
		if len(actual) != expected {
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

func TestMultiFuzzStatic(t *testing.T) {
	test := [][]string{
		[]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
		[]string{"1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th", "10th"},
		[]string{"is1", "is2", "is3", "is4", "is5", "is6", "is7", "is8", "is9", "is10"},
	}
	gx := source.NewStaticList(test[0])
	sources := []source.Generator{gx}
	rqpaths := []string{"SOURCE_1"}
	var request Request

	for i := 2; i < 4; i++ {
		rqpaths = append(rqpaths, fmt.Sprintf("SOURCE_%d", i))
		sources = append(sources, source.NewStaticList(test[i-1]))
		request = NewRequest(strings.Join(rqpaths, "/"))
		f := Fuzzer{original: request}

		actual := f.Fuzz(sources)
		for idx, req := range actual {
			t.Log(req.Url)
			if strings.Contains(req.Url, "SOURCE_") {
				t.Fatalf("position %d: expected fuzzed url: %s", idx, req.Url)
			}
		}
		expected := int(math.Pow(float64(gx.Size()), float64(i)))
		if len(actual) != expected {
			t.Fatalf("expected set amount of fuzzed requests: %d, got %d", expected, len(actual))
		}
	}
}
