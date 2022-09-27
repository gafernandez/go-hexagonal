package asset

import (
	"time"

	"github.com/gafernandez/go-hexagonal/internal/infra/definancesources"

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
	asset.LastUpdate = time.Now()
	asset, err := srv.assetRepository.UpdateAsset(asset)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return asset, nil
}

func (srv *service) Create(asset domain.Asset) (domain.Asset, error) {
	asset.LastUpdate = time.Now()
	createdAsset, err := srv.assetRepository.CreateAsset(asset)
	if err != nil {
		//TODO: Error handlering
		return domain.Asset{}, err
	}
	return createdAsset, nil
}

func (srv *service) GetAll() ([]domain.Asset, error) {
	assets, err := srv.assetRepository.GetAllAsset()
	if err != nil {
		//TODO: error handlering
		return []domain.Asset{}, err
	}
	return assets, nil
}

func (srv *service) Refresh(symbol string) (domain.Asset, error) {
	assetRepo, err := srv.assetRepository.GetAsset(symbol)
	if err != nil {
		return domain.Asset{}, err
	}

	price, err := definancesources.GetPrice(symbol, assetRepo.Source)
	if err != nil {
		return domain.Asset{}, err
	}

	assetRepo.Price = price
	assetUpdated, err := srv.Update(assetRepo)
	if err != nil {
		return domain.Asset{}, err
	}

	return assetUpdated, nil
}
