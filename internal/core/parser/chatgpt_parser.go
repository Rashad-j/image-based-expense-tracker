package parser

import (
	"context"
	"fmt"
	"github.com/Rashad-j/image-based-expense-tracker/config"
	"github.com/sashabaranov/go-openai"
)

type ChatGPTParser struct {
	cfg *config.Config
}

const prompt = `
	You are a receipt parser. Given the following receipt text, extract and return a JSON object containing items (name, count and price), and include the total.
	Note there could be duplicate items, make sure to count them as well.
	Also, add a category to the json object to classify this purchase.
	Receipt text:`

func NewChatGPTParser(cfg *config.Config) *ChatGPTParser {
	return &ChatGPTParser{cfg: cfg}
}

func (c *ChatGPTParser) Parse(ocrText string) (string, error) {
	// TODO: make sure to break this down into separate functions
	// and also add interfaces where needed.
	client := openai.NewClient(c.cfg.ChatGPTApiKey)
	ctx := context.Background()

	prompt := fmt.Sprintf(`
	%s
	%s
	`, prompt, ocrText)

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: c.cfg.ChatGPTModel,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "You are a helpful assistant that parses receipts."},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	})

	if err != nil {
		return "", fmt.Errorf("ChatGPT API call failed: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

// getCleanJSON will receive a chatgpt text with a json object in it
// this function will get the json object out of this generated text and parse it into
// the expenses struct
func getCleanJSON(chatGPTText string) Expenses {
	// TODO: implement
	return Expenses{}
}

// askChatGPT will receive an ocr text and ask chatgpt to get the items and their price
func askChatGPT(ocrText string) (string, error) {
	// TODO: implement
	return "", nil
}
