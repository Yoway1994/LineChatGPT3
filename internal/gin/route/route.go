package route

import (
	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/Yoway1994/LineChatGPT3/internal/gin/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config *config.Config) *gin.Engine {
	if config.Gin.Mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	// r.Use(middleware.CORSMiddleware())

	r.Static("/api/assets", "./assets")
	r.POST("/callback", handler.Callback)
	return r
}
