package mocks

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"fmt"
)

type UsersRepoMock struct {
    mp map[string]models.User
}

func NewUsersRepoMock() *UsersRepoMock {
    return &UsersRepoMock{mp: make(map[string]models.User)}
}

func (r *UsersRepoMock) GetUser(usrId string) (*models.User, error) {
    usr, ok := r.mp[usrId]

    if !ok {
        return nil, fmt.Errorf("repo: getting user: %w", repository.ErrorUserKeyNotFound)
    }

    return &usr, nil
}

func (r *UsersRepoMock) GetUserByCred(login string, password string) (*models.User, error) {
    for _, v := range r.mp {
        if v.Login == login && v.Password == password {
            usr := v
            return &usr, nil
        }
    }

    return nil, fmt.Errorf("repo: getting user by cred: %w", repository.ErrorNoSuchCredExists)
}

func (r *UsersRepoMock) PostUser(usr *models.User) error {
    if _, ok := r.mp[usr.Id]; ok {
        return fmt.Errorf("repo: posting user: %w", repository.ErrorUserKeyExists)
    }

    r.mp[usr.Id] = *usr

    return nil
}
