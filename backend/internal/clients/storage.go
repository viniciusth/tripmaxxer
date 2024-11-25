package clients

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage() (*Storage, error) {

    url := fmt.Sprintf("%s&pool_max_conns=10&pool_max_conn_lifetime=5s", os.Getenv("DB_URL"))
    pool, err := pgxpool.New(context.Background(), url)

    _, err = pool.Exec(context.Background(), "SELECT 1;")
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    return &Storage{pool: pool}, nil
}

func (s *Storage) Close() {
    s.pool.Close()
}

