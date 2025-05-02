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
	txRepo repository.TransactionsRepo
	collectClient pb.CollectionClient
	collSvcConfig config.GrpcConfig
}

func NewTransactionsService(
	txRepo repository.TransactionsRepo,
	collSvcConfig config.GrpcConfig,
) (*TransactionsService, error) {
	const op = "NewTransactionsService"

	addr := fmt.Sprintf("%s:%s", collSvcConfig.Host, collSvcConfig.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &TransactionsService{
		txRepo: txRepo,
		collectClient: pb.NewCollectionClient(conn),
		collSvcConfig: collSvcConfig,
	}, nil
}

func (s *TransactionsService) GetTransaction(txId uuid.UUID) (*models.Transaction, error) {
	return s.txRepo.GetTransaction(txId)
}

func (s *TransactionsService) PostTransaction(tx *models.Transaction) (uuid.UUID, error) {
	const op = "TransactionsService.PostTransaction"

	var txId uuid.UUID
	var err error
	
	if tx.WithAutoCategory {
		txId, err = s.txRepo.PostTransactionWithoutCategory(tx)
		if err != nil {
			return uuid.Nil, err
		}

		go func() {
			pbTx, err := ConvertModelTransactionToProtobuf(tx)
			if err != nil {
				log.Printf("%s: %s", op, err.Error())
			}
	
			pbTx.Id = txId.String()

			//ctx, cancel := context.WithTimeout(context.Background(), s.grpcConfig.RequestExpireTimeout)
			ctx := context.Background()
			//defer cancel()
		
			_, err = s.collectClient.GetTransactionCategory(ctx, &pb.GetTransactionCategoryRequest{Tx: pbTx})
			if err != nil {
				log.Printf("received error from collector: %s", err.Error())
			}
		}()
	} else {
		txId, err = s.txRepo.PostTransaction(tx)
		if err != nil {
			return uuid.Nil, err
		}
	}

	return txId, nil
}
