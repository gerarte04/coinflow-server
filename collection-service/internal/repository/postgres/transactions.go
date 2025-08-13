package postgres

import (
	"coinflow/coinflow-server/collection-service/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type TransactionRepo struct {
	conn driver.Conn
}

func NewTransactionRepo(conn driver.Conn) *TransactionRepo {
	return &TransactionRepo{conn: conn}
}

func (r *TransactionRepo) GetSummaryInPeriod(
	ctx context.Context,
	userId uuid.UUID,
	begin time.Time, end time.Time,
) (*models.Summary, error) {
	const op = "TransactionRepo.GetSummaryInPeriod"

	row := r.conn.QueryRow(
		ctx,
		`SELECT count(cost), sum(cost), avg(cost)
		FROM transactions FINAL
		WHERE user_id = $1 AND timestamp >= $2 AND timestamp <= $3`,
		userId, begin, end,
	)

	var sm models.Summary

	if err := row.Scan(&sm.Count, &sm.Sum, &sm.Avg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sm, nil
}

func (r *TransactionRepo) GetSummaryInLastNMonths(
	ctx context.Context,
	userId uuid.UUID,
	n int,
	curTime time.Time,
	timezone string,
) ([]*models.Summary, error) {
	const op = "TransactionRepo.GetSummaryInLastNMonths"

	rows, err := r.conn.Query(
		ctx,
		`SELECT dateTrunc('month', timestamp, $1) AS month,
		count(cost), sum(cost), avg(cost)
		FROM transactions FINAL
		WHERE user_id = $2
		AND timestamp >= dateTrunc('month', addMonths($3, $4), $1)
		AND timestamp <= $3
		GROUP BY month
		ORDER BY month DESC`,
		timezone, userId, curTime, 1 - n,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sm := []*models.Summary{}

	for rows.Next() {
		var month time.Time
		var s models.Summary

		if err := rows.Scan(&month, &s.Count, &s.Sum, &s.Avg); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		sm = append(sm, &s)
	}

	return sm, nil
}

func (r *TransactionRepo) GetSummaryByCategories(
	ctx context.Context,
	userId uuid.UUID,
	begin time.Time, end time.Time,
) (map[string]*models.Summary, error) {
	const op = "TransactionRepo.GetSummaryByCategories"

	rows, err := r.conn.Query(
		ctx,
		`SELECT category, count(cost), sum(cost), avg(cost)
		FROM transactions FINAL
		WHERE user_id = $1 AND timestamp >= $2 AND timestamp <= $3
		GROUP BY category
		ORDER BY sum(cost) DESC`,
		userId, begin, end,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sm := map[string]*models.Summary{}

	for rows.Next() {
		var category string
		var s models.Summary

		if err := rows.Scan(&category, &s.Count, &s.Sum, &s.Avg); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		sm[category] = &s
	}

	return sm, nil
}
