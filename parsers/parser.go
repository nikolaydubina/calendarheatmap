package parsers

// Parser transform input files into format ready for render
type Parser interface {
	Parse(data []byte) (int, map[int]int, error)
}
