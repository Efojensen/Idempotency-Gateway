package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

type ApiServer struct {
	Port string
	Db *sql.DB
}

func NewApiServer(port string, db *sql.DB) *ApiServer {
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