package usecases

import (
	"coinflow/coinflow-server/storage-service/internal/models"
	"time"

	"github.com/google/uuid"
)

type TransactionsService interface {
	GetTransaction(txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(begin time.Time, end time.Time) ([]*models.Transaction, error)
	PostTransaction(tx *models.Transaction) (uuid.UUID, error)
}
