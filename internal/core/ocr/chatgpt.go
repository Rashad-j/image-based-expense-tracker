package ocr

import (
	"encoding/base64"
	"fmt"
	"github.com/Rashad-j/image-based-expense-tracker/config"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/chatgpt"
	"io"
)

const prompt = `
	You are a receipt parser. Given the following receipt as a mobile taken image, extract and return a all items names and price and include the total.
	Note there could be duplicate items, make sure to count them as well.
`

type chatgptClient interface {
	SendRequest(payload chatgpt.OpenAIRequest) (chatgpt.OpenAIResponse, error)
}

// ChatgptOCR represents an ChatgptOCR processor using ChatGPT.
type ChatgptOCR struct {
	client chatgptClient
	cfg    *config.Config
}

// NewChatgptOCR creates a new ChatgptOCR processor.
func NewChatgptOCR(cfg *config.Config, client chatgptClient) *ChatgptOCR {
	return &ChatgptOCR{
		cfg:    cfg,
		client: client,
	}
}

// EncodeImageToBase64 reads an image and encodes it as a base64 string.
func EncodeImageToBase64(imagePath io.Reader) (string, error) {
	imageBytes, err := io.ReadAll(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
	}
	return base64.StdEncoding.EncodeToString(imageBytes), nil
}

// ProcessImage processes an image with ChatGPT and returns the result.
func (o *ChatgptOCR) ProcessImage(imagePath io.Reader) (string, error) {
	base64Image, err := EncodeImageToBase64(imagePath)
	if err != nil {
		return "", err
	}

	messages := []chatgpt.Message{
		{
			Role: "user",
			Content: []interface{}{
				chatgpt.TextContent{
					Type: "text",
					Text: prompt,
				},
				chatgpt.ImageContent{
					Type: "image_url",
					ImageURL: chatgpt.ImageURLInfo{
						URL: fmt.Sprintf("data:image/jpeg;base64,%s", base64Image),
					},
				},
			},
		},
	}

	payload := chatgpt.CreateRequestPayload(o.cfg.ChatGPTModel, 200, messages)
	response, err := o.client.SendRequest(payload)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no content found in OpenAI response")
}
