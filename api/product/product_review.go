package product

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities/token"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"strconv"
	"strings"
)

type FormInputProductReview struct {
	ProductId string `json:"product_id" binding:"required"`
	Rating    string `json:"rating" binding:"required"`
}

// CreateProductReview godoc
// @Summary CreateProductReview Private
// @Description Product Review hanya diizinkan setiap user 1 review, ke setiap Product, rating dibatasi dari 1 hingga 5
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData FormInputProductReview true "set 'product_id' and 'rating' to create a new product review"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/create/productreview [post]
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

// UpdateProductReview godoc
// @Summary UpdateProductReview Private
// @Description Product Review hanya di update oleh user yang berkaitan dengan review
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData FormInputProductReview true "set 'product_id' and 'rating' to create a new product review"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/update/productreview [post]
func (a Product) UpdateProductReview(c *gin.Context) {
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

	response := a.ProductUsecase.UpdateProductReview(prodRev, email)
	c.JSON(response.Status, response)
	return

}
