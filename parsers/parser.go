package parsers

type Parser interface {
	Parse(string) error
	GetLines() []Line
}

//func CalculateMaxWpm(p Parser) int {
//	p.g
//}
