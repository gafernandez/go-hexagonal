package asset

import (
	"github.com/gafernandez/go-hexagonal/internal/core/domain"
	"github.com/gafernandez/go-hexagonal/internal/core/ports"
)

type service struct {
	assetRepository ports.DefinanceRepository
}

func NewService(assetRepo ports.DefinanceRepository) *service {
	return &service{
		assetRepository: assetRepo,
	}
}

func (srv *service) Get(symbol string) (domain.Asset, error) {
	asset, err := srv.assetRepository.GetAsset(symbol)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return asset, nil
}

func (srv *service) Update(asset domain.Asset) (domain.Asset, error) {
	asset, err := srv.assetRepository.UpdateAsset(asset)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return asset, nil
}

func (srv *service) Create(asset domain.Asset) (domain.Asset, error) {
	asset, err := srv.assetRepository.CreateAsset(asset)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return asset, nil
}
