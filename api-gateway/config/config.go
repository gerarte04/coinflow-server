package config

import pkgConfig "coinflow/coinflow-server/pkg/config"

type Config struct {
	HttpCfg 	pkgConfig.HttpConfig	`yaml:"http" env-required:"true"`
	StorageCfg 	pkgConfig.GrpcConfig	`yaml:"storage_service" env-required:"true"`
}
