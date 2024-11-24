package db

import (
	"context"
	"fmt"
	"javacode-test/internal/env"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPostgresDb(ctx context.Context) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, env.POSTGRES_CONN)
	if err != nil {
		return nil, fmt.Errorf("new postgres pool: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("postgres db ping: %v", err)
	}

	return db, nil
}
