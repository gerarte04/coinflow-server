package repository

import "coinflow/coinflow-server/restful-api/internal/models"

//go:generate mockgen -source users.go -destination mocks/users_mock.go -package mocks
type UsersRepo interface {
	GetUser(usrId string) (*models.User, error)
	GetUserByCred(login string, password string) (*models.User, error)
	PostUser(usr *models.User) error
}
