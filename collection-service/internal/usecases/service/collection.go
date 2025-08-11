package service

import (
	"coinflow/coinflow-server/collection-service/internal/repository"
	"context"
	"fmt"
	"time"
)

type CollectionService struct {
	catsRepo repository.CategoriesRepo
	categories []string
}

func NewCollectionService(
	catsRepo repository.CategoriesRepo,
) (*CollectionService, error) {
	const op = "NewCollectionService"

	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Millisecond)
	defer cancel()

	categories, err := catsRepo.GetCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &CollectionService{
		catsRepo: catsRepo,
		categories: categories,
	}, nil
}
