package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hobee/service"
)

func GetAllCategory(c *gin.Context) {

	var categoryService = service.NewCategoryService()
	category, err := categoryService.GetAllCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": category})
}
