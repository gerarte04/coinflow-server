package service

import (
	"coinflow/coinflow-server/collection-service/config"
	"coinflow/coinflow-server/collection-service/internal/models"
	"log"
	"net/http"
)

type CollectionService struct {
	cli *http.Client
	svcCfg config.ServicesConfig
}

func NewCollectionService(svcCfg config.ServicesConfig) *CollectionService {
	return &CollectionService{
		cli: &http.Client{},
		svcCfg: svcCfg,
	}
}

func (s *CollectionService) CollectCategory(ts *models.Transaction) error {
	text, err := TranslateToLanguage(s.cli, ts.Description, LanguageEnglish, s.svcCfg)

	if err != nil {
		return err
	}

	log.Printf("successful translation:\n%s\n", text)

	return nil
}
