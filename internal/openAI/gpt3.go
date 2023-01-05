package openAI

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

// 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
func (o openAI) Chat(text string) (string, error) {
	ctx := context.Background()
	zap.S().Info("input text: ", text)
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 1024,
		Prompt:    text,
	}
	resp, err := o.gpt3.CreateCompletion(ctx, req)
	if err != nil {
		zap.S().Error(err)
		return "openAI api 似乎出了點問題, 請稍後再試", err
	}
	return resp.Choices[0].Text, nil
}
