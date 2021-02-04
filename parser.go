package gosub

import (
	"github.com/arman-aminian/gosub/parsers"
)

func Parse(path string) error {
	srt := parsers.NewSrt()
	err := srt.Parse(path)
	if err != nil {
		panic(err)
	}
	return nil
}
