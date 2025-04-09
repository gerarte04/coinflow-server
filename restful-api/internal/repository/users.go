package repository

import "coinflow/coinflow-server/restful-api/internal/models"

type UsersRepo interface {
    GetUser(usrId string) (*models.User, error)
    GetUserByCred(login string, password string) (*models.User, error)
    PostUser(usr *models.User) error
}
