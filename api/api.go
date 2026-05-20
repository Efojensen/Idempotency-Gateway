package api

import (
	"fmt"
	"net/http"
)

type ApiServer struct {
	Port string
	Db *map[string]string
}

func NewApiServer(port string, db *map[string]string) *ApiServer {
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