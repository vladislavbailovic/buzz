package request

import (
	"encoding/json"
	"testing"

	"buzz/pkg/source"
)

func TestReplace(t *testing.T) {
	gx := source.NewRandomInt(10)
	test := "nanana SOURCE_1 nanana"
	if test == Replace(test, []source.Generator{gx}) {
		t.Fatalf("no replacement took place")
	}
}

func TestReplaceJson(t *testing.T) {
	gx := source.NewRandomInt(10)
	test := []byte(`{ "whatever": "nanana SOURCE_1 nanana" }`)

	var jsonData map[string]interface{}
	json.Unmarshal(test, &jsonData)
	if string(test) == ReplaceJson(jsonData, []source.Generator{gx}) {
		t.Fatalf("no replacement took place")
	}
}
