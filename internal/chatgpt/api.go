package chatgpt

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func chat() {
	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model: gogpt.GPT3TextDavinci003,
		MaxTokens: 1024,
		Prompt:    "現在假裝你是一隻貓"+"叫兩聲",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}