package ports

import "github.com/gafernandez/go-hexagonal/internal/core/domain"

type DefinanceRepository interface {
	GetAsset(symbol string) (domain.Asset, error)
	CreateAsset(asset domain.Asset) (domain.Asset, error)
	UpdateAsset(asset domain.Asset) (domain.Asset, error)
	GetAllAsset() ([]domain.Asset, error)
}
