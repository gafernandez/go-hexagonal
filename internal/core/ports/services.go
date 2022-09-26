package ports

import "github.com/gafernandez/go-hexagonal/internal/core/domain"

type AssetServices interface {
	Get(symbol string) (domain.Asset, error)
	GetAll() ([]domain.Asset, error)
	Create(asset domain.Asset) (domain.Asset, error)
	Update(asset domain.Asset) (domain.Asset, error)
	Refresh(symbol string, source DefinanceSourceServices) (domain.Asset, error)
}

type DefinanceSourceServices interface {
	FillAsset(symbol string) domain.Asset
}
