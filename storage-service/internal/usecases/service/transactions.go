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

type CategoryQuery struct {
	ctx context.Context
	cancel context.CancelFunc
	tx *models.Transaction
}

type TransactionsService struct {
	txRepo repository.TransactionsRepo
	collectClient pb.CollectionClient
	collSvcConfig pkgConfig.GrpcConfig
	svcCfg config.ServiceConfig

	categoryChan chan *CategoryQuery
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

	categoryChan := make(chan *CategoryQuery, svcCfg.CategoryChanBuffer)
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
	for query := range s.categoryChan {
		select {
		case <-query.ctx.Done():
			log.Printf("[WARN] failed to put category: %s", query.ctx.Err())
			break
		default:
			err := s.GetAndPutCategory(query.ctx, query.tx)

			if err != nil {
				log.Printf("[WARN] failed to put category: %s", err.Error())
			} else {
				log.Printf("[INFO] successfully got and put category")
			}

			break
		}

		query.cancel()
	}
}

func (s *TransactionsService) GetTransaction(ctx context.Context, userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error) {
	return s.txRepo.GetTransaction(ctx, userId, txId)
}

func (s *TransactionsService) GetTransactionsInPeriod(ctx context.Context, userId uuid.UUID, begin time.Time, end time.Time) ([]*models.Transaction, error) {
	return s.txRepo.GetTransactionsInPeriod(ctx, userId, begin, end)
}

func (s *TransactionsService) GetAndPutCategory(ctx context.Context, tx *models.Transaction) error {
	const op = "TransactionsService.GetAndPutCategory"

	pbTx, err := ConvertModelTransactionToProtobuf(tx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	resp, err := s.collectClient.GetTransactionCategory(ctx, &pb.GetTransactionCategoryRequest{Tx: pbTx})
	if err != nil {
		return fmt.Errorf("%s: received error from collector: %w", op, err)
	}

	err = s.txRepo.PutCategory(ctx, tx.Id, resp.Category)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TransactionsService) PostTransaction(ctx context.Context, tx *models.Transaction, withAutoCategory bool) (uuid.UUID, error) {
	const op = "TransactionsService.PostTransaction"

	var txId uuid.UUID
	var err error
	
	if withAutoCategory {
		txId, err = s.txRepo.PostTransactionWithoutCategory(ctx, tx)
		if err != nil {
			return uuid.Nil, err
		}

		tx.Id = txId

		ctx, cancel := context.WithTimeout(context.Background(), s.svcCfg.CategoryTimeout)
		s.categoryChan <- &CategoryQuery{ctx: ctx, cancel: cancel, tx: tx}
	} else {
		txId, err = s.txRepo.PostTransaction(ctx, tx)
		if err != nil {
			return uuid.Nil, err
		}
	}

	return txId, nil
}
