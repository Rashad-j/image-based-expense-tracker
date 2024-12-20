package parser

import (
	"context"

	expenses "github.com/Rashad-j/image-based-expense-tracker/internal/core/expense"
)

// Parser is the interface for all parsers.
type Parser interface {
	Parse(ctx context.Context, ocrText string) (expenses.Expenses, error)
}
