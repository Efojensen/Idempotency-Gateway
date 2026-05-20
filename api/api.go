package api

import (
	"fmt"
	"net/http"

	"github.com/Efojensen/Idempotency-Gateway/types"
)

type ApiServer struct {
	Port string
	Db *map[string]types.CachedResponse
}

func NewApiServer(port string, db *map[string]types.CachedResponse) *ApiServer {
	return &ApiServer{
		Port: port,
		Db: db,
	}
}

func (server *ApiServer) Run() error {
	mux := http.NewServeMux()

	fmt.Println("Server running on localhost" + server.Port)
	return http.ListenAndServe(server.Port, mux)
}