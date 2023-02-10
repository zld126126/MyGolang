package main

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	token := "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"
	ask := "Please write a article About how to live in the USA?"
	example(token, ask)
}

func example(token string, prompt string) {
	c := gogpt.NewClient(token)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci001,
		Temperature:      0.4,
		MaxTokens:        1000,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		BestOf:           1,
		Prompt:           prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
