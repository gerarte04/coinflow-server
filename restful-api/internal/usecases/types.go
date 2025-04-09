package usecases

import (
	"coinflow/coinflow-server/restful-api/internal/models"

	"github.com/google/uuid"
)

type TransactionsService interface {
    GetTransaction(tsId uuid.UUID) (*models.Transaction, error)
    PostTransaction(ts *models.Transaction) (uuid.UUID, error)
}
