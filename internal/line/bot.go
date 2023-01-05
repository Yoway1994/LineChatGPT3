package line

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/Yoway1994/LineChatGpt3/config"
)

type line struct {
	Bot *linebot.Client
}

func NewBot(cg config.Config) *linebot.Client {
	bot, err := linebot.New()
	if err != nil {
		panic("line bot fail")
	}
	return bot
}
