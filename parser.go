package gosub

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Line struct {
	Text  string
	Start time.Time
	End   time.Time
	WPM   int
}

func Parse(path string) error {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("next line:")
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
