package clients

import (
	"os"
	"time"

	"github.com/jackc/pgx"
)

type Storage struct {
	client *pgx.ConnPool
}

func NewStorage() (*Storage, error) {
    cfg, err := pgx.ParseURI(os.Getenv("DB_URL"))
    if err != nil {
        return nil, err
    }

    client, err := pgx.NewConnPool(pgx.ConnPoolConfig{
    	ConnConfig: cfg,
    	MaxConnections: 5,
    	AcquireTimeout: 5 * time.Second,
    })

    return &Storage{client: client}, nil
}

func (s *Storage) Close() {
    s.client.Close()
}

