package line

import (
	"net/http"

	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.uber.org/zap"
)

func (l line) GetMessage(r *http.Request) (*domain.MessageEvent, error) {
	events, err := l.BotProd.ParseRequest(r)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	var msg domain.MessageEvent
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg = domain.MessageEvent{
					User:  event.Source.UserID,
					Text:  message.Text,
					Token: event.ReplyToken,
				}
				break
			}
		}
	}
	return &msg, nil
}

func (l line) ReplyMessage(reply *domain.MessageEvent) error {
	_, err := l.BotProd.ReplyMessage(
		reply.Token,
		linebot.NewTextMessage(reply.Text)).Do()
	if err != nil {
		zap.S().Error(err)
		return err
	}
	return nil
}

func (l line) SendDevMessage(msg string) error {
	_, err := l.BotDev.PushMessage(
		"Ufdfa8423b5962e5d05b5e89b61dc43d4",
		linebot.NewTextMessage(msg),
	).Do()
	if err != nil {
		zap.S().Error(err)
		return err
	}
	return nil
}
