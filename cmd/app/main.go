package main

import (
	"os"

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
	server.Start(env, port)
}
