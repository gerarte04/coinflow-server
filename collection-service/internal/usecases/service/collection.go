package service

import (
	"coinflow/coinflow-server/collection-service/internal/models"
	"coinflow/coinflow-server/collection-service/internal/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type CollectionService struct {
	txRepo repository.TransactionRepo
}

func NewCollectionService(
	txRepo repository.TransactionRepo,
) (*CollectionService, error) {
	const op = "NewCollectionService"

	return &CollectionService{
		txRepo: txRepo,
	}, nil
}

func (s *CollectionService) GetSummaryInPeriod(
	ctx context.Context,
	userId uuid.UUID,
	begin time.Time, end time.Time,
) (*models.Summary, error) {
	return s.txRepo.GetSummaryInPeriod(ctx, userId, begin, end)
}

func (s *CollectionService) GetSummaryInLastNMonths(
	ctx context.Context,
	userId uuid.UUID,
	n int,
	curTime time.Time, timezone string,
) ([]*models.Summary, error) {
	return s.txRepo.GetSummaryInLastNMonths(ctx, userId, n, curTime, timezone)
}

func (s *CollectionService) GetSummaryByCategories(
	ctx context.Context,
	userId uuid.UUID,
	begin time.Time, end time.Time,
) (map[string]*models.Summary, error) {
	return s.txRepo.GetSummaryByCategories(ctx, userId, begin, end)
}
