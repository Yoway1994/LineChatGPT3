package handler

import (
	"fmt"

	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/provider"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Callback(c *gin.Context) {
	line, _ := provider.NewLine()
	//
	msg, err := line.GetMessage(c.Request)
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, fmt.Sprintf("line.GetMessage: %s", err), msg)
		return
	}
	//
	openAI, err := provider.NewOpenAI()
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, fmt.Sprintf("provider.NewOpenAI: %s", err), msg)
		return
	}
	//
	msgAI, err := openAI.Chat(msg)
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, fmt.Sprintf("openAI.Chat: %s", err), msgAI)
		return
	}
	//
	err = line.ReplyMessage(msgAI)
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, fmt.Sprintf("line.ReplyMessage: %s", err), msgAI)
		return
	}
	//
	Success(c, nil)
}
