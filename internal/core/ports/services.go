package ports

import "github.com/gafernandez/go-hexagonal/internal/core/domain"

type AssetServices interface {
	Get(symbol string) (domain.Asset, error)
	GetAll() ([]domain.Asset, error)
	Create(asset domain.Asset) (domain.Asset, error)
	Update(asset domain.Asset) (domain.Asset, error)
	Refresh(symbol string) (domain.Asset, error)
}

type DefinanceSourcesServices interface {
	GetPrice(symbol string, source string) (float64, error)
}
