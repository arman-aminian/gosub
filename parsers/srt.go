package parsers

import (
	"bufio"
	"errors"
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

func (s *Srt) GetLines() []Line {
	return s.lines
}

func (*Srt) Parse(path string) ([]Line, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, SrtTimestampSeparator) {
			var l Line
			l.Start, l.End, err = parseTimestamps(line)
			if err != nil {
				return nil, err
			}

			if scanner.Scan() {
				l.Text = scanner.Text()
			}

			lines = append(lines, l)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
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

	sl := strings.Split(s, ",")
	if len(sl[1]) > 3 {
		sl[1] = sl[1][0:3]
	}
	for len(sl[1]) < 3 {
		sl[1] = sl[1] + "0"
	}
	return sl[0] + "." + sl[1]
}
