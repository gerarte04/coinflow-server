package main

import (
	"coinflow/coinflow-server/collection-service/config"
	apiGrpc "coinflow/coinflow-server/collection-service/internal/api/grpc"
	"coinflow/coinflow-server/collection-service/internal/repository/postgres"
	"coinflow/coinflow-server/collection-service/internal/usecases/service"
	pb "coinflow/coinflow-server/gen/collection_service/golang"
	pkgPostgres "coinflow/coinflow-server/pkg/database/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
	flg := config.ParseFlags()
	config.MustLoadConfig(flg.ConfigPath, &cfg)

	lis, err := net.Listen("tcp", cfg.CollectionSvcCfg.Host + ":" + cfg.CollectionSvcCfg.Port)

	if err != nil {
		log.Fatalf("failed to listen address: %s", err.Error())
	}

	dbConn, err := pkgPostgres.NewPostgresConn(cfg.PostgresCfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	catsRepo := postgres.NewCategoriesRepo(dbConn)
	collectSvc, err := service.NewCollectionService(cfg.SvcCfg, cfg.ClassificationSvcCfg, catsRepo)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	cfServer := apiGrpc.NewCollectionServer(collectSvc)

	svr := grpc.NewServer()
	pb.RegisterCollectionServer(svr, cfServer)

	if err := svr.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
