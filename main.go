package main

import (
	"github.com/Efojensen/Idempotency-Gateway/api"
)

func main() {
	apiServer := api.NewApiServer(":8080", nil)

	apiServer.Run()
}