package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/viniciusth/tripmaxxer/internal/clients"
	"github.com/viniciusth/tripmaxxer/internal/server"
)


func main() {
    godotenv.Load()

    storage, err := clients.NewStorage()
    if err != nil {
        panic(err)
    }
    defer storage.Close()

    deps := server.ServerDependencies{Storage: storage}
    hs := server.SetupServer(&deps)

    fmt.Println("Server running on port 8080")
    hs.ListenAndServe()
}
