package handler

import (
	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Callback(c *gin.Context) {
	line, _ := provider.NewLine()
	openAI, err := provider.NewOpenAI()
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, "")
		return
	}
	//
	msg, err := line.GetMessage(c.Request)
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, "")
		return
	}
	zap.S().Info("input message: ", msg)
	//
	msgAI, err := openAI.Chat(msg)
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, "")
		return
	}
	zap.S().Info("output message: ", msgAI)
	err = line.ReplyMessage(msgAI)
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorServer, "")
		return
	}
	//
	Success(c, nil)
	return
}
