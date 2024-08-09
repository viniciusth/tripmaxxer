package clients

import (
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx"
)

type Storage struct {
	pool *pgx.ConnPool
}

func NewStorage() (*Storage, error) {
    cfg, err := pgx.ParseURI(os.Getenv("DB_URL"))
    if err != nil {
        return nil, fmt.Errorf("failed to parse DB_URL: %w", err)
    }

    pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
    	ConnConfig: cfg,
    	MaxConnections: 5,
    	AcquireTimeout: 5 * time.Second,
    })

    _, err = pool.Exec("SELECT 1;")
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    return &Storage{pool: pool}, nil
}

func (s *Storage) Close() {
    s.pool.Close()
}

