package repository

import (
	"coinflow/coinflow-server/storage-service/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source transactions.go -destination mocks/transactions_mock.go -package mocks
type TransactionsRepo interface {
	GetTransaction(ctx context.Context, userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(ctx context.Context, userId uuid.UUID, begin time.Time, end time.Time, limit int) ([]*models.Transaction, error)
	PostTransaction(ctx context.Context, tx *models.Transaction) (*models.Transaction, error)
	PostTransactionWithoutCategory(ctx context.Context, tx *models.Transaction) (*models.Transaction, error)
	PutCategory(ctx context.Context, txId uuid.UUID, category string) error
}
