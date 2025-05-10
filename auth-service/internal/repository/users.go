package repository

import (
	"coinflow/coinflow-server/auth-service/internal/models"

	"github.com/google/uuid"
)

type UsersRepo interface {
	GetUser(id uuid.UUID) (*models.User, error)
	GetUserByCred(login, password string) (*models.User, error)
	PostUser(usr *models.User) (uuid.UUID, error)
}
