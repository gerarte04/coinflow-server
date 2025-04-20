package postgres

import (
	"coinflow/coinflow-server/collection-service/internal/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TransactionsRepo struct {
	conn *pgx.Conn
}

func NewTransactionsRepo(conn *pgx.Conn) *TransactionsRepo {
	return &TransactionsRepo{conn: conn}
}

func (r *TransactionsRepo) PutCategory(tsId uuid.UUID, category string) error {
	const method = "TransactionsRepo.PutCategory"

	tag, err := r.conn.Exec(context.Background(),
		"UPDATE transactions SET category = $1 WHERE id = $2",
		category, tsId,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	} else if tag.RowsAffected() != 1 {
		return fmt.Errorf("%s: %w", method, repository.ErrorNoTsIdExists)
	}

	return nil
}
