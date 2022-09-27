package definancesources

import (
	"errors"

	"github.com/gafernandez/go-hexagonal/internal/infra/definancesources/kucoin"
)

func GetPrice(symbol string, source string) (float64, error) {
	if source == "kucoin" {
		return kucoin.GetPrice(symbol)
	}
	return -1, errors.New("Invalid source")
}
