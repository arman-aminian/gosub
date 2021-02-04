package parsers

type parser interface {
	Parse(string) error
}
