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

	for _, line := range s.Lines {
		fmt.Println(line.Start, " -> ", line.End)
		fmt.Println(line.Text)
		fmt.Println()
	}

	fmt.Println(parsers.CalculateMaxWpm(s, 0, s.Size))
	fmt.Println(parsers.CalculateMinWpm(s, 0, s.Size))
	fmt.Println(parsers.CalculateMeanWpm(s, 0, s.Size))
	fmt.Println(parsers.TotalWordCount(s))

	fmt.Println(cleaners.LowerCase(s.GetLines()))
	fmt.Println()
	fmt.Println(cleaners.RemoveBrackets(s.GetLines()))
	fmt.Println()
	fmt.Println(cleaners.RemoveTags(s.GetLines()))
}
