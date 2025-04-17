package usecases

import "coinflow/coinflow-server/collection-service/internal/models"

type CollectionService interface {
	CollectCategory(ts *models.Transaction) error
}
