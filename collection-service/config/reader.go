package config

import (
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/pkg/utils"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type GrpcConfig struct {
	Host 					string 						`yaml:"host" env-required:"true"`
	Port 					string 						`yaml:"port" env-required:"true"`
}

type Config struct {
	ClassificationSvcCfg 	GrpcConfig 					`yaml:"classification_service" env-required:"true"`
	CollectionSvcCfg 		GrpcConfig 					`yaml:"collection_service" env-required:"true"`
	SvcCfg 					utils.TranslateConfig		`yaml:"translate"`
	PostgresCfg 			postgres.PostgresConfig 	`yaml:"postgres" env-required:"true"`
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
