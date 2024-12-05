package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	expenses "github.com/Rashad-j/image-based-expense-tracker/internal/core/expense"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/config"
	"github.com/sashabaranov/go-openai"
)

type chatgptClient interface {
	Request(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error)
}

type ChatGPTParser struct {
	cfg    *config.Config
	client chatgptClient
}

const prompt = `
	Given the following receipt text, extract and return a JSON object containing items (name, count and price), and include the total.
	Note there could be duplicate items, make sure to count them as well.
	Make sure the prices are per item.
	Also, add a category (in English) to the json object to classify this purchase.
	Receipt text:`

func NewChatGPTParser(cfg *config.Config, client chatgptClient) *ChatGPTParser {
	return &ChatGPTParser{
		cfg:    cfg,
		client: client,
	}
}

func (c *ChatGPTParser) Parse(ctx context.Context, ocrText string) (*expenses.Expenses, error) {
	chatgptText, err := c.askChatGPT(ctx, ocrText)
	if err != nil {
		return nil, fmt.Errorf("failed to ask chatgpt: %w", err)
	}

	jsonText, err := extractJSONFromResponse(chatgptText)
	if err != nil {
		return nil, fmt.Errorf("failed to extract JSON from response: %w", err)
	}

	fmt.Println("json text", jsonText)

	expense, err := parseExpense(jsonText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expense: %w", err)
	}

	return expense, nil
}

// askChatGPT will receive an ocr text and ask chatgpt to get the items and their price
func (c *ChatGPTParser) askChatGPT(ctx context.Context, ocrText string) (string, error) {
	msg := fmt.Sprintf(`
		%s 
		%s
		`, prompt, ocrText)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: msg,
		},
	}

	response, err := c.client.Request(ctx, messages)
	if err != nil {
		return "", fmt.Errorf("failed to ask chatgpt: %w", err)
	}

	return response, nil
}

func extractJSONFromResponse(responseText string) (string, error) {
	re := regexp.MustCompile(`(?s)\{.*\}`) // Matches the first JSON object in the response
	matches := re.FindString(responseText)

	if matches == "" {
		return "", fmt.Errorf("no JSON object found in the response")
	}

	return matches, nil
}

func parseExpense(jsonText string) (*expenses.Expenses, error) {
	var expense expenses.Expenses
	err := json.Unmarshal([]byte(jsonText), &expense)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return &expense, nil
}
