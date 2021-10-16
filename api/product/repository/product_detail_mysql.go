package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
)

type IProductDetailMysql interface {
	GetProductDetailById(id string) (*models.ProductDetail, error)
	GetProductDetailByProductId(product_id string) (*models.ProductDetail, error)
	DeleteProductDetailById(id string) error
	CreateProductDetail(detail *models.ProductDetail) error
}
