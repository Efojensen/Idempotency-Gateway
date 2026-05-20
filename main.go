package main

import (
	"github.com/Efojensen/Idempotency-Gateway/api"
	"github.com/Efojensen/Idempotency-Gateway/types"
)

func main() {
	store := make(map[string]types.CachedResponse)
	apiServer := api.NewApiServer(":8080", &store)

	apiServer.Run()
}