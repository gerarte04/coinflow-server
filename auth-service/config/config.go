package config

import (
	"coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/pkg/infra/cache/redis"
	"time"
)

type JwtConfig struct {
	PrivateKeyBase64 		string			`env:"JWT_PRIVATE_KEY" env-required:"true"`
	PublicKeyBase64 		string			`env:"JWT_PUBLIC_KEY" env-required:"true"`

	AccessExpirationTime 	time.Duration	`yaml:"access_expiration_time" env:"ACCESS_EXPIRATION_TIME" env-default:"30m"`
	RefreshExpirationTime 	time.Duration	`yaml:"refresh_expiration_time" env:"REFRESH_EXPIRATION_TIME" env-default:"7d"`
}

type Config struct {
	AuthSvcCfg 		config.GrpcConfig			`yaml:"auth_service"`
	PostgresCfg 	postgres.PostgresConfig		`yaml:"postgres"`
	RedisCfg 		redis.RedisConfig			`yaml:"redis"`
	JwtCfg			JwtConfig					`yaml:"jwt"`
}
