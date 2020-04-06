package common

import (
	"strings"

	te "github.com/muesli/termenv"
)

// KeyValueView renders key-value pairs
func KeyValueView(stuff ...string) string {
	if len(stuff) == 0 {
		return ""
	}
	var s string
	for i := 0; i < len(stuff); i++ {
		if i%2 == 0 {
			// even
			s += te.String("│ ").Foreground(Color("241")).String()
			s += stuff[i] + ": "
			continue
		}
		// odd
		s += te.String(stuff[i]).Foreground(Color(purpleFg)).String()
		s += "\n"
	}
	return strings.TrimSpace(s)
}
