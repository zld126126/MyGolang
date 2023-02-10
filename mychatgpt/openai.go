package main

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type OpenAIService struct {
	Token string
}

var _OpenAI = &OpenAIService{
	Token: "Your Key",
}

// GetOpenAIService 获取service
func GetOpenAIService() *OpenAIService {
	return _OpenAI
}

// GetToken 获取token
func (p *OpenAIService) GetToken() string {
	return p.Token
}

// GetTextResult 获取文本结果
func (p *OpenAIService) GetTextResult(prompt string) (string, error) {
	c := gogpt.NewClient(p.GetToken())
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
		return "", err
	}
	fmt.Println(resp.Choices[0].Text)
	return resp.Choices[0].Text, nil
}

// GetCodeResult 获取代码结果
func (p *OpenAIService) GetCodeResult(prompt string) (string, error) {
	c := gogpt.NewClient(p.GetToken())
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.CodexCodeDavinci002,
		Temperature:      0,
		MaxTokens:        2048,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		BestOf:           1,
		Prompt:           prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	fmt.Println(resp.Choices[0].Text)
	return resp.Choices[0].Text, nil
}
