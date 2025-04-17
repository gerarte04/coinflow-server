package service

import (
	"coinflow/coinflow-server/restful-api/config"
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"context"
	"fmt"
	"log"

	pb "coinflow/coinflow-server/gen/collect_service"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TransactionsService struct {
	tsRepo repository.TransactionsRepo
	collectClient pb.CollectClient
	grpcConfig config.GrpcConfig
}

func NewTransactionsService(
	tsRepo repository.TransactionsRepo,
	grpcConfig config.GrpcConfig,
) (*TransactionsService, error) {
	addr := fmt.Sprintf("%s:%s", grpcConfig.Host, grpcConfig.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("service.NewTransactionsService: %w", err)
	}

	return &TransactionsService{
		tsRepo: tsRepo,
		collectClient: pb.NewCollectClient(conn),
		grpcConfig: grpcConfig,
	}, nil
}

func (s *TransactionsService) GetTransaction(tsId uuid.UUID) (*models.Transaction, error) {
	return s.tsRepo.GetTransaction(tsId)
}

func (s *TransactionsService) PostTransaction(ts *models.Transaction) (uuid.UUID, error) {
	const method = "TransactionsService.PostTransaction"

	
	pbTs, err := ConvertModelTransactionToProtobuf(ts)
	
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", method, err)
	}

	go func() {
		//ctx, cancel := context.WithTimeout(context.Background(), s.grpcConfig.RequestExpireTimeout)
		ctx := context.Background()
		//defer cancel()
	
		_, err := s.collectClient.GetTransactionCategory(ctx, &pb.GetTransactionCategoryRequest{Ts: pbTs})
	
		if err != nil {
			log.Printf("received error from collector: %s", err.Error())
		}
	}()

	id, err := s.tsRepo.PostTransaction(ts)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
