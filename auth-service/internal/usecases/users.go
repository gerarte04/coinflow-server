package usecases

import (
	"coinflow/coinflow-server/auth-service/internal/models"

	"github.com/google/uuid"
)

type TokenPair struct {
	Access string
	Refresh string
}

type UserService interface {
	Login(login, password string) (*TokenPair, error)
	Refresh(refreshToken string) (*TokenPair, error)
	Register(usr *models.User) (uuid.UUID, error)
	GetUserData(usrId uuid.UUID) (*models.User, error)
}
