package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Ping(c *gin.Context) {
	zap.S().Info("Ping")
}
