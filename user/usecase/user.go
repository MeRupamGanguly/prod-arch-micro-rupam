package usecase

import (
	"rupx/user/contracts"
	"rupx/user/domain"
)

type service struct {
	db contracts.RepoContracts
}

func NewUserSvc(d contracts.RepoContracts) *service {
	return &service{db: d}
}

func (svc *service) Add(name string, email string) error {
	err := svc.db.CreateUser(name, email)
	if err != nil {
		return err
	}
	return nil
}
func (svc *service) Get(id uint) (user domain.User, err error) {
	user, err = svc.db.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (svc *service) List() (users []domain.User, err error) {
	users, err = svc.db.ListUser()
	if err != nil {
		return users, err
	}
	return users, nil
}
