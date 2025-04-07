package main

import (
	"coinflow/coinflow-server/restful-api/config"
	api "coinflow/coinflow-server/restful-api/internal/api/http"
	tsRepo "coinflow/coinflow-server/restful-api/internal/repository/mocks"
	tsService "coinflow/coinflow-server/restful-api/internal/usecases/service"
	"fmt"
	"log"

	_ "coinflow/coinflow-server/restful-api/docs"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Coinflow API
// @version 1.0
// @description Restful API for Coinflow service

// @host localhost:8080
// @BasePath /
func main() {
    var cfg config.Config
    flg := config.ParseFlags()
    config.MustLoadConfig(flg.ConfigPath, &cfg)

    tsRepo := tsRepo.NewTransactionsRepoMock()
    tsSvc := tsService.NewTransactionsService(tsRepo)
    cfServer := api.NewCoinflowServer(tsSvc)
    
    engine := gin.Default()
    cfServer.RouteHandlers(engine,
        cfServer.WithStandardUserHandlers(),
    )
    engine.GET("/swagger/*path", gin.WrapF(httpSwagger.WrapHandler))

    addr := fmt.Sprintf("%s:%s", cfg.HttpCfg.Host, cfg.HttpCfg.Port)
    
    if err := engine.Run(addr); err != nil {
        log.Fatalf("failed to serve: %s", err.Error())
    }
}
