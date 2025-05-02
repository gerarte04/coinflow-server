package repository

import (
	"github.com/google/uuid"
)

type TransactionsRepo interface {
	PutCategory(txId uuid.UUID, category string) error
}
