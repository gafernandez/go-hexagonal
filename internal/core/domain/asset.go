package domain

import "time"

type Asset struct {
	Symbol     string    `json:"symbol,omitempty"`
	Source     string    `json:"source,omitempty"`
	Price      float64   `json:"price,omitempty"`
	LastUpdate time.Time `json:"last_update,omitempty"`
	Worker     string    `json:"worker,omitempty"`
}
