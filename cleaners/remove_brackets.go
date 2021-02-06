package cleaners

import (
	"github.com/arman-aminian/gosub/parsers"
	"regexp"
	"strings"
)

func RemoveBrackets(lines []parsers.Line) []parsers.Line {
	re := regexp.MustCompile(`\[[^[]*\]`)
	for i := range lines {
		m := re.FindAllString(lines[i].Text, -1)
		for _, b := range m {
			lines[i].Text = strings.ReplaceAll(lines[i].Text, b, "")
		}
	}
	return lines
}
