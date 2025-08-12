package service

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/storage-service/config"
	"coinflow/coinflow-server/storage-service/internal/models"
	"coinflow/coinflow-server/storage-service/internal/repository"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "coinflow/coinflow-server/gen/classification_service/golang"

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
	httpCli *http.Client
	clfCli pb.ClassificationClient
	clfSvcCfg pkgConfig.GrpcConfig
	svcCfg config.ServiceConfig
	
	txRepo repository.TransactionsRepo
	catsRepo repository.CategoriesRepo
	categories []string

	categoryChan chan *CategoryQuery
}

func NewTransactionsService(
	clfSvcCfg pkgConfig.GrpcConfig,
	svcCfg config.ServiceConfig,
	txRepo repository.TransactionsRepo,
	catsRepo repository.CategoriesRepo,
) (*TransactionsService, error) {
	const op = "NewTransactionsService"

	addr := fmt.Sprintf("%s:%s", clfSvcCfg.Host, clfSvcCfg.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Millisecond)
	defer cancel()

	categories, err := catsRepo.GetCategories(ctx)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	categoryChan := make(chan *CategoryQuery, svcCfg.CategoryChanBuffer)
	txService := &TransactionsService{
		httpCli: &http.Client{},
		clfCli: pb.NewClassificationClient(conn),
		clfSvcCfg: clfSvcCfg,
		svcCfg: svcCfg,

		txRepo: txRepo,
		catsRepo: catsRepo,
		categories: categories,

		categoryChan: categoryChan,
	}
	
	go txService.listenCategoryChannel()

	return txService, nil
}

func (s *TransactionsService) listenCategoryChannel() {
	for query := range s.categoryChan {
		select {
		case <-query.ctx.Done():
			log.Printf("[WARN] failed to put category: %s", query.ctx.Err())
		default:
			err := s.GetAndPutCategory(query.ctx, query.tx)

			if err != nil {
				log.Printf("[WARN] failed to put category: %s", err.Error())
			} else {
				log.Printf("[INFO] successfully got and put category")
			}
		}

		query.cancel()
	}
}

func (s *TransactionsService) GetTransaction(ctx context.Context, userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error) {
	return s.txRepo.GetTransaction(ctx, userId, txId)
}

func (s *TransactionsService) GetTransactionsInPeriod(
	ctx context.Context,
	userId uuid.UUID,
	begin time.Time, end time.Time,
	limit int,
) ([]*models.Transaction, error) {
	return s.txRepo.GetTransactionsInPeriod(ctx, userId, begin, end, limit)
}

func (s *TransactionsService) GetAndPutCategory(ctx context.Context, tx *models.Transaction) error {
	const op = "TransactionsService.GetAndPutCategory"

	text := tx.Description

	if s.svcCfg.DoTranslate {
		var err error
		text, err = utils.TranslateToLanguage(s.httpCli, tx.Description, utils.LanguageEnglish, s.svcCfg.TranslateCfg)
		
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	resp, err := s.clfCli.GetTextCategory(ctx, &pb.GetTextCategoryRequest{
		Text: text,
		Labels: s.categories,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.txRepo.PutCategory(ctx, tx.Id, resp.Category)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TransactionsService) PostTransaction(ctx context.Context, tx *models.Transaction, withAutoCategory bool) (*models.Transaction, error) {
	const op = "TransactionsService.PostTransaction"

	var err error
	
	if withAutoCategory {
		tx, err = s.txRepo.PostTransactionWithoutCategory(ctx, tx)
		if err != nil {
			return nil, err
		}

		ctx, cancel := context.WithTimeout(context.Background(), s.svcCfg.CategoryTimeout)
		s.categoryChan <- &CategoryQuery{ctx: ctx, cancel: cancel, tx: tx}
	} else {
		tx, err = s.txRepo.PostTransaction(ctx, tx)
		if err != nil {
			return nil, err
		}
	}

	return tx, nil
}
