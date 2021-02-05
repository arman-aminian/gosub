package main

import "github.com/arman-aminian/gosub"

func main() {
	//str := "asdfghjk"
	//str = str[0:4]
	//println(str)
	//t, err := time.Parse("15:04:05.000", "00:03:47.305")
	//t2, err := time.Parse("15:04:05.000", "00:03:44.305")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(t.Sub(t2).Seconds())
	//fmt.Println(t.Sub(t2).Minutes())
	//x := t.Sub(t2).Minutes()
	//y := 1
	//z := float64(y) / x
	//println(z)
	//if t2.After(t) {
	//	println("inja")
	//}
	err := gosub.Parse("./Example/adele.srt")
	if err != nil {
		panic(err)
	}
}
