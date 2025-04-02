package transactions

import (
	"coinflow/coinflow-server/internal/models"
	"coinflow/coinflow-server/internal/repository"

	"github.com/google/uuid"
)

type TransactionsService struct {
    tsRepo repository.TransactionsRepo
}

func NewTransactionsService(tsRepo repository.TransactionsRepo) *TransactionsService {
    return &TransactionsService{tsRepo: tsRepo}
}

func (s *TransactionsService) GetTransaction(tsId string) (*models.Transaction, error) {
    return s.tsRepo.GetTransaction(tsId)
}

func (s *TransactionsService) PostTransaction(ts *models.Transaction) (*models.Transaction, error) {
    id := uuid.New()
    ts.Id = id.String()

    if err := s.tsRepo.PostTransaction(ts); err != nil {
        return nil, err
    }

    return ts, nil
}
