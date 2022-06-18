package controller

import (
	"github.com/gin-gonic/gin"
	"hobee/helper"
	"net/http"
)

func GetAllCategory(c *gin.Context) {

	response := helper.Response(200, "Succes", "data berhasil masuk", "")
	c.JSON(http.StatusOK, response)
}
