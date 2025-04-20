package config

import (
	"coinflow/coinflow-server/pkg/database/postgres"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type GrpcConfig struct {
	Host string `yaml:"host" env:"GRPC_COLLECTION_SERVICE_HOST"`
	Port string `yaml:"port" env:"GRPC_COLLECTION_SERVICE_PORT"`
	ClassificationServiceHost string `yaml:"classification_service_host" env:"GRPC_CLASSIFICATION_SERVICE_HOST"`
	ClassificationServicePort string `yaml:"classification_service_port" env:"GRPC_CLASSIFICATION_SERVICE_PORT"`
}

type ServicesConfig struct {
	TranslateApiAddress string `yaml:"translate_api_address" env:"TRANSLATE_API_ADDRESS"`
	TranslateApiKey string `yaml:"translate_api_key" env:"TRANSLATE_API_KEY"`
}

type Config struct {
	GrpcCfg GrpcConfig `yaml:"grpc"`
	SvcCfg ServicesConfig `yaml:"services"`
	PostgresCfg postgres.PostgresConfig `yaml:"postgres"`
}

func MustLoadConfig(path string, cfg any) {
	if path == "" {
		log.Fatalf("path is not set")
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist by this path: %s", path)
	} else if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf("error reading config: %s", err)
	}
}
