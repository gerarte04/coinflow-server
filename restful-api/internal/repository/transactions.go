package repository

import (
	"coinflow/coinflow-server/restful-api/internal/models"

	"github.com/google/uuid"
)

//go:generate mockgen -source transactions.go -destination mocks/transactions_mock.go -package mocks
type TransactionsRepo interface {
	GetTransaction(txId uuid.UUID) (*models.Transaction, error)
	PostTransaction(tx *models.Transaction) (uuid.UUID, error)
	PostTransactionWithoutCategory(tx *models.Transaction) (uuid.UUID, error)
	PutCategory(txId uuid.UUID, category string) error
}
