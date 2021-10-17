package product

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities/token"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (a Product) CreateProductReview(c *gin.Context) {
	tokenstring := token.ExtractToken(c)
	metadata, _ := token.ExtractTokenMetadata(tokenstring)
	email := metadata.Email
	product_id := c.PostForm("product_id")
	rating := c.PostForm("rating")

	prodtrim := strings.Trim(product_id, " ")
	ratingtring := strings.Trim(rating, " ")

	intprodtrim, _ := strconv.Atoi(prodtrim)
	intratingtring, _ := strconv.Atoi(ratingtring)

	prodRev := models.ProductReview{
		Rating:    intratingtring,
		ProductId: intprodtrim,
	}

	response := a.ProductUsecase.CreateProductReview(prodRev, email)
	c.JSON(response.Status, response)
	return

}

// func (a Product) UpdateProductReview(c *gin.Context) {
// 	var productReq models.ProductReq
// 	err := c.ShouldBindJSON(&productReq)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	response := a.ProductUsecase.UpdateProductReview(&productReq)
// 	c.JSON(response.Status, response)
// 	return
// }

// func (a Product) DeleteProductReview(c *gin.Context) {
// 	var productReq models.ProductReq
// 	err := c.ShouldBindJSON(&productReq)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	response := a.ProductUsecase.DeleteProductReview(&productReq)
// 	c.JSON(response.Status, response)
// 	return
// }
