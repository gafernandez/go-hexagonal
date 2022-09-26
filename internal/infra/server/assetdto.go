package server

import (
	"time"

	"github.com/gafernandez/go-hexagonal/internal/core/domain"
)

type AssetGetResponse struct {
	Symbol     string    `json:"symbol,omitempty"`
	Source     string    `json:"source,omitempty"`
	Price      float64   `json:"price,omitempty"`
	LastUpdate time.Time `json:"last_update,omitempty"`
	Worker     string    `json:"worker,omitempty"`
}

type AssetPostRequest struct {
	Symbol string `json:"symbol,omitempty"`
	Source string `json:"source,omitempty"`
	Worker string `json:"worker,omitempty"`
}

type AssetPostReponse struct {
	Symbol     string    `json:"symbol,omitempty"`
	Source     string    `json:"source,omitempty"`
	Price      float64   `json:"price,omitempty"`
	LastUpdate time.Time `json:"last_update,omitempty"`
	Worker     string    `json:"worker,omitempty"`
}

func (body AssetPostRequest) BuildAsset() domain.Asset {
	asset := domain.Asset{}
	asset.Symbol = body.Symbol
	asset.Source = body.Source
	asset.Worker = body.Worker
	return asset
}

func BuildPostResponse(asset domain.Asset) interface{} {
	return asset
}
