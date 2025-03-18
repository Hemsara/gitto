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

	prompt := fmt.Sprintf("Generate a professional and concise Git commit message for the following diff:\n\n%s\n\n➡️ Use conventional commit prefixes (e.g., feat, fix, chore) and relevant emojis (e.g., 🐛 for fixes, ✨ for features). Keep it brief and consistent.", diff)

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
