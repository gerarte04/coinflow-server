package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
)

type ServiceConfig struct {
	CategoryChanBuffer 	int							`yaml:"category_chan_buffer" env:"CATEGORY_CHAN_BUFFER" env-default:"128"`
}

type Config struct {
	CollectionSvcCfg 	pkgConfig.GrpcConfig 		`yaml:"collection_service" env-required:"true"`
	StorageSvcCfg		pkgConfig.GrpcConfig		`yaml:"storage_service" env-required:"true"`
	PostgresCfg 		postgres.PostgresConfig 	`yaml:"postgres" env-required:"true"`
	SvcCfg				ServiceConfig				`yaml:"service"`
}
