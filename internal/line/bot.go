package line

import (
	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/line/line-bot-sdk-go/linebot"
)

func NewBot(cg *config.Config) *linebot.Client {
	bot, err := linebot.New(cg.Line.Secret, cg.Line.Token)
	if err != nil {
		panic("new linebot fail")
	}
	return bot
}
