package repository

import "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"

type IProductCategoryMysql interface {
	GetProductCategoryById(id string) (*models.ProductCategory, error)
	GetProductCategoryByName(name string) (*models.ProductCategory, error)
	GetAllProductCategory(orderby string) ([]*models.ProductCategory, error)
	CreateProductCategory(form *models.FormName) error
	UpdateProductName(id string, name string) error
	DeleteProductById(id string) error
}
