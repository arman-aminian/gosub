# gosub
Utility to extract the contents of a subtitle file
* `srt`: [SubRip](https://en.wikipedia.org/wiki/SubRip)
* other types will be supported as soon as possible :)

## Installation
Installation is done using `go get`.
```
go get -u github.com/arman-aminian/gosub
```

## Usage
The method parse requires the following parameters:
* `path`: location of the subtitle file.

```golang
s, err := gosub.Parse("./Example/test.srt")
if err != nil {
	panic(err)
}
```
___

## Line Struct

Each line of a subtitle is represented with a `Line` struct with the following properties:

* `start`: timestamp of the start of the dialog.
* `end`: timestamp of the end of the dialog.
* `text`: dialog contents.

```golang
for _, line := range s.Lines {
	fmt.Println(line.Start, " -> ", line.End)
	fmt.Println(line.Text)
	fmt.Println()
}
```

Output:
```text
0000-01-01 00:00:00 +0000 UTC  ->  0000-01-01 00:00:30.05 +0000 UTC
first line [test test] of subtitle

0000-01-01 00:01:21.062 +0000 UTC  ->  0000-01-01 00:01:26.361 +0000 UTC
second line <i> test test<\i> of subtitle

0000-01-01 00:01:26.565 +0000 UTC  ->  0000-01-01 00:01:32.564 +0000 UTC
third line of subtitle
third 2

0000-01-01 00:01:32.768 +0000 UTC  ->  0000-01-01 00:01:37.967 +0000 UTC
Make IT lowerCase

```
## Cleaners

Currently, 3 cleaners are provided:

* `RemoveBrackets` will remove anything between them (e.g., `[test test]`)
* `RemoveTags` will remove formatting keys like `<i>` and `</i>` and also anything between them.
* `LowerCase` will lower case all text. 

## WPM

Calculate wpm(word per minute):

* `CalculateMaxWpm` will find the Max WPM.
* `CalculateMinWpm` will find the min WPM.
* `CalculateMeanWpm` will find the mean WPM.
