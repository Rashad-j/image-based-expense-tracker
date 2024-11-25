// Package chatgpt provides a generic interface for sending requests to the OpenAI API.
package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Rashad-j/image-based-expense-tracker/config"
	"io"

	"net/http"
)

// Client represents a ChatGPT client.
type Client struct {
	cfg    *config.Config
	client httpClient
}

// httpClient abstracts the HTTP client for testability.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// OpenAIRequest represents the payload sent to the OpenAI API.
type OpenAIRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

// Message represents a chat message for OpenAI.
type Message struct {
	Role    string        `json:"role"`
	Content []interface{} `json:"content"`
}

// TextContent represents text content in OpenAI messages.
type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ImageContent represents image content in OpenAI messages.
type ImageContent struct {
	Type     string       `json:"type"`
	ImageURL ImageURLInfo `json:"image_url"`
}

// ImageURLInfo represents the URL of the image.
type ImageURLInfo struct {
	URL string `json:"url"`
}

// OpenAIResponse represents the OpenAI API response.
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// NewClient creates a new ChatGPT client.
func NewClient(cfg *config.Config, httpClient httpClient) *Client {
	return &Client{
		cfg:    cfg,
		client: httpClient,
	}
}

// CreateRequestPayload constructs a generic OpenAI request payload.
func CreateRequestPayload(model string, maxTokens int, messages []Message) OpenAIRequest {
	return OpenAIRequest{
		Model:     model,
		Messages:  messages,
		MaxTokens: maxTokens,
	}
}

// SendRequest sends a request to the OpenAI API and returns the response.
func (c *Client) SendRequest(payload OpenAIRequest) (OpenAIResponse, error) {
	const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.cfg.ChatGPTApiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return OpenAIResponse{}, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var r OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return r, nil
}
