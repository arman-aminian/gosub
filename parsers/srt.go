package parsers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Srt struct {
	lines []Line
}

func NewSrt() *Srt {
	return &Srt{
		lines: []Line{},
	}
}

func (*Srt) Parse(path string) error {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("next line:")
		fmt.Println(scanner.Text())
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, SrtTimestampSeparator) {
			if scanner.Scan() {
				text := scanner.Text()
				fmt.Println("value : ", text)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
