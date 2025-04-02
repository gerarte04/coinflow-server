package main

import (
	"coinflow/coinflow-server/config"
	pb "coinflow/coinflow-server/gen/cfapi"
	api "coinflow/coinflow-server/internal/api/grpc"
	tsRepo "coinflow/coinflow-server/internal/repository/mocks/transactions"
	tsService "coinflow/coinflow-server/internal/usecases/transactions"
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

    tsRepo := tsRepo.NewTransactionsRepoMock()
    tsSvc := tsService.NewTransactionsService(tsRepo)
    cfServer := api.NewCoinflowServer(tsSvc)

    svr := grpc.NewServer()
    pb.RegisterCoinflowServer(svr, cfServer)

    if err := svr.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %s", err.Error())
    }
}
