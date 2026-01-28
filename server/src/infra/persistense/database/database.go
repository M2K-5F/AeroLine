package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnectionPoolFromEnv(ctx context.Context) (*pgxpool.Pool, error) {
	dsn, err := GetDatabaseDSNFromEnv()
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
