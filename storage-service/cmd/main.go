package main

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	pkgPostgres "coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/storage-service/config"
	"coinflow/coinflow-server/storage-service/internal/repository/postgres"
	"coinflow/coinflow-server/storage-service/internal/usecases/service"
	"log"
	"net"

	pb "coinflow/coinflow-server/gen/storage_service/golang"
	apiGrpc "coinflow/coinflow-server/storage-service/internal/api/grpc"

	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
	flg := pkgConfig.ParseFlags()
	pkgConfig.MustLoadConfig(flg.ConfigPath, &cfg)

	lis, err := net.Listen("tcp", cfg.StorageSvcCfg.Host + ":" + cfg.StorageSvcCfg.Port)

	if err != nil {
		log.Fatalf("failed to listen address: %s", err.Error())
	}

	pool, err := pkgPostgres.NewPostgresPool(cfg.PostgresCfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	txRepo := postgres.NewTransactionsRepo(pool)
	txSvc, err := service.NewTransactionsService(txRepo, cfg.CollectionSvcCfg, cfg.SvcCfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	stServer := apiGrpc.NewStorageServer(txSvc, cfg.SvcCfg)

	svr := grpc.NewServer()
	pb.RegisterStorageServer(svr, stServer)

	if err := svr.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
