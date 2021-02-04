package parsers

import "time"

type Line struct {
	Text  string
	Start time.Time
	End   time.Time
	WPM   int
}

func NewLine(text string, start, end time.Time, wpm int) *Line {
	return &Line{
		Text:  text,
		Start: start,
		End:   end,
		WPM:   wpm,
	}
}
