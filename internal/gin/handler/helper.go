package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var errorText string = "伺服器似乎遇到一點問題...Seems like we encountered with some issue here."

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func Failed(c *gin.Context, err domain.ErrorFormat, customMessage string, msg *domain.MessageEvent) {
	line, _ := provider.NewLine()
	message := err.Message
	if customMessage != "" {
		message = customMessage
	}
	errLine := line.SendDevMessage(message)
	if errLine != nil {
		zap.S().Error(errLine)
	}
	if msg != nil && msg.Token != "" {
		msg.Text = errorText
		errLine = line.ReplyMessage(msg)
		if errLine != nil {
			zap.S().Error(errLine)
		}
	}

	switch err {
	case domain.ErrorServer:
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    err.Code,
			"message": message,
			"time":    fmt.Sprintf("%d", time.Now().Unix()),
		})
	case domain.ErrorUnauthorized:
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    err.Code,
			"message": message,
			"time":    fmt.Sprintf("%d", time.Now().Unix()),
		})
	case domain.ErrorForbidden:
		c.JSON(http.StatusForbidden, gin.H{
			"code":    err.Code,
			"message": message,
			"time":    fmt.Sprintf("%d", time.Now().Unix()),
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    err.Code,
			"message": message,
			"time":    fmt.Sprintf("%d", time.Now().Unix()),
		})
	}
}
