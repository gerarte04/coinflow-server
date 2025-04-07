package mocks

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"fmt"
	"time"
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
        return nil, fmt.Errorf("repo: getting transaction: %w", repository.ErrorTransactionKeyNotFound)
    }

    return &ts, nil
}

func (r *TransactionsRepoMock) GetUserTransactionsAfterTimestamp(usrId string, tm time.Time) ([]*models.Transaction, error) {
    tss := make([]*models.Transaction, 0)

    for _, v := range r.mp {
        if v.UserId == usrId && v.Timestamp.After(tm) {
            ts := v
            tss = append(tss, &ts)
        }
    }

    return tss, nil
}

func (r *TransactionsRepoMock) PostTransaction(ts *models.Transaction) error {
    if _, ok := r.mp[ts.Id]; ok {
        return fmt.Errorf("repo: posting transaction: %w", repository.ErrorTransactionKeyExists)
    }

    ts.Timestamp = time.Now()
    r.mp[ts.Id] = *ts

    return nil
}
