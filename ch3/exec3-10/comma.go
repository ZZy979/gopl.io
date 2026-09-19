package exec3_10

import (
	"bytes"
	"strings"
)

// Comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	if len(s) == 0 {
		return ""
	}

	var buf bytes.Buffer
	// optional sign
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	// decimal point
	n := strings.LastIndex(s, ".")
	if n == -1 {
		n = len(s)
	}

	// integer part
	first := n % 3
	if first == 0 {
		first = 3
	}
	buf.WriteString(s[:first])
	for i := first; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	// decimal part
	buf.WriteString(s[n:])

	return buf.String()
}
