package repository

import (
	"coinflow/coinflow-server/auth-service/internal/models"
	"context"

	"github.com/google/uuid"
)

type UsersRepo interface {
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByCred(ctx context.Context, login, password string) (*models.User, error)
	PostUser(ctx context.Context, usr *models.User) (uuid.UUID, error)
}
