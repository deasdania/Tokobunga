package repository

import "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"

type IProductMysql interface {
	GetProductByName(name *string) (*models.Product, error)
	GetProductById(id string) (*models.Product, error)
	GetAllProduct(orderby string) ([]*models.Product, error)

	CreateProductReq(name *models.ProductReq) error
	CreateProductDetail(detail *models.ProductDetail) error
	CreateProduct(name *models.Product) error

	UpdateProductName(id string, name string) error
	UpdateProductCategoryId(id string, category_id string) error
	DeleteProductById(id string) error
}
