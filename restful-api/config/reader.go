package config

import (
	"coinflow/coinflow-server/pkg/database/postgres"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpConfig struct {
	Host				string 						`yaml:"host" env:"HTTP_HOST" env-required:"true"`
	Port 				string 						`yaml:"port" env:"HTTP_PORT" env-required:"true"`
}

type GrpcConfig struct {
	Host 				string 						`yaml:"host" env-required:"true"`
	Port 				string 						`yaml:"port" env-required:"true"`
	RequestTimeout 		time.Duration 				`yaml:"req_timeout" env-default:"300ms"`
}

type Config struct {
	HttpCfg 			HttpConfig 					`yaml:"http" env-required:"true"`
	CollectionSvcCfg 	GrpcConfig 					`yaml:"collection_service" env-required:"true"`
	PostgresCfg 		postgres.PostgresConfig 	`yaml:"postgres" env-required:"true"`
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
