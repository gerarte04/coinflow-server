package grpc

import (
	"coinflow/coinflow-server/api-gateway/internal/models"

	"github.com/google/uuid"
)

type AuthClient interface {
	Login(login, password string) (string, string, error)
	Refresh(refreshToken string) (string, string, error)
	Register(usr *models.User) (uuid.UUID, error)
	GetUserData(usrId uuid.UUID) (*models.User, error)
}
