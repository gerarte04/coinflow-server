package usecases

import "coinflow/coinflow-server/collection-service/internal/models"

type CollectionService interface {
	CollectCategory(tx *models.Transaction) error
}
