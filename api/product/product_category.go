package product

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
)

// CreateProductCategory godoc
// @Summary CreateProductCategory Private
// @Description Category tidak bisa dibuat dengan nama yang sama
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData models.FormName true "set 'name' to create a new category"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/create/productcategory [post]
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

// UpdateProductCategory godoc
// @Summary UpdateProductCategory Private
// @Description Update dilakukan menurut id yang dilempar dan name yang berbeda
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData models.FormName true "set 'name' and 'category_id' to update Product Category"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/update/productcategory [put]
func (a Product) UpdateProductCategory(c *gin.Context) {
	category_id := c.PostForm("category_id")
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name cannot be Empty",
		})
		return
	}
	response := a.ProductUsecase.UpdateProductCategory(category_id, name)
	c.JSON(response.Status, response)
	return
}
