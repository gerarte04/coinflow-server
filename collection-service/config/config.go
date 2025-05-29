package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/pkg/utils"
)

type ServiceConfig struct {
	DoTranslate 			bool						`yaml:"do_translate" env:"DO_TRANSLATE" env-default:"false"`
	TranslateCfg			utils.TranslateConfig		`yaml:"translate"`
}

type Config struct {
	ClassificationSvcCfg 	pkgConfig.GrpcConfig 		`yaml:"classification_service" env-required:"true"`
	CollectionSvcCfg 		pkgConfig.GrpcConfig 		`yaml:"collection_service" env-required:"true"`
	SvcCfg					ServiceConfig				`yaml:"service"`
	PostgresCfg 			postgres.PostgresConfig 	`yaml:"postgres" env-required:"true"`
}
