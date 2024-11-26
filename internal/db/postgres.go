package db

import (
	"context"
	"fmt"
	"javacode-test/internal/env"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPostgresDb(ctx context.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(env.POSTGRES_CONN)
	if err != nil {
		return nil, fmt.Errorf("parse config: %v", err)
	}

	config.MaxConns = 40
	config.MaxConnIdleTime = time.Minute * 5
	config.MaxConnLifetime = time.Minute * 10

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("new postgres pool: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("postgres db ping: %v", err)
	}

	return db, nil
}
