package asset

import (
	"github.com/gafernandez/go-hexagonal/internal/core/domain"
	"github.com/gafernandez/go-hexagonal/internal/core/ports"
)

type service struct {
	assetRepository ports.AssetRepository
}

func New(assetRepo ports.AssetRepository) *service {
	return &service{
		assetRepository: assetRepo,
	}
}

func (srv *service) Get(symbol string) (domain.Asset, error) {
	asset, err := srv.assetRepository.Get(symbol)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return asset, nil
}

func (srv *service) Update(asset domain.Asset) (domain.Asset, error) {
	asset, err := srv.assetRepository.Update(asset)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return asset, nil
}

func (srv *service) Create(asset domain.Asset) (domain.Asset, error) {
	asset, err := srv.assetRepository.Save(asset)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return asset, nil
}
