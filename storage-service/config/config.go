package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/pkg/utils"
	"time"
)

type ServiceConfig struct {
	DoTranslate  bool                  `yaml:"do_translate" env:"DO_TRANSLATE" env-default:"false"`
	TranslateCfg utils.TranslateConfig `yaml:"translate"`

	CategoryChanBuffer int           `yaml:"category_chan_buffer" env:"CATEGORY_CHAN_BUFFER" env-default:"128"`
	CategoryTimeout    time.Duration `yaml:"category_timeout" env:"CATEGORY_TIMEOUT" env-default:"5m"`

	HttpCodeHeaderName string `yaml:"http_code_header_name" env:"HTTP_CODE_HEADER_NAME" env-default:"x-http-code"`
}

type Config struct {
	ClfSvcCfg     pkgConfig.GrpcConfig    `yaml:"classification_service" env-required:"true"`
	StorageSvcCfg pkgConfig.GrpcConfig    `yaml:"storage_service" env-required:"true"`
	PostgresCfg   postgres.PostgresConfig `yaml:"postgres" env-required:"true"`
	SvcCfg        ServiceConfig           `yaml:"service"`
}
