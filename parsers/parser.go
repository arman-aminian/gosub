package parsers

import "math"

type Parser interface {
	Parse(string) error
	GetLines() []Line
}

func CalculateMaxWpm(p Parser, from, to int) int {
	lines := p.GetLines()[from:to]
	maxWpm := 0
	for _, line := range lines {
		if line.WPM > maxWpm {
			maxWpm = line.WPM
		}
	}
	return maxWpm
}

func CalculateMinWpm(p Parser, from, to int) int {
	lines := p.GetLines()[from:to]
	minWpm := math.MaxInt32
	for _, line := range lines {
		if line.WPM < minWpm {
			minWpm = line.WPM
		}
	}
	return minWpm
}

func CalculateMeanWpm(p Parser, from, to int) int {
	lines := p.GetLines()[from:to]
	sum := 0
	for _, line := range lines {
		sum += line.WPM
	}
	return sum / len(lines)
}

func TotalWordCount(p Parser) int {
	sum := 0
	for _, line := range p.GetLines() {
		sum += WordsCount(line.Text)
	}
	return sum
}
