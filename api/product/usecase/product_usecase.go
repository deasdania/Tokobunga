package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
)

type IProductUsecase interface {
	GetProduct(productid string, order string) *response.Response
	CreateProduct(productReq *models.ProductReq) *response.Response
	GetProductReview(user_id string, product_id string) *response.Response

	CreateProductReview(prodRev models.ProductReview, email string) *response.Response
	UpdateProductReview(prodRev models.ProductReview, email string) *response.Response

	CreateProductCategory(form_name models.FormName) *response.Response
	UpdateProductCategory(category_id string, name string) *response.Response

	UpdateProduct(productReq *models.ProductReq) *response.Response
}
