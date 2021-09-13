package request

import (
	"encoding/json"
	"fmt"
	"strings"
	"buzz/pkg/source"
)

func Replace(haystack string, sources []source.Generator) string {
	for idx, source := range sources {
		for source.HasNext() {
			rpl := source.GetNext()
			needle := fmt.Sprintf("SOURCE_%d", idx+1)
			haystack = strings.Replace(haystack, needle, rpl, -1)
		}
	}
	return haystack
}

func ReplaceJson(haystack map[string]interface{}, sources []source.Generator) string {
	jsonString, _ := json.Marshal(haystack)
	return Replace(string(jsonString), sources)
}
