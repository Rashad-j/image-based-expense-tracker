package parser

import (
	"fmt"
	"github.com/Rashad-j/image-based-expense-tracker/config"
	expenses "github.com/Rashad-j/image-based-expense-tracker/internal/core/expense"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/chatgpt"
)

type chatgptClient interface {
	SendRequest(payload chatgpt.OpenAIRequest) (chatgpt.OpenAIResponse, error)
}

type ChatGPTParser struct {
	cfg    *config.Config
	client chatgptClient
}

const prompt = `
	Given the following receipt text, extract and return a JSON object containing items (name, count and price), and include the total.
	Note there could be duplicate items, make sure to count them as well.
	Also, add a category to the json object to classify this purchase.
	Receipt text:`

func NewChatGPTParser(cfg *config.Config, client chatgptClient) *ChatGPTParser {
	return &ChatGPTParser{
		cfg:    cfg,
		client: client,
	}
}

func (c *ChatGPTParser) Parse(ocrText string) (string, error) {
	fmt.Printf("parsing text", ocrText)
	return c.askChatGPT(ocrText)
}

// getCleanJSON will receive a chatgpt text with a json object in it
// this function will get the json object out of this generated text and parse it into
// the expenses struct
func getCleanJSON(chatGPTText string) expenses.Expenses {
	// TODO: implement
	return expenses.Expenses{}
}

// askChatGPT will receive an ocr text and ask chatgpt to get the items and their price
func (c *ChatGPTParser) askChatGPT(ocrText string) (string, error) {
	text := fmt.Sprintf(`
		%s 
		%s
		`, prompt, ocrText)
	messages := []chatgpt.Message{
		{
			Role: "user",
			Content: []interface{}{
				chatgpt.TextContent{
					Type: "text",
					Text: text,
				},
			},
		},
	}

	payload := chatgpt.CreateRequestPayload(c.cfg.ChatGPTModel, 200, messages)
	response, err := c.client.SendRequest(payload)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no content found in OpenAI response")
}
