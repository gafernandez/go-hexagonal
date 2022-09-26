package server

import (
	"github.com/gafernandez/go-hexagonal/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HandlerHTTP struct {
	assetServices ports.AssetServices
}

func NewHandlerHTTP(service ports.AssetServices) *HandlerHTTP {
	return &HandlerHTTP{
		assetServices: service,
	}
}

func (handler *HandlerHTTP) Post(c *gin.Context) {
	body := AssetPostRequest{}
	c.BindJSON(&body)

	asset := body.BuildAsset()
	result, err := handler.assetServices.Create(asset)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, BuildAssetResponse(result))
}

func (handler *HandlerHTTP) Get(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.AbortWithStatusJSON(400, gin.H{"message": "invalid symbol param"})
		return
	}

	asset, err := handler.assetServices.Get(symbol)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildAssetResponse(asset))
}
