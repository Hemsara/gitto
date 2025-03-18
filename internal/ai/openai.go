package ai

import (
	"context"
	"fmt"

	"github.com/Hemsara/gitto/internal/keys"
	"github.com/sashabaranov/go-openai"
)

func GenerateCommitMessage(diff string) (string, error) {
	apikey, err := keys.LoadAPIKey()
	if err != nil {
		return "", err
	}
	client := openai.NewClient(apikey)

	prompt := fmt.Sprintf("🔎 Generate a concise and meaningful git commit message (with emojis) for the following diff:\n\n%s\n\n➡️ Use relevant emojis for context (e.g., 🐛 for fixes, ✨ for features) and maintain a consistent style.", diff)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "You are a helpful assistant that writes commit messages for git based on the provided diff.",
				},
				{
					Role:    "user",
					Content: prompt,
				},
			},
			MaxTokens: 50,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
