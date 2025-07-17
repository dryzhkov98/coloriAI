package postgres

import (
	"coloriAI/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, cfg config.DbConfig, logger *zap.Logger) (*Database, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DbUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres url is not empty: %w", err)
	}

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdle

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		logger.Error("error while creating pool", zap.Error(err))
		return nil, fmt.Errorf("error while creating pool: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctxPing); err != nil {
		pool.Close()
		return nil, fmt.Errorf("error while conntect to DB: %w", err)
	}

	logger.Info("DB is connected successfully\n")

	return &Database{
		Pool: pool,
	}, nil
}
