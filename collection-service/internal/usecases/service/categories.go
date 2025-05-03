package service

import (
	"coinflow/coinflow-server/collection-service/internal/models"
	"coinflow/coinflow-server/collection-service/internal/repository"
	pb "coinflow/coinflow-server/gen/classification_service/golang"
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/utils"
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CollectionService struct {
	httpCli *http.Client
	grpcCli pb.ClassificationClient
	tlsCfg utils.TranslateConfig
	clfSvcCfg pkgConfig.GrpcConfig
	catsRepo repository.CategoriesRepo
	categories []string
}

func NewCollectionService(
	tlsCfg utils.TranslateConfig,
	clfSvcCfg pkgConfig.GrpcConfig,
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
		tlsCfg: tlsCfg,
		clfSvcCfg: clfSvcCfg,
		catsRepo: catsRepo,
		categories: categories,
	}, nil
}

func (s *CollectionService) CollectCategory(ctx context.Context, tx *models.Transaction) (string, error) {
	const op = "CollectionService.CollectCategory"

	text, err := utils.TranslateToLanguage(s.httpCli, tx.Description, utils.LanguageEnglish, s.tlsCfg)
	if err != nil {
		return "", err
	}

	resp, err := s.grpcCli.GetTextCategory(ctx, &pb.GetTextCategoryRequest{
		Text: text,
		Labels: s.categories,
	})

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.Category, nil
}
