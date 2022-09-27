package kucoin

import (
	"errors"
	"strconv"

	"github.com/Kucoin/kucoin-go-sdk"
)

func GetPrice(symbol string) (float64, error) {
	ticker, err := getTicker(symbol)
	if err != nil {
		return -1, err
	}
	price, err := strconv.ParseFloat(ticker.Price, 64)
	if err != nil {
		return -1, errors.New("Fail to parse price.")
	}
	return price, nil
}

func getTicker(symbol string) (kucoin.TickerLevel1Model, error) {
	// API key version 2.0
	s := kucoin.NewApiService()
	resultJson := &kucoin.TickerLevel1Model{}

	result, err := s.TickerLevel1(symbol)

	if err != nil {
		return *resultJson, errors.New("Fail to get price from kucoin.")
	}

	if err2 := result.ReadData(resultJson); err2 != nil {
		return *resultJson, errors.New("Fail to get price from kucoin.")
	}

	return *resultJson, nil
}
