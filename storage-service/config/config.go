package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
)

type Config struct {
	CollectionSvcCfg 	pkgConfig.GrpcConfig 		`yaml:"collection_service" env-required:"true"`
	StorageSvcCfg		pkgConfig.GrpcConfig		`yaml:"storage_service" env-required:"true"`
	PostgresCfg 		postgres.PostgresConfig 	`yaml:"postgres" env-required:"true"`
}
