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
	}
	c.JSON(200, BuildPostResponse(result))
}

func (handler *HandlerHTTP) Get(c *gin.Context) {

}
