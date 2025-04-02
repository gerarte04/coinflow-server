package usecases

import "coinflow/coinflow-server/internal/models"

type TransactionsService interface {
    GetTransaction(tsId string) (*models.Transaction, error)
    PostTransaction(ts *models.Transaction) (*models.Transaction, error)
}
