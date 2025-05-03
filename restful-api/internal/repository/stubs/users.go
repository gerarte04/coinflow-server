package stubs

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"fmt"
)

type UsersRepoStub struct {
	mp map[string]models.User
}

func NewUsersRepoStub() *UsersRepoStub {
	return &UsersRepoStub{mp: make(map[string]models.User)}
}

func (r *UsersRepoStub) GetUser(usrId string) (*models.User, error) {
	const method = "UsersRepoStub.GetUser"

	usr, ok := r.mp[usrId]

	if !ok {
		return nil, fmt.Errorf("%s: %w", method, repository.ErrorUserIdNotFound)
	}

	return &usr, nil
}

func (r *UsersRepoStub) GetUserByCred(login string, password string) (*models.User, error) {
	const method = "UsersRepoStub.GetUserByCred"

	for _, v := range r.mp {
		if v.Login == login && v.Password == password {
			usr := v
			return &usr, nil
		}
	}

	return nil, fmt.Errorf("%s: %w", method, repository.ErrorNoSuchCredExists)
}

func (r *UsersRepoStub) PostUser(usr *models.User) error {
	const method = "UsersRepoStub.PostUser"

	if _, ok := r.mp[usr.Id]; ok {
		return fmt.Errorf("%s: %w", method, repository.ErrorUserIdAlreadyExists)
	}

	r.mp[usr.Id] = *usr

	return nil
}
