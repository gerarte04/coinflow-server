package repository

import (
	"github.com/google/uuid"
)

type TransactionsRepo interface {
	PutCategory(tsId uuid.UUID, category string) error
}
