package main

import (
	"github.com/gin-gonic/gin"
	"hobee/controller"
	"hobee/config"
	"hobee/seeder"
)

func main() {
	config.Migration()
	seeder.Category()
	router := gin.Default()

	// router.POST("v1/get_all_category", controller.GetAllCategory)
	router.POST("v1/register", controller.Register)
	router.POST("v1/login", controller.Login)
	router.GET("v1/get_all_category", controller.GetAllCategory)
	router.Static("/assets/icon", "./assets/icons")
	router.Run(":10")

}
 