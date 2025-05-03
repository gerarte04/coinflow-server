package main

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/database/postgres"
	pkgHandlers "coinflow/coinflow-server/pkg/http/handlers"
	"coinflow/coinflow-server/restful-api/config"
	api "coinflow/coinflow-server/restful-api/internal/api/http"
	txRepo "coinflow/coinflow-server/restful-api/internal/repository/postgres"
	txService "coinflow/coinflow-server/restful-api/internal/usecases/service"
	"fmt"
	"log"

	_ "coinflow/coinflow-server/restful-api/docs"

	"github.com/gin-gonic/gin"
)

// @title Coinflow API
// @version 1.0
// @description Restful API for Coinflow service

// @host localhost:8080
// @BasePath /
func main() {
	var cfg config.Config
	flg := pkgConfig.ParseFlags()
	pkgConfig.MustLoadConfig(flg.ConfigPath, &cfg)

	dbConn, err := postgres.NewPostgresConn(cfg.PostgresCfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	txRepo := txRepo.NewTransactionsRepo(dbConn)
	txSvc, err := txService.NewTransactionsService(txRepo, cfg.CollectionSvcCfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	cfServer := api.NewCoinflowServer(txSvc)

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
