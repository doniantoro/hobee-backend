package service

import (
	"github.com/dgrijalva/jwt-go"
	"hobee/model"
)

var SECRET_KEY = []byte("DoNiAn1234")

type serviceInterface interface {
	GenerateToken(user model.User) (string, error)
}

type jwtService struct {
}

func JwtService() serviceInterface {
	return &jwtService{}
}

func (j jwtService) GenerateToken(user model.User) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = user.Id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, err

}
