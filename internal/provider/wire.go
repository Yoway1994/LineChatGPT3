package provider

import (
	"log"
	"sync"

	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/Yoway1994/LineChatGPT3/internal/line"
	"github.com/line/line-bot-sdk-go/linebot"
)

var cg *config.Config
var configOnce sync.Once

func NewConfig() *config.Config {
	configOnce.Do(func() {
		log.Println("read config")
		cg = config.NewConfig()
		log.Println("read config success")
	})
	return cg
}

var bot *linebot.Client
var botOnce sync.Once

func NewLineBot() *linebot.Clinet {
	botOnce.Do(func() {
		log.Println("init line bot")
		bot = line.NewBot(cg.)
		log.Println("init line bot success")
	})
	return bot
}
