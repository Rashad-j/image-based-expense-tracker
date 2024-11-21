package ocr

import (
	"errors"
	"os/exec"
)

type TesseractsProcessor struct{}

func NewTesseractsProcessor() *TesseractsProcessor {
	return &TesseractsProcessor{}
}

// Process uses Tesseract to extract text from the given image file.
func (t *TesseractsProcessor) Process(filePath string) (string, error) {
	cmd := exec.Command("tesseract", filePath, "stdout")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New("failed to execute Tesseract OCR: " + err.Error())
	}
	return string(output), nil
}
