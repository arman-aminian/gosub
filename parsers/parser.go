package parsers

type parser interface {
	parse(string) error
}
