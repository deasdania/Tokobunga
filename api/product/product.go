package product

import (
	accountusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	productusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strings"
)

type Product struct {
	ProductUsecase productusecase.IProductUsecase
	AccountUsecase accountusecase.IAccountUsecase
}

func (a Product) Product(r *gin.RouterGroup) {
	r.POST(utilities.CREATE_PRODUCT, a.CreateProduct)
	r.POST(utilities.CREATE_PRODUCT_CATEGORY, a.CreateProductCategory)

	r.POST(utilities.CREATE_PRODUCT_REVIEW, a.CreateProductReview)
	// r.PUT(utilities.UPDATE_PRODUCT_REVIEW, a.UpdateProductReview)
	// r.DEL(utilities.DELETE_PRODUCT_REVIEW, a.DeleteProductReview)
}

func (a Product) CreateProduct(c *gin.Context) {
	var productReq models.ProductReq
	err := c.ShouldBindJSON(&productReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := a.ProductUsecase.CreateProduct(&productReq)
	c.JSON(response.Status, response)
	return
}

func (a Product) UpdateProduct(c *gin.Context) {
	var productReq models.ProductReq
	err := c.ShouldBindJSON(&productReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := a.ProductUsecase.UpdateProduct(&productReq)
	c.JSON(response.Status, response)
	return
}
