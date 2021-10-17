package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
)

type IProductGalleryMysql interface {
	GetProductGalleryById(id *int) (*models.ProductGallery, error)
	GetProductGalleryByProductId(product_id string) (*models.ProductGallery, error)
	DeleteProductGalleryById(id string) error
	CreateProductGallery(gallery *models.ProductGallery) error
}
