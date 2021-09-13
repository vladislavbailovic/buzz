package request

import (
	"testing"

	"buzz/pkg/source"
)

func TestInit(t *testing.T) {
	gx := source.NewRandomInt(10)
	test := "nanana SOURCE_1 nanana"
	if test == Replace(test, []source.Generator{gx}) {
		t.Fatalf("no replacement took place")
	}
}
