package service

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
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
	collSvcConfig pkgConfig.GrpcConfig
}

func NewTransactionsService(
	txRepo repository.TransactionsRepo,
	collSvcConfig pkgConfig.GrpcConfig,
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

func (s *TransactionsService) GetAndPutCategory(tx *models.Transaction) error {
	const op = "TransactionsService.GetAndPutCategory"

	pbTx, err := ConvertModelTransactionToProtobuf(tx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	ctx := context.Background()

	resp, err := s.collectClient.GetTransactionCategory(ctx, &pb.GetTransactionCategoryRequest{Tx: pbTx})
	if err != nil {
		return fmt.Errorf("%s: received error from collector: %w", op, err)
	}

	err = s.txRepo.PutCategory(tx.Id, resp.Category)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
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

		tx.Id = txId

		go func() { 
			err := s.GetAndPutCategory(tx)

			if err != nil {
				log.Printf("[WARN] failed to put category: %s", err.Error())
			} else {
				log.Printf("[INFO] successfully got and put category")
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
