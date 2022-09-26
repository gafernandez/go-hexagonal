package server

import (
	"net/http"

	"github.com/gafernandez/go-hexagonal/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type ServerFactory struct {
	Eenvironment string
	Port         string
	AssetService ports.AssetServices
}

func (s *ServerFactory) Start() {
	handler := NewHandlerHTTP(s.AssetService)
	engine := gin.Default()

	s.registerRoutes(engine, handler)

	if err := engine.Run(":" + s.Port); err != nil {
		panic(err)
	}

}

func (s *ServerFactory) registerRoutes(engine *gin.Engine, handler *HandlerHTTP) {
	//ping
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.POST("/asset", handler.Post)
	engine.GET("/asset/:symbol", handler.Get)
}
