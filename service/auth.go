package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	request "hobee/Request"
	"hobee/model"
	"hobee/repository"
)

type Service interface {
	Register(user request.UserInput) (model.User, error)
	Login(input request.UserInput) (string, error)
}

type service struct {
}

func AuthService() Service {
	return &service{}
}

func (s service) Login(input request.UserInput) (string, error) {

	var userRepo = repository.AuthRepository()
	getUser, _ := userRepo.Login(input.Email)
	if getUser.Id == 0 {

		return "failed", errors.New("No User Found")
	}

	var err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(input.Password))

	if err != nil {

		return "failed", err
	}
	var jwtService = JwtService()
	token, err := jwtService.GenerateToken(getUser)
	return token, nil

}

func (s service) Register(input request.UserInput) (model.User, error) {

	user := model.User{}
	user.Email = input.Email
	user.Password = getHash([]byte(input.Password))

	var register = repository.AuthRepository()
	user, err := register.Register(user)
	if err != nil {
		return user, err
	}

	return user, err
}

func getHash(password []byte) string {
	password, _ = bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	return string(password)
}
