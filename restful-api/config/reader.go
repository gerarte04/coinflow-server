package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpConfig struct {
	Host string `yaml:"host" env:"HTTP_HOST"`
	Port string `yaml:"port" env:"HTTP_PORT"`
}

type GrpcConfig struct {
	Host string `yaml:"host" env:"GRPC_COLLECTION_SERVICE_HOST"`
	Port string `yaml:"port" env:"GRPC_COLLECTION_SERVICE_PORT"`
	RequestExpireTimeout time.Duration `yaml:"req_expire_timeout" env-default:"300ms"`
}

type Config struct {
	HttpCfg HttpConfig `yaml:"http"`
	GrpcCfg GrpcConfig `yaml:"grpc"`
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
