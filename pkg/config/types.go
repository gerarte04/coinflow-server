package config

import "time"

type HttpConfig struct {
	Host				string 						`yaml:"host" env:"HTTP_HOST" env-required:"true"`
	Port 				string 						`yaml:"port" env:"HTTP_PORT" env-required:"true"`
}

type GrpcConfig struct {
	Host 				string 						`yaml:"host" env-required:"true"`
	Port 				string 						`yaml:"port" env-required:"true"`
	RequestTimeout 		time.Duration 				`yaml:"req_timeout" env-default:"300ms"`
}
