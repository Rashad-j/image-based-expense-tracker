// package ocr
package ocr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Rashad-j/image-based-expense-tracker/config"
	"io"
	"net/http"
)

const (
	openAIEndpoint = "https://api.openai.com/v1/chat/completions"
	defaultModel   = "gpt-4o-mini"
)

type chatGPT struct {
	cfg *config.Config
}

// httpClient is an interface to abstract HTTP client for testability.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// openAIRequest represents the payload sent to the OpenAI API.
type openAIRequest struct {
	Model     string    `json:"model"`
	Messages  []message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

// message represents a chat message for OpenAI.
type message struct {
	Role    string        `json:"role"`
	Content []interface{} `json:"content"`
}

// textContent represents text content in OpenAI messages.
type textContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// imageContent represents image content in OpenAI messages.
type imageContent struct {
	Type     string       `json:"type"`
	ImageURL imageURLInfo `json:"image_url"`
}

// imageURLInfo represents the URL of the image.
type imageURLInfo struct {
	URL string `json:"url"`
}

// openAIResponse represents the OpenAI API response.
type openAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewChatGPT(cfg *config.Config) *chatGPT {
	return &chatGPT{cfg: cfg}
}

// encodeImageToBase64 reads a file from disk and encodes it as a base64 string.
func encodeImageToBase64(imagePath io.Reader) (string, error) {
	imageBytes, err := io.ReadAll(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
	}
	return base64.StdEncoding.EncodeToString(imageBytes), nil
}

// createOpenAIRequestPayload constructs the OpenAI request payload.
func createOpenAIRequestPayload(base64Image string) (openAIRequest, error) {
	return openAIRequest{
		Model: defaultModel,
		Messages: []message{
			{
				Role: "user",
				Content: []interface{}{
					textContent{
						Type: "text",
						Text: "What’s in this receipt? Extract items and their prices. Return a json object with items (name and price) and total",
					},
					imageContent{
						Type: "image_url",
						ImageURL: imageURLInfo{
							URL: fmt.Sprintf("data:image/jpeg;base64,%s", base64Image),
						},
					},
				},
			},
		},
		MaxTokens: 200,
	}, nil
}

// sendOpenAIRequest sends a request to OpenAI API using a given HTTP client.
func sendOpenAIRequest(client httpClient, apiKey string, payload openAIRequest) (openAIResponse, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Avoiding potential ReadAll errors in logs.
		return openAIResponse{}, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var r openAIResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return openAIResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return r, nil
}

// Process implements the processor interface. Sends a request to chatgpt and returns the result as a string
func (c *chatGPT) Process(imagePath io.Reader) (string, error) {
	base64Image, err := encodeImageToBase64(imagePath)
	if err != nil {
		return "", err
	}

	payload, err := createOpenAIRequestPayload(base64Image)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	response, err := sendOpenAIRequest(client, c.cfg.ChatGPTApiKey, payload)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no content found in OpenAI response")
}
