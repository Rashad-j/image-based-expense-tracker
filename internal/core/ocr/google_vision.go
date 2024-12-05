package ocr

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/config"
	"google.golang.org/api/option"
)

// GoogleVisionProcessor implements OCRProcessor using the Google Vision API.
type GoogleVisionProcessor struct {
	cfg *config.Config
}

// NewGoogleVisionProcessor creates a new GoogleVisionProcessor.
func NewGoogleVisionProcessor(cfg *config.Config) *GoogleVisionProcessor {
	return &GoogleVisionProcessor{
		cfg: cfg,
	}
}

// Process uses the Google Vision API to extract text from the given image file.
func (g *GoogleVisionProcessor) Process(file io.Reader) ([]string, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()

	// Create an authenticated Vision API client using the provided credentials file
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile(g.cfg.GoogleApplicationCredentials))
	if err != nil {
		return nil, errors.New("failed to create client: " + err.Error())
	}
	defer client.Close()

	// Read the image from the io.Reader
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return nil, errors.New("failed to read image file: " + err.Error())
	}

	// Perform text detection
	lines, err := client.DetectTexts(ctx, image, nil, 1)
	if err != nil {
		return nil, errors.New("failed to detect text using Google Vision: " + err.Error())
	}

	// Check if any text was detected
	if len(lines) == 0 {
		return nil, errors.New("no text detected in the image")
	}

	result := make([]string, len(lines))
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		result = append(result, line.Description)
		fmt.Println("response ", i, line.Description)
	}

	return result, nil
}
