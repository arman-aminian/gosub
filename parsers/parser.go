package parsers

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

func CalculateMeanWpm(p Parser, from, to int) int {
	lines := p.GetLines()[from:to]
	sum := 0
	for _, line := range lines {
		sum += line.WPM
	}
	return sum / len(lines)
}
