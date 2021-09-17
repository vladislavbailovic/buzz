package output

import (
	"buzz/pkg/report"
	"fmt"
)

func Format(item report.Item) string {
	return fmt.Sprintf("%s - %d", item.GetSource().Url, item.GetResult().StatusCode)
}
