package repository

import (
	"rupx/user/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepo(d *gorm.DB) *repository {
	return &repository{
		db: d,
	}
}

func (r *repository) CreateUser(name string, email string) (err error) {
	err = r.db.Create(domain.User{Name: name, EMail: email}).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) GetUser(id uint) (user domain.User, err error) {
	err = r.db.First(user, id).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (r *repository) ListUser() (users []domain.User, err error) {
	err = r.db.Find(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}
