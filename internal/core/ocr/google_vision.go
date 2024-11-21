package ocr

import (
	"context"
	"errors"
	"io"

	vision "cloud.google.com/go/vision/apiv1"
)

// GoogleVisionProcessor implements OCRProcessor using the Google Vision API.
type GoogleVisionProcessor struct {
	client *vision.ImageAnnotatorClient
}

// NewGoogleVisionProcessor creates a new GoogleVisionProcessor.
func NewGoogleVisionProcessor() (*GoogleVisionProcessor, error) {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, errors.New("failed to create Google Vision client: " + err.Error())
	}
	return &GoogleVisionProcessor{client: client}, nil
}

// Process uses the Google Vision API to extract text from the given image file.
func (g *GoogleVisionProcessor) Process(file io.Reader) (string, error) {
	ctx := context.Background()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return "", errors.New("failed to open image file: " + err.Error())
	}

	response, err := g.client.DetectTexts(ctx, image, nil, 1)
	if err != nil {
		return "", errors.New("failed to detect text using Google Vision: " + err.Error())
	}

	if len(response) == 0 {
		return "", errors.New("no text detected in the image")
	}

	return response[0].Description, nil
}
