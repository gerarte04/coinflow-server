package repository

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source transactions.go -destination mocks/transactions_mock.go -package mocks
type TransactionsRepo interface {
    GetTransaction(tsId uuid.UUID) (*models.Transaction, error)
    GetUserTransactionsAfterTimestamp(usrId string, tm time.Time) ([]*models.Transaction, error)
    PostTransaction(ts *models.Transaction) (uuid.UUID, error)
}
