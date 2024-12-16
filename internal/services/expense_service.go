package services

import (
	"context"
	"fmt"
	"io"

	expenses "github.com/Rashad-j/image-based-expense-tracker/internal/core/expense"
	"github.com/Rashad-j/image-based-expense-tracker/internal/core/ocr"
	"github.com/Rashad-j/image-based-expense-tracker/internal/core/parser"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/config"
)

type ExpenseService struct {
	cfg          config.Config
	OCRProcessor ocr.Processor
	TextParser   parser.Parser
}

func NewExpenseService(cfg config.Config, OCRProcessor ocr.Processor, TextParser parser.Parser) ExpenseService {
	return ExpenseService{
		cfg:          cfg,
		OCRProcessor: OCRProcessor,
		TextParser:   TextParser,
	}
}

// AnalyzeReceipt takes an image and runs OCR on it, followed by parsing the extracted
// text into an Expenses object. It will return an error if either step fails.
func (e ExpenseService) AnalyzeReceipt(ctx context.Context, image io.Reader) (expenses.Expenses, error) {
	text, err := e.OCRProcessor.Process(image)
	if err != nil {
		return expenses.Expenses{}, fmt.Errorf("failed to ocr process in AnalyzeReceipt: %v", err)
	}
	exp, err := e.TextParser.Parse(ctx, text)
	if err != nil {
		return expenses.Expenses{}, fmt.Errorf("failed to parse in AnalyzeReceipt: %v", err)
	}
	return exp, nil
}
