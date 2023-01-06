package openAI

import (
	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/redis"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type openAI struct {
	gpt3  *gogpt.Client
	redis *redis.GoRedis
}

func NewOpenAI(c *gogpt.Client, r *redis.GoRedis) domain.OpenAI {
	return openAI{
		gpt3:  c,
		redis: r,
	}
}

func NewClient(cg *config.Config) *gogpt.Client {
	c := gogpt.NewClient(cg.Gpt3.ApiKey)
	return c
}
