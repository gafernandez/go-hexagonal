package repositories

import (
	"encoding/json"
	"errors"

	"github.com/gafernandez/go-hexagonal/internal/core/domain"
)

type memkvs struct {
	kvsAsset map[string][]byte
}

func NewDefinanceMemRepository() *memkvs {
	return &memkvs{kvsAsset: map[string][]byte{}}
}

func (repo *memkvs) GetAsset(symbol string) (domain.Asset, error) {
	if value, result := repo.kvsAsset[symbol]; result {
		asset := domain.Asset{}
		err := json.Unmarshal(value, &asset)
		if err != nil {
			return domain.Asset{}, errors.New("Fail to get value from kvs")
		}

		return asset, nil
	}
	return domain.Asset{}, errors.New("Asset not found in kvs")
}

func (repo *memkvs) GetAllAsset() ([]domain.Asset, error) {
	assets := make([]domain.Asset, len(repo.kvsAsset))

	i := 0
	for _, asset := range repo.kvsAsset {
		json.Unmarshal(asset, &assets[i])
		i++
	}

	return assets, nil
}

func (repo *memkvs) CreateAsset(asset domain.Asset) (domain.Asset, error) {

	if repo.kvsAsset[asset.Symbol] != nil {
		return domain.Asset{}, errors.New("Duplicate key " + asset.Symbol)
	}

	bytes, err := json.Marshal(asset)
	if err != nil {
		return domain.Asset{}, errors.New("Asset fails to marshal into json string")
	}
	repo.kvsAsset[asset.Symbol] = bytes
	return asset, nil
}

func (repo *memkvs) UpdateAsset(asset domain.Asset) (domain.Asset, error) {
	bytes, err := json.Marshal(asset)
	if err != nil {
		return domain.Asset{}, errors.New("Asset fails to marshal into json string")
	}
	repo.kvsAsset[asset.Symbol] = bytes
	return asset, nil
}
