package usecases

import (
	"coinflow/coinflow-server/restful-api/internal/models"

	"github.com/google/uuid"
)

type TransactionsService interface {
	GetTransaction(txId uuid.UUID) (*models.Transaction, error)
	PostTransaction(tx *models.Transaction) (uuid.UUID, error)
}
