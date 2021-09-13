package request

import (
	"buzz/pkg/source"
	"fmt"
	"strings"
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
