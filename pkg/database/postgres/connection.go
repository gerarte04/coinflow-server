package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PostgresConfig struct {
	Host string `yaml:"host" env:"POSTGRES_HOST"`
	Port string `yaml:"port" env:"POSTGRES_PORT"`
	DbName string `yaml:"db" env:"POSTGRES_DB"`
	User string `yaml:"user" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
}

func NewPostgresConn(cfg PostgresConfig) (*pgx.Conn, error) {
	const method = "postgres.NewPostgresConn"

	conn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.DbName, cfg.User, cfg.Password,
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	} else if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return conn, nil
}
