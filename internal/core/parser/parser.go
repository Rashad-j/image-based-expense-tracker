package parser

// Expenses represents the output structure.
type Expenses struct {
	Items []struct {
		Name  string `json:"name"`
		Price string `json:"price"`
	} `json:"items"`
	Total string `json:"total"`
}

// Parser is the interface for all parsers.
type Parser interface {
	Parse(ocrText string) (string, error) // Parses OCR text and returns JSON string of Expenses.
}
