package postgres

import (
	"coinflow/coinflow-server/pkg/database"
	"coinflow/coinflow-server/pkg/database/postgres"
	"coinflow/coinflow-server/storage-service/internal/models"
	"coinflow/coinflow-server/storage-service/internal/repository"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TransactionsRepo struct {
	conn *pgx.Conn
}

func NewTransactionsRepo(conn *pgx.Conn) *TransactionsRepo {
	return &TransactionsRepo{conn: conn}
}

func (r *TransactionsRepo) GetTransaction(userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error) {
	const op = "TransactionsRepo.GetTransaction"

	row := r.conn.QueryRow(
		context.Background(),
		"SELECT * FROM transactions WHERE id = $1", txId,
	)

	var tx models.Transaction
	err := row.Scan(&tx.Id, &tx.UserId, &tx.Type, &tx.Target, &tx.Description, &tx.Category, &tx.Cost, &tx.Timestamp)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrorTxIdNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if tx.UserId != userId {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrorPermissionDenied)
	}

	return &tx, nil
}

func (r *TransactionsRepo) GetTransactionsInPeriod(userId uuid.UUID, begin time.Time, end time.Time) ([]*models.Transaction, error) {
	const op = "TransactionsRepo.GetTransactionsInPeriod"

	rows, err := r.conn.Query(
		context.Background(),
		"SELECT * FROM transactions WHERE (user_id = $1 AND timestamp >= $2 AND timestamp <= $3)",
		userId, begin, end,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	txs := make([]*models.Transaction, 0)

	for rows.Next() {
		var tx models.Transaction
		err := rows.Scan(&tx.Id, &tx.UserId, &tx.Type, &tx.Target, &tx.Description, &tx.Category, &tx.Cost, &tx.Timestamp)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		txs = append(txs, &tx)
	}

	return txs, nil
}

func (r *TransactionsRepo) PostTransaction(tx *models.Transaction) (uuid.UUID, error) {
	const op = "TransactionsRepo.PostTransaction"

	row := r.conn.QueryRow(
		context.Background(),
		`INSERT INTO transactions (
			user_id, type, target, description, category, cost
		) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		tx.UserId, tx.Type, tx.Target, tx.Description, tx.Category, tx.Cost,
	)

	var txId uuid.UUID
	err := row.Scan(&txId)

	if dbErr := postgres.DetectError(err); dbErr == database.ErrorUniqueViolation {
		return uuid.Nil, fmt.Errorf("%s: %w", op, repository.ErrorTxIdAlreadyExists)
	} else if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return txId, nil
}

func (r *TransactionsRepo) PostTransactionWithoutCategory(tx *models.Transaction) (uuid.UUID, error) {
	const op = "TransactionsRepo.PostTransaction"

	row := r.conn.QueryRow(
		context.Background(),
		`INSERT INTO transactions (
			user_id, type, target, description, cost
		) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		tx.UserId, tx.Type, tx.Target, tx.Description, tx.Cost,
	)

	var txId uuid.UUID
	err := row.Scan(&txId)

	if dbErr := postgres.DetectError(err); dbErr == database.ErrorUniqueViolation {
		return uuid.Nil, fmt.Errorf("%s: %w", op, repository.ErrorTxIdAlreadyExists)
	} else if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return txId, nil
}

func (r *TransactionsRepo) PutCategory(tsId uuid.UUID, category string) error {
	const op = "TransactionsRepo.PutCategory"

	tag, err := r.conn.Exec(context.Background(),
		"UPDATE transactions SET category = $1 WHERE id = $2",
		category, tsId,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	} else if tag.RowsAffected() != 1 {
		return fmt.Errorf("%s: %w", op, repository.ErrorTxIdNotFound)
	}

	return nil
}
