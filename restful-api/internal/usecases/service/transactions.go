package service

import (
	"coinflow/coinflow-server/restful-api/config"
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"context"
	"fmt"
	"log"

	pb "coinflow/coinflow-server/gen/collection_service/golang"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TransactionsService struct {
	tsRepo repository.TransactionsRepo
	collectClient pb.CollectionClient
	collSvcConfig config.GrpcConfig
}

func NewTransactionsService(
	tsRepo repository.TransactionsRepo,
	collSvcConfig config.GrpcConfig,
) (*TransactionsService, error) {
	addr := fmt.Sprintf("%s:%s", collSvcConfig.Host, collSvcConfig.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("service.NewTransactionsService: %w", err)
	}

	return &TransactionsService{
		tsRepo: tsRepo,
		collectClient: pb.NewCollectionClient(conn),
		collSvcConfig: collSvcConfig,
	}, nil
}

func (s *TransactionsService) GetTransaction(tsId uuid.UUID) (*models.Transaction, error) {
	return s.tsRepo.GetTransaction(tsId)
}

func (s *TransactionsService) PostTransaction(ts *models.Transaction) (uuid.UUID, error) {
	const method = "TransactionsService.PostTransaction"
	
	id, err := s.tsRepo.PostTransaction(ts)
	
	if err != nil {
		return uuid.Nil, err
	}

	if ts.WithAutoCategory {
		go func() {
			pbTs, err := ConvertModelTransactionToProtobuf(ts)
			
			if err != nil {
				log.Printf("%s: %s", method, err.Error())
			}
	
			pbTs.Id = id.String()

			//ctx, cancel := context.WithTimeout(context.Background(), s.grpcConfig.RequestExpireTimeout)
			ctx := context.Background()
			//defer cancel()
		
			_, err = s.collectClient.GetTransactionCategory(ctx, &pb.GetTransactionCategoryRequest{Ts: pbTs})
		
			if err != nil {
				log.Printf("received error from collector: %s", err.Error())
			}
		}()
	}

	return id, nil
}
