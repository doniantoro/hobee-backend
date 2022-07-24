package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hobee/service"
	"hobee/helper"
)

func GetAllCategory(c *gin.Context) {

	var categoryService = service.NewCategoryService()
	category, err := categoryService.GetAllCategory()
	if err != nil {
		
		response := helper.Response(http.StatusInternalServerError , "Failed", err.Error(), "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	
	response := helper.Response(200, "Succes", "Succes Get All Data", category)
	
	c.JSON(http.StatusOK, gin.H{"data": response})
}
