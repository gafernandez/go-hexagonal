package ports

import "github.com/gafernandez/go-hexagonal/internal/core/domain"

type AssetRepository interface {
	Get(symbol string) (domain.Asset, error)
	Save(asset domain.Asset) (domain.Asset, error)
	Update(asset domain.Asset) (domain.Asset, error)
	GetAll() ([]domain.Asset, error)
}
