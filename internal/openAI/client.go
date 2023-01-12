package openAI

import (
	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/redis"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type openAI struct {
	gpt3        *gogpt.Client
	gpt3Clients *[]*gogpt.Client
	redis       *redis.GoRedis
	num         int
}

func NewOpenAI(c *gogpt.Client, g *[]*gogpt.Client, r *redis.GoRedis, n int) domain.OpenAI {
	return openAI{
		gpt3:        c,
		gpt3Clients: g,
		redis:       r,
		num:         n,
	}
}

func NewClient(apikey string) *gogpt.Client {
	c := gogpt.NewClient(apikey)
	return c
}
