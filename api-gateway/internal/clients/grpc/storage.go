package grpc

import (
	"coinflow/coinflow-server/api-gateway/internal/models"

	"github.com/google/uuid"
)

type StorageClient interface {
	GetTransaction(userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(userId uuid.UUID, begin string, end string) ([]*models.Transaction, error)
	PostTransaction(tx *models.Transaction, withAutoCategory bool) (string, error)
}
