package parsers

import (
	"bufio"
	"errors"
	"math"
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

			l.Text = ""
			for scanner.Scan() {
				text := scanner.Text()
				if text == "" {
					break
				}
				if l.Text == "" {
					l.Text += text
				} else {
					l.Text += "\n" + text
				}
			}
			wn := WordsCount(l.Text)
			t := l.End.Sub(l.Start).Minutes()
			l.WPM = int(math.Round(float64(wn) / t))

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
		return time.Time{}, time.Time{}, errors.New(InvalidTimeErr)
	}

	times[0] = convertToTimeFormat(times[0])
	times[1] = convertToTimeFormat(times[1])

	s, err := time.Parse(TimeLayout, times[0])
	if err != nil {
		return time.Time{}, time.Time{}, errors.New(InvalidTimeErr)
	}

	e, err := time.Parse(TimeLayout, times[1])
	if err != nil {
		return time.Time{}, time.Time{}, errors.New(InvalidTimeErr)
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
