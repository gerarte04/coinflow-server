package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepo struct {
	pool *pgxpool.Pool
}

func NewCategoriesRepo(pool *pgxpool.Pool) *CategoriesRepo {
	return &CategoriesRepo{pool: pool}
}

func (r *CategoriesRepo) GetCategories() ([]string, error) {
	const op = "CategoriesRepo.GetCategories"

	rows, err := r.pool.Query(context.Background(), "SELECT * FROM categories")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	cats := make([]string, 0)

	for rows.Next() {
		var cat string

		if err := rows.Scan(&cat); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		cats = append(cats, cat)
	}

	return cats, nil
}
