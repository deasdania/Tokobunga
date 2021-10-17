package repository

import "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"

type IProductCategoryMysql interface {
	GetProductCategoryById(id *int) (*models.ProductCategory, error)
	GetProductCategoryByName(name string) (*models.ProductCategory, error)
	GetProductCategoryByProductId(id *int) (*models.ProductCategory, error)
	GetAllProductCategory(orderby string) ([]*models.ProductCategory, error)
	CreateProductCategory(form *models.FormName) error
	UpdateProductCategory(id int, name string) error
	DeleteProductById(id string) error
}
