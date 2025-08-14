package repository

import (
	"coinflow/coinflow-server/collection-service/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
)

type TransactionRepo interface {
	GetSummaryInPeriod(
		ctx context.Context,
		userId uuid.UUID,
		begin time.Time, end time.Time,
	) (*models.Summary, error)
	
	GetSummaryInLastNMonths(
		ctx context.Context,
		userId uuid.UUID,
		n int,
		curTime time.Time, timezone string,
	) ([]*models.Summary, error)

	GetSummaryByCategories(
		ctx context.Context,
		userId uuid.UUID,
		begin time.Time, end time.Time,
	) (map[string]*models.Summary, error)
}
