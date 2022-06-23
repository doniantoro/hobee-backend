package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hobee/helper"
	"log"
	"net/http"
	"os"
)

func GetAllCategory(c *gin.Context) {

	//var response = helper.Response(200, "Succes", "data berhasil masuk", c.GetHeader("Authorization"))

	//c.JSON(http.StatusOK, response)
	//return
	var getToken = c.GetHeader("Authorization")
	c.JSON(http.StatusOK, getToken)
	return
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SECRET_KEY := []byte(os.Getenv("SECRET_KEY"))

	parse, err := jwt.Parse(getToken, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
	}
	if parse.Valid {
		fmt.Println("Valid")
		fmt.Println("Valid")
		fmt.Println("Valid")
		fmt.Println("Valid")
	}
	var response = helper.Response(200, "Succes", "data berhasil masuk", parse)

	c.JSON(http.StatusOK, response)

	//return parse, err
	//parse, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return SECRET_KEY, nil
	//})
	//if claims, ok := parse.Claims.(jwt.MapClaims); ok && parse.Valid {
	//	fmt.Println(claims["foo"], claims["nbf"])
	//} else {
	//	fmt.Println(err)
	//}

	//response := helper.Response(200, "Succes", "data berhasil masuk", "")
	//c.JSON(http.StatusOK, response)
}
