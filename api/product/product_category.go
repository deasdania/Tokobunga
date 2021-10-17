package product

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
)

func (a Product) CreateProductCategory(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name cannot be Empty",
		})
		return
	}
	form_product := models.FormName{
		Name: name,
	}
	response := a.ProductUsecase.CreateProductCategory(form_product)
	c.JSON(response.Status, response)
	return
}
