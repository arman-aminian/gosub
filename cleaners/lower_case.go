package cleaners

import (
	"github.com/arman-aminian/gosub/parsers"
	"strings"
)

func LowerCase(lines []parsers.Line) []parsers.Line {
	for i := range lines {
		lines[i].Text = strings.ToLower(lines[i].Text)
	}
	return lines
}
