package usecases

import (
	"coinflow/coinflow-server/storage-service/internal/models"
	"time"

	"github.com/google/uuid"
)

type TransactionsService interface {
	GetTransaction(userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error)
	GetTransactionsInPeriod(userId uuid.UUID, begin time.Time, end time.Time) ([]*models.Transaction, error)
	PostTransaction(tx *models.Transaction, withAutoCategory bool) (uuid.UUID, error)
}
