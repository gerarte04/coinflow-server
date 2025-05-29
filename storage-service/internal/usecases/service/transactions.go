package service

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/storage-service/config"
	"coinflow/coinflow-server/storage-service/internal/models"
	"coinflow/coinflow-server/storage-service/internal/repository"
	"context"
	"fmt"
	"log"
	"time"

	pb "coinflow/coinflow-server/gen/collection_service/golang"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TransactionsService struct {
	txRepo repository.TransactionsRepo
	collectClient pb.CollectionClient
	collSvcConfig pkgConfig.GrpcConfig
	svcCfg config.ServiceConfig

	categoryChan chan *models.Transaction
}

func NewTransactionsService(
	txRepo repository.TransactionsRepo,
	collSvcConfig pkgConfig.GrpcConfig,
	svcCfg config.ServiceConfig,
) (*TransactionsService, error) {
	const op = "NewTransactionsService"

	addr := fmt.Sprintf("%s:%s", collSvcConfig.Host, collSvcConfig.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	categoryChan := make(chan *models.Transaction, svcCfg.CategoryChanBuffer)
	txService := &TransactionsService{
		txRepo: txRepo,
		collectClient: pb.NewCollectionClient(conn),
		collSvcConfig: collSvcConfig,
		svcCfg: svcCfg,

		categoryChan: categoryChan,
	}
	
	go txService.ListenCategoryChannel()

	return txService, nil
}

func (s *TransactionsService) ListenCategoryChannel() {
	for tx := range s.categoryChan {
		err := s.GetAndPutCategory(tx)

		if err != nil {
			log.Printf("[WARN] failed to put category: %s", err.Error())
		} else {
			log.Printf("[INFO] successfully got and put category")
		}
	}
}

func (s *TransactionsService) GetTransaction(userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error) {
	return s.txRepo.GetTransaction(userId, txId)
}

func (s *TransactionsService) GetTransactionsInPeriod(userId uuid.UUID, begin time.Time, end time.Time) ([]*models.Transaction, error) {
	return s.txRepo.GetTransactionsInPeriod(userId, begin, end)
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

func (s *TransactionsService) PostTransaction(tx *models.Transaction, withAutoCategory bool) (uuid.UUID, error) {
	const op = "TransactionsService.PostTransaction"

	var txId uuid.UUID
	var err error
	
	if withAutoCategory {
		txId, err = s.txRepo.PostTransactionWithoutCategory(tx)
		if err != nil {
			return uuid.Nil, err
		}

		tx.Id = txId
		s.categoryChan <- tx
	} else {
		txId, err = s.txRepo.PostTransaction(tx)
		if err != nil {
			return uuid.Nil, err
		}
	}

	return txId, nil
}
