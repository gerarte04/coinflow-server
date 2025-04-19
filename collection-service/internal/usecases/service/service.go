package service

import (
	"coinflow/coinflow-server/collection-service/config"
	"coinflow/coinflow-server/collection-service/internal/models"
	"context"
	"fmt"
	"log"
	"net/http"

	pb "coinflow/coinflow-server/gen/classification_service/golang"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	labels = []string{"food", "sport", "health", "investment", "entertainment", "other"}
)

type CollectionService struct {
	httpCli *http.Client
	grpcCli pb.ClassificationClient
	svcCfg config.ServicesConfig
	grpcCfg config.GrpcConfig
}

func NewCollectionService(
	svcCfg config.ServicesConfig,
	grpcCfg config.GrpcConfig,
) (*CollectionService, error) {
	addr := fmt.Sprintf("%s:%s", grpcCfg.ClassificationServiceHost, grpcCfg.ClassificationServicePort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("service.NewCollectionService: %w", err)
	}

	return &CollectionService{
		httpCli: &http.Client{},
		grpcCli: pb.NewClassificationClient(conn),
		svcCfg: svcCfg,
		grpcCfg: grpcCfg,
	}, nil
}

func (s *CollectionService) CollectCategory(ts *models.Transaction) error {
	const method = "CollectionService.CollectCategory"

	text, err := TranslateToLanguage(s.httpCli, ts.Description, LanguageEnglish, s.svcCfg)

	if err != nil {
		return err
	}

	resp, err := s.grpcCli.GetTextCategory(context.Background(), &pb.GetTextCategoryRequest{
		Text: text,
		Labels: labels,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	log.Printf("successfully detected category: desc = \"%s\", cat = \"%s\"", ts.Description, resp.Category)

	return nil
}
