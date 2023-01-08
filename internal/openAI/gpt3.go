package openAI

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   1024,
		Temperature: 0.75,
		Prompt:      msg2AI.Text,
	}

	resp, err := o.gpt3.CreateCompletion(ctx, req)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	if resp.Choices[0].Text == "" {
		err = errors.New("open AI resp text empty")
		return nil, err
	}
	msg.Text = resp.Choices[0].Text

	return msg, nil
}

func (o openAI) GetTextRecord(msg *domain.MessageEvent) (*domain.MessageEvent, error) {
	if msg.Text == "/end" {
		count := 0
		for count < 4 {
			count++
			key := fmt.Sprintf(msg.User, count)
			err := o.redis.Del(key)
			if err != nil {
				zap.S().Error(err)
				return nil, err
			}
		}
		return nil, nil
	}
	// 取出數字代表最舊的槽位
	recordNumStr, err := o.redis.Get(msg.User)
	if err == redis.Nil {
		recordNumStr = "0"
	} else if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	recordNum, err := strconv.Atoi(recordNumStr)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	zap.S().Info("目前最舊的record槽位: ", recordNum)
	// 設定redis記憶欄位只有4筆
	// pointer會從最舊的記憶欄位, 指向最新的 (0 , 1, 2, 3)
	count := 0
	textRecord := ""
	for count < 4 {
		pointer := (recordNum + count) % 4
		count++
		redisKey := fmt.Sprintf(msg.User, pointer)
		record, err := o.redis.Get(redisKey)
		zap.S().Info("目前槽位: ", pointer)
		zap.S().Info("目前槽位key: ", redisKey)
		if err == redis.Nil {
			zap.S().Info("該槽位無資訊")
			// 找不到跳下一個槽位
			continue
		} else if err != nil {
			zap.S().Error(err)
			return nil, err
		}
		textRecord += (record + " ")
		zap.S().Info("取得該槽位record: ", record)
	}
	// 更新記憶槽位資訊
	// 把最新的資訊覆蓋在舊的資訊的槽位
	newMsgKey := fmt.Sprintf(msg.User, recordNumStr)
	err = o.redis.Set(newMsgKey, msg.Text)
	zap.S().Info("最新槽位key: ", newMsgKey)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	// recordNumStr指向最舊的槽位
	recordNumStr = strconv.Itoa(recordNum + 1)
	err = o.redis.Set(msg.User, recordNumStr)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	//
	msg.Text = textRecord + msg.Text
	zap.S().Info("最終訊息: ", msg.Text)
	return msg, nil
}
