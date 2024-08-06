package main

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()

    cfg, err := pgx.ParseURI(os.Getenv("DB_URL"))
    if err != nil {
        panic(err)
    }

    fmt.Println("Connecting to database")
    conn, err := pgx.Connect(cfg)
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    fmt.Println("Connected to database")


    _, err = conn.Exec("CREATE TABLE IF NOT EXISTS migrations (name text PRIMARY KEY);")
    if err != nil {
        panic(err)
    }

    files, err := os.ReadDir("./migrations")
    if err != nil {
        panic(err)
    }

    migrations := make(map[string]bool)
    rows, err := conn.Query("SELECT name FROM migrations;")
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err != nil {
            panic(err)
        }
        migrations[name] = true
    }

    for _, file := range files {
        if migrations[file.Name()] {
            fmt.Println("Skipping already applied migration", file.Name())
            continue
        }
        fmt.Println("Running migration", file.Name())

        f, err := os.ReadFile("./migrations/" + file.Name())
        if err != nil {
            panic(err)
        }
        _, err = conn.Exec(string(f))
        if err != nil {
            panic(err)
        }
        _, err = conn.Exec("INSERT INTO migrations (name) VALUES ($1);", file.Name())
        if err != nil {
            panic(err)
        }
    }
}
