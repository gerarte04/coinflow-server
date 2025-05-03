package usecases

import (
	"coinflow/coinflow-server/collection-service/internal/models"
	"context"
)

type CollectionService interface {
	CollectCategory(ctx context.Context, tx *models.Transaction) (string, error)
}
