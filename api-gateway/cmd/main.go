package main

import (
	"coinflow/coinflow-server/api-gateway/config"
	api "coinflow/coinflow-server/api-gateway/internal/api/http"
	"coinflow/coinflow-server/api-gateway/internal/clients/grpc/collect"
	"coinflow/coinflow-server/api-gateway/internal/clients/grpc/storage"
	pkgConfig "coinflow/coinflow-server/pkg/config"
	pkgHandlers "coinflow/coinflow-server/pkg/http/handlers"
	"fmt"
	"log"

	_ "coinflow/coinflow-server/api-gateway/docs"

	"github.com/gin-gonic/gin"
)

// @title Coinflow API
// @version 1.0
// @description API Gateway for Coinflow service

// @host localhost:8080
// @BasePath /
func main() {
	var cfg config.Config
	flg := pkgConfig.ParseFlags()
	pkgConfig.MustLoadConfig(flg.ConfigPath, &cfg)

	collClient := collect.NewCollectionClient()
	stClient, err := storage.NewStorageClient(cfg.StorageCfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	cfServer := api.NewCoinflowServer(stClient, collClient)

	engine := gin.New()
	cfServer.RouteHandlers(engine,
		pkgHandlers.WithLogger(),
		pkgHandlers.WithRecovery(),
		pkgHandlers.WithHealthCheck(),
		pkgHandlers.WithSwagger(),
		cfServer.WithStandardUserHandlers(),
	)

	addr := fmt.Sprintf("%s:%s", cfg.HttpCfg.Host, cfg.HttpCfg.Port)

	if err := engine.Run(addr); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
