package cleaners

import (
	"github.com/arman-aminian/gosub/parsers"
	"regexp"
	"strings"
)

func RemoveTags(lines []parsers.Line) []parsers.Line {
	reO := regexp.MustCompile(`<[^[]*>`)
	reC := regexp.MustCompile(`</[^[]*>`)
	for i := range lines {
		m := reO.FindAllString(lines[i].Text, -1)
		m = append(m, reC.FindAllString(lines[i].Text, -1)...)
		for _, b := range m {
			lines[i].Text = strings.ReplaceAll(lines[i].Text, b, "")
		}
	}
	return lines
}
