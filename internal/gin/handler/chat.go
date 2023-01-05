package handler

import (
	"bytes"
	"io/ioutil"

	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/Yoway1994/LineChatGPT3/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type request struct {
	Text string `json:"text" binding:"required"`
}

type respond struct {
	Msg string `json:"message"`
}

func Chat(c *gin.Context) {
	//
	openAI, err := provider.NewOpenAI()
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorBadRequest, "")
		return
	}
	//
	var req request
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.S().Infow("", "err", err, "reqBody", string(reqBody))
		Failed(c, domain.ErrorBadRequest, "")
		return
	}
	//
	resmassage, err := openAI.Chat(req.Text)
	if err != nil {
		zap.S().Error(err)
		Failed(c, domain.ErrorBadRequest, "")
		return
	}
	//
	Success(c, &respond{
		Msg: resmassage,
	})
}
