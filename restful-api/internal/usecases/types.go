package usecases

import "coinflow/coinflow-server/restful-api/internal/models"

type TransactionsService interface {
    GetTransaction(tsId string) (*models.Transaction, error)
    PostTransaction(ts *models.Transaction) (string, error)
}
