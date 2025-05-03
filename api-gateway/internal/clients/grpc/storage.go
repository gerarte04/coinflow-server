package grpc

import (
	"coinflow/coinflow-server/api-gateway/internal/models"
)

type StorageClient interface {
	GetTransactionsInPeriod(begin string, end string) ([]*models.Transaction, error)
	PostTransaction(tx *models.Transaction) (string, error)
}
