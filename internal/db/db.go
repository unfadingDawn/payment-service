package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cfg, _ := pgxpool.ParseConfig(dsn)
	return pgxpool.NewWithConfig(ctx, cfg)
}

func Ping(ctx context.Context, pool *pgxpool.Pool) error {
	if err := pool.Ping(ctx); err != nil {
		return err
	}

	log.Println("connected")
	return nil
}
