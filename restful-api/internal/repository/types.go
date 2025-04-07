package repository

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"time"
)

type TransactionsRepo interface {
    GetTransaction(tsId string) (*models.Transaction, error)
    GetUserTransactionsAfterTimestamp(usrId string, tm time.Time) ([]*models.Transaction, error)
    PostTransaction(ts *models.Transaction) error
}

type UsersRepo interface {
    GetUser(usrId string) (*models.User, error)
    GetUserByCred(login string, password string) (*models.User, error)
    PostUser(usr *models.User) error
}
