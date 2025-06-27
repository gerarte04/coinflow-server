package config

import (
	"coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/pkg/infra/cache/redis"
	"time"
)

type JwtConfig struct {
	PrivateKeyPEM 			string			`env:"PRIVATE_KEY_PEM" env-required:"true"`
	PublicKeyPEM			string			`env:"PUBLIC_KEY_PEM" env-required:"true"`

	Issuer					string			`eng:"JWT_ISSUER" env-required:"true"`

	AccessExpirationTime 	time.Duration	`yaml:"access_expiration_time" env:"ACCESS_EXPIRATION_TIME" env-default:"30m"`
	RefreshExpirationTime 	time.Duration	`yaml:"refresh_expiration_time" env:"REFRESH_EXPIRATION_TIME" env-default:"7d"`
}

type ServiceConfig struct {
	AuthCookieName			string			`yaml:"auth_cookie_name" env:"AUTH_COOKIE_NAME" env-default:"accessToken"`
	HttpCodeHeaderName		string			`yaml:"http_code_header_name" env:"HTTP_CODE_HEADER_NAME" env-default:"x-http-code"`
}

type Config struct {
	AuthSvcCfg 		config.GrpcConfig			`yaml:"auth_service"`
	PostgresCfg 	postgres.PostgresConfig		`yaml:"postgres"`
	RedisCfg 		redis.RedisConfig			`yaml:"redis"`
	JwtCfg			JwtConfig					`yaml:"jwt"`
	SvcCfg			ServiceConfig				`yaml:"service"`
}
