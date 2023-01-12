package openAI

import (
	"context"

	"github.com/Yoway1994/LineChatGPT3/domain"

	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

var userPrefix string = "user:"
var aiPrefix string = "response:"

func (o openAI) Chat(msg *domain.MessageEvent) (*domain.MessageEvent, error) {
	msg2AI, err := o.GetTextRecord(msg)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	if msg2AI == nil {
		msg.Text = "Ending Chat, 對話結束."
		return msg, nil
	}
	ctx := context.Background()
	// AI原始輸入
	zap.S().Info("AI Prompt: ", msg2AI.Text)
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   1024,
		Temperature: 0.87,
		Prompt:      msg2AI.Text,
	}
	var count int = 0
	var resp gogpt.CompletionResponse
	for {
		resp, err = o.gpt3.CreateCompletion(ctx, req)
		if err == nil {
			break
		} else if err != nil && count == o.index {
			zap.S().Error(err)
			return nil, err
		} else {
			count++
			zap.S().Error(err)
			o.NextClient(count)
		}
	}
	// AI原始輸出
	msg.Text = resp.Choices[0].Text
	zap.S().Info("AI Completion: ", msg.Text)

	// 美化AI字串輸出
	err = o.BeautifyAiOutput(msg)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	// 存入AI response到redis
	err = o.RecordAiResp(msg)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	return msg, nil
}
