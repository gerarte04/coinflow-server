package grpc

import (
	"coinflow/coinflow-server/api-gateway/internal/models"
	"context"

	"github.com/google/uuid"
)

type StorageClient interface {
	GetTransaction(ctx context.Context, userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(ctx context.Context, userId uuid.UUID, begin string, end string) ([]*models.Transaction, error)
	PostTransaction(ctx context.Context, tx *models.Transaction, withAutoCategory bool) (string, error)
}
