package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(environment string, port string) *gin.Engine {
	engine := gin.Default()

	registerRoutes(engine)

	engine.Run(":" + port)
	return engine
}

func registerRoutes(engine *gin.Engine) {
	//ping
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//engine.POST("/asset", protocol.WrapperHTTP(protocol.Handler_post_asset))
	//engine.GET("/asset/:symbol", protocol.WrapperHTTP(protocol.Handler_get_asset))
}
