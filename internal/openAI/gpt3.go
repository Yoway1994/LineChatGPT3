package openAI

import (
	"context"

	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/go-redis/redis/v8"

	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

// 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
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
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 1024,
		Prompt:    msg2AI.Text,
	}

	resp, err := o.gpt3.CreateCompletion(ctx, req)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	msg.Text = resp.Choices[0].Text
	return msg, nil
}

func (o openAI) GetTextRecord(msg *domain.MessageEvent) (*domain.MessageEvent, error) {
	if msg.Text == "/end" {
		zap.S().Info("del: ", msg.User)
		err := o.redis.Del(msg.User)
		if err != nil {
			zap.S().Error(err)
			return nil, err
		}
		return nil, nil
	}
	//
	textRecord, err := o.redis.Get(msg.User)
	if err == redis.Nil {
		textRecord = ""
	} else if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	//
	msg.Text = textRecord + msg.Text
	zap.S().Info("輸入redis訊息:", msg.Text)
	zap.S().Info("輸入redis key:", msg.User)
	o.redis.Set(msg.User, msg.Text)
	return msg, nil
}
