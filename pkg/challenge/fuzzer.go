package challenge

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"buzz/pkg/source"
)

type Fuzzer struct {
	original Request
}

func (fuzz Fuzzer) getShortestSourceSize(sources []source.Generator) int {
	shortest := float64(sources[0].Size())
	for _, source := range sources {
		shortest = math.Min(float64(shortest), float64(source.Size()))
	}
	return int(shortest)
}

func (fuzz Fuzzer) Fuzz(sources []source.Generator) []Request {
	var result []Request
	var tmp Request
	haystack, _ := json.Marshal(fuzz.original)
	shortest := fuzz.getShortestSourceSize(sources)

	for i := 0; i < shortest; i++ {
		replaced := string(haystack)
		for idx, source := range sources {
			for source.HasNext() {
				rpl := source.GetNext()
				needle := fmt.Sprintf("SOURCE_%d", idx+1)
				replaced = strings.Replace(replaced, needle, rpl, -1)
			}
			source.Reset()
		}
		json.Unmarshal([]byte(replaced), &tmp)
		result = append(result, tmp)
	}
	return result
}
