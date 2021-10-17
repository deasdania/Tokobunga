package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
)

type IProductSizeMysql interface {
	GetProductSizeByName(name string) (*models.ProductSize, error)
	GetProductSizeById(id int) (*models.ProductSize, error)
	GetAllProductSize(orderby string) ([]*models.ProductSize, error)
	CreateProductSize(name *models.FormName) error
	UpdateProductSizeName(id string, name string) error
	DeleteProductSizeById(id string) error
	Cobaajadulu() error
}
