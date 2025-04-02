package transactions

import (
	"coinflow/coinflow-server/internal/models"
	"coinflow/coinflow-server/internal/repository"
)

type TransactionsRepoMock struct {
    mp map[string]models.Transaction
}

func NewTransactionsRepoMock() *TransactionsRepoMock {
    return &TransactionsRepoMock{mp: make(map[string]models.Transaction)}
}

func (r *TransactionsRepoMock) GetTransaction(tsId string) (*models.Transaction, error) {
    ts, ok := r.mp[tsId]

    if !ok {
        return nil, repository.ErrorTransactionKeyNotFound
    }

    return &ts, nil
}

func (r *TransactionsRepoMock) PostTransaction(ts *models.Transaction) error {
    if _, ok := r.mp[ts.Id]; ok {
        return repository.ErrorTransactionKeyExists
    }

    r.mp[ts.Id] = *ts

    return nil
}
