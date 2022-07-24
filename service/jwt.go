package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"hobee/model"
	"log"
	"os"
)

type serviceInterface interface {
	GenerateToken(user model.User) (string, error)
}

type jwtService struct {
}

func JwtService() serviceInterface {
	return &jwtService{}
}

func (j jwtService) GenerateToken(user model.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SECRET_KEY := []byte(os.Getenv("SECRET_KEY"))

	claim := jwt.MapClaims{}
	claim["user_id"] = user.Id
	claim["name"] = user.Name
	claim["email"] = user.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, err

}
