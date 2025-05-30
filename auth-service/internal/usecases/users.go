package usecases

import (
	"coinflow/coinflow-server/auth-service/internal/models"
	"context"

	"github.com/google/uuid"
)

type TokenPair struct {
	Access string
	Refresh string
}

type UserService interface {
	Login(ctx context.Context, login, password string) (*TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (*TokenPair, error)
	Register(ctx context.Context, usr *models.User) (uuid.UUID, error)
	GetUserData(ctx context.Context, usrId uuid.UUID) (*models.User, error)
}
