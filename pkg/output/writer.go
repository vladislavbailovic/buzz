package output

import (
	"buzz/pkg/report"
	"fmt"
)

type Writer struct {
}

func (w Writer) Listen(what ...interface{}) {
	assembly := what[0].(*report.Assembly)
	for idx, item := range assembly.GetItems() {
		fmt.Printf("[%d]: %s\n", idx, Format(item))
	}
}
