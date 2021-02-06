package cleaners

import (
	"github.com/arman-aminian/gosub/parsers"
	"strings"
)

func LowerCase(lines []parsers.Line) []parsers.Line {
	for _, l := range lines {
		l.Text = strings.ToLower(l.Text)
	}
	return lines
}
