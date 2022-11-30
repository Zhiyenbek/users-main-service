package connection

import (
	"context"
	"fmt"

	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDB(cfg *config.DBConf) (*pgxpool.Pool, error) {
	dbURI := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.TimeOut)
	defer cancel()
	pool, err := pgxpool.Connect(ctx, dbURI)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
