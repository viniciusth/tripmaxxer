package server

import "net/http"

type ServerDependencies struct {
}

func SetupServer(deps ServerDependencies) *http.Server {
    http.HandleFunc("/", helloWorld)
	return &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}
