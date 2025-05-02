package service

import (
	"coinflow/coinflow-server/collection-service/config"
	"coinflow/coinflow-server/collection-service/internal/models"
	"coinflow/coinflow-server/collection-service/internal/repository"
	"context"
	"fmt"
	"net/http"

	pb "coinflow/coinflow-server/gen/classification_service/golang"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CollectionService struct {
	httpCli *http.Client
	grpcCli pb.ClassificationClient
	svcCfg config.ServicesConfig
	clfSvcCfg config.GrpcConfig
	txRepo repository.TransactionsRepo
	catsRepo repository.CategoriesRepo
	categories []string
}

func NewCollectionService(
	svcCfg config.ServicesConfig,
	clfSvcCfg config.GrpcConfig,
	txRepo repository.TransactionsRepo,
	catsRepo repository.CategoriesRepo,
) (*CollectionService, error) {
	const method = "service.NewCollectionService"

	addr := fmt.Sprintf("%s:%s", clfSvcCfg.Host, clfSvcCfg.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	categories, err := catsRepo.GetCategories()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return &CollectionService{
		httpCli: &http.Client{},
		grpcCli: pb.NewClassificationClient(conn),
		svcCfg: svcCfg,
		clfSvcCfg: clfSvcCfg,
		txRepo: txRepo,
		catsRepo: catsRepo,
		categories: categories,
	}, nil
}

func (s *CollectionService) CollectCategory(tx *models.Transaction) error {
	const op = "CollectionService.CollectCategory"

	text, err := TranslateToLanguage(s.httpCli, tx.Description, LanguageEnglish, s.svcCfg)
	if err != nil {
		return err
	}

	resp, err := s.grpcCli.GetTextCategory(context.Background(), &pb.GetTextCategoryRequest{
		Text: text,
		Labels: s.categories,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.txRepo.PutCategory(tx.Id, resp.Category)
	if err != nil {
		return err
	}

	return nil
}
