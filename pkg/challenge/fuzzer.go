package challenge

import (
	"encoding/json"
	"fmt"
	"strings"

	"buzz/pkg/source"
)

type Fuzzer struct {
	original Request
}

func (fuzz Fuzzer) Fuzz(sources []source.Generator) []Request {
	var result []Request
	var tmp Request
	original, _ := json.Marshal(fuzz.original)

	for _, replacements := range carthesianProduct(sources) {
		haystack := string(original)
		for idx, rpl := range replacements {
			needle := fmt.Sprintf("SOURCE_%d", idx+1)
			haystack = strings.Replace(haystack, needle, rpl, -1)
		}
		json.Unmarshal([]byte(haystack), &tmp)
		result = append(result, tmp)
	}

	return result
}

func getSources(sources []source.Generator) [][]string {
	result := make([][]string, 0)
	for _, source := range sources {
		lines := []string{}
		for source.HasNext() {
			lines = append(lines, source.GetNext())
		}
		source.Reset()
		if len(lines) == 0 {
			continue
		}
		result = append(result, lines)
	}
	return result
}

func carthesianProduct(raw []source.Generator) [][]string {
	sources := getSources(raw)
	if len(sources) == 0 {
		return [][]string{}
	}

	result := [][]string{}
	lenghts := func(i int) int { return len(sources[i]) }
	for ix := make([]int, len(sources)); ix[0] < lenghts(0); nextIndex(ix, lenghts) {
		var temp []string
		for j, k := range ix {
			temp = append(temp, sources[j][k])
		}
		result = append(result, temp)
	}
	return result
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}
