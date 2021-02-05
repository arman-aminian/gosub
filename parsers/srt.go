package parsers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Srt struct {
	lines []Line
}

func NewSrt() *Srt {
	return &Srt{
		lines: []Line{},
	}
}

func (*Srt) Parse(path string) ([]Line, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
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
			start, end, err := parseTimestamps(line)
			if err != nil {
				return nil, err
			}
			fmt.Println(start, end)
			if scanner.Scan() {
				text := scanner.Text()
				fmt.Println("value : ", text)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nil
}

func parseTimestamps(line string) (time.Time, time.Time, error) {
	times := strings.Split(line, SrtTimestampSeparator)
	if len(times) != 2 {
		return time.Time{}, time.Time{}, errors.New("invalid timestamps")
	}

	times[0] = convertToTimeFormat(times[0])
	times[1] = convertToTimeFormat(times[1])

	s, err := time.Parse("15:04:05.000", times[0])
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("invalid timestamps")
	}

	e, err := time.Parse("15:04:05.000", times[1])
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("invalid timestamps")
	}
	return s, e, nil
}

func convertToTimeFormat(s string) string {
	if !strings.Contains(s, ",") {
		return s + ".000"
	}
	return ""
}
