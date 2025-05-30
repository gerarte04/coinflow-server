package grpc

import (
	"coinflow/coinflow-server/api-gateway/internal/models"
	"context"

	"github.com/google/uuid"
)

type AuthClient interface {
	Login(ctx context.Context, login, password string) (string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)
	Register(ctx context.Context, usr *models.User) (uuid.UUID, error)
	GetUserData(ctx context.Context, usrId uuid.UUID) (*models.User, error)
}
