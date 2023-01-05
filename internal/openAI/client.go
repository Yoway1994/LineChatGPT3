package openAI

import (
	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/Yoway1994/LineChatGPT3/domain"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type openAI struct {
	gpt3 *gogpt.Client
}

func NewOpenAI(c *gogpt.Client) domain.OpenAI {
	return openAI{
		gpt3: c,
	}
}

func NewClient(cg *config.Config) *gogpt.Client {
	c := gogpt.NewClient(cg.Gpt3.ApiKey)
	return c
}
