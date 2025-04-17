package service

import (
	"coinflow/coinflow-server/collect-service/config"
	"coinflow/coinflow-server/collect-service/internal/models"
	"log"
	"net/http"
)

type CollectService struct {
	cli *http.Client
	svcCfg config.ServicesConfig
}

func NewCollectService(svcCfg config.ServicesConfig) *CollectService {
	return &CollectService{
		cli: &http.Client{},
		svcCfg: svcCfg,
	}
}

func (s *CollectService) CollectCategory(ts *models.Transaction) error {
	text, err := TranslateToLanguage(s.cli, ts.Description, LanguageEnglish, s.svcCfg)

	if err != nil {
		return err
	}

	log.Printf("successful translation:\n%s\n", text)

	return nil
}
