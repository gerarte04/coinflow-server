package config

import pkgConfig "coinflow/coinflow-server/pkg/config"

type SwaggerConfig struct {
	HttpCfg     pkgConfig.HttpConfig `yaml:"http"`
	SwaggerPath string               `yaml:"swagger_path" env-default:"/v1/swagger"`
}
