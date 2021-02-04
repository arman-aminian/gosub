package main

import (
	"github.com/arman-aminian/gosub"
)

func main() {
	err := gosub.Parse("./Example/Chainz.srt")
	if err != nil {
		panic(err)
	}
}
