package chatgpt

import (
	"context"
	"fmt"

	"github.com/Rashad-j/image-based-expense-tracker/pkg/config"
	openai "github.com/sashabaranov/go-openai"
)

// ClientUsingLib represents a ChatGPT client.
type ClientUsingLib struct {
	apiClient *openai.Client
	cfg       *config.Config
}

// NewClientUsingOpenapiLib creates a new ChatGPT client using the go-openai library.
func NewClientUsingOpenapiLib(cfg *config.Config) *ClientUsingLib {
	return &ClientUsingLib{
		apiClient: openai.NewClient(cfg.ChatGPTApiKey),
		cfg:       cfg,
	}
}

func (c *ClientUsingLib) Request(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	response, err := c.apiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    c.cfg.ChatGPTModel,
		Messages: messages,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no content in response")
}
