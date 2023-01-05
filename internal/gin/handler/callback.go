package handler

import "github.com/gin-gonic/gin"

func Callback(c *gin.Context) {
	events, err := bot.ParseRequest(c.Request)


}
