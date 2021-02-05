package parsers

type parser interface {
	Parse(string) ([]Line, error)
}
