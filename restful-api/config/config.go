package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
)

type Config struct {
	HttpCfg 			pkgConfig.HttpConfig 		`yaml:"http" env-required:"true"`
	CollectionSvcCfg 	pkgConfig.GrpcConfig 		`yaml:"collection_service" env-required:"true"`
	PostgresCfg 		postgres.PostgresConfig 	`yaml:"postgres" env-required:"true"`
}
