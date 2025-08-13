package clickhouse

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type ClickhouseConfig struct {
	Addr              string        `yaml:"addr" env:"CLICKHOUSE_ADDR" env-required:"true"`
	DbName            string        `yaml:"db" env:"CLICKHOUSE_DB" env-required:"true"`
	User              string        `yaml:"user" env:"CLICKHOUSE_USER" env-required:"true"`
	Password          string        `yaml:"password" env:"CLICKHOUSE_PASSWORD" env-required:"true"`
	ConnectionTimeout time.Duration `yaml:"connection_timeout" env-default:"500ms"`
}

func NewClickhouseConn(cfg ClickhouseConfig) (driver.Conn, error) {
	const op = "NewClickhouseConn"

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
	defer cancel()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.Addr},
		Auth: clickhouse.Auth{
			Database: cfg.DbName,
			Username: cfg.User,
			Password: cfg.Password,
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return conn, nil
}
