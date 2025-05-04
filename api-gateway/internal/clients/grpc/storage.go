package grpc

import (
	"coinflow/coinflow-server/api-gateway/internal/models"

	"github.com/google/uuid"
)

type StorageClient interface {
	GetTransaction(txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(begin string, end string) ([]*models.Transaction, error)
	PostTransaction(tx *models.Transaction) (string, error)
}
