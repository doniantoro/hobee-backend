package controller

import (
	"github.com/gin-gonic/gin"
	request "hobee/Request"
	"hobee/helper"
	"hobee/service"
	"net/http"
)

func Login(c *gin.Context) {
	var UserInput request.UserInput

	err := c.ShouldBindJSON(&UserInput)

	if err != nil {
		err := helper.Response(http.StatusBadGateway, "Failed", "There is required failed", err.Error())

		c.JSON(http.StatusBadGateway, err)
		return
	}
	var authservice = service.AuthService()

	responseLogin, err := authservice.Login(UserInput)
	if err != nil {

		response := helper.Response(http.StatusBadGateway, "Failed", err.Error(), "")

		c.JSON(http.StatusBadGateway, response)
		return
	}
	response := helper.Response(200, "Succes", "Login Succes", responseLogin)
	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context) {
	var UserInput request.UserInput
	err := c.ShouldBindJSON(&UserInput)

	if err != nil {
		err := helper.Response(http.StatusBadGateway, "Failed", "There is required failed", err.Error())
		c.JSON(502, err)
		return
	}

	var register = service.AuthService()
	register.Register(UserInput)

	response := helper.Response(200, "Success", "Register Succes", "")

	c.JSON(200, response)

}
