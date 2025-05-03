package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/pkg/utils"
)

type Config struct {
	ClassificationSvcCfg 	pkgConfig.GrpcConfig 		`yaml:"classification_service" env-required:"true"`
	CollectionSvcCfg 		pkgConfig.GrpcConfig 		`yaml:"collection_service" env-required:"true"`
	SvcCfg 					utils.TranslateConfig		`yaml:"translate"`
	PostgresCfg 			postgres.PostgresConfig 	`yaml:"postgres" env-required:"true"`
}
