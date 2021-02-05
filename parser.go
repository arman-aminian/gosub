package gosub

import (
	"fmt"
	"github.com/arman-aminian/gosub/parsers"
)

func Parse(path string) error {
	srt := parsers.NewSrt()
	lines, err := srt.Parse(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(lines[2].WPM)
	fmt.Println(lines[2].Text)
	return nil
}
