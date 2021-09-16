package challenge

import (
	"encoding/json"
	"fmt"
	"strings"

	"buzz/pkg/source"
)

type Fuzzer struct {
	original Data
}

func (fuzz Fuzzer) Fuzz(sources []source.Generator) []Data {
	var result []Data
	var tmp Data
	haystack, _ := json.Marshal(fuzz.original)
	for idx, source := range sources {
		for source.HasNext() {
			rpl := source.GetNext()
			needle := fmt.Sprintf("SOURCE_%d", idx+1)
			replaced := strings.Replace(string(haystack), needle, rpl, -1)
			json.Unmarshal([]byte(replaced), &tmp)
			result = append(result, tmp)
		}
	}
	return result
}
