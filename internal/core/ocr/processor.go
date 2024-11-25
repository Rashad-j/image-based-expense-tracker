// Package ocr_test is responsible for taking images and extracting the text using an ocr_test provider.
// uses the Strategy Pattern to allow the application to support multiple ChatgptOCR
// engines (e.g., Tesseract, Google Vision API, etc) with interchangeable implementations.
package ocr

import "io"

// Processor defines the interface for ChatgptOCR processing strategies.
type Processor interface {
	// Process takes the file as bytes of an image and extracts text from it.
	Process(file io.Reader) (string, error)
}
