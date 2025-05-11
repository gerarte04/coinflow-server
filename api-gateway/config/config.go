package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"time"
)

type SecurityConfig struct {
	JwtPublicKeyBase64		string					`env:"JWT_PUBLIC_KEY" env-required:"true"`
	AccessExpirationTime	time.Duration 			`yaml:"access_expiration_time" env-default:"15m"`	
}

type Config struct {
	HttpCfg 				pkgConfig.HttpConfig	`yaml:"http" env-required:"true"`
	StorageCfg 				pkgConfig.GrpcConfig	`yaml:"storage_service" env-required:"true"`
	AuthCfg					pkgConfig.GrpcConfig	`yaml:"auth_service" env-required:"true"`
	SecurityCfg				SecurityConfig			`yaml:"security"`
}
