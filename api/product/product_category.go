package product

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a Product) CreateProductCategory(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errA.Error(),
		})
		return
	}
	fmt.Println(metadata)
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
	if isAdmin {
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
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})

}

func (a Product) UpdateProductCategory(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errA.Error(),
		})
		return
	}
	fmt.Println(metadata)
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
	if isAdmin {
		name := c.PostForm("name")
		category_id := c.PostForm("category_id")

		if name == "" || category_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "name and category_id cannot be Empty",
			})
			return
		}

		tint, _ := strconv.Atoi(category_id)

		form_product := models.ProductCategory{
			Name: name,
			Id:   tint,
		}
		response := a.ProductUsecase.UpdateProductCategory(form_product)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}
