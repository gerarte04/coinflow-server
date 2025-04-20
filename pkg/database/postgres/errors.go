package postgres

import (
	"coinflow/coinflow-server/pkg/database"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	PgUniqueViolation = "23505"
)

var (
	codeErrors = map[string]error {
		PgUniqueViolation: database.ErrorUniqueViolation,
	}
)

func DetectError(err error) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return database.ErrorUndocumented
	}

	if dbErr, ok := codeErrors[pgErr.Code]; ok {
		return dbErr
	}

	return database.ErrorUndocumented
}
