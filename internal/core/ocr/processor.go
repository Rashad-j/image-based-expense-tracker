// Package ocr is responsible for taking images and extracting the text using an ocr provider.
// uses the Strategy Pattern to allow the application to support multiple OCR
// engines (e.g., Tesseract, Google Vision API, etc) with interchangeable implementations.
package ocr

import "io"

// Processor defines the interface for OCR processing strategies.
type Processor interface {
	// Process takes the file as bytes of an image and extracts text from it.
	Process(file io.Reader) (string, error)
}
