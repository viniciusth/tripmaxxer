package main

import (
	"fmt"

	"github.com/viniciusth/tripmaxxer/internal/server"
)


func main() {
    deps := server.ServerDependencies{}
    hs := server.SetupServer(deps)

    fmt.Println("Server running on port 8080")
    hs.ListenAndServe()
}
