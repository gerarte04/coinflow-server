package redis

import (
	"coinflow/coinflow-server/pkg/infra/cache"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host			string		`yaml:"host" env:"REDIS_HOST" env-required:"true"`
	Port 			string		`yaml:"port" env:"REDIS_PORT" env-required:"true"`
	User 			string		`yaml:"user" env:"REDIS_USER" env-required:"true"`
	UserPassword 	string		`yaml:"user_password" env:"REDIS_USER_PASSWORD" env-required:"true"`
	DBNumber 		int			`yaml:"db_number" env:"REDIS_DB_NUMBER" env-default:"0"`
}

type RedisCache struct {
	cli *redis.Client
    redisCfg RedisConfig
}

func NewRedisCache(redisCfg RedisConfig) (*RedisCache, error) {
	const op = "NewRedisCache"

	cli := redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port),
        Username: redisCfg.User,
        Password: redisCfg.UserPassword,
        DB: redisCfg.DBNumber,
    })

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

    err := cli.Ping(ctx).Err()
    if err != nil {
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    return &RedisCache{
        cli: cli,
        redisCfg: redisCfg,
    }, nil
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	const op = "RedisCache.Get"

	val, err := c.cli.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("%s: %w", op, cache.ErrorKeyNotFound)
	} else if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return val, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	const op = "RedisCache.Set"

	val, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = c.cli.Set(ctx, key, val, expiration).Result()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
