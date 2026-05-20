package main

import (
	"github.com/Efojensen/Idempotency-Gateway/api"
)

func main() {
	store := make(map[string]string)
	apiServer := api.NewApiServer(":8080", &store)

	apiServer.Run()
}