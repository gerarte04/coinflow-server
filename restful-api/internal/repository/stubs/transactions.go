package stubs

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TransactionsRepoStub struct {
	mp map[uuid.UUID]*models.Transaction
}

func NewTransactionsRepoStub() *TransactionsRepoStub {
	return &TransactionsRepoStub{mp: make(map[uuid.UUID]*models.Transaction)}
}

func (r *TransactionsRepoStub) GetTransaction(tsId uuid.UUID) (*models.Transaction, error) {
	const method = "TransactionsRepoStub.GetTransaction"

	ts, ok := r.mp[tsId]

	if !ok {
		return nil, fmt.Errorf("%s: %w", method, repository.ErrorTxIdNotFound)
	}

	return ts, nil
}

func (r *TransactionsRepoStub) GetUserTransactionsAfterTimestamp(usrId uuid.UUID, tm time.Time) ([]*models.Transaction, error) {
	tss := make([]*models.Transaction, 0)

	for _, v := range r.mp {
		if v.UserId == usrId && v.Timestamp.After(tm) {
			ts := v
			tss = append(tss, ts)
		}
	}

	return tss, nil
}

func (r *TransactionsRepoStub) PostTransaction(ts *models.Transaction) (uuid.UUID, error) {
	const method = "TransactionsRepoStub.PostTransaction"

	id := uuid.New()

	if _, ok := r.mp[id]; ok {
		return uuid.Nil, fmt.Errorf("%s: %w", method, repository.ErrorTxIdAlreadyExists)
	}

	tsCopy := *ts
	tsCopy.Id = id
	tsCopy.Timestamp = time.Now()
	r.mp[id] = &tsCopy

	return id, nil
}
