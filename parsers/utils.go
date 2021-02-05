package parsers

import (
	"regexp"
)

const (
	SrtTimestampSeparator = " --> "
	TimeLayout            = "15:04:05.000"
	InvalidTimeErr        = "invalid timestamps syntax"
)

func WordsCount(s string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(s, -1)
	return len(results)
}
