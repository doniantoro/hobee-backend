package repository

import (
	"hobee/config"
	"hobee/model"
)

var db = config.ConnectionDatabase()

type Repository interface {
	Login(email string) (model.User, error)
	Register(model.User) (model.User, error)
}
type repository struct {
}

func AuthRepository() Repository {
	return &repository{}
}

func (s repository) Login(email string) (model.User, error) {
	var user model.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s repository) Register(user model.User) (model.User, error) {
	err := db.Create(&user).Error

	return user, err
}
