package handler

import (
	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/provider"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Callback(c *gin.Context) {
	bot := provider.NewLineBot()
	openAI, err := provider.NewOpenAI()
	if err != nil {
		zap.S().Error(err)
	}
	//
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		zap.S().Warnw("Callback", "bot.ParseRequest", "err", errors.WithStack(err))
		Failed(c, domain.ErrorBadRequest, "")
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				resmassage, err := openAI.Chat(message.Text)
				if err != nil {
					zap.S().Error(err)
					Failed(c, domain.ErrorBadRequest, "")
					return
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(resmassage)).Do(); err != nil {
					zap.S().Warnw("Callback", " bot.ReplyMessage", "err", errors.WithStack(err))

				}
			}
		}
	}
	Success(c, nil)
	return
}
