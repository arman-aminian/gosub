package main

import "github.com/arman-aminian/gosub"

func main() {
	//str := "asdfghjk"
	//str = str[0:4]
	//println(str)
	//t, err := time.Parse("15:04:05.000", "00:03:46.307")
	//t2, err := time.Parse("15:04:05.000", "00:03:44.305")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(t.Sub(t2))
	//if t2.After(t) {
	//	println("inja")
	//}
	err := gosub.Parse("./Example/Chainz.srt")
	if err != nil {
		panic(err)
	}
}
