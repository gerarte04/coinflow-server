package config

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/clickhouse"
)

type Config struct {
	CollectionSvcCfg pkgConfig.GrpcConfig        `yaml:"collection_service" env-required:"true"`
	ClickhouseCfg    clickhouse.ClickhouseConfig `yaml:"clickhouse" env-required:"true"`
}
