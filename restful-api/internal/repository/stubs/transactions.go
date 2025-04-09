package stubs

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TransactionsRepoMock struct {
    mp map[uuid.UUID]*models.Transaction
}

func NewTransactionsRepoMock() *TransactionsRepoMock {
    return &TransactionsRepoMock{mp: make(map[uuid.UUID]*models.Transaction)}
}

func (r *TransactionsRepoMock) GetTransaction(tsId uuid.UUID) (*models.Transaction, error) {
    ts, ok := r.mp[tsId]

    if !ok {
        return nil, fmt.Errorf("repo: getting transaction: %w", repository.ErrorTransactionKeyNotFound)
    }

    return ts, nil
}

func (r *TransactionsRepoMock) GetUserTransactionsAfterTimestamp(usrId string, tm time.Time) ([]*models.Transaction, error) {
    tss := make([]*models.Transaction, 0)

    for _, v := range r.mp {
        if v.UserId == usrId && v.Timestamp.After(tm) {
            ts := v
            tss = append(tss, ts)
        }
    }

    return tss, nil
}

func (r *TransactionsRepoMock) PostTransaction(ts *models.Transaction) (uuid.UUID, error) {
    id := uuid.New()

    if _, ok := r.mp[id]; ok {
        return uuid.Nil, fmt.Errorf("repo: posting transaction: %w", repository.ErrorTransactionKeyExists)
    }

    tsCopy := *ts
    tsCopy.Id = id
    tsCopy.Timestamp = time.Now()
    r.mp[id] = &tsCopy

    return id, nil
}
