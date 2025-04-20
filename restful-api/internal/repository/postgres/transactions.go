package postgres

import (
	"coinflow/coinflow-server/pkg/database"
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
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

func (r *TransactionsRepo) GetTransaction(tsId uuid.UUID) (*models.Transaction, error) {
	const method = "TransactionsRepo.GetTransaction"

	row := r.conn.QueryRow(
		context.Background(),
		"SELECT * FROM transactions WHERE id = $1", tsId,
	)

	var ts models.Transaction
	err := row.Scan(&ts.Id, &ts.UserId, &ts.Type, &ts.Target, &ts.Description, &ts.Category, &ts.Cost, &ts.Timestamp)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("%s: %w", method, repository.ErrorTransactionKeyNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return &ts, nil
}

func (r *TransactionsRepo) PostTransaction(ts *models.Transaction) (uuid.UUID, error) {
	const method = "TransactionsRepo.PostTransaction"

	tsId := uuid.New()
	var err error

	if ts.WithAutoCategory {
		_, err = r.conn.Exec(
			context.Background(),
			`INSERT INTO transactions (
				id, user_id, type, target, description, cost
			) VALUES ($1, $2, $3, $4, $5, $6)`,
			tsId, ts.UserId, ts.Type, ts.Target, ts.Description, ts.Cost,
		)
	} else {
		_, err = r.conn.Exec(
			context.Background(),
			`INSERT INTO transactions (
				id, user_id, type, target, description, category, cost
			) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			tsId, ts.UserId, ts.Type, ts.Target, ts.Description, ts.Category, ts.Cost,
		)
	}

	if dbErr := postgres.DetectError(err); dbErr == database.ErrorUniqueViolation {
		return uuid.Nil, fmt.Errorf("%s: %w", method, repository.ErrorTransactionKeyExists)
	} else if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", method, err)
	}

	return tsId, nil
}
