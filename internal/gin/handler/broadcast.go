package handler

import (
	"github.com/Yoway1994/LineChatGPT3/internal/provider"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Broadcast(c *gin.Context) {
	bot := provider.NewLineBot()
	//
	bot.BroadcastMessage(linebot.NewTextMessage("廣播訊息測試")).Do()
}
