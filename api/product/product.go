package product

import (
	accountusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	productusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
)

type Product struct {
	ProductUsecase productusecase.IProductUsecase
	AccountUsecase accountusecase.IAccountUsecase
	AuthUsecase    usecase.IAuthUsecase
}

func (a Product) Product(r *gin.RouterGroup) {
	r.POST(utilities.CREATE_PRODUCT, a.CreateProduct)
	// r.PUT(utilities.UPDATE_PRODUCT, a.UpdateProduct)

	r.POST(utilities.CREATE_PRODUCT_CATEGORY, a.CreateProductCategory)
	// r.PUT(utilities.UPDATE_PRODUCT_CATEGORY, a.UpdateProductCategory)

	// r.POST(utilities.CREATE_PRODUCT_DETAIL, a.CreateProductDetail)
}

func (a Product) CreateProduct(c *gin.Context) {
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
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})

}

// func (a Product) UpdateProduct(c *gin.Context) {
// 	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
// 	if errA != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": errA.Error(),
// 		})
// 		return
// 	}
// 	fmt.Println(metadata)
// 	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
// 	if isAdmin {
// 		id := c.PostForm("product_id")
// 		name := c.PostForm("name")
// 		category_id := c.PostForm("category_id")

// 		if id == "" || name == "" || category_id == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": "product_id, name, and category_id cannot be Empty",
// 			})
// 			return
// 		}

// 		tint, _ := strconv.Atoi(category_id)
// 		tid, _ := strconv.Atoi(id)

// 		form_product := models.Product{
// 			Id:         tid,
// 			Name:       name,
// 			CategoryId: tint,
// 		}
// 		response := a.ProductUsecase.UpdateProduct(form_product)
// 		c.JSON(response.Status, response)
// 		return
// 	}
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"message": "you are not allowed",
// 	})
// }
