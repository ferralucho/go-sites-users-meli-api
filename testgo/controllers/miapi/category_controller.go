package miapi

import (
	"../../services/miapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	paramCategoryID = "categoryID"
)

func GetCategory(c *gin.Context){
	categoryID := c.Param(paramCategoryID)

	category, apiError := miapi.GetCategoryFromApi(categoryID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, category)
}