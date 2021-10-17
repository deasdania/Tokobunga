package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
)

type IProductUsecase interface {
	CreateProduct(productReq *models.ProductReq) *response.Response
	// CreateProduct(product models.Product) *response.Response
	// UpdateProduct(product models.Product) *response.Response
	CreateProductCategory(form_name models.FormName) *response.Response
	// UpdateProductCategory(prodcat models.ProductCategory) *response.Response
	// CreateProductSize(form_name models.FormName) *response.Response
	// UpdateProductSize(prodsize models.ProductSize) *response.Response
	// CreateProductDetail(detail models.FormProductDetail) *response.Response
}
