package postgres

import (
	"github.com/jackc/pgx/v5"
)

type TransactionsRepo struct {
	conn *pgx.Conn
}

func NewTransactionsRepo(conn *pgx.Conn) *TransactionsRepo {
	return &TransactionsRepo{conn: conn}
}
