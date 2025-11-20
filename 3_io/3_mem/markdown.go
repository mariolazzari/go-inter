package markdown

import (
	"bytes"
	"fmt"
)

// List renders a slice of item to a markdown list.
func List(items []string) string {
	var buf bytes.Buffer

	for _, item := range items {
		fmt.Fprintf(&buf, "- %s\n", item)
	}
	return buf.String()
}
