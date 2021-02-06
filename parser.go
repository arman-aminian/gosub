package gosub

import (
	"fmt"
	"github.com/arman-aminian/gosub/parsers"
)

func Parse(path string) error {
	srt := parsers.NewSrt()
	err := srt.Parse(path)
	if err != nil {
		panic(err)
	}
	max := parsers.CalculateMaxWpm(srt, 0, len(srt.Lines))
	fmt.Println(max)
	//fmt.Println(srt.Lines[2].WPM)
	//fmt.Println(srt.Lines[2].Text)
	return nil
}
