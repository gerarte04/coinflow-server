package repository

import "context"

type CategoriesRepo interface{
	GetCategories(ctx context.Context) ([]string, error)
}
