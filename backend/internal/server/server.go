package server

import (
	"net/http"

	"github.com/viniciusth/tripmaxxer/internal/clients"
)

type ServerDependencies struct {
    Storage *clients.Storage
}

func SetupServer(deps *ServerDependencies) *http.Server {
    http.HandleFunc("/health", deps.health)
	return &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
}

func (s *ServerDependencies) health(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}
