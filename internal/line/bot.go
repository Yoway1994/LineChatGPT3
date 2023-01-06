package line

import (
	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/line/line-bot-sdk-go/linebot"
)

type line struct {
	BotProd *linebot.Client
	BotDev  *linebot.Client
}

func NewLine(botp, botd *linebot.Client) domain.Line {
	return line{
		BotProd: botp,
		BotDev:  botd,
	}
}

func NewBot(cg *config.Config) (*linebot.Client, *linebot.Client) {
	botProd, err := linebot.New(cg.Line.Secret, cg.Line.Token)
	if err != nil {
		panic("new linebot fail")
	}
	botDev, err := linebot.New(cg.Line.DevS, cg.Line.DevT)
	if err != nil {
		panic("new linebot fail")
	}
	return botProd, botDev
}
