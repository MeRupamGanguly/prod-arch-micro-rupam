package contracts

import "rupx/user/domain"

type RepoContracts interface {
	CreateUser(name string, email string) (err error)
	GetUser(id uint) (user domain.User, err error)
	ListUser() (users []domain.User, err error)
}
