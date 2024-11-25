package parser

// Parser is the interface for all parsers.
type Parser interface {
	Parse(ocrText string) (string, error) // Parses OCR text and returns JSON string of Expenses.
}
