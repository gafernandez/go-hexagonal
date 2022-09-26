package main

import (
	"os"

	"github.com/gafernandez/go-hexagonal/internal/core/service/asset"
	"github.com/gafernandez/go-hexagonal/internal/infra/repositories"

	"github.com/gafernandez/go-hexagonal/internal/infra/server"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		panic("ENVIRONMENT EMPTY")
	}
	port := os.Getenv("PORT")
	if env == "" {
		panic("PORT EMPTY")
	}
	repository := repositories.NewMemKVS()
	service := asset.NewService(repository)
	server := server.ServerFactory{
		Environment:  env,
		Port:         port,
		AssetService: service,
	}
	server.Start()
}
