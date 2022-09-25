package ports

import "github.com/gafernandez/go-hexagonal/internal/core/domain"

type AssetServices interface {
	Get(symbol string) (domain.Asset, error)
	Create(asset domain.Asset) (domain.Asset, error)
	Update(asset domain.Asset) (domain.Asset, error)
}

type SourceServices interface {
	FillAsset(symbol string) domain.Asset
}

type WorkerServices interface {
	Start() error
	Stop() error
}
