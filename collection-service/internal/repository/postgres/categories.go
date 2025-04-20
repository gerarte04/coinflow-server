package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CategoriesRepo struct {
	conn *pgx.Conn
}

func NewCategoriesRepo(conn *pgx.Conn) *CategoriesRepo {
	return &CategoriesRepo{conn: conn}
}

func (r *CategoriesRepo) GetCategories() ([]string, error) {
	const method = "CategoriesRepo.GetCategories"

	rows, err := r.conn.Query(context.Background(), "SELECT * FROM categories")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	defer rows.Close()

	cats := make([]string, 0)

	for rows.Next() {
		var cat string

		if err := rows.Scan(&cat); err != nil {
			return nil, fmt.Errorf("%s: %w", method, err)
		}

		cats = append(cats, cat)
	}

	return cats, nil
}
