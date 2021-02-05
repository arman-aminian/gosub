package parsers

type Parser interface {
	Parse(string) ([]Line, error)
	GetLines() []Line
}

//func CalculateMaxWpm(p Parser) int {
//	p.g
//}
