package repository

import (
	"coinflow/coinflow-server/internal/models"
)

type TransactionsRepo interface {
    GetTransaction(tsId string) (*models.Transaction, error)
    PostTransaction(ts *models.Transaction) error
}
