package usecases

import (
	"coinflow/coinflow-server/storage-service/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
)

type TransactionsService interface {
	GetTransaction(ctx context.Context, userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(ctx context.Context, userId uuid.UUID, begin time.Time, end time.Time) ([]*models.Transaction, error)
	PostTransaction(ctx context.Context, tx *models.Transaction, withAutoCategory bool) (uuid.UUID, error)
}
