package server

import (
	"net/http"

	"github.com/gafernandez/go-hexagonal/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type ServerFactory struct {
	Environment  string
	Port         string
	AssetService ports.AssetServices
}

func (s *ServerFactory) Start() {
	assetHandler := NewAssetHandlerHTTP(s.AssetService)
	engine := gin.Default()

	s.registerRoutes(engine, assetHandler)

	if err := engine.Run(":" + s.Port); err != nil {
		panic(err)
	}

}

func (s *ServerFactory) registerRoutes(engine *gin.Engine, aHandler *AssetHandlerHTTP) {
	//ping
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.POST("/asset", aHandler.Post)
	engine.GET("/asset/:symbol", aHandler.Get)
	engine.POST("/asset/:symbol/refresh", aHandler.Refresh)
}
