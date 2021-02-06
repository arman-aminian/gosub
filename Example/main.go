package main

import (
	"fmt"
	"github.com/arman-aminian/gosub"
	"github.com/arman-aminian/gosub/cleaners"
	"github.com/arman-aminian/gosub/parsers"
)

func main() {
	s, err := gosub.Parse("./Example/test.srt")
	if err != nil {
		panic(err)
	}

	fmt.Println(parsers.CalculateMaxWpm(s, 0, s.Size))
	fmt.Println(parsers.CalculateMinWpm(s, 0, s.Size))
	fmt.Println(parsers.CalculateMeanWpm(s, 0, s.Size))

	fmt.Println(cleaners.LowerCase(s.GetLines()))
	fmt.Println(cleaners.RemoveBrackets(s.GetLines()))
	fmt.Println(cleaners.RemoveTags(s.GetLines()))
}
