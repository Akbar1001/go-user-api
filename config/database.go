package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(cfg *Config) (*pgx.Conn, error) {

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	conn, err := pgx.Connect(
		context.Background(),
		connString,
	)

	if err != nil {
		return nil, err
	}

	return conn, nil
}