package main

import (
	"github.com/gin-gonic/gin"
	"hobee/config"
	"hobee/controller"
)

func main() {

	config.ConnectionDatabase()
	router := gin.Default()

	router.POST("v1/register", controller.Register)
	router.POST("v1/login", controller.Login)
	router.GET("v1/get_all_category", controller.GetAllCategory)
	router.Run(":10")

}
