package openAI

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/go-redis/redis/v8"

	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

var slot int = 4
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
	zap.S().Info("AI Prompt: ", msg2AI.Text)
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

	// 存入AI response到redis
	msg.Text = resp.Choices[0].Text
	zap.S().Info("AI Completion: ", msg.Text)

	err = o.RecordAiResp(msg)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	// 美化AI字串輸出
	err = o.BeautifyAiOutput(msg)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	//

	return msg, nil
}

func (o openAI) GetTextRecord(msg *domain.MessageEvent) (*domain.MessageEvent, error) {
	if msg.Text == "/end" {
		count := 0
		err := o.redis.Del(msg.User)
		if err != nil {
			zap.S().Error(err)
			return nil, err
		}
		for count < slot {
			count++
			key := fmt.Sprintf("%s%d", msg.User, count)
			err = o.redis.Del(key)
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
	// 設定redis記憶欄位slot筆數
	// pointer會從最舊的記憶欄位, 指向最新的 (0, 1, 2, 3)
	count := 0
	textRecord := ""
	for count < slot {
		pointer := (recordNum + count) % slot
		count++
		redisKey := fmt.Sprintf("%s%d", msg.User, pointer)
		record, err := o.redis.Get(redisKey)
		if err == redis.Nil {
			// 找不到跳下一個槽位
			continue
		} else if err != nil {
			zap.S().Error(err)
			return nil, err
		}
		textRecord += (record + " ")
	}
	// 更新記憶槽位資訊
	// 把最新的資訊覆蓋在舊的資訊的槽位
	newMsgKey := fmt.Sprintf("%s%d", msg.User, recordNum)
	err = o.redis.Set(newMsgKey, msg.Text)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	// recordNumStr指向最舊的槽位
	recordNumStr = strconv.Itoa((recordNum + 1) % slot)
	err = o.redis.Set(msg.User, recordNumStr)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	//
	msg.Text = textRecord + userPrefix + msg.Text
	return msg, nil
}

func (o openAI) RecordAiResp(msg *domain.MessageEvent) error {
	recordNumStr, err := o.redis.Get(msg.User)
	if err == redis.Nil {
		recordNumStr = "0"
	} else if err != nil {
		zap.S().Error(err)
		return err
	}
	recordNum, err := strconv.Atoi(recordNumStr)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	// 取的剛存入的最新訊息
	pointer := (recordNum + 3) % slot
	redisKey := fmt.Sprintf("%s%d", msg.User, pointer)
	record, err := o.redis.Get(redisKey)
	// 存入AI的回覆
	respRecord := record + aiPrefix + msg.Text
	err = o.redis.Set(redisKey, respRecord)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	return nil
}

func (o openAI) BeautifyAiOutput(msg *domain.MessageEvent) error {
	originText := msg.Text
	// 先檢查prefix
	pattern1 := "(?s)" + aiPrefix + "(.+)"
	r1, err := regexp.Compile(pattern1)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	// 有時候會有大寫Response:
	pattern2 := "(?s)Response:(.+)"
	r2, err := regexp.Compile(pattern2)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	// 檢查雙空格
	pattern3 := "\n\n" + "(.+)"
	r3, err := regexp.Compile(pattern3)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	var beautifiedText string
	if r1.MatchString(originText) {
		// 發現prefix就移除
		rawText := r1.FindString(originText)
		beautifiedText = strings.TrimPrefix(rawText, aiPrefix)
	} else if r2.MatchString(originText) {
		rawText := r2.FindString(originText)
		beautifiedText = strings.TrimPrefix(rawText, "Response:")
	} else if r3.MatchString(originText) {
		// 沒發現prefix, 但發現雙空格
		rawText := r3.FindString(originText)
		beautifiedText = strings.TrimPrefix(rawText, "\n\n")
	} else {
		// 都沒發現, 不做美化
		beautifiedText = msg.Text
	}
	// 再去一次雙空格
	beautifiedText = strings.TrimPrefix(beautifiedText, "\n\n")
	// 輸出美化
	if msg.Text == "" {
		msg.Text = ":)"
	} else {
		msg.Text = beautifiedText
	}
	return nil
}
