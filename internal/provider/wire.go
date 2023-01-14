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
	"github.com/Yoway1994/LineChatGPT3/internal/redis"

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

var botProd *linebot.Client
var botDev *linebot.Client
var botOnce sync.Once

func NewLineBot(cfg *config.Config) (*linebot.Client, *linebot.Client) {
	botOnce.Do(func() {
		log.Println("init line bot")
		botProd, botDev = line.NewBot(cfg)
		log.Println("init line bot success")
	})
	return botProd, botDev
}

var gpt *gogpt.Client
var gptOnce sync.Once

func NewChatGpt3(cfg *config.Config) *gogpt.Client {
	gptOnce.Do(func() {
		log.Println("init gpt3")
		apikey := cfg.Gpt3.ApiKey
		gpt = openAI.NewClient(apikey)
		log.Println("init gpt3 success")
	})
	return gpt
}

var gpt3 *[]*gogpt.Client
var gpt3Once sync.Once

func NewChatGpt3List(cfg *config.Config) *[]*gogpt.Client {
	gpt3Once.Do(func() {
		gpt3New := make([]*gogpt.Client, 0)
		log.Println("init gpt3 list")
		for _, key := range cfg.Gpt3.ApiKeys {
			gpti := openAI.NewClient(key)
			gpt3New = append(gpt3New, gpti)
		}
		log.Println("init gpt3 list")
		gpt3 = &gpt3New
	})
	return gpt3
}

func NewChatGpt3Length(gpt3 *[]*gogpt.Client) int {
	return len(*gpt3)
}

func NewLine() (domain.Line, error) {
	cfg := NewConfig()
	botp, botd := NewLineBot(cfg)
	line := line.NewLine(botp, botd)
	if line == nil {
		panic(line)
	}
	return line, nil
}

//
func NewOpenAI() (domain.OpenAI, error) {
	panic(wire.Build(openAI.NewOpenAI, redis.NewGoRedis, NewConfig, NewChatGpt3, NewChatGpt3List, NewChatGpt3Length))
}
