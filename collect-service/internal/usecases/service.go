package usecases

import "coinflow/coinflow-server/collect-service/internal/models"

type CollectService interface {
	CollectCategory(ts *models.Transaction) error
}
