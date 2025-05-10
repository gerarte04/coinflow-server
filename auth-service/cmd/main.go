package main

import (
	"coinflow/coinflow-server/auth-service/config"
	apiGrpc "coinflow/coinflow-server/auth-service/internal/api/grpc"
	"coinflow/coinflow-server/auth-service/internal/repository/postgres"
	"coinflow/coinflow-server/auth-service/internal/usecases/service"
	pb "coinflow/coinflow-server/gen/auth_service/golang"
	pkgConfig "coinflow/coinflow-server/pkg/config"
	pkgPostgres "coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/pkg/infra/cache/redis"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
	flg := pkgConfig.ParseFlags()
	pkgConfig.MustLoadConfig(flg.ConfigPath, &cfg)

	lis, err := net.Listen("tcp", cfg.AuthSvcCfg.Host + ":" + cfg.AuthSvcCfg.Port)
	if err != nil {
		log.Fatalf("failed to listen address: %s", err.Error())
	}

	cache, err := redis.NewRedisCache(cfg.RedisCfg)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	dbConn, err := pkgPostgres.NewPostgresConn(cfg.PostgresCfg)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	usersRepo := postgres.NewUsersRepo(dbConn)
	userSvc, err := service.NewUserService(usersRepo, cache, cfg.JwtCfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	authServer := apiGrpc.NewAuthServer(userSvc)

	svr := grpc.NewServer()
	pb.RegisterAuthServer(svr, authServer)

	if err := svr.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
