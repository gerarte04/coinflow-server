package repository

import (
	"coinflow/coinflow-server/storage-service/internal/models"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source transactions.go -destination mocks/transactions_mock.go -package mocks
type TransactionsRepo interface {
	GetTransaction(userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(userId uuid.UUID, begin time.Time, end time.Time) ([]*models.Transaction, error)
	PostTransaction(tx *models.Transaction) (uuid.UUID, error)
	PostTransactionWithoutCategory(tx *models.Transaction) (uuid.UUID, error)
	PutCategory(txId uuid.UUID, category string) error
}
