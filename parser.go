package gosub

import (
	"github.com/arman-aminian/gosub/parsers"
)

func Parse(path string) (*parsers.Srt, error) {
	srt := parsers.NewSrt()
	err := srt.Parse(path)
	if err != nil {
		return nil, err
	}

	//max := cleaners.RemoveTags(srt.Lines)

	//max := parsers.CalculateMaxWpm(srt, 0, len(srt.Lines))
	//fmt.Println(max)
	//mean := parsers.CalculateMeanWpm(srt, 0, len(srt.Lines))
	//fmt.Println(mean)

	//fmt.Println(srt.Lines[2].WPM)
	//fmt.Println(srt.Lines[2].Text)
	return srt, nil
}

func ParseByFile(f []byte) (*parsers.Srt, error) {
	srt := parsers.NewSrt()
	err := srt.ParseByFile(f)
	if err != nil {
		return nil, err
	}

	//max := cleaners.RemoveTags(srt.Lines)

	//max := parsers.CalculateMaxWpm(srt, 0, len(srt.Lines))
	//fmt.Println(max)
	//mean := parsers.CalculateMeanWpm(srt, 0, len(srt.Lines))
	//fmt.Println(mean)

	//fmt.Println(srt.Lines[2].WPM)
	//fmt.Println(srt.Lines[2].Text)
	return srt, nil
}
