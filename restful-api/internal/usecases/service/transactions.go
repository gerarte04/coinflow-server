package service

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"

	"github.com/google/uuid"
)

type TransactionsService struct {
	tsRepo repository.TransactionsRepo
}

func NewTransactionsService(tsRepo repository.TransactionsRepo) *TransactionsService {
	return &TransactionsService{tsRepo: tsRepo}
}

func (s *TransactionsService) GetTransaction(tsId uuid.UUID) (*models.Transaction, error) {
	return s.tsRepo.GetTransaction(tsId)
}

func (s *TransactionsService) PostTransaction(ts *models.Transaction) (uuid.UUID, error) {
	id, err := s.tsRepo.PostTransaction(ts)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
