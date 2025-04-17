package main

import (
	"coinflow/coinflow-server/collection-service/config"
	apiGrpc "coinflow/coinflow-server/collection-service/internal/api/grpc"
	"coinflow/coinflow-server/collection-service/internal/usecases/service"
	pb "coinflow/coinflow-server/gen/collection_service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
    flg := config.ParseFlags()
    config.MustLoadConfig(flg.ConfigPath, &cfg)

    lis, err := net.Listen("tcp", cfg.GrpcCfg.Host + ":" + cfg.GrpcCfg.Port)
    if err != nil {
        log.Fatalf("failed to listen address: %s", err.Error())
    }

	collectSvc := service.NewCollectionService(cfg.SvcCfg)
    cfServer := apiGrpc.NewCollectionServer(collectSvc)

    svr := grpc.NewServer()
    pb.RegisterCollectionServer(svr, cfServer)

    if err := svr.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %s", err.Error())
    }
}
