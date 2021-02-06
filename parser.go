package gosub

import (
	"fmt"
	"github.com/arman-aminian/gosub/cleaners"
	"github.com/arman-aminian/gosub/parsers"
)

func Parse(path string) error {
	srt := parsers.NewSrt()
	err := srt.Parse(path)
	if err != nil {
		panic(err)
	}

	max := cleaners.RemoveBrackets(srt.Lines)
	for _, line := range max {
		fmt.Println(line.Text)
	}
	//max := parsers.CalculateMaxWpm(srt, 0, len(srt.Lines))
	//fmt.Println(max)
	//mean := parsers.CalculateMeanWpm(srt, 0, len(srt.Lines))
	//fmt.Println(mean)

	//fmt.Println(srt.Lines[2].WPM)
	//fmt.Println(srt.Lines[2].Text)
	return nil
}
