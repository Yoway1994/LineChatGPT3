//go:build wireinject
// +build wireinject

package provider

import (
	"log"
	"sync"

	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/line"
	"github.com/Yoway1994/LineChatGPT3/internal/openAI"
	"github.com/google/wire"
	"github.com/line/line-bot-sdk-go/linebot"
	gogpt "github.com/sashabaranov/go-gpt3"
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

func NewLineBot() *linebot.Client {
	botOnce.Do(func() {
		log.Println("init line bot")
		bot = line.NewBot(cg)
		log.Println("init line bot success")
	})
	return bot
}

var gpt *gogpt.Client
var gptOnce sync.Once

func NewChatGpt3(cg *config.Config) *gogpt.Client {
	gptOnce.Do(func() {
		log.Println("init gpt3")
		gpt = openAI.NewClient(cg)
		log.Println("init gpt3")
	})
	return gpt
}

func NewOpenAI() (domain.OpenAI, error) {
	panic(wire.Build(openAI.NewOpenAI, NewChatGpt3, NewConfig))
}
